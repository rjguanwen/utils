// @Title  PassListGener
// @Description  密码表生成器
// @Author  zhengbin  2020/10/23 17:04
// @Update  姓名（需要改）  2020/10/23 17:04
package utils

import (
	"bufio"
	"log"
	"math"
	"os"
)

// 结构体，密码生成器
type PassListGener struct {
	words string
	// 开始字符与结束字符暂不使用
	startWords string
	endWords   string
	digits     int
	outChan    chan string
	overChan   chan bool
	//// 计数器，记录一共生成了多少个密码
	counter int
}

// 创建密码表生成器
func NewPassListGener(words string, digits int) PassListGener {
	pg := PassListGener{}
	pg.words = words
	//pg.startWords = startWords
	//pg.endWords = endWords
	pg.digits = digits
	pg.outChan = make(chan string)
	pg.overChan = make(chan bool)
	pg.counter = 0
	return pg
}

func (pg PassListGener) OutChan() chan string {
	return pg.outChan
}

func (pg PassListGener) OverChan() chan bool {
	return pg.overChan
}

func (pg PassListGener) Counter() int {
	return pg.counter
}

// 通过 chan 获取密码列表
// 根据指定的位数与可选字符，生成密码表
// 此方法用了一些小技巧，通过将可选字符进行笛卡尔积可以得到所需结果
// 而可选字符的索引的笛卡尔积正好是 N 进制数的表示
// 所以可以考虑通过索引反向拼接得到元素值
// 而索引的计算可以通过十进制索引向 N 进制数的转换来快速达成
func (pg *PassListGener) GetPassListByChan() {
	//pg.counter = 0
	// 将可选字符转换为数组，方便后续操作
	wordsArr := make([]string, 0)
	for _, v := range pg.words {
		wordsArr = append(wordsArr, string(v))
	}
	// 可选字符长度
	wordsLen := len(wordsArr)
	// 返回字符串数组长度
	// 因为是笛卡尔积，所以结果长度是可选字符长度的 N 次方
	newWordsLen := math.Pow(float64(wordsLen), float64(pg.digits))
	// 通过结果数组的索引计算得到其 N 进制数的表示，
	// 通过每个表示拼接得到结果数组的每个元素
	for i := 0; i < int(newWordsLen); i++ {
		// 将结果数组下标转换为相应的 N 进制，以数组方式存储
		decimalArr := tenToAnyDecimal(i, wordsLen)
		// 对数组进行反转，以方便补位
		decimalArr = reverseArr(decimalArr)

		var passStr string
		// 因为对数组进行了反转，所以索引从大到小进行循环
		for j := pg.digits - 1; j >= 0; j-- {
			passStr = passStr + wordsArr[getIntValueByIndex(decimalArr, j)]
		}
		pg.counter = pg.counter + 1
		//fmt.Println(pg.counter)
		pg.outChan <- passStr
	}
	// 写入完成标志
	pg.overChan <- true
}

// 生成密码列表并写入文件，每行以“\n”结尾
func (pg *PassListGener) GetPassListByFile(filePath string) (err error) {
	//// 创建文件夹
	//err = os.MkdirAll(filePath, os.ModePerm)
	//if err != nil {
	//	log.Fatalf("创建日志目录出错(%s)：%v", filePath, err)
	//	return
	//}
	// 打开密码列表文件
	outFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("无法打开文件(%s)：%v", filePath, err)
		return
	}
	defer outFile.Close()
	// 开始生成密码，并放入输出管道
	go pg.GetPassListByChan()

	fileWriter := bufio.NewWriter(outFile)
	// 从管道中获取密码并写入文件
	// 接收到完成标志则退出
Loop:
	for {
		select {
		case pass := <-pg.outChan:
			fileWriter.WriteString(pass)
			fileWriter.WriteString("\n")
		case <-pg.overChan:
			break Loop
		}
	}
	fileWriter.Flush()

	return nil
}

// 生成密码列表，直接以数组的形式返回
// 该方法不适合结果集过于大的情况
func (pg *PassListGener) GetPassListByArray() (outArray []string) {
	// 开始生成密码，并放入输出管道
	go pg.GetPassListByChan()
	// 从输出管道中选好获取密码并放入结果数组
Loop:
	for {
		select {
		case pass := <-pg.outChan:
			outArray = append(outArray, pass)
		case <-pg.overChan:
			break Loop
		}
	}
	return outArray
}

// 根据索引，获取数值，超出索引的直接返回 0
func getIntValueByIndex(originArr []int, index int) int {
	if index >= len(originArr) {
		return 0
	} else {
		return originArr[index]
	}
}

// 数组反转，将数组整个反转
func reverseArr(originArr []int) []int {
	length := len(originArr)
	for i := 0; i < length/2; i++ {
		originArr[i], originArr[length-i-1] = originArr[length-i-1], originArr[i]
	}
	return originArr
}

// 将十进制数字转化为任意进制的数组，每一个元素代表一位
// 本方法只处理非负整数
func tenToAnyDecimal(num, n int) (decimalArr []int) {
	if num == 0 {
		decimalArr = append(decimalArr, 0)
		return decimalArr
	}
	if num < 0 {
		return nil
	}
	var remainder int
	for num != 0 {
		remainder = num % n
		decimalArr = append([]int{remainder}, decimalArr...)
		num = num / n
	}
	return decimalArr
}
