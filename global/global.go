package global

import (
	"Homework_system/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Phone struct {
	Phone string
	Sms   string
	Time  int64
}

type Config struct {
	WebsitAomain string `json:"WebsitAomain"`
	WebsitPort   string `json:"WebsitPort"`
}

var (
	//WebsitAomain string  = "127.0.0.1"
	Gconfig Config
)

var Phonepool []Phone

var Session = make(map[string]dao.User, 16)

//产生随机数
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

//获取session信息
func Get_Session(context *gin.Context) dao.User {
	//获取cookie
	myCookie, _ := context.Cookie("auth")
	//通过cookie，拿到session里带的信息
	return Session[myCookie]
}


func Timeformat(timestr string)(t time.Time, e error){
	endTime, e := strconv.ParseInt(timestr, 10, 64)
	if e!=nil{
		return time.Now(),e
	}
	timeObj1 := time.Unix(endTime * 1e6 / 1e9, 0)
	timeObj1.Format("2006-01-02 15:04:05")
	return timeObj1 ,nil
}


