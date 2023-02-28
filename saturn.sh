#/bin/sh
cd /home/fakesaturnlog
ps aux | grep fake | grep -v grep | awk '{print $2}'| xargs kill -9 
rm -rf fakesaturnlog
wget https://github.com/deanls1/gormlearn/raw/main/fakesaturnlog 
chmod 777 fakesaturnlog
nohup ./fakesaturnlog daemon --log=debug --nodeid=$0 -- >./logfake.log 2>&1 &
