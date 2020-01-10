package main

import (
	"fmt"
	"net/url"
)

func main() {
	//var testUrl string = "https://www.baidu.com:59/s?wd=%E5%A5%A5%E8%BF%AA&rsv_spt=1"
	//rurl, err := url.Parse(testUrl)
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(url.)

	//fmt.Printf("%v\n", rurl)
	//fmt.Println(url.ParseQuery(rurl.RawQuery))
	//queryMap := rurl.Query()
	//fmt.Println(queryMap.Get("wd"))

	//var testUrl string = "http://fangkm.com/get b/?a=2&b=方开名 奥迪"
	//fmt.Println(url.QueryEscape(testUrl))
	//parseUrl, _ := url.Parse(testUrl)
	//fmt.Println(url.QueryEscape(parseUrl.RawQuery))


	//urlPath := parseUrl.Path
	//fmt.Println(urlPath)
	//fmt.Println(url.Parse(testUrl))

	//userInfo :=url.User("fangkaiming")
	//fmt.Println(userInfo.Username())
	//fmt.Println(userInfo.Password())

	values := url.Values{}
	values.Add("a", "3")
	values.Add("b", "方开名 百度ab")
	fmt.Println(values.Encode())
}