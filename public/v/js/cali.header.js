function commonData() {
    const data = new FormData();
    if (data === undefined){
        alert("unsupport web browser html5");
    }

    if (store.get("session") !== undefined){
        data.append("session",store.get("session"));
    }else {
        data.append("session","watcher");
    }
    return data;
}

function UrlSearch(){
    let name, value;
    let str = location.href; //取得整个地址栏
    let num = str.indexOf("?");
    str=str.substr(num+1); //取得所有参数stringvar.substr(start [, length ]

    const arr = str.split("&"); //各个参数放到数组里
    for(let i=0; i < arr.length; i++){
        num=arr[i].indexOf("=");
        if(num>0){
            name=arr[i].substring(0,num);
            value=arr[i].substr(num+1);
            this[name]=value;
        }
    }
}

function tips(title,html) {
    $("#tipsModalLabel").text(title);
    $("#tipsModelBody").html(html);
    $("#tipsModal").modal({
        keyboard: false
    });
}

function markdown2html(m) {
    showdown.setOption('simpleLineBreaks', true);
    //showdown.setOption('\n', '<br/>');
    let converter = new showdown.Converter();
    return converter.makeHtml(m);
}

$(document).ready(function(){
    _.mixin(s.exports());

    //set the current location
    if (window.location.href.indexOf("login")<=0 && window.location.href.indexOf("signup")<=0){
        store.set("location",window.location.href);
    }else if (store.get("location") === undefined){
        store.set("location","/");
    }
});