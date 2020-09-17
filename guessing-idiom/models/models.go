/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/9/17 4:34 下午
*/

package models


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"wonderful-games/guessing-idiom/conf"
)

var (
	// 全局连接句柄
	dbs = make(map[string]*gorm.DB)

)

func InitModel() {

	// db_main
	err := registerDB("mysql")
	if err!=nil {
		panic(err)
	}

}

// 获取主库实例
func GetDbInst() *gorm.DB {
	return dbs["mysql"]
}

func registerDB(sectionName string) error {
	//"user:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local"
	//&parseTime=True&loc=Local 解析 time 类型
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.GetConfigString(sectionName, "user"),
		conf.GetConfigString(sectionName, "password"),
		conf.GetConfigString(sectionName, "host"),
		conf.GetConfigString(sectionName, "port"),
		conf.GetConfigString(sectionName, "name"))

	var err error
	dbs[sectionName], err = gorm.Open("mysql", connstr)
	if err!=nil {
		return err
	}

	// todo 设置连接池，实际设置参考正式上线使用量
	//db.DB().SetMaxIdleConns()

	return nil
}
