const apiURL = 'https://fanke.net.cn/ghost/api/content';
const key = '3af611d8878cdbf8d3316ff1f6';



/**
 * 获取文章列表url
 */
const getArticleListUrl = (params) => {
  var url= `${apiURL}/posts/?key=${key}&page=${params.page}&limit=${params.limit}&fields=${params.fields}&include=${params.include}&filter=${params.filter}`;
  return url
};

/**
 * 获取文章详情url
 */
const getArticleDetailUrl = (params) => {
  var url= `${apiURL}/posts/${params.id}/?key=${key}&page=${params.page}&limit=${params.limit}&fields=${params.fields}&include=${params.include}&filter=${params.filter}`;
  return url
};


/**
 * 获取标签列表url
 */
const getTagListUrl = (params) => {
  var url= `${apiURL}/tags/?key=${key}&page=${params.page}&limit=${params.limit}&fields=${params.fields}&filter=${params.filter}`;
  return url
};


module.exports = {
  getArticleListUrl,
  getTagListUrl,
  getArticleDetailUrl
};
