```shell
#create
docker build -t golang-work2:v1 .
docker images golang-work2:v1

#use
docker run --net host -itd  --name golang-work2  golang-work2:v1
docker exec -it golang-work2 sh

#test
test curl -I 127.0.0.1/
curl -I http://127.0.0.1/healthz
```