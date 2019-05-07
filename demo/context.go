package main

import (
	"net/http"
	"github.com/gorilla/context"
	"strconv"
)
// 在一个 request 的处理期间，保存一些值
func main() {
	// 启动一个Web服务
	http.Handle("/", http.HandlerFunc(myHandle))
	http.ListenAndServe(":9999", nil)
}

func myHandle(rw http.ResponseWriter, r *http.Request)  {
	// 模拟为Request附加值
	context.Set(r, "user", "张大炮")
	context.Set(r, "age", 23)
	context.Set(r, "job", "programmer")
	context.Set(r, "gender", "man")
	// 这个模拟一个方法或者函数的调用，大部分情况下可能不在一个包里
	doHandle(rw, r)
}
func doHandle(rw http.ResponseWriter, r *http.Request)  {
	user := context.Get(r, "user").(string)
	age  := context.Get(r, "age").(int)
	job  := context.Get(r, "job").(string)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("this user is " + user + ", age is " + strconv.Itoa(age) + " job is " + job))
}