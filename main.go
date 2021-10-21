package main

import (
	"net/http"
	"io/ioutil"

	"RankSearch/utils"
)
type stu struct {
	name 		string
	userName    string
	hduRank 	int
	hduSol 		int
	wustRank 	int
}

func main() {
	httpString := "http://acm.hdu.edu.cn/userstatus.php?user="
	stus := make([]stu, 0)

	var curStu stu 
	username := "pyoxiao"
	curStu.userName = username
	curHttp := httpString + username
	r, err := http.Get(curHttp)
	if err != nil {
		panic(err)
	}
	defer func(){
		_ = r.Body.Close()
	}()
	body, _ := ioutil.ReadAll(r.Body)
	str := string(body)
	curStu.hduRank = utils.GetNumber(str, utils.StrRank)
	curStu.hduSol = utils.GetNumber(str, utils.StrSol)
	stus = append(stus, curStu)
	
}