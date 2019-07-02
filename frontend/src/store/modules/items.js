import http from '@/utilities/http'

const state = {
  items: [],
  itemsNumber: 0,
  limit: 0,
  offset: 0,
  currentPage: 1,
  error: {},
  isLoading: false,
  perPageValues: [5, 10, 20, 30, 50]
}

const getters = {
  strLimit: state => state.limit.toString(),
  showError: state => state.error.status != null
}

const mutations = {
  SET_ITEMS: (state, newItems) => {
    state.items = newItems
  },
  SET_ITEMS_NUMBER: (state, newItemsNumber) => {
    state.itemsNumber = newItemsNumber
  },
  SET_LIMIT: (state, newLimit) => {
    state.limit = newLimit
  },
  SET_OFFSET: (state, newOffset) => {
    state.offset = newOffset
  },
  SET_CURRENT_PAGE: (state, newCurrentPage) => {
    state.currentPage = newCurrentPage
  },
  SET_ERROR: (state, newError) => {
    state.error = newError
  },
  SET_LOAD_STATUS: (state, newLoadStatus) => {
    state.isLoading = newLoadStatus
  }
}

const actions = {
  getItems: async ({ state, commit }) => {
    commit('SET_ERROR', {})
    commit('SET_LOAD_STATUS', true)
    const data = await http.get(`items?limit=${state.limit}&offset=${state.offset}`)
    if (data.status == null) {
      commit('SET_ITEMS', data['items'])
      commit('SET_ITEMS_NUMBER', data['totalItems'])
    } else {
      commit('SET_ERROR', data)
    }
    commit('SET_LOAD_STATUS', false)
  },
  updateLimit: ({ commit, dispatch }, newLimit) => {
    commit('SET_LIMIT', newLimit)
    dispatch('getItems')
  },
  updateOffset: ({ commit, dispatch }, newOffset) => {
    commit('SET_OFFSET', newOffset)
    dispatch('getItems')
  },
  updateCurrentPage: ({ state, commit, dispatch }, newCurrentPage) => {
    if (newCurrentPage == null) { return }
    commit('SET_CURRENT_PAGE', newCurrentPage)
    dispatch('updateOffset', state.limit * (newCurrentPage - 1))
  },
  initValues: ({ state, commit }) => {
    commit('SET_LIMIT', state.perPageValues[0])
    commit('SET_OFFSET', 0)
    commit('SET_CURRENT_PAGE', 1)
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
