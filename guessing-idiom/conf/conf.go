package conf

import (
	"gopkg.in/ini.v1"
)

var (
	configFp *ini.File


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


	return nil
}
