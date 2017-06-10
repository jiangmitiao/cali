# Cali

[![Build Status](https://www.travis-ci.org/jiangmitiao/cali.svg?branch=master)](https://www.travis-ci.org/jiangmitiao/cali)
[![GitHub release](https://img.shields.io/github/release/jiangmitiao/cali.svg)](https://github.com/jiangmitiao/cali/releases)
[![license](https://img.shields.io/github/license/jiangmitiao/cali.svg)](LICENSE)
[![Language](https://img.shields.io/badge/language-go1.8.1-brightgreen.svg)](https://github.com/golang/go/tree/release-branch.go1.8)
[![framework](https://img.shields.io/badge/framework-revel0.16.0-brightgreen.svg)](https://github.com/revel/revel/tree/v0.16.0)

# Welcome to Cali

A Web Library developed by [Revel](http://revel.github.io/).

# Usages

### Get Calibre

you shoud have [calibre](https://calibre-ebook.com/) to manage your books like *.epub,*.mobi ,.etc

then remenber your calibre library's path and the `metadata.db`.


### Download SourceCode or Releases

#### SourceCode

```shell
go get -u -v github.com/revel/revel
go get -u -v github.com/revel/cmd/revel
go get -u -v github.com/jiangmitiao/cali
```
#### Releases

[Releases](https://github.com/jiangmitiao/cali/releases)

### Modify Configuration

you should open `conf/app.conf`

then modify there:
```ini
books.path = your library        #/home/gavin/Calibre 书库/
sqlitedb.path = the calibre's db #/home/gavin/Calibre 书库/metadata.db
``` 

### Start the web server:

source code:
```
revel run github.com/jiangmitiao/cali
```
or releases:
```
sh run.sh
```


### Go to http://localhost:9000/ and you'll see:
```
 your library 
```



# Help


