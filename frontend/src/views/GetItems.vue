<template>
  <div>
    <h2 class="vue-green">Get Items</h2>
    <b-row align-h="start">
    <b-col sm="12" md="4" lg="2">
      <b-dropdown :text="strLimit" variant="primary" class="m-2">
        <b-dropdown-item @click="setLimit(10)">10</b-dropdown-item>
        <b-dropdown-item @click="setLimit(20)">20</b-dropdown-item>
        <b-dropdown-item @click="setLimit(30)">30</b-dropdown-item>
        <b-dropdown-item @click="setLimit(50)">50</b-dropdown-item>
      </b-dropdown>
    </b-col>
    <b-col sm="12" md="6">
      <b-pagination
        v-model="currentPage"
        :total-rows="itemsNumber"
        :per-page="limit"
        class="m-2"
      ></b-pagination>
    </b-col>
  </b-row>
    <ApiViewer :result="items" :loading="loading"/>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import ApiViewer from '@/components/ApiViewer.vue'

export default {
  data () {
    return {
      items: {},
      itemsNumber: 0,
      limit: 0,
      skip: 0,
      currentPage: 1,
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
      let url = `items?limit=${this.limit}&offset=${this.skip}`
      let data = await this.get(url)
      this.items = { 'items': data['items'] }
      this.itemsNumber = data['total_items']
      this.loading = false
    },
    setLimit (limit) {
      this.limit = limit
      this.getItems()
    }
  },
  watch: {
    currentPage (value) {
      this.skip = this.limit * (value - 1)
      this.getItems()
    }
  },
  computed: {
    strLimit () {
      return this.limit.toString()
    }
  },
  mounted () {
    this.setLimit(10)
  }
}
</script>
