$(document).ready(function(){
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            loginName:"",
            loginPassword:""
        },
        methods: {
            login:function () {
                loadingStart();
                let form = commonData();
                form.append("loginName",this.loginName);
                form.append("loginPassword",this.loginPassword);
                fetch('/api/user/login',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        tips("info","after 3 seconds, turn to "+response.url);
                        setTimeout("window.location.href = response.url",3000);
                        loadingStop();
                    }else {
                        return response.json();
                    }
                }).then(function(json) {
                    if (json.statusCode ===200){
                        store.set("session", json.info);
                        let form = commonData();
                        form.append("session",json.info);
                        fetch('/api/user/info',{method:'post',body:form}).then(function(response) {
                            return response.json()
                        }).then(function(user) {
                            if (user.statusCode ===200){
                                store.set("user", user.info);
                                store.set("session", json.info);
                                if (store.get("location") !== undefined && store.get("location") !== ""){
                                    tips("info","after 3 seconds, turn to "+store.get("location"));
                                    setTimeout("window.location.href = store.get('location')",3000);
                                }else {
                                    tips("info","after 3 seconds, turn to index");
                                    setTimeout("window.location.href = '/'",3000);
                                }
                            }else {
                                tips("warn",user.message);
                            }
                        }).catch(function(ex) {
                            tips("error",ex);
                        });
                    }else {
                        tips("error",json.message);
                    }
                    loadingStop();
                }).catch(function(ex) {
                    tips("error",ex);
                    loadingStop();
                });
            }
        },
        created: function() {
            if (!_.isUndefined(store.get("session")) && !_.isUndefined(store.get("user"))){
                tips("info","<p>you has login.</p><p>after 3 seconds, turn to index</p>");
                setTimeout("window.location.href = '/'",3000);
            }
        }
    });
});