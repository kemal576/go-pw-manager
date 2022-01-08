
# Go Password Manager 

This project is the API part of the password management software I developed for Yemeksepeti Go Bootcamp.
## Technologies

**Database & Driver:** [PostgreSQL](https://www.postgresql.org) & [Pq](https://github.com/lib/pq)

**Secret Management:** [HashiCorp Vault](https://github.com/hashicorp/vault)

**Router:** [HttpRouter (julienschmidt)](https://github.com/julienschmidt/httprouter)

**CORS:** [Cors (rs)](https://github.com/rs/cors)

**JWT:** [jwt-go (dgrijalva)](https://github.com/dgrijalva/jwt-go)

## Installation & Run

To get this project:

```bash
  $ go get github.com/kemal576/go-pw-manager
```

To run this project:
> :warning: This project uses Vault to store hidden variables. Therefore, you must first add the following variables into the Vault.

```bash
  $ go run main.go
```

  
## Secret Variables & Vault


---------------- KEYS -----------------     ------- PATHS -------

`DB_NAME` , `USERNAME` , `PASSWORD` --> `secret/DB_SECRETS`

`JWT_KEY` ----------------------------> `secret/JWT`

`ENC_KEY` ----------------------------> `secret/AES`

```bash
  //command prompt sample code
  $ vault kv put secret/JWT JWT_KEY=mysecretjwtkey123
```
## Environment Variables

To run this project, you will need to create a config.yaml file in the root project directory and add the following environment variables into it.

`VAULT_ADRESS`

`VAULT_PORT`

`VAULT_TOKEN`


  
## => Folder structure and endpoint documentation will be added.
