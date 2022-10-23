package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	engine "github.com/rpturbina/assigment-go-3/config/gin"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})

	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger(),
	)

	startTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "server up and running",
			"start_time": startTime,
		})
	})

	ginEngine.GetGin().LoadHTMLFiles("template/index.html")

	type DataPoint struct {
		Water     int       `json:"water"`
		Wind      int       `json:"wind"`
		Timestamp time.Time `json:"timestamp"`
	}

	data := []DataPoint{}
	ginEngine.GetGin().GET("/status", func(c *gin.Context) {
		newData := DataPoint{
			Water:     rand.Intn(100),
			Wind:      rand.Intn(100),
			Timestamp: time.Now().Local(),
		}

		waterStatus := "AMAN"

		if newData.Water >= 6 && newData.Water <= 8 {
			waterStatus = "SIAGA"
		}
		if newData.Water > 8 {
			waterStatus = "BAHAYA"
		}

		windStatus := "AMAN"
		if newData.Wind >= 7 && newData.Wind <= 15 {
			windStatus = "SIAGA"
		}
		if newData.Wind > 15 {
			windStatus = "BAHAYA"
		}

		data = append(data, newData)

		file, _ := json.MarshalIndent(data, "", " ")
		_ = os.WriteFile("data.json", file, 0644)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"water_status": waterStatus,
			"wind_status":  windStatus,
			"data":         data,
		})
	})

	ginEngine.Serve()
}
