<template>
  <no-ssr placeholder="Loading...">
    <v-layout column>
      <v-container>
        <v-tooltip top>
          <template v-slot:activator="{ on }">
            <v-btn color="deep-purple" @click="run" v-on="on">Run</v-btn>
          </template>
          <span class="tooltip">{{ runShourtcutTooltip() }}</span>
        </v-tooltip>
        <v-flex>
          <TalangEditor v-model="code" @run="run" />
        </v-flex>
      </v-container>
      <v-container>
        <v-flex>
          <v-alert v-if="!!error" type="error" value="error">
            {{ error }}
          </v-alert>
          <ResultPane v-if="!error" :value="String(result)" />
        </v-flex>
      </v-container>
    </v-layout>
  </no-ssr>
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
.result-pane {
  background: #151515;
  color: #e0e0e0;
}
.tooltip {
  font-family: 'Courier New', Courier, monospace;
}
</style>
