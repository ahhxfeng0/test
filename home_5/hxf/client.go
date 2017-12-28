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


type PhoneSorter interface {
	getName()
	// sortByName()
}


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

	c, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json err:", err)
	}

	body := bytes.NewBuffer([]byte(c))
	res, err := http.Post("http://localhost:9001/sort?type="+ phonetype, "application/json;charset=utf-8", body)
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
