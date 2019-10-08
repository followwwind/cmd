// Copyright 2017 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/go-xorm/core"
)

type LangTmpl struct {
	Funcs      template.FuncMap
	Formater   func(string) (string, error)
	GenImports func([]*core.Table) map[string]string
}

var (
	mapper    = &core.SnakeMapper{}
	langTmpls = map[string]LangTmpl{
		"go":   GoLangTmpl,
		"c++":  CPlusTmpl,
		"objc": ObjcTmpl,
	}
)

func loadConfig(f string) map[string]string {
	bts, err := ioutil.ReadFile(f)
	if err != nil {
		return nil
	}
	configs := make(map[string]string)
	lines := strings.Split(string(bts), "\n")
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		vs := strings.Split(line, "=")
		if len(vs) == 2 {
			configs[strings.TrimSpace(vs[0])] = strings.TrimSpace(vs[1])
		}
	}
	return configs
}

func unTitle(src string) string {
	if src == "" {
		return ""
	}

	if len(src) == 1 {
		return strings.ToLower(string(src[0]))
	} else {
		return strings.ToLower(string(src[0])) + src[1:]
	}
}

func upTitle(src string) string {
	if src == "" {
		return ""
	}

	return strings.ToUpper(src)
}

//驼峰命名
func toCamel(src string) string {
	if src == "" {
		return ""
	}

	arr := strings.Split(src, "_")

	result := ""
	for i := range arr {
		result += upFirst(arr[i])
	}
	return result
}

//json,首单词小写驼峰命名
func toJsonCamel(src string) string {
	if src == "" {
		return ""
	}

	arr := strings.Split(src, "_")

	result := ""
	for i := range arr {
		if i == 0 {
			result += arr[i]
		} else{
			result += upFirst(arr[i])
		}
	}
	return result
}

//首字母大写
func upFirst(src string) string {
	if src == "" {
		return ""
	}

	if len(src) == 1 {
		return strings.ToUpper(string(src[0]))
	} else {
		return strings.ToUpper(string(src[0])) + src[1:]
	}
}
