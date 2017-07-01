$(document).ready(function(){
    //the instance is only one html's Vue's instance on public.html
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            categories :[],
            categoryid:"",
            categoryname:"",
            // the books control one div where is books which has 8 items
            books:[],
        },
        methods: {
            showbooks:function (c) {
                //books展示分页
                app.categoryid = c.id;
                app.categoryname = c.category;
                if (app.categoryid === "") {
                    return
                }
                app.books = [];
                let form = commonData();
                form.append("categoryId", app.categoryid);
                fetch('/api/book/bookscount', {method: 'post', body: form}).then(function (response) {
                    if (response.redirected) {
                        let tmpJson = {};
                        tmpJson.statusCode = 500;
                        return tmpJson;
                    }else {
                        return response.json();
                    }
                }).then(function (json) {
                    if (json.statusCode === 200) {
                        $('#bookspage').pagination({
                            dataSource: function (done) {
                                let tmp = [];
                                for (let i = 0; i < json.info; i++) {
                                    tmp.push(i);
                                }
                                return done(tmp);
                            },
                            pageRange: 1,
                            totalNumber: json.info,
                            pageSize: 12,
                            showGoInput: true,
                            showGoButton: true,
                            callback: function (data, pagination) {
                                let form = commonData();
                                form.append("start", _.min(data));
                                form.append("limit", data.length);
                                form.append("categoryId", app.categoryid);
                                fetch('/api/book/books', {method: 'post', body: form}).then(function (response) {
                                    if (response.redirected) {
                                        let tmpJson = {};
                                        tmpJson.statusCode = 500;
                                        return tmpJson;
                                    }else {
                                        return response.json();
                                    }
                                }).then(function (json) {
                                    if (json.statusCode === 200) {
                                        app.books = json.info;
                                    }else {
                                        tips("error","server error");
                                    }
                                }).catch(function (ex) {
                                    tips("error",ex);
                                });
                            }
                        });
                    }
                }).catch(function (ex) {
                    tips("error",ex);
                });
            }
        },
        created: function() {
            fetch('/api/category/all',{method:'post',body:commonData()}).then(function(response) {
                if (response.redirected){
                    let tmpJson = {};
                    tmpJson.statusCode = 500;
                    return tmpJson;
                }else {
                    return response.json();
                }
            }).then(function(json) {
                if (json.statusCode ===200){
                    app.categories = json.info;
                    if (app.categories.length !==0){
                        app.categoryid = app.categories[0].id;
                        app.categoryname = app.categories[0].category;
                        app.showbooks(app.categories[0]);
                    }
                }else {
                    tips("error","server error");
                }
            }).catch(function(ex) {
                tips("error",ex);
            });
        }
    });
});