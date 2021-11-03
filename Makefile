default:
	cat Makefile

b:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w"
	scp iot-echo root@192.90.3.204:/usr/local/bin/
