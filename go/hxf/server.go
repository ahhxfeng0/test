package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strings"
	"sort"
	"bytes"
)
type phonenumbers struct{
	Items []*PhoneNumberA
}
type PhoneNumberA struct{
	TYPE string
	Name string
	Phone_number string

}

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers   []Server
	ServersID string
}

type Items struct {
    // Ono string `json:ono`
    OrderItem []Item `json:item`
    // OrderRefund []Refund `json:refund`
}
type Item struct {
    Name string `json:name`
    Phone_number string    `json:phone_number`
}


type sortByName []Item
func (a sortByName) Len() int { return len(a)}
func (a sortByName) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a sortByName) Less(i, j int) bool {
	if a[i].Name < a[j].Name {
		return 	true
	}
	if a[i].Name > a[j].Name {
		return 	false
	}
	return a[i].Name < a[j].Name
}


func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Fprintf(w, "Hi, wait for sort %s", html.EscapeString(r.URL.Path[1:]))
	if r.Method == "GET" {
		fmt.Println("method:", r.Method) //获取请求的方法

		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])

		for k, v := range r.Form {
			fmt.Print("key:", k, "; ")
			fmt.Println("val:", strings.Join(v, ""))
		}
	} else if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)
		var items Items
		decoder := json.NewDecoder(bytes.NewBuffer(result))
		err := decoder.Decode(&items)
		if err != nil {
			panic(err)
		}
		sort.Sort(sortByName(items.OrderItem))
		json.NewEncoder(w).Encode(items)
		if err := json.Unmarshal(result, &items); err == nil{
			fmt.Printf("%+v", items)
			// fmt.Fprint(w, items)
			// fmt.Printf(phoneData.TYPE)
			// resp,_ := json.Marshal(phoneData)

		}

		// //未知类型的推荐处理方法

		// var f interface{}
		// json.Unmarshal(result, &f)
		// m := f.(map[string]interface{})
		// for k, v := range m {
		// 	switch vv := v.(type) {
		// 	case string:
		// 		fmt.Println(k, "is string", vv)
		// 	case int:
		// 		fmt.Println(k, "is int", vv)
		// 	case float64:
		// 		fmt.Println(k, "is float64", vv)
		// 	case []interface{}:
		// 		fmt.Println(k, "is an array:")
		// 		for i, u := range vv {
		// 			fmt.Println(i, u)
		// 		}
		// 	default:
		// 		fmt.Println(k, "is of a type I don't know how to handle")
		// 	}
		// }

		//结构已知，解析到结构体

		// var s PhoneNumberA
		// json.Unmarshal([]byte(result), &s)
		
		// fmt.Println("hello world !")
		// fmt.Println(s.TYPE)
		// fmt.Println(s.Name)
		// fmt.Println(s.Phone_number)
		// fmt.Printf("%+v", s)
		// fmt.Fprintf(w, "Items:")
		// fmt.Fprint(w, phoneData)

		// for i := 0; i < len(s.Servers); i++ {
		// 	fmt.Println(s.Servers[i].ServerName)
		// 	fmt.Println(s.Servers[i].ServerIP)
		// 	fmt.Fprintf(w, s.Servers[i].ServerName)
		// 	fmt.Fprintf(w, s.Servers[i].ServerIP)
		// }
	}
}
