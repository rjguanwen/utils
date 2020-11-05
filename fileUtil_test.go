// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2020/8/26 15:37
// @Update  姓名（需要改）  2020/8/26 15:37
package utils

import (
	"fmt"
	"log"
	"os"
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

func TestCheckIsHidden(t *testing.T) {
	beginTime := time.Now()
	//filePath := "c:/tmp/彩色圆环图.7z"
	//filePath := "c:/tmp/flexganttfx-11.8.1-bin.zip"
	filePath := "c:/tmp/ss"
	fi, err := os.Stat(filePath)
	if err != nil {
		log.Fatal("读出文件信息报错：", err)
	}
	isHidden := CheckIsHidden(fi)
	timeUse := time.Since(beginTime)
	fmt.Println("文件为隐藏文件：", isHidden)
	fmt.Println("方法用时：", timeUse)
}
