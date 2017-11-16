# myapp
## 功能介绍
### 用户登录
用户通过输出用户名和密码点击SIGN_IN进行登录，如果用户是新用户，可以点击new user and I want to regist进入注册页面进行注册。如果输入的用户名或者密码不正确，页面会跳转至错误页面提示err并在三秒后返回到登录页面。
<br> ![Image text]( https://github.com/muhongwei/myapp/blob/master/static/picture/login.png)<br>
### 用户注册
用户通过输入用户名，密码和自我介绍进行注册。注册成功后页面会跳转是登录界面，用户输入注册的用户名和密码进行登录。如果注册失败，页面会跳转至错误页面提示err并在三秒后返回登录页面。
<br> ![Image text]( https://github.com/muhongwei/myapp/blob/master/static/picture/regist.png)<br>
### 用户登录信息显示
用户通过输入正确的用户名和密码进入用户登录信息显示页面。页面中会显示访问用户的useragent信息，访问用户的ip信息，服务端主机名，访问用户的名字，访问用户的自我介绍。
<br> ![Image text]( https://github.com/muhongwei/myapp/blob/master/static/picture/showInfo.png)<br>
### 产品名称和版本号信息显示在每一个页面
在用户登录页面，用户注册页面以及用户登录信息显示页面均通过调用模板函数在页面右上角显示应用名称和版本号。
### echo命令行服务
通过客服端访问服务端grpc开放的端口来实现echo服务。连接成功后会提示“输入exit退出：”信息。之后通过客服端输入信息与服务端交互返回客服端输入信息，并且服务端会在连接建立之后每20秒主动推送当前时间到客服端并显示出来。
<br> ![Image text]( https://github.com/muhongwei/myapp/blob/master/static/picture/client.png)<br>
## 接口说明
### grpc接口调用
在实现命令行echo服务的时候调用的grpc的接口，通过grpc.NewServer()来新建一个grpc服务，随时监听端口号的连接。
### mysql接口调用
通过导入_ "github.com/Go-SQL-Driver/MySQL" 驱动，在beego框架models模块建立数据库连接，和sql命令执行查询操作远程数据库。
## 部署方法
### myapp部署方法
编写dockerfile和makefiel文件，在linux环境makefile同级目录执行make package对代码进行进行build生成可执行的二进制文件并通过dockerfile打包成镜像。执行ecos平台上上传镜像的代码将镜像上传到ecos平台。
进入ecos平台进入镜像管理可以查看已经拥有的镜像列表，点击添加镜像可以弹出三种不同的添加方法，上一步是在在linux平台上的centos虚拟机进行的为第三种添加方法。 
在ecos平台进入应用管理模块点击添加应用，设置应用的名称和应用描述。点击应用，添加服务设置相应的信息，点击下一步进入如下所示界面：定义容器和服务的端口映射
<br> ![Image text]( https://github.com/muhongwei/myapp/blob/master/static/picture/myappPort.png)<br>
点击下一步，设置容器运行的名称，设置前面上传的镜像。我的应用需要加启动命令：./myapp -logtostderr=true。可以在此处设置环境变量。我添加了三个环境变量：OEM,VER,GO15VENDOREXPERIMENT。OEM定义软件的名字，VER定义软件的版本号，GO15VENDOREXPERIMENT是govendor所需要的环境变量，值为1。如下图所示：
点击下一步，添加。
<br>![Image text]( https://github.com/muhongwei/myapp/blob/master/static/picture/myappConfig.png)<br>
### mysql部署方法
首先新建一台虚拟机作为nfs挂载的服务端配置nfs挂载相关设置，在ecos中储存模块将nfs主机添加上去。
通过docker pull mysql将mysql镜像拉取到本地，并通过上述方法将镜像上传到ecos。类似的添加应用和服务。在基础设置那里需要将服务状态设置为不共享磁盘，容器配置那里需要添加环境变量MYSQL_ROOT_PASSWORD初始化mysql的root账户密码，以及添加网络挂载卷。
### 添加负载均衡
在应用管理模块下有一个负载均衡模块，可以在此处添加转发规则将集群内部的端口映射到某一台主机上，我的设置如下图：
<br>![Image text]( https://github.com/muhongwei/myapp/blob/master/static/picture/myappser.png)<br>
接下来就可以通过负载均衡生成的网址访问myapp了。
## 遇到的问题和解决方法