Tasks
* Create a producer service
* Create a consumer service
* S3 file uploader service
* DB structure

# Bill generator
The project demonstrate async task processing and interaction between producer and consumer

## Description
We have users who can require to generate pdf for they bills. We have a service that accepts a request from the user,
puts it to the ques and returns the taskID to the user. The taskID is used to check the status of the task.

Consumer service takes the task from the queue and process it. Consumer generates a pds bill, put it to s3 storage and 
store in the db



# Building and running the project
