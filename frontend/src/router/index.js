import Vue from 'vue'
import Router from 'vue-router'

import domains from './modules/domains'
import items from './modules/items'
import about from './modules/about'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    { ...domains },
    { ...items },
    { ...about },
    {
      path: '*',
      redirect: '/'
    }
  ]
})
