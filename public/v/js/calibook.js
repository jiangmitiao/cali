window.onload = function() {
    console.log("start");
    function UrlSearch(){
        var name,value;
        var str=location.href; //取得整个地址栏
        var num=str.indexOf("?");
        str=str.substr(num+1); //取得所有参数   stringvar.substr(start [, length ]

        var arr=str.split("&"); //各个参数放到数组里
        for(var i=0;i < arr.length;i++){
            num=arr[i].indexOf("=");
            if(num>0){
                name=arr[i].substring(0,num);
                value=arr[i].substr(num+1);
                this[name]=value;
            }
        }
    }
    var Request=new UrlSearch(); //实例化

    //<div class="panel-heading"><div class="panel-title"><span v-text="book.title"></span></div></div><div class="panel-body"><div class="row"><div class="col-md-3 col-md-offset-1"><img width="100%" height="100%" src="/book/bookimage?bookid=2"></div><div class="col-md-5"><p>Book's Name:<span v-text="book.title"></span></p><p>Book's Author:<span v-text="book.name"></span></p><p>Book's Publish Time:<span v-text="book.pubdate"></span></p><p>Book's Update Time:<span v-text="book.timestamp"></span></p><p>Book's ISBN:<span v-text="book.isbn"></span></p><p>Book's Last Modified Time:<span v-text="book.last_modified"></span></p><p>Book's Rating:<span v-text="book.rating"></span></p></div></div></div>
    //https://developers.douban.com/wiki/?title=book_v2#get_isbn_book

    // 定义名为 bookinfodiv 的新组件
    Vue.component('bookinfodiv', {
        // bookinfodiv 组件现在接受一个
        // "prop"，类似于一个自定义属性
        // 这个属性名为 book。
        props: ['book'],
        template: '\
        <div class="content-box-large">\
            <div class="panel-heading">\
                    <div class="panel-title"><span v-text="book.title"></span></div>\
            </div>\
            <div class="panel-body">\
                <div class="row">\
                    <div class="col-md-3 col-md-offset-1">\
                        <img width="100%" height="100%" :src="\'/book/bookimage?bookid=\'+book.id"/>\
                    </div>\
                    <div class="col-md-5">\
                        <p>{{ $t("lang.bookname") }}: <span v-text="book.title"></span></p>\
                        <p>{{ $t("lang.bookauthor") }}: <span v-text="book.name"></span></p>\
                        <p>{{ $t("lang.bookpublishtime") }}: <span v-text="formatdate(book.pubdate)"></span></p>\
                        <p>{{ $t("lang.bookupdatetime") }}: <span v-text="formatdate(book.timestamp)"></span></p>\
                        <p>{{ $t("lang.bookisbn") }}: <span v-text="book.isbn"></span></p>\
                        <p>{{ $t("lang.bookmodifiedtime") }}: <span v-text="formatdate(book.last_modified)"></span></p>\
                        <p>{{ $t("lang.bookrating") }}: <span v-text="book.rating"></span></p>\
                        <p>{{ $t("lang.bookdownloadlink") }}: <a :href="\'/book/bookdown?bookid=\'+book.id">{{ $t("lang.download") }}</a></p>\
                    </div>\
                </div>\
            </div>\
        </div>\
        ',
        methods:{
            formatdate:function (d) {
                return StringFormatter.format("{Date:{format:'YYYY-MM-DD'}}", new Date(d))
            }
        }
    });


    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            book:{},
            bookseen:false,
            doubanbook:{},
            doubanbookseen:false
        },
        methods: {
            markdown2html: function (m) {
                showdown.setOption('simpleLineBreaks', true);
                //showdown.setOption('\n', '<br/>');
                var converter = new showdown.Converter();
                var html      = converter.makeHtml(m);
                return html;
            },
            formatdate:function (d) {
                return StringFormatter.format("{Date:{format:'YYYY-MM-DD'}}", new Date(d))
            }
        },
        computed: {

        },
        created: function() {
            console.log("created");
            console.log(Request.bookid);
            fetch('/book/book?bookid='+Request.bookid).then(function(response) {
                return response.json()
            }).then(function(json) {
                console.log('parsed json', json);
                if (json.statusCode ==200){
                    app.book = json.info;
                    app.bookseen = true;
                    fetch('/book/doubanbook?bookid='+json.info.id).then(function(response) {
                        return response.json()
                    }).then(function(json) {
                        if (json.statusCode ==200){
                            console.log(JSON.parse(json.info));
                            var info = JSON.parse(json.info);
                            if (info.count != undefined & info.count!=0){
                                console.log(info.books[0]);
                                app.doubanbook = info.books[0];
                                app.doubanbookseen = true;
                            }
                        }
                    }).
                    catch(function(ex) {
                        console.log('parsing failed', ex)
                    });
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });
        },
        beforeMount: function () {
            console.log("beforeMount");
            //this.book.title="oookkk"
        },
        mounted: function () {
            console.log("mounted");
        }
    });



}