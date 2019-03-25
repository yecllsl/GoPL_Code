// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 160.

// The urlvalues command demonstrates a map type with methods.
package main

/*
//!+values
package url

// Values maps a string key to a list of values.映射字符串到字符串列表
type Values map[string][]string

// Get returns the first value associated with the given key,Get返回第一个具有给定Key的值
// or "" if there are none.如不存在，则返回空字符串。
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add adds the value to key.
// It appends to any existing values associated with key.
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
//!-values
*/

import (
	"fmt"
	"net/url"
)

func main() {
	//!+main
	m := url.Values{"lang": {"en"}} // direct construction直接构造
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"      (first value)
	fmt.Println(m["item"])     // "[1 2]"  (direct map access)

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic: assignment to entry in nil map。宕机：赋值给空的map类型。
	//!-main
}
