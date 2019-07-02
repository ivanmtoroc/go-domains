<template>
  <div>
    <h2 class="vue-green">Get Domain</h2>
    <b-form @submit="getDomain" inline>
      <b-form-group>
        <input
          :value="domainName"
          @input="updateDomainName"
          class="form-control mb-2 mr-sm-2 mb-sm-0"
          placeholder="Domain name"
          pattern="^[a-zA-Z0-9-.]*$"
          required
        />
        <b-button :disabled="isLoading || !isOnline" type="submit" variant="success">
          Get!
        </b-button>
      </b-form-group>
    </b-form>
    <div v-if="isOnline">
      <error-component v-if="showError" :message="error.statusText" />
      <h3 v-else-if="!showDomain" class="vue-blue mt-3">
        Look for the information of some domain!
      </h3>
      <b-tabs v-else class="mt-3" content-class="mt-3" pills justified>
        <b-tab title="Table">
          <domain-component :domain="domain" :isLoading="isLoading" />
        </b-tab>
        <b-tab title="JSON">
          <json-viewer-component :data="domain" :isLoading="isLoading" />
        </b-tab>
      </b-tabs>
    </div>
    <error-component v-else message="No internet connection 😿" />
  </div>
</template>

<script>
import { mapState, mapGetters, mapActions } from 'vuex'

import DomainComponent from '@/components/Domain'
import JsonViewerComponent from '@/components/JsonViewer'
import ErrorComponent from '@/components/Error'

export default {
  name: 'domain',
  components: {
    DomainComponent,
    JsonViewerComponent,
    ErrorComponent
  },
  computed: {
    ...mapState('app', [
      'isOnline'
    ]),
    ...mapState('domains', [
      'domainName',
      'domain',
      'error',
      'isLoading'
    ]),
    ...mapGetters('domains', [
      'showDomain',
      'showError'
    ])
  },
  methods: {
    ...mapActions('domains', [
      'getDomain',
      'updateDomainName'
    ])
  }
}
</script>
