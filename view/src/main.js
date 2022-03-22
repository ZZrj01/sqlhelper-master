/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2021-10-30 15:53:11
 * @Description: 
 */
import Vue from 'vue'
import App from './App.vue'
// ElementUI
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);
// vue-router
import VueRouter from 'vue-router'
Vue.use(VueRouter)
// http
import {
  POST
} from '@/utils/http.js'
Vue.prototype.$POST = POST

import router from '@/utils/router.js'

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')