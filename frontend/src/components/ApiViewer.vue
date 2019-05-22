<template>
  <b-card class="mt-3" header="API Result">
    <div v-if="loading" class="text-center">
      <b-spinner variant="success" label="Spinning"></b-spinner>
    </div>
    <pre v-else class="m-0">{{ result }}</pre>
  </b-card>
</template>

<script>
export default {
  props: {
    result: Object,
    loading: Boolean
  }
}
</script>
