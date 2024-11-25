package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"oracle/interfaces"
)

type IdentityController struct {
	VerificationService interfaces.VerificationService
	IdentityService     interfaces.IdentityVerificationService
}

func (i *IdentityController) self(ctx *gin.Context) {
	var num int
	if err := ctx.Bind(&num); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
	}

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
