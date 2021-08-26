/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/25 21:50
 */
package main

import (
	"io/ioutil"
	"log"
	"os"

	_ "kirito-layout/ecode"

	"github.com/go-kirito/pkg/errors"
)

func main() {

	md, err := errors.GenErrCodeDoc("app")
	if err != nil {
		log.Panic(err)
	}

	w, err := os.Getwd()

	if err != nil {
		log.Panic(err)
	}

	filename := w + "/doc/errors/errors.md"

	err = ioutil.WriteFile(filename, md, 0644)

	if err != nil {
		log.Panic(err)
	}

	log.Println("生成errors文档成功")
}
