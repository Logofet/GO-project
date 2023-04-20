# Short explanation

In this project, I made a simple web application using Golang and MySQL that allows users to "Create", "Read", "Update", and "Delete" their information in the database. Firstly, I created a simple database that takes users' information such as ID, Name, Email, and Age using the code below:
```
create database mojabaza;
create table users (
`ID` integer auto_increment primary key,
`Name` varchar(20),
`Email` varchar(50),
`Age` integer
);
```

Next, I started working on the Golang part. I used github.com/go-sql-driver/mysql to connected the Golang application to Mysql database and checked whether the connection was successful by pinging the database.

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

After that, I created a new router in order to handle the HTTP requests. For path "/users" I handled Get, Post and for "/users/{id}" i handled Put and Delete requests.
Finally, I made the router listen on port 6969.

# How to run

In MySql run the script mentioned above. 
Next, in terminal run: go run main.go .
I tested the implementation using Postman and added the files that i've tested in the "Postman" folder.

Post - You can add new user by adding "name, "email" and "age" in params and filling the information up.
Get- You can get information from all existing users, or search them individualy.
Put- You can change user's information by typing the ID of the user in the URL, and typing the new information in Body (JSON)
Delete - You can delete an existing user by typing user's ID and pressing DELETE
