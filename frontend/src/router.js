import Vue from 'vue'
import Router from 'vue-router'

import GetDomain from './views/GetDomain.vue'
import GetItems from './views/GetItems.vue'
import About from './views/About.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'domains',
      component: GetDomain
    },
    {
      path: '/items',
      name: 'items',
      component: GetItems
    },
    {
      path: '/about',
      name: 'about',
      component: About
    },
    {
      path: '*',
      redirect: '/'
    }
  ]
})
