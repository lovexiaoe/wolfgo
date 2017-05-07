# wolfgo 

### 添加用户信息 

请求地址：wolfgo/user/addUser

请求方式：Post

Content-Type:application/x-www-form-urlencoded

入参：
 
| 名称             | 表单类型        | 样例               |
| :--------------- | :--------------| :------------      |
| nickName         | text           | 阿里开发            |
| openId           | text           | 斯蒂芬四分          |
| sex              | text           | 1                  |
| city             | text           | 成都市             |
| province         | text           | 四川省             |
| country          | text           | 高新区             |
| headImg          | text           | http://12323.jpg   |

正确返回：

```
{ "result":true,
  "user":{
        "UserId":0,
        "NickName":"阿里开发",
        "OpenId":"123123213",
        "Sex":1,
        "City":" 阿娇等级分",
        "Province":"四川省",
        "Country":"高新区",
        "HeadImg":"http://829384.jpg",
        "Created":"2017-05-07T23:12:21.333156333+08:00"
   }
}
```


