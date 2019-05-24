import axios from 'axios'

const http = axios.create({
  baseURL: 'http://localhost:3333/api/v1/'
})

const actions = {
  async get (context, url) {
    const response = await http.get(url)
      .catch(errors => {
        return errors.response.data
      })
    return response.data
  }
}

export default {
  namespaced: true,
  actions
}
