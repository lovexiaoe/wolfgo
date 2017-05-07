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

### 获取活动详情 

请求地址：wolfgo/activity/getActivity

请求方式：Post

Content-Type:application/x-www-form-urlencoded

入参：
 
| 名称             | 表单类型        | 样例               |
| :--------------- | :--------------| :------------      |
| actId         | text           | 1            |


正确返回：

```
{"activity":
          {"ActId":1,
          "Title":"新建活动1",
          "ActUserId":1,
          "ActTime":"2017-05-01T06:40:00+08:00",
          "ActAddress":"阿斯地方打算速度",
          "ActUserAcount":0,
          "ActStatus":0,
          "ImgUrl":"asdf.img",
          "ActRemark":"",
          "Yn":1,
          "Created":"2017-04-30T22:53:38+08:00",
          "Modified":"2017-04-30T22:53:38+08:00"
          },
 "result":true
}
```

### 添加活动

请求地址：/wolfgo/activity/add

请求方式：Post

Content-Type:application/x-www-form-urlencoded

入参：
 
| 名称             | 表单类型        | 样例               |
| :--------------- | :--------------| :------------      |
| title         | text           | 新建活动1            |
| actUserId           | text           | 1          |
| actTime              | text           | 2017-04-30 22:40:00                 |
| actAddress             | text           | 阿斯地方打算速度             |
| imgUrl         | text           | asdf.img             |
| actRemark          | text           | 活动描述内容             |


正确返回：

```
{"activity":
      {"ActId":0,
      "Title":"新建活动1",
      "ActUserId":1,
      "ActTime":"2017-04-30T22:40:00Z",
      "ActAddress":"阿斯地方打算速度",
      "ActUserAcount":0,
      "ActStatus":0,
      "ImgUrl":"asdf.img",
      "ActRemark":"活动描述内容",
      "Yn":1,
      "Created":"2017-05-07T23:45:19.0354748+08:00",
      "Modified":"2017-05-07T23:45:19.0354748+08:00"
      },
 "result":true
 }
```

