# gRPCinDockerDemo
a test repository

测试环境:  
1. ubuntu 16.04  
2. docker 18.03.1-ce
3. go 1.10.2
  
Usage:  
1. cd /path/to/gRPCinDockerDemo/dockerServer  
2. docker build -t dserver .  
3. docker run --net=host -d dserver  
4. cd /path/to/gRPCinDockerDemo/dockerClient  
5. docker build -t dclient .  
6. docker run --net=host -d dclient  
7. use url: curl "http://localhost:8000/?m=3&n=6"  
   output will be: s: 9  
   Note that double quotation marks.  
8. stop containers:  
![image](https://github.com/jacenr/gRPCinDockerDemo/blob/master/stop_containers.png)  

