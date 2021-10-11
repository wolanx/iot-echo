# modbus rtu

```yaml
channels:
  - name: mbRtu
    network: rtu
    endpoint: /dev/ttyO2
    protocol: modbus
  - name: mbTcp
    network: tcp
    endpoint: 192.168.30.66:502
    protocol: modbus
```

# modbus tcp

```yaml
attributes:
  - name: r6
    channelRefName: mbRtu
    slaveId: 10
    address: 6
    dataType: float
```

# GeNi

```yaml
protocol: rtu
endpoint: /dev/ttyO4
slaveId: 1 # 0x01
address: 20
dataType: float # len 2
```
