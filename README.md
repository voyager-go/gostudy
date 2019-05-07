# gostudy
## 学习golang的一些编程基础
### 在Linux系统中安装Golang
#### 首先我们找到golang的官方下载地址https://golang.google.cn/dl/
#### 然后我们找到对应的Linux tar.gz包 下载
#### 下载后提取go文件夹tar -xzf go1.12.4.linux-amd64.tar.gz // 注意修改为自己的包名称 或者直接使用归档器打开提取
#### 然后把go这个文件夹转移到/usr/local/文件夹下面 sudo mv go/ /usr/local/go
#### 找个自己存放go项目的文件夹 如 ~/project/go
#### 在这个go文件夹下面新建三个文件夹 pkg bin src 可以使用命令 mkdir pkg bin src
#### 然后我们给单个用户进行设置环境变量
#### vim ~/.bashrc
```
export GOPATH=~/project/go
export GOROOT=/usr/local/go
export GOARCH=386
export GOOS=linux
export GOBIN=$GOROOT/bin/
export GOTOOLS=$GOROOT/pkg/tool/
export PATH=$PATH:$GOBIN:$GOTOOLS
```
#### :wq保存退出 接着使环境变量立即生效 source ~/.bashrc
#### 这时我们使用go version // go version go1.12.4 linux/amd64
#### 使用go env可以查看更多的信息
