
function commonData() {
    data = new FormData();
    data.set("session",store.get("session"));
    return data;
}

$(document).ready(function(){
    // 定义名为 headerdiv 的新组件
    Vue.component('headerdiv', {
        // headerdiv 组件现在接受一个
        // 这个属性名为 book。
        props: ['book'],
        template: '\
        <div class="row">\
            <nav class="navbar navbar-default navbar-inverse navbar-fixed-top nav-background-color">\
                <div class="container-fluid">\
                <!-- Brand and toggle get grouped for better mobile display -->\
                    <div class="navbar-header">\
                    <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">\
                        <span class="sr-only">Toggle navigation</span>\
                        <span class="icon-bar"></span>\
                        <span class="icon-bar"></span>\
                        <span class="icon-bar"></span>\
                    </button>\
                    <a class="navbar-brand" href="/" target="_blank">Cali</a>\
                    </div>\
                    <!-- Collect the nav links, forms, and other content for toggling -->\
                    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">\
                        <ul class="nav navbar-nav">\
                            <li v-bind:class="{ active: isindex }"><a href="/" target="_blank"><span v-text="$t(\'lang.index\')"></span> <span class="sr-only">(current)</span></a></li>\
                            <li v-bind:class="{ active: islibrary }"><a href="/public" target="_blank"><span v-text="$t(\'lang.library\')"></span> <span class="sr-only">(current)</span></a></li>\
                            <li><a href="https://github.com/jiangmitiao/cali/blob/master/README.md" target="_blank"><span v-text="$t(\'lang.help\')"></span></a></li>\
                            <li><a href="http://blog.gavinzh.com" target="_blank"><span v-text="$t(\'lang.blog\')"></span></a></li>\
                            <li v-if="leftdropdownseen" class="dropdown">\
                                <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span class="caret"></span></a>\
                                    <ul class="dropdown-menu">\
                                        <li><a href="#">Action</a></li>\
                                        <li><a href="#">Another action</a></li>\
                                        <li><a href="#">Something else here</a></li>\
                                        <li role="separator" class="divider"></li>\
                                        <li><a href="#">Separated link</a></li>\
                                        <li role="separator" class="divider"></li>\
                                        <li><a href="#">One more separated link</a></li>\
                                    </ul>\
                            </li>\
                        </ul>\
                        <form class="navbar-form navbar-left" method="get" action="/search">\
                            <div class="form-group">\
                                <input name="q" type="text" class="form-control" :placeholder="$t(\'lang.searchholder\')">\
                            </div>\
                            <button type="submit" class="btn btn-default"><span v-text="$t(\'lang.search\')"></span></button>\
                        </form>\
                        <ul class="nav navbar-nav navbar-right">\
                            <li><a href="https://github.com/jiangmitiao/cali"  target="_blank">Github</a></li>\
                            <li v-if="rightdropdownseen" class="dropdown">\
                                <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false"><span v-if="islogin" v-text="user.userName"></span><span v-if="!islogin" v-text="$t(\'lang.login\') + $t(\'lang.signup\')"></span><span class="caret"></span></a>\
                                <ul class="dropdown-menu">\
                                    <li v-if="islogin"><a href="/person" target="_blank"><span v-text="$t(\'lang.personcenter\')" target="_blank"></span></a></li>\
                                    <li role="separator" class="divider"></li>\
                                    <li v-if="!islogin"><a href="/login" target="_blank"><span v-text="$t(\'lang.login\')"></span></a></li>\
                                    <li v-if="!islogin"><a href="/signup" target="_blank"><span v-text="$t(\'lang.signup\')"></span></a></li>\
                                    <li v-if="islogin"><a href="#" @click="logout"><span v-text="$t(\'lang.logout\')"></span></a></li>\
                                </ul>\
                            </li>\
                        </ul>\
                    </div><!-- /.navbar-collapse -->\
                </div><!-- /.container-fluid -->\
            </nav>\
    </div>\
        ',
        methods:{
            logout:function () {
                fetch('/api/user/logout',{method:'post',body:commonData()}).then(function(response) {
                    if (response.redirected){
                        window.location.href = response.url;
                    }
                    return response.json()
                }).then(function(json) {
                    store.remove('user');
                    store.remove('session');
                    window.location.reload(true);
                }).
                catch(function(ex) {
                    console.log('parsing failed', ex);
                    alert("server error");
                });
                return false;
            }
        },
        data:function () {
            return {
                rightdropdownseen: function () {
                    return true;
                }(),
                leftdropdownseen: function () {
                    return false
                }(),
                isindex:function () {
                    if (window.location.pathname == "/"){
                        return true;
                    } else {
                        return false;
                    }
                }(),
                islibrary:function () {
                    if (window.location.pathname.indexOf("public") >=0){
                        return true;
                    } else {
                        return false;
                    }
                }(),
                islogin:function () {
                    if (_.isUndefined(store.get("session")) || _.isUndefined(store.get("user"))){
                        return false;
                    }else {
                        fetch('/api/user/islogin',{method:'post',body:commonData()}).then(function(response) {
                            if (response.redirected){
                                window.location.href = response.url;
                            }
                            return response.json()
                        }).then(function(json) {
                            if (json.statusCode != 200){
                                store.remove('user');
                                store.remove('session');
                                window.location.reload(true);
                            }
                        }).
                        catch(function(ex) {
                            console.log('parsing failed', ex);
                            alert("server error");
                        });
                        return true;
                    }
                }(),
                user:function () {
                    if (_.isUndefined(store.get("user"))){
                        return {};
                    }else {
                        return store.get("user");
                    }
                }()

            };
        }
    });
});