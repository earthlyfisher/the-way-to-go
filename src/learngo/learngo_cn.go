// 单行注释
/*多行
注释 */

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"net/http"
	"io/ioutil"
	"log"
	"reflect"
)

func main() {
	fmt.Println("welcome!")

	p := pair{1, 2}
	//类型断言
	typeAssertion(p)

	//beyondHello()
}

func beyondHello() {

	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "string type"
	s2 := `这是一个
	可以换行的字符串` // 同样是String类型

	g := 'x'

	f := 3.14195
	c := 3 + 4i

	var u uint = 7
	var pi float32 = 22. / 7

	n := byte('\n')

	//数组
	var a4 [4]int
	a3 := [...]int{2, 1, 3}

	//slice
	s3 := []int{2, 3, 4}
	s4 := make([]int, 4)
	var d2 [][]float64
	bs := []byte("a slice")

	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	// 除了向append()提供一组原子元素（写死在代码里的）以外，我们
	// 还可以用如下方法传递一个slice常量或变量，并在后面加上省略号，
	// 用以表示我们将引用一个slice、解包其中的元素并将其添加到s数组末尾。
	s = append(s, []int{7, 8, 9}...)
	fmt.Println(s)

	p, q := learnMemory()
	fmt.Println(*p, *q)

	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	// 在Go语言中未使用的变量在编译的时候会报错，而不是warning。
	// 下划线 _ 可以使你“使用”一个变量，但是丢弃它的值。
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a3, s4, bs

	// 通常的用法是，在调用拥有多个返回值的函数时，
	// 用下划线抛弃其中的一个参数。下面的例子就是一个脏套路，
	// 调用os.Create并用下划线变量扔掉它的错误代码。
	// 因为我们觉得这个文件一定会成功创建。
	file, _ := os.Create("D:\\123.txt")
	fmt.Fprint(file, "这句话示范如何写入文件")
	file.Close()

	// 输出变量
	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl()
}

func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return z
}

func learnMemory() (p, q *int) {
	p = new(int)
	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r

}

func expensiveComputation() int {
	return 1e6
}

func learnFlowControl() {

	if true {
		fmt.Println("这句话肯定被执行")
	}

	if false {

	} else {

	}

	x := 1
	switch x {
	case 0:
	case 1:
		// 隐式调用break语句，匹配上一个即停止
	case 2:
		// 不会运行
	}

	for x := 0; x < 3; x++ {
		fmt.Println("遍历", x)
	}

	x = 0
	for x < 3 {
		x = x + 1
	}

	// 如果你只想要值，那就用前面讲的下划线扔掉没用的
	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("你是。。 %s\n", name)
	}

	// 和for一样，if中的:=先给y赋值，然后再和x作比较。
	if y := expensiveComputation(); y > x {
		x = y
	}

	// 闭包函数
	xBig := func() bool {
		return x > 100 // x是上面声明的变量引用
	}
	fmt.Println("xBig:", xBig()) // true （上面把y赋给x了）
	x /= 1e5                     // x变成10
	fmt.Println("xBig:", xBig()) // 现在是false

	// 除此之外，函数体可以在其他函数中定义并调用，
	// 满足下列条件时，也可以作为参数传递给其他函数：
	//   a) 定义的函数被立即调用
	//   b) 函数返回值符合调用者对类型的要求
	fmt.Println("两数相加乘二: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2)) // Called with args 10 and 2
	// => Add + double two numbers: 24

	goto love

love:

	learnFunctionFactory() // 返回函数的函数多棒啊
	learnDefer()           // 对defer关键字的简单介绍
	learnInterfaces()      // 好东西来了！
}

func learnFunctionFactory() {
	// 空行分割的两个写法是相同的，不过第二个写法比较实用
	fmt.Println(sentenceFactory("原谅")("当然选择", "她！"))

	d := sentenceFactory("原谅")
	fmt.Println(d("当然选择", "她！"))
	fmt.Println(d("你怎么可以", "她？"))
}

// Decorator在一些语言中很常见，在go语言中，
// 接受参数作为其定义的一部分的函数是修饰符的替代品
func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
	}
}

func learnDefer() (ok bool) {
	// defer表达式在函数返回的前一刻执行
	defer fmt.Println("defer表达式执行顺序为后进先出（LIFO）")
	defer fmt.Println("\n这句话比上句话先输出，因为")
	// 关于defer的用法，例如用defer关闭一个文件，
	// 就可以让关闭操作与打开操作的代码更近一些
	return true
}

type Stringer interface {
	String() string
}

type pair struct {
	x, y int
}

func (p pair) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func learnInterfaces() {
	p := pair{1, 2}
	fmt.Println(p.String())

	var i Stringer
	i = p
	fmt.Println(i.String())

	fmt.Println(p)
	fmt.Println(i)

	learnVariadicParams("great", "learning", "here!")
}

// 有变长参数列表的函数
func learnVariadicParams(myStrings ...interface{}) {
	// 枚举变长参数列表的每个参数值
	// 下划线在这里用来抛弃枚举时返回的数组索引值
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	// 将可变参数列表作为其他函数的参数列表
	fmt.Println("params:", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	// ", ok"用来判断有没有正常工作
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok { // ok 为false，因为m中没有1
		fmt.Println("别找了真没有")
	} else {
		fmt.Print(x) // 如果x在map中的话，x就是那个值喽。
	}
	// 错误可不只是ok，它还可以给出关于问题的更多细节。
	if _, err := strconv.Atoi("non-int"); err != nil { // _ discards value
		// 输出"strconv.ParseInt: parsing "non-int": invalid syntax"
		fmt.Println(err)
	}

	learnConcurrency()
}

func inc(i int, c chan int) {
	c <- i + 1
}

func learnConcurrency() {
	c := make(chan int)

	go inc(0, c)
	go inc(10, c)
	go inc(-805, c)
	// 从channel中读取结果并打印。
	// 打印出什么东西是不可预知的。
	fmt.Println(<-c, <-c, <-c) // channel在右边的时候，<-是读操作。

	cs := make(chan string)
	cc := make(chan chan string)

	go func() {
		time.Sleep(10) // sleep one second
		c <- 84
	}()

	go func() {
		cs <- "wordy"
	}()

	// Select类似于switch，但是每个case包括一个channel操作。
	// 它随机选择一个准备好通讯的case。
	select {
	case i := <-c:
		fmt.Println("这是……", i)
	case z := <-cs: // 或者直接丢弃
		fmt.Println("这是个字符串:", z)
	case <-cc: // 空的，还没作好通讯的准备
		fmt.Println("别瞎想")
	}
	// 上面c或者cs的值被取到，其中一个goroutine结束，另外一个一直阻塞。

	learnWebProgramming()
}

func learnWebProgramming() {
	go func() {
		requestServer()
	}()

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello,this version 1!"))
	})
	http.Handle("/bye", pair{})

	log.Println("Starting server...v1")
	http.ListenAndServe(":8080", nil)
	//requestServer()
}

// 使pair实现http.Handler接口的ServeHTTP方法。
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 使用http.ResponseWriter返回数据
	w.Write([]byte("Y分钟golang速成!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080/bye")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\n服务器消息： `%s`", string(body))

	resp, err = http.Get("http://localhost:8080/test")
	body, err = ioutil.ReadAll(resp.Body)
	fmt.Printf("\n服务器消息test： `%s`", string(body))

}

func typeAssertion(object interface{}) {
	value, ok := object.(Stringer)
	if ok {
		fmt.Printf("接口类型断言成功，类型为 <%v> ", reflect.TypeOf(value))
	} else {
		fmt.Println("接口类型断言失败")
	}
}
