// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2020/10/28 8:12
// @Update  姓名（需要改）  2020/10/28 8:12
package utils

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"os/exec"
	"strings"
	"time"
)

var successChan = make(chan bool) //判断是否退出

// 尝试密码暴力破解
// @param	zipFile: 带破解文件
// @param	passWords: 密码可选字符
// @param 	digits: 密码位数
func DoUnZip(zipFile string, passWords string, min_digits, max_digits int, outPath string) {
	log.Println("----- 密码破解开始 -----")

	beginTime := time.Now()

	// 创建线程池并设定大小，线程任务为尝试解压
	poolSize := 10000
	routinePool, _ := ants.NewPoolWithFunc(poolSize, func(zipFile_pass interface{}) {
		doUnZipTry(zipFile_pass)
		//wg.Done()
	})
	defer routinePool.Release()

	Loop1:
		for digits := min_digits; digits <= max_digits; digits++ {
			log.Printf("==> 开始 %d 位密码破解尝试...\n", digits)
			// 创建密码生成器
			pg := NewPassListGener(passWords, digits)
			outChan := pg.OutChan()
			overChan := pg.OverChan()
			// 启动密码生成 goroutine
			go pg.GetPassListByChan()

			// 循环密码，尝试破解
			Loop2:
				for {
					select {
					// 接收到密码则启动 goroutine 尝试解压
					case pass := <-outChan:
						routinePool.Invoke([]string{zipFile, pass, outPath})
						// 所有密码均尝试完成，则退出
					case <-overChan:
						log.Printf("==> 所有 %d 位密码均已尝试...\n", digits)
						break Loop2
						// 如果成功解压，则退出
					case <-successChan:
						log.Println("==> 已成功解压...")
						break Loop1
					}
				}
		}
	timeUse := time.Since(beginTime)
	log.Println("---- 工作已完成，共用时：", timeUse, ",程序退出！ ----")
}

// @param	zipFile_pass: string数组，一共两个元素，[]string{zipFile, pass}
func doUnZipTry(zipFile_pass interface{}) {
	zfp := zipFile_pass.([]string)
	cmdshell7zip(zfp[0], zfp[1], zfp[2])
}

// 调用 7zip 命令行，尝试破解
// 需要本机安装 7zip
func cmdshell7zip(zipFile string, pass string, outPath string) {
	cmd := exec.Command("7z", "e", "-aot", "-o"+outPath, "-p"+pass, zipFile)
	out, _ := cmd.Output()
	successFlag := "Everything is Ok"
	outStr := string(out)

	if strings.Contains(outStr, successFlag) {
		fmt.Println("---------- OK --------")
		fmt.Println("Pass: ", pass)
		fmt.Println(outStr)
		successChan <- true
	}
}
