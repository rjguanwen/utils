// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2020/10/29 14:46
// @Update  姓名（需要改）  2020/10/29 14:46
package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestDoUnZip(t *testing.T) {
	fmt.Println("开始---")
	beginTime := time.Now()
	// 待破解文件
	zipFile := "C:/tmp/zip1-de.7z" // rar 文件路径
	passWords := "0123456789"
	outPath := "c:/tmp/rar2"

	// 开始破解
	DoUnZip(zipFile, passWords, 1, 4, outPath)

	timeUse := time.Since(beginTime)
	fmt.Println("执行时长：", timeUse)
}