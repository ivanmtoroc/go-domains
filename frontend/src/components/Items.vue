<template>
  <b-card>
    <div v-if="isLoading" class="text-center">
      <b-spinner variant="success"></b-spinner>
    </div>
    <b-card-text v-else>
      <b-list-group>
        <b-list-group-item
          v-for="(item, index) in items"
          :key="index"
          :href="`http://${item.name}`"
          target="_blank"
        >
          <div class="d-flex w-100 justify-content-between align-items-center">
            <p class="m-0">
              <b-img :src="item.info.logo" width="30" height="30"></b-img>
              {{ item.info.title }}
            </p>
            <b-badge :variant="item.info.isDown ? 'danger' : 'success'" pill>
              {{ item.info.isDown | status }}
            </b-badge>
          </div>
        </b-list-group-item>
      </b-list-group>
    </b-card-text>
  </b-card>
</template>

<script>
export default {
  name: 'items-component',
  props: {
    items: {
      type: Array,
      default: () => []
    },
    isLoading: {
      type: Boolean,
      default: false
    }
  },
  filters: {
    status: value => value ? 'Down' : 'Up'
  }
}
</script>
