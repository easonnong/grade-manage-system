package model

import (
	"github.com/jinzhu/gorm"
)

//////////////////////////////////////
//1. 学号 2. 姓名 3. 性别 4. 个人学分
//////////////////////////////////////
type Students struct {
	gorm.Model
	StudentId uint
	Password  string `gorm:"size:255;not null"`
	Name      string `gorm:"type:varchar(5);not null"`
	Sex       string `gorm:"type:varchar(1);not null"`
	Credit    uint
}
