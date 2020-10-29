// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2020/8/26 15:37
// @Update  姓名（需要改）  2020/8/26 15:37
package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestReadLastLineFromFile(t *testing.T) {
	filePath := "C:/tmp/sparrow_redo/Index.redo"
	lastLine, _ := ReadLastLineFromFile(filePath, 100, true)
	fmt.Println("最后一行==>", lastLine)
}

func TestGetFileMD5(t *testing.T) {
	fmt.Println("--- 开始 ----")
	begin := time.Now()
	filePath := "C:/tmp/zip1-de.7z"
	fileMD5, err := GetFileMD5(filePath)
	if err != nil {
		fmt.Printf("获取文件 MD5 码出错：%v\n", err)
	}
	timeUse := time.Since(begin)
	fmt.Println("文件 MD5 码：", fileMD5)
	fmt.Println("共用时：", timeUse)
}
