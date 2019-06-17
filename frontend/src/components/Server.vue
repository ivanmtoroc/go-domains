<template>
  <b-list-group-item href="#" class="flex-column align-items-start">
    <div class="d-flex w-100 justify-content-between">
      <h5 class="mb-1">{{ server['address'] }}</h5>
      <small>{{ `Server #${index}` }}</small>
    </div>
    <b-table class="mt-2" hover :items="serverInfo"></b-table>
  </b-list-group-item>
</template>

<script>
export default {
  props: {
    server: Object,
    index: Number
  },
  computed: {
    serverInfo () {
      return [
        {
          metric: 'SSL grade',
          value: this.server.sslGrade
        },
        {
          metric: 'Country',
          value: this.server.country
        },
        {
          metric: 'Owner',
          value: this.server.owner
        }
      ]
    }
  }
}
</script>
