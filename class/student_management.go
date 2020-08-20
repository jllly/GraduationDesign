//老师管理班级学生
package class

import (
	"Homework_system/dao"
	"Homework_system/db"
	"Homework_system/reg"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//学生列表
type Userlistreq struct {
	Code        int             `json:"code"`
	Msg         string          `json:"msg"`
	Count       int             `json:"count"`       //数据总条数
	Currentpage int             `json:"Currentpage"` //当前页
	Totalpage   int             `json:"Totalpage"`   //最大页数
	Data        []Query_student `json:"data"`        //数据

}

type Query_student struct {
	Phonenumber int    `json:"phonenumber"`
	Id          int    `json:"id"`
	UserName    string `json:"user_name"`
}

type SearchUserList struct {
	Msg   string          `json:"msg"`
	Code  int             `json:"code"`
	Count int             `json:"count"` //数据总条数
	Data  []Query_student `json:"data"`
}

//班级学生管理
func Student_list(context *gin.Context) {
	//根据邀请码拿到学生ID名单。
	invitation := context.Query("invitation")
	var membersID string
	var number int
	//拿到班级下学生ID名单，人数
	db.Db.QueryRow("select membersID,number from homework.class where invitation = ?;", invitation).Scan(&membersID, &number)

	//	db.Db.QueryRow("SELECT phonenumber,id  from user where id IN ?;",a).Scan(&u.Phonenumber,&u.Id)
	//fmt.Println(u.Phonenumber,u.Id)
	//一页总数
	pagesize := context.DefaultQuery("limit", "10")
	//当前页
	pageNumber := context.DefaultQuery("page", "1")
	//初始化Data切片
	ulist := make([]Query_student, 0, 0)
	//初始化结构体
	resdata := Userlistreq{
		Code:        0,
		Msg:         "学生列表",
		Currentpage: 0,
		Totalpage:   0,
		Count:       0,
		Data:        ulist,
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
	//查询当前页数据
	//返回姓名，手机号，ID号
	//SELECT  phonenumber,id,userName FROM homework.user WHERE  FIND_IN_SET(id,(select membersID from homework.class where invitation = 9841));
	row1, _ := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS phonenumber,id,userName FROM homework.user WHERE FIND_IN_SET(id,?) ORDER BY id DESC LIMIT ?, ?; ", membersID, (pageNumberInt-1)*pagesizeInt, pagesizeInt)
	defer row1.Close()
	if row1 == nil {
		return
	}
	var i int
	for row1.Next() {
		//填充数据
		var row Query_student
		row1.Scan(&row.Phonenumber, &row.Id, &row.UserName)
		//每行数据
		//放入结果集
		resdata.Data = append(resdata.Data, row)
		i++
	}
	resdata.Code = 0
	resdata.Msg = "学生列表成功获取"
	resdata.Count = number //总条数
	resdata.Currentpage = pageNumberInt
	//总页码
	//传过去姓名，ID，手机号。json数据
	//db.Db.QueryRow("SELECT count(*)  FROM homework.class WHERE createID = ? ;",information.Id).Scan(&count)
	totalPageNum := (number + pagesizeInt - 1) / pagesizeInt
	resdata.Totalpage = totalPageNum
	context.JSON(http.StatusOK, resdata)
	return
}

//添加学生
func Add_student(context *gin.Context) {
	invitation := context.Query("invitation")
	data := context.PostForm("data") //获取要加班级的学生的ID或手机号
	var rs dao.Rep

	//判断传过来的数据的格式
	reg := regexp.MustCompile(reg.Mobilephone)
	//加入班级,传邀请码，人数，和ID
	if reg.MatchString(data) {
		//是手机号的格式,手机号加入班级
		var phone int
		var id string
		var position int
		err := db.Db.QueryRow("select phonenumber,id,position from homework.user where phonenumber =?", data).Scan(&phone, &id, &position)
		if err != nil {
			//手机号不存在
			rs.Msg = "手机号不存在，请重新确定！"
			rs.Code = 500
			context.JSON(http.StatusOK, rs)
			return
		} else if position == 1 {
			rs.Msg = "该用户是老师，不能进行此操作！"
			rs.Code = 500
			context.JSON(http.StatusOK, rs)
			return
		}
		if dao.Join_class(invitation, id) {
			rs.Code = 200
			rs.Msg = "加入班级成功"
			context.JSON(http.StatusOK, rs)
			return
		}
	} else {
		//是ID号的格式,ID加入班级
		var id string
		err := db.Db.QueryRow("select id from homework.user where id =?", data).Scan(&id)
		if err != nil {
			//手机号不存在
			rs.Msg = "id号不存在，请重新确定！"
			rs.Code = 500
			context.JSON(http.StatusOK, rs)
			return
		}
		if dao.Join_class(invitation, data) {
			rs.Code = 200
			rs.Msg = "加入班级成功"
			context.JSON(http.StatusOK, rs)
			return
		}
	}
}

//批量删除学生
func Batch_delete(context *gin.Context) {
	var rs dao.Rep
	data := context.PostForm("setID") //获取要删除的学生的ID集
	//拿到要删除的membersID，然后进行判断，删除。全部删除完成之后把membersID字段更新
	var membersID string
	var number int
	invitation := context.Query("invitation")
	err := db.Db.QueryRow("select number,membersID from homework.class where invitation=?", invitation).Scan(&number, &membersID)
	if err != nil {
		//邀请码错误
		rs.Msg = "邀请码错误"
		rs.Code = 500
		context.JSON(http.StatusOK, rs)
		return
	}
	for index, v := range membersID {
		fmt.Printf("%T,%s,%T,%s", index, index, v, v)
	}
	mysqldata := strings.Split(membersID, ",") //数据库数据
	deletedata := strings.Split(data, ",")     //前台数据
	//差集
	var count int
	for _, set := range deletedata {
		for index, set2 := range mysqldata {
			if set == set2 {
				mysqldata = append(mysqldata[:index], mysqldata[index+1:]...)
				count = count + 1
			}
		}
	}
	number = number - count
	fmt.Println(mysqldata)
	sjk1 := strings.Join(mysqldata, ",")
	fmt.Println(sjk1)
	//更新数据库数据
	_, error := db.Db.Exec("UPDATE homework.class SET number=?,membersID=? WHERE invitation=?", number, sjk1, invitation)
	if error != nil {
		rs.Msg = "删除发生错误！"
		rs.Code = 500
		context.JSON(http.StatusOK, rs)
		return
	}
	rs.Msg = "删除成功"
	rs.Code = 200
	context.JSON(http.StatusOK, rs)
}

//搜索
func Search_student(context *gin.Context) {
	var rs SearchUserList
	//判断通过什么搜索，姓名/ID/手机号
	way := context.Query("way")      //通过什么方式查询，0是姓名，1是手机号
	data := context.Query("content") //搜索框数据
	invitation := context.Query("invitation")

	var membersID string
	//通过邀请码拿到班级下学生ID名单
	db.Db.QueryRow("select membersID from homework.class where invitation = ?;", invitation).Scan(&membersID)
	//拿到所有学生信息
	rows, _ := db.Db.Query("SELECT phonenumber,id,userName FROM homework.user WHERE FIND_IN_SET(id,?); ", membersID)
	defer rows.Close()
	if rows == nil {
		fmt.Printf("没有")
		return
	}
	//初始化Data切片
	ulist := make([]Query_student, 0, 0)
	var i int
	for rows.Next() {
		//填充数据
		var row Query_student
		rows.Scan(&row.Phonenumber, &row.Id, &row.UserName)
		if way == "0" {
			//0:通过名字查询
			if strings.Contains(row.UserName, data) {
				ulist = append(ulist, row)
				i++
			}
		} else {
			//通过手机号查询
			if strings.Contains(strconv.Itoa(row.Phonenumber), data) {
				ulist = append(ulist, row)
				i++
			}
		}

	}

	rs.Code = 0
	rs.Msg = "查询列表成功获取"
	rs.Count = len(ulist)
	rs.Data = ulist
	//resdata.Count = number //总条数
	//resdata.Currentpage = pageNumberInt
	//totalPageNum := (number + pagesizeInt - 1) / pagesizeInt
	//resdata.Totalpage = totalPageNum
	context.JSON(http.StatusOK, rs)
	return

}

//下载名单
func Download_list(context *gin.Context) {
	//判断通过什么搜索，姓名/ID/手机号
	invitation := context.Query("invitation")
	var membersID string
	var className string
	//通过邀请码拿到班级下学生ID名单
	db.Db.QueryRow("select membersID,className from homework.class where invitation = ?;", invitation).Scan(&membersID, &className)
	//拿到所有学生信息
	rows, _ := db.Db.Query("SELECT phonenumber,id,userName FROM homework.user WHERE FIND_IN_SET(id,?); ", membersID)
	defer rows.Close()
	if rows == nil {
		fmt.Printf("没有")
		return
	}
	//初始化Data切片
	ulist := make([]Query_student, 0, 0)
	for rows.Next() {
		//填充数据
		var row Query_student
		rows.Scan(&row.Phonenumber, &row.Id, &row.UserName)
		ulist = append(ulist, row)
	}
	//生成Excel 文件
	err, filepath := GeneratePeopleExcel(ulist, className)
	if err != nil {
		fmt.Printf("生成文件失败")
		return
	}
	filepath = filepath
	//filepath
	context.Header("Content-Disposition", "attachment; filename=file.xls")
	context.Header("Content-Type", "application/octet-stream")
	context.File(filepath)

	return

}

type XlsxRow struct {
	Row  *xlsx.Row
	Data []string
}

func NewRow(row *xlsx.Row, data []string) *XlsxRow {
	return &XlsxRow{
		Row:  row,
		Data: data,
	}
}
func (row *XlsxRow) SetRowTitle() error {
	return generateRow(row.Row, row.Data)
}
func (row *XlsxRow) GenerateRow() error {
	return generateRow(row.Row, row.Data)
}

func generateRow(row *xlsx.Row, rowStr []string) error {
	if rowStr == nil {
		return errors.New("no data to generate xlsx!")
	}
	for _, v := range rowStr {
		cell := row.AddCell()
		cell.SetString(v)
	}
	return nil
}

func GeneratePeopleExcel(peo []Query_student, classname string) (error, string) {

	t := make([]string, 0)
	t = append(t, "id")
	t = append(t, "姓名")
	t = append(t, "电话")
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet")
	if err != nil {
		return err, ""
	}
	titleRow := sheet.AddRow()
	xlsRow := NewRow(titleRow, t)
	err = xlsRow.SetRowTitle()
	if err != nil {
		return err, ""
	}
	for _, v := range peo {
		currentRow := sheet.AddRow()
		tmp := make([]string, 0)
		tmp = append(tmp, strconv.Itoa(v.Id))
		tmp = append(tmp, v.UserName)
		tmp = append(tmp, strconv.Itoa(v.Phonenumber))

		xlsRow := NewRow(currentRow, tmp)
		err := xlsRow.GenerateRow()
		if err != nil {
			return err, ""
		}
	}

	nowtime := time.Now()
	path := "./" + classname + strconv.Itoa(int(nowtime.Unix())) + ".xlsx"

	err = file.Save(path)
	if err != nil {
		return err, ""
	}
	return nil, path
}
