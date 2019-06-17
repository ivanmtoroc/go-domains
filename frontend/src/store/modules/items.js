import http from '@/utilities/http'

const state = {
  items: {},
  itemsNumber: 0,
  limit: 0,
  offset: 0,
  currentPage: 1,
  perPageValues: [5, 10, 20, 30, 50],
  isLoading: true
}

const getters = {
  strLimit: state => state.limit.toString()
}

const mutations = {
  setItems: (state, newItems) => { state.items = newItems },
  setItemsNumber: (state, newItemsNumber) => { state.itemsNumber = newItemsNumber },
  setLimit: (state, newLimit) => { state.limit = newLimit },
  setOffset: (state, newOffset) => { state.offset = newOffset },
  setCurrentPage: (state, newCurrentPage) => { state.currentPage = newCurrentPage },
  setIsLoading: (state, newIsLoading) => { state.isLoading = newIsLoading }
}

const actions = {
  async getItems ({ state, commit }) {
    commit('setIsLoading', true)
    let url = `items?limit=${state.limit}&offset=${state.offset}`
    let data = await http.get(url)
    commit('setItems', data['items'])
    commit('setItemsNumber', data['totalItems'])
    commit('setIsLoading', false)
  },
  updateLimit ({ commit, dispatch }, newLimit) {
    commit('setLimit', newLimit)
    dispatch('getItems')
  },
  updateOffset ({ commit, dispatch }, newOffset) {
    commit('setOffset', newOffset)
    dispatch('getItems')
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
