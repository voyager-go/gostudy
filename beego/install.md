### 快速安装beego
### 为了方便起见你可以在go目录的src下面获取
### 请使用命令go get -v -u github.com/astaxie/beego
### 获取Bee开发工具go get -v -u github.com/beego/bee
### 有些人在这一步遇到了错误(linux系统)go get: cannot install cross-compiled binaries when GOBIN is set
#### 这里有一个解决方案 首先进入go目录也就是cd $GOPATH 然后临时删除GOBIN设置unset GOBIN
#### 此时再来执行go get -v -u github.com/beego/bee就会成功了
#### 但是还有一个问题 就是你会发现当前目录下面没有找到bee
#### 那是因为bee工具在~/go/bin/linux_386文件夹下
#### 那么我们把它添加到环境变量中
#### vim ~/.bashrc 然后把export PATH=$GOPATH/bin/linux_386:$PATH配置一下 
#### 再来执行source ~/.bashrc就可以成功了
###  没有遇到错误的
#### 如果您还没添加 $GOPATH 变量
$ echo 'export GOPATH="$HOME/go"' >> ~/.profile # 或者 ~/.zshrc, ~/.cshrc, 您所使用的sh对应的配置文件

#### 如果您已经添加了 $GOPATH 变量
$ echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.profile # 或者 ~/.zshrc, ~/.cshrc, 您所使用的sh对应的配置文件
$ exec $SHELL
### 此时使用bee version来查看bee的版本
artist@artist-pc:~/develop/go$ bee version
##### ______
##### | ___ \
##### | |_/ /  ___   ___
##### | ___ \ / _ \ / _ \
##### | |_/ /|  __/|  __/
##### \____/  \___| \___| v1.10.0

##### ├── Beego     : 1.11.2
##### ├── GoVersion : go1.12.4
##### ├── GOOS      : linux
##### ├── GOARCH    : 386
##### ├── NumCPU    : 4
##### ├── GOPATH    : /home/artist/develop/go/
##### ├── GOROOT    : /usr/local/go
##### ├── Compiler  : gc
##### └── Date      : Thursday, 18 Apr 2019
### 我们可以尝试创建一个项目并运行
#### 进入/go/src目录下
#### bee new hello
#### cd hello
#### bee run
#### 此时我们打开浏览器访问127.0.0.1:8080
#### 就可以看到请求的go页面
#### 更多内容可以到beego的官网查看https://beego.me/quickstart
