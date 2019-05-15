<template>
  <div>
    <h2 class="vue-green">Get Domain</h2>
    <b-form @submit="getDomain" inline>
      <b-form-group>
      <b-input
        class="mb-2 mr-sm-2 mb-sm-0"
        placeholder="Domain"
        v-model="domain"
        required
      >
      </b-input>
      <b-button type="submit" variant="success">Get</b-button>
    </b-form-group>
  </b-form>
  <b-card class="mt-3" header="API Result">
    <pre class="m-0">{{ result }}</pre>
  </b-card>
</div>
</template>

<script>
export default {
  data () {
    return {
      result: {},
      domain: ''
    }
  },
  methods: {
    getDomain (evt) {
      evt.preventDefault()
      this.result = { 'data': this.domain }
    }
  }
}
</script>
