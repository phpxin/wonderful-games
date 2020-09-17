/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/9/17 4:31 下午
*/

package main


import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"wonderful-games/guessing-idiom/conf"
	"wonderful-games/guessing-idiom/core"
	"wonderful-games/guessing-idiom/httpserver"
	"wonderful-games/guessing-idiom/models"
	"wonderful-games/guessing-idiom/tools/log"
	"wonderful-games/guessing-idiom/tools/logger"
)

var (
	sigs = make(chan os.Signal)
	done = make(chan bool)
)

func main() {

	// runtime.NumCPU()

	if len(os.Args) < 2 {
		fmt.Println("usage : ./pushserver {path of config}")
		os.Exit(1)
	}
	confpath := os.Args[1]

	// 初始化配置

	err := conf.ParseConfigINI(confpath)
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

	// 开启 RPC 服务
	go func() {
		defer log.PrintPanicStackError()

		// 启动程序
		httpserver.Run()
	}()

	// serve
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // ctrl+c, kill, kill -2,
	go sigAwaiter()

	<-done
}

// kill , ctrl+c 可以， kill -9 不行
func sigAwaiter() {
	sig := <-sigs
	fmt.Println(fmt.Sprintf("recv signal %s, process will exit 3 second later", sig.String()))

	// 进程退出需要处理的逻辑
	logger.DestroyLogger() // 刷日志

	// time.Sleep(3*time.Second) // 暂停 3 秒钟，等待后台任务结束 todo 延迟关闭会有不同步问题，上线看实际情况定

	fmt.Println("process exit.")

	done<-true
}
