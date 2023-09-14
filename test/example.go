package main

import "fmt"

type Example struct {
}

type UserBaseInfo struct {
	Id       int
	UserName string
}

type UserInfRequest struct {
	Ids []int
}

type UserInfResponse struct {
	Users []UserBaseInfo
}

// ctx *context.Context,
func (imp *Example) getUserListByIds(req UserInfRequest, res *UserInfResponse) (int32, error) {
	return 0, nil
}

func invoke() {
	req := new(UserInfRequest)
	res := new(UserInfResponse)
	e := new(Example)
	ids, err := e.getUserListByIds(*req, res)
	if err != nil {
		fmt.Printf("err %s", err)
	}
	fmt.Printf("ids %s", ids)
}

func main() {
	invoke()
}
