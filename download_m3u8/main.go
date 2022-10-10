package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Message struct {
	Success bool
	Data    struct {
		New []struct {
			Md_url string
			Title  string
		}
		Course struct {
			Name string
		}
	}
}

func main() {

	jsonFiles := map[string]string{
		//"1": "data1.json",
		"2": "data2.json",
		"3": "data3.json",
		"4": "data4.json",
		"5": "data5.json",
		"6": "data6.json",
		"7": "data7.json",
		"8": "data8.json",
	}

	Md_urls := ""

	for _, s2 := range jsonFiles {
		/**
		读取文件
		*/
		file, err := os.Open("domain/json/" + s2)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)

		dec := json.NewDecoder(strings.NewReader(string(content)))
		for {
			var m Message
			if err := dec.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
			}
			for _, s := range m.Data.New {

				fmt.Println(s.Md_url)
				//fmt.Println(s.Title)
				Md_urls = Md_urls + s.Md_url + "\n"
			}
		}
	}

	writeFile(Md_urls)
}

func writeFile(wireteString string) {
	var filename = "./CourseLink.txt"
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()
	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("写入 %d 个字节n", n)
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
