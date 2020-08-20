//发布作业
package homework

import (
	"Homework_system/dao"
	"Homework_system/db"
	"Homework_system/global"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Classlist struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data [](dao.Class) `json:"data"`
}

//布置作业当前班级
func Choose_class(context *gin.Context) {

	var rs Classlist
	//拿到当前班级id
	classID := context.Query("id")

	//初始化Data切片
	dlist := make([]dao.Class, 0, 0)
	var class dao.Class
	err := db.Db.QueryRow("select * from homework.class where classID=?", classID).Scan(&class.ClassName, &class.ClassID, &class.CreateID, &class.Invitation, &class.Number, &class.StartTime, &class.MembersID)
	if err != nil {
		rs.Code = 500
		rs.Msg = "信息出错！"
		context.JSON(http.StatusOK, rs)
		return
	}
	dlist = append(dlist, class)
	rs.Code = 200
	rs.Msg = "成功"
	rs.Data = dlist
	context.JSON(http.StatusOK, rs)
}

//搜索班级下列表
func Search_Class_list(context *gin.Context) {
	var rs Classlist
	captial := global.Get_Session(context) //获取session信息

	dlist := make([]dao.Class, 0, 0)
	rows, _ := db.Db.Query("SELECT * FROM homework.class WHERE createID = ?  ORDER BY startTime; ", captial.Id)
	defer rows.Close()
	if rows == nil {
		return
	}
	for rows.Next() {
		//填充数据
		var row dao.Class
		rows.Scan(&row.ClassName, &row.ClassID, &row.CreateID, &row.Invitation, &row.Number, &row.StartTime, &row.MembersID)
		//每行数据
		//放入结果集
		dlist = append(dlist, row)

	}
	rs.Code = 200
	rs.Msg = "成功"
	rs.Data = dlist
	context.JSON(http.StatusOK, rs)
}

//搜索班级
func Search_Class(context *gin.Context) {
	data := context.Query("keyword")

	//SELECT * FROM [user] WHERE u_name LIKE '%三%'
	var rs Classlist
	captial := global.Get_Session(context) //获取session信息

	dlist := make([]dao.Class, 0, 0)

	rows, _ := db.Db.Query("SELECT * FROM homework.class WHERE createID = ? and className like CONCAT('%',?,'%') ORDER BY startTime; ", captial.Id, data)
	defer rows.Close()
	if rows == nil {
		return
	}
	for rows.Next() {
		//填充数据
		var row dao.Class
		rows.Scan(&row.ClassName, &row.ClassID, &row.CreateID, &row.Invitation, &row.Number, &row.StartTime, &row.MembersID)
		//每行数据
		//放入结果集
		dlist = append(dlist, row)
	}
	rs.Code = 200
	rs.Msg = "成功"
	rs.Data = dlist
	context.JSON(http.StatusOK, rs)
}


//完成作业的布置
//根据提交过来的时间，作业名称，完成作业的布置，并通知当前班级的所有学生
func Complete_assignment(context *gin.Context) {


	var rs dao.Rep
	rs.Msg ="作业布置失败"
	rs.Code = 500

	var StartTimestr, EndTimestr, HomeworkName, classID,topicID string
	StartTimestr = context.PostForm("StartTime")       //开始时间
	EndTimestr = context.PostForm("EndTime")           //结束时间
	HomeworkName = context.PostForm("HomeworkName") //作业名称
	classID = context.PostForm("classID")           //要布置作业的班级ID的集合
	topicID = context.PostForm("topicID")         //布置作业的题目
	//captial := global.Get_Session(context)          //获取session信息

	StartTime, err := global.Timeformat(StartTimestr)
	if err !=nil{
		context.JSON(http.StatusOK, rs)
		return
	}
	EndTime, err := global.Timeformat(EndTimestr)
	if err !=nil{
		context.JSON(http.StatusOK, rs)
		return
	}
	creator := global.Get_Session(context)

	//拆分班级
	class:=strings.Split(classID,",")
	for _, v := range class {
		_, e := db.Db.Exec("insert into homework.teacher_homework (StartTime,EndTime,HomeworkName,classID,topicID,creator) values (?,?,?,?,?,?);", StartTime, EndTime, HomeworkName, v, topicID,creator.Id) //插入信息
		if e != nil {
			context.JSON(http.StatusOK, rs)
			return
		}

		var membersID string
		db.Db.QueryRow("select membersID from homework.class where classID=?",v).Scan(&membersID)   //v班级ID
		membersIDs:=strings.Split(membersID,",")
		for _,d:=range membersIDs{
			a:=0
			uid, _:=strconv.Atoi(d)
			hid, _:=strconv.Atoi(v)
			_, e := db.Db.Exec("insert into homework.student_homework (userID,homeworkId,state,topic) values (?,?,?,?);",uid,hid,a,topicID) //插入学生id:d;作业id：HomeworkID;状态a(默认情况下0)
			if e != nil {
				log.Panicln("user insert error", e.Error())
				context.JSON(http.StatusOK, rs)
				return
			}
		}
	}
	//var HomeworkID string
	//db.Db.QueryRow("select @@identity").Scan(&HomeworkID)
	rs.Code = 200
	rs.Msg = "成功"
	context.JSON(http.StatusOK, rs)
	//作业布置成功，给作业班级的所有学生发送通知
	//关于微信发通知可以使用消息队列，但当前只要存储数据
	//插入成功，拿到作业ID，
	//go add_homework(HomeworkID,classID)
	//		？班级表要加状态吗，1发布   2批改中   3已完成
}
//起一个goroutine，做学生班级表添加数据
//func addStu_homework(HomeworkID,classlist string){
//	class:=strings.Split(classlist,",")
//	for _,v:=range class{
//		var membersID string
//		//v班级，通过班级jiao
//		db.Db.QueryRow("select membersID from homework.class where classID=?",v).Scan(&membersID)   //v班级ID
//		membersIDs:=strings.Split(membersID,",")
//		for _,d:=range membersIDs{
//			a:='0'
//			_, e := db.Db.Exec("insert into homework.student_homework (homeworkId,userID,state) values (?,?,?);",d,HomeworkID,a) //插入学生id:d;作业id：HomeworkID;状态a(默认情况下0)
//			if e != nil {
//				log.Panicln("user insert error", e.Error())
//				return
//			}
//		}
//	}
//}