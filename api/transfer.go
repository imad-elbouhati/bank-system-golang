package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/imad-elbouhati/bank/db/sqlc"
	"github.com/imad-elbouhati/bank/token"
)



type transferRequest struct {
	FromAccountID    int64 `json:"from_account_id" binding:"required,min=1"`
	ToAccountID int64 `json:"to_account_id" binding:"required,min=1"`
	Amount int `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	
	var req transferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}
	

	fromAccount, valid := server.validateAccount(ctx,req.FromAccountID,req.Currency) 

	authPayload := ctx.MustGet(authorizationPayloadkey).(*token.Payload)

	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account doesn't belong to authenticated user")
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	if !valid{
		return
	}

	_, valid = server.validateAccount(ctx,req.ToAccountID,req.Currency) 

	if !valid {
		err := errors.New("to account isn't valid")
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	arg := db.TransferTxParams {
		FromAccountID: req.FromAccountID,
		ToAccountID: req.ToAccountID,
		Amount: int64(req.Amount),
	}

	result,err := server.store.TranserTx(ctx,arg)
	if(err != nil) {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,result)
}

func (server *Server) validateAccount(ctx *gin.Context, accoutID int64, currency string) (db.Account, bool) {

	account, err := server.store.GetAccount(ctx,accoutID)

	if(err != nil) {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return account,false
		}

		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return account,false
	}

	if account.Currency != currency {
		err = fmt.Errorf("account [%d] currency mismatch: %s vs %s", accoutID,currency,account.Currency)
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return account, false
	}

	return account,true
}