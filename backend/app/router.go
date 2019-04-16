/**
* @author [Hilton Li]
* @email [liminghilton@gmail.com]
* @create date 2019-04-16 15:03:43
* @modify date 2019-04-16 15:03:43
* @desc [description]
*/

package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lmlala/Devent/backend/module"
	"net/http"
)

func (r *ginApp)addEvent() {
	v1 := r.client.Group("/v1")
	{		
		v1.GET("/incident", func(c *gin.Context) {
			// 注意:在前后端分离过程中，需要注意跨域问题，因此需要设置请求头
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			r, err := module.ReadIncident()
			
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": 500,
					"error":  "Get Incident Error!",
				})
			} else {
				c.JSON(http.StatusOK, r)
			}
			
		})

		v1.GET("/deployment", func(c *gin.Context) {
			// 注意:在前后端分离过程中，需要注意跨域问题，因此需要设置请求头
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			r, err := module.ReadDepolyment()
			
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": 500,
					"error":  "Get Deployment Error!",
				})
			} else {
				c.JSON(http.StatusOK, r)
			}
			
		})

		v1.GET("/list", func(c *gin.Context) {
			// 注意:在前后端分离过程中，需要注意跨域问题，因此需要设置请求头
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			r, err := module.ListEvent()
			
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": 500,
					"error":  "Get All Event Error!",
				})
			} else {
				c.JSON(http.StatusOK, r)
			}
			
		})
		
	}
}

//定义默认路由
func (r *ginApp)addDefault() {
	r.client.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
}