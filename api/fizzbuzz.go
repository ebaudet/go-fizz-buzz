package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ebaudet/go-fizz-buzz/util"
	"github.com/gin-gonic/gin"
)

type fizzBuzzRequest struct {
	Int1  int    `json:"int1" binding:"min=1"`
	Int2  int    `json:"int2" binding:"min=1"`
	Limit int    `json:"limit" binding:"min=1,max=1000000"`
	Str1  string `json:"str1" binding:"required"`
	Str2  string `json:"str2" binding:"required"`
}

func (server *Server) fizzBuzz(ctx *gin.Context) {
	var request fizzBuzzRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	f := util.FizzParams{
		Int1: request.Int1,
		Int2: request.Int2,
		Str1: request.Str1,
		Str2: request.Str2,
	}

	var arrayString []string
	for i := 1; i <= request.Limit; i++ {
		res, err := util.FizzBuzz(i, f)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		arrayString = append(arrayString, res)
	}

	output := strings.Join(arrayString, ",") + "."

	json_request, err := json.Marshal(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	_, err = server.store.IncrementRequest(ctx, json_request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(200, output)
}
