$(document).ready(function(){



    //public
    // 定义名为 bookdiv 的新组件
    Vue.component('bookdiv', {
        // bookdiv 组件现在接受一个
        // 这个属性名为 book。
        props: ['book'],
        template: '\
        <div class="col-lg-2 col-md-3 col-sm-3 col-xs-6">\
            <div class="content-box">\
                <div class="panel-body text-center">\
                    <a :href="\'/book?bookid=\'+book.id" target="_blank" class="text-center">\
                        <canvas :id="book.id" :src="fakePic(toJson(book.douban_json).image,book.id)"></canvas>\
                    </a>\
                    <p class="text-center">\
                        <a :href="\'/book?bookid=\'+book.id" target="_blank">\
                            <nobr v-text="maxstring(book.title,5)" :title="book.title" style="word-break: keep-all;white-space: nowrap;"></nobr>\
                        </a>\
                    </p>\
                    <p class="text-center">\
                        <a :href="\'/search?q=\'+book.author" target="_blank">\
                        <nobr v-text="maxstring(book.author,5)"></nobr>\
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
                if (result.length !== str.length){
                    result+="...";
                }
                return result;
            },
            toJson : function (str) {
                try {
                    var json = JSON.parse(str);
                    return json
                }catch (e){
                    var json = {};
                    json.image = "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png";
                    json.rating = {};
                    json.rating.average = 0.0;
                    return json;
                }
            },
            fakePic:function (src,id) {
                var img = new Image();
                img.src = src;
                img.onload = function(){
                    myCanvas = document.getElementById(id);
                    myCanvas.width = 100;
                    myCanvas.height = 150;
                    var context = myCanvas.getContext('2d');
                    context.drawImage(img, 0, 0);
                    //var imgdata = context.getImageData(0, 0, img.width, img.height);
                    // 处理imgdata
                };
                return "src"
            }
        }
    });

    // 定义名为 categorydiv 的新组件
    Vue.component('categorydiv', {
        // categorydiv 组件现在接受一个
        // 这个属性名为 category。
        props: ['category','categoryid'],
        template: '\
        <button type="button" @click="categoryclick(category)" :class="\'list-group-item \'+active(category,categoryid)"><i class="glyphicon glyphicon-star"></i><span v-text="category.category"></span></button>\
        ',
        methods:{
            //return a sub string ,sub's length is max .if src string not equals result the result add '...'
            categoryclick : function (c) {
                this.$emit('categoryclick',c);
            },
            active:function (category,categoryid) {
                if (category.id === categoryid){
                    return "active"
                }else {
                    return ""
                }
            }
        }
    });



    //book
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
                        <p><span v-text="$t(\'lang.booksummary\')"></span>: <span v-html="markdown2html(toJson(book.douban_json).summary)"></span></p>\
                    </div>\
                </div>\
                <div class="row">\
                    <div class="col-md-10 col-md-offset-1">\
                        <h4 v-text="$t(\'lang.bookdownloadlink\')"></h4>\
                    </div>\
                </div>\
                <div class="row">\
                    <div class="col-md-10 col-md-offset-1" v-for="item in book.formats">\
                        <a :href="\'/api/book/bookdown?formatid=\'+item.id+withSession"><h4 v-text="item.title+\'.\'+item.format"></h4></a><a v-if="item.format==\'EPUB\'" :href="\'/read?formatid=\'+item.id" target="_blank"><span v-text="$t(\'lang.read\')"></span></a></p>\
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
                try {
                    var json = JSON.parse(str);
                    return json
                }catch (e){
                    var json = {};
                    json.image = "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png";
                    json.rating = {};
                    json.rating.average = 0.0;
                    json.pubdate = "0";
                    json.isbn13 = "0";
                    json.format = "";
                    return json;
                }
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





    //person
    // 定义名为 userdiv 的新组件
    Vue.component('usersdiv', {
        // bookinfodiv 组件现在接受一个
        // 这个属性名为 book。
        props: ['userlist'],
        template: '\
        <table class="table">\
            <thead>\
                <tr>\
                    <th>#</th>\
                    <th>#</th>\
                    <th>#</th>\
                    <th>#</th>\
                </tr>\
            </thead>\
            <tbody>\
                <tr v-for="item in userlist">\
                    <td v-text="item.userName"></td>\
                    <td v-text="item.loginName"></td>\
                    <td v-text="item.email"></td>\
                    <td><a class="btn btn-danger" @click="deleteuser" :id="item.id">delete</a></td>\
                </tr>\
            </tbody>\
        </table>\
        ',
        methods:{
            deleteuser:function (t) {
                var form = commonData();
                form.append("userId",t.target.id);
                fetch('/api/user/delete',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        var tmpJson = {};
                        tmpJson.statusCode = 500;
                        return tmpJson;
                    }
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ==200){
                        //refresh this page
                        $('#userlistpage').pagination($('#userlistpage').pagination('getSelectedPageNum'))
                    }else {
                    }
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex)
                });
            }
        }
    });


    // 定义名为 sysconfigdiv 的新组件
    Vue.component('sysconfigdiv', {
        // sysconfigdiv 组件现在接受一个
        // 这个属性名为 sysconfiglist。
        props: ['sysconfiglist'],
        template: '\
        <table class="table">\
            <thead>\
                <tr>\
                    <th>#</th>\
                    <th>#</th>\
                    <th>#</th>\
                    <th>#</th>\
                </tr>\
            </thead>\
            <tbody>\
                <tr v-for="item in sysconfiglist">\
                    <td v-text="item.key"></td>\
                    <td><input type="text" maxlength="64" v-model="item.value" class="form-control"/></td>\
                    <td v-text="item.comments"></td>\
                    <td><a class="btn btn-info" @click="update" :id="item.id">update</a></td>\
                </tr>\
            </tbody>\
        </table>\
        ',
        methods:{
            update:function (t) {
                app.sysconfigupdate(t.target.id);
            }
        }
    });

    // 定义名为 sysstatusdiv 的新组件
    Vue.component('sysstatusdiv', {
        // sysstatusdiv 组件现在接受一个
        // 这个属性名为 sysstatuslist。
        props: ['sysstatuslist'],
        template: '\
        <table class="table">\
            <thead>\
                <tr>\
                    <th>#</th>\
                    <th>#</th>\
                    <th>#</th>\
                    <th>#</th>\
                </tr>\
            </thead>\
            <tbody>\
                <tr v-for="item in sysstatuslist">\
                    <td v-text="item.key"></td>\
                    <td v-text="item.value"></td>\
                    <td v-text="item.comments"></td>\
                    <td><a class="btn btn-danger" @click="deletestatus" :id="item.id">delete</a></td>\
                </tr>\
            </tbody>\
        </table>\
        ',
        methods:{
            deletestatus:function (t) {
                app.sysstatusdelete(t.target.id);
            }
        }
    });


    // 定义名为 categoriesdiv 的新组件
    Vue.component('categoriesdiv', {
        // categoriesdiv 组件现在接受一个
        // 这个属性名为 categorylist。
        props: ['categorylist'],
        template: '\
        <table class="table">\
            <thead>\
                <tr>\
                    <th>#</th>\
                    <th>#</th>\
                    <th>#</th>\
                    <th>#</th>\
                </tr>\
            </thead>\
            <tbody>\
                <tr v-for="item in categorylist">\
                    <td v-text="item.id"></td>\
                    <td><input type="text" maxlength="12" v-model="item.category" class="form-control"></td>\
                    <td><a class="btn btn-warning" @click="updatecategory(item)">update</a></td>\
                    <td><a class="btn btn-danger" @click="deletecategory(item)">delete</a></td>\
                </tr>\
            </tbody>\
        </table>\
        ',
        methods:{
            updatecategory:function (c) {
                this.$emit('updatecategory',c);
            },
            deletecategory:function (c) {
                this.$emit('deletecategory',c);
            }
        }
    });

    // 定义名为 categoriesselectdiv 的新组件
    Vue.component('categoriesselectdiv', {
        // categoriesselectdiv 组件现在接受一个
        // 这个属性名为 categorylist。
        props: ['categorylist'],
        template: '\
        <div class="form-group">\
            <label for="categoryid" class="col-sm-2 control-label"><span v-text="$t(\'lang.categories\')"></span>:</label>\
            <div class="col-sm-10">\
                <select class="form-control" id="categoryid" name="categoryid">\
                    <option v-if="item.id!=\'default\'" v-for="item in categorylist" :value="item.id"><span v-text=item.category></span></option>\
                </select>\
            </div>\
        </div>\
        ',
        methods:{
            updatecategory:function (c) {
                this.$emit('updatecategory',c);
            },
            deletecategory:function (c) {
                this.$emit('deletecategory',c);
            }
        }
    });

});