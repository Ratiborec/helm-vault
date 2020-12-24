package main

import (
	"encoding/json"
	"fmt"
	yaml "github.com/ghodss/yaml"
	vault "github.com/hashicorp/vault/api"
	"os"
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
	 y, _ := yaml.JSONToYAML(j)
	 return string(y)
}

