This code is created for Golang learner !!!

Application : coursemanagement 

Tech stack:

Frontend : Needs to create

Backend Server:
# Golang
# MySQL
# Docker
# GIT
# AWS
# postman for API testing

The code is ready for Deployment.

Deployment requirement:

#create an instance in AWS
#Allow only 3306 and 8080 ports

# Install Docker 

Please refer the below link for Docker installation.

https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04

# Check GIT installed or not.

#if yes, Please clone the URL given below or Install GIT.

git clone https://github.com/9500073161/coursemanagement.git

# go to Coursemanagement Directory

# check Docker status

sudo systemctl status docker

# build docker image

 docker-compose up --build

 # wait for few min and manualy start/run app container.

 docker start coursemanagement-app-1

 # login to DB bash

 docker exec -it coursemanagement-db-1 bash

 mysql -u root -p1234

 # check Database 

 show databases;

---------
 +--------------------+
| Database           |
+--------------------+
| coursemanagement   |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)


# select coursemanagement 

use coursemanagement -A;

# check all tables are created or not

show tables;

+----------------------------+
| Tables_in_coursemanagement |
+----------------------------+
| courses                    |
| entrollments               |
| students                   |
| teachers                   |
+----------------------------+


# now we have to test Restapi with postman.

   install postman from below link.

   https://www.postman.com/downloads/

# configure postman and do the testing for Crud operations.

create ccollection and name coursemanagement

add request for course, student, teacher and entrollment

example for course

POST http://localhost:8080/course

go to header

add below configuration as a bulkedit

Content-Type: application/json
Accept: application/json

go to Body and use below format and send.

{
    "id": 3,
    "teacher_id": 103,
    "name": "Devops"
}

now request will send to server and created row in course. Do it for all table and test it. 

This will work with GET, POST, PUT, and DELETE.

Thank you. Wish you all the best. 

####

Creating CICD pipeline for Coursemanagement System

1.GIT
2.Jenkins











