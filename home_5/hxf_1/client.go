package main

import (
	"time"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"math/rand"
	// "sort"
	// "net/url"
)
const PhoneType = "ABC"
const name = "abcdefghijklmnopqrstuvwxyz"
const phone = "0123456789"
func  RandomString(pixff string, strlen int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = pixff[rand.Intn(len(pixff))]
	}
	return string(result)
}



// type Server struct {
// 	ServerName string
// 	ServerIP   string
// }

// type Serverslice struct {
// 	Servers   []Server
// 	ServersID string
// }
// type JsonPostSample struct{

// }
// // phoneData := make(map[string]string)
// type phonenumbers struct{
// 	Items []PhoneNumberA

// }
// type Items struct{
// 	Item []PhoneNumberA
// }
// type PhoneNumberA struct{
// 	Name string
// 	Phone_number string

// }

// type PhoneNumberB struct{
// 	TYPE string
// 	Nickname string
// 	Phone string
// 	Address string

// }
// type PhoneNumberC struct{
// 	TYPE string
// 	Fullname string
// 	Telephone string
// 	Note string

// }

// type Items struct {
//     // Ono string `json:ono`
//     OrderItem []Item `json:item`
//     // OrderRefund []Refund `json:refund`
// }

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
type Items struct{
	Item []interface{}
}
type ItemsA struct{
	ItemA []ItemA
}

// func (i *ItemsA)getName()()  {
// 	name := &i.Name
// 	return * name
// }
// func (i *ItemsA)sortByName()  {
// 	sort.Sort(ItemsA.ItemA)
// }
// // type nameList []ItemA
// func (a *ItemA) Len() int {return len(a)}
// func (a *ItemA) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
// func (a *ItemA) Less(i, j int) bool {
// 	return a[i].name < a[j].name
// }
// func (i *Items)sortByName()  {
// 	name, itemName := i.getName()
// 	type nameList []itemName
// 	// len := func (a nameList) Len() (int) {return len(a)}
// 	// return func (a nameList) Len() int {
// 	// 	return len(a)
// 	// }
// 	return func (a nameList)int {
// 		return len(a)
// 	}
// 	return func (a nameList)  {

// 	}
// 	return func(x int) int {
// 		sum += x
// 		return sum
//    }
// 	// return func (a nameList) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
// 	// func (a nameList) Less(i, j int) bool {
// 	// 	return a[i].name < a[j].name
// 	// }
// }
// type i *Items
// var attr, itemName = (*Items).getName(i)

// type nameList []itemName
// func (a nameList) Len() int { return len(a)}
// func (a nameList) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
// func (a nameList) Less(i, j int) bool {
// 	return a[i].name < a[j].name
// }

// var attr, itemName = i.getName()
// type namelist []itemName
// func (a nameList) Len() int { return len(a)}
// func (a nameList) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
// func (a nameList) Less(i, j int) bool {
// 	return a[i].attr < a[j].attr
// }
// type nameLista []ItemA
// 	func (a nameLista) Len() int { return len(a)}
// 	func (a nameLista) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
// 	func (a nameLista) Less(i, j int) bool {
// 		return a[i].name < a[j].name
// 	}
type PhoneSorter interface {
	getName()
	// sortByName()
}

// type ItemsA struct{
// 	Item []Item
// }

func main() {


	var m Items
	// m.Item = append(m.Item, ItemA{Name: RandomString(name, 1), PhoneNumber: RandomString(phone, 4)})
	phonetype := RandomString(PhoneType, 1)
	for i := 0; i < 4; i++ {
		switch phonetype {
		case "A":
			m.Item = append(m.Item, ItemA{Name: RandomString(name, 1), PhoneNumber: RandomString(phone, 4)})
		case "B":
			m.Item = append(m.Item, ItemB{NickName: RandomString(name, 1), Phone: RandomString(phone, 4), Address: RandomString(name, 5)})
		case "C":
			m.Item = append(m.Item, ItemC{FullName: RandomString(name, 1), Telephone: RandomString(phone, 4), Note: RandomString(name, 5)})
		}
		time.Sleep(1000)
	}
	// var b int
	// // fmt.Println("接口类型：", m.Item.(type))
	// // switch v := m.Item[0].(type) {
	// // case struct:
	// // 	fmt.Println("结构体")
	// // }
	// type Element interface {}
	// var e Element = 100
	switch value := m.Item[0].(type) {
	case int:
		fmt.Println("int", value)
	case string:
		fmt.Println("string", value)
	default:
		fmt.Println("unknown", value)
		// var Itemss struct{}
		// Itemss = m.Item
		fmt.Printf(m.Item.(struct))
	}
	
	c, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json err:", err)
	}

	body := bytes.NewBuffer([]byte(c))
	res, err := http.Post("http://localhost:9001/sort", "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
}
