# DESAFIO - Criando Primeira API REST com GO
Escrever uma APIREST que a vovó possa guardar dados dos seus clientes

---

### tools:
- linux
- goland ide
- go1.22.0
- package gorilla/mux

---

### routes

| API ROUTE		            | DESCRIPTION           | STATUS |
|:-----------------------|:----------------------------|:-------|
| [GET] /customers	      | Retrieve all the customers  | 200    |
| [GET] /customer/{id}   | Retrieve a customer by ID   | 200    |
| [POST] /customers      | Add a new customer          | 200    |

### initial set of data

| id | firstName | lastName   | email | address           |
| :--- | :--- |:-----------| :--- |:--------------------|
| 1 | Ana | Maria   | ana@email.com | São Paulo SP        |
| 2 | Maria | Clara | maria@email.com | Belo Horizonte MG |

---

#### homepage
```bash
curl --location 'http://localhost:8010/'
```
###### 200 OK
``` json
Bem vinde à página da Vovó
```

#### get all customers
```bash
curl --location 'http://localhost:8010/customers'
```
###### 200 OK
``` json
[
  {
    "id": 1,
    "firstName": "Ana",
    "lastName": "Maria",
    "email": "ana@email.com",
    "address": {
      "city": "São Paulo",
      "state": "SP"
    }
  },
  {
    "id": 2,
    "firstName": "Maria",
    "lastName": "Clara",
    "email": "maria@email.com",
    "address": {
      "city": "Belo Horizonte",
      "state": "MG"
    }
  }
]
```

#### get all customer by ID
```bash
curl --location 'http://localhost:8010/customer/1'
```
###### 200 OK
``` json
{
  "id": 1,
  "firstName": "Ana",
  "lastName": "Maria",
  "email": "ana@email.com",
  "address": {
    "city": "São Paulo",
    "state": "SP"
  }
}
```

#### add a customer
```bash
curl --location 'http://localhost:8010/customers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "john@email.com",
    "address": {
        "city": "Salvador",
        "state": "BA"
    }
}'
```

###### 200 OK
``` json
{
    "id": 3,
    "firstName": "Joao",
    "lastName": "Doe",
    "email": "joao@email.com",
    "address": {
        "city": "Salvador",
        "state": "BA"
    }
}
```
