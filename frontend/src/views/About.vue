<template>
  <div>
    <h2 class="vue-green">About</h2>
    <p>
      Rest API that allows us to obtain information about a server
      and to know if the configurations have changed.<br/>
      <br/>
      <router-link class="vue-green" :to="{ name: 'domains' }">Get domains</router-link>
      receive a domain like
      <a class="vue-green" href="http://truora.com" target="_blank">truora.com</a>
      and return a JSON with the information.<br/>
      <br/>
      <router-link class="vue-green" :to="{ name: 'items' }">Get items</router-link>
      lists the servers that have been previously consulted.
    </p>
  </div>
</template>

<style scoped>
p {
  font-size: 20px;
}
</style>
