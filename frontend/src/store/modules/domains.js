import http from '@/utilities/http'

const REGEX = /^[a-zA-Z0-9-.]*$/

const state = {
  domainName: '',
  domain: {},
  error: {},
  isLoading: false
}

const getters = {
  showDomain: state => state.domain.title != null || state.isLoading,
  showError: state => state.error.status != null
}

const mutations = {
  SET_DOMAIN_NAME: (state, newDomainName) => {
    state.domainName = newDomainName
  },
  SET_DOMAIN: (state, newDomain) => {
    state.domain = newDomain
  },
  SET_ERROR: (state, newError) => {
    state.error = newError
  },
  SET_LOAD_STATUS: (state, newLoadStatus) => {
    state.isLoading = newLoadStatus
  }
}

const actions = {
  getDomain: async ({ state, commit }, event) => {
    event.preventDefault()
    commit('SET_ERROR', {})
    commit('SET_LOAD_STATUS', true)
    const data = await http.get(`domains/${state.domainName}`)
    if (data.status == null) {
      commit('SET_DOMAIN', data)
    } else {
      commit('SET_ERROR', data)
    }
    commit('SET_LOAD_STATUS', false)
  },
  updateDomainName: ({ commit }, event) => {
    let newDomainName = event.target.value
    if (REGEX.test(newDomainName)) {
      commit('SET_DOMAIN_NAME', newDomainName)
    }
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
