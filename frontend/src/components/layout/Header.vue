<template>
  <b-navbar class="navbar">
    <b-navbar-brand>
      <h1>Go Domains!</h1>
      <p>
        💻 with ❤️ by
        <a class="vue-green" href="http://ivanmtoroc.me" target="_blank">
          @ivanmtoroc
        </a>
      </p>
    </b-navbar-brand>
  </b-navbar>
</template>

<script>
export default {
  name: 'header-component'
}
</script>

<style scoped>
.navbar {
  background-color: #1e3044;
}
a:hover {
  color: #2edc8e;
}
p, h1 {
  color: white;
  margin: 0px 10px;
}
</style>
