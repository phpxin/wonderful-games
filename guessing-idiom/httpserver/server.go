/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/9/17 4:34 下午
*/

package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wonderful-games/guessing-idiom/conf"
)

func Run() {
	e := gin.Default()

	e.LoadHTMLGlob("html/*")
	e.GET("/test", func(c *gin.Context) {
		//log.Debug("", "123")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	game := new(GameController)
	e.Handle(http.MethodGet, "/game/stage/get", game.Getstage)

	gm := new(GamemasterController)
	e.Handle(http.MethodGet, "/gm/test", gm.Test)
	e.Handle(http.MethodGet, "/gm/games/create", gm.Creategames)

	host := conf.GetConfigString("app", "host")
	port := conf.GetConfigString("app", "port")
	err := e.Run(host+":"+port)

	if err != nil {
		panic(err)
	}
}
