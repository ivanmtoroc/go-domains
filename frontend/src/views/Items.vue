<template>
  <div>
    <h2 class="vue-green">Get Items</h2>
    <b-row align-h="start">
      <b-col sm="12" md="4" lg="2">
        <b-dropdown :text="strLimit" variant="primary" class="m-2">
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
          page="currentPage"
          @input="updateCurrentPage"
          :total-rows="itemsNumber"
          :per-page="limit"
          class="m-2"
        ></b-pagination>
      </b-col>
    </b-row>
    <b-tabs class="mt-3" content-class="mt-3" pills justified>
      <b-tab title="Web">
        <b-card>
          <div v-if="isLoading" class="text-center">
            <b-spinner variant="success" label="Spinning"></b-spinner>
          </div>
          <b-card-text v-else>
            <b-list-group>
              <Item
                v-for="(item, index) in items"
                :key="index"
                :item="item"
              />
            </b-list-group>
          </b-card-text>
        </b-card>
      </b-tab>
      <b-tab title="JSON">
        <JsonViewer :data="{ items: items }" :isLoading="isLoading"/>
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>
import { mapState, mapGetters, mapMutations, mapActions } from 'vuex'
import Item from '@/components/Item'
import JsonViewer from '@/components/JsonViewer'

export default {
  components: {
    Item,
    JsonViewer
  },
  computed: {
    ...mapState('items', [
      'items',
      'itemsNumber',
      'limit',
      'offset',
      'currentPage',
      'perPageValues',
      'isLoading'
    ]),
    ...mapGetters('items', [
      'strLimit'
    ])
  },
  methods: {
    ...mapMutations('items', [
      'setLimit',
      'setOffset',
      'setCurrentPage'
    ]),
    ...mapActions('items', [
      'getItems',
      'updateLimit',
      'updateCurrentPage'
    ])
  },
  mounted () {
    this.setLimit(this.perPageValues[0])
    this.setOffset(0)
    this.setCurrentPage(1)
    this.getItems()
  }
}
</script>
