import http from '@/utilities/http'

const state = {
  domainInfo: {},
  domainName: '',
  isLoading: false,
  showInfo: false
}

const mutations = {
  setDomainInfo (state, newDomainInfo) {
    state.domainInfo = newDomainInfo
  },
  setDomainName (state, newDomainName) {
    var regex = /^[a-zA-Z0-9-.]*$/
    state.domainName = regex.test(newDomainName) ? newDomainName : state.domainName
  },
  setIsLoading (state, newIsLoading) {
    state.isLoading = newIsLoading
  },
  setShowInfo (state, newShowInfo) {
    state.showInfo = newShowInfo
  }
}

const actions = {
  async getDomain ({ state, commit }, event) {
    event.preventDefault()
    commit('setShowInfo', true)
    commit('setIsLoading', true)
    const data = await http.get(`domains/${state.domainName}`)
    commit('setIsLoading', false)
    commit('setDomainInfo', data)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
