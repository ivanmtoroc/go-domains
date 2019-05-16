<template>
  <div>
    <h2 class="vue-green">Get Domain</h2>
    <b-form @submit="getDomain" inline>
      <b-form-group>
        <b-input
        class="mb-2 mr-sm-2 mb-sm-0"
        placeholder="Domain"
        v-model="domain"
        required
        >
      </b-input>
      <b-button type="submit" variant="success">Get</b-button>
    </b-form-group>
  </b-form>
  <b-card class="mt-3" header="API Result">
    <pre class="m-0">{{ result }}</pre>
  </b-card>
</div>
</template>

<script>
import axios from 'axios'

const http = axios.create({
  baseURL: 'http://localhost:3333'
})

export default {
  data () {
    return {
      result: {},
      domain: '',
      loading: true,
      error: false
    }
  },
  methods: {
    async getDomain (evt) {
      evt.preventDefault()
      this.loading = true
      this.error = false
      const response = await http.get('domains/' + this.domain)
        .catch(errors => {
          this.result = errors.response.data
          this.error = true
        })
      if (!this.error) {
        this.result = response.data
      }
      this.loading = false
    }
  }
}
</script>
