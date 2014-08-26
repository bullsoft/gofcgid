package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/guweigang/fcgiclient"
)

func main() {
	postParams := "name=value"
	queryParams := "foo=bar&hello=world"

        env := make(map[string]string)
        env["REQUEST_METHOD"] = "POST"
        env["SCRIPT_FILENAME"] = "/Users/guweigang/wwwroot/test.php"
        env["SERVER_SOFTWARE"] = "go / fcgiclient "
        env["REMOTE_ADDR"] = "127.0.0.1"
        env["SERVER_PROTOCOL"] = "HTTP/1.1"
        env["QUERY_STRING"] = queryParams
	env["CONTENT_TYPE"] = "application/x-www-form-urlencoded"
	env["CONTENT_LENGTH"] = strconv.FormatInt(10, 10)

        fcgi, err := fcgiclient.New("127.0.0.1", 9000)
        if err != nil {
                fmt.Printf("err: %v", err)
        }

        content, err := fcgi.Request(env, postParams)
        if err != nil {
                fmt.Printf("err: %v", err)
        }
	
	data := strings.SplitAfter(string(content[:]), "\r\n\r\n")
	fmt.Printf("%v\n", data[1])
	// fmt.Printf("content: %s", content)

}
