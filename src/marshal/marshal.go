package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "Ram",
		Age:   32,
	}
	u2 := user{
		First: "Sham",
		Age:   30,
	}

	users := []user{u1, u2}
	fmt.Println(users)

	bs, err := json.Marshal(&users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	users1 := []user{}
	err = json.Unmarshal(bs, &users1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users1)

}
