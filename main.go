package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"io"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port int
	Db   string
	User string
}

var Cnf Config

func init() {
	c, _ := ini.Load("config.ini")
	Cnf = Config{
		Port: c.Section("web").Key("port").MustInt(),
		Db:   c.Section("db").Key("name").MustString("hogehoge.sql"),
		User: c.Section("db").Key("user").String(),
	}
}

func LoggingSettings(logFile string) {
	// RDWRはreadとwrite。パーミッションで0666は読み書きができるユーザーその他。
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {

	fmt.Printf("%v \n", Cnf.Port)
	fmt.Printf("%v \n", Cnf.Db)
	fmt.Printf("%v \n", Cnf.User)

	LoggingSettings("test.log")
	// テキトーなファイル名を開く処理
	_, err := os.Open("iqsgfiquegf")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	res, _ := http.Get("https://example.com")

	fmt.Printf("[\x1b[32m%s\x1b[0m]\n", "OK   ")
	fmt.Printf("[\x1b[31m%s\x1b[0m]\n", "   NG")

	log.Println("aaa")

	//fmt.Println(res.StatusCode)
	fmt.Println(res.Proto)
	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])

	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Print(string(body))

}
