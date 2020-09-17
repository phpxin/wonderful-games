package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"wonderful-games/guessing-idiom/conf"
	"wonderful-games/guessing-idiom/tools/log"
)

type IdiomIntro struct {
	Pron string
	Word string
	Source string
	Expl string
}

type IdiomSlot struct {
	Idiom string
	Slot int32
}

var (
	//
	// 字 -》 词 倒排索引，一对多
	InvertedIndex = make(map[string][]string)
	Idioms = make(map[string]*IdiomIntro)
)

func InitInvertedIndex() {
	content,err := ioutil.ReadFile(conf.GetConfigString("app", "data_path")+"/wordInfo.json")
	if err!=nil {
		panic(err)
		return
	}

	err = json.Unmarshal(content, &Idioms)
	if err!=nil {
		panic(err)
		return
	}

	for k,_ := range Idioms {
		words := strings.Split(k,"")
		//fmt.Println(words)

		for _,w := range words  {
			_,ok := InvertedIndex[w]
			if !ok {
				InvertedIndex[w] = make([]string, 0)
			}
			InvertedIndex[w] = append(InvertedIndex[w], k)
		}
	}

	log.Debug("", "init idioms done")
}

func PrintInvertedIndex() {
	for k,v := range InvertedIndex {
		fmt.Println(k, " : ", v)
	}
}

