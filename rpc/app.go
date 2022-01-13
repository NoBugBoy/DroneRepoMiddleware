package rpc

import (
	"fmt"
	"log"
	"net/http"
)

func init(){
	http.HandleFunc("/", dronePullYaml)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	log.Fatal(err)
}

type Repo struct {
	namespace string
	name string
}

func dronePullYaml(res http.ResponseWriter,req *http.Request)  {
	fmt.Println(req.Body)
}
