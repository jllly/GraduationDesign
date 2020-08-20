//班级管理
package class

import (
	"Homework_system/dao"
	"Homework_system/db"
	"Homework_system/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"

	"time"
)

//班级列表
type Classlistreq struct {
	Count       int           `json:"Count"`       //数据总条数
	Currentpage int           `json:"Currentpage"` //当前页
	Totalpage   int           `json:"Totalpage"`   //最大页数
	Data        [](dao.Class) `json:"Data"`        //数据
}

//查询班级列表 学生
func Claa_List_students(context *gin.Context) {
	//一页总数
	pagesize := context.DefaultQuery("page_size", "10")
	//当前页
	pageNumber := context.DefaultQuery("page_number", "1")
	//初始化Data切片
	dlist := make([]dao.Class, 0, 0)
	//初始化结构体
	resdata := Classlistreq{
		Currentpage: 0,
		Totalpage:   0,
		Count:       0,
		Data:        dlist,
	}
	//类型转换
	pagesizeInt, err := strconv.Atoi(pagesize)
	if err != nil || pagesizeInt < 0 {
		context.JSON(http.StatusOK, resdata)
		return
	}
	//类型转换
	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil || pageNumberInt < 0 {
		context.JSON(http.StatusOK, resdata)
		return
	}

	if pageNumberInt <= 0 {
		pageNumberInt = 1
	}
	if pagesizeInt <= 0 {
		pagesizeInt = 10
	}
	//limitInt
	//pageNumberInt
	//拿到老师ID，把班级表中创建者老师ID拿到
	information := global.Get_Session(context)
	//查询当前页数据
	//rows, _ := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS * FROM homework.class WHERE membersID like  ?  ORDER BY startTime DESC  LIMIT ?, ?; ", information.Id, (pageNumberInt-1)*pagesizeInt, pagesizeInt)
	rows, _ := db.Db.Query("SELECT * FROM class WHERE FIND_IN_SET(?,membersID)  ORDER BY startTime DESC  LIMIT ?, ?; ", information.Id, (pageNumberInt-1)*pagesizeInt, pagesizeInt)

	defer rows.Close()
	if rows == nil {
		return
	}
	var i int
	for rows.Next() {
		//填充数据
		var row dao.Class
		rows.Scan(&row.ClassName, &row.ClassID, &row.CreateID, &row.Invitation, &row.Number, &row.StartTime, &row.MembersID)
		//每行数据
		//放入结果集
		resdata.Data = append(resdata.Data, row)
		i++
	}
	resdata.Currentpage = pageNumberInt
	//总页码
	var count int
	db.Db.QueryRow("SELECT count(*)  FROM homework.class WHERE membersID like ? ;", information.Id).Scan(&count)
	totalPageNum := (count + pagesizeInt - 1) / pagesizeInt
	resdata.Totalpage = totalPageNum
	resdata.Count = count
	context.JSON(http.StatusOK, resdata)
	return
}



//班级列表
type THPClasslistreq struct {
	Count       int           `json:"Count"`       //数据总条数
	Currentpage int           `json:"Currentpage"` //当前页
	Totalpage   int           `json:"Totalpage"`   //最大页数
	Data        [](dao.Worklist_Tea) `json:"Data"`        //数据
}

//The teacher's homework is in progress
func Teacher_homework_progress_List(context *gin.Context) {
	//一页总数
	pagesize := context.DefaultQuery("page_size", "10")
	//当前页
	pageNumber := context.DefaultQuery("page_number", "1")
	//初始化Data切片
	dlist := make([]dao.Worklist_Tea, 0, 0)
	//初始化结构体
	resdata := THPClasslistreq{
		Currentpage: 0,
		Totalpage:   0,
		Count:       0,
		Data:        dlist,
	}
	//类型转换
	pagesizeInt, err := strconv.Atoi(pagesize)
	if err != nil || pagesizeInt < 0 {
		context.JSON(http.StatusOK, resdata)
		return
	}
	//类型转换
	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil || pageNumberInt < 0 {
		context.JSON(http.StatusOK, resdata)
		return
	}

	if pageNumberInt <= 0 {
		pageNumberInt = 1
	}
	if pagesizeInt <= 0 {
		pagesizeInt = 10
	}
	//limitInt
	//pageNumberInt
	//拿到老师ID，把班级表中创建者老师ID拿到
	information := global.Get_Session(context)
	//查询当前页数据
	rows, _ := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS * FROM homework.class WHERE createID = ?  ORDER BY startTime DESC  LIMIT ?, ?; ", information.Id, (pageNumberInt-1)*pagesizeInt, pagesizeInt)
	defer rows.Close()
	if rows == nil {
		return
	}
	var i int
	for rows.Next() {
		//填充数据
		var row dao.Class
		rows.Scan(&row.ClassName, &row.ClassID, &row.CreateID, &row.Invitation, &row.Number, &row.StartTime, &row.MembersID)
		//每行数据
		//放入结果集
		//resdata.Data = append(resdata.Data, row)
		i++
	}
	resdata.Currentpage = pageNumberInt
	//总页码
	var count int
	db.Db.QueryRow("SELECT count(*)  FROM homework.class WHERE createID = ? ;", information.Id).Scan(&count)
	totalPageNum := (count + pagesizeInt - 1) / pagesizeInt
	resdata.Totalpage = totalPageNum
	resdata.Count = count
	context.JSON(http.StatusOK, resdata)
	return
}

