import 'babel-polyfill'
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Vant from 'vant';
import 'vant/lib/index.css';
import { Lazyload } from 'vant';
import Fastclick from 'fastclick'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import Fragment from 'vue-fragment';
import './assets/css/music_icon/iconfont.css'


Fastclick.attach(document.body)
Vue.use(Vant);
Vue.use(ElementUI);
Vue.use(Fragment.Plugin);
Vue.use(Lazyload);
Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
