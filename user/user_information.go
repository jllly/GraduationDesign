//用户信息
package user

import (
	"Homework_system/class"
	"Homework_system/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取名字,显示
func Obtain_name(context *gin.Context) {
	var rs class.SearchUserList
	//获取cookie
	myCookie, _ := context.Cookie("auth") //string格式
	var name string
	err := db.Db.QueryRow("select userName from homework.user where phonenumber=?", myCookie).Scan(&name)
	if err != nil {
		rs.Code = 500
		rs.Msg = "数据库错误"
		context.JSON(http.StatusOK, rs)
		return
	}
	rs.Code = 200
	rs.Msg = name
	context.JSON(http.StatusOK, rs)
}
