$(document).ready(function(){
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            loginName:"",
            loginPassword:"",
            repeatLoginPassword:""
        },
        methods: {
            regist:function () {
                if (this.tipinfo !== ""){
                    tips("warn","please enter correct info");
                    return;
                }
                let form = commonData();
                form.append("loginName",app.loginName);
                form.append("loginPassword",app.loginPassword);
                fetch('/api/user/regist',{method:'post',body:form}).then(function(response) {
                    if (response.redirected){
                        window.location.href = response.url;
                    }
                    return response.json();
                }).then(function(json) {
                    if (json.statusCode ===200){
                        tips("info","<p>"+json.message+"</p><p>after 5 seconds, turn to login</p>");
                        setTimeout("window.location.href = '/login'",5000);
                    }else {
                        tips("error",json.message);
                    }
                }).catch(function(ex) {
                    tips("error",ex);
                });
            }
        },
        computed: {
            tipinfo:function () {
                if (this.loginName === ""){
                    return "lang.peln"
                }
                if (this.loginPassword === ""){
                    return "lang.pelp"
                }
                if (this.loginPassword !== this.repeatLoginPassword){
                    return "lang.pcp"
                }
                return ""
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