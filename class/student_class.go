package class

import (
	"Homework_system/dao"
	"Homework_system/db"
	"Homework_system/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//学生加入班级
func Student_class(context *gin.Context) {
	var rs dao.Rep
	var stu dao.Class
	//通过邀请码加入，并在班级表中number+1，学生ID加入membersID
	invitation := context.PostForm("invitation")
	row := db.Db.QueryRow("select * from homework.class where invitation = ?;", invitation)
	err := row.Scan(&stu.ClassName, &stu.ClassID, &stu.CreateID, &stu.Invitation, &stu.Number, &stu.StartTime, &stu.MembersID)
	if err != nil {
		rs.Code = 500
		rs.Msg = "没有此邀请码请重新确认"
		context.JSON(http.StatusOK, rs)
		return
	}
	//获取学生信息
	captial := global.Get_Session(context)
	id := strconv.Itoa(captial.Id)
	if dao.Join_class(invitation, id) {
		rs.Code = 200
		rs.Msg = "加入班级成功"
		context.JSON(http.StatusOK, rs)
		return
	}
	rs.Code = 500
	rs.Msg = "加入班级失败"
	context.JSON(http.StatusOK, rs)
}

//班级学生管理
func Student_management(context *gin.Context) {
	//根据邀请码拿到学生ID名单。
	invitation := context.PostForm("invitation")
	//获取当前班级学生ID
	var membersID string
	db.Db.QueryRow("select membersID from homework.class where invitation = ?;", invitation).Scan(&membersID)
	fmt.Println(membersID)
}
