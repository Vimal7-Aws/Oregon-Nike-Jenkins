package model

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func DisplayState() State {
	oregon := State{
		Name:    "Oregon",
		Rivers:  []string{"Coloumbia Gorge -v2", "Tualatin - v2"},
		Created: time.Now(),
	}

	var jsonData []byte
	jsonData, err := json.MarshalIndent(oregon, "", "    ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	return oregon
}
