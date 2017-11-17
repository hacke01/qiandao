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

func openconf(s string) string { //读取本地文件
	fi, err := os.Open(s)
	if err != nil {
		log.Println("没有这个文件！")
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}
func regexp0(s, v string) []string { //正则寻找字符串
	reg := regexp.MustCompile(s)
	regs := reg.FindAllString(v, -1)

	return regs
}
func regexp1(s, v string) string { //正则删除字符
	reg := regexp.MustCompile(s)
	regs := reg.ReplaceAllString(v, "")
	return regs
}

func main() {
	var cookie string
	cookie0 := openconf("cookie.conf") //读取本地conf

	c0 := regexp0(`"tieba"=.+`, cookie0)
	c1 := regexp0(`"bilibili"=.+`, cookie0)
	c2 := regexp0(`"smzdm"=.+`, cookie0)
	c3 := regexp0(`"hostloc"=.+`, cookie0)
	c4 := regexp0(`"v2ex"=.+`, cookie0)

	cookie = regexp1(`("tieba"=|\r)`, c0[0])
	if cookie != "" {
		tieba.Tieba(cookie) //把贴吧cookie传递到GO/qiandao/tieba.go包使用
		cookie = ""
	}
	cookie = regexp1(`("bilibili"=|\r)`, c1[0])
	if cookie != "" {
		bilibili.Bilibili(cookie) //把贴吧cookie传递到GO/qiandao/bilibili.go包使用
		cookie = ""
	}
	cookie = regexp1(`("smzdm"=|\r)`, c2[0])
	if cookie != "" {
		smzdm.SMZDM(cookie) //把贴吧cookie传递到GO/qiandao/smzdm.go包使用
		cookie = ""
	}
	cookie = regexp1(`("hostloc"=|\r)`, c3[0])
	if cookie != "" {

		hostloc.Hostloc(cookie) //把贴吧cookie传递到GO/qiandao/hostloc.go包使用
		cookie = ""
	}
	cookie = regexp1(`("v2ex"=|\r)`, c4[0])
	if cookie != "" {

		v2ex.V2ex(cookie) //把贴吧cookie传递到GO/qiandao/v2ex.go包使用
		cookie = ""
	}

}
