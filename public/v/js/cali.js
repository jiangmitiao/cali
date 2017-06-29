$(document).ready(function(){
    //console.log("start");
    // 定义名为 bookdiv 的新组件
    Vue.component('bookdiv', {
        // bookdiv 组件现在接受一个
        // 这个属性名为 book。
        props: ['book'],
        template: '\
        <div class="col-lg-3 col-md-3 col-sm-3 col-xs-6">\
            <div class="content-box">\
                <div class="panel-body text-center">\
                    <a :href="\'/book?bookid=\'+book.id" target="_blank">\
                        <img class="cover" :src="toJson(book.douban_json).image" width="80%" height="80%"/>\
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

    // 定义名为 categorydiv 的新组件
    Vue.component('categorydiv', {
        // categorydiv 组件现在接受一个
        // 这个属性名为 category。
        props: ['category'],
        template: '\
        <li @click="categoryclick(category)"><a href="#"><i class="glyphicon glyphicon-star"></i><span v-text="category.category"></span></a></li>\
        ',
        methods:{
            //return a sub string ,sub's length is max .if src string not equals result the result add '...'
            categoryclick : function (c) {
                //alert(c.id);
                app.categoryid = c.id;
                app.categoryname = c.category;
                app.showbooks();
            }
        }
    });

    //the instance is only one html's Vue's instance on public.html
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            categories :[],
            categoryid:"",
            categoryname:"",

            // the books control one div where is books which has 8 items
            books:[],

        },
        methods: {
            showbooks:function () {
                //books展示分页
                if (app.categoryid == "") {
                    return
                }
                var form = commonData();
                form.append("categoryid", app.categoryid);
                fetch('/api/book/bookscount', {method: 'post', body: form}).then(function (response) {
                    if (response.redirected) {
                        window.location.href = response.url;
                    }
                    return response.json();
                }).then(function (json) {
                    if (json.statusCode == 200) {
                        $('#bookspage').pagination({
                            dataSource: function (done) {
                                var tmp = [];
                                for (var i = 0; i < json.info; i++) {
                                    tmp.push(i);
                                }
                                return done(tmp);
                            },
                            pageRange: 1,
                            totalNumber: json.info,
                            pageSize: 10,
                            showGoInput: true,
                            showGoButton: true,
                            callback: function (data, pagination) {
                                var form = commonData();
                                form.append("start", _.min(data));
                                form.append("limit", data.length);
                                form.append("categoryid", app.categoryid);
                                fetch('/api/book/books', {method: 'post', body: form}).then(function (response) {
                                    if (response.redirected) {
                                        window.location.href = response.url;
                                    }
                                    return response.json();
                                }).then(function (json) {
                                    if (json.statusCode == 200) {
                                        app.books = json.info
                                    }
                                }).catch(function (ex) {
                                    console.log('parsing failed', ex)
                                });
                            }
                        });
                    }
                }).catch(function (ex) {
                    console.log('parsing failed', ex)
                });
            }
        },
        computed: {

        },
        watched:{

        },
        created: function() {
            fetch('/api/category/all',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    window.location.href = response.url;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    app.categories = json.info
                    if (app.categories.length !=0){
                        //app.categoryid = app.categories[0].id;
                        //app.categoryname = app.categories[0].category;
                    }
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });

        },
        beforeMount: function () {

        },
        mounted: function () {
            //console.log("mounted");
        },
        activated:function () {
            //console.log("activated");

        }
    });
});