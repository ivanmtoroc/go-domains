<template>
  <b-card class="mt-3">
    <div v-if="isLoading" class="text-center">
      <b-spinner variant="success"></b-spinner>
    </div>
    <pre v-else class="m-0">{{ data }}</pre>
  </b-card>
</template>

<script>
export default {
  name: 'json-viewer-component',
  props: {
    data: {
      type: Object,
      default: () => {}
    },
    isLoading: {
      type: Boolean,
      default: false
    }
  }
}
</script>
