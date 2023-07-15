package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/imad-elbouhati/bank/db/sqlc"
	"github.com/imad-elbouhati/bank/token"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency`
}

func (server *Server) createAccount(ctx *gin.Context) {
	
	var req CreateAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}
	
	authPayload := ctx.MustGet(authorizationPayloadkey).(*token.Payload)

	arg := db.CreateAccountParams {
		Owner: authPayload.Username,
		Currency: req.Currency,
		Balance: 0,
	}

	account,err := server.store.CreateAccount(ctx,arg)
	if(err != nil) {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,account)
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	
	var req GetAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	account,err := server.store.GetAccount(ctx,req.ID)

	if(err != nil) {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadkey).(*token.Payload)

	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to authenticated user")
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK,account)
}



type ListAccountRequest struct {
	PageID  int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=5,max=10"`
}


func (server *Server) listAccount(ctx *gin.Context) {
	
	var req ListAccountRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListAccountsParams {
		Limit: int32(req.PageSize),
		Offset: int32(req.PageID),
	}

	accounts,err := server.store.ListAccounts(ctx,arg)

	if(err != nil) {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK,accounts)
}
