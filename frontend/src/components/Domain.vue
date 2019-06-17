<template>
  <b-card>
    <div v-show="showInfo">
      <div v-if="isLoading" class="text-center">
        <b-spinner variant="success" label="Spinning"></b-spinner>
      </div>
      <b-card-text v-else>
        <h3>
          <b-img :src="domainInfo.logo" width="30" height="30"></b-img>
          {{ domainInfo.title }}
        </h3>
        <b-table class="mt-2" hover :items="metrics"></b-table>
        <b-list-group>
          <Server
            v-for="(server, index) in domainInfo.servers"
            :key="index"
            :server="server"
            :index="index"
          />
        </b-list-group>
      </b-card-text>
    </div>
  </b-card>
</template>

<script>
import parser from '@/utilities/parser'
import Server from '@/components/Server'

export default {
  props: {
    domainInfo: Object,
    isLoading: Boolean,
    showInfo: Boolean
  },
  components: {
    Server
  },
  computed: {
    metrics () {
      return [
        {
          metric: 'Is down',
          value: parser.booleanToHuman(this.domainInfo.isDown)
        },
        {
          metric: 'Servers changed',
          value: parser.booleanToHuman(this.domainInfo.serversChanged)
        },
        {
          metric: 'SSL grade',
          value: this.domainInfo.sslGrade
        },
        {
          metric: 'Previous SSL grade',
          value: this.domainInfo.previousSslGrade
        }
      ]
    }
  }
}
</script>
