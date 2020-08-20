package homework

import (
	"Homework_system/dao"
	"Homework_system/db"
	"github.com/gin-gonic/gin"
)
//给前台一个总数据
type work struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data [](dao.Worklist_Tea) `json:"data"`
}
type topic struct {

}

//做作业
func Do_homework(context *gin.Context)  {
	//查对应的题
	//学生做作业，通过作业ID，找到对应的题目并显示
	var HomeworkID string   //HomeworkID,通过地址栏拿到作业ID，
	var topicID string
	err := db.Db.QueryRow("select topicID from homework.teacher_homework where HomeworkID=?",HomeworkID).Scan(&topicID)
	if err != nil {
		//rs.Code = 500
		//rs.Msg = "信息出错！"
		//context.JSON(http.StatusOK, rs)
		//return
	}
	//传给前台数据，
	//1。判断题型
	//2。根据不同的题型自增编号
	//topic:=strings.Split(topicID,",")
	//for i,v:=range topic{
	//	//v题目号
	//}

}
