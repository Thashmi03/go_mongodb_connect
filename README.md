# go_mongodb_connect

## Steps:
************

1. Install mongo db driver

2. Defining/creating the connnection string

3. Creating a connection context

4. Call the connection get the instance of connected object

5. Prepare the query 

6. Execute the query again as the connected instance

7. Convert the result set into golang structures


## To connect to database
******************************
1. we need to have the name of the database 

2. we need to have the collection name

# Mongo connections to do
**********************************
## steps for employeeCollection 

1. create a common method that takes two parameters 

    first-database name

    second- collection name

  this method will return mmongo collection



## operators

1. $unwind-used to split the array documents or array fields into separate documents for each limit in the array.
eg:db.getCollection("students").aggregate([{$unwind:"$course"}])

2. $pipeline - when we are performing the aggreegation pipeline will give you the different stages to perform the aggregation.

3. $match-will be acting as the search query that can be intergrated with aggregation pipeline. 

4. $sort,$skip-we can use in aggregation pipeline
5. $lookup-will be used to perform the joint between two tables
eg-db.Books.aggregate([{
    $lookup:{
        from:"Authors",
        localField:"author_id",
        foreignField:"_id",
        as:"AuthorDetail"
    }
}])
## Update query
# query
***************
## a>7000,tc==100
db.getCollection("transactions").find({"$and":[
    {"transaction_count":{"$eq":100}},
    {"account_id":{"$lte":700000}}
    ]})

## sum
db.getCollection("transactions").aggregate([{
    $group:{
        _id:null,
        sum:{
            $sum:"$transaction_count"
        }
    }
}])
## sum of all transacations(amounts)

db.getCollection("transactions").aggregate([{
    $project:{
        sum:{
            $sum:"$transactions.amount"
        }
    }
}])