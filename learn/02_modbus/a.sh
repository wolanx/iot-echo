CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o modbus .
scp modbus root@172.16.0.29:/home/root/
