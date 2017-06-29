$(document).ready(function(){
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
                    <td><input type="text" maxlength="64" v-model="item.value" class="form-control"></td>\
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


    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            session:"",
            user:{},
            discover:"",
            listseen:{},
            changeuserinfotipinfo:"",

            //change password
            oldLoginPassword:"",
            loginPassword:"",
            repeatLoginPassword:"",
            changepasswordtipinfo:"",

            //userlist
            userlistseen:false,
            userlist:[],
            searchloginname:"",

            //sysconfig
            sysconfigseen:false,
            sysconfiglist:[],

            //sysstatus
            sysstatusseen:false,
            sysstatuslist:[],

            //download
            downloadstats:{},

            //upload
            uploadfileconfirmseen:false,
            formatid:"",
            uploadbook:{}
        },
        methods: {
            changeseen:function (e) {
                this.listseen = {};
                this.listseen["discover"] = false;
                this.listseen["userlist"] = false;
                this.listseen["sysconfig"] = false;
                this.listseen["sysstatus"] = false;
                this.listseen["download"] = false;
                this.listseen["upload"] = false;
                this.listseen["changeuserinfo"] = false;
                this.listseen["changepassword"] = false;
                this.listseen[e] = true;
            },
            markdown2html: function (m) {
                showdown.setOption('simpleLineBreaks', true);
                //showdown.setOption('\n', '<br/>');
                var converter = new showdown.Converter();
                var html      = converter.makeHtml(m);
                return html;
            },
            uploadfile:function () {
                var file = document.getElementById("book").value;
                if (file == null || (file.match(/.epub$/g)==null && file.match(/.mobi$/g)==null && file.match(/.pdf$/g)==null && file.match(/.txt$/g)==null)){
                    alert("please open file or choose .epub/.mobi/.pdf/.txt");
                    return;
                }
                var form = new FormData(document.getElementById("uploadfile"));
                form.append('session',store.get("session"));
                fetch("/api/book/uploadbook", {
                    method: "post",
                    body: form
                }).then(function(response) {
                    if (response.redirected){
                        alert("role action setting error !");
                        return;
                    }
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ==200){
                        alert(json.message);
                        app.formatid = json.info.id;
                        app.uploadbook.title = json.info.title;
                        app.uploadbook.author = json.info.author;
                        app.uploadbook.douban_id = "";
                        app.uploadfileconfirmseen = true;
                    }else {
                        alert("upload error "+json.info);
                    }
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex)
                });
                alert("please wait...");
            },
            uploadfileconfirm :function () {
                var form = new FormData(document.getElementById("uploadfileconfirm"));
                form.append('session',store.get("session"));
                fetch("/api/book/uploadbookconfirm", {
                    method: "post",
                    body: form
                }).then(function(response) {
                    if (response.redirected){
                        alert("role action setting error !");
                        return;
                    }
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ==200){
                        alert("success");
                        app.uploadfileconfirmseen = false;
                    }else {
                        alert("upload error "+json.info);
                    }
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex)
                });
                //alert("please wait...");
            },
            searchloginnameclick:function () {
                var form = commonData();
                form.append("loginName",app.searchloginname);
                fetch('/api/user/queryusercount',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        app.userlistseen = false;
                        var tmpJson = {};
                        tmpJson.statusCode = 500;
                        return tmpJson;
                    }
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ==200){
                        $('#userlistpage').pagination({
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
                                form.append("start",_.min(data));
                                form.append("limit",data.length);
                                form.append("loginName",app.searchloginname);
                                fetch('/api/user/queryuser',{method:'post',body:form}).then(function(response) {
                                    if (response.redirected){
                                        app.userlistseen = false;
                                        var tmpJson = {};
                                        tmpJson.statusCode = 500;
                                        return tmpJson;
                                    }
                                    return response.json();
                                }).then(function(json) {
                                    if (json.statusCode ==200){
                                        app.userlistseen = true;
                                        app.userlist = json.info
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
            changeuserinfo:function () {
                if (app.user.userName == null || app.user.userName==""){
                    app.changeuserinfotipinfo = "lang.notnull";
                    return;
                }else {
                    app.changeuserinfotipinfo = "";
                    var form = commonData();
                    form.append("userName",app.user.userName);
                    fetch('/api/user/update',{method:'post',body:form}).then(function(response) {
                        if (response.redirected){
                            alert("role action setting error !");
                            return;
                        }
                        return response.json();
                    }).then(function(json) {
                        if (json.statusCode ==200){
                            fetch('/api/user/info',{method:'post',body:commonData()}).then(function(response) {
                                return response.json();
                            }).then(function(user) {
                                if (user.statusCode ==200){
                                    store.set("user", user.info);
                                    alert("update success");
                                    window.location = "/person"
                                }else {
                                    alert("密码错误:"+user.message);
                                }
                            }).
                            catch(function(ex) {
                                console.log('parsing failed', ex)
                            });
                        }else {
                            alert("update error "+json.info);
                        }
                    }).
                    catch(function(ex) {
                        console.log('parsing failed', ex)
                    });
                }
            },
            checkchangepassword:function () {
                if (app.oldPassword=="" || app.loginPassword==""){
                    app.changepasswordtipinfo = "lang.notnull";
                    return ;
                }
                if(app.loginPassword!=app.repeatLoginPassword){
                    app.changepasswordtipinfo = "lang.pcp";
                    return ;
                }
                app.changepasswordtipinfo = "";
                return;
            },
            changepassword:function () {
                if (app.changepasswordtipinfo != ""){
                    return;
                }
                var form = commonData();
                form.append("oldLoginPassword",app.oldLoginPassword);
                form.append("loginPassword",app.loginPassword);
                fetch('/api/user/changepassword',{method:'post',body:form}).then(function(response) {
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ==200){
                        store.remove("user");
                        store.remove("session");
                        alert("update success");
                        window.location = "/login"
                    }else {
                        alert("密码错误:"+json.message);
                    }
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex)
                });

            },
            sysconfigupdate : function (id) {
                var find = false;
                for (var i=0;i< app.sysconfiglist.length; i++){
                    if (app.sysconfiglist[i].id == id){
                        find = true;
                        var form = commonData();
                        form.append("id",id);
                        form.append("key",app.sysconfiglist[i].key);
                        form.append("value",app.sysconfiglist[i].value);
                        fetch('/api/sysconfig/update',{method:'post',body:form}).then(function(response) {
                            return response.json();
                        }).then(function(json) {
                            if (json.statusCode ==200){
                                alert("update success");
                            }else {
                                alert("update error");
                            }
                        }).
                        catch(function(ex) {
                            console.log('parsing failed', ex)
                        });
                    }
                }
                if (!find){
                    alert("update error");
                }
            },
            sysstatusdelete : function (id) {
                var find = false;

                for (var i=0;i< app.sysstatuslist.length; i++){
                    if (app.sysstatuslist[i].id == id){
                        var tmp = app.sysstatuslist[i] ;
                        find = true;
                        var form = commonData();
                        form.append("id",id);
                        form.append("key",tmp.key);
                        form.append("value",tmp.value);
                        fetch('/api/sysstatus/delete',{method:'post',body:form}).then(function(response) {
                            return response.json();
                        }).then(function(json) {
                            if (json.statusCode ==200){
                                app.sysstatuslist = _.difference(app.sysstatuslist,[tmp]);
                                alert("delete success");
                            }else {
                                alert("delete error");
                            }
                        }).
                        catch(function(ex) {
                            console.log('parsing failed', ex)
                        });
                    }
                }
                if (!find){
                    alert("delete error");
                }
            }
        },
        computed: {
            computed_discover:function () {
                showdown.setOption('simpleLineBreaks', true);
                //showdown.setOption('\n', '<br/>');
                var converter = new showdown.Converter();
                var html      = converter.makeHtml(this.discover);
                return html;
            },
            userlist_computed : function () {
                return this.userlist;
            },
            computed_userstatus : function () {
                if (_.has(app.downloadstats,"count") && _.has(app.downloadstats,"maxcount")){
                    return "已下载 "+app.downloadstats.count+" 本.可下载 "+(app.downloadstats.maxcount-app.downloadstats.count)+"本";
                }
                return "暂无信息"
            }
        },
        created: function() {
            //console.log("created");
            if (_.isUndefined(store.get("session")) || _.isUndefined(store.get("user"))){
                window.location = "/login";
            }
            this.session = store.get('session');
            this.user = store.get('user');

            //https://raw.githubusercontent.com/jiangmitiao/cali/master/README.md
            var url = "";
            if (get_language()=="zh-CN"){
                url = "https://raw.githubusercontent.com/jiangmitiao/cali/master/README_CN.md";
            }else {
                url = "https://raw.githubusercontent.com/jiangmitiao/cali/master/README.md";
            }
            fetch(url).then(function(response) {
                return response.text();
            }).then(function(text) {
                app.discover = text;
            }).catch(function(ex) {
                console.log('parsing failed', ex)
            });


            fetch('/api/user/queryusercount',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    app.userlistseen = false;
                    var tmpJson = {};
                    tmpJson.statusCode = 500;
                    return tmpJson;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    $('#userlistpage').pagination({
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
                            form.append("start",_.min(data));
                            form.append("limit",data.length);
                            fetch('/api/user/queryuser',{method:'post',body:form}).then(function(response) {
                                if (response.redirected){
                                    app.userlistseen = false;
                                    var tmpJson = {};
                                    tmpJson.statusCode = 500;
                                    return tmpJson;
                                }
                                return response.json();
                            }).then(function(json) {
                                if (json.statusCode ==200){
                                    app.userlistseen = true;
                                    app.userlist = json.info
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

            fetch('/api/sysconfig/configs',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    app.sysconfigseen = false;
                    var tmpJson = {};
                    tmpJson.statusCode = 500;
                    return tmpJson;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    app.sysconfigseen = true;
                    app.sysconfiglist = json.info;
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });


            fetch('/api/sysstatus/status',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    app.sysstatusseen = false;
                    var tmpJson = {};
                    tmpJson.statusCode = 500;
                    return tmpJson;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    app.sysstatusseen = true;
                    app.sysstatuslist = json.info;
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });

            fetch('/api/user/userstatus',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    app.sysstatusseen = false;
                    var tmpJson = {};
                    tmpJson.statusCode = 500;
                    return tmpJson;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ==200){
                    app.downloadstats = json.info;
                }
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });
        },
        beforeMount: function () {
            //console.log("beforeMount");
            this.listseen = {};
            this.listseen["discover"] = true;
            this.listseen["sysconfig"] = false;
            this.listseen["sysstatus"] = false;
            this.listseen["userlist"] = false;
            this.listseen["upload"] = false;
            this.listseen["changeuserinfo"] = false;
            this.listseen["changepassword"] = false;
        },
        mounted: function () {
            //console.log("mounted");
        }
    });
});