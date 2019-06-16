<template>
  <div>
    <b-card title="Title" sub-title="domain.url">
      <b-card-text>
        {{ result }}
      </b-card-text>
    </b-card>
  </div>
</template>

<script>
export default {
  props: {
    result: Object
  }
}
</script>
