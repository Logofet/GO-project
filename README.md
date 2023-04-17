# Go-assingnment

In the next project i made a simple web application using Golang and Mysql that allows users to "Create", "Read", "Update" and "Delete" their informations on the database. 
First, I made a simple database that takes the users information (ID, Name, Email and Age) with the code below:

create database mojabaza;
create table users (
`ID` integer auto_increment primary key,
`Name` varchar(20),
`Email` varchar(50),
`Age` integer
);

Next i started working on the Golang part. 
