package business

import (
	"log"
	"net/http"
	"time"
	"wolfgo/base"
)

type Activities struct {
	ActId         int64
	Title         string
	ActUserId     int64
	ActTime       time.Time
	ActAddress    string
	ActUserAcount int64
	ActStatus     int
	ImgUrl        string
	ActRemark     string
	Yn            int
	Created       time.Time
	Modified      time.Time
}

//添加活动
func AddActivity(w http.ResponseWriter, r *http.Request, ctx *iris.Context) {
	if r.Method != "POST" {
		base.WrongCodeJsonToRsp(w, "请使用POST请求", base.Not_post_req)
		return
	}
	resultmap := make(map[string]interface{})
	rendezvous := new(Activities)

	r.ParseForm()
	title := r.Form["title"][0]
	actUserId := r.Form("actUserId")[0]
	actTimeStr := r.Form["actTime"][0]
	actAddress := r.Form["actAddress"][0]
	imgUrl := r.Form["imgUrl"][0]
	actRemark := base.Getformvalue("actRemark", r)

	rendezvous.Title = title
	rendezvous.ActUserId = actUserId
	timefmt := "2006-01-02 15:04:05"

	actTime, err := time.Parse(timefmt, actTimeStr)
	if err != nil {
		log.Println(err.Error())
		base.WrongCodeJsonToRsp(w, "日期格式错误", base.Dateformat_err)
		return
	}
	rendezvous.ActTime = actTime
	rendezvous.ActAddress = actAddress
	rendezvous.ImgUrl = imgUrl
	rendezvous.ActRemark = actRemark
	rendezvous.ActUserAcount = 0
	rendezvous.Created = time.Now()
	rendezvous.Modified = time.Now()
	rendezvous.Yn = 1

	base.RightJsonToRsp(w, resultmap)

	affected, err := base.Engine.Insert(rendezvous)
	if err != nil {
		base.WrongCodeJsonToRsp(w, "服务器错误", base.Server_run_err)
		return
	}
	if affected == 0 {
		base.WrongCodeJsonToRsp(w, "未有数据更改", base.Affected_isnull_err)
		return
	}

	resultmap["activity"] = rendezvous
	base.RightJsonToRsp(w, resultmap)
}

//添加活动
func ListActivities(w http.ResponseWriter, r *http.Request, ctx *iris.Context) {
	if r.Method != "POST" {
		base.WrongCodeJsonToRsp(w, "请使用POST请求", base.Not_post_req)
		return
	}
	resultmap := make(map[string]interface{})
	rendezvous := new(Activities)

	r.ParseForm()
	title := r.Form["title"][0]
	actUserId := r.Form("actUserId")[0]
	actTimeStr := r.Form["actTime"][0]
	actAddress := r.Form["actAddress"][0]
	imgUrl := r.Form["imgUrl"][0]
	actRemark := base.Getformvalue("actRemark", r)

	rendezvous.Title = title
	rendezvous.ActUserId = actUserId
	timefmt := "2006-01-02 15:04:05"

	actTime, err := time.Parse(timefmt, actTimeStr)
	if err != nil {
		log.Println(err.Error())
		base.WrongCodeJsonToRsp(w, "日期格式错误", base.Dateformat_err)
		return
	}
	rendezvous.ActTime = actTime
	rendezvous.ActAddress = actAddress
	rendezvous.ImgUrl = imgUrl
	rendezvous.ActRemark = actRemark
	rendezvous.ActUserAcount = 0
	rendezvous.Created = time.Now()
	rendezvous.Modified = time.Now()
	rendezvous.Yn = 1

	base.RightJsonToRsp(w, resultmap)

	affected, err := base.Engine.Insert(rendezvous)
	if err != nil {
		base.WrongCodeJsonToRsp(w, "服务器错误", base.Server_run_err)
		return
	}
	if affected == 0 {
		base.WrongCodeJsonToRsp(w, "未有数据更改", base.Affected_isnull_err)
		return
	}

	resultmap["activity"] = rendezvous
	base.RightJsonToRsp(w, resultmap)
}
