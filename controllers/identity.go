package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IdentityController struct {
}

func (i *IdentityController) self(ctx *gin.Context) {
	var num int
	if err := ctx.Bind(&num); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
	}

}

func (i *IdentityController) Init(r *gin.RouterGroup) {
	controller := r.Group("identity")

	controller.POST("/self", i.self)
}
