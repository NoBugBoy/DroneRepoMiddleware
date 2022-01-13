package main

import (
	"encoding/json"
	"fmt"
	. "github.com/NoBugBoy/httpgo/http"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("开始监听8080")
	http.HandleFunc("/", dronePullYaml)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	log.Fatal(err)
}

type Drone struct {
	Build struct{} `json:"build"`
	Repo struct{
		NameSpace string `json:"namespace"`
		RepoName string `json:"name"`
		Branch string `json:"default_branch"`
	} `json:"repo"`

}
type Data struct {
	Data string
}
var droneRepoUrl = "http://172.16.3.130:31886/%s/dronerepo/raw/branch/%s/%s/%s/drone.yaml"
func dronePullYaml(res http.ResponseWriter,req *http.Request)  {
	defer req.Body.Close()

	b, err := ioutil.ReadAll(req.Body)
	//fmt.Println(string(b))
	if err != nil {
		fmt.Println("read request.Body failed, err", err)
		return
	}
	drone := &Drone{}
	json.Unmarshal(b,drone)
	marshal, _ := json.Marshal(drone)
	fmt.Println(string(marshal))
	url := fmt.Sprintf(droneRepoUrl,drone.Repo.NameSpace,drone.Repo.Branch,drone.Repo.RepoName,drone.Repo.Branch)
	fmt.Println(url)

	httpR := &Req{}
	body, err := httpR.Url(url).
		Method(http.MethodGet).Go().Body()
	fmt.Println(body)
	d := &Data{
		Data: body,
	}
	marshaler,_ := json.Marshal(d)
	res.Write(marshaler)
}

