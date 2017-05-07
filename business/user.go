package business

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"wolfgo/base"

	"gopkg.in/macaron.v1"
)

type User struct {
	UserId   int64
	NickName string
	OpenId   string
	Sex      int
	City     string
	Province string
	Country  string
	HeadImg  string
	Created  time.Time
}

//增加活动
func AddUser(w http.ResponseWriter, r *http.Request, ctx *macaron.Context) {
	if r.Method != "POST" {
		base.WrongCodeJsonToRsp(w, "请使用POST请求", base.Not_post_req)
		return
	}
	resultmap := make(map[string]interface{})
	user := new(User)

	r.ParseForm()
	nickName := r.Form["nickName"][0]
	openId := r.Form["openId"][0]
	sexStr := r.Form["sex"][0]
	city := base.Getformvalue("city", r)
	province := base.Getformvalue("province", r)
	country := base.Getformvalue("country", r)
	headImg := base.Getformvalue("headImg", r)

	sex, err := strconv.Atoi(sexStr)
	if err != nil {
		base.WrongCodeJsonToRsp(w, "传入参数错误", base.Params_illegal_err)
		return
	}

	user.City = city
	user.Country = country
	user.Province = province
	user.HeadImg = headImg
	user.Sex = sex
	user.OpenId = openId
	user.NickName = nickName
	user.Created = time.Now()

	user1 := new(User)
	has, err := base.Engine.Where("openId=?", user.OpenId).Get(user1)
	if err != nil {
		base.WrongCodeJsonToRsp(w, "服务器错误", base.Server_run_err)
		return
	}
	if has == true {
		//如果已经有openid，则更新资料
		//coach.Id = coa.Id
		user.Created = user1.Created

		_, err := base.Engine.Where("id=?", user1.UserId).Update(user)
		if err != nil {
			log.Println("runerr:", err.Error())
		}
		user.UserId = user1.UserId
	} else {
		_, err = base.Engine.Insert(user)
		if err != nil {
			base.WrongCodeJsonToRsp(w, "用户信息添加失败", base.Server_run_err)
			return
		}
	}

	resultmap["user"] = user
	base.RightJsonToRsp(w, resultmap)
}
