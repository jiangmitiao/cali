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
                <div class="panel-body">\
                    <a :href="\'/public/v/book.html?bookid=\'+book.id" target="_blank">\
                        <img class="cover" :src="\'/book/bookimage?bookid=\'+book.id" width="100%" height="100%"/>\
                    </a>\
                    <p class="text-center">\
                        <span v-text="maxstring(book.title,10)" :title="book.title" style="word-break: keep-all;white-space: nowrap;"></span>\
                    </p>\
                    <p class="text-center"><span v-text="maxstring(book.name,10)"></span></p>\
                    <p class="text-center"><span v-text="$t(\'lang.rating\')"></span>:<span v-text="book.rating"></span></p>\
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
            }
        }
    });

    // 定义名为 tagdiv 的新组件
    Vue.component('tagdiv', {
        // tagdiv 组件现在接受一个
        // 这个属性名为 tag。
        props: ['tag'],
        template: '\
            <a @click="tagclick(tag.id)" class="btn btn-default" style="margin-bottom: 25px;margin-left: 25px;">\
                <span v-text="tag.name"></span>\
            </a>\
        '
        ,
        methods:{
            //the tagdiv component's method.if click the <a/> ,then invoke the methods ,then invoke the Vue's instance app's methods tagclick
            tagclick:function (tagid) {
                app.tagclick(tagid);
            }
        }
    });

    // 定义名为 authordiv 的新组件
    Vue.component('authordiv', {
        // tagdiv 组件现在接受一个
        // 这个属性名为 tag。
        props: ['author'],
        template: '\
        <a @click="authorclick(author.id)" class="btn btn-default" style="margin-bottom: 25px;margin-left: 25px;">\
            <span v-text="author.name"></span>\
        </a>\
        ',
        methods:{
            //the authordiv component's method.if click the <a/> ,then invoke the methods ,then invoke the Vue's instance app's methods authorclick
            authorclick:function (tagid) {
                app.authorclick(tagid);
            }
        }
    });

    // 定义名为 languagediv 的新组件
    Vue.component('languagediv', {
        // languagediv 组件现在接受一个
        // 这个属性名为 language。
        props: ['language'],
        template: '<a @click="languageclick(language.id)" class="btn btn-default"><span v-text="language.lang_code"></span></a>',
        methods:{
            //the languagediv component's method.if click the <a/> ,then invoke the methods ,then invoke the Vue's instance app's methods languageclick
            languageclick:function (lang_code) {
                app.languageclick(lang_code);
            }
        }
    });

    //the instance is only one html's Vue's instance on public.html
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            // the hotbooks control one div where is hotbooks which has 8 items
            hotbooks:[],
            // the newbooks control one div where is newbooks which has 8 items
            newbooks:[],
            // the discover control one div where is discover which has 8 items
            discover:[],
            // the categories control one div where is categories which has 8 items
            categories:[],
            // the authors control one div where is authors which has 8 items
            authors:[],
            // the language control one div where is language which has 8 items
            language:[],
            // booksseen is a object which has 6 attrs.one boolean attr control one div which in hotbooks,newbooks,discover,categories,authors,language.
            booksseen:{},
            // the tags is an array ,which has a struct like '{id:0,name:"history"}'.and it is first display in categories.
            tags:[],  //tags
            // tagsseen is a condition for dispaly or hide the tags.if tagsseen is true the display the tags then hide the categories.
            tagsseen:true,
            // one page display 20 tags
            tagssize:20,
            // the authornames is an array,like the tags
            authornames:[],//authors
            // like tagsize
            authorsize:20,
            // authorsseen like tagsseen
            authorsseen : true,
            // the languagenames like tags .item struct like '{id:0,lang_code:"zho"}'
            languagenames:[],//languages
            // the languagesseen like tagsseen
            languagesseen:true
        },
        methods: {
            // when user click one item on left bar,then invoke this method.to display one div
            changeseen:function (e) {
                this.booksseen = {};
                this.booksseen["hotbooks"] = false;
                this.booksseen["newbooks"] = false;
                this.booksseen["discover"] = false;
                this.booksseen["categories"] = false;
                this.booksseen["authors"] = false;
                this.booksseen["language"] = false;
                this.booksseen[e] = true;
                // when user click the three items ,then display first display
                if (e=="categories"){
                    this.tagsseen = true;
                }
                if (e=="authors"){
                    this.authorsseen = true;
                }
                if (e=="language"){
                    this.languagesseen = true;
                }
            },
            // the categories first display,when click tag's item ,then hide first div,fetch 8 books items which has the click's item'tag to display
            tagclick:function (tagid) {
                var form = commonData();
                form.set("tagid",tagid);
                fetch('/book/tagbookscount',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        window.location.href = response.url;
                    }
                    return response.json();
                }).then(function(json) {
                    // when result code is 200, then rending div
                    if (json.statusCode ==200){
                        $('#tagpage').pagination({
                            dataSource:function (done) {
                                var tmp = [];
                                for(var i=0; i<json.info;i++){
                                    tmp.push(i)
                                }
                                return done(tmp);
                            },
                            pageRange:1,
                            totalNumber:json.info,
                            pageSize: 8,
                            showGoInput: true,
                            showGoButton: true,
                            callback: function(data, pagination) {
                                var form = commonData();
                                form.set("start",_.min(data));
                                form.set("limit",data.length);
                                form.set("tagid",tagid);
                                fetch('/book/tagbooks',{method:'post',body:form}).then(function(response) {
                                    if (response.redirected){
                                        window.location.href = response.url;
                                    }
                                    return response.json();
                                }).then(function(json) {
                                    if (json.statusCode ==200){
                                        app.categories = json.info
                                    }
                                }).
                                catch(function(ex) {
                                    console.log('parsing failed', ex)
                                });
                            }
                        });
                    }
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex)
                });
                this.tagsseen = false;
            },
            // like tags click
            authorclick:function (authorid) {
                var form = commonData();
                form.set("authorid",authorid);
                fetch('/book/authorbookscount',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        window.location.href = response.url;
                    }
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ==200){
                        $('#authorspage').pagination({
                            dataSource:function (done) {
                                var tmp = [];
                                for(var i=0; i<json.info;i++){
                                    tmp.push(i)
                                }
                                return done(tmp);
                            },
                            pageRange:1,
                            totalNumber:json.info,
                            pageSize: 8,
                            showGoInput: true,
                            showGoButton: true,
                            callback: function(data, pagination) {
                                var form = commonData();
                                form.set("start",_.min(data));
                                form.set("limit",data.length);
                                form.set("authorid",authorid);
                                fetch('/book/authorbooks',{method:'post',body:form}).then(function(response) {
                                    if (response.redirected){
                                        window.location.href = response.url;
                                    }
                                    return response.json();
                                }).then(function(json) {
                                    if (json.statusCode ==200){
                                        app.authors = json.info
                                    }
                                }).
                                catch(function(ex) {
                                    console.log('parsing failed', ex)
                                });
                            }
                        });
                    }
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex)
                });
                this.authorsseen = false;
            },
            // like tags click
            languageclick:function (lang_code) {
                var form = commonData();
                form.set("lang_code",lang_code);
                fetch('/book/languagebookscount',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        window.location.href = response.url;
                    }
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ==200){
                        $('#languagespage').pagination({
                            dataSource:function (done) {
                                var tmp = [];
                                for(var i=0; i<json.info;i++){
                                    tmp.push(i)
                                }
                                return done(tmp);
                            },
                            pageRange:1,
                            totalNumber:json.info,
                            pageSize: 8,
                            showGoInput: true,
                            showGoButton: true,
                            callback: function(data, pagination) {
                                var form = commonData();
                                form.set("start",_.min(data));
                                form.set("limit",data.length);
                                form.set("lang_code",lang_code);
                                fetch('/book/languagebooks',{method:'post',body:form}).then(function(response) {
                                    if (response.redirected){
                                        window.location.href = response.url;
                                    }
                                    return response.json();
                                }).then(function(json) {
                                    if (json.statusCode ==200){
                                        app.language = json.info
                                    }
                                }).
                                catch(function(ex) {
                                    console.log('parsing failed', ex)
                                });
                            }
                        });
                    }
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex)
                });
                this.languagesseen = false;
            }
        },
        computed: {
            // we not use the data,we use the computed data
            hotbooks_computed : function () {
                return this.hotbooks;
            },
            newbooks_computed : function () {
                return this.newbooks;
            },
            discover_computed : function () {
                return this.discover;
            },
            categories_computed : function () {
                return this.categories;
            },
            authors_computed:function () {
                return this.authors;
            },
            languages_computed:function () {
                return this.language;
            }
        },
        created: function() {
            //when page's instance created,we should get the data to render the page
            //console.log("created");
            //hotbooks展示分页
            fetch('/book/bookscount',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    window.location.href = response.url;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    $('#hotbookspage').pagination({
                        dataSource:function (done) {
                            var tmp = [];
                            for(var i=0; i<json.info;i++){
                                tmp.push(i)
                            }
                            return done(tmp);
                        },
                        pageRange:1,
                        totalNumber:json.info,
                        pageSize: 8,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            var form = commonData();
                            form.set("start",_.min(data));
                            form.set("limit",data.length);
                            fetch('/book/ratingbooks',{method:'post',body:form}).then(function(response) {
                                if (response.redirected){
                                    window.location.href = response.url;
                                }
                                return response.json();
                            }).then(function(json) {
                                if (json.statusCode ==200){
                                    app.hotbooks = json.info
                                }
                            }).
                            catch(function(ex) {
                                console.log('parsing failed', ex)
                            });
                        }
                    });
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });

            //newbooks展示分页
            fetch('/book/bookscount',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    window.location.href = response.url;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    $('#newbookspage').pagination({
                        dataSource:function (done) {
                            var tmp = [];
                            for(var i=0; i<json.info;i++){
                                tmp.push(i)
                            }
                            return done(tmp);
                        },
                        pageRange:1,
                        totalNumber:json.info,
                        pageSize: 8,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            var form = commonData();
                            form.set("start",_.min(data));
                            form.set("limit",data.length);
                            fetch('/book/newbooks',{method:'post',body:form}).then(function(response) {
                                if (response.redirected){
                                    window.location.href = response.url;
                                }
                                return response.json();
                            }).then(function(json) {
                                if (json.statusCode ==200){
                                    app.newbooks = json.info
                                }
                            }).
                            catch(function(ex) {
                                console.log('parsing failed', ex)
                            });
                        }
                    });
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });

            //discover展示分页
            fetch('/book/bookscount',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    window.location.href = response.url;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    $('#discoverpage').pagination({
                        dataSource:function (done) {
                            var tmp = [];
                            for(var i=0; i<json.info;i++){
                                tmp.push(i)
                            }
                            return done(tmp);
                        },
                        pageRange:1,
                        totalNumber:json.info,
                        pageSize: 8,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            var form = commonData();
                            form.set("start",_.min(data));
                            form.set("limit",data.length);
                            fetch('/book/discoverbooks',{method:'post',body:form}).then(function(response) {
                                if (response.redirected){
                                    window.location.href = response.url;
                                }
                                return response.json();
                            }).then(function(json) {
                                if (json.statusCode ==200){
                                    app.discover = json.info
                                }
                            }).
                            catch(function(ex) {
                                console.log('parsing failed', ex)
                            });
                        }
                    });
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });

            //tags展示分页
            fetch('/tag/tagscount',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    window.location.href = response.url;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    $('#tagspage').pagination({
                        dataSource:function (done) {
                            var tmp = [];
                            for(var i=0; i<json.info;i++){
                                tmp.push(i)
                            }
                            return done(tmp);
                        },
                        pageRange:1,
                        totalNumber:json.info,
                        pageSize: app.tagssize,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            var form = commonData();
                            form.set("start",_.min(data));
                            form.set("limit",data.length);
                            fetch('/tag/tags',{method:'post',body:form}).then(function(response) {
                                if (response.redirected){
                                    window.location.href = response.url;
                                }
                                return response.json();
                            }).then(function(json) {
                                if (json.statusCode ==200){
                                    app.tags = json.info
                                }
                            }).
                            catch(function(ex) {
                                console.log('parsing failed', ex)
                            });
                        }
                    });
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });

            //authornames展示分页
            fetch('/author/authorscount',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    window.location.href = response.url;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    $('#authornamespage').pagination({
                        dataSource:function (done) {
                            var tmp = [];
                            for(var i=0; i<json.info;i++){
                                tmp.push(i)
                            }
                            return done(tmp);
                        },
                        pageRange:1,
                        totalNumber:json.info,
                        pageSize: app.authorsize,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            var form = commonData();
                            form.set("start",_.min(data));
                            form.set("limit",data.length);
                            fetch('/author/authors',{method:'post',body:form}).then(function(response) {
                                if (response.redirected){
                                    window.location.href = response.url;
                                }
                                return response.json();
                            }).then(function(json) {
                                if (json.statusCode ==200){
                                    app.authornames = json.info
                                }
                            }).
                            catch(function(ex) {
                                console.log('parsing failed', ex)
                            });
                        }
                    });
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });

            //languagenames展示分页
            fetch('/language/languagescount',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    window.location.href = response.url;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    $('#languagenamespage').pagination({
                        dataSource:function (done) {
                            var tmp = [];
                            for(var i=0; i<json.info;i++){
                                tmp.push(i)
                            }
                            return done(tmp);
                        },
                        pageRange:1,
                        totalNumber:json.info,
                        pageSize: 8,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            var form = commonData();
                            form.set("start",_.min(data));
                            form.set("limit",data.length);
                            fetch('/language/languages',{method:'post',body:form}).then(function(response) {
                                if (response.redirected){
                                    window.location.href = response.url;
                                }
                                return response.json();
                            }).then(function(json) {
                                if (json.statusCode ==200){
                                    app.languagenames = json.info
                                }
                            }).
                            catch(function(ex) {
                                console.log('parsing failed', ex)
                            });
                        }
                    });
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });


        },
        beforeMount: function () {
            // when instance created and prepare to render page ,we should confirm we display one div
            //console.log("beforeMount");
            this.booksseen["hotbooks"] = true;
            this.booksseen["newbooks"] = false;
            this.booksseen["discover"] = false;
            this.booksseen["categories"] = false;
            this.booksseen["authors"] = false;
            this.booksseen["language"] = false;
        },
        mounted: function () {
            //console.log("mounted");
        },
        activated:function () {
            //console.log("activated");

        }
    });
});