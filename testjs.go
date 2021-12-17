package main

import (
	"fmt"
	"github.com/dop251/goja"
)

type stu struct {
	Name string
	Age  int64
}

func (s *stu) Add(x int, y int) int {
	return x + y
}

func MyObject(call goja.ConstructorCall) *goja.Object {
	// call.This contains the newly created object as per http://www.ecma-international.org/ecma-262/5.1/index.html#sec-13.2.2
	// call.Arguments contain arguments passed to the function

	//call.This.Set("Add", Add)

	// If return value is a non-nil *Object, it will be used instead of call.This
	// This way it is possible to return a Go struct or a map converted
	// into goja.Value using ToValue(), however in this case
	// instanceof will not work as expected, unless you set the prototype:
	//
	instance := &stu{Name: "jy", Age: 15}
	instanceValue := goja.New().ToValue(instance).(*goja.Object)
	instanceValue.SetPrototype(call.This.Prototype())
	return instanceValue

}
func main() {
	// js 只认Value，r.ToValue--需要把go转换为Value
	//传递值到js vm.set
	//js 导入go vm.ExportTo/ vm.Export
	var num int64
	vm := goja.New()
	v, err := vm.RunString("2 + 2")
	if err != nil {
		panic(err)
	}
	//从js 导出值到go
	if num = v.Export().(int64); num != 4 {
		panic(num)
	}
	fmt.Println(num)

	f, err := vm.RunString(`
		function sum(a, b) {
			return a+b;
		}
    `)
	_ = f

	if err != nil {
		panic(err)
	}
	sum, ok := goja.AssertFunction(vm.Get("sum"))
	if !ok {
		panic("Not a function")
	}
	// 导入的js fun
	res, err := sum(goja.Undefined(), vm.ToValue(40), vm.ToValue(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	const SCRIPT = `
	function f(param) {
		return +param + 2;
	}
	`

	ff, err := vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}
	_ = ff
	var fn func(string) string
	err = vm.ExportTo(vm.Get("f"), &fn)
	if err != nil {
		panic(err)
	}

	fmt.Println(fn("40")) // note, _this_ value in the function will be undefined.
	// Output: 42

	//注意传递指针和不传递指针的区别
	a := []interface{}{1}
	vm.Set("a", &a)
	vm.RunString(`a.push(2); a[0] = 0;`)
	fmt.Println(a[0]) // prints "1"
	fmt.Println(a[1]) // prints "1"

	err = vm.Set("MyObject", MyObject)
	if err != nil {
		panic(err)
	}
	const SCRIPT2 = `
        var o = new MyObject();
        o.Name
        o.Add(3,5)
	`

	v, err = vm.RunString(SCRIPT2)
	if err != nil {
		panic(err)
	}
	s := v.Export().(int64)
	fmt.Println(s)
}
