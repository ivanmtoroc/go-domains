<template>
  <b-alert
    show
    variant="danger"
    class="text-center mt-3"
  >
    <p>{{ message }}</p>
    <!-- <img
      class="error-img"
      src="@/assets/images/error.jpg"
      alt="Error image"
    > -->
  </b-alert>
</template>

<script>
export default {
  name: 'error-component',
  props: {
    message: {
      type: String,
      default: ''
    }
  }
}
</script>

<style scoped>
p {
  margin: 0px;
}
.error-img {
  border-radius: 10px;
}
</style>
