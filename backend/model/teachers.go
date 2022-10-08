package model

import (
	"github.com/jinzhu/gorm"
)

///////////////////////////////
//1. 教师编号 2. 姓名 3. 学院
///////////////////////////////
type Teachers struct {
	gorm.Model
	TeacherId uint
	Password  string `gorm:"size:255;not null"`
	Name      string `gorm:"type:varchar(5);not null"`
	College   string `gorm:"type:varchar(10);not null"`
}
