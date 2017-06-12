function get_language() {
    if (navigator.language) {
        var language = navigator.language;
    }
    else {
        var language = navigator.browserLanguage;
    }
    return language;
}

// If using a module system (e.g. via vue-cli), import Vue and VueI18n and then call Vue.use(VueI18n).
// import Vue from 'vue'
// import VueI18n from 'vue-i18n'
//
    Vue.use(VueI18n);

// Ready translated locale messages
const messages = {
    en: {
        lang: {
            search:"Search",
            browse: 'BROWSE',
            hotbooks: "Hot Books",
            newbooks:"New Books",
            discover:"Discover",
            categories:"Categories",
            authors:"Authors",
            language:"Language",
            ifd:"infomation from douban.com",
            rating:"rating",
            bookname:"Book's Name",
            bookauthor:"Book's Author",
            bookpage:"Book's Pages",
            bookdblink:"Book's DoubanLink",
            bookpublisher:"Book's Publisher",
            bookpublishtime:"Book's Publish Time",
            bookupdatetime:"Book's Update Time",
            bookmodifiedtime:"Book's Last Modified Time",
            bookisbn:"Book's ISBN",
            bookrating:"Book's Rating",
            bookprice:"Book's Price",
            bookauthorinfo:"Book's AuthorIntro",
            booksummary:"Book's Summary",
            bookcatalog:"Book's Catalog",
            bookdownloadlink:"Book's DownloadAddress",
            download:"download",
            loginName:"login name",
            loginPassword:"login password",
            login:"Login",
            logout:"Logout",
            signin:"SIGN IN",
            dhaay:"Don't have an account yet?",
            signup:"Sign Up",
            help:"Help",
            searchholder:"Search...",
            index:"Index",
            personcenter:"PersonCenter"
        }
    },
    cn: {
        lang: {
            search:"搜索",
            browse: '浏览',
            hotbooks: "热门书籍",
            newbooks:"新书推荐",
            discover:"探索发现",
            categories:"标签",
            authors:"作者",
            language:"语言",
            rating:"评分",
            ifd:"以下信息来自 douban.com",
            bookname:"书名",
            bookauthor:"作者",
            bookpage:"页码",
            bookdblink:"豆瓣链接",
            bookpublisher:"出版者",
            bookpublishtime:"出版时间",
            bookupdatetime:"更新时间",
            bookmodifiedtime:"修改时间",
            bookisbn:"ISBN",
            bookrating:"评分",
            bookprice:"价格",
            bookauthorinfo:"作者介绍",
            booksummary:"本书概要",
            bookcatalog:"目录",
            bookdownloadlink:"下载链接",
            download:"点击下载",
            loginName:"登录名",
            loginPassword:"登录密码",
            login:"登录",
            logout:"退出",
            signin:"登录",
            dhaay:"没有帐号?",
            signup:"注册",
            help:"帮助",
            searchholder:"查找...",
            index:"首页",
            personcenter:"个人中心"
        }
    }
};

// Create VueI18n instance with options
if (get_language()=="zh-CN"){
    loc = "cn"
}else {
    loc = "en"
}
const i18n = new VueI18n({
    locale: loc, // set locale
    messages, // set locale messages
});


// new Vue({
//     el:"#app"
// });


// Create a Vue instanc with `i18n` option
//new Vue({i18n}).$mount('#app');

// Now the app has started!