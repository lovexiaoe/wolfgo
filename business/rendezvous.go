package business

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"wolfgo/base"

	"gopkg.in/macaron.v1"
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

type UserActivity struct {
	UserActId  int64
	UserId     int64
	Activities `xorm:"extends"`
}

//增加活动
func AddActivity(w http.ResponseWriter, r *http.Request, ctx *macaron.Context) {
	if r.Method != "POST" {
		base.WrongCodeJsonToRsp(w, "请使用POST请求", base.Not_post_req)
		return
	}
	resultmap := make(map[string]interface{})
	rendezvous := new(Activities)

	r.ParseForm()
	title := r.Form["title"][0]
	actUserId := ctx.QueryInt64("actUserId")
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

//我发起的活动列表
func ActUserList(w http.ResponseWriter, r *http.Request, ctx *macaron.Context) {
	if r.Method != "POST" {
		base.WrongCodeJsonToRsp(w, "请使用POST请求", base.Not_post_req)
		return
	}
	resultmap := make(map[string]interface{})

	r.ParseForm()

	actUserIdStr := base.Getformvalue("actUserId", r)
	if actUserIdStr == "" {
		base.WrongCodeJsonToRsp(w, "请求参数为空！", base.Parmeter_isnull)
		return
	}

	actUserId, err := strconv.Atoi(actUserIdStr)
	if err != nil {
		base.WrongCodeJsonToRsp(w, "传入参数错误", base.Params_illegal_err)
		return
	}

	var activities []Activities
	err = base.Engine.Where("actUserId = ?", actUserId).Find(&activities)
	if err != nil {
		base.WrongCodeJsonToRsp(w, "服务器错误", base.Server_run_err)
		return
	}

	resultmap["activities"] = activities
	base.RightJsonToRsp(w, resultmap)
}

//我参与的活动列表
func UserList(w http.ResponseWriter, r *http.Request, ctx *macaron.Context) {
	if r.Method != "POST" {
		base.WrongCodeJsonToRsp(w, "请使用POST请求", base.Not_post_req)
		return
	}
	resultmap := make(map[string]interface{})

	r.ParseForm()

	userIdStr := base.Getformvalue("userId", r)
	if userIdStr == "" {
		base.WrongCodeJsonToRsp(w, "请求参数为空！", base.Parmeter_isnull)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		base.WrongCodeJsonToRsp(w, "传入参数错误", base.Params_illegal_err)
		return
	}

	var activities []Activities
	err = base.Engine.Where("userId = ?", userId).Join("INNER", "userActivity", "userActivity.actId = activities.actId").Find(&activities)
	//	sql := "select orders.*,DATE_FORMAT(orders.created,'%Y-%m-%d %H:%k:%s') as ordercreated,DATE_FORMAT(orders.updated,'%Y-%m-%d %H:%k:%s') as orderupdated,customerName,customerPhone,goodsname,imageurl," +
	//		"user.nickname,user.name from  orders inner join goods on orders.goodsid=goods.id left join customers on orders.customerid=customers.customerId left join user on orders.pardnerid=user.id"
	//	err := engine.Sql(sql).Find(&orders)

	if err != nil {
		base.WrongCodeJsonToRsp(w, "服务器错误", base.Server_run_err)
		return
	}

	resultmap["activities"] = activities
	base.RightJsonToRsp(w, resultmap)
}

//活动详情
func GetActivity(w http.ResponseWriter, r *http.Request, ctx *macaron.Context) {
	if r.Method != "POST" {
		base.WrongCodeJsonToRsp(w, "请使用POST请求", base.Not_post_req)
		return
	}
	resultmap := make(map[string]interface{})

	r.ParseForm()

	actIdStr := base.Getformvalue("actId", r)
	if actIdStr == "" {
		base.WrongCodeJsonToRsp(w, "请求参数为空！", base.Parmeter_isnull)
		return
	}

	actId, err := strconv.Atoi(actIdStr)
	if err != nil {
		base.WrongCodeJsonToRsp(w, "传入参数错误", base.Params_illegal_err)
		return
	}
	activity := new(Activities)

	has, err := base.Engine.Where("actId=?", int64(actId)).Get(activity)

	if err != nil {
		base.WrongCodeJsonToRsp(w, "服务器错误", base.Server_run_err)
		log.Println("服务器错误", err.Error())
		return
	}
	if has == false {
		base.WrongCodeJsonToRsp(w, "没有记录对象", base.Server_run_err)
		return
	}

	resultmap["activity"] = activity
	base.RightJsonToRsp(w, resultmap)
}

//参加活动
//func GetActivity(w http.ResponseWriter, r *http.Request, ctx *macaron.Context) {
//	if r.Method != "POST" {
//		base.WrongCodeJsonToRsp(w, "请使用POST请求", base.Not_post_req)
//		return
//	}
//	resultmap := make(map[string]interface{})
//	r.ParseForm()

//	actIdStr := base.Getformvalue("actId", r)
//	if actIdStr == "" {
//		base.WrongCodeJsonToRsp(w, "请求参数为空！", base.Parmeter_isnull)
//		return
//	}

//	actId, err := strconv.Atoi(actIdStr)
//	if err != nil {
//		base.WrongCodeJsonToRsp(w, "传入参数错误", base.Params_illegal_err)
//		return
//	}

//	userIdStr := base.Getformvalue("userId", r)
//	if userIdStr == "" {
//		base.WrongCodeJsonToRsp(w, "请求参数为空！", base.Parmeter_isnull)
//		return
//	}

//	userId, err := strconv.Atoi(userIdStr)
//	if err != nil {
//		base.WrongCodeJsonToRsp(w, "传入参数错误", base.Params_illegal_err)
//		return
//	}

//	activity := new(Activities)

//	has, err := base.Engine.Where("actId=?", int64(actId)).Get(activity)

//	if err != nil {
//		base.WrongCodeJsonToRsp(w, "服务器错误", base.Server_run_err)
//		log.Println("服务器错误", err.Error())
//		return
//	}
//	if has == false {
//		base.WrongCodeJsonToRsp(w, "没有记录对象", base.Server_run_err)
//		return
//	}

//	resultmap["activity"] = activity
//	base.RightJsonToRsp(w, resultmap)
//}
