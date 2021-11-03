json、map、struct 之间的相互转化 https://www.cnblogs.com/liang1101/p/6741262.html

GOOS=linux GOARCH=arm GOARM=7 go build main.go


# ota
流程 https://help.aliyun.com/document_detail/85700.html
js https://developer.aliyun.com/article/718838

# 启动
ls -l /etc/rc* | grep iot
update-rc.d -f iot-echo remove
update-rc.d iot-echo defaults 99
