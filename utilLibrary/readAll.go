package utillibrary

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

func ReadAll() {
	// r := strings.NewReader("abcde")

	// buf, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(buf)

	v := "Manikanta"
	strings.Count("manika", v)
	fmt.Print(v)

	// var ct int

	// for range v {
	// 	ct++
	// }
	// fmt.Println(ct)

	// for i :=0;i<= len(v); i++{
	// 	fmt.Print(ct)
	// }

	b := "abcd"

	var ct int
	for range b {
		ct++
	}
	fmt.Println(ct)

	//------------------------------------
	m := map[string]bool{
		"Go is awesome":             true,
		"I drink and I know things": true,
		"PI IS EXACTLY THREE!":      false,
	}
	data, err := yaml.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
