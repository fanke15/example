const api = require('../../utils/api.js');
const http = require('../../utils/http.js');

Page({
  data: {
    active: 0,
    articleTags:[],
    bannerArticles:[],
    articles:[]
  },
  
   /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function () {
    // 获取tags
    this.getArticleTags()

    // 加载全部文章
    this.getArticles()

    // 加载Banner文章
    this.getBannerArticles()
  },

  onClickArticleDetail:function(event){
    wx.navigateTo({
      url: '../article/index?articleID='+event.currentTarget.dataset.info.id,
      success:function(){
        
      },
      fail:function(){

      },
      complete:function(){
      },
    })
  },
  onTagChange:function(event){
    console.log(event.detail)
    if (event.detail.index==0){
      this.getArticles()
    }else{
      this.getArticlesByTag(event.detail.name)
    }
  },
  getArticleTags:function(){
    var query = {
      page:1,
      limit:20,
      fields:'id,name,slug',
      filter:''
    }
    var tags = [{name:"全部"}]
    var req = http.getRequest(api.getTagListUrl(query));
    req.then(res=>{
      res.data.tags.forEach(function(value, key, iterable){
        tags.push(value)
      })
      this.setData({
        articleTags:tags
      })
    })
  },
  getArticles:function(){
    var query = {
      page:1,
      limit:100,
      fields:'id,title,feature_image,updated_at,excerpt',
      filter:'',
      include:'tags'
    }
    var req = http.getRequest(api.getArticleListUrl(query));
    req.then(res=>{
      this.setData({
        articles:res.data.posts
      })
    })
  },
  getBannerArticles:function(){
    var query = {
      page:1,
      limit:9,
      fields:'id,title,feature_image,updated_at,excerpt',
      filter:'featured:true'
    }
    var req = http.getRequest(api.getArticleListUrl(query));
    req.then(res=>{
      this.setData({
        bannerArticles:res.data.posts
      })
    })
  },
  getArticlesByTag:function(tagName){
    var query = {
      page:1,
      limit:100,
      fields:'id,title,feature_image,updated_at,excerpt',
      filter:'tag:'+tagName,
      include:'tags'
    }
    var req = http.getRequest(api.getArticleListUrl(query));
    req.then(res=>{
      this.setData({
        articles:res.data.posts
      })
    })
  }
})
