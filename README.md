### Plugin description

Purpose of plugin to get values from vault server and without any middle steps
run ```helm install/upgrade```

### Variables

- To set specific vault address set environment variable
  ```export VAULT_ADDR=https://vault.com:8200```

- To set specific namespace set environment variable
  ```export VAULT_NAMESPACE==development```

- In case of problems with ssl
  ```VAULT_SKIP_VERIFY=1```
  
- Search secrets in needed location
  ```export VAULT_SECRET_PATH=secret/development```
