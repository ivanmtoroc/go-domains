<template>
  <b-card class="mt-3" header="JSON viewer">
    <div v-if="isLoading" class="text-center">
      <b-spinner variant="success" label="Spinning"></b-spinner>
    </div>
    <pre v-else class="m-0">{{ data }}</pre>
  </b-card>
</template>

<script>
export default {
  props: {
    data: Object,
    isLoading: Boolean
  }
}
</script>
