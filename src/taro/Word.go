package main

import "fmt"

type TarsusStream struct {
	_Data *[]any
}

func (t TarsusStream) StorageData(data *[]any) {
	t._Data = data
}

func (t TarsusStream) readInt(index int) int {
	return 0
}
func (t TarsusStream) readString(index int) string {
	return "1"
}

type Word struct {
	Id   int    `json:"id"`
	Name string `json:"name`
}

func (parse Word) Constructor(data []any) *Word {
	tarsusStream := new(TarsusStream)
	tarsusStream.StorageData(&data)
	w := new(Word)
	w.Id = tarsusStream.readInt(0)
	w.Name = tarsusStream.readString(1)
	return w
}

type WordTranslate struct {
	Id          int    `json:"id"`
	Cn_name     string `json:"cn_name"`
	En_type     string `json:"en_type"`
	Own_mark    string `json:"own_mark"`
	Create_time string `json:"create_time"`
	Word_id     int    `json:"word_id"`
	User_id     int    `json:"user_id"`
}

func (parse WordTranslate) Constructor(data []any) {

}

type DelOrSaveRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type getWordListReq struct {
	Desc       string `json:"desc"`
	Keyword    string `json:"keyword"`
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
	User_id    int    `json:"user_id"`
	Type       int    `json:"type"`
}

type getWordListRes struct {
	Code  int    `json:"code"`
	List  []Word `json:"list"`
	Total int    `json:"total"`
}

type getTranslateListReq struct {
	Desc       string `json:"desc"`
	Keyword    string `json:"keyword"`
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
	User_id    int    `json:"user_id"`
	Type       int    `json:"type"`
}

type getTranslateListRes struct {
	Code  int             `json:"code"`
	List  []WordTranslate `json:"list"`
	Total int             `json:"total"`
}

func main() {
	w := new(Word)
	ints := []any{1, 2, 3, 4, 5, 6, 7}
	constructor := w.Constructor(ints)
	fmt.Println(constructor.Id)
	fmt.Println(constructor.Name)
}
