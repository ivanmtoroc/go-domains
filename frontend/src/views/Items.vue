<template>
  <div>
    <h2 class="vue-green">Get Items</h2>
    <div v-if="isOnline">
      <error-component v-if="showError" :message="error.statusText" />
      <div v-else>
        <b-row align-h="start" class="mt-2">
          <b-col sm="12" md="4" lg="2">
            <b-dropdown :text="strLimit" variant="primary">
              <b-dropdown-item
                v-for="(value, index) in perPageValues"
                :key="index"
                @click="updateLimit(value)"
              >
                {{ value }}
              </b-dropdown-item>
            </b-dropdown>
          </b-col>
          <b-col sm="12" md="6">
            <b-pagination
              :page="currentPage"
              @input="updateCurrentPage"
              :total-rows="itemsNumber"
              :per-page="limit"
            ></b-pagination>
          </b-col>
        </b-row>
        <b-tabs class="mt-2" content-class="mt-3" pills justified>
          <b-tab title="List">
            <items-component :items="items" :isLoading="isLoading" />
          </b-tab>
          <b-tab title="JSON">
            <json-viewer-component :data="{ 'items': items }" :isLoading="isLoading" />
          </b-tab>
        </b-tabs>
      </div>
    </div>
    <error-component v-else message="No internet connection 😿" />
  </div>
</template>

<script>
import { mapState, mapGetters, mapActions } from 'vuex'

import ItemsComponent from '@/components/Items'
import JsonViewerComponent from '@/components/JsonViewer'
import ErrorComponent from '@/components/Error'

export default {
  name: 'items',
  components: {
    ItemsComponent,
    JsonViewerComponent,
    ErrorComponent
  },
  computed: {
    ...mapState('app', [
      'isOnline'
    ]),
    ...mapState('items', [
      'items',
      'itemsNumber',
      'limit',
      'offset',
      'currentPage',
      'error',
      'isLoading',
      'perPageValues'
    ]),
    ...mapGetters('items', [
      'strLimit',
      'showError'
    ])
  },
  methods: {
    ...mapActions('items', [
      'getItems',
      'updateLimit',
      'updateCurrentPage',
      'initValues'
    ])
  },
  mounted () {
    this.initValues()
    this.getItems()
  }
}
</script>
