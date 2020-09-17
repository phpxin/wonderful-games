/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/7/8 11:36 上午
*/

package tools

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func GetCaller(skip int) string {
	_,file,line,_ := runtime.Caller(skip+1)
	file = file[strings.LastIndex(file, "/")+1:]
	return fmt.Sprintf("%s:%d", file, line)
}

func GetMillisecond(t time.Time) int64 {
	return t.UnixNano()/int64(time.Millisecond)
}