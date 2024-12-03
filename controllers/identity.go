package controllers

import (
	"encoding/hex"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"oracle/interfaces"
)

type IdentityController struct {
	VerificationService interfaces.VerificationService
	IdentityService     interfaces.IdentityVerificationService
	WalletService       interfaces.WalletService
}

func (i *IdentityController) self(ctx *gin.Context) {
	var challenge []byte
	if err := ctx.Bind(&challenge); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
	}

	signature, err := i.WalletService.SignMessage(challenge)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Failed to generate signature"})
		return
	}
	sig := hex.EncodeToString(signature)
	ctx.JSON(http.StatusOK, sig)
}

func (i *IdentityController) challange(ctx *gin.Context) {
	certificate, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Denied": "Failed to read certificate"})
		return
	}

	challange, err := i.VerificationService.GenerateChallenge(certificate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
		return
	}
	sig := hex.EncodeToString(challange)
	ctx.JSON(http.StatusOK, sig)
}

func (i *IdentityController) Init(r *gin.RouterGroup) {
	controller := r.Group("identity")

	controller.POST("/self", i.self)
	controller.POST("/certificate", i.challange)
}
