1.第一部分  设备操作抽象
hwdemo/ds18b20  golang版本的温度传感器 Demo
hwdemo/led      golang+wiringPi版本的LED Demo
hwdemo/led2      纯golang版本的LED Demo


2.第二部分  调试工具
hug      golang版本数据监控Demo 源码  

tools    编译好的工具
         hug -->模拟下行控制对应的树莓派板子    
         sug -->监控MQTT所有通信的上下行消息


3.第三部分  Pi 温度和LED远程控制系统

monitor  远程控制程序（自动上报温度和开关状态， 读取远程指令进行相关动作）
publish  monitor编译发布程序

monitor 自启动实现：
cp -f monitor /home/pi/monitor/
cp -f appserver.conf /home/pi/monitor/
sudo cp -f monitor.service /lib/systemd/system/
sudo chmod 755 /lib/systemd/system/monitor.service
sudo systemctl enable monitor.service

其他命令：
#启动webtest服务
sudo systemctl start monitor.service

#设置开机自启动
sudo systemctl enable monitor.service

#停止开机自启动
sudo systemctl disable monitor.service

#查看服务当前状态
sudo systemctl status monitor.service

#重新启动服务
sudo systemctl restart monitor.service

#查看所有已启动的服务
sudo systemctl list-units --type=service


4.参考文档
https://coldnew.github.io/f7349436/
https://gobot.io/documentation/platforms/raspi/
http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html