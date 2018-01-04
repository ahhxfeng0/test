package main

import (
	// "sync"
	"encoding/json"
	"fmt"
	// "html"
	"io/ioutil"
	"net/http"
	// "strings"
	"sort"
	// "bytes"
)



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
func (i ItemA)getName() string{
	return i.Name
}
func (i ItemB)getName() string{
	return i.NickName
}
func (i ItemC)getName() string{
	return i.FullName
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
type sortList []PhoneGetNamer
func (a sortList) Len() int {return len(a)}
func (a sortList) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a sortList) Less(i, j int) bool {
	return a[i].getName() < a[j].getName()
}
// type sortLista []ItemA
// func (a sortLista) Len() int {return len(a)}
// func (a sortLista) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
// func (a sortLista) Less(i, j int) bool {
// 	return a[i].getName() < a[j].getName()
// }

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
type PhoneSorter interface{
	sorter() sortList
}


type Factory interface{
	Create() PhoneSorter
}
var Factories = make(map[string]Factory)
// type itemMapa struct{}
// func (i itemMapa)getMap()  map[string]interface{}{
// 	return map[string][]ItemA
// }

func (i ItemA)Create() PhoneSorter  {
	instance := make(PhoneSorterA)
	a := &instance
	return a
}

func (i ItemB)Create() PhoneSorter {
	instance := make(PhoneSorterB)
	a := &instance
	return a
}
func (i ItemC)Create() PhoneSorter {
	instance := make(PhoneSorterC)
	a := &instance
	return a
}

type PhoneSorterA map[string][]ItemA

func (s PhoneSorterA)sorter() sortList {
	// items := make(sortList, len)
	list := s["items"]
	var newList sortList
    for _, value := range list {
	newList = append(newList, value)
    }
    sort.Sort(sortList(newList))
    return newList
	
}

type PhoneSorterB map[string][]ItemB

func (s PhoneSorterB)sorter() sortList {
	// items := make(sortList, len)
	list := s["items"]
	var newList sortList
    for _, value := range list {
	newList = append(newList, value)
    }
    sort.Sort(sortList(newList))
    return newList
	
}

type PhoneSorterC map[string][]ItemC

func (s PhoneSorterC)sorter() sortList {
	// items := make(sortList, len)
	list := s["items"]
	var newList sortList
    for _, value := range list {
	newList = append(newList, value)
    }
    sort.Sort(sortList(newList))
    return newList
	
}
func Register(name string, factory Factory)  {
	if factory == nil {
		panic("Must not provide nil Factory")
	}
	_, registered := Factories[name]
	if registered {
		panic(fmt.Sprintf("Factory named %s already registered", name))
	}
	Factories[name] = factory
}

func Create(name string) PhoneSorter	 {
	Factory, ok := Factories[name]
	if !ok {
		return nil
	}
	return Factory.Create()
}

// 
// type phonesorters = struct{
// 	m map[string]Factory
// 	sync.RWMutex
// }{m: make(map[string]Factory)}

// func Create(id string) PhoneSorter {
// 	return phonesorters.m[id].Create()
// }

// func Register(id string, newfactory Factory)  {
// 	phonesorters.m[id] = newfactory
// }

// type Itemamap struct{}
// func (s *Itemamap)Create() PhoneSorter {
	
// }

// func itemOperater(result []byte, itemMap map[string]type)  {
// 	// result, _ = ioutil.ReadAll(r.Body)
// 	json.Unmarshal(result, &itemMap)
// 	items := make(sortList, len(itemMap["items"]))
// 	for i, item := range itemMap["items"]{
// 		items[i] = item
// 	}
// 	sort.Sort(sortList(items))
// }
func main() {

	Register("A", ItemA{})
    Register("B", ItemB{})
    Register("C", ItemC{})
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	// fmt.Fprintf(w, "Hi, wait for sort %s", html.EscapeString(r.URL.Path[1:]))
	if r.Method == "GET" {
		fmt.Println("method:", r.Method) //获取请求的方法
	} else if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		// fmt.Printf("%s\n", result)
		phoneType := r.FormValue("type")
		mapData := Create(phoneType)
		// var Data Items
		fmt.Println("result", result)
		json.Unmarshal(result, &mapData)
		// sort.Sort(sortList(mapData))
		// fmt.Println("Data", Data)
		// mapData := make(map[string]ite)
		// mapData["items"] = Data
		fmt.Println("mapData", mapData)
		listData := mapData.sorter()
		mapData_res := make(map[string]sortList)
		mapData_res["items"] = listData
		fmt.Println(listData)
        response, _ := json.Marshal(mapData_res)
		w.Write([]byte(fmt.Sprintf(string(response) + "\n")))
		// switch sortType := r.FormValue("type"); sortType {
		// case "A":
		// 	var itemMap map[string][]ItemA
		// 	json.Unmarshal(result, &itemMap)	
		// 	items := make(sortList, len(itemMap["items"]))		
		// 	for i, item := range itemMap["items"] {
		// 		items[i] = item	
		// 	}
		// 	sort.Sort(sortList(items))		
		// 	for i, item := range items {
		// 		itemMap["items"][i] = item.(ItemA)	
		// 	}
		// 	result, _ := json.Marshal(itemMap)
		// 	w.Write([]byte(fmt.Sprintf(string(result) + "\n")))
		// case "B":
		// 	var itemMap map[string][]ItemB
		// 	json.Unmarshal(result, &itemMap)	
		// 	items := make(sortList, len(itemMap["items"]))		
		// 	for i, item := range itemMap["items"] {
		// 		items[i] = item	
		// 	}
		// 	sort.Sort(sortList(items))		
		// 	for i, item := range items {
		// 		itemMap["items"][i] = item.(ItemB)	
		// 	}
		// 	result, _ := json.Marshal(itemMap)
		// 	w.Write([]byte(fmt.Sprintf(string(result) + "\n")))
		// case "C":
		// 	var itemMap map[string][]ItemC
		// 	json.Unmarshal(result, &itemMap)	
		// 	items := make(sortList, len(itemMap["items"]))		
		// 	for i, item := range itemMap["items"] {
		// 		items[i] = item	
		// 	}
		// 	sort.Sort(sortList(items))		
		// 	for i, item := range items {
		// 		itemMap["items"][i] = item.(ItemC)	
		// 	}
		// 	result, _ := json.Marshal(itemMap)
		// 	w.Write([]byte(fmt.Sprintf(string(result) + "\n")))
		// default:
		// 	w.Write([]byte(fmt.Sprintf("unknow sortType: %s.\n", sortType)))
		// }
		fmt.Printf("%+v", phoneType)
		// var items Items
		// decoder := json.NewDecoder(bytes.NewBuffer(result))
		// err := decoder.Decode(&items)
		// if err != nil {
		// 	panic(err)
		// }
		// sort.Sort(sortLista(items.ItemA))
		// json.NewEncoder(w).Encode(items)
		var Items Items
		if err := json.Unmarshal(result, &Items); err == nil{
			fmt.Printf("%+v", Items)
			// fmt.Fprint(w, items)
			// fmt.Printf(phoneData.TYPE)
			// resp,_ := json.Marshal(phoneData)

		}
	}
}
