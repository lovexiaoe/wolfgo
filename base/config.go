package base

import (
	"log"
)

const (
	//---返回给前端的处理错误code
	Server_run_err = -1 //服务器程序处理错误
	//user_or_pass_err       = -2  //账户或者密码错误
	Params_illegal_err = -3 //参数非法
	//unknown_err            = -4  //未知错误
	//start_session_err      = -5  //启动sesseion失败
	Not_post_req = -6 //非POST请求
	//get_wxopenid_fail      = -7  //获取openid失败
	//get_wxtoken_fail       = -8  //获取access_token失败
	//get_wxuserinfo_fail    = -9  //获取微信用户信息失败
	//wx_no_subscribe        = -10 // 用户未关注公众号
	//get_wxjsticket_fail    = -11 //获取微信jsticket失败
	Affected_isnull_err = -12 //未有数据更改
	//already_has_result_err = -13 //已经是最终状态
	//transaction_err        = -14 //事务处理失败
	Dateformat_err  = -15 //日期格式错误
	Parmeter_isnull = -16 //未有数据更改

	server_online_1 = "" //上线服务器，db
	mysql_url       = ""
)

func init() {

	//初始化log
	log.SetPrefix("logger:")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}
