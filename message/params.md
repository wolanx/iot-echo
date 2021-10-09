# modbus rtu

```yaml
protocol: rtu
endpoint: /dev/ttyO2
slaveId: 10 # 0x0A
address: 6
dataType: float # len 2
```

# modbus tcp

```yaml
protocol: tcp
endpoint: 192.168.30.66:502
slaveId: 10 # 0x0A
address: 6
dataType: float # len 2
```

# GeNi

```yaml
protocol: rtu
endpoint: /dev/ttyO4
slaveId: 1 # 0x01
address: 20
dataType: float # len 2
```
