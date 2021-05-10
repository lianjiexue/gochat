package models

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
func GetFUllMoods(uid int, page int) []MoodTime {
	var moods []Mood
	var newmoods []MoodTime

	offset := (page - 1) * 10
	db.Where("uid", uid).Select("id", "content", "time").Limit(10).Offset(offset).Order("id DESC").Find(&moods)
	var temp MoodTime
	for _, value := range moods {
		temp.MoodMixin = value.MoodMixin
		temp.DateTime = time.Unix(int64(value.Time), 0).Format("2006.1.2 15:04")
		newmoods = append(newmoods, temp)
	}
	return newmoods
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

	if mood.Id != 0 {
		db.Model(&Mood{}).Where("id", mood.Id).Update("is_read", 1)
	}
	moodtime.MoodMixin = mood.MoodMixin
	moodtime.DateTime = time.Unix(int64(mood.Time), 0).Format("2006.1.2")
	return moodtime
}

func DelMood(id int) bool {
	affected := db.Model(&Mood{}).Where("id", id).Update("is_del", 1)
	return affected.RowsAffected != 0
}
