package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muathendirangu/hex/internal/core/domain"
	"github.com/muathendirangu/hex/internal/core/services"
)

type HTTPHandler struct {
	messageSVC services.MessageService
}

func NewHTTPHandler(messageSVC services.MessageService) *HTTPHandler {
	return &HTTPHandler{messageSVC}
}

func (h *HTTPHandler) SaveMessage(ctx *gin.Context) {
	var message domain.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"Error": err.Error(),
			},
		)
		return
	}
	err := h.messageSVC.SaveMessage(message)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"Error": err.Error(),
			},
		)
		return
	}
	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"Message": "message was successfully created",
		},
	)
}

func (h *HTTPHandler) ReadMessage(ctx *gin.Context) {
	id := ctx.Param("id")
	message, err := h.messageSVC.ReadMessage(id)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"Error": err.Error(),
			},
		)
		return
	}
	ctx.JSON(http.StatusOK, message)
}

func (h *HTTPHandler) ReadMessages(ctx *gin.Context) {

	messages, err := h.messageSVC.ReadMessages()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, messages)
}
