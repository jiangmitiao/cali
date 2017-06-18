require.config({
    paths: {
        //https://github.com/jquery/jquery
        "jquery": "https://cdn.bootcss.com/jquery/3.1.1/jquery.min",
        //https://github.com/twbs/bootstrap
        "bootstrap": "https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min",
        //https://github.com/jashkenas/underscore
        "underscore": "https://cdn.bootcss.com/underscore.js/1.8.3/underscore-min",
        //https://github.com/showdownjs/showdown
        "showdown":"https://cdn.bootcss.com/showdown/1.7.1/showdown.min",
        //https://github.com/superRaytin/paginationjs
        "pagination":"https://cdn.bootcss.com/paginationjs/2.0.8/pagination.min",
        //https://github.com/marcuswestin/store.js
        "store":"https://cdn.bootcss.com/store.js/1.3.20/store.min",
        //https://github.com/github/fetch
        "fetch":"https://cdn.bootcss.com/fetch/2.0.3/fetch.min",
        //https://github.com/vuejs/vue
        "Vue":"https://cdn.bootcss.com/vue/2.3.4/vue.min"
    },
    shim: {
        'jquery': {
            exports: 'jquery'
        },
        'bootstrap': {
            deps: ['jquery']
        },
        'underscore':{
            exports: '_'
        },
        'backbone': {
            deps: ['underscore', 'jquery'],
            exports: 'Backbone'
        }
    }
});
require(['jquery', 'bootstrap', 'underscore','showdown','pagination','store'], function ($, bootstrap, _, showdown,pagination,store){
    // some code here
    var math = _.min([0,1,2,3,4,-1]);
    //alert(math);

    var converter = new showdown.Converter(),
        text      = '#hello, markdown!',
        html      = converter.makeHtml(text);
    //alert(html);

    // $('#testpage').pagination({
    //     dataSource:function (done) {
    //         var tmp = [];
    //         for(var i=0; i<100;i++){
    //             tmp.push(i);
    //         }
    //         return done(tmp);
    //     },
    //     pageRange:1,
    //     totalNumber:100,
    //     pageSize: 8,
    //     showGoInput: true,
    //     showGoButton: true,
    //     callback: function(data, pagination) {
    //         alert(data);
    //     }
    // });

    store.set("sss","sssok");
    //lalert(store.get("sss"))

    var data = new FormData();
    data.set("bookid",Request.bookid);
    fetch('/api/book/book',{method:'post',body:data}).then(function(response) {
        if (response.redirected){
            window.location.href = response.url;
        }
        return response.json();
    }).then(function(json) {
        alert(json.statusCode);
    }).
    catch(function(ex) {
        console.log('parsing failed', ex)
    });
});