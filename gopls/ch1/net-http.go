// net/http 简单使用


package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	// handler 可以是nil，原因是net/http为了方便，定义了一个默认处理器DefaultServeMux
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// GET http://localhost:8000/?q=1&b=2&b=3
//GET /?q=1&b=2&b=3 HTTP/1.1
//Header["Sec-Ch-Ua-Mobile"] = ["?0"]
//Header["Sec-Fetch-Mode"] = ["navigate"]
//Header["Sec-Fetch-Dest"] = ["document"]
//Header["Accept-Encoding"] = ["gzip, deflate, br"]
//Header["Sec-Ch-Ua"] = ["\"Not_A Brand\";v=\"99\", \"Google Chrome\";v=\"109\", \"Chromium\";v=\"109\""]
//Header["Connection"] = ["keep-alive"]
//Header["Sec-Ch-Ua-Platform"] = ["\"macOS\""]
//Header["Upgrade-Insecure-Requests"] = ["1"]
//Header["User-Agent"] = ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"]
//Header["Sec-Fetch-Site"] = ["none"]
//Header["Accept-Language"] = ["en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"]
//Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"]
//Header["Sec-Fetch-User"] = ["?1"]
//Host = "localhost:8000"
//RemoteAddr = "127.0.0.1:59039"
//Form["b"] = ["2" "3"]
//Form["q"] = ["1"]