Simple webserver that has the ability to update a counter per request.
where webservers request counter should stop at exactly 100.



To Run server by using docker file, you must have docker installed.
Run the following commands.
```$xslt
    cd web_server
    docker build -t web_server .
    docker run -d -p 8080:8081 --name server web_server

```
To test the webserver, we will use the tool Apache benchmark (ab) 
Run command in your terminal
```$xslt
    ab -n 100 -c 10 http://127.0.0.1:8080/increment
```
And see the logs by typing
```$xslt
    docker logs server
```
*******************************************

Run without Dockerfile
```$xslt
    go build main.go
    ./main
```
To test the webserver, we will use the tool Apache benchmark (ab) 
Run command in your terminal
```$xslt
    ab -n 100 -c 10 http://127.0.0.1:8081/increment
```
You can see the logs from the terminal where server is running.