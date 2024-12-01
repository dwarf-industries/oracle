package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	"oracle/models"
)

type StatusController struct {
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

	var data models.Resources

	total := v.Total / 1024 / 1024
	available := v.Available / 1024 / 1024
	data.Memory = strings.Join([]string{
		strconv.Itoa(int(total)),
		"MB",
	}, "")
	data.Available = strings.Join([]string{
		strconv.Itoa(int(available)),
		"MB",
	}, "")

	for _, c := range cpuData {
		str := strconv.FormatFloat(c, 'f', 2, 64)
		data.Cpu += str
	}

	ctx.JSON(http.StatusOK, data)
}

func (s *StatusController) Init(r *gin.RouterGroup) {
	controller := r.Group("status")
	controller.GET("", s.resources)
}
