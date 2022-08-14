package http

import (
	"github.com/avinashb98/blaast/domain/blaast"
	"github.com/avinashb98/blaast/errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type BlaastHandler interface {
	CreateBlaast(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type handler struct {
	s blaast.Service
}

func NewBlaast(s blaast.Service) BlaastHandler {
	return handler{
		s: s,
	}
}

func (h handler) CreateBlaast(ctx *gin.Context) {
	input, err := h.parseCreateBlaast(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	newBlaast, err := h.s.CreateBlaast(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, newBlaast)
}

func (h handler) parseCreateBlaast(ctx *gin.Context) (*blaast.CreateBlaastInput, error) {
	var payload blaast.CreateBlaastInput
	requestBodyErr := ctx.BindJSON(&payload)
	if requestBodyErr != nil {
		log.Printf("invalid create blaast input format: %s", requestBodyErr.Error())
		return nil, errors.ValidationError{Message: "invalid input format"}
	}

	if payload.ReceiverID == "" {
		log.Printf("empty receiver id")
		return nil, errors.ValidationError{Message: "empty receiver id"}
	}
	if payload.SenderID == "" {
		log.Printf("empty sender id")
		return nil, errors.ValidationError{Message: "empty sender id"}
	}

	if payload.Text == "" {
		log.Printf("empty text")
		return nil, errors.ValidationError{Message: "empty text"}
	}

	return &payload, nil
}

func (h handler) GetByID(ctx *gin.Context) {
	blaastId, err := h.parseGetBlaastByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	foundBlaast, err := h.s.GetByID(blaastId)
	if err != nil {
		if _, ok := err.(*errors.NotFoundError); ok {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, foundBlaast)
}

func (h handler) parseGetBlaastByID(ctx *gin.Context) (string, error) {
	blaastId := ctx.Query("blaast-id")
	if blaastId == "" {
		log.Printf("empty blaast id")
		return "", errors.ValidationError{Message: "empty blaast id"}
	}

	return blaastId, nil
}
