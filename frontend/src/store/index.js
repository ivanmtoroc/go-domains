import Vue from 'vue'
import Vuex from 'vuex'

import app from './modules/app'
import domains from './modules/domains'
import items from './modules/items'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    app,
    domains,
    items
  }
})
