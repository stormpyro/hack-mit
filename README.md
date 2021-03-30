# Golang clean architecture

This is a project with 4 services in order to apply clean architecture in backend project using echo, mongodb and swagger.

## Run project

```
go run main.go
```

## Run with nodemon

```
nodemon --exec go run . --ext go
```

## Open swagger

```
http://localhost:3000/swagger/index.html
```

## Clean architecture layers

- <b>Database</b>  
  Is the repository folder. Here you can do
  all your databases operations. For example create a file called "collectionA". Here you need to write a function to get the collection and all the operations like retrieve all documents, insert a new document, update a document, delete a document and delete all documents in the collection.
- <b>Usecase</b> \
  Is the usecase folder. Here you can use all your repository functions(database folder) for write the logic business. For example if you need to get data from different collections to show them in a table you can call your repository operations(database folder) and manage the data in the form that you need.
- <b>Delivery</b> \
  In the project there is a folder called services. Here you can write all the services that you want, in this case there is a service called fue. Each service needs to handle data sent by the user. This is exactly what deliver layer do. So each service needs a file called deliver.go wich will handle all data sent by the user, then call usecase functions to manage this data and retrieves the result or a http error.

## Packages used

---

- MongoDB driver: \
  https://github.com/mongodb/mongo-go-driver
- Swaggo: \
  https://github.com/swaggo/swag
- Echo swagger: \
  https://github.com/swaggo/echo-swagger
- Echo: \
  https://github.com/labstack/echo
- Viper: \
  https://github.com/spf13/viper

## Note

---

You need a config.json file with the mongo db uri from your mongo atlas cluster or your local mongo like this.

```
{
    "mongodb": {
        "uri": "mongodb+srv://contacto:<password>@cluster0.v3zeq.mongodb.net/<yourDB>?retryWrites=true&w=majority"
    }
}
```
