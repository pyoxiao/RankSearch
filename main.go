package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)
type stu struct {
	name string
	id string
	hduRank int
	wustRank int
}

func main() {
	httpString := "http://acm.hdu.edu.cn/userstatus.php?user="
	user := "pyoxiao"
	curHttp := httpString + user
	r, err := http.Get(curHttp)
	if err != nil {
		panic(err)
	}
	defer func(){
		_ = r.Body.Close()
	}()
	body, _ := ioutil.ReadAll(r.Body)
	s := string(body)
	strNeed := "<tr><td>Rank</td><td align=center>"
	index := strings.Index(s, strNeed)
	if index == -1 {
		panic(-1)
	} 
	num := 0
	for pos := index + len(strNeed); s[pos] >= '0' && s[pos] <= '9'; pos += 1 {
		num = num * 10 + int(s[pos] - '0');
	}
	fmt.Println(num)
}