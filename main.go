package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// 指定端口、仓库、mmhook地址
// https://blog.csdn.net/tower888/article/details/115718980

var (
	FlagPort int    // http 监听端口
	FlagRepo string // 仓库名称
	FlagHook string // mattermost webhook
)

func init() {
	flag.IntVar(&FlagPort, "port", 9090, "Http listen port.")
	flag.StringVar(&FlagRepo, "repo", "", "Gitlab repository name.")
	flag.StringVar(&FlagHook, "webhook", "", "Mattermost webhook.")
}

func main() {
	// flag
	flag.Parse()
	if flag.NFlag() <= 0 || len(FlagRepo) == 0 || len(FlagHook) == 0 {
		flag.PrintDefaults()
		return
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/alert", hookHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", FlagPort), nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Action, Module")

	// test mm
	//http.Post(FlagHook, "application/json", strings.NewReader("{\"text\": \"hello\"}"))
	io.WriteString(w, "Hello, world!\n")
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	var hook HeaderType
	// 调用json包的解析，解析请求body
	// if err := json.NewDecoder(r.Body).Decode(&reqdata); err != nil {
	// 	r.Body.Close()
	// 	log.Fatal(err)
	// }
	var data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request: %s", err)
		r.Body.Close()
		return
	}
	//log.Println("req json: ", string(data))

	err = json.Unmarshal(data, &hook)
	if err != nil {
		log.Printf("Failed to parse request: %s", err)
		return
	}

	if hook.ObjectKind == "push" && hook.Repository.Name == FlagRepo {
		handlePush(&hook, data)
	}

	// 返回json字符串给客户端
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(reqdata)
	io.WriteString(w, "hookHandler!\n")
}
