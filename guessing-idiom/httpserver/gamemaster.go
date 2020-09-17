package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"wonderful-games/guessing-idiom/core"
	"wonderful-games/guessing-idiom/models"
	"wonderful-games/guessing-idiom/tools/log"
)

type GamemasterController struct {
	core.BaseController
}

func (c *GamemasterController) Test(ctx *gin.Context){
	c.JsonSuccess(ctx, map[string]interface{}{
		"msg": "success",
	})
}

func (c *GamemasterController) Creategames(ctx *gin.Context) {

	go func() {
		i:=0
		for k,_ := range core.Idioms {

			g := core.NewGame(k)
			g.FindResult(true)
			//g.PrintMatrix()
			stage := g.GetResult()
			stageStr,_ := json.Marshal(stage)
			fmt.Println(string(stageStr))

			// save
			models.SaveStage(int32(i)+1, string(stageStr))

			log.Debug("", "create new stage done %d", i)

			i++
			if i>=20 {
				break
			}
		}

	}()

	c.JsonSuccess(ctx, map[string]interface{}{
		"msg": "success",
	})
}
