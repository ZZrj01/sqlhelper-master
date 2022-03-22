/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:30:54
 * @Description:
 */
package source

import (
	"fmt"
	"net/http"

	"com.fanyunai.sqlHelper/source/api"
	"com.fanyunai.sqlHelper/source/config"
	"github.com/gin-gonic/gin"
)

/**
 * @description: 创建路由
 * @param {*}
 * @return {*}
 */
func CreateRouter() {
	// gin 模式
	gin.SetMode(gin.DebugMode)
	// 创建路由
	router := gin.Default()
	// 解决跨域
	router.Use(Cors())
	// 初始化
	router.POST("/initialize", api.Initialize)
	// 测试连接
	router.POST("/testConnect", api.TestConnect)
	// 添加连接
	router.POST("/addconnect", api.AddConnect)
	// 获取所有的
	router.POST("/getconnect", api.GetConnect)
	// 获取所有表
	router.POST("/gettable", api.GetTables)
	// 到出模型
	router.POST("/exportmodel", api.ExportModel)
	// 删除连接
	router.POST("/deleteconnect", api.DeleteConnect)
	// 监听端口
	router.Run(fmt.Sprintf(":%d", config.Conf.Server.Port))
}

/**
 * @description: 解决跨域
 * @param {*}
 * @return {*}
 */
//解决跨域问题的中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println(c.Request.Host)
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))             // 允许访问域名                               // 允许访问域名
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,token")        // 允许自定义头信息，如Content-Type、Token、Authorization 必填
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE") // 允许方法
		c.Header("Access-Control-Allow-Credentials", "true")                                // 响应头是否可以把请求头暴露给页面
		c.Header("Access-Control-Expose-Headers", "*")                                      // 响应暴露头部
		c.Header("Access-Control-Max-Age", "172800")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		// 处理请求
		c.Next()
	}
}
