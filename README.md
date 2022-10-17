# Anime CRUD Microservice

This is a simple CRUD microservice which is a rest api written in Golang, Gorilla as development framework and MongoDB as the database. It is deseigned for saving and manipulating anime records

![Real File](https://raw.githubusercontent.com/Debetome/Anime-CRUD-microservice/master/assets/records.png)

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