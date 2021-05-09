package model

import (
	"time"
)

type MoodMixin struct {
	Id      int    `gorm:"id" json:"id"`
	Content string `gorm:"content" json:"content"`
	Uid     int    `gorm:"uid" json:"uid"`
	IsDel   int    `gorm:"is_del" json:"is_del"`
	IsRead  int    `gorm:"is_read" json:"is_read"`
}

type Mood struct {
	MoodMixin
	Time int `gorm:"time" json:"time"`
}
type MoodTime struct {
	MoodMixin
	DateTime string `gorm:"time" json:"datetime"`
}

func (m *Mood) TableName() string {
	return "gc_mood"
}
func (mt *MoodTime) TableName() string {
	return "gc_mood"
}
func AddMood(uid int, content string) bool {
	var mood Mood

	mood.Uid = uid
	mood.Content = content
	mood.Time = int(time.Now().Unix())
	affected := db.Save(&mood)
	return affected.RowsAffected != 0
}
func OneMood() MoodTime {
	var mood Mood
	var moodtime MoodTime
	//读取记录
	db.Where("is_read=0").First(&mood)
	moodtime.MoodMixin = mood.MoodMixin
	moodtime.DateTime = time.Unix(int64(mood.Time), 0).Format("2006.1.2")
	return moodtime
}

func DelMood(id int) bool {
	affected := db.Model(&Mood{}).Where("id", id).Update("is_del", 1)
	return affected.RowsAffected != 0
}
