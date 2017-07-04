$(document).ready(function(){
    var app = new Vue({
        i18n,
        el: "#root",
        data: {},
        methods: {
            test1:function () {
                appLoading.setColor('#f83');
                appLoading.start();
                setTimeout(function () {
                    appLoading.stop();
                },3000)
            }
        },
        omputed: {},
        created: function() {
            //console.log("created");
        },
        beforeMount: function () {
            //console.log("beforeMount");
        },
        mounted: function () {
            //console.log("mounted");
        }
    });
});