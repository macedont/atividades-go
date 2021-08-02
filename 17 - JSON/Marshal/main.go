package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"spike", "filaBr", 1}
	//fmt.Println(c)

	cachorroJson, erro := json.Marshal(c)
	if erro != nil{
		log.Fatalln("Houve algum erro ao colocar o struct como json")
	}
	fmt.Println(bytes.NewBuffer(cachorroJson))

	usuario := map[string]string{
		"nome": "Macedo",
		"idade": "22",
		"trabalho": "ifood",
	}

	usuarioJSON, erro := json.Marshal(usuario)
	if erro != nil{
		log.Fatalln("Houve algum erro ao colocar o map como json")
	}
	fmt.Println(bytes.NewBuffer(usuarioJSON))

	gato := []string{"amora", "misturada", "alegre"}

	gatoJSON, erro := json.Marshal(gato)
	if erro != nil{
		log.Fatalln("Houve um erro ao colocar o slice como json")
	}

	fmt.Println(bytes.NewBuffer(gatoJSON))
}
