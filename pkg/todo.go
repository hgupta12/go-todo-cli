package todo

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
var fileName = "todos.txt"

type ToDo struct{
	id string
	todo string
	completed bool
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
	f, err:= os.Open(fileName)

	if err!=nil{
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	list := []ToDo{}
	for scanner.Scan(){
		text := strings.Split(scanner.Text(),",")
		completed, _ := strconv.ParseBool(text[2])
		todo := ToDo{text[0],text[1], completed}
		list = append(list,todo)
	}

	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}

	return list
}

func AddToDo(todos []string){
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	
	if err!= nil{
		log.Fatal(err)
	}
	defer f.Close()
	for _,str := range todos{
		id := generateId()
		todo := ToDo{
			id,str,false,
		}
		if _,err := f.WriteString(fmt.Sprintf("%s,%s,%v\n",todo.id,todo.todo,todo.completed)); err != nil{
			log.Fatal(err)
		}
	}
}
func ShowToDoList(all bool){
	list := getAllTodos()
	var output bool
	for _,item := range list{
		if all || !item.completed{
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
	list := getAllTodos()
	var data string
	for _,item := range list{
		if item.id != id{
			data = data + fmt.Sprintf("%s,%s,%v\n",item.id,item.todo,item.completed)
		}
	}

	if err:=ioutil.WriteFile(fileName,[]byte(data),0644); err!=nil{
		log.Fatal(err)
	}
	fmt.Println("ToDo deleted!")
}

func MarkToDo(id string,f bool){
	list := getAllTodos()
	var data string
	for _,item := range list{
		if item.id == id{
			data = data + fmt.Sprintf("%s,%s,%v\n",item.id,item.todo,!f)
		}else{
			data = data + fmt.Sprintf("%s,%s,%v\n",item.id,item.todo,item.completed)
		}
	}

	if err:=ioutil.WriteFile(fileName,[]byte(data),0644); err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("ToDo Marked as %v!\n",!f)
}
