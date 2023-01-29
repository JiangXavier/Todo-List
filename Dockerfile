
FROM golang:latest    

RUN mkdir /app     //创建一个文件夹
 
WORKDIR /app       //移动到指定的目录

ADD . /app         //将当前的目录的文件都拷贝到/app目录

RUN go build -o main ./main.go //构建go二进制文件

EXPOSE 5080        //输出的端口号

CMD /app/main      //执行,有点像shell
 