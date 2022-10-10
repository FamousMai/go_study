package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
		"1": "data1.json",
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
		file, err := os.Open("../domain/json/" + s2)
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

				// 创建文件夹
				if !checkFileIsExist("../file/" + m.Data.Course.Name) {
					os.Mkdir("../file/"+m.Data.Course.Name, 0666) //创建文件
					//fmt.Println("文件不存在" + err1.Error())
				}

				// 重命名
				fileName := strings.Replace(s.Md_url, "https://oss-hqwx-video.hqwx.com/", "", -1)
				Original_Path := "../file/" + fileName + ".mp4"
				New_Path := "../file/" + m.Data.Course.Name + "/" + s.Title + ".mp4"
				if checkFileIsExist(Original_Path) {
					e := os.Rename(Original_Path, New_Path)
					if e != nil {
						log.Fatal(e)
					}
				}

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

func rename() {

}
