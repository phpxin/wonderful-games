/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/9/18 4:13 下午
*/

package tester

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"wonderful-games/guessing-idiom/conf"
	"wonderful-games/guessing-idiom/core"
	"wonderful-games/guessing-idiom/models"
	"wonderful-games/guessing-idiom/tools/log"
	"wonderful-games/guessing-idiom/tools/logger"
)

func TestInit(t *testing.T) {
	err := conf.ParseConfigINI("../conf/conf.ini")
	if err!=nil {
		fmt.Println("err : parse config failed", err.Error())
		os.Exit(1)
	}

	logPath := conf.GetConfigString("app", "log_path")

	logger.InitLogger(logPath, "20060102") // 日志
	el,err := conf.GetConfigInt("log", "level")
	if err!=nil {
		el = 0
	}
	log.SetLogErrorLevel(int(el))

	// 初始化数据库
	models.InitModel()

	core.InitInvertedIndex()
}

func TestCreateGame(t *testing.T) {
	for k,_ := range core.Idioms {

		g := core.NewGame(k)
		g.FindResult(true)
		g.PrintMatrix()
		stage := g.GetResult()
		stageStr,_ := json.Marshal(stage)
		fmt.Println(string(stageStr))

		break
	}
}