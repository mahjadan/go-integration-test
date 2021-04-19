# go-integration-test
using dockertest to run multiple containers for integration test

# About
This is a demo API to show how to use **dockertest** for writing integrations tests.
this API has 2 end-points :
* /planets/:id (which makes a call to Star War API and get a specific Planet Info)
* /people/:id  (which access MongoDB and retreive a sample document)

the idea is to write integration test using docker container package 'dockertest' which will start 3 containers:
1. **wiremock** container to mock the external API call ( Star War in our example)
2. **MongoDB** container to mock the second end point and retrieve something from database.
3. **API Itself** container by making a build and run it using the wiremock and mongodb mocked urls.

this is all done using golang **MainTest** func which will setup all theses containers before running any test.
we have also added to this project **build tags**  to avoid running integration tests with regular unit tests.

# To Run the Tests
```
 cd integration 
 go test -tags integration
```
 
 # To run the API 
 using docker-compose we can run the API with MongoDB and Mongo-Express
```
 docker-compose up --build --remove-orphans
```

