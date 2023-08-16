# go_mongodb_connect

Steps:
************

1. Install mongo db driver

2. Defining/creating the connnection string

3. Creating a connection context

4. Call the connection get the instance of connected object

5. Prepare the query 

6. Execute the query again as the connected instance

7. Convert the result set into golang structures


To connect to database
***********************
1. we need to have the name of the database 

2. we need to have the collection name

steps for employeeCollection 
**********************************
1. create a common method that takes two parameters 

    first-database name

    second- collection name

  this method will return mmongo collection
