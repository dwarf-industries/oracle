package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"oracle/models"
)

type NodesController struct {
	oracles []models.Oracle
}

func (n *NodesController) all(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, n.oracles)
}

func (n *NodesController) near(ctx *gin.Context) {
	var nodeLatency []struct {
		ip      string
		latency time.Duration
	}

	for _, o := range n.oracles {
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
