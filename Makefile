default:
	cat Makefile

build:
	#CGO_ENABLED=0 go build -o filebeat
	docker build -f Dockerfile -t wolanx/iot-echo .
	#docker push wolanx/iot-echo

arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w"
	scp iot-echo root@192.90.3.204:/usr/local/bin/
