
# Go Password Manager API

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


## API Usage
> :warning: A valid JWT is required for all endpoints except adding users and logging in. The JWT should be added to the Authorization header of the request to be sent, without the "Bearer" part.


- **User Endpoints:**

#### Get all users:

```http
  GET /users
```
------
#### Get specified user:

```http
  GET /users/${id}
```

| Parameter | Explanation                       |
| :-------- | :-------------------------------- |
| `id`      | **Necessary**. Key value of the item to be called. |

------
#### Create user:

```http
  POST /users
```
Sample request body:
```json
{
 "firstname": "Kemal",
 "lastname": "Şahin",
 "password": "test123",
 "email": "kemalsahin@gmail.com"
}

```
---
#### Update user:

```http
  PUT /users
```
Sample request body:
```json
{
 "id": "1",
 "firstname": "Kemal",
 "lastname": "Şahin",
 "password": "test123",
 "email": "kemalsahin@gmail.com"
}

```
-----
#### Delete user:

```http
  DELETE /users/${id}
```

| Parameter | Explanation                       |
| :-------- | :-------------------------------- |
| `id`      | **Necessary**. Key value of the item to be deleted. |

----

- **Login Endpoints:**
  
#### Get all logins:

```http
  GET /logins
```
------
#### Get specified login:

```http
  GET /logins/${id}
```

| Parameter | Explanation                       |
| :-------- | :-------------------------------- |
| `id`      | **Necessary**. Key value of the item to be called. |

------

#### Get logins for specified user

```http
  GET /user/${id}/logins
```

| Parameter | Explanation                       |
| :-------- | :-------------------------------- |
| `id`      | **Necessary**. Key value of the item to be called. |

------
#### Create login:

```http
  POST /logins
```
Sample request body:
```json
{
  "url": "kommunity.com",
  "identity": "kemalsahin",
  "password": "test123",
  "userId": 1
}

```
---
#### Update login:

```http
  PUT /logins
```
Sample request body:
```json
{
  "id": "3",
  "url": "kommunity.com",
  "identity": "kemalsahin576",
  "password": "test1234",
}

```
-----
#### Delete user:

```http
  DELETE /logins/${id}
```

| Parameter | Explanation                       |
| :-------- | :-------------------------------- |
| `id`      | **Necessary**. Key value of the item to be deleted. |

----

- **Auth Endpoints:**

#### Sign-In:

```http
  POST /signin
```
Sample request body:
```json
{
  "email": "kemalsahin@gmail.com",
  "password": "test123"
}

```
---
