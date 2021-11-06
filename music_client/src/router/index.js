import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeRouter from './home/homeRouter.js'
import HomeMRouter from './home/homeMRouter.js'

Vue.use(VueRouter)

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

const routes = [
  HomeRouter,
  HomeMRouter,
  {
    path: '/*',
    redirect: '/'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (!localStorage.getItem("token") && to.meta.isAuth === true) {
    var appType = localStorage.getItem("appType")
    if (appType == "m") {
      router.push('/m/login')
      Vue.prototype.$toast.fail("请重新登录")
      return
    } else {
      router.push('/login')
      Vue.prototype.$toast.fail("请重新登录")
      return
    }
  }
  next()
})

export default router