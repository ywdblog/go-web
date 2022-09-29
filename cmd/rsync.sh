#!/bin/bash

#cd /home/work/go/src/github.com/xiwujie/article/
cd /home/work/go-release/article/
./cmd/build.sh
if [ $? != 0 ]; then
    exit
fi

tmpDir=`date '+%Y%m%d%H%M%S'`
buildDir=build-tmp/$tmpDir
mkdir -p $buildDir
cp -rf cmd $buildDir/
cp -rf conf $buildDir/
rm -f $tmpDir/cmd/*.go


sshpass -p "zz@xiw#2" ssh work@172.16.104.133 "mkdir -p /home/work/www/article-logs; mkdir -p /home/work/www/article-releases/;"
sshpass -p "zz@xiw#2" ssh work@172.16.104.132 "mkdir -p /home/work/www/article-logs; mkdir -p /home/work/www/article-releases/;"


echo "transfer file"
sshpass -p "zz@xiw#2" scp -r $buildDir work@172.16.104.133:/home/work/www/article-releases/
sshpass -p "zz@xiw#2" scp -r $buildDir work@172.16.104.132:/home/work/www/article-releases/

echo "restart server"
sshpass -p "zz@xiw#2" ssh work@172.16.104.133 "rm -f /home/work/www/article; ln -sf /home/work/www/article-releases/$tmpDir /home/work/www/article; cd /home/work/www/article; ln -sf /home/work/www/article-logs logs; ./cmd/start.sh -e production -k restart"
sshpass -p "zz@xiw#2" ssh work@172.16.104.132 "rm -f /home/work/www/article; ln -sf /home/work/www/article-releases/$tmpDir /home/work/www/article; cd /home/work/www/article; ln -sf /home/work/www/article-logs logs; ./cmd/start.sh -e production -k restart"
