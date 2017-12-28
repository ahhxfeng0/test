package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	// "strings"
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


type ItemA struct {
    Name string `json:name`
    PhoneNumber string    `json:phone_number`
}

type ItemB struct{
	NickName string `json:name`
	Phone string `json:phone`
	Address string `json:address`
}

type ItemC struct{
	FullName string `json:name`
	Telephone string `json:phone`
	Note string `json:note`
}

type PhoneGetNamer interface {
	getName() string  
	// sortByName() struct
}
// type PhoneSorter interface{
// 	PhoneGetNamer
// 	sortByName() struct
// }

// func (i Items)sortByName() struct{
// 	var i[0] PhoneGetNamer = new(type) 
// }

type Items struct{
	ItemA []ItemA
	ItemB []ItemB
	ItemC []ItemC
}
type ItemsA struct{
	ItemA []ItemA
}
// type Items struct {
//     // Ono string `json:ono`
//     OrderItem []Item `json:item`
//     // OrderRefund []Refund `json:refund`
// }
// type Item struct {
//     Name string `json:name`
//     Phone_number string    `json:phone_number`
// }
func (i *ItemA)getName() string{
	name := i.Name
	return name
}
func (i *ItemB)getName() string{
	name := i.NickName
	return name
}
func (i *ItemC)getName() string{
	name := i.FullName
	return name
}

// func (i ItemA)sortByName() struct {
// 	sort.Sort(sortLista(i))
// 	return i
// }
// func (i *ItemsA)sortByName()  struct{
// 	sort.Sort(sortList(ItemsA.ItemA))
// 	return ItemsA
// }
// type nameList []ItemA
// type sortList []PhoneGetNamer
// func (a sortList) Len() int {return len(a)}
// func (a sortList) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
// func (a sortList) Less(i, j int) bool {
// 	return a[i].getName() < a[j].getName()
// }
type sortLista []ItemA
func (a sortLista) Len() int {return len(a)}
func (a sortLista) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a sortLista) Less(i, j int) bool {
	return a[i].getName() < a[j].getName()
}

// type sortListb []ItemB
// func (a sortListb) Len() int {return len(a)}
// func (a sortListb) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
// func (a sortListb) Less(i, j int) bool {
// 	return a[i].getName() < a[j].getName()
// }

// type sortListc []ItemC
// func (a sortListc) Len() int {return len(a)}
// func (a sortListc) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
// func (a sortListc) Less(i, j int) bool {
// 	return a[i].getName() < a[j].getName()
// }
// type sortByName []Item
// func (a sortByName) Len() int { return len(a)}
// func (a sortByName) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
// func (a sortByName) Less(i, j int) bool {
// 	if a[i].Name < a[j].Name {
// 		return 	true
// 	}
// 	if a[i].Name > a[j].Name {
// 		return 	false
// 	}
// 	return a[i].Name < a[j].Name
// }


func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Fprintf(w, "Hi, wait for sort %s", html.EscapeString(r.URL.Path[1:]))
	if r.Method == "GET" {
		fmt.Println("method:", r.Method) //获取请求的方法
	} else if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)
		phoneType := r.FormValue("type")
		fmt.Printf("%+v", phoneType)
		var items Items
		// switch phoneType {
		// case "A":
		// 	// json.Unmarshal(result, &items.ItemA)
		// 	// sort.Sort(sortList(items.ItemA))
		// 	decoder := json.NewDecoder(bytes.NewBuffer(result))
		// 	err := decoder.Decode(&items.ItemA)
		// 	if err != nil {
		// 		panic(err)
		// 	}
			
		// 	sort.Sort(sortLista(items.ItemA))
		// 	json.NewEncoder(w).Encode(items)
		// case "B":
		// 	// json.Unmarshal(result, &items.ItemA)
		// 	// sort.Sort(sortList(items.ItemA))
		// 	decoder := json.NewDecoder(bytes.NewBuffer(result))
		// 	err := decoder.Decode(&items.ItemB)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	sort.Sort(sortListb(items.ItemB))
		// 	json.NewEncoder(w).Encode(items)
		// case "C":
		// 	// json.Unmarshal(result, &items.ItemA)
		// 	// sort.Sort(sortList(items.ItemA))
		// 	decoder := json.NewDecoder(bytes.NewBuffer(result))
		// 	err := decoder.Decode(&items.ItemC)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	sort.Sort(sortListc(items.ItemC))
		// 	json.NewEncoder(w).Encode(items)
		// }
		decoder := json.NewDecoder(bytes.NewBuffer(result))
		err := decoder.Decode(&items)
		if err != nil {
			panic(err)
		}
		sort.Sort(sortLista(items.ItemA))
		json.NewEncoder(w).Encode(items)
		if err := json.Unmarshal(result, &items); err == nil{
			fmt.Printf("%+v", items)
			// fmt.Fprint(w, items)
			// fmt.Printf(phoneData.TYPE)
			// resp,_ := json.Marshal(phoneData)

		}
	}
}
