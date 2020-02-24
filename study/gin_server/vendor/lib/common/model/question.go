/*
 * author:        liujun
 * created:       2018-04-23 22:39
 * last modified: 2018-04-23 22:39
 * filename:      question.go
 * description:   题库相关的公用模型
 */
package model

type QuestionCategory struct {
	Id    int64  `gorm:"column:id"`
	Level int    `gorm:"column:level"`
	Cname string `gorm:"column:c_name"`
}

func (self QuestionCategory) TableName() string {
	return "question_category"
}

type QuestionTag struct {
	Id  int64  `gorm:"column:id"`
	Tag string `gorm:"column:tag"`
}

func (self QuestionTag) TableName() string {
	return "question_tag"
}

type Question struct {
	Id             int64  `gorm:"column:id"`
	QTitle         string `gorm:"column:qtitle"`
	QPic           string `gorm:"column:qpic"`
	Aa             string `gorm:"column:a_a"`
	Ap             string `gorm:"column:a_pic"`
	Ba             string `gorm:"column:b_a"`
	Bp             string `gorm:"column:b_pic"`
	Ca             string `gorm:"column:c_a"`
	Cp             string `gorm:"column:c_pic"`
	Da             string `gorm:"column:d_a"`
	Dp             string `gorm:"column:d_pic"`
	Answer         string `gorm:"column:answer"`
	Star           int    `gorm:"column:star"`
	IsBase         int    `gorm:"column:is_base"`
	Atime          int    `gorm:"column:a_time"`
	Category       int    `gorm:"column:category"`
	SecondCategory int
	Tag            int `gorm:"column:tag"`
}

func (self Question) TableName() string {
	return "question"
}

type QuestionInfo struct {
	Id     int64  `json:"id"`
	QTitle string `json:"title"`
	QPic   string `json:"pic"`
	Aa     string `json:"aa"`
	Ap     string `json:"ap"`
	Ba     string `json:"ba"`
	Bp     string `json:"bp"`
	Ca     string `json:"ca"`
	Cp     string `json:"cp"`
	Da     string `json:"da"`
	Dp     string `json:"dp"`
	Answer string `json:"answer"`
	Tag    string `json:"tag"`
	Atime  int    `json:"a_time"`
}

const (
	SettleBefore = 1
	SettleAfter  = 2
)

type RoomInfo struct {
	RoomId     string `json:"room_id"`
	Uid        string `json:"uid"`
	SettleType int64  `json:"settle_type"`
}

type CreateRoomReq struct {
	Quess    []QuestionInfo `json:"quess"`
	Eid      int            `json:"eid"`
	Uid      string         `json:"uid"`
	RoomType int            `json:"room_type"`
	Category string         `json:"category"`
	Gold     int64          `json:"gold"`
}
type CreateRoomRes struct {
	RoomId   string `json:"room_id"`
	RoomAddr string `json:"room_addr"`
}

type RoomInitInfo struct {
	FirstTag  string   `json:"first_tag"`
	SecondTag string   `json:"second_tag"`
	ThreeTags []string `json:"three_tags"`
	FourTags  []string `json:"four_tags"`
}

type PassRoomCondition struct {
	Score   int64 `json:"score"`
	Correct int64 `json:"correct"`
}

type PassRoomAward struct {
	Gold int64 `json:"gold"`
}

type RoomExternalInfo struct {
	Rtag    string            `json:"rtag"`
	Rii     RoomInitInfo      `json:"rii"`
	Prc     PassRoomCondition `json:"prc"`
	Pra     PassRoomAward     `json:"pra"`
	Count   int64             `json:"count"`
	Quess   []QuestionInfo    `json:"quess"`
	ResType int               `json:"res_type"`
}
