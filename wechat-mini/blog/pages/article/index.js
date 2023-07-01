const WxParse = require('../../wxParse/wxParse.js');
const api = require('../../utils/api.js');
const http = require('../../utils/http.js');



Page({

  /**
   * 页面的初始数据
   */
  data: {
    article:{}
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    this.getArticleByID(options.articleID)
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  },
  getArticleByID:function(articleID){
    var that = this;
    var query = {
      id:articleID,
      page:1,
      limit:100,
      fields:'id,title,html,feature_image,updated_at,excerpt',
      filter:'',
      include:'tags'
    }
    var req = http.getRequest(api.getArticleDetailUrl(query));
    req.then(res=>{
      var info = res.data.posts[0]
      WxParse.wxParse('info', 'html', info.html, that, 10);
      this.setData({
        article:info
      })
    })
  }
})