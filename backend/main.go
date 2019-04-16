/**
 * @author [Hilton Li]
 * @email [liminghilton@gmail.com]
 * @create date 2019-04-15 14:27:58
 * @modify date 2019-04-15 14:27:58
 * @desc [description]
 */

package main

import (
    "fmt"
	"github.com/gin-gonic/gin"
	"github.com/lmlala/Devent/backend/events"
    "math/rand"
	"net/http"
)

// type Event struct{
// 	StartTime time.Time	`"json:start_time"`
// 	EndTime	time.Time	`"json:end_time"`
// 	EventType	string	`"json:event_type"`
// 	EventTag	string	`"json:event_tag"`
// 	Description	string	`"json:description"`
// }

func HelloPage(c *gin.Context) {
	
    c.JSON(http.StatusOK, gin.H{
        "message": "welcome to bgops,please visit https://xxbandy.github.io!",
    })
}



func main() {
    r := gin.Default()
    v1 := r.Group("/v1")
    {
        v1.GET("/hello", HelloPage)
        v1.GET("/hello/:name", func(c *gin.Context) {
            name := c.Param("name")
            c.String(http.StatusOK, "Hello %s", name)
        })

        v1.GET("/line", func(c *gin.Context) {
            // 注意:在前后端分离过程中，需要注意跨域问题，因此需要设置请求头
            c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
            legendData := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
            xAxisData := []int{120, 240, rand.Intn(500), rand.Intn(500), 150, 230, 180}
            c.JSON(200, gin.H{
                "legend_data": legendData,
                "xAxis_data":  xAxisData,
            })
		})
		
		v1.GET("/incident", func(c *gin.Context) {
            // 注意:在前后端分离过程中，需要注意跨域问题，因此需要设置请求头
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			r, err := Events.ReadIncident()
			
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
			r, err := Events.ReadDepolyment()
			
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
			r, err := Events.ListEvent()
			
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
    //定义默认路由
    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "status": 404,
            "error":  "404, page not exists!",
        })
    })
    r.Run(":8000")
}