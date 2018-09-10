package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"regexp"
)

func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil{
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte){
	re, err := regexp.Compile(
		`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)

	if err != nil{
		panic(err)
	}

	{
	//	matchs := re.FindAll(contents, -1)
		//fmt.Println(matchs)
	//	for _, m := range matchs{
	//		fmt.Printf("%s\n", m)
	//	}
	}

	{
		matchs := re.FindAllSubmatch(contents, -1)
		for _, submatch := range matchs{
			fmt.Printf("city:%s url:%s \n", submatch[2], submatch[1])
		}
		//fmt.Printf("size=%d\n", len())
	}




}




func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Error: status code: ",
			resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil{
		panic(err)
	}

	printCityList(all)

}
