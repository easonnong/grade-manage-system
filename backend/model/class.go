package model

import (
	"github.com/jinzhu/gorm"
)

////////////////////////////////////////////
//1. 学院 2. 专业 3. 年级(2020级 2021级...)
////////////////////////////////////////////
type Class struct {
	gorm.Model
	College string `gorm:"type:varchar(10);not null"`
	Major   string `gorm:"type:varchar(10);not null"`
	Grade   uint
}
