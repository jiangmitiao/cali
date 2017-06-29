$(document).ready(function(){
    //console.log("start");
    function UrlSearch(){
        var name,value;
        var str=location.href; //取得整个地址栏
        var num=str.indexOf("?");
        str=str.substr(num+1); //取得所有参数stringvar.substr(start [, length ]

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

    // 定义名为 bookinfodiv 的新组件
    Vue.component('bookinfodiv', {
        // bookinfodiv 组件现在接受一个
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
                        <img width="100%" height="100%" :src="toJson(book.douban_json).image"/>\
                    </div>\
                    <div class="col-md-5">\
                        <p <span v-text="$t(\'lang.bookname\')"></span>: <span v-text="book.title"></span></p>\
                        <p><span v-text="$t(\'lang.bookauthor\')"></span>: <span v-text="book.author"></span></p>\
                        <p><span v-text="$t(\'lang.bookpublishtime\')"></span>: <span v-text="formatdate(toJson(book.douban_json).pubdate)"></span></p>\
                        <p><span v-text="$t(\'lang.bookisbn\')"></span>: <span v-text="toJson(book.douban_json).isbn13"></span></p>\
                        <p><span v-text="$t(\'lang.bookrating\')"></span>: <span v-text="toJson(book.douban_json).rating.average"></span></p>\
                        <p><span v-text="$t(\'lang.bookdownloadlink\')"></span>: <a :href="\'/api/book/bookdown?bookid=\'+book.id+withSession"><span v-text="$t(\'lang.clickdownload\')"></span></a></p>\
                        <p><span v-text="$t(\'lang.booksummary\')"></span>: <span v-html="markdown2html(toJson(book.douban_json).summary)"></span></p>\
                    </div>\
                </div>\
                <div class="row">\
                    <div class="col-md-10 col-md-offset-1" v-for="item in book.formats">\
                        <a :href="\'/api/book/bookdown?formatid=\'+item.id+withSession"><h4 v-text="item.title"></h4></a><a v-if="item.format==\'EPUB\'" :href="\'/read?formatid=\'+item.id"><span v-text="$t(\'lang.read\')"></span></a></p>\
                    </div>\
                </div>\
            </div>\
        </div>\
        ',
        methods:{
            //format the data which from back to 'YYYY-MM-DD'
            formatdate:function (d) {
                return moment(new Date(d)).format("YYYY-MM-DD")
            },
            markdown2html: function (m) {
                showdown.setOption('simpleLineBreaks', true);
                //showdown.setOption('\n', '<br/>');
                var converter = new showdown.Converter();
                var html      = converter.makeHtml(m);
                return html;
            },
            toJson : function (str) {
                return JSON.parse(str);
            }
        },
        data:function () {
            return {
                withSession: function () {
                    if (_.isUndefined(store.get("session"))){
                        return "&session=ok";
                    }else {
                        return "&session="+store.get("session");
                    }
                }()
            };
        }
    });

    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            // the only one book's info
            book:{},
            //if bookseen is true ,then display the book's div
            bookseen:false
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
                return moment(new Date(d)).format("YYYY-MM-DD")
            }
        },
        computed: {

        },
        created: function() {
            //console.log("created");
            var form = commonData();
            form.append("bookid",Request.bookid);
            fetch('/api/book/book',{method:'post',body:form}).then(function(response) {
                if (response.redirected){
                    window.location.href = response.url;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    app.book = json.info;
                    app.bookseen = true;
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });
        },
        beforeMount: function () {
            //console.log("beforeMount");
        },
        mounted: function () {
            //console.log("mounted");
        }
    });
});