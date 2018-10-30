package fetcher

import (
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
	"bufio"
	"log"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("Wrong status code: %d",
				resp.StatusCode)
	}

	e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(
		resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	// Reader实现了给一个io.Reader接口对象附加缓冲。
	// func (*Reader) Peek 
	// func (b *Reader) Peek(n int) ([]byte, error)
	// Peek返回输入流的下n个字节，而不会移动读取位置。
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}