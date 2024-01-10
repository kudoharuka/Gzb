import LoginVue from '@/views/Login.vue'
import Vue from 'vue'
import VueRouter from 'vue-router'
import AdminHomeView from '../views/AdminHomeView.vue'
import UserView from "@/views/UserView";
import RecipeView from "@/views/RecipeView";
import CommentView from "@/views/CommentView";
import FeedBackView from "@/views/FeedBackView";
import DashBoard from "@/components/DashBoard";
import NewsView from "@/views/NewsView";
import PostView from "@/views/PostView";
import ProfileView from "@/views/ProfileView";
import AboutView from "@/views/AboutView";
import enterpriseHomeView from "@/views/EnterpriseHomeView";
import etpListView from "@/views/EtpListView";
import etpProfileView from "@/views/EtpProfileView";
import etpAboutView from "@/views/EtpAboutView";
import etpWorkbenchView from "@/views/EtpWorkbenchView";
import EnterpriseView from "@/views/EnterpriseView";
import Carousel from "@/components/Carousel";
import Register from "@/views/Register";


Vue.use(VueRouter)

// 管理员路由
export const adminRouters=[
    {
        path: '/adminHome',
        name: 'adminHome',
        component: AdminHomeView,
        children: [
            {
                path: '/enterprise',
                name: 'enterprise',
                component: EnterpriseView,
            },
            {
                path: '/comment',
                name: 'comment',
                component: CommentView,
            },
            {
                path: '/feedback',
                name: 'feedback',
                component: FeedBackView,
            },
            {
                path: '/recipe',
                name: 'recipe',
                component: RecipeView,
            },
            {
                path: '/post',
                name: 'post',
                component: PostView,
            },
            {
                path: '/news',
                name: 'news',
                component: NewsView,
            },
            {
                path: '/user',
                name: 'user',
                component: UserView,
            },
            {
                path: '',
                name: 'dashBoard',
                component: DashBoard,
            },
            {
                path: '/profile',
                name: 'profile',
                component: ProfileView,
            },
            {
                path: '/about',
                name: 'about',
                component: AboutView,
            },
        ]
    },
]

// 企业路由
export const enterpriseRouters=[
    {
        path: '/enterpriseHome',
        name: 'enterpriseHome',
        component: enterpriseHomeView,
        children: [
            {
                path: '/etpWorkbench',
                name: 'etpWorkbench',
                component: etpWorkbenchView,
            },
            {
                path: '/etpAbout',
                name: 'etpAbout',
                component: etpAboutView,
            },
            {
                path: '/etpList',
                name: 'etpList',
                component: etpListView,
            },
            {
                path: '/etpProfile',
                name: 'etpProfile',
                component: etpProfileView,
            },
            {
                path: '',
                name: 'carousel',
                component: Carousel,
            },
        ]
    },
]

//公共路由
const publicRoutes = [
    {
        path: '/login',
        name: 'login',
        component: LoginVue,
    },
    {
        path: "/",
        redirect: "/login",
    },
    {
        path: '/register',
        name: 'register',
        component: Register,
    },
]

// 初始路由
const router = new VueRouter({
    mode: 'history', // 即可去掉#
    routes: publicRoutes,
})

function resetToRoutes() {
    router.matcher = new VueRouter().matcher;
    router.addRoutes(publicRoutes);
}
export {resetToRoutes}
const isadmin =window.sessionStorage.getItem('admin')
if(!isadmin)  router.addRoutes(enterpriseRouters);
else router.addRoutes(adminRouters);

// 挂载路由导航守卫
router.beforeEach((to,from,next)=>{

    if(to.path==='/login' || to.path==='/forget' || to.path==='/register') return next();
    const tokenStr =window.sessionStorage.getItem('token')
    if(!tokenStr) return next('/login')
    next()
})

// const router = new VueRouter({
//     mode: 'history',
//     base: process.env.BASE_URL,
//     routes
// })

export default router
