package main

//xiugai
import (
	"Homework_system/db"
	"Homework_system/global"
	"Homework_system/tool"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var configPath string = "./config.json"
	tool.LoadConfigFile(configPath)
	db.Dbinit()                           //连接和初始化数据库
	r := gin.Default()                    //一个gin的实例
	r.LoadHTMLGlob("htmlfile/*")          //gin加载HTML模板
	r.Static("/resources", "./resources") //设置静态资源路径
	routerinit(r)                         //路由

	r.Run(":" + global.Gconfig.WebsitPort)

}
