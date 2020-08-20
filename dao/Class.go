//对班级数据库的一些操作进行封装
package dao

import (
	"Homework_system/db"
	"log"
)

type Class struct {
	ClassName  string //班级名称
	ClassID    int    //班级ID
	CreateID   int    //创建者id
	Invitation int    //邀请码
	Number     int    //人数
	StartTime  string //创建时间
	MembersID  string //学生ID
}

//班级表,添加,修改,删除

//增加，调用添加函数，把名字，创建者ID，邀请码传入进来，生成新记录

//加入班级,传邀请码，人数，和ID
func Join_class(invitation string, id string) bool {
	//接收人数，邀请码
	//通过邀请码加入，并在班级表中number+1，学生ID加入membersID

	var number int
	err := db.Db.QueryRow("select number from homework.class where invitation = ?;", invitation).Scan(&number)
	if err != nil {
		log.Panicln()
		return false
	}

	userIDStr := "," + id
	if number == 0 {
		_, error := db.Db.Exec("UPDATE homework.class SET number=?,membersID = CONCAT(membersID , ?) where invitation=?", number+1, id, invitation)
		if error != nil {
			return false
		}
	} else {
		_, error := db.Db.Exec("UPDATE homework.class SET number=?,membersID = CONCAT(membersID , ?) where invitation=?", number+1, userIDStr, invitation)
		if error != nil {
			return false
		}
	}
	return true

}

//查询
func Query(invitation int) bool {
	//查找是否有invitation为新生成的数，有则false，没有则为true
	rows := db.Db.QueryRow("select * from homework.class where invitation = ?;", invitation).Scan(&invitation)
	if rows == nil {
		return false
	}
	return true
}
