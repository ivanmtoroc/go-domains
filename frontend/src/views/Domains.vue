<template>
  <div>
    <h2 class="vue-green">Get Domain</h2>
    <b-form @submit="getDomain" inline>
      <b-form-group>
        <input
        class="form-control mb-2 mr-sm-2 mb-sm-0"
        placeholder="Domain"
        v-model="domain"
        pattern="^[a-zA-Z0-9-.]*$"
        required
        />
        <b-button :disabled="loading" type="submit" variant="success">Get</b-button>
      </b-form-group>
    </b-form>
    <b-tabs class="mt-3" content-class="mt-3" pills justified>
      <b-tab title="Web">
        <Domain :result="result"/>
      </b-tab>
      <b-tab title="JSON">
        <JsonViewer :result="result" :loading="loading"/>
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>
import http from '@/utilities/http'
import JsonViewer from '@/components/JsonViewer'
import Domain from '@/components/Domain'

var re = /^[a-zA-Z0-9-.]*$/

export default {
  data () {
    return {
      result: {},
      domain: '',
      loading: false
    }
  },
  components: {
    JsonViewer,
    Domain
  },
  methods: {
    async getDomain (evt) {
      evt.preventDefault()
      this.loading = true
      this.result = await http.get('domains/' + this.domain)
      this.loading = false
    }
  },
  watch: {
    domain (newValue, oldValue) {
      this.domain = re.test(newValue) ? newValue : oldValue
    }
  }
}
</script>
