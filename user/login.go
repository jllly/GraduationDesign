package user

import (
	"Homework_system/dao"
	"Homework_system/db"
	"Homework_system/global"
	"Homework_system/reg"
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Rep struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}


func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	fmt.Println("file : " + file)
	fmt.Println("ext : " + ext)
	fmt.Println("id : " + id)
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("reload : " + r.FormValue("reload"))
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

//生成验证码图片
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}
	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"` //验证码Id
	ImageUrl  string `json:"imageUrl"`  //验证码图片url
}

//生成4位数的验证码图片
func Verificationcode(context *gin.Context) {
	length := 4
	captchaId := captcha.NewLen(length)
	var captcha CaptchaResponse
	captcha.CaptchaId = captchaId
	captcha.ImageUrl = "/captcha/" + captchaId + ".png"
	context.JSON(200, captcha)
}

func Captchaview(context *gin.Context) {
	captchaId := context.Param("captchaId")
	fmt.Println("GetCaptchaPng : " + captchaId)
	ServeHTTP(context.Writer, context.Request)
}

//提交注册表单
func Submitregistered(context *gin.Context) {
	phonenumber := context.PostForm("phonenumber")   //电话号码
	password := context.PostForm("password")         //密码
	verification := context.PostForm("verification") //输入的验证码
	captchaId := context.PostForm("captchaId")       //图片验证码
	position := context.PostForm("position")         //身份
	var rs Rep
	rs.Code = 200
	rs.Msg = ""
	reg := regexp.MustCompile(reg.Mobilephone)
	if phonenumber == "" || !reg.MatchString(phonenumber) {
		rs.Msg = "手机号码格式错误"
		rs.Code = 500
		context.JSON(http.StatusOK, rs)
		return
	}
	if password == "" {
		rs.Msg = "密码不能为空"
		rs.Code = 500
		context.JSON(http.StatusOK, rs)
		return
	}
	if verification == "" {
		rs.Msg = "验证码不能为空"
		rs.Code = 500
		context.JSON(http.StatusOK, rs)
		return
	}
	//if position == "0" {
	//	//	rs.Msg = "学生"
	//	//	context.JSON(http.StatusOK, rs)
	//	//	return
	//	//}
	//查询是否有账号，有则请登录
	phone, _ := strconv.Atoi(phonenumber)
	var num int
	db.Db.QueryRow("select phonenumber from user where phonenumber = ?;", phone).Scan(&num)
	if num == phone {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "已有账号，请登录",
			"code": "500",
		})
		return
	}

	//图片验证
	if captcha.VerifyString(captchaId, verification) {
		//db.Db.Exec();
		userName := ""
		_, e := db.Db.Exec("insert into homework.user (phonenumber,password,position,userName) values (?,?,?,?);", phone, password, position, userName) //插入信息
		if e != nil {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "注册失败",
				"code": "500",
			})
			log.Panicln("user insert error", e.Error())
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"msg":  "注册成功",
			"code": "200",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "验证码错误",
			"code": "500",
		})
		return
	}
}

//登录
func Login(context *gin.Context) {
	phonenumber := context.PostForm("phonenumber")
	password := context.PostForm("password")
	var rs dao.Rep
	rs.Code = 200
	rs.Msg = ""
	reg := regexp.MustCompile(reg.Mobilephone)
	if !reg.MatchString(phonenumber) {
		rs.Msg = "手机号码格式错误"
		rs.Code = 500
		context.JSON(http.StatusOK, rs)
		return
	}
	if password == "" {
		rs.Msg = "密码不能为空"
		rs.Code = 501
		context.JSON(http.StatusOK, rs)
		return
	}
	u := dao.User{}
	//查询是否有该账号
	row := db.Db.QueryRow("select * from user where phonenumber = ?;", phonenumber)
	err := row.Scan(&u.Phonenumber, &u.Password, &u.Id, &u.Position, &u.UserName)
	if err != nil {
		rs.Code = 500
		rs.Msg = "没有查询到账号，请注册"
		context.JSON(http.StatusOK, rs)
		return
	}
	if u.Password != password {
		rs.Code = 501
		rs.Msg = "密码错误"
		context.JSON(http.StatusOK, rs)
		return
	}
	//设置cookie
	context.SetCookie("auth", phonenumber, 86400, "/", global.Gconfig.WebsitAomain, false, false) //http://35.241.123.232:8080/
	//设置带数据的session
	global.Session[phonenumber] = u
	switch u.Position {
	case 0:
		rs.Msg = "/Student"
	case 1:
		rs.Msg = "/Teacher"
	}
	//重定向
	rs.Code = 304
	context.JSON(http.StatusOK, rs)
}

//忘记密码时获取验证码
func Submitphone(context *gin.Context) {
	phone := context.PostForm("phone")
	var ph global.Phone
	ph.Phone = phone
	ph.Sms = global.GenValidateCode(6)
	ph.Time = time.Now().Unix()
	global.Phonepool = append(global.Phonepool, ph)

}

//忘记密码
func Forget(context *gin.Context) {
	phone := context.PostForm("phone")
	sms := context.PostForm("sms")
	newpassword := context.PostForm("newpassword")
	nowtime := time.Now().Unix()
	var removelist []int
	var rs dao.Rep
	rs.Code = 200
	rs.Msg = ""
	var state int = 0
	for index, value := range global.Phonepool {
		//value.Time
		if (nowtime - value.Time) > 60 {
			removelist = append(removelist, index)
			continue
		}
		if value.Phone == phone {
			//这就要验证的对应号码
			if sms == value.Sms {
				//验证吗正确
				state = 2
			} else {
				//不正确
				state = 1
			}
			break
		}
	}
	removelist = reg.Reverse(removelist)
	for _, value := range removelist {
		global.Phonepool = append(global.Phonepool[:value], global.Phonepool[value+1:]...)
	}
	if state == 2 {
		//执行逻辑
		_, error := db.Db.Exec("UPDATE user SET password=? WHERE phonenumber=?", newpassword, phone)
		if error != nil {
			rs.Code = 500
			rs.Msg = "修改密码失败"
			context.JSON(http.StatusOK, rs)
			return
		}
		rs.Code = 200
		rs.Msg = "修改密码成功"
		context.JSON(http.StatusOK, rs)
	} else if state == 1 {
		rs.Code = 500
		rs.Msg = "验证码错误"
		context.JSON(http.StatusOK, rs)
	} else if state == 0 {
		rs.Code = 500
		rs.Msg = " 验证码过期了"
		context.JSON(http.StatusOK, rs)
	}

}

//退出登录
func LoginOut(context *gin.Context) {
	var rs dao.Rep
	//获取cookie
	myCookie, _ := context.Cookie("auth")
	//删除session
	delete(global.Session, myCookie)
	context.SetCookie("auth", myCookie, -1, "/", global.Gconfig.WebsitAomain, false, false)

	myCookie1, _ := context.Cookie("myCookie")
	fmt.Println(myCookie1)

	rs.Code = 304
	rs.Msg = "/"
	context.JSON(http.StatusOK, rs)
}

//注销账号
func Cancellation_name(context *gin.Context) {
	var rs dao.Rep
	//获取cookie
	myCookie, _ := context.Cookie("auth") //string格式
	// sqlStr := "delete from user where id=?"
	_, err := db.Db.Exec("delete from user where phonenumber=?", myCookie) //删除user数据库内容
	if err != nil {
		rs.Code = 500
		rs.Msg = "注销失败，发生错误！"
		context.JSON(http.StatusOK, rs)
		return
	}
	rs.Code = 304
	rs.Msg = "/"
	context.JSON(http.StatusOK, rs)
}
