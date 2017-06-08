# Cali

[![Build Status](https://www.travis-ci.org/jiangmitiao/cali.svg?branch=master)](https://www.travis-ci.org/jiangmitiao/cali)

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


