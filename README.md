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
  
### How to use

Plugin in development.

You can build ```./cmd/vault/vault.go``` and leave at ```$HELM_PLUGIN_DIR/bin/```

Also leave ```plugin.yaml``` in ```$HELM_PLUGIN_DIR```

You need to export variables that described higher and run ```helm vault```

Example of using:

```
filename=$(mktemp /tmp/value.yaml.XXXXXXX)
helm value > $filename
helm install release_name chart -f $filename
```

### TODO

1. Installation script to automate install process
   
2. Make possibility to run in such way of similar

```helm vault install release_name chart```