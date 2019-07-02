const state = {
  isOnline: navigator.onLine || false,
  events: ['online', 'offline']
}

const mutations = {
  UPDATE_ONLINE_STATUS: state => {
    state.isOnline = navigator.onLine || false
  }
}

export default {
  namespaced: true,
  state,
  mutations
}
