package controllers

import (
	"encoding/hex"
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
	var certificate []byte
	if err := ctx.Bind(&certificate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Denied": "Bad Request"})
		return
	}

	challange, err := i.VerificationService.GenerateChallenge(certificate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
		return
	}

	ctx.JSON(http.StatusOK, challange)
}

func (i *IdentityController) Init(r *gin.RouterGroup) {
	controller := r.Group("identity")

	controller.POST("/self", i.self)
	controller.POST("/certificate", i.challange)
}
