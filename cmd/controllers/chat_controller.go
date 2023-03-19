package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func InitializeWS(ctx *gin.Context) {

	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)

	ws.WriteMessage(100, []byte("hey little stranger"))

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"cause": "Error on initialize websocket",
		})
		return
	}

	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		fmt.Println(string(message))

		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"cause":   "Error on read message",
			})
			break
		}

		if string(message) == "ping" {
			message = []byte("pong")
		}

		err = ws.WriteMessage(mt, message)

		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"cause":   "Error on send message",
			})
			break
		}

	}

}
