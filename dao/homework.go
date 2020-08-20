package dao


//学生作业
type Worklist_Stu struct {
	StartTime int	   		`json:"start_time"`
	EndTime  string  		`json:"end_time"`
	HomeworkName string 	`json:"homework_name"`
	ClassID  int 			`json:"class_id"`
	State    string 		`json:"state"`
	Topic    string			`json:"topic"`
	Score	 int 			`json:"score"`
	Answer   string			`json:"answer"`
}

//老师作业
type Worklist_Tea struct {
	HomeworkName string 	`json:"homework_name"`
	StartTime int	   		`json:"start_time"`
	EndTime  string  		`json:"end_time"`
	ClassName string 		`json:"class_name"`
	Number	int 			`json:"number"`
	SubmitNum  int 			`json:"submit_num"`

}

