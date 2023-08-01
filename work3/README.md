### 创建镜像
```shell
docker build -t golang-work2:v1 .
docker images golang-work2:v1
```

### 启动镜像
```shell
docker run --net host -itd  --name golang-work2  golang-work2:v1
docker exec -it golang-work2 sh
```

### 测试httpserver
```shell
curl -I 127.0.0.1/
curl -I http://127.0.0.1/healthz
```

### 推送官方镜像仓库
```shell
docker login 
docker tag <IMAGE ID> singleyuan/golang-work2:v1
docker push singleyuan/golang-work2:v1

#确认上传是否成功
docker search golang-work2 
```

### nsenter查看IP
```shell
#获取容器pid,也可用 ps -aux|grep demo
netstat -tlnp|grep 80 

#因容器以Host模式运行，看到的ip信息与主机相同
nsenter -t <PID> -n ip a
```