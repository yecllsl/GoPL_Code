// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 130.

// The wait program waits for an HTTP server to start responding.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//!+
// WaitForServer attempts to contact the server of a URL.尝试连接URL对应的服务器
// It tries for one minute using exponential back-off.在一分钟内使用指数退避策略进行重试
// It reports an error if all attempts fail.所有尝试失败后返回错误
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off指数退避策略。超出一定时间后或次数后再退出
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

//!-

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)//优雅的停止程序
	}
	url := os.Args[1]
	//!+main
	// (In function main.)
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
	//!-main
}
