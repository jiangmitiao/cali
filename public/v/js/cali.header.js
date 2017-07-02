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

/**
 * @return {string}
 */
function DrawCanvasByImg(src,id,width,heigth) {
    let style;
    let app = new PIXI.Application(width, heigth, {backgroundColor: 0xc2c2c2});
    //document.getElementById(id).innerHTML = "";
    //document.getElementById(id).appendChild(app.view);

    if (src.length === 0){
        style = new PIXI.TextStyle({
            align: 'center',
            fontSize: width/5,
            breakWords: true,
            wordWrap: true,
            padding: 5,
        });
        let basicText = new PIXI.Text('\nCali',style);
        basicText.x = 0;
        basicText.y = 0;

        app.stage.addChild(basicText);
    }else {
        let texture = PIXI.Texture.fromImage(src);
        let sprite1 = new PIXI.Sprite(texture);
        sprite1.height = heigth;
        sprite1.width = width;
        sprite1.scale.x = texture.width/width;
        sprite1.scale.y = texture.height/heigth;
        app.stage.addChild(sprite1);

    }
    console.log(app.view.toDataURL("image/png"));
    return app.view.toDataURL("image/png");
}

/**
 * @return {string}
 */
function DrawCanvasByText(text,id,width,heigth) {
    let style;
    let app = new PIXI.Application(width, heigth, {backgroundColor: 0xc2c2c2});
    document.getElementById(id).innerHTML = "";
    document.getElementById(id).appendChild(app.view);


    style = new PIXI.TextStyle({
        align: 'center',
        fontSize: width/20,
        breakWords: true,
        wordWrap: true,
        padding: 5,
    });
    let basicText = new PIXI.Text('\n'+text,style);
    basicText.x = 0;
    basicText.y = 0;

    app.stage.addChild(basicText);
    console.log(app.view.toDataURL("image/png"));
    return app.view.toDataURL("image/png");
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