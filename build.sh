#/usr/bin/bash


SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
echo $SCRIPTPATH
rm -r $SCRIPTPATH/build/*
cd $SCRIPTPATH/build
revel package github.com/jiangmitiao/cali prod 
