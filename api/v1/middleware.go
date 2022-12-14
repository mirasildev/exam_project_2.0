package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirasildev/exam_project_2.0/pkg/utils"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationPayloadKey = "authorization_payload"
)

func (h *handlerV1) AuthMiddleware(c *gin.Context) {
	accessToken := c.GetHeader(authorizationHeaderKey)

	if len(accessToken) == 0 {
		err := errors.New("authorization header is not provided")
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	payload, err := utils.VerifyToken(h.cfg, accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	c.Set(authorizationPayloadKey, payload)
	c.Next()
}

func (h *handlerV1) GetAuthPayload(c *gin.Context) (*utils.Payload, error) {
	i, exists := c.Get(authorizationPayloadKey)
	if !exists {
		return nil, errors.New("")
	}

	payload, ok := i.(*utils.Payload)
	if !ok {
		return nil, errors.New("unknown user")
	}
	return payload, nil
}
