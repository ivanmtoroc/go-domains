<template>
  <div>
    <h2 class="vue-green">About</h2>
    <p>
      Web application to obtain information about a domain and its servers.
      For example: the SSL grade of the servers and if the configurations
      of the servers have changed one hour beforehand.<br />
      <br />
      In
      <router-link class="vue-green" :to="{ name: 'domains' }">
        Get Domain
      </router-link>
      a domain, like
      <a class="vue-green" href="http://truora.com" target="_blank">
        truora.com
      </a>
      , is entered and the information of the domain and its servers is obtained.<br />
      <br />
      In
      <router-link class="vue-green" :to="{ name: 'items' }">
        Get Items
      </router-link>
      , the domains that have been consulted previously are listed.
    </p>
  </div>
</template>

<script>
export default {
  name: 'about'
}
</script>
