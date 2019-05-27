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
  <ApiViewer :result="result" :loading="loading"/>
</div>
</template>

<script>
import { mapActions } from 'vuex'
import ApiViewer from '@/components/ApiViewer.vue'

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
    ApiViewer
  },
  methods: {
    ...mapActions('utilities', ['get']),
    async getDomain (evt) {
      evt.preventDefault()
      this.loading = true
      this.result = await this.get('domains/' + this.domain)
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
