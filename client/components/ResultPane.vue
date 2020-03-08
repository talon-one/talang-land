<template>
  <code class="result-pane">
    {{ value }}
  </code>
</template>

<script>
export default {
  props: {
    value: {
      type: String,
      default: null,
    },
  },
}
</script>

<style>
.result-pane {
  width: 100%;
  height: 330px;
}
</style>
