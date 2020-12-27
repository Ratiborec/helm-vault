package main

import (
	"encoding/json"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"os"
	"strings"
)

func  main()  {
	fmt.Print(getVaultSecret())
}

func getVaultSecret() string {
	 client, err := vault.NewClient(&vault.Config{Address: os.Getenv("VAULT_ADDR")})
	 if err != nil {
		 panic(err)
	 }
	 data, err := client.Logical().Read(os.Getenv("VAULT_SECRET_PATH"))
	 if err != nil {
	 	panic(err)
	 }
	 j, _ := json.Marshal(data.Data)
	 return propToYaml(string(j))
}
func propToYaml (raw string) string {
	var data map[string]string
	err := json.Unmarshal([]byte(raw), &data)
	if err != nil {
		panic(err)
	}
	var myYaml string = ""
	for k, v := range data {
		splitedData := strings.Split(k, ".")
		var tmpStr string = splitedData[0] + ":"
		if len(splitedData) > 1 {
			tmpStr = tmpStr + "\n"
		} else {
			tmpStr = tmpStr + " " + v + "\n"
			myYaml = myYaml + tmpStr
		}
		var counter int = 1
		for a := 1; a < len(splitedData); a++ {
			for i := 0; i < counter; i++ {
				tmpStr = tmpStr + "  "
			}
			if a == len(splitedData) - 1 {
				tmpStr = tmpStr + splitedData[1] + ": " + v + "\n"
				myYaml = myYaml + tmpStr
			}
			tmpStr = tmpStr + splitedData[1] + ":" + "\n"
			counter++
		}
	}
	return myYaml
}

