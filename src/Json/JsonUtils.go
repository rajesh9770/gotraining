package Json

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	First string
	Last string
	Age int `json:"-"`   // this is like tagging field transient
	PublicAge int `json:"PAge"`
	notExported int
}

func JsonEx(){
	p1 := Person{"Rajesh", "Manjrekar", 44, 40, 007}
	
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
	
	var p2 Person
	json.Unmarshal(bs, &p2)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.PublicAge, p2.notExported)
	
	var p3 Person
	bs = [] byte(`{"First" : "Rajesh1", "Last" : "Manjrekar1", "Age" : 44, "PAge": 40, "notExported" : 100}`)
	json.Unmarshal(bs, &p3)
	fmt.Println(p3.First, p3.Last, p3.Age, p3.PublicAge, p3.notExported)
}