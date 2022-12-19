package api

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"database/sql"
)


type renewAccessTokenRequest struct {
	RefreshToken    string `json:"refresh_token" binding:"required"`
}
 
type renewAccessTokenResponce struct{
	AccessToken string `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context){
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
		return 
	}

	refreshPayload,err :=  server.tokenMaker.VerifyToken(req.RefreshToken)
	if(err != nil){
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if(err != nil){
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return 
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	if(session.IsBlocked){
		err := fmt.Errorf("Blocked session !!")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err) )
		return
	}

	if(session.Username != refreshPayload.Username){
		err := fmt.Errorf("incorrect session user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err) )
		return
	}

	if(session.RefreshToken != req.RefreshToken){
		err := fmt.Errorf("missmatch of token !!")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err) )
		return
	}

	if(time.Now().After(session.ExpiresAt)){
		err:= fmt.Errorf("the token has expired")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		server.config.AccessTokenDuration,
	)
	if(err != nil){
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp:= renewAccessTokenResponce{
		AccessToken: accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK,rsp) 
}