<template>
  <div id="app">
    <header-component />
    <b-container fluid>
      <b-row>
        <b-col sm="12" md="4">
          <div class="section">
            <menu-component />
          </div>
        </b-col>
        <b-col sm="12" md="8">
          <div class="section">
            <router-view />
          </div>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import { mapState, mapMutations } from 'vuex'

import HeaderComponent from '@/components/layout/Header'
import MenuComponent from '@/components/layout/Menu'

export default {
  name: 'app',
  components: {
    HeaderComponent,
    MenuComponent
  },
  computed: {
    ...mapState('app', [
      'events'
    ])
  },
  methods: {
    ...mapMutations('app', [
      'UPDATE_ONLINE_STATUS'
    ])
  },
  mounted () {
    this.events.forEach(event => {
      window.addEventListener(event, this.UPDATE_ONLINE_STATUS)
    })
  },
  beforeDestroy () {
    this.events.forEach(event => {
      window.removeEventListener(event, this.UPDATE_ONLINE_STATUS)
    })
  }
}
</script>

<style>
a {
  text-decoration: none;
}
p {
  font-size: 20px;
}
a:hover {
  text-decoration: none !important;
}
h2 {
  font-size: 40px !important;
  margin: 10px 0px !important;
}
.btn:focus {
  box-shadow: none !important;
}
.section {
  padding: 25px;
}
.section .btn-lg {
  margin: 20px 0px;
}
.vue-green {
  color: #42b883;
}
.vue-blue {
  color: #35495e;
}
</style>
