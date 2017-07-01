$(document).ready(function(){
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            listseen :"discover",
            //首页展示readme
            discover:"",

            session:"",
            user:{},

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

            //category change
            categoryseen:false,
            categories:[],
            newcategory:{},

            //download
            downloadstats:{},

            //upload
            uploadfileconfirmseen:false,
            formatid:"",
            uploadbook:{}
        },
        methods: {
            changeseen:function (e) {
                this.listseen = e;
            },
            needseen:function (e) {
                return this.listseen===e;
            },
            markdown2html: function (m) {
                return markdown2html(m);
            },
            //上传文件
            uploadfile:function () {
                let file = document.getElementById("book").value;
                if (file === null || (file.match(/.epub$/g)===null && file.match(/.mobi$/g)===null)){
                    tips("info","please open file or choose .epub/.mobi");
                    return;
                }
                let form = new FormData(document.getElementById("uploadfile"));
                form.append('session',store.get("session"));
                fetch("/api/book/uploadbook", {method: "post",body: form}).then(function(response) {
                    if (response.redirected){
                        let json = {};
                        json.statusCode = 500;
                        json.message = "role action setting error !";
                        return json;
                    }else {
                        return response.json();
                    }
                }).then(function(json) {
                    if (json.statusCode ===200){
                        app.formatid = json.info.id;
                        app.uploadbook.title = _(json.info.title).chain().trim().value();
                        app.uploadbook.author = _(json.info.author).chain().trim().value();
                        app.uploadbook.douban_id = "";
                        app.uploadfileconfirmseen = true;
                    }else {
                        tips("error",json.message);
                    }
                }).catch(function(ex) {
                    tips("error",ex);
                });
            },
            //上传文件确认
            uploadfileconfirm :function () {
                let form = new FormData(document.getElementById("uploadfileconfirm"));
                form.append('session',store.get("session"));
                fetch("/api/book/uploadbookconfirm", {method: "post",body: form}).then(function(response) {
                    if (response.redirected){
                        let json = {};
                        json.statusCode = 500;
                        json.message = "role action setting error !";
                        return json;
                    }else {
                        return response.json();
                    }
                }).then(function(json) {
                    if (json.statusCode ===200){
                        tips("info","upload success!");
                        document.getElementById("book").value = "";
                        app.uploadfileconfirmseen = false;
                    }else {
                        tips("error",json.message);
                    }
                }).catch(function(ex) {
                    tips("error",ex);
                });
            },
            //搜索用户
            searchloginnameclick:function () {
                let form = commonData();
                form.append("loginName",app.searchloginname);
                fetch('/api/user/queryusercount',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        app.userlistseen = false;
                        let tmpJson = {};
                        tmpJson.statusCode = 500;
                        return tmpJson;
                    }else{
                        return response.json();
                    }
                }).then(function(json) {
                    if (json.statusCode ===200){
                        $('#userlistpage').pagination({
                            dataSource:function (done) {
                                let tmp = [];
                                for(let i=0; i<json.info; i++){
                                    tmp.push(i)
                                }
                                return done(tmp);
                            },
                            pageRange:1,
                            totalNumber:json.info,
                            pageSize: 10,
                            showGoInput: true,
                            showGoButton: true,
                            callback: function(data, pagination) {
                                let form = commonData();
                                form.append("start",_.min(data));
                                form.append("limit",data.length);
                                form.append("loginName",app.searchloginname);
                                fetch('/api/user/queryuser',{method:'post',body:form}).then(function(response) {
                                    if (response.redirected){
                                        app.userlistseen = false;
                                        let tmpJson = {};
                                        tmpJson.statusCode = 500;
                                        return tmpJson;
                                    }else {
                                        return response.json();
                                    }
                                }).then(function(json) {
                                    if (json.statusCode ===200){
                                        app.userlistseen = true;
                                        app.userlist = json.info
                                    }else {
                                        tips("error","query count ok, and query user list error");
                                    }
                                }).catch(function(ex) {
                                    tips("error",ex);
                                });
                            }
                        });
                    }
                }).catch(function(ex) {
                    tips("error",ex);
                });
            },
            //修改用户昵称
            changeuserinfo:function () {
                if (app.user.userName === null || app.user.userName==="" ){
                    tips("warn","can not null");
                }else if (app.user.userName.length >64){
                    tips("warn","can not more than 64 chars");
                }else {
                    let form = commonData();
                    form.append("userName",app.user.userName);
                    fetch('/api/user/update',{method:'post',body:form}).then(function(response) {
                        if (response.redirected){
                            let tmpJson = {};
                            tmpJson.statusCode = 500;
                            tmpJson.message = "role action setting error !";
                            return tmpJson;
                        }else {
                            return response.json();
                        }
                    }).then(function(json) {
                        if (json.statusCode ===200){
                            fetch('/api/user/info',{method:'post',body:commonData()}).then(function(response) {
                                return response.json();
                            }).then(function(user) {
                                if (user.statusCode ===200){
                                    store.set("user", user.info);
                                    tips("info","<p>upload success</p><p>after 3 seconds, reload</p>");
                                    setTimeout("window.location.href = '/person'",3000);
                                }else {
                                    tips("error","密码错误:"+user.message);
                                }
                            }).catch(function(ex) {
                                tips("error",ex);
                            });
                        }else {
                            tips("error",json.message);
                        }
                    }).catch(function(ex) {
                        tips("error",ex);
                    });
                }
            },
            checkchangepassword:function () {
                if (app.oldPassword==="" || app.loginPassword===""){
                    app.changepasswordtipinfo = "lang.notnull";
                    return ;
                }
                if(app.loginPassword!==app.repeatLoginPassword){
                    app.changepasswordtipinfo = "lang.pcp";
                    return ;
                }
                app.changepasswordtipinfo = "";
            },
            changepassword:function () {
                if (app.changepasswordtipinfo !== ""){
                    tips("info",app.changepasswordtipinfo);
                    return;
                }
                let form = commonData();
                form.append("oldLoginPassword",app.oldLoginPassword);
                form.append("loginPassword",app.loginPassword);
                fetch('/api/user/changepassword',{method:'post',body:form}).then(function(response) {
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ===200){
                        store.remove("user");
                        store.remove("session");
                        tips("info","<p>update success</p><p>after 3 seconds, turn to login</p>");
                        setTimeout("window.location.href = '/login'",3000);
                    }else {
                        tips("error","密码错误:"+json.message);
                    }
                }).catch(function(ex) {
                    tips("error",ex);
                });

            },
            sysconfigupdate : function (id) {
                let find = false;
                for (let i=0; i< app.sysconfiglist.length; i++){
                    if (app.sysconfiglist[i].id === id){
                        find = true;
                        let form = commonData();
                        form.append("id",id);
                        form.append("key",app.sysconfiglist[i].key);
                        form.append("value",app.sysconfiglist[i].value);
                        fetch('/api/sysconfig/update',{method:'post',body:form}).then(function(response) {
                            return response.json();
                        }).then(function(json) {
                            if (json.statusCode ===200){
                                tips("info","update success");
                            }else {
                                tips("error","update error");
                            }
                        }).catch(function(ex) {
                            tips("error",ex);
                        });
                    }
                }
                if (!find){
                    tips("error","update error");
                }
            },
            //删除系统状态
            sysstatusdelete : function (id) {
                let find = false;
                for (let i=0; i< app.sysstatuslist.length; i++){
                    if (app.sysstatuslist[i].id === id){
                        let tmp = app.sysstatuslist[i];
                        find = true;
                        let form = commonData();
                        form.append("id",id);
                        form.append("key",tmp.key);
                        form.append("value",tmp.value);
                        fetch('/api/sysstatus/delete',{method:'post',body:form}).then(function(response) {
                            return response.json();
                        }).then(function(json) {
                            if (json.statusCode ===200){
                                app.sysstatuslist = _.difference(app.sysstatuslist,[tmp]);
                                tips("info","delete success");
                            }else {
                                tips("error","delete error");
                            }
                        }).catch(function(ex) {
                            tips("error",ex);
                        });
                    }
                }
                if (!find){
                    tips("error","delete error");
                }
            },
            addcategory:function () {
                let form = commonData();
                form.append("categoryName",this.newcategory.category);
                fetch('/api/category/add',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        let tmpJson = {};
                        tmpJson.statusCode = 500;
                        tmpJson.message = "role action setting error !";
                        return tmpJson;
                    }else {
                        return response.json();
                    }
                }).then(function(json) {
                    if (json.statusCode ===200){
                        app.showcategories();
                        tips("info","add categoy success");
                    }else {
                        tips("error",json.message);
                    }
                }).catch(function(ex) {
                    tips("error", ex);
                });
            },
            deletecategory:function (c) {
                let form = commonData();
                form.append("categoryId",c.id);
                form.append("categoryName",c.category);
                fetch('/api/category/delete',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        let tmpJson = {};
                        tmpJson.statusCode = 500;
                        tmpJson.message = "role action setting error !";
                        return tmpJson;
                    }else {
                        return response.json();
                    }
                }).then(function(json) {
                    if (json.statusCode ===200){
                        app.showcategories();
                        tips("info","delete categoy success");
                    }else {
                        tips("error",json.message);
                    }
                }).catch(function(ex) {
                    tips("error", ex);
                });
            },
            updatecategory:function (c) {
                let form = commonData();
                form.append("categoryId",c.id);
                form.append("categoryName",c.category);
                fetch('/api/category/update',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        let tmpJson = {};
                        tmpJson.statusCode = 500;
                        tmpJson.message = "role action setting error !";
                        return tmpJson;
                    }else {
                        return response.json();
                    }
                }).then(function(json) {
                    if (json.statusCode ===200){
                        tips("info","update categoy success");
                        app.showcategories();
                    }else {
                        tips("error",json.message);
                    }
                }).catch(function(ex) {
                    tips("error", ex);
                });
            },
            showcategories:function () {
                fetch('/api/category/all',{method:'post',body:commonData()}).then(function(response) {
                    if (response.redirected){
                        let tmpJson = {};
                        tmpJson.statusCode = 500;
                        tmpJson.message = "role action setting error !";
                        return tmpJson;
                    }else {
                        return response.json();
                    }
                }).then(function(json) {
                    if (json.statusCode ===200){
                        app.categories = json.info;
                    }
                }).catch(function(ex) {
                    tips("error", ex);
                });
            }
        },
        computed: {
            computed_discover:function () {
                return markdown2html(this.discover);
            },
            computed_userstatus : function () {
                if (_.has(this.downloadstats,"count") && _.has(this.downloadstats,"maxcount")){
                    return "已下载 "+this.downloadstats.count+" 本.可下载 "+(this.downloadstats.maxcount-this.downloadstats.count)+"本";
                }
                return "暂无信息";
            }
        },
        created: function() {
            if (_.isUndefined(store.get("session")) || _.isUndefined(store.get("user"))){
                window.location = "/login";
            }
            this.session = store.get('session');
            this.user = store.get('user');

            //https://raw.githubusercontent.com/jiangmitiao/cali/master/README.md
            let url = "";
            if (get_language()==="zh-CN"){
                url = "https://raw.githubusercontent.com/jiangmitiao/cali/master/README_CN.md";
            }else {
                url = "https://raw.githubusercontent.com/jiangmitiao/cali/master/README.md";
            }
            fetch(url).then(function(response) {
                return response.text();
            }).then(function(text) {
                app.discover = text;
            }).catch(function(ex) {
                tips("error",'parsing failed');
            });


            fetch('/api/user/queryusercount',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    app.userlistseen = false;
                    let tmpJson = {};
                    tmpJson.statusCode = 500;
                    return tmpJson;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ===200){
                    $('#userlistpage').pagination({
                        dataSource:function (done) {
                            let tmp = [];
                            for(var i=0; i<json.info;i++){
                                tmp.push(i)
                            }
                            return done(tmp);
                        },
                        pageRange:1,
                        totalNumber:json.info,
                        pageSize: 10,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            let form = commonData();
                            form.append("start",_.min(data));
                            form.append("limit",data.length);
                            fetch('/api/user/queryuser',{method:'post',body:form}).then(function(response) {
                                if (response.redirected){
                                    app.userlistseen = false;
                                    let tmpJson = {};
                                    tmpJson.statusCode = 500;
                                    return tmpJson;
                                }
                                return response.json();
                            }).then(function(json) {
                                if (json.statusCode ===200){
                                    app.userlistseen = true;
                                    app.userlist = json.info
                                }
                            }).catch(function(ex) {
                                tips("error", ex);
                            });
                        }
                    });
                }
            }).catch(function(ex) {
                tips("error", ex);
            });

            fetch('/api/sysconfig/configs',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    app.sysconfigseen = false;
                    let tmpJson = {};
                    tmpJson.statusCode = 500;
                    return tmpJson;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ===200){
                    app.sysconfigseen = true;
                    app.sysconfiglist = json.info;
                }
            }).catch(function(ex) {
                tips("error", ex);
            });


            fetch('/api/sysstatus/status',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    app.sysstatusseen = false;
                    let tmpJson = {};
                    tmpJson.statusCode = 500;
                    return tmpJson;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ===200){
                    app.sysstatusseen = true;
                    app.sysstatuslist = json.info;
                }
            }).catch(function(ex) {
                tips("error", ex);
            });

            fetch('/api/user/userstatus',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    let tmpJson = {};
                    tmpJson.statusCode = 500;
                    return tmpJson;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ===200){
                    app.downloadstats = json.info;
                }
            }).catch(function(ex) {
                tips("error", ex);
            });

            //是否显示类目
            fetch('/api/sysstatus/status',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    app.categoryseen = false;
                    let tmpJson = {};
                    tmpJson.statusCode = 500;
                    app.showcategories();
                    return tmpJson;
                }
                return response.json();
            }).then(function(json) {
                if (json.statusCode ===200){
                    app.categoryseen = true;
                    app.showcategories();
                }
            }).catch(function(ex) {
                tips("error", ex);
            });

        }
    });
});