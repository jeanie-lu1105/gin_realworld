package handler

import (
	"net/http"
	"time"

	"example.com/gin_realworld/logger"
	"example.com/gin_realworld/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func AddWebsocket(r *gin.Engine) {
	wsGroup := r.Group("/api/ws")
	wsGroup.GET("", ws)
}

func ws(ctx *gin.Context) {
	log := logger.New(ctx)
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.WithError(err).Error("websocket upgrade failed")
		ctx.Status(http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			err := conn.WriteJSON(map[string]interface{}{
				"type": "heartbeat",
			})
			if err != nil {
				return
			}
		}
	}()

	for {
		reqMsg := make(map[string]interface{})
		err := conn.ReadJSON(&reqMsg)
		if err != nil {
			log.WithError(err).Errorf("read json failed")
			return
		}

		log.Infof("read msg: %v\n", utils.JsonMarshal(reqMsg))

		if reqMsg["exit"] != nil {
			return
		}

		err = conn.WriteJSON(reqMsg)
		if err != nil {
			log.WithError(err).Errorf("write json failed")
			return
		}
	}
}
