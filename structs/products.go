package structs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Products struct {
	Products []Product
}

func (prods *Products) ReadJson(name string) {

	jsonprods, err := os.Open(name)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonprods)

	if err1 := json.Unmarshal(byteValue, &prods); err1 != nil {
		log.Fatal(err1)
	}

}