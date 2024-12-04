package controllers

import (
	"encoding/hex"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"oracle/di"
	"oracle/interfaces"
	"oracle/models"
)

type DataSocketController struct {
	PaymentProcessor    interfaces.PaymentProcessor
	VerificationService interfaces.VerificationService
	dataStore           map[string][]models.Data
	activeSessions      map[string]string
	connMutex           sync.Mutex
	upgrader            websocket.Upgrader
}

func (d *DataSocketController) HandleWebSocket(ctx *gin.Context) {
	conn, err := d.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to upgrade to WebSocket"})
		return
	}
	defer conn.Close()

	for {
		var message map[string]interface{}
		err := conn.ReadJSON(&message)
		if err != nil {
			conn.WriteJSON(gin.H{"Error": "Failed to read message"})
			return
		}

		action, ok := message["action"].(string)
		if !ok {
			conn.WriteJSON(gin.H{"Error": "Invalid action"})
			continue
		}

		authenticated := d.isAuthenticated(message["sessionToken"])
		if !authenticated && action != "authenticate" {
			conn.WriteJSON(gin.H{"Error": "Invalid handshake"})
			return
		}

		switch action {
		case "authenticate":
			d.authenticate(conn, message)

		case "pop-request":
			d.issueTransfer(conn, message)
		case "store":
			transferPayment := d.verifyTransferPayment(conn, message)
			if !transferPayment {
				conn.WriteJSON(gin.H{"Error": "Invalid Payment Details"})
				return
			}

			d.storeDataSocket(conn, message)
		case "retrieve":
			d.retrieveMessageSocket(conn, message)
		default:
			conn.WriteJSON(gin.H{"Error": "Unknown action"})
		}
	}
}

func (d *DataSocketController) issueTransfer(conn *websocket.Conn, data map[string]interface{}) {
	size, ok := data["size"].(float64)
	if !ok {
		conn.WriteJSON(gin.H{"Error": "Invalid size"})
		return
	}
	popDetails := d.PaymentProcessor.GeneratePaymentRequest(int(size))
	conn.WriteJSON(popDetails)
}

func (d *DataSocketController) verifyTransferPayment(conn *websocket.Conn, data map[string]interface{}) bool {
	payment := data["pop"].(string)

	decoded, err := hex.DecodeString(payment)
	if err != nil {
		return false
	}

	paymentId := string(decoded)
	allow, err := di.PaymentProcessor().VerifyPayment(paymentId)
	if !allow || err != nil {
		conn.WriteJSON(gin.H{"Error": "Rejected, can't verify payment"})
		return false
	}

	return true
}

func (d *DataSocketController) isAuthenticated(authentication interface{}) bool {
	_, ok := d.activeSessions[authentication.(string)]
	return ok
}

func (d *DataSocketController) authenticate(conn *websocket.Conn, handshake map[string]interface{}) {

	address, addressOk := handshake["address"].(string)
	signedChallengeData, signOk := handshake["signedChallenge"].(string)
	sessionToken, tokenOk := handshake["sessionToken"].(string)

	if !signOk || !tokenOk || !addressOk {
		conn.WriteJSON(gin.H{"Error": "Missing required fields"})
		return
	}

	signatureBytes, _ := hex.DecodeString(signedChallengeData)

	challenge := d.VerificationService.GetChallenge([]byte(address))
	verifiedMessageID, err := d.VerificationService.VerifyChallange(challenge, signatureBytes, address)
	if err != nil {
		conn.WriteJSON(gin.H{"Error": "Challenge verification failed"})
		return
	}

	d.connMutex.Lock()
	d.activeSessions[sessionToken] = *verifiedMessageID
	d.connMutex.Unlock()

	conn.WriteJSON(gin.H{"State": "Authenticated"})
}

func (d *DataSocketController) storeDataSocket(conn *websocket.Conn, message map[string]interface{}) {
	sessionToken, ok := message["sessionToken"].(string)
	if !ok {
		conn.WriteJSON(gin.H{"Error": "Missing session token"})
		return
	}

	d.connMutex.Lock()
	messageID, exists := d.activeSessions[sessionToken]
	d.connMutex.Unlock()

	if !exists {
		conn.WriteJSON(gin.H{"Error": "Invalid session token"})
		return
	}

	encryptedMessage, ok := message["encryptedMessage"].(string)
	if !ok {
		conn.WriteJSON(gin.H{"Error": "Invalid message format"})
		return
	}

	if _, exists := d.dataStore[messageID]; !exists {
		d.dataStore[messageID] = []models.Data{}
	}

	newMessage := models.Data{ID: messageID, Content: []byte(encryptedMessage)}
	d.dataStore[messageID] = append(d.dataStore[messageID], newMessage)

	conn.WriteJSON(gin.H{"State": "Stored"})
}

func (d *DataSocketController) retrieveMessageSocket(conn *websocket.Conn, message map[string]interface{}) {
	sessionToken, ok := message["sessionToken"].(string)
	if !ok {
		conn.WriteJSON(gin.H{"Error": "Missing session token"})
		return
	}

	d.connMutex.Lock()
	messageID, exists := d.activeSessions[sessionToken]
	d.connMutex.Unlock()

	if !exists {
		conn.WriteJSON(gin.H{"Error": "Invalid session token"})
		return
	}

	messages, exists := d.dataStore[messageID]
	if !exists {
		conn.WriteJSON([]models.Data{})
		return
	}

	delete(d.dataStore, messageID)
	conn.WriteJSON(messages)
}

func (d *DataSocketController) Init(r *gin.RouterGroup) {
	d.activeSessions = map[string]string{}
	d.dataStore = make(map[string][]models.Data)

	controller := r.Group("rlt")
	controller.GET("/ws", func(ctx *gin.Context) {
		d.HandleWebSocket(ctx)
	})
}
