

Page({
  data: {
    post: {},
  },
   /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function () {
    console.log("demo")


wx.request({
  //这里的url用的是新视觉实训的一个测试接口
  url: 'https://fanke.net.cn/ghost/api/content/posts/?key=3af611d8878cdbf8d3316ff1f6',
  //success是接口调用成功的回调函数,这里习惯用res去接收返回值
  success:res=>{
    console.log(res)
  }
})
  },
})
