export default {
    path:'/',
    name:'home',
    component:() => import('@/views/home/home.vue'),
    children:[
        {
            path:'/home/commend',
            name:'homeCommend',
            component:() => import('@/views/home/recommend.vue')
        },
        {
            path:'/home/singer',
            name:'singer',
            component:() => import('@/views/singer/singer.vue')
        },
        {
            path:'/',
            redirect:'/home/commend'
        },
    ]
}