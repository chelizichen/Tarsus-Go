package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func printBody(r *http.Response)  {
	content,err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	println(content)
}

func GetDemo(){
	request,err := http.NewRequest(http.MethodGet,"http://httpbin.org/get",nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent","chrome")

	values := make(url.Values)
	values.Add("NAME","LEEMULUS")
	values.Add("AGE","22")
	// 设置插叙你参数
	request.URL.RawQuery = values.Encode()

	do, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	all, err := ioutil.ReadAll(do.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s",all)
}

func PostDemo(){
	request,err := http.NewRequest(http.MethodPost,"http://httpbin.org/post",nil)
	if err != nil {
		panic(err)
	}
	values := make(url.Values)
	request.Header.Add("User-Agent","chrome")
	values.Add("NAME","LEEMULUS")
	values.Add("AGE","22")
	// 设置插叙你参数
	//request.Body = values

	do, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	all, err := ioutil.ReadAll(do.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s",all)
}

func main(){
	GetDemo()
}