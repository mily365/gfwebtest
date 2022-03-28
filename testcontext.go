package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

type MyContext struct {
	// 这里的 Context 是我 copy 出来的，所以前面不用加 context.
	context.Context
}
type MyContext2 struct {
	// 这里的 Context 是我 copy 出来的，所以前面不用加 context.
	T
}
type T struct {
}

func newMyCt2() *MyContext2 {
	return &MyContext2{T{}}
}
func main() {
	childCancel := true

	parentCtx, parentFunc := context.WithCancel(context.Background())
	mctx := MyContext{parentCtx}

	childCtx, childFun := context.WithCancel(mctx)

	if childCancel {
		childFun()
	} else {
		parentFunc()
	}
	fmt.Println(reflect.TypeOf(parentCtx))
	fmt.Println(parentCtx)
	fmt.Println(mctx)
	fmt.Println(childCtx)
	ct2 := newMyCt2()
	fmt.Println(ct2)

	// 防止主协程退出太快，子协程来不及打印
	time.Sleep(10 * time.Second)
}
