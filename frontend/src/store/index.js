
import Vue from 'vue'
import Vuex from 'vuex'

import domains from './modules/domains'
import items from './modules/items'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    domains,
    items
  }
})
