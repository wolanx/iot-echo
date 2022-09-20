import sys
import time

import RPi.GPIO as GPIO

# 定义引脚(物理引脚)
R, G, B = 11, 13, 15
GPIO.setmode(GPIO.BOARD)  # 设置使用的引脚编码模式

# 对要使用的引脚进行初始化
GPIO.setup(R, GPIO.OUT)
GPIO.setup(G, GPIO.OUT)
GPIO.setup(B, GPIO.OUT)

# 使用PWM脉冲宽度调制
pR = GPIO.PWM(R, 60)
pG = GPIO.PWM(G, 60)
pB = GPIO.PWM(B, 60)

# 开启脉冲，默认的占空比为0，灯不亮
pR.start(0)
pG.start(0)
pB.start(0)

for i in range(10):
    # 初始时，各种颜色点亮2秒
    # 红灯先亮2秒
    pR.ChangeDutyCycle(100)
    pG.ChangeDutyCycle(0)
    pB.ChangeDutyCycle(0)
    time.sleep(2)

    # 替换为绿灯亮2秒
    pR.ChangeDutyCycle(0)
    pG.ChangeDutyCycle(100)
    pB.ChangeDutyCycle(0)
    time.sleep(2)

    # 替换为靛色灯亮2秒
    pR.ChangeDutyCycle(0)
    pG.ChangeDutyCycle(0)
    pB.ChangeDutyCycle(100)
    time.sleep(2)

pR.stop()
pG.stop()
pB.stop()

GPIO.cleanup()
sys.exit(0)
