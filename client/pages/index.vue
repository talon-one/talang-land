<template>
  <client-only placeholder="Loading...">
    <v-layout column>
      <v-container class="px-6 py-6">
        <v-tooltip top>
          <template v-slot:activator="{ on }">
            <v-btn
              class="run-button"
              color="deep-purple"
              @click="run"
              v-on="on"
            >
              Run
            </v-btn>
          </template>
          <span class="tooltip">{{ runShourtcutTooltip() }}</span>
        </v-tooltip>
        <TalangEditor v-model="code" autofocus @run="run" />
      </v-container>
      <v-container class="px-6 py-6">
        <v-alert v-if="error !== null" type="error">
          {{ error }}
        </v-alert>
        <ResultPane v-if="error === null" :value="String(result)" />
      </v-container>
    </v-layout>
  </client-only>
</template>

<script>
import axios from 'axios'
import TalangEditor from '~/components/TalangEditor.vue'
import ResultPane from '~/components/ResultPane.vue'

export default {
  components: {
    TalangEditor,
    ResultPane,
  },
  data: function() {
    return {
      code: '(join (list "Welcome" "to" "Talang" "Land" "👋") " ")',
      runShourtcutTooltip: () => {
        const defaultTooltip = '⌘ + Enter'
        try {
          if (!navigator) return defaultTooltip
          return navigator.platform === 'MacIntel'
            ? defaultTooltip
            : 'Ctrl + Enter'
        } catch {
          return defaultTooltip
        }
      },
      result: '',
      error: null,
    }
  },
  methods: {
    run() {
      axios
        .get(`/api/parse?talang=${encodeURIComponent(this.code)}`)
        .then(response => {
          this.result = response.data
          this.error = null
        })
        .catch(err => {
          this.error = err.response ? err.response.data : err
          this.result = null
        })
    },
  },
}
</script>

<style>
.tooltip {
  font-family: 'Courier New', Courier, monospace;
}
.run-button {
  margin: 0.5rem;
}
</style>
