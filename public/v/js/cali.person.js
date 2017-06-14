$(document).ready(function(){
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
            changepasswordtipinfo:""
        },
        methods: {
            changeseen:function (e) {
                this.listseen = {};
                this.listseen["discover"] = false;
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
                if (file == null || (file.match(/.epub$/g)==null && file.match(/.mobi$/g)==null)){
                    alert("please open file or choose .epub/.mobi");
                    return;
                }
                alert("please wait...");
                var form = new FormData(document.getElementById("uploadfile"));
                fetch("/api/book/uploadbook", {
                    method: "POST",
                    body: form
                }).then(function(response) {
                    if (response.redirected){
                        alert("role action setting error !");
                        return;
                        //window.location.href = response.url;
                    }
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ==200){
                        alert(json.info);
                    }else {
                        alert("upload error "+json.info);
                    }
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex)
                });
            },
            changeuserinfo:function () {
                if (app.user.userName == null || app.user.userName=="" || app.user.email == null || app.user.email==""){
                    app.changeuserinfotipinfo = "lang.notnull";
                    return;
                }else {
                    app.changeuserinfotipinfo = "";
                    fetch("/api/user/update?session="+app.session+"&userName="+app.user.userName+"&email="+app.user.email).then(function(response) {
                        if (response.redirected){
                            alert("role action setting error !");
                            return;
                            //window.location.href = response.url;
                        }
                        return response.json();
                    }).then(function(json) {
                        if (json.statusCode ==200){
                            fetch('/api/user/info?session='+app.session).then(function(response) {
                                return response.json();
                            }).then(function(user) {
                                if (user.statusCode ==200){
                                    console.log(user.info);
                                    store.set("user", user.info);
                                    alert("update success");
                                    window.location = "/public/v/person.html"
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
                fetch('/api/user/changepassword?session='+app.session+"&oldLoginPassword="+app.oldLoginPassword+"&loginPassword="+app.loginPassword).then(function(response) {
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ==200){
                        console.log(json.info);
                        store.remove("user");
                        store.remove("session");
                        alert("update success");
                        window.location = "/public/v/login.html"
                    }else {
                        alert("密码错误:"+json.message);
                    }
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex)
                });

            }
        },
        computed: {
            computed_discover:function () {
                showdown.setOption('simpleLineBreaks', true);
                //showdown.setOption('\n', '<br/>');
                var converter = new showdown.Converter();
                var html      = converter.makeHtml(this.discover);
                return html;
            }
        },
        created: function() {
            console.log("created");
            if (_.isUndefined(store.get("session")) || _.isUndefined(store.get("user"))){
                window.location = "/public/v/login.html"
                ///public/v/login.html
            }
            this.session = store.get('session');
            this.user = store.get('user');
            //store.remove('session')

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
            }).
            catch(function(ex) {
                console.log('parsing failed', ex)
            });
        },
        beforeMount: function () {
            console.log("beforeMount");
            this.listseen = {};
            this.listseen["discover"] = true;
            this.listseen["upload"] = false;
            this.listseen["changeuserinfo"] = false;
            this.listseen["changepassword"] = false;
        },
        mounted: function () {
            console.log("mounted");
        }
    });
});