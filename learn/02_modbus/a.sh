CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o modbus .
scp modbus root@172.16.1.175:/home/root/

#CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w"
#scp iot-echo root@172.16.1.175:/usr/local/bin/
