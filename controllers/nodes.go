package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"oracle/di"
)

type NodesController struct {
}

func (n *NodesController) all(ctx *gin.Context) {
	oracles, err := di.RegisterService().Oracles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Error"})
		return
	}

	ctx.JSON(http.StatusOK, oracles)
}

func (n *NodesController) near(ctx *gin.Context) {
	var nodeLatency []struct {
		ip      string
		latency time.Duration
	}
	oracles, err := di.RegisterService().Oracles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Error"})
		return
	}

	for _, o := range oracles {
		start := time.Now()
		resp, err := http.Get(o.Ip)
		if err != nil {
			fmt.Print("Node appears to be offline, didn't notify the network")
		}
		defer resp.Body.Close()

		latency := time.Since(start)
		if latency.Abs().Milliseconds() < 60 {
			nodeLatency = append(nodeLatency, struct {
				ip      string
				latency time.Duration
			}{
				ip:      o.Ip,
				latency: latency,
			})
		}
	}

	ctx.JSON(http.StatusOK, nodeLatency)
}

func (n *NodesController) Init(r *gin.RouterGroup) {
	nodesController := r.Group("nodes")
	nodesController.GET("/", n.all)
	nodesController.GET("/near", n.near)
}
