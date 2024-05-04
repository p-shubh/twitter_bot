package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var 

type webhookHandler struct{}

// This handler will handle the events from twitter. You can print it to console, forward it to yourself,
// reply using twitter api. possibilities are endless.
func (wh webhookHandler) Handler(c *gin.Context) {
	webhookBody, _ := c.GetRawData()
	fmt.Println(string(webhookBody))
}

func Router() {

	router := gin.Default()

	wh := webhookHandler{}

	router.GET("/webhookPath", func(c *gin.Context) {
		Bot.HandleCRCResponse(c.Writer, c.Request)
	})

	router.POST("/webhookPath", func(c *gin.Context) {
		wh.Handler(c)
	})

	go func() {
		fmt.Println(http.ListenAndServe(":8080", router))
	}()

	select {}
}
