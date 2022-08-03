package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	db "github.com/ebaudet/go-fizz-buzz/db/sqlc"
	"github.com/gin-gonic/gin"
)

type statisticResponse struct {
	Request json.RawMessage `json:"request"`
	Hints   int64           `json:"hints"`
}

func newStatisticResponse(data db.FizzbuzzStatistic) statisticResponse {
	return statisticResponse{
		Request: data.Request,
		Hints:   data.Count,
	}
}

func (server *Server) statistics(ctx *gin.Context) {
	fizz_stat, err := server.store.GetMostUsedRequest(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, statisticResponse{})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newStatisticResponse(fizz_stat)
	ctx.JSON(200, rsp)
}
