package main

import (
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

const line int = 200000

func main() {
	// 配置文件保存路径
	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkErr(err)
	fileDir := path.Clean(path.Join(curDir, "file-store"))
	//fileDir = strings.Replace(fileDir, "\\", "/", -1)

	// 创建文件
	filename := "big_input_file.txt"
	inputFile, err := os.Create(path.Join(fileDir, filename))
	checkErr(err)

	// 随机生成字符串
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < line; i++ {
		words := r.Intn(10) // 长度不超过10
		var line string
		for j := 0; j < words; j++ {
			letters := 0 // 每个串的长度在[3,10)
			for ; letters < 3; {
				letters = r.Intn(10)
			}
			var word string
			for k := 0; k < letters; k++ {
				ch := r.Intn(26)
				word += string(ch + 97)
			}
			line += word + " "
		}
		_, _ = inputFile.WriteString(line + "\n")
	}
	_ = inputFile.Close()
}
