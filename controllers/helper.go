package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ErrorResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func ReturnJsonStruct(c *gin.Context, code int, genericStruct interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	_ = json.NewEncoder(c.Writer).Encode(genericStruct)
}
