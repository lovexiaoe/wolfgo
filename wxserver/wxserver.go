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

	//m.Post("/wolfgo/index", business.WxGetSignature)
	m.Post("/wolfgo/activity/add", business.AddActivity)

	m.Run(8090)
}
