<template>
  <div id="app">
    <HeaderComponent/>
    <b-container fluid>
      <b-row>
        <b-col class="column" sm="12" md="4">
          <div class="section">
            <router-link to="/">
              <b-button pill block size="lg" variant="outline-primary">
                Get Domain
              </b-button>
            </router-link>
            <router-link to="/items">
              <b-button pill block size="lg" variant="outline-danger">
                Get Items
              </b-button>
            </router-link>
            <router-link to="/about">
              <b-button pill block size="lg" variant="outline-success">
                About
              </b-button>
            </router-link>
            <a href="http://github.com/ivanmtoroc/go-domains" target="_blank">
              <b-button size="lg" variant="outline-secondary">
                Github
              </b-button>
            </a>
          </div>
        </b-col>
        <b-col sm="12" md="8">
          <div class="section">
            <router-view/>
          </div>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import HeaderComponent from '@/components/Header'

export default {
  components: {
    HeaderComponent
  }
}
</script>

<style>
a {
  text-decoration: none;
}
a:hover {
  text-decoration: none !important;
}
h2 {
  font-size: 40px !important;
}
.btn:focus {
  box-shadow: none !important;
}
.section .btn-lg {
  margin: 20px 0px;
}
.section {
  padding: 25px;
}
.vue-green {
  color: #42B883;
}
.vue-blue {
  color: #35495E;
}
</style>
