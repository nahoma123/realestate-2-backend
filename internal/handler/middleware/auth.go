package middleware

import (
	"crypto/rsa"
	"net/http"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/module"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware interface {
	Authentication(admin ...bool) gin.HandlerFunc
	BindUser(id string) gin.HandlerFunc
}

type authMiddleware struct {
	auth         module.UserModule
	ssoPublicKey *rsa.PublicKey
}

func InitAuthMiddleware(
	auth module.UserModule, ssoPublicKey *rsa.PublicKey) AuthMiddleware {
	return &authMiddleware{
		auth,
		ssoPublicKey,
	}
}

func (a *authMiddleware) Authentication(admin ...bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			Err := errors.ErrInvalidToken.New("Unauthorized")
			ctx.Error(Err)
			ctx.Abort()
			return
		}

		tokenString := authHeader[len(bearer):]
		valid, claims, _ := a.auth.VerifyToken(jwt.SigningMethodHS256, tokenString)
		if !valid {
			Err := errors.ErrAuthError.New("Unauthorized")
			ctx.Error(Err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userId, ok := (*claims)["user_id"].(string)
		if !ok {
			// userId is not present in the claims or not in the correct format
		}

		if len(admin) > 0 && admin[0] && (*claims)["role"] != "ADMIN_ROLE" {
			Err := errors.ErrAuthError.New("Unauthorized. Only ADMIN_ROLE allowed.")
			ctx.Error(Err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// userStatus, err := a.auth.GetUserStatus(ctx.Request.Context(), userId)
		// if err != nil {
		// 	ctx.Error(err)
		// 	ctx.Abort()
		// 	return
		// }

		// if userStatus != constant.Active {
		// 	Err := errors.ErrAuthError.Wrap(nil, "Your account has been deactivated, Please activate your account.")
		// 	ctx.Error(Err)
		// 	ctx.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }
		ctx.Set("x-user-id", userId)
		ctx.Next()
	}
}

func (a *authMiddleware) BindUser(id string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("x-user-id", id)
		ctx.Next()
	}
}
