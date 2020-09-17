package conf

import (
	"gopkg.in/ini.v1"
)

var (
	configFp *ini.File

	//ServerID int64
	//ServerName string
	//
	//Env = EnvDev
)

func GetConfigInt(section,name string) (int64,error) {
	return configFp.Section(section).Key(name).Int64()
}

func GetConfigString(section,name string) string {
	return configFp.Section(section).Key(name).String()
}

func ParseConfigINI(cpath string) (err error) {

	configFp,err = ini.Load(cpath)
	if err!=nil {
		return err
	}

	//rand.Seed(time.Now().Unix())
	//ServerID = rand.Int63()

	//hostName,e := os.Hostname()
	//
	//if e!=nil {
	//	hostName = fmt.Sprintf("no-hostname-%d", ServerID)
	//}

	//ServerName = fmt.Sprintf("%s:%s",
	//	hostName,
	//	GetConfigString("ws", "port"))
	//
	//if appEnv := GetConfigString("app", "environment"); appEnv!="" {
	//	Env = appEnv
	//}

	return nil
}
