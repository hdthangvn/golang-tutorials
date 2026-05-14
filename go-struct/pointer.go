package main

import (
	"fmt"
	"reflect"
)

// 1. profileGao
type Person struct {
	Name string
	Age  int
}

type BInfo struct {
	Address string
	Gender  string
}

type PersonFull struct {
	Name string
	Age  int
	BInfo
	// ... có thể thêm nhiều trường khác nếu cần
}

func updateProfile(profile PersonFull, newAge int) PersonFull {
	profile.Age = newAge
	return profile
}

func updateProfileOnline(p *PersonFull, newAge int) {
	p.Age = newAge
}

func runExample() {
	// 1. profileGao
	profileGao := PersonFull{
		Name: "Mr Rice",
		Age:  5,
		BInfo: BInfo{
			Address: "123 Main St",
			Gender:  "Male",
		},
	}
	fmt.Println(profileGao.Age) // in ra tuổi của Gao
	fmt.Println("type profileGao:", reflect.TypeOf(profileGao))

	// 2. profileKen
	var profileKen Person
	profileKen.Name = "Mr Ken"
	profileKen.Age = 6
	fmt.Println(profileKen.Name)                                // in ra tên của Ken
	fmt.Println("type profileKen:", reflect.TypeOf(profileKen)) //reflect: thông tin về kiểu biến.

	// 3. updateProfileRice
	profileGao = updateProfile(profileGao, 6)
	fmt.Println("Updated age of profileGao:", profileGao.Age)

	// 4. updateProfileOnline
	updateProfileOnline(&profileGao, 7)
	fmt.Println("Updated age of profileGao (online):", profileGao.Age)

	// 5. nested struct
	profileGaoFull := PersonFull{
		Name: "Mr Rice",
		Age:  6,
		BInfo: BInfo{
			Address: "VN",
			Gender:  "Male",
		},
	}
	fmt.Println("Address of profileGaoFull:", profileGaoFull.BInfo.Address) // VN
}
