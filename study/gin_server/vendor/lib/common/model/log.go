package model

// UserOper 用户行为记录表
type UserOper struct {
	ID       int64  `gorm:"column:id"`
	UID      string `gorm:"column:uid"`
	RoomType int64  `gorm:"column:room_type"`
	RoomID   string `gorm:"column:room_id"`
	Oper     int32  `gorm:"column:oper"`
	Str1     string `gorm:"column:str1"`
	Str2     string `gorm:"column:str2"`
	Str3     string `gorm:"column:str3"`
	Int1     int64  `gorm:"column:int_1"`
	Int2     int64  `gorm:"column:int_2"`
	Int3     int64  `gorm:"column:int_3"`
}

// TableName user_oper
func (uo UserOper) TableName() string {
	return "user_oper"
}
