# HIT动漫展会 现场检票 服务端
## 开发部署过程

* nodejs版本：10.0.0 （建议使用**nvm**对nodejs运行时版本进行管理）

* 配置文件 **./config.js**

* 数据库服务器：MySQL 5.7 （不要使用其他版本的MySQL服务器）

  默认数据库配置：

  * 地址：localhost
  * 用户名：root
  * 密码：my-secret-pw
  * 数据库名称：msxh-dev
  * ⚠️无需配置 数据库初始化表格，由Sequelize 数据库ORM自动完成

* 包管理器：Yarn

  安装命令：`yarn install`

  启动命令：`yarn start` （不推荐）

  自动启动：`yarn develop`（推荐，监测到文件变化后，自动更新服务端）

  

  注意⚠️：在项目启动中 Sequelize 会自动执行`sync()`功能，根据./models/目录下的模型创建数据表

  

# hitapp-local-server API 文档



## 功能

`/` - 请求访问首页

`/ping` - 服务端存活检测

`/query` - 票务信息查询(包括 纸质票、电子票、证件)

`/ticket` - 纸质票/电子票 核销接口

`/staff` - Staff证件 核销接口



## 数据库

### Tickets - 门票数据表

| 字段名 | 数据结构 | 说明         |
| ------ | -------- | ------------ |
| Id     | int(11)  | Id           |
| Key    | String   | 票码内容     |
| Type   | int(11)  | 票类型       |
| Times  | int(11)  | 剩余使用次数 |



### Logs - 票务查询核销日志

| 字段名 | 数据结构 | 说明     |
| ------ | -------- | -------- |
| Id     | int(11)  | Id       |
| Key    | String   | 票码内容 |
| Result | int(11)  | 检票状态 |





## API详细细节



### 请求访问首页

* URl : `/`

* 功能：

  通过 Render 相应首页

* 请求报文：`GET /`

* 响应报文：

  ```html
  Content-type: text/html
  HTML 页面
  ```



### ping 检测 服务器存活

* URL: `/ping`

* 功能：

  检测 服务器存活

* 请求报文：`GET /ping`

* 响应报文：

  ```json
  Content-Type: "application/json"
  {	
      "status":"success",
      "result":"pong"
  }
  ```



### 纸质票/电子票 核销接口

* URL： `/ticket`

* 功能：

  查询票务扫描信息

* 请求报文：`POST /ticket`

  ```json
  Content-Type: "application/json"
  {	
      "token":"76a9408e-a8ca-4f37-a27c-2bcf1f1d1cef",
      "key":"101691071aad710c58a4bf72f4802180"
  }
  ```

  | 字段名 | 数据类型 | 说明           |
  | ------ | -------- | -------------- |
  | token  | string   | 鉴权、认证     |
  | key    | string   | 进行查询的票码 |

  

* 响应报文：200

  ```json
  Content-Type: "application/json"
  {	
      "result":"success"
  }
  ```

  | 字段名 | 数据类型 | 说明                 |
  | ------ | -------- | -------------------- |
  | result | String   | 返回票务信息查询状态 |

  返回值示例说明：

  * "success" - 检票成功
  * "invalid" - 无效的票（次数为0无法继续使用）
  * "fake" - 假票（不存在数据库内的票）
  * "fuckyou" - 证件码进入检票口，票码进入证件检票口

  

### Staff证件 核销接口

- URL： `/staff`

- 功能：

  查询票务扫描信息

- 请求报文：`POST /staff`

  ```json
  Content-Type: "application/json"
  {	
      "token":"76a9408e-a8ca-4f37-a27c-2bcf1f1d1cef",
      "key":"101691071aad710c58a4bf72f4802180"
  }
  ```

  | 字段名 | 数据类型 | 说明           |
  | ------ | -------- | -------------- |
  | token  | string   | 鉴权、认证     |
  | key    | string   | 进行查询的票码 |

  

- 响应报文：200

  ```json
  Content-Type: "application/json"
  {	
      "result":"success"
  }
  ```

  | 字段名 | 数据类型 | 说明                 |
  | ------ | -------- | -------------------- |
  | result | String   | 返回票务信息查询状态 |

  返回值示例说明：

  - "success" - 检票成功
  - "invalid" - 无效的票（次数为0无法继续使用）
  - "fake" - 假票（不存在数据库内的票）
  - "fuckyou" - 证件码进入检票口，票码进入证件检票口



### 查询票务信息 - 查询票务日志 

- URL： `/query`

- 功能：

  查询票务扫描信息

- 请求报文：`POST /staff`

  ```json
  Content-Type: "application/json"
  {	
      "key":"101691071aad710c58a4bf72f4802180"
  }
  ```

  | 字段名 | 数据类型 | 说明           |
  | ------ | -------- | -------------- |
  | key    | string   | 进行查询的票码 |

  

- 响应报文：200

  ```json
  Content-Type: "application/json"
  {	
      "result":"success",
      "key":"301691071aad710c58a4bf72f4802180",
      "logs":[
          {
              "time":"2018-06-18 00:00:00",
              "result":0,
          },
          {
              "time":"2018-06-18 00:00:00",
              "result":0,
          },
          {
              "time":"2018-06-18 00:00:00",
              "result":1,
          }
      ]
  }
  ```





