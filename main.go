package main

import (
	"github.com/easonnong/grade-manage/model"
	"github.com/easonnong/grade-manage/routes"
)

func main() {
	// 引用数据库
	model.InitDb()
	// 引入路由组件
	routes.InitRouter()

}
