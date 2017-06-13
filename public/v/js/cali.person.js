$(document).ready(function(){
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            session:""
        },
        methods: {

        },
        computed: {

        },
        created: function() {
            console.log("created");
            if (_.isUndefined(store.get("session")) || _.isUndefined(store.get("user"))){
                window.location = "/public/v/login.html"
                ///public/v/login.html
            }
            this.session = store.get('session');
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