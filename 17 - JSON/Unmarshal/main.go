package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` //se quiser ignorar o campo é so colocar o -
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}


func main(){
	cachorroEmJSON := `{"nome":"Amora", "raca":"misturada", "idade":1}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil{
		log.Fatalln(erro)
	} //precisa converter o json em slice de byte

	fmt.Println(c)

	jsonCat := `{"nome":"Malu", "raca": "angorá", "idade": 1}`
	mapCat := make(map[string]interface{})

	if erro := json.Unmarshal([]byte(jsonCat), &mapCat); erro != nil{
		log.Fatalln(erro)
	}
	fmt.Println(mapCat)
}
