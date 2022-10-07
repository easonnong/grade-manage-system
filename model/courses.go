package model

import (
	"github.com/jinzhu/gorm"
)

//////////////////////////////
//1. 课程编号 2. 教师编号
//3. 课程名称 4. 课程学分
//////////////////////////////
type Courses struct {
	gorm.Model
	Teachers  Teachers `gorm:"ForeignKey:TeacherID;AssociationForeignKey:ID"`
	TeacherID uint
	Name      string `gorm:"type:varchar(10);not null"`
	Credit    string `gorm:"type:varchar(5);not null"`
}
