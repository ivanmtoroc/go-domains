
import Vue from 'vue'
import Vuex from 'vuex'

import domains from './modules/domains'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    domains
  }
})
