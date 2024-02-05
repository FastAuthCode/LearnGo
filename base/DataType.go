package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	//整数
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	//浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	//复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	//布尔类型变量的默认值为false。
	//Go 语言中不允许将整型强制转换为布尔型.
	//布尔型无法参与数值运算，也无法与其他类型进行转换。

	var flag bool
	fmt.Printf("%v\n", flag)
	flag = true
	fmt.Printf("%v\n", flag)

	//字符串
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	s1 := `第一行
第二行
第三行
`
	fmt.Println(s1)

	//len(str)	求长度
	fmt.Println(strconv.Itoa(len(s1)))
	//+或fmt.Sprintf	拼接字符串
	fmt.Printf(fmt.Sprintf("这是一个格式化输出:%s", s1))
	//strings.Split	分割
	split := strings.Split(s1, "行")
	for i := range split {
		fmt.Println(split[i])
	}
	//strings.contains	判断是否包含
	fmt.Println(strings.Contains(s1, "第一"))
	//strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
	fmt.Println(strings.HasSuffix(s1, "第一"))
	fmt.Println(strings.HasPrefix(s1, "第一"))
	//strings.Index(),strings.LastIndex()	子串出现的位置
	fmt.Println(strings.Index(s1, "第一"))
	fmt.Println(strings.LastIndex(s1, "第一"))

	//strings.Join(a[]string, sep string)	join操作
	fmt.Println(strings.Join(split, "行"))

	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

	s2 := "big"
	// 强制类型转换
	byteS2 := []byte(s2)
	byteS2[0] = 'p'
	fmt.Println(string(byteS2))

	s3 := "白萝卜"
	runeS3 := []rune(s3)
	runeS3[0] = '红'
	fmt.Println(string(runeS3))
	sqrtDemo()
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

	TestDataType()

	str := "hello沙河小王子"
	count := countChineseCharacters(str)
	fmt.Printf("字符串 \"%s\" 中汉字的数量为：%d\n", str, count)
}

//写代码分别定义一个整型、浮点型、布尔型、字符串型变量，
//使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。

func TestDataType() {
	// 定义整型变量
	var integerVar int = 42
	fmt.Printf("整型变量值：%v，类型：%T\n", integerVar, integerVar)

	// 定义浮点型变量
	var floatVar float64 = 3.14
	fmt.Printf("浮点型变量值：%v，类型：%T\n", floatVar, floatVar)

	// 定义布尔型变量
	var boolVar bool = true
	fmt.Printf("布尔型变量值：%v，类型：%T\n", boolVar, boolVar)

	// 定义字符串型变量
	var stringVar string = "Hello, Go!"
	fmt.Printf("字符串型变量值：%v，类型：%T\n", stringVar, stringVar)
}

func countChineseCharacters(s string) int {
	count := 0
	for _, r := range s {
		if utf8.RuneLen(r) > 1 {
			count++
		}
	}
	return count
}
