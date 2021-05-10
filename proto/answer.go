package proto

type NewAnswerReq struct {
	Qid          int    `json:"qid" form:"qid"`
	AnswerDetail string `json:"answerDetail" form:"answerDetail"`
}

type ReadAnswerReq struct {
	Qid int `json:"qid" form:"qid"`
}
