package main

import (
	"blog/dao"
	"blog/model"
)

func main() {

	user := model.User{
		Username: "tom",
		Password: "123",
	}

	dao.Mgr.AddUser(&user)

}
