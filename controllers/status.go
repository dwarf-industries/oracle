package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type StatusController struct {
}

func (s *StatusController) memory(ctx *gin.Context) {

	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data struct {
		total     string
		available string
	}
	total := v.Total / 1024 / 1024
	available := v.Available / 1024 / 1024
	data.total = strconv.Itoa(int(total))
	data.available = strconv.Itoa(int(available))

	ctx.JSON(http.StatusOK, data)
}

func (s *StatusController) cpu(ctx *gin.Context) {
	cpuData, err := cpu.Percent(time.Duration(0), false)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
		return
	}

	ctx.JSON(http.StatusOK, cpuData)
}

func (s *StatusController) resources(ctx *gin.Context) {
	cpuData, err := cpu.Percent(time.Duration(0), false)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Bad Request"})
		return
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data struct {
		total     string
		available string
		cpu       string
	}
	total := v.Total / 1024 / 1024
	available := v.Available / 1024 / 1024
	data.total = strconv.Itoa(int(total))
	data.available = strconv.Itoa(int(available))

	for _, c := range cpuData {
		str := strconv.FormatFloat(c, 'f', 2, 64)
		data.cpu += str
	}

	ctx.JSON(http.StatusOK, data)
}

func (s *StatusController) Init(r *gin.RouterGroup) {
	controller := r.Group("status")
	controller.GET("memory", s.memory)
	controller.GET("cpu", s.cpu)
	controller.GET("all", s.resources)
}
