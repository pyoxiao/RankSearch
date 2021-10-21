package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"io/ioutil"
	"os"
	"unicode"
	"sort"
	"strconv"

	"RankSearch/utils"
)

func main() {
	httpString := "http://acm.hdu.edu.cn/userstatus.php?user="
	stus := make([]utils.Stu, 0)
	path := "/root/pyo/stus.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file open err")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		var username []byte
		var name []byte
		for i, v := range line {
			if unicode.IsSpace(rune(v)) {
				name = line[:i]
				username = line[(i + 1):]
				break
			} 
		}
		var curStu utils.Stu 
		curStu.Name = string(name)
		curStu.UserName = string(username)
		curHttp := httpString + string(username)
		r, err := http.Get(curHttp)
		if err != nil {
			fmt.Println("http err")
			return
		}
		defer func(){
			_ = r.Body.Close()
		}()
		body, _ := ioutil.ReadAll(r.Body)
		str := string(body)
		curStu.HduRank = utils.GetNumber(str, utils.StrRank)
		curStu.HduSol = utils.GetNumber(str, utils.StrSol)
		stus = append(stus, curStu)
	}
	sort.Sort(utils.Stus(stus))
	result, err := os.Create("/root/pyo/result.txt")
	if err != nil {
		fmt.Println("create file err")
		return
	}
	defer result.Close()
	result.WriteString("姓 名 hdu账号名 HduRank HduSolve WustRank\n")
	for i, stu := range stus {
		result.WriteString(stu.Name) 
		result.WriteString(" ")
		result.WriteString(stu.UserName) 
		result.WriteString(" ")
		result.WriteString(strconv.Itoa(stu.HduRank)) 
		result.WriteString(" ")
		result.WriteString(strconv.Itoa(stu.HduSol)) 
		result.WriteString(" ")
		result.WriteString(strconv.Itoa(i + 1)) 
		result.WriteString("\n")
	}
}