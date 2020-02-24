package serverapi

import (
	"lib/common/chttp"
	"lib/common/errcode"
)

var load_es_url string = "http://base_question.in.netwa.cn:8004/v1/basequestion/inner/back/load_elasticsearch"
var add_question_url string = "http://base_question.in.netwa.cn:8004/v1/basequestion/inner/back/add_question"
var update_question_url string = "http://base_question.in.netwa.cn:8004/v1/basequestion/inner/back/update_question"
var user_get_question_url string = "http://base_question.in.netwa.cn:8004/v1/basequestion/inner/user/get_question"

type QuestionJsonV1 struct {
	Id        int64    `json:"id"`
	Star      int32    `json:"star"`
	Title     string   `json:"title"`
	AnswerA   string   `json:"answer_a"`
	AnswerB   string   `json:"answer_b"`
	AnswerC   string   `json:"answer_c"`
	AnswerD   string   `json:"answer_d"`
	IsA       int8     `json:"is_a"`
	IsB       int8     `json:"is_b"`
	IsC       int8     `json:"is_c"`
	IsD       int8     `json:"is_d"`
	FirstTag  string   `json:"first_tag"`
	SecondTag string   `json:"second_tag"`
	ThirdTag  []string `json:"third_tag"`
	FourthTag []string `json:"fourth_tag"`
}

type QuestionRes struct {
	Index int `json:"index"`
	Code  int `json:"code"`
}

type UserGetQuestionReq struct {
	FirstTag string `json:"first_tag"`
	Size     int    `json:"size"`
}

type UserGetQuestionRes struct {
	Questions []QuestionJsonV1 `json:"questions"`
}

type UpdateQuestionReq struct {
	Questions []QuestionJsonV1 `json:"questions"`
}

type UpdateQuestionRes struct {
	Questions []QuestionRes `json:"questions"`
}

type AddQuestionReq struct {
	Questions []QuestionJsonV1 `json:"questions"`
}

type AddQuestionRes struct {
	Questions []QuestionRes `json:"questions"`
}

type BackLoadDataToEsReq struct {
	Cmd int `json:"cmd"`
}

type BackLoadDataToEsRes struct {
	Code int `json:"code"`
}

func LoadElasticSearch(req *BackLoadDataToEsReq, res *BackLoadDataToEsRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}
	return chttp.InnerRequest(load_es_url, req, res)
}

func AddQuestion(req *AddQuestionReq, res *AddQuestionRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}
	return chttp.InnerRequest(add_question_url, req, res)
}

func UpdateQuestion(req *UpdateQuestionReq, res *UpdateQuestionRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}
	return chttp.InnerRequest(update_question_url, req, res)
}

func UserGetQuestion(req *UserGetQuestionReq, res *UserGetQuestionRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}
	return chttp.InnerRequest(user_get_question_url, req, res)
}
