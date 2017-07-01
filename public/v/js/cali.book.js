$(document).ready(function(){
    var Request=new UrlSearch();
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            // the only one book's info
            book:{},
            //if bookseen is true ,then display the book's div
            bookseen:false
        },
        methods: {
            markdown2html: function (m) {
                showdown.setOption('simpleLineBreaks', true);
                //showdown.setOption('\n', '<br/>');
                const converter = new showdown.Converter();
                return converter.makeHtml(m);
            },
            formatdate:function (d) {
                return moment(new Date(d)).format("YYYY-MM-DD");
            }
        },
        created: function() {
            let form = commonData();
            form.append("bookid",Request.bookid);
            fetch('/api/book/book',{method:'post',body:form}).then(function(response) {
                if (response.redirected){
                    tips("info","after 3 seconds, turn to "+response.url);
                    setTimeout("window.location.href = response.url",3000);
                }else {
                    return response.json();
                }
            }).then(function(json) {
                if (json.statusCode ===200){
                    app.book = json.info;
                    app.bookseen = true;
                }else {
                    tips("error",json.message);
                }
            }).catch(function(ex) {
                tips("error",ex);
            });
        }
    });
});