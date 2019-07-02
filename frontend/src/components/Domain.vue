<template>
  <b-card>
    <div v-if="isLoading" class="text-center">
      <b-spinner variant="success"></b-spinner>
    </div>
    <b-card-text v-else>
      <h3 class="mb-2">
        <b-img :src="domain.logo" width="30" height="30"></b-img>
        {{ domain.title }}
      </h3>
      <b-table :items="items" stacked class="mt-3"></b-table>
      <servers-component :servers="domain.servers" />
    </b-card-text>
  </b-card>
</template>

<script>
import converter from '@/utilities/converter'
import ServersComponent from '@/components/Servers'

export default {
  name: 'domain-component',
  components: {
    ServersComponent
  },
  props: {
    domain: {
      type: Object,
      default: () => {}
    },
    isLoading: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    items () {
      return [
        {
          'Is down': converter.affirmation(this.domain.isDown),
          'Servers changed': converter.affirmation(this.domain.serversChanged),
          'SSL grade': this.domain.sslGrade,
          'Previous SSL grade': this.domain.previousSslGrade
        }
      ]
    }
  }
}
</script>
