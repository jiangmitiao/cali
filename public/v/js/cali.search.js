$(document).ready(function(){
    const Request = new UrlSearch(); //实例化

    //the instance is only one html's Vue's instance on search.html
    var app = new Vue({
        i18n,
        el: "#root",
        data: {
            searchbooks:[],
        },
        methods: {
        },
        computed: {
        },
        created: function() {
            //when page's instance created,we should get the data to render the page
            //console.log("created");
            //searchbooks展示分页
            loadingStart();
            let form = commonData();
            form.append("q",Request.q);
            fetch('/api/book/searchcount',{method:'post',body:form}).then(function(response) {
                if (response.redirected){
                    let tmpJson = {};
                    tmpJson.statusCode = 500;
                    return tmpJson;
                }else {
                    return response.json();
                }
            }).then(function(json) {
                if (json.statusCode ===200 && json.info !==0){
                    $('#searchbookspage').pagination({
                        dataSource:function (done) {
                            let tmp = [];
                            for(let i=0; i<json.info; i++){
                                tmp.push(i)
                            }
                            return done(tmp);
                        },
                        pageRange:1,
                        totalNumber:json.info,
                        pageSize: 12,
                        showGoInput: true,
                        showGoButton: true,
                        callback: function(data, pagination) {
                            let form = commonData();
                            form.append("start",_.min(data));
                            form.append("limit",data.length);
                            form.append("q",Request.q);
                            fetch('/api/book/search',{method:'post',body:form}).then(function(response) {
                                loadingStop();
                                if (response.redirected){
                                    let tmpJson = {};
                                    tmpJson.statusCode = 500;
                                    return tmpJson;
                                }else {
                                    return response.json();
                                }
                            }).then(function(json) {
                                if (json.statusCode ===200){
                                    app.searchbooks = json.info
                                }else {
                                    tips("error","server error");
                                }
                            }).
                            catch(function(ex) {
                                tips("error",ex);
                                loadingStop();
                            });
                        }
                    });
                }else if (json.statusCode ===500){
                    tips("error","server error");
                }else {
                    loadingStop();
                }
            }).catch(function(ex) {
                tips("error",ex);
                loadingStop();
            });
        }
    });
});