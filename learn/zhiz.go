package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func (e *Books) name() {
	e.title = "qweasd"
}

func main2() {
	var b1 Books /* 声明 b1 为 Books 类型 */
	var b2 Books /* 声明 b2 为 Books 类型 */

	/* book 1 描述 */
	b1.title = "Go 语言"
	b1.author = "www.runoob.com"
	b1.subject = "Go 语言教程"
	b1.book_id = 6495407

	/* book 2 描述 */
	b2.title = "Python 教程"
	b2.author = "www.runoob.com"
	b2.subject = "Python 语言教程"
	b2.book_id = 6495700

	fmt.Println(b1)
	fmt.Println(b2)

	b1.name()

	fmt.Println(b1)
	fmt.Println(b2)
}
