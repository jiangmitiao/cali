$(document).ready(function(){
    commonData();
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            loginName:"",
            loginPassword:""
        },
        methods: {
            login:function () {
                var form = commonData();
                form.set("loginName",this.loginName);
                form.set("loginPassword",this.loginPassword);
                fetch('/api/user/login',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        window.location.href = response.url;
                    }
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ==200){
                        var form = commonData();
                        form.set("session",json.info);
                        fetch('/api/user/info',{method:'post',body:form}).then(function(response) {
                            return response.json()
                        }).then(function(user) {
                            if (user.statusCode ==200){
                                store.set("user", user.info);
                                store.set("session", json.info);
                                window.location = "/"
                            }else {
                                alert("密码错误:"+user.message);
                            }
                        }).
                        catch(function(ex) {
                            console.log('parsing failed', ex)
                        });
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

        },
        created: function() {
            //console.log("created");
            if (!_.isUndefined(store.get("session")) && !_.isUndefined(store.get("user"))){
                window.location = "/"
            }
        },
        beforeMount: function () {
            //console.log("beforeMount");
        },
        mounted: function () {
            //console.log("mounted");
        }
    });
});