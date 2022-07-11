package models

import "github.com/gin-gonic/gin"

type TaskResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    gin.H  `json:"data"`
}
