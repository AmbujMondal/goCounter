# goCounter
sample counter api implemented in go lang with Kubernettes cluster

build the application

launch the exe file generated-
in terminal write command `` .\goCounter.exe``

if no error is displayed in the terminal, means the applicaiton is running.

Three endpoints are exposed by this application (see he main.go file)
can be tested in postman-
1) HTTP GET - get the current counter value : Default value is 0. 
```http://localhost:8082/api/counter```

2) HTTP PUT - increment the counter by 1
```http://localhost:8082/api/counter```

3) HTTP DELETE - decrement the counter by 1
```http://localhost:8082/api/counter```