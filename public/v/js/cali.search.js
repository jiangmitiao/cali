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

    // 定义名为 bookdiv 的新组件
    Vue.component('bookdiv', {
        // bookdiv 组件现在接受一个
        // 这个属性名为 book。
        props: ['book'],
        template: '\
        <div class="col-lg-2 col-md-2 col-sm-3 col-xs-6">\
            <div class="content-box">\
                <div class="panel-body text-center">\
                    <a :href="\'/book?bookid=\'+book.id" target="_blank">\
                        <img class="cover" :src="toJson(book.douban_json).image" width="95%" height="95%"/>\
                    </a>\
                    <p class="text-center">\
                        <a :href="\'/book?bookid=\'+book.id" target="_blank">\
                            <span v-text="maxstring(book.title,10)" :title="book.title" style="word-break: keep-all;white-space: nowrap;"></span>\
                        </a>\
                    </p>\
                    <p class="text-center">\
                        <a :href="\'/search?q=\'+book.author" target="_blank">\
                        <span v-text="maxstring(book.author,10)"></span>\
                        </a>\
                    </p>\
                    <p class="text-center badge" style="background-color: #2c3742"><span v-text="$t(\'lang.rating\')"></span>:<span  v-text="toJson(book.douban_json).rating.average"></span></p>\
                    <br>\
                </div>\
            </div>\
        </div>\
        ',
        methods:{
            //return a sub string ,sub's length is max .if src string not equals result the result add '...'
            maxstring : function (str,max) {
                var result = str.substr(0,max);
                if (result != str){
                    result+="...";
                }
                return result;
            },
            toJson : function (str) {
                return JSON.parse(str);
            }
        }
    });



    //the instance is only one html's Vue's instance on search.html
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            searchbooks:[],
        },
        methods: {
        },
        computed: {
        },
        created: function() {
            //when page's instance created,we should get the data to render the page
            //console.log("created");
            //searchbooks展示分页
            var form = commonData();
            form.append("q",Request.q);
            fetch('/api/book/searchcount',{method:'post',body:form}).then(function(response) {
                if (response.redirected){
                    window.location.href = response.url;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200 && json.info !=0){
                    $('#searchbookspage').pagination({
                        dataSource:function (done) {
                            var tmp = [];
                            for(var i=0; i<json.info;i++){
                                tmp.push(i)
                            }
                            return done(tmp);
                        },
                        pageRange:1,
                        totalNumber:json.info,
                        pageSize: 24,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            var form = commonData();
                            form.append("start",_.min(data));
                            form.append("limit",data.length);
                            form.append("q",Request.q);
                            fetch('/api/book/search',{method:'post',body:form}).then(function(response) {
                                if (response.redirected){
                                    window.location.href = response.url;
                                }
                                return response.json();
                            }).then(function(json) {
                                if (json.statusCode ==200){
                                    app.searchbooks = json.info
                                }
                            }).
                            catch(function(ex) {
                                console.log('parsing failed', ex)
                            });
                        }
                    });
                }{

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
        },
        activated:function () {
            //console.log("activated");

        }
    });
});