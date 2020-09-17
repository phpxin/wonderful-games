package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"wonderful-games/guessing-idiom/core"
	"wonderful-games/guessing-idiom/models"
	"wonderful-games/guessing-idiom/tools/rpchelper"
)

type GameController struct {
	core.BaseController
}

func (c *GameController) Getstage(ctx *gin.Context) {
	level, ok := rpchelper.RequestParameterInt(ctx, "level")
	if !ok || level <= 0 {
		c.JsonError(ctx, core.ApiErrMsg, "invalid param level")
		return
	}

	limit, ok := rpchelper.RequestParameterInt(ctx, "limit")
	if !ok || limit <= 0 {
		c.JsonError(ctx, core.ApiErrMsg, "invalid param limit")
		return
	}

	stages, err := models.GetStages(int32(level), int32(limit))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//c.apiError(errcode.NotFound, "stage not found")
			c.JsonError(ctx, core.ApiErrMsg, "stage not found")
		}
		//c.apiError(errcode.InnerError, "")
		c.JsonError(ctx, core.ApiErrMsg, "inner error")
		return
	}

	c.JsonSuccess(ctx, map[string]interface{}{
		"list":  stages,
		"level": level,
		"limit": limit,
	})
}
