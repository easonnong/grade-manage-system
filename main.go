package main

import (
	"os"

	"github.com/easonnong/grade-manage-system/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	//读取配置信息
	InitConfig()
	//连接mysql
	common.InitDB()
	//获取数据库
	db := common.GetDB()
	defer db.Close()
	//获取路由
	router := gin.Default()
	router = CollectRoutes(router)
}

//从配置文件中读取
func InitConfig() {
	//获取当前工作目录
	workDir, err := os.Getwd()
	if err != nil {
		panic("获取当前工作目录失败")
	}
	//设置读取文件的类型
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		panic("读取文件失败")
	}
}
