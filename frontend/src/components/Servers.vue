<template>
  <b-list-group>
    <b-list-group-item
      class="flex-column align-items-start"
      v-for="(item, index) in items"
      :key="index"
    >
      <h4 class="mb-2">
        {{ `Server #${index + 1}` }}
      </h4>
      <b-table :items="item" stacked class="mt-3"></b-table>
    </b-list-group-item>
  </b-list-group>
</template>

<script>
export default {
  name: 'servers-component',
  props: {
    servers: {
      type: Array,
      default: () => []
    }
  },
  computed: {
    items () {
      return this.servers.map(server => [
        {
          'Address': server.address,
          'SSL grade': server.sslGrade,
          'Country': server.country,
          'Owner': server.owner
        }
      ])
    }
  }
}
</script>
