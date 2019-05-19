<template>
  <codemirror
    :value="value"
    :options="cmOptions"
    @input="$emit('input', $event)"
  />
</template>

<script>
export default {
  props: {
    value: {
      type: String,
      default: '(+ 1 2 3 4)',
    },
  },
  data: function() {
    return {
      cmOptions: {
        tabSize: 2,
        mode: 'text/x-common-lisp',
        theme: 'base16-dark',
        lineNumbers: true,
        keyMap: 'sublime',
        extraKeys: {
          'Cmd-Enter': () => this.$emit('run'),
          'Control-Enter': () => this.$emit('run'),
        },
      },
    }
  },
}
</script>
