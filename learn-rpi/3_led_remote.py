"""
红外遥控器
    NEC协议
    文档 https://pan.baidu.com/s/1MteiSMpdFRif2KoJpL5TNg    提取码：t52l
    淘宝 https://item.taobao.com/item.htm?spm=a1z09.2.0.0.67212e8dNqGDJW&id=548823910606&_u=knhqssn86e7

控制 led RGB
    “1”    “2”    “3”
     45    46     47
"""
import RPi.GPIO as GPIO

GPIO.setmode(GPIO.BOARD)  # 设置使用的引脚编码模式

# 定义开关引脚
sw = 40
# 定义蜂鸣器引脚
fm = 11
# 进行开关引脚的初始化，设置为输入引脚，且默认为高电平
GPIO.setup(sw, GPIO.IN, pull_up_down=GPIO.PUD_UP)
# 进行蜂鸣器引脚的初始化，因为是低电平触发，初始时设置为高电平
GPIO.setup(fm, GPIO.OUT, initial=GPIO.HIGH)


# 定义状态变化的回调函数
def switch(channel):
    val = GPIO.input(channel)
    print(val)
    if val:
        # 高电平的开关松开
        GPIO.output(fm, GPIO.HIGH)
    else:
        # 低电平为开关按下
        GPIO.output(fm, GPIO.LOW)


# 添加输入引脚电平变化的回调函数
GPIO.add_event_detect(sw, GPIO.BOTH, callback=switch, bouncetime=200)

# 开启循环
while True:
    pass
