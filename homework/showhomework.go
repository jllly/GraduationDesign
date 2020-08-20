//学生界面显示班级

package homework

import (
	"Homework_system/dao"
	"Homework_system/db"
	"Homework_system/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


type Student_worklist struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data [](dao.Worklist_Stu) `json:"data"`
}
type Teacher_worklist struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data [](dao.Worklist_Tea) `json:"data"`
}


//查询正在进行中的作业(学生)
func Fornow_Student(context *gin.Context){
	var rs Student_worklist
	//获取用户ID
	captial := global.Get_Session(context) //获取session信息
	//获取当前时间,（现在还没有对接口，无法测试）
	now:=time.Now()
	now1:=now.Format("2006-01-02 15:04:05")  //格式当前时间好做比较
	//找到作业的时间、班级、状态和作业名称，在两表作业ID一样，用户ID为当前用户，结束时间与当前时间做比较，当前作业是，结束时间比当前时间大
	rows, _ := db.Db.Query("SELECT StartTime,EndTime,HomeworkName,classID,state FROM homework.student_homework,homework.teacher_homework WHERE homeworkId=HomeworkID and userID=? and Endtime >=?",captial.Id,now1)
	dlist := make([]dao.Worklist_Stu, 0, 0)
	defer rows.Close()
	if rows == nil {
		return
	}
	//返回正在进行的作业的集合数据给前台
	for rows.Next() {
		//填充数据
		var row dao.Worklist_Stu
		rows.Scan(&row.StartTime, &row.EndTime, &row.HomeworkName, &row.ClassID, &row.State)
		//每行数据
		//放入结果集
		dlist = append(dlist, row)

	}
	rs.Code = 200
	rs.Msg = "成功"
	rs.Data = dlist
	context.JSON(http.StatusOK, rs)
}

//往期作业
func Before_Student(context *gin.Context){
	var rs Student_worklist
	//获取用户ID
	captial := global.Get_Session(context) //获取session信息
	//获取当前时间,（现在还没有对接口，无法测试）
	now:=time.Now()
	now1:=now.Format("2006-01-02 15:04:05")  //格式当前时间好做比较
	//找到作业的时间、班级、状态和作业名称，在两表作业ID一样，用户ID为当前用户，结束时间与当前时间做比较，，往期作业，
	rows, _ := db.Db.Query("SELECT StartTime,EndTime,HomeworkName,classID,state FROM homework.student_homework,homework.teacher_homework WHERE homeworkId=HomeworkID and userID=? and Endtime >=?",captial.Id,now1)
	dlist := make([]dao.Worklist_Stu, 0, 0)
	defer rows.Close()
	if rows == nil {
		return
	}
	for rows.Next() {
		//填充数据
		var row dao.Worklist_Stu
		rows.Scan(&row.StartTime, &row.EndTime, &row.HomeworkName, &row.ClassID, &row.State)
		//每行数据
		//放入结果集
		dlist = append(dlist, row)

	}
	rs.Code = 200
	rs.Msg = "成功"
	rs.Data = dlist
	context.JSON(http.StatusOK, rs)
}

//查询正在进行中的作业(老师)
func fornow_Teacher(context *gin.Context)  {
	//给前台数据，时间，作业名称，提交人数			 (老师作业表)
	// 班级名称，班级人数，未提交人数（总人数减提交人数）  （班级表）
	//现在，通过老师ID，查询到班级ID，再在表中查班级ID
	var rs Teacher_worklist
	captial := global.Get_Session(context) //获取用户信息
	//获取当前时间,（现在还没有对接口，无法测试）
	now:=time.Now()
	now1:=now.Format("2006-01-02 15:04:05")  //格式当前时间好做比较

	rows, _ := db.Db.Query("SELECT StartTime,EndTime,HomeworkName,submitNum,className,number FROM homework.teacher_homework,homework.calss WHERE class.createID=? and Endtime >=?",captial.Id,now1)
	dlist := make([]dao.Worklist_Tea, 0, 0)
	defer rows.Close()
	if rows == nil {
		return
	}
	for rows.Next() {

		//填充数据
		var row dao.Worklist_Tea
		rows.Scan(&row.HomeworkName,&row.StartTime, &row.EndTime, &row.ClassName, &row.Number, &row.SubmitNum)
		//每行数据
		//放入结果集
		dlist = append(dlist, row)

	}
	rs.Code = 200
	rs.Msg = "成功"
	rs.Data = dlist
	context.JSON(http.StatusOK, rs)
}