window.onload = function() {
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            loginName:"",
            loginPassword:""
        },
        methods: {
            login:function () {
                console.log(this.loginName);
                console.log(this.loginPassword);
                fetch('/api/user/login?loginName='+this.loginName+'&loginPassword='+this.loginPassword).then(function(response) {
                    return response.json()
                }).then(function(json) {
                    if (json.statusCode ==200){
                        console.log(json.info);
                        store.set("session", json.info);
                        //lwindow.location.href = "/public/v/public.html"
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
            console.log("created");
        },
        beforeMount: function () {
            console.log("beforeMount");
            //this.book.title="oookkk"
        },
        mounted: function () {
            console.log("mounted");
        }
    });
};