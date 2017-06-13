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
                console.log(this.loginName);
                console.log(this.loginPassword);
                if (this.tipinfo != ""){
                    alert("please enter correct info");
                    return;
                }
                fetch('/api/user/regist?loginName='+this.loginName+'&loginPassword='+this.loginPassword).then(function(response) {
                    return response.json()
                }).then(function(json) {
                    if (json.statusCode ==200){
                        console.log(json.info);
                        window.location = "/public/v/login.html"
                    }else {
                        alert("注册失败:"+json.message);
                    }
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex)
                });
            }
        },
        computed: {
            tipinfo:function () {
                if (this.loginName == ""){
                    return "lang.peln"
                }
                if (this.loginPassword == ""){
                    return "lang.pelp"
                }
                if (this.loginPassword != this.repeatLoginPassword){
                    return "lang.pcp"
                }
                return ""
            }
        },
        created: function() {
            console.log("created");
            //store.remove('user');
            //store.remove('session')
        },
        beforeMount: function () {
            console.log("beforeMount");
            //this.book.title="oookkk"
        },
        mounted: function () {
            console.log("mounted");
        }
    });
});