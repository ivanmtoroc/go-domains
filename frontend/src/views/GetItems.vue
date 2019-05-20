<template>
  <div>
    <h2 class="vue-green">Get Items</h2>
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
    async getItems () {
      this.loading = true
      this.error = false
      const response = await http.get('items')
        .catch(errors => {
          this.result = errors.response.data
          this.error = true
        })
      if (!this.error) {
        this.result = response.data
      }
      this.loading = false
    }
  },
  mounted () {
    this.getItems()
  }
}
</script>
