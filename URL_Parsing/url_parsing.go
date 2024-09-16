package main

import (
	"fmt"
	"net"
	"net/url"
)

/*
	โค้ดตัวอย่างนี้แสดงการทำงานของการ "URL Parsing" หรือการแยก

องค์ประกอบต่างๆ ของ URL โดยใช้แพ็คเกจ net/url ของภาษา Go
ซึ่งมีการแยกส่วนต่างๆ ของ URL ออกมา เช่น scheme, user, info, host, port, path,
query parameters, และ fragment
*/
func main() {
	s := "postgres://user:5bJwO@example.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	f(u.Scheme)

	f(u.User)
	f(u.User.Username())
	p, _ := u.User.Password()
	f(p)

	f(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	f(host)
	f(port)

	f(u.Path)
	f(u.Fragment)

	f(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	f(m)
	f(m["k"][0]) //
}

func f(s ...interface{}) {
	fmt.Println(s...)
}
