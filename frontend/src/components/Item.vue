<template>
  <b-list-group-item :href="'http://' + item.name" target="_blank" class="flex-column align-items-start">
    <div class="d-flex w-100 justify-content-between align-items-center">
      <h5 class="mb-1">
        <b-img :src="item.info.logo" width="30" height="30"></b-img>
        {{ item.info.title }}
      </h5>
      <b-badge
        :variant="item.info.isDown ? 'danger' : 'success'"
        pill
      >
        {{ isDown }}
      </b-badge>
    </div>
  </b-list-group-item>
</template>

<script>
export default {
  props: {
    item: Object
  },
  computed: {
    isDown () {
      return this.item.info.isDown ? 'Down' : 'Up'
    }
  }
}
</script>
