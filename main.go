package main

import (
	"GO/qiandao/bilibili"
	"GO/qiandao/hostloc"
	"GO/qiandao/smzdm"
	"GO/qiandao/tieba"
	"GO/qiandao/v2ex"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

const (
	s1 = `("tieba"=|"(bilibili|smzdm|hostloc|v2ex)"=.*|\r|\f|\n|########################END################)`
	s2 = `("bilibili"=|"(tieba|smzdm|hostloc|v2ex)"=.*|\r|\f|\n|########################END################)`
	s3 = `("smzdm"=|"(tieba|bilibili|hostloc|v2ex)"=.*|\r|\f|\n|########################END################)`
	s4 = `("hostloc"=|"(tieba|bilibili|smzdm|v2ex)"=.*|\r|\f|\n|########################END################)`
	s5 = `("v2ex"=|"(tieba|bilibili|smzdm|hostloc)"=.*|\r|\f|\n|########################END################)`
)

func openconf(s string) string { //读取本地文件
	fi, err := os.Open(s)
	if err != nil {
		log.Println("没有这个文件！")
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func regexp1(s, v string) string { //正则删除字符
	reg := regexp.MustCompile(s)
	regs := reg.ReplaceAllString(v, "")
	return regs
}

func main() {
	var cookie string
	cookie0 := openconf("cookie.conf") //读取本地conf

	cookie = regexp1(s1, cookie0)

	if cookie != "" {
		tieba.Tieba(cookie) //把贴吧cookie传递到GO/qiandao/tieba.go包使用
		cookie = ""
	}
	cookie = regexp1(s2, cookie0)

	if cookie != "" {
		bilibili.Bilibili(cookie) //把贴吧cookie传递到GO/qiandao/bilibili.go包使用
		cookie = ""
	}
	cookie = regexp1(s3, cookie0)

	if cookie != "" {
		smzdm.SMZDM(cookie) //把贴吧cookie传递到GO/qiandao/smzdm.go包使用
		cookie = ""
	}
	cookie = regexp1(s4, cookie0)

	if cookie != "" {
		hostloc.Hostloc(cookie) //把贴吧cookie传递到GO/qiandao/hostloc.go包使用
		cookie = ""
	}
	cookie = regexp1(s5, cookie0)

	if cookie != "" {
		v2ex.V2ex(cookie) //把贴吧cookie传递到GO/qiandao/v2ex.go包使用
		cookie = ""
	}
	log.Println("签到完成，退出程序!")
}
