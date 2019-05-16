import Vue from 'vue'
import Router from 'vue-router'

import GetDomain from './views/GetDomain.vue'
import About from './views/About.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    { path: '/', name: 'home', component: GetDomain },
    { path: '/about', name: 'about', component: About },
    { path: '*', redirect: '/' }
  ]
})
