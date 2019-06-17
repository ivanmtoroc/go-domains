<template>
  <div>
    <h2 class="vue-green">Get Domain</h2>
    <b-form @submit="getDomain" inline>
      <b-form-group>
        <input
        class="form-control mb-2 mr-sm-2 mb-sm-0"
        placeholder="Domain name"
        :value="domainName"
        @input="updateDomainName"
        pattern="^[a-zA-Z0-9-.]*$"
        required
        />
        <b-button :disabled="isLoading" type="submit" variant="success">
          Get
        </b-button>
      </b-form-group>
    </b-form>
    <b-tabs class="mt-3" content-class="mt-3" pills justified>
      <b-tab title="Web">
        <Domain
        :domainInfo="domainInfo"
        :isLoading="isLoading"
        :showInfo="showInfo"
        />
      </b-tab>
      <b-tab title="JSON">
        <JsonViewer :data="domainInfo" :isLoading="isLoading"/>
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>
import { mapState, mapMutations, mapActions } from 'vuex'
import JsonViewer from '@/components/JsonViewer'
import Domain from '@/components/Domain'

export default {
  components: {
    JsonViewer,
    Domain
  },
  computed: {
    ...mapState('domains', [
      'domainInfo',
      'domainName',
      'isLoading',
      'showInfo'
    ])
  },
  methods: {
    ...mapMutations('domains', [
      'setDomainName'
    ]),
    ...mapActions('domains', [
      'getDomain'
    ]),
    updateDomainName (event) {
      this.setDomainName(event.target.value)
    }
  }
}
</script>
