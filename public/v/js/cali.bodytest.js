require.config({
    paths: {
        //https://github.com/jquery/jquery
        "jquery": "https://cdn.bootcss.com/jquery/3.1.1/jquery.min",
        //https://github.com/twbs/bootstrap
        "bootstrap": "https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min",
        //https://github.com/superRaytin/paginationjs
        "pagination":"https://cdn.bootcss.com/paginationjs/2.0.8/pagination.min",
        //https://github.com/jashkenas/underscore
        "underscore": "https://cdn.bootcss.com/underscore.js/1.8.3/underscore-min",
        //https://github.com/showdownjs/showdown
        "showdown":"https://cdn.bootcss.com/showdown/1.7.1/showdown.min",
        //https://github.com/marcuswestin/store.js
        "store":"https://cdn.bootcss.com/store.js/1.3.20/store.min",
        //https://github.com/github/fetch
        "fetch":"https://cdn.bootcss.com/fetch/2.0.3/fetch.min",
        //https://github.com/vuejs/vue
        "vue":"https://cdn.bootcss.com/vue/2.3.4/vue.min",
        //https://github.com/kazupon/vue-i18n
        "vue-i18n":"https://cdn.bootcss.com/vue-i18n/7.0.3/vue-i18n",
        //https://github.com/moment/moment
        "moment":"https://cdn.bootcss.com/moment.js/2.15.1/moment.min",
        //self
        // "header":"/public/v/js/cali.header",
        // "footer":"/public/v/js/cali.footer",
        // 'i18n':"/public/v/js/cali.i18n"
    },
    shim: {
        'jquery': {
            exports: 'jquery'
        },
        'bootstrap': {
            deps: ['jquery']
        },
        'pagination':{
            deps:['jquery']
        },
        'underscore':{
            exports: '_'
        },
        'fetch':{
            exports:'fetch'
        },
        'vue-i18n': {
            deps: ['vue']
        },
        'moment':{
            exports:'moment'
        },
        //self
        // 'header':{
        //     exports:'commonData',
        //     deps:["jquery","store","vue"]
        // },
        // 'footer':{
        //     deps:["jquery","store","vue"]
        // },
        // 'i18n':{
        //     //exports:'getI18n',
        //     deps:["jquery","store","vue","vue-i18n"]
        // }
    }
});
require(['jquery', 'bootstrap', 'underscore','showdown','pagination','store','vue','vue-i18n','fetch','moment'], function ($, bootstrap, _, showdown,pagination,store,Vue,VueI18n,fetch,moment){
    // some code here
    var app = new Vue({
        el: "#root",
        data: {},
        methods: {},
        computed: {},
        created: function() {
            console.log("created");
        },
        beforeMount: function () {
            console.log("beforeMount");
        },
        mounted: function () {
            console.log("mounted");
        }
    });
});