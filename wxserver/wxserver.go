package main

import (
	"wolfgo/base"
	"wolfgo/business"

	"gopkg.in/macaron.v1"
)

func main() {
	//xorm database init
	base.InitDB()
	//macaron route
	m := macaron.Classic()

	//activity
	m.Post("/wolfgo/activity/add", business.AddActivity)
	m.Post("/wolfgo/activity/actUserList", business.ActUserList)
	m.Post("/wolfgo/activity/userList", business.UserList)
	m.Post("/wolfgo/activity/getActivity", business.GetActivity)

	//user
	m.Post("/wolfgo/user/addUser", business.AddUser)

	m.Run(8090)
}
