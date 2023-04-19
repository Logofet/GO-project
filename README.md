#Short explanation

In this project i made a simple web application using Golang and Mysql that allows users to "Create", "Read", "Update" and "Delete" their informations on the database. 
First, I made a simple database that takes the users information (ID, Name, Email and Age) with the code below:
```
create database mojabaza;
create table users (
`ID` integer auto_increment primary key,
`Name` varchar(20),
`Email` varchar(50),
`Age` integer
);
```

Next I started working on the Golang part. 
I used github.com/go-sql-driver/mysql to connected the Golang application to Mysql database and checked whether the connection was successful by pinging the database.

```
db, err := sql.Open("mysql", "root:mysqlmysql@tcp(127.0.0.1:3306)/mojabaza")
if err != nil {
	log.Fatal(err)
}
defer db.Close()

err = db.Ping()
if err != nil {
	log.Fatal(err)
}
```

After that i created a new router in order to handle the http requests. For path "/users" i handle Get, Post and for "/users/{id}" i handle Put and Delete.
Finally I make the router listen on port 6969.

#How to run

In MySql run the above mentioned script. 
Next, in terminal run: go run main.go .
I tested the implementation using Postman, and added those files that i've tested in the "Postman" folder.

Post - you can , add new user by adding "name, "email" and "age" in params and filling the information up.
Get- you can get information from all users that are stored already exist, or search them individualy.
Put- you can change users information by typing the ID of the user in the URL, and typing the new information in Body (JSON)
Delete - you can delete existing user by typing users ID and pressing DELETE