//查询班级列表
func Claa_List(context *gin.Context) {
	//一页总数
	pagesize := context.DefaultQuery("page_size", "10")
	//当前页
	pageNumber := context.DefaultQuery("page_number", "1")
	//初始化Data切片
	dlist := make([]dao.Class, 0, 0)
	//初始化结构体
	resdata := Classlistreq{
		Currentpage: 0,
		Totalpage:   0,
		Count:       0,
		Data:        dlist,
	}
	//类型转换
	pagesizeInt, err := strconv.Atoi(pagesize)
	if err != nil || pagesizeInt < 0 {
		context.JSON(http.StatusOK, resdata)
		return
	}
	//类型转换
	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil || pageNumberInt < 0 {
		context.JSON(http.StatusOK, resdata)
		return
	}

	if pageNumberInt <= 0 {
		pageNumberInt = 1
	}
	if pagesizeInt <= 0 {
		pagesizeInt = 10
	}
	//limitInt
	//pageNumberInt
	//拿到老师ID，把班级表中创建者老师ID拿到
	information := global.Get_Session(context)
	//查询当前页数据
	rows, _ := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS * FROM homework.class WHERE createID = ?  ORDER BY startTime DESC  LIMIT ?, ?; ", information.Id, (pageNumberInt-1)*pagesizeInt, pagesizeInt)
	defer rows.Close()
	if rows == nil {
		return
	}
	var i int
	for rows.Next() {
		//填充数据
		var row dao.Class
		rows.Scan(&row.ClassName, &row.ClassID, &row.CreateID, &row.Invitation, &row.Number, &row.StartTime, &row.MembersID)
		//每行数据
		//放入结果集
		resdata.Data = append(resdata.Data, row)
		i++
	}
	resdata.Currentpage = pageNumberInt
	//总页码
	var count int
	db.Db.QueryRow("SELECT count(*)  FROM homework.class WHERE createID = ? ;", information.Id).Scan(&count)
	totalPageNum := (count + pagesizeInt - 1) / pagesizeInt
	resdata.Totalpage = totalPageNum
	resdata.Count = count
	context.JSON(http.StatusOK, resdata)
	return
}

//添加班级
func Claa_Add(context *gin.Context) {
	var rs dao.Rep
	rs.Code = 200
	rs.Msg = "班级表添加成功"
	//创建班级，班级表:班级名称,创建者ID,班级ID(主键),邀请码
	//老师输入班级名称，创建班级，自动生成随机
	classname := context.PostForm("classname")
	captial := global.Get_Session(context) //获取session信息
	//生成随机唯一数
	var invitation1 int
	for {
		invitation := global.GenValidateCode(4)   //随机四位数
		invitation1, _ = strconv.Atoi(invitation) //string强制转换成int
		if dao.Query(invitation1) {               //如果没有，即传过来的是true，则跳出循环
			break
		}
	}
	number := 0
	//往班级表里插入信息,班级名称，创建者ID，邀请码,创始人数。班级ID，自增
	startTime := time.Now().Format("2006-01-02 15:04:05")
	var membersID string
	membersID = ""
	_, e := db.Db.Exec("insert into homework.class (className,createID,invitation,number,startTime,membersID ) values (?,?,?,?,?,?);", classname, captial.Id, invitation1, number, startTime, membersID)
	if e != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "班级表生成失败",
			"code": "500",
		})
		log.Panicln("user insert error", e.Error())
	}
	context.JSON(http.StatusOK, rs)
}

//修改班级名称
func ClassName_editor(context *gin.Context) {
	newName := context.PostForm("classname")
	invitation := context.PostForm("invitation")
	invitationint, err := strconv.Atoi(invitation)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "修改班级名称失败",
			"code": "500",
		})
		return
	}
	//通过invitation找到要修改的，将className改为新传进来的
	_, error := db.Db.Exec("UPDATE homework.class SET className=? where invitation=?", newName, invitationint)
	if error != nil {
		log.Println(error)
		fmt.Println("update data fail,err:", error)
		context.JSON(http.StatusOK, gin.H{
			"msg":  "修改班级名称失败",
			"code": "500",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "修改班级名称成功",
		"code": "200",
	})
}

//删除班级
func Class_delete(context *gin.Context) {
	//删除班级
	var invitation string
	//邀请码
	invitation = context.PostForm("invitation")
	invitationint, err := strconv.Atoi(invitation)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "删除班级失败",
			"code": "500",
		})
		return
	}
	//删除
	results, err := db.Db.Exec("DELETE FROM homework.class  where invitation=?", invitationint)
	results = results
	if err != nil {
		fmt.Println("delete data fail,err:", err)
		context.JSON(http.StatusOK, gin.H{
			"msg":  "删除班级2失败",
			"code": "500",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "删除班级表成功",
		"code": "200",
	})
}
