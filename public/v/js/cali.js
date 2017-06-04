window.onload = function() {
    console.log("start");
    // 定义名为 bookdiv 的新组件
    Vue.component('bookdiv', {
        // bookdiv 组件现在接受一个
        // "prop"，类似于一个自定义属性
        // 这个属性名为 book。
        props: ['book'],
        template: '\
        <div class="col-md-3">\
            <div class="content-box">\
                <div class="panel-body">\
                    <a :href="\'/public/v/book.html?bookid=\'+book.id" target="_blank">\
                        <img :src="\'/book/bookimage?bookid=\'+book.id" width="100%" height="100%"/>\
                    </a>\
                    <p class="text-center">\
                        <span v-text="book.title" style="word-break: keep-all;white-space: nowrap;"></span>\
                    </p>\
                    <p class="text-center"><span v-text="book.name"></span></p>\
                    <p class="text-center">{{ $t("lang.rating") }}:<span v-text="book.rating"></span></p>\
                    <br>\
                </div>\
            </div>\
        </div>\
        '
    });

    // 定义名为 tagdiv 的新组件
    Vue.component('tagdiv', {
        // tagdiv 组件现在接受一个
        // "prop"，类似于一个自定义属性
        // 这个属性名为 tag。
        props: ['tag'],
        template: '<a @click="tagclick(tag.id)" class="btn btn-default"><span v-text="tag.name"></span></a>',
        methods:{
            tagclick:function (tagid) {
                //console.log(tagid);
                app.tagclick(tagid);
            }
        }
    });

    // 定义名为 authordiv 的新组件
    Vue.component('authordiv', {
        // tagdiv 组件现在接受一个
        // "prop"，类似于一个自定义属性
        // 这个属性名为 tag。
        props: ['author'],
        template: '<a @click="authorclick(author.id)" class="btn btn-default"><span v-text="author.name"></span></a>',
        methods:{
            authorclick:function (tagid) {
                //console.log(tagid);
                app.authorclick(tagid);
            }
        }
    });

    // 定义名为 languagediv 的新组件
    Vue.component('languagediv', {
        // languagediv 组件现在接受一个
        // "prop"，类似于一个自定义属性
        // 这个属性名为 language。
        props: ['language'],
        template: '<a @click="languageclick(language.id)" class="btn btn-default"><span v-text="language.lang_code"></span></a>',
        methods:{
            languageclick:function (lang_code) {
                //console.log(lang_code);
                app.languageclick(lang_code);
            }
        }
    });

    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            hotbooks:[],
            newbooks:[],
            discover:[],
            categories:[],
            authors:[],
            language:[],
            booksseen:{},
            tags:[],  //tags
            tagsseen:true,
            authornames:[],//authors
            authorsseen : true,
            languagenames:[],//languages
            languagesseen:true
        },
        methods: {
            changeseen:function (e) {
                //console.log(this);
                //console.log(e);
                this.booksseen = {};
                this.booksseen["hotbooks"] = false;
                this.booksseen["newbooks"] = false;
                this.booksseen["discover"] = false;
                this.booksseen["categories"] = false;
                this.booksseen["authors"] = false;
                this.booksseen["language"] = false;
                this.booksseen[e] = true;
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
            tagclick:function (tagid) {
                //console.log("tagid"+tagid)
                fetch('/book/tagbookscount?tagid='+tagid).then(function(response) {
                    return response.json()
                }).then(function(json) {
                    //console.log('parsed json', json);
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
                                // template method of yourself
                                //console.log(data);
                                //console.log(pagination);
                                fetch('/book/tagbooks?start='+_.min(data)+'&limit='+data.length+'&tagid='+tagid).then(function(response) {
                                    return response.json()
                                }).then(function(json) {
                                    //console.log('parsed json', json);
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
            authorclick:function (authorid) {
                //console.log("authorid"+authorid);
                fetch('/book/authorbookscount?authorid='+authorid).then(function(response) {
                    return response.json()
                }).then(function(json) {
                    //console.log('parsed json', json);
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
                                // template method of yourself
                                //console.log(data);
                                //console.log(pagination);
                                fetch('/book/authorbooks?start='+_.min(data)+'&limit='+data.length+'&authorid='+authorid).then(function(response) {
                                    return response.json()
                                }).then(function(json) {
                                    //console.log('parsed json', json);
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
            languageclick:function (lang_code) {
                //console.log("lang_code"+lang_code);
                fetch('/book/languagebookscount?lang_code='+lang_code).then(function(response) {
                    return response.json()
                }).then(function(json) {
                    //console.log('parsed json', json);
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
                                // template method of yourself
                                //console.log(data);
                                //console.log(pagination);
                                fetch('/book/languagebooks?start='+_.min(data)+'&limit='+data.length+'&lang_code='+lang_code).then(function(response) {
                                    return response.json()
                                }).then(function(json) {
                                    //console.log('parsed json', json);
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
            console.log("created");
            //hotbooks展示分页
            fetch('/book/bookscount').then(function(response) {
                return response.json()
            }).then(function(json) {
                //console.log('parsed json', json);
                if (json.statusCode ==200){
                    //app.hotbooks = json.info
                    //return done(tmp);
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
                            // template method of yourself
                            //console.log(data);
                            //console.log(pagination);
                            fetch('/book/ratingbooks?start='+_.min(data)+'&limit='+data.length).then(function(response) {
                                return response.json()
                            }).then(function(json) {
                                //console.log('parsed json', json);
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
            fetch('/book/bookscount').then(function(response) {
                return response.json()
            }).then(function(json) {
                //console.log('parsed json', json);
                if (json.statusCode ==200){
                    //app.hotbooks = json.info
                    //return done(tmp);
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
                            // template method of yourself
                            //console.log(data);
                            //console.log(pagination);
                            fetch('/book/newbooks?start='+_.min(data)+'&limit='+data.length).then(function(response) {
                                return response.json()
                            }).then(function(json) {
                                //console.log('parsed json', json);
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
            fetch('/book/bookscount').then(function(response) {
                return response.json()
            }).then(function(json) {
                //console.log('parsed json', json);
                if (json.statusCode ==200){
                    //app.hotbooks = json.info
                    //return done(tmp);
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
                            // template method of yourself
                            //console.log(data);
                            //console.log(pagination);
                            fetch('/book/discoverbooks?start='+_.min(data)+'&limit='+data.length).then(function(response) {
                                return response.json()
                            }).then(function(json) {
                                //console.log('parsed json', json);
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
            fetch('/tag/tagscount').then(function(response) {
                return response.json()
            }).then(function(json) {
                //console.log('parsed json', json);
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
                        pageSize: 8,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            // template method of yourself
                            //console.log(data);
                            //console.log(pagination);
                            fetch('/tag/tags?start='+_.min(data)+'&limit='+data.length).then(function(response) {
                                return response.json()
                            }).then(function(json) {
                                //console.log('parsed json', json);
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
            fetch('/author/authorscount').then(function(response) {
                return response.json()
            }).then(function(json) {
                //console.log('parsed json', json);
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
                        pageSize: 8,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            // template method of yourself
                            //console.log(data);
                            //console.log(pagination);
                            fetch('/author/authors?start='+_.min(data)+'&limit='+data.length).then(function(response) {
                                return response.json()
                            }).then(function(json) {
                                //console.log('parsed json', json);
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
            fetch('/language/languagescount').then(function(response) {
                return response.json()
            }).then(function(json) {
                //console.log('parsed json', json);
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
                            // template method of yourself
                            //console.log(data);
                            //console.log(pagination);
                            fetch('/language/languages?start='+_.min(data)+'&limit='+data.length).then(function(response) {
                                return response.json()
                            }).then(function(json) {
                                //console.log('parsed json', json);
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
            console.log("beforeMount");
            this.booksseen["hotbooks"] = true;
            this.booksseen["newbooks"] = false;
            this.booksseen["discover"] = false;
            this.booksseen["categories"] = false;
            this.booksseen["authors"] = false;
            this.booksseen["language"] = false;
        },
        mounted: function () {
            console.log("mounted");
            //this.booksseen["hotbooks"] = true;
        }
    });
}