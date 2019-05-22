<template>
  <div>
    <h2 class="vue-green">Get Items</h2>
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
      loading: true
    }
  },
  components: {
    ApiViewer
  },
  methods: {
    ...mapActions('utilities', ['get']),
    async getItems () {
      this.loading = true
      this.result = await this.get('items')
      this.loading = false
    }
  },
  mounted () {
    this.getItems()
  }
}
</script>
