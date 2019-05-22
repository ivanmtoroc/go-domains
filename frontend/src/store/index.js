
import Vue from 'vue'
import Vuex from 'vuex'

import utilities from './modules/utilities'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    utilities
  }
})
