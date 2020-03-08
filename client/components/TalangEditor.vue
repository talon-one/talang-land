<template>
  <MonacoEditor
    ref="editor"
    class="editor"
    :value="value"
    :options="opts"
    theme="vs-dark"
    language="clojure"
    @change="$emit('input', $event)"
  />
</template>

<script>
import MonacoEditor from 'vue-monaco'

export default {
  components: {
    MonacoEditor,
  },
  props: {
    value: {
      type: String,
      default: '(+ 1 2 3 4)',
    },
    autofocus: {
      type: Boolean,
      default: false,
    },
  },
  data: function() {
    return {
      opts: {
        minimap: {
          enabled: false,
        },
      },
    }
  },
  mounted() {
    const editorInstance = this.$refs.editor.getEditor()
    if (!editorInstance || !window.monaco) return
    editorInstance.addAction({
      id: 'run-talang-code',
      label: 'Run Talang',
      keybindings: [window.monaco.KeyMod.CtrlCmd | window.monaco.KeyCode.Enter],
      run: () => this.$emit('run'),
    })

    if (this.$props.autofocus) {
      this.$refs.editor.focus()
    }
  },
}
</script>

<style>
.editor {
  width: 100%;
  height: 300px;
}
</style>
