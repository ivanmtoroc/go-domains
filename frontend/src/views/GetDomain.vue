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
      <b-button :disabled="loading" type="submit" variant="success">Get</b-button>
    </b-form-group>
  </b-form>
  <ApiViewer :result="result" :loading="loading"/>
</div>
</template>

<script>
import { mapActions } from 'vuex'
import ApiViewer from '@/components/ApiViewer.vue'

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
  }
}
</script>
