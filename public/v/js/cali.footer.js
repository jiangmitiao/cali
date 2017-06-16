$(document).ready(function(){
    // 定义名为 footerdiv 的新组件
    Vue.component('footerdiv', {
        // headerdiv 组件现在接受一个
        // 这个属性名为 book。
        props: ['book'],
        template: '\
        <footer class="navbar-fixed-bottom">\
            <div class="container">\
                <div class="copy text-center">\
                    Copyright 2017 <a href="/">Cali</a>\
                </div>\
            </div>\
        </footer>\
        ',
        methods:{}
    });
});