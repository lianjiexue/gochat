package model

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
func AddMood(ctx *gin.Context) {
	var mood Mood
	uid, err := strconv.Atoi(ctx.PostForm("uid"))
	if err != nil {
		return
	}
	mood.Uid = uid
	mood.Content = ctx.PostForm("content")
	mood.Time = int(time.Now().Unix())
	db.Save(&mood)
	ctx.JSON(200, gin.H{
		"code": 200,
	})
}
func OneMood(ctx *gin.Context) {
	var mood Mood
	var moodtime MoodTime

	//读取记录
	db.Where("is_read=0").First(&mood)
	moodtime.MoodMixin = mood.MoodMixin
	moodtime.DateTime = time.Unix(int64(mood.Time), 0).Format("2006.1.2")
	log.Println(moodtime)
	if mood.Id != 0 {
		//将记录设置为已读
		db.Model(&Mood{}).Where("id", mood.Id).Update("is_read", 1)

		var user OneUser
		db.Where("id", mood.Uid).First(&user)
		ctx.JSON(200, gin.H{
			"code": 200,
			"data": gin.H{
				"user": user,
				"mood": moodtime,
			},
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}

}

func DelMood(ctx *gin.Context) {
	id := ctx.PostForm("id")
	db.Model(&Mood{}).Where("id", id).Update("is_del", 1)
	ctx.JSON(200, gin.H{
		"code": 200,
	})
}
