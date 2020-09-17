package models

import (
	"time"
)

const (
	STAGE_STATUS_NORMAL = 1
	STAGE_STATUS_OFF    = 2
)

type Stage struct {
	Id        int32
	Level     int32
	Contents  string
	Status    int32
	CreatedAt int64
}

func (stage *Stage) TableName() string {
	return "stages"
}

func GetStages(level int32, limit int32) ([]*Stage, error) {
	o := GetDbInst()

	stages := make([]*Stage, 0)
	err := o.Raw("select * from stages where level>=? order by level asc limit ?", level, limit).Find(&stages).Error

	if err != nil {
		return nil, err
	}

	return stages, nil
}

func GetStage(level int32) (*Stage, error) {
	o := GetDbInst()
	stage := new(Stage)
	o.Where("level = ?", level)
	err := o.First(stage).Error

	if err != nil {
		return nil, err
	}

	return stage, nil
}

func SaveStage(level int32, contents string) {
	stage := new(Stage)
	stage.Level = level
	stage.Contents = contents
	stage.Status = STAGE_STATUS_NORMAL
	stage.CreatedAt = time.Now().Unix()

	o := GetDbInst()
	o.Save(stage)
}
