# Anime CRUD Microservice

---

This is a simple CRUD microservice which is a rest api written in Golang, Gorilla as development framework and MongoDB as the database. It is deseigned for saving and manipulating anime records

![Real File](https://raw.githubusercontent.com/Debetome/Anime-CRUD-microservice/master/assets/records.png)

## Endpoints

---

```Bash
/getAnimes           It fetches all anime records from the database
```

```Bash
/getAnimes/{id}      It fetches only one anime record
```

```Bash
/newAnime            It creates a new anime record (it receives a json body)
```

```Bash
/updateAnime/{id}    It updates or edits a anime record (it receives a json body)
```

```Bash
/deleteAnime/{id}    It deletes a anime record from the database
```