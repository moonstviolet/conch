package proto

type NewQuestionReq struct {
	QuestionTitle  string `json:"questionTitle" form:"questionTitle"`
	QuestionDetail string `json:"questionDetail" form:"questionDetail"`
}

type ReadQuestionReq struct {
	Qid int `json:"qid" form:"qid"`
}
