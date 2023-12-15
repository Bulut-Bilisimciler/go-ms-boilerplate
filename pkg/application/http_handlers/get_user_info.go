package http_handlers

import (
	"errors"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/domain"
	"github.com/gin-gonic/gin"
)

// HandleGetUserInformation godoc
// @Summary get user info from jwt
// @Schemes
// @Description get user info from jwt
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} api.RespondJson "user info"
// @Failure 401 {object} api.RespondJson "invalid jwt token, unauthorized"
// @Failure 500 {object} api.RespondJson "internal server error"
// @Router /me [get]
func (svc *HTTPHandlerService) HandleGetUserInformation(c *gin.Context) (int, interface{}, error) {
	// get user from context after jwt middleware
	jwt := c.MustGet("jwt")
	jwtClaims, ok := jwt.(*domain.JWTClaims)
	if !ok {
		return http.StatusUnauthorized, nil, errors.New("jwt user is invalid")
	}

	// return user info
	return http.StatusOK, jwtClaims, nil
}

// example usage of redis: try to get redis keys *
// keys, err := svc.rdb.Keys(c.Request.Context(), "*").Result()
// if err != nil {
// 	return http.StatusInternalServerError, nil, err
// }
