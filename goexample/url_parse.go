package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	//urlStr := "postgres://user:pass@host.com:5432/path?k=v#f"
	urlStr1 := "https://gobyexample.com/url-parsing"

	u, err := url.Parse(urlStr1)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	if u.User != nil {
		fmt.Println(u.User.Username())

		p, _ := u.User.Password()
		fmt.Println(p)
	} /* else {
		panic("use is empty")
	}*/

	fmt.Println(u.Host)
	host, post, _ := net.SplitHostPort(u.Host)
	if len(host) > 0 {
		fmt.Println(host)
	}

	if len(post) > 0 {
		fmt.Println(post)
	}

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	/*	m, _ := urlStr1.ParseQuery(u.RawQuery)
		fmt.Println(m)
		fmt.Println(m["k"][0])
	*/
}
