package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"

	"oracle/interfaces"
	"oracle/models"
)

type DataController struct {
	VerificationService interfaces.VerificationService
	dataStore           map[string][]models.Data
}

func (d *DataController) storeData(ctx *gin.Context) {
	var data models.DataDTO

	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
		return
	}

	hash := sha256.Sum256(data.PublicKey)
	messageID := hex.EncodeToString(hash[:])

	if _, exists := d.dataStore[messageID]; !exists {
		d.dataStore[messageID] = []models.Data{}
	}

	newMessage := models.Data{ID: messageID, Content: data.EncryptedMessage}
	d.dataStore[messageID] = append(d.dataStore[messageID], newMessage)

	ctx.JSON(http.StatusOK, gin.H{"Sate": "Stored"})
}

func (d *DataController) retrieveMessage(ctx *gin.Context) {
	var dataRequest models.DataRequest
	if err := ctx.BindJSON(dataRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
		return
	}

	messageID, err := d.VerificationService.VerifyCertificateChallange(dataRequest.Certificate, dataRequest.SignedChallenge)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Denied": "Bad Credentails"})
		return
	}
	messages, exists := d.dataStore[*messageID]
	if !exists {
		ctx.JSON(http.StatusOK, []models.Data{})
		return
	}

	delete(d.dataStore, *messageID)
	ctx.JSON(http.StatusOK, messages)
}

func (d *DataController) Init(r *gin.RouterGroup) {
	d.dataStore = make(map[string][]models.Data)
	controller := r.Group("data")
	controller.POST("/", d.storeData)
	controller.GET("/new", d.retrieveMessage)
}
