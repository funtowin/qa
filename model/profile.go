package model

// Profile 用户简介
type Profile struct {
	GORMBase
	Nickname  string `json:"nickname" gorm:"type:varchar(500)"`
	Gender    int    `json:"gender" gorm:"type:int;DEFAULT:0;"`
	Desc      string `json:"desc" gorm:"type:varchar(1000)"`
	AvatarUrl string `json:"avatarUrl" gorm:"type:varchar(1000)"`
	UserID    uint   `json:"userId"`
}
