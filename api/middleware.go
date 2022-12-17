package api

import (
	"fmt"
	"github.com/techschool/simplebank/token"
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
	"strings"
)

const(
	authorizationHeaderKey = "autorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "autorization_payloadKey"
)

func authMiddleWare(tokenMaker token.Maker) gin.HandlerFunc{
	return func(ctx *gin.Context){
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if(len(authorizationHeader) == 0 ){
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorResponse(err))
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2{
			err := errors.New("authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorResponse(err))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer{
			err := fmt.Errorf("Unsupported authorization tyoe %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if(err != nil){
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
		
	}
}