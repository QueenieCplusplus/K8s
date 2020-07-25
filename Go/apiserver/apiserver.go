// 2020, 7/25, am 9:55
// mux is a type that can register http handler

package main

import "net/http"

// 接手可以放入任何資料型別的資料
type Mux interface {
	Handle(pattern string, handler http.Handler)
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}
