# Anime CRUD Microservice

This is a simple CRUD microservice which is a rest api written in Golang, Gorilla as development framework and MongoDB as the database. It is deseigned for saving and manipulating anime records

![Real File](https://raw.githubusercontent.com/Debetome/Anime-CRUD-microservice/master/assets/records.png)

## Setup

### Clone repository

```Bash
git clone https://github.com/Debetome/Anime-CRUD-microservice.git
```

### Declare environment variables

#### Windows

```Batch
set MONGO_HOST=localhost
set MONGO_PORT=27017
set MONGO_USER=user
set MONGO_PASSWORD=password123
```

#### Linux

```Bash
export MONGO_HOST=localhost
export MONGO_PORT=27017
export MONGO_USER=user
export MONGO_PASSWORD=password123
```

**NOTE**: The 'MONGO_PORT' variable is just optional, in case you have mongodb running on a different

## Endpoints


```Bash
/get-animes           
```
It fetches all anime records from the database

---

```Bash
/get-anime/{id}      
```
It fetches only one anime record

---

```Bash
/new-anime
```
It creates a new anime record (it receives a json body)

---

```Bash
/update-anime/{id}
```
It updates or edits a anime record (it receives a json body)

---

```Bash
/delete-anime/{id}
```
It deletes a anime record from the database