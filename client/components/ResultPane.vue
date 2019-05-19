<template>
  <codemirror :value="value" :options="cmOptions" />
</template>

<script>
export default {
  props: {
    value: {
      type: String,
      default: null,
    },
  },
  data: function() {
    return {
      cmOptions: {
        tabSize: 2,
        mode: 'text/javascript',
        theme: 'solarized light',
        readOnly: true,
      },
    }
  },
}
</script>

<style>
.cm-s-solarized.CodeMirror {
  box-shadow: none;
}
</style>
