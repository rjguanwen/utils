// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2020/10/27 17:48
// @Update  姓名（需要改）  2020/10/27 17:48
package utils

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestNewPassListGener(t *testing.T) {
	pg := NewPassListGener("AB", 3)
	fmt.Println("密码生成器：", pg)
}

func TestPassListGener_GetPassListByArray(t *testing.T) {
	fmt.Println("开始---")
	beginTime := time.Now()
	pg := NewPassListGener("ABC", 3)
	outArr := pg.GetPassListByArray()
	timeUse := time.Since(beginTime)
	fmt.Println(outArr)
	fmt.Println("执行时长：", timeUse)
	fmt.Println("生成密码个数：", pg.counter)
}

func TestPassListGener_GetPassListByFile(t *testing.T) {
	fmt.Println("开始---")
	beginTime := time.Now()
	pg := NewPassListGener("ABCD", 3)
	filePath := "c:/tmp/pass-1.txt"
	err := pg.GetPassListByFile(filePath)
	if err != nil {
		log.Fatalf("写入密码文件出错(%s)：%v", filePath, err)
	}
	timeUse := time.Since(beginTime)
	fmt.Println("执行时长：", timeUse)
	fmt.Println("生成密码个数：", pg.counter)
}

func TestPassListGener_GetPassListByChan(t *testing.T) {
	fmt.Println("开始---")
	beginTime := time.Now()
	pg := NewPassListGener("012", 3)
	outChan := pg.OutChan()
	overChan := pg.OverChan()
	go pg.GetPassListByChan()
Loop:
	for {
		select {
		case pass := <-outChan:
			fmt.Println(pass)
		case <-overChan:
			break Loop
		}
	}
	timeUse := time.Since(beginTime)
	fmt.Println("执行时长：", timeUse)
	fmt.Println("生成密码个数：", pg.counter)
}
