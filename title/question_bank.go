package title

import (
	"Homework_system/dao"
	"Homework_system/db"
	"Homework_system/global"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Question struct {
	Id             int    `json:"id"`
	Topic_describe string `json:"topic_describe"`
	Category       int    `json:"category"`
	Answer         string `json:"answer"`
	Annotation     string `json:"annotation"`
	Knowledge      int    `json:"knowledge"`
	CreateID       int    `json:"create_id"`
}

type Questionlist struct {
	question Question
}

//添加题目
func Add_question(context *gin.Context) {
	information := global.Get_Session(context)
	var rs dao.Rep
	createID := information.Id                           //创建者id
	category := context.PostForm("category")             //题目类型 0选择，1填空，2解答
	answer := context.PostForm("answer")                 //题目答案
	annotation := context.PostForm("annotation")         //题目注释
	knowledge := context.PostForm("knowledge")           //知识点
	topic_describe := context.PostForm("topic_describe") //题目描述

	if answer == "" {
		rs.Msg = "答案不能为空"
		rs.Code = 500
		context.JSON(http.StatusOK, rs)
		return
	}
	if topic_describe == "" {
		rs.Msg = "题目描述不能为空"
		rs.Code = 500
		context.JSON(http.StatusOK, rs)
		return
	}

	_, e := db.Db.Exec("insert into homework.title_bank (topic_describe,category,answer,annotation,knowledge,createID) values (?,?,?,?,?,?);", topic_describe, category, answer, annotation, knowledge, createID) //插入信息
	if e != nil {
		rs.Code = 500
		rs.Msg = "添加题目出错！"
		context.JSON(http.StatusOK, rs)
		log.Panicln("user insert error", e.Error())
		return
	}
	rs.Code = 200
	rs.Msg = "添加题目成功！"
	context.JSON(http.StatusOK, rs)
}

//显示题目
func Question_List(context *gin.Context) {
	//拿到老师ID，把题目表中创建者老师ID和知识点为当前的显示出来
	information := global.Get_Session(context)
	knowledge := context.Query("knowledg")
	rows, _ := db.Db.Query("SELECT  * FROM homework.title_bank WHERE createID=? and knowledge=?; ", information.Id, knowledge)
	defer rows.Close()
	if rows == nil {
		return
	}
	ulist := make([]Question, 0, 0)
	for rows.Next() {
		//填充数据
		var row Question
		//每行数据
		rows.Scan(&row.Id, &row.Topic_describe, &row.Category, &row.Answer, &row.Annotation, &row.Knowledge, &row.CreateID)

		//放入结果集
		ulist = append(ulist, row)
	}
	context.JSON(http.StatusOK, ulist)
	return
}
