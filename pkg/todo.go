package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
var fileName = "todos.txt"

type Response struct{
	Status string
	Code string
	Todos []ToDo `json:"todos"`
}
type ToDo struct{
	Id string `json:"id"`
	Todo string `json:"todo"`
	Completed bool `json:"completed"`
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func generateId() string{
	id := make([]rune,5)
	for i:=0;i<5;i++{
		id[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(id)
}

func getAllTodos() []ToDo{
	res,err := http.Get("http://localhost:8080/todos")
	if err !=nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body,err := ioutil.ReadAll(res.Body)
	if err !=nil{
		log.Fatal(err)
	}
	var response Response
	json.Unmarshal(body,&response)
	return response.Todos
}

func AddToDo(todos []string){
	for _,str := range todos{
		id := generateId()
		todo := url.Values{
			"id": {id},
			"todo":{str},
		}
		fmt.Println(todo)
		_,err := http.PostForm("http://localhost:8080/todos",todo)
		if err != nil{
			log.Fatal(err)
		}
	}
}
func ShowToDoList(all bool){
	list := getAllTodos()
	var output bool
	for _,item := range list{
		if all || !item.Completed{
			fmt.Printf("%v\n",item)
			output = true
		}
	}
	if len(list)==0 || !output{
		fmt.Println("No ToDos to show")
		return
	}
}

func DeleteToDo(id string){
	url := fmt.Sprintf("http://localhost:8080/todos/%s",id)
	client := &http.Client{}

	req,err := http.NewRequest(http.MethodDelete,url,nil)
	if err!=nil{
		log.Fatal(err)
	}
	if _,err :=client.Do(req); err != nil{
		log.Fatal(err)
	}

	fmt.Println("ToDo deleted!")
}

func MarkToDo(id string,f bool){
	url := fmt.Sprintf("http://localhost:8080/todos/%s?value=%v",id,!f)

	client := &http.Client{}

	req,err := http.NewRequest(http.MethodPatch, url,nil)
	if err!=nil{
		log.Fatal(err)
	}

	if _,err :=client.Do(req); err != nil{
		log.Fatal(err)
	}

	fmt.Printf("ToDo Marked as %v!\n",!f)
}
