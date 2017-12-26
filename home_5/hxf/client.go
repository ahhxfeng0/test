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
	ItemA []ItemA
	ItemB []ItemB
	ItemC []ItemC
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
	phonetype := RandomString(PhoneType, 1)
	for i := 0; i < 4; i++ {
		switch phonetype {
		case "A":
			m.ItemA = append(m.ItemA, ItemA{Name: RandomString(name, 1), PhoneNumber: RandomString(phone, 4)})
		case "B":
			m.ItemB = append(m.ItemB, ItemB{NickName: RandomString(name, 1), Phone: RandomString(phone, 4), Address: RandomString(name, 5)})
		case "C":
			m.ItemC = append(m.ItemC, ItemC{FullName: RandomString(name, 1), Telephone: RandomString(phone, 4), Note: RandomString(name, 5)})
		}
		time.Sleep(1000)
	}
	// var m1 ItemsA
	// for i := 0; i < 4; i++ {
	// 	m1 = append(m1, ItemA{Name: RandomString(name, 1), PhoneNumber: RandomString(phone, 4)})
	// }

	// var s Serverslice

	// var newServer Server
	// newServer.ServerName = "Guangzhou_VPN"
	// newServer.ServerIP = "127.0.0.1"
	// s.Servers = append(s.Servers, newServer)

	// s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.2"})
	// s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.3"})

	// s.ServersID = "team1"
	// phoneData := make(map[string]string)
	// var PA PhoneNumberA
	// // PA.TYPE = "A"
	// PA.Name = "zhangsan"
	// PA.Phone_number = "641324643"
	// fmt.Printf("%#v\n", PA)
	// // var PhoneNumberAs phonenumbers
	// PhoneNumberAs := new(phonenumbers)

	// // PhoneNumberAs.Items = [{
	// // 	{Name:"d", Phone_number:"516"},
	// // 	{Name:"C", Phone_number:"576"},
	// // 	{Name:"e", Phone_number:"526"},
	// // 	{Name:"a", Phone_number:"536"},

	// // }]
	// // PhoneNumberA := []
	// PhoneNumberA := [{"Name":"d", "Phone_number":"516"},
	// 	{"Name":"C", "Phone_number":"576"},
	// 	// {Name:"e", Phone_number:"526"},
	// 	// {Name:"a", Phone_number:"536"},

	// ]

	// var s Items
	// s.Item = append(s.Item, PhoneNumberA{Name:"a", Phone_number:"456"})

	// var s phonenumbers
	// s.Items = append(s.Items, PhoneNumberA{Name:"d", Phone_number:"516"})

	// tmp := `{"TYPE":"A", "name":"lisi", "phone_number":"564123416546"}`
	// req = bytes.NewBuffer([]byte(tmp))
	// resp, _ = http.Post("http://localhost:9001/sort", "application/json;charset=utf-8", body)



	// var PA1 PhoneNumberA
	// PA.TYPE = "A"
	// PA.phoneData["name"] = "lisi"
	// PA.phoneData["phone_number"] = "314651882"
	// PA = PhoneNumberPost()


	// b, err := json.Marshal(s)
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
