#!/bin/sh
src_dir=/mnt/hgfs/winshare/pi_client_wifi/src/  
dest_dir=/home/pi/pi_client_wifi/src  
host=192.168.199.221
username=root  
password=root  
  
# 目录不存在，则创建，如果存在先删除再创建  
if [ ! -d $dest_dir ]; then  
  mkdir -p $dest_dir  
else  
  rm -rf $dest_dir  
 # mkdir -p $dest_dir  
fi  
  
# 将远程服务器上的文件拷贝到本机  
./expect_scp $host $port $username $password $src_dir $dest_dir  
  
echo "end" 
