package main

import (
	"net/http"
	"github.com/gorilla/context"
	"strconv"
	"sync"
)

/* desc
** context 包可以在一个 request 的处理期间，保存一些值
** func Clear(r *http.Request)
** func ClearHandler(h http.Handler) http.Handler
** func Delete(r *http.Request, key interface{})
** func Get(r *http.Request, key interface{}) interface{}
** func GetAll(r *http.Request) map[interface{}]interface{}
** func GetAllOk(r *http.Request) (map[interface{}]interface{}, bool)
** func GetOk(r *http.Request, key interface{}) (interface{}, bool)
** func Purge(maxAge int) int
** func Set(r *http.Request, key, val interface{})
*/
func main() {
	// 启动一个Web服务
	http.Handle("/", http.HandlerFunc(MyHandler))
	http.ListenAndServe(":9999", nil)
}
var mutex sync.RWMutex
const MyKey int = 0
func MyHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许一个Go协程 避免竞态条件
	mutex.Lock()
	context.Set(r, MyKey, "bar")
	// 删除MyKey
	context.Delete(r, MyKey)
	if  rv := context.Get(r, MyKey); rv != nil{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The MyKey " + strconv.Itoa(MyKey) + " is " + rv.(string)))
	}else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The MyKey " + strconv.Itoa(MyKey) + " is nothing"))
	}
	mutex.Unlock()
}
