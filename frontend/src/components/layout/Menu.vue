<template>
  <div>
    <router-link to="/domains">
      <b-button
        pill
        block
        size="lg"
        variant="outline-primary"
      >
        Get Domain
      </b-button>
    </router-link>
    <router-link to="/items">
      <b-button
        pill
        block
        size="lg"
        variant="outline-danger"
      >
        Get Items
      </b-button>
    </router-link>
    <router-link to="/about">
      <b-button
        pill
        block
        size="lg"
        variant="outline-success"
      >
        About
      </b-button>
    </router-link>
    <b-button
      size="lg"
      variant="outline-secondary"
      href="http://github.com/ivanmtoroc/go-domains"
      target="_blank"
    >
      Github
    </b-button>
  </div>
</template>

<script>
export default {
  name: 'menu-component'
}
</script>
