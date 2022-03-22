/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:30:17
 * @Description:
 */
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

/**
 * @description: 成功返回
 * @param {*gin.Context} c
 * @param {interface{}} data
 * @return {*}
 */
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Res{Code: 200, Message: "success", Data: data})
}

/**
 * @description: 错误返回
 * @param {*gin.Context} c
 * @param {error} err
 * @return {*}
 */
func Error(c *gin.Context, err error) {
	if e, ok := err.(*Err); ok {
		c.JSON(http.StatusOK, Res{Code: e.Code, Message: e.Message})
	} else {
		c.JSON(http.StatusOK, Res{Code: -1, Message: err.Error()})
	}
}

/**
 * @description: 自定义返回
 * @param {*gin.Context} c
 * @param {int} code
 * @param {string} message
 * @return {*}
 */
func ErrorE(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Res{Code: code, Message: message})
}
