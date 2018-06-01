# gRPCinDockerDemo
a test repository
  
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


由于二进制程序文件太占用GitHub空间, 二进制文件将在一天后删除.
