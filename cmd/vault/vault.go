package main

import (
	"encoding/json"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"os"
	"strings"
)

var filename string

func  main()  {
	if len(os.Args) > 2 {
		for i := 1; i < len(os.Args); i += 2 {
			getParameter(os.Args[i], i)
		}
		if filename != "" {
			writeVaultSecret()
			fmt.Printf(filename)
		}
	} else {
		fmt.Print(getVaultSecret())
	}
}

func checkError (err error){
	if err != nil {
		panic(err)
	}
}

func getParameter (action string, seq int) {
	if seq + 1 < len(os.Args) {
		switch action {
			case "-f":
				filename = os.Args[seq+1]
		}
	}
}

func getVaultSecret() string {
	 client, err := vault.NewClient(&vault.Config{Address: os.Getenv("VAULT_ADDR")})
	 checkError(err)
	 data, err := client.Logical().Read(os.Getenv("VAULT_SECRET_PATH"))
	 checkError(err)
	 j, _ := json.Marshal(data.Data)
	 return propToYaml(string(j))
}

func propToYaml (raw string) string {
	var data map[string]string
	err := json.Unmarshal([]byte(raw), &data)
	checkError(err)
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

func writeVaultSecret () {
	file, err := os.Create(filename)
	checkError(err)
	file.WriteString(getVaultSecret())
	file.Sync()
}