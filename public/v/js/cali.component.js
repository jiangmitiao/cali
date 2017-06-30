$(document).ready(function(){

    // 定义名为 bookdiv 的新组件
    Vue.component('bookdiv', {
        // bookdiv 组件现在接受一个
        // 这个属性名为 book。
        props: ['book'],
        template: '\
        <div class="col-lg-2 col-md-3 col-sm-3 col-xs-6">\
            <div class="content-box">\
                <div class="panel-body text-center">\
                    <a :href="\'/book?bookid=\'+book.id" target="_blank" class="text-center">\
                        <canvas :id="book.id" :src="fakePic(toJson(book.douban_json).image,book.id)"></canvas>\
                    </a>\
                    <p class="text-center">\
                        <a :href="\'/book?bookid=\'+book.id" target="_blank">\
                            <nobr v-text="maxstring(book.title,5)" :title="book.title" style="word-break: keep-all;white-space: nowrap;"></nobr>\
                        </a>\
                    </p>\
                    <p class="text-center">\
                        <a :href="\'/search?q=\'+book.author" target="_blank">\
                        <nobr v-text="maxstring(book.author,5)"></nobr>\
                        </a>\
                    </p>\
                    <p class="text-center badge" style="background-color: #2c3742"><span v-text="$t(\'lang.rating\')"></span>:<span  v-text="toJson(book.douban_json).rating.average"></span></p>\
                    <br>\
                </div>\
            </div>\
        </div>\
        ',
        methods:{
            //return a sub string ,sub's length is max .if src string not equals result the result add '...'
            maxstring : function (str,max) {
                var result = str.substr(0,max);
                if (result.length !== str.length){
                    result+="...";
                }
                return result;
            },
            toJson : function (str) {
                return JSON.parse(str);
            },
            fakePic:function (src,id) {
                var img = new Image();
                img.src = src;
                img.onload = function(){
                    myCanvas = document.getElementById(id);
                    myCanvas.width = 100;
                    myCanvas.height = 150;
                    var context = myCanvas.getContext('2d');
                    context.drawImage(img, 0, 0);
                    //var imgdata = context.getImageData(0, 0, img.width, img.height);
                    // 处理imgdata
                };
                return "src"
            }
        }
    });

    // 定义名为 categorydiv 的新组件
    Vue.component('categorydiv', {
        // categorydiv 组件现在接受一个
        // 这个属性名为 category。
        props: ['category','categoryid'],
        template: '\
        <button type="button" @click="categoryclick(category)" :class="\'list-group-item \'+active(category,categoryid)"><i class="glyphicon glyphicon-star"></i><span v-text="category.category"></span></button>\
        ',
        methods:{
            //return a sub string ,sub's length is max .if src string not equals result the result add '...'
            categoryclick : function (c) {
                this.$emit('categoryclick',c);
            },
            active:function (category,categoryid) {
                if (category.id === categoryid){
                    return "active"
                }else {
                    return ""
                }
            }
        }
    });

});