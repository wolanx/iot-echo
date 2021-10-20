# iot-echo

> iot 客户端，默认支持 `阿里云 iothub`，支持 `私有 iothub-echo`

项目 | 地址 | 描述
----|-----|-----
iot-echo | [github.com/zx5435/iot-echo](https://github.com/zx5435/iot-echo) | 设备端(go) 采集数据，发数据
iothub-echo | [github.com/zx5435/iothub-echo](https://github.com/zx5435/iothub-echo) | 服务端(java) 收数据，处理数据


## Feature
- [x] config
  - config
  - inputs
- [x] rrpc /sys/{pk}/{dn}/rrpc/request/{uuid} - /sys/{pk}/{dn}/rrpc/response/{uuid}
- [ ] ota
- protocol
  - [x] modbus rtu tpc
  - [ ] genibus



## Config demo
```yaml
provider: aliyun
server:
  host: xxx.iot-as-mqtt.cn-shanghai.aliyuncs.com
  tls: true
device:
  productKey: xxx
  deviceName: xxx
  deviceSecret: xxx
params:
metric:
```
```yaml
channels:
  - name: tcpModbusLocal
    network: tcp
    endpoint: localhost:502
    protocol: modbus
attributes:
  - name: c
    channelRefName: tcpModbusLocal
    slaveId: 1
    address: 5
    dataType: int
  - name: d
    channelRefName: tcpModbusLocal
    slaveId: 1
    address: 5
    dataType: float
```
