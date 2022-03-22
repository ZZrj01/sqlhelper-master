import VueRouter from 'vue-router'
// 路由数组
const routes = [{
    path: "/",
    name: "main",
    component: () => import('@/views/main/index.vue')
}]
// 创建路由
const router = new VueRouter({
    routes,
})

// 路由前置守卫
router.beforeEach((to, from, next) => {
    // 非登录页面 
    // if (to.path != "/login") {

    // }
    next()
})

export default router