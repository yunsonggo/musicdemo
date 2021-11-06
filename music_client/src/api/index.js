import ajax from './ajax'
const BASE_URL = "http://192.168.1.102:8080/api"

// export const reqCities = () => ajax(BASE_URL + "/consumer/cities");
// export const reqEmailCode = (consumer_email) => ajax(BASE_URL + '/consumer/email/captcha',{consumer_email},'POST')
// 网站信息
export const reqWebsite = () => ajax(BASE_URL + "/website/info")
// 网站菜单
export const reqMenus = () => ajax(BASE_URL + "/menus")
// 轮播图数据
export const reqSwiper = () => ajax(BASE_URL + "/swiper")
// 获取歌单数据
export const reqTopCategory = (count,start) => ajax(BASE_URL + "/category",{count,start},'POST')
// 获取歌手数据
export const reqSinger = () => ajax(BASE_URL + "/singer")
// 根据歌手ID 获取 歌曲
export const reqMusicBySinger = (id,count,start) => ajax(BASE_URL + "/music/singer",{id,count,start})