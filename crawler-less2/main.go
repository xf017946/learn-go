package main

import (
	// net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket。
	"net/http"
	"io/ioutil"
	"io"
	"bufio"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"regexp"
)

func main () {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",
			resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(
		resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
	
	printCityList(all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	// Reader实现了给一个io.Reader接口对象附加缓冲。
	// func (*Reader) Peek 
	// func (b *Reader) Peek(n int) ([]byte, error)
	// Peek返回输入流的下n个字节，而不会移动读取位置。
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}

// 城市列表解析器
func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n",
		m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n",
			len(matches))
}