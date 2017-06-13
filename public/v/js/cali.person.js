$(document).ready(function(){
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            session:"",
            user:{},
            discover:"",
            listseen:{}
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