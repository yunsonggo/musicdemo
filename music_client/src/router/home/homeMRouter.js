export default {
    path:'/m',
    name:'Mhome',
    component:() => import('@/views/Mhome/Mhome.vue'),
    children:[
        {
            path:'/m/home/commend',
            name:'MhomeCommend',
            component:() => import('@/views/Mhome/Mrecommend.vue')
        },
        {
            path:'/m/home/singer',
            name:'Msinger',
            component:() => import('@/views/singer/mSinger.vue')
        },
        {
            path:'/m/home/music/list',
            name:'Mmusiclist',
            component:() => import('@/views/music/mMusicList.vue')
        },
        {
            path:'/m',
            redirect:'/m/home/commend'
        },
    ]
}