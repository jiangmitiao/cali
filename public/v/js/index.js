$(document).ready(function(){
    var app = new Vue({
        i18n,
        el: "#root",
        data: {},
        methods: {
            ooo:function () {
                tips("hello","<h4>world</h4>");
            }
        },
        computed: {},
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