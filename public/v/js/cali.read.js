$(document).ready(function(){
    function UrlSearch(){
        var name,value;
        var str=location.href; //取得整个地址栏
        var num=str.indexOf("?");
        str=str.substr(num+1); //取得所有参数stringvar.substr(start [, length ]

        var arr=str.split("&"); //各个参数放到数组里
        for(var i=0;i < arr.length;i++){
            num=arr[i].indexOf("=");
            if(num>0){
                name=arr[i].substring(0,num);
                value=arr[i].substr(num+1);
                this[name]=value;
            }
        }
    }
    var Request=new UrlSearch(); //实例化

    Book = ePub({
        restore: true ,
        spreads : false});

    var app = new Vue({
        i18n,
        el: "#root",
        data: {

        },
        methods: {
        },
        computed: {

        },
        created: function() {
            //console.log("created");

        },
        beforeMount: function () {
            //console.log("beforeMount");
        },
        mounted: function () {
            //console.log("mounted");
            //document.body.clientHeight

            $("#area").css("height",""+(document.body.clientHeight-200)+"px");
            $("#area").css("height",""+(600)+"px");
            $("#area").css("width",""+(500)+"px");
            var bookid=Request.bookid;
            var url = "/book/bookdown.epub?bookid="+bookid+"&session="+store.get("session");
            //console.log(url);
            Book.open(url);
            Book.getMetadata().then(function(meta){
                document.title = meta.bookTitle+" – "+meta.creator;
            });
            // Book.getToc().then(function(toc){
            //     var $select = document.getElementById("toc"),
            //         docfrag = document.createDocumentFragment();
            //     toc.forEach(function(chapter) {
            //         var option = document.createElement("option");
            //         option.textContent = chapter.label;
            //         option.ref = chapter.href;
            //         docfrag.appendChild(option);
            //     });
            //     $select.appendChild(docfrag);
            //     $select.onchange = function(){
            //         var index = $select.selectedIndex,
            //             url = $select.options[index].ref;
            //         Book.goto(url);
            //         return false;
            //     }
            // });
            Book.ready.all.then(function(){
                document.getElementById("loader").style.display = "none";
            });
            Book.renderTo("area");

            Mousetrap.bind('left', function() {
                //Book.prevPage();
            });
            Mousetrap.bind('right', function() {
                //Book.nextPage();
            });

            // $('#root').mousewheel(function(event) {
            //     console.log(event.deltaX, event.deltaY, event.deltaFactor);
            //     if($(document).scrollTop() + $(window).height() == $(document).height()){
            //         if(event.deltaY == -1){
            //             Book.nextPage();
            //             scroll(0,0);
            //
            //         }
            //     }
            //     if($(document).scrollTop() == 0 && event.deltaY == s1){
            //         Book.prevPage();
            //         scroll(0,0);
            //     }
            // });
        }
    });
});