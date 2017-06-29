# Cali

[![Build Status](https://www.travis-ci.org/jiangmitiao/cali.svg?branch=master)](https://www.travis-ci.org/jiangmitiao/cali)
[![GitHub release](https://img.shields.io/github/release/jiangmitiao/cali.svg)](https://github.com/jiangmitiao/cali/releases)
[![license](https://img.shields.io/github/license/jiangmitiao/cali.svg)](https://github.com/jiangmitiao/cali/blob/master/LICENSE)
[![Language](https://img.shields.io/badge/language-go1.8.1-brightgreen.svg)](https://github.com/golang/go/tree/release-branch.go1.8)
[![framework](https://img.shields.io/badge/framework-revel0.16.0-brightgreen.svg)](https://github.com/revel/revel/tree/v0.16.0)

# 欢迎来到Cali

[English](https://github.com/jiangmitiao/cali/blob/master/README.md)

这是一个基于 [Revel](http://revel.github.io/) 开发的在线图书馆.

你可以通过 [Cali](https://github.com/jiangmitial/cali) 管理你的书籍.

它的功能有:

* 添加书籍
* 分类查看书籍
* 从 douban.com 获取更多关于查看书籍的信息
* 下载你想阅读的书籍
* 在线阅读(目前仅支持epub格式)
* 有限的用户注册和管理功能
* 搜索
* 删除书籍 (计划中)
* 改变书籍的信息 (计划中)
* 其他...

# 使用说明


### 下载

#### 源码

```shell
go get -u -v github.com/revel/revel
go get -u -v github.com/revel/cmd/revel
go get -u -v github.com/jiangmitiao/cali
```
#### 发行版（linux X64）

[Releases](https://github.com/jiangmitiao/cali/releases)

### 修改配置文件

你肯定能找到这个 `conf/app.conf`

接下来，修改它:
```ini
books.path = 书籍存放地址 #/home/gavin/uploadpath/
books.uploadpath =书籍上传地址 #/tmp/
```

### 打开程序:

源码:
```
revel run github.com/jiangmitiao/cali
```
or 发行版（linux X64）:
```
sh run.sh
```


### 在浏览器中打开 http://localhost:9000/ 你会看到:

![index_cn.png](index_cn.png "")

# 修改日志

* **v0.1.0**
    * **升级系统** 不再需要额外的程序支持，例如Calibre
    * 在首页展示更多的书

* **v0.0.4**
    * 增加下载和浏览记录
    * 修正注册bug

* **v0.0.3**
    * 搜索功能
    * 修正移动端bug
    * 其他

* **v0.0.2**
    * 上传书籍功能
    * 在线阅读 (目前仅支持epub)
    * 用户注册与管理

* **v0.0.1**
    * 可以以6种分类方式查看书籍
    * 展示书籍信息，下载书籍
    * 从 [douban](douban.com) 获取额外的书籍信息并展示



# 开发环境

Gogland 1.0
