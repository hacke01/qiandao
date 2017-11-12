package main

import (
	"github.com/Caeseason/qiandao/qiandao/bilibili"
	"github.com/Caeseason/qiandao/qiandao/smzdm"
	"github.com/Caeseason/qiandao/qiandao/tieba"

	"log"
	"os"
	"regexp"
	"time"
)

func openconf(s string) string { //读取本地文件
	z, err := os.Open("cookie.conf")
	if err != nil {
		log.Println("没有cookie.conf文件！")
	}
	buf := make([]byte, 1024)
	n, _ := z.Read(buf)
	cookie := string(buf[:n])
	defer z.Close()
	return cookie
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
func date(cookie string) string {
	cookie = ""
	time.Sleep(1 * time.Second)
	return cookie
}
func main() {
	var cookie string
	cookie0 := openconf("cookie.conf")                           //读取本地conf
	cookie1 := regexp0(`\"(tieba|bilibili|smzdm)\"=.+`, cookie0) //func用regexp0,传递cookie里的每一行

	cookie = regexp1(`(\"tieba\"=|\r)`, cookie1[0])
	if cookie != "" {
		tieba.Tieba(cookie) //把贴吧cookie传递到GO/qiandao/tieba.go包使用
		cookie = date(cookie)
	}

	cookie = regexp1("(\"bilibili\"=|\r)", cookie1[1])
	if cookie != "" {
		bilibili.Bilibili(cookie) //把贴吧cookie传递到GO/qiandao/bilibili.go包使用
		cookie = date(cookie)
	}
	cookie = regexp1("(\"smzdm\"=|\r)", cookie1[1])
	if cookie != "" {
		smzdm.SMZDM(cookie) //把贴吧cookie传递到GO/qiandao/smzdm.go包使用
		cookie = date(cookie)
	}

}
