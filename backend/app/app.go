/**
 * @author [Hilton Li]
 * @email [liminghilton@gmail.com]
 * @create date 2019-04-16 15:03:43
 * @modify date 2019-04-16 15:03:43
 * @desc [description]
 */

package app

import (
	"github.com/gin-gonic/gin"
	
)

type ginApp struct {
	client	*gin.Engine
}

func Run() {
	var app ginApp
	app.client = gin.Default()
	app.addDefault()
	app.addEvent()

	app.client.Run(":8000")
}

