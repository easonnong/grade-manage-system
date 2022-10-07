package model

/////////////////////////////////
//1. 课程编号 2. 学生学号 3. 分数
/////////////////////////////////
type Grade struct {
	Courses   Courses `gorm:"ForeignKey:CourseID;AssociationForeignKey:ID;primaryKey"`
	CourseID  uint
	Students  Students `gorm:"ForeignKey:StudentID;AssociationForeignKey:ID;primaryKey"`
	StudentID uint
	Score     uint
}
