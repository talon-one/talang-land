<template>
  <v-app dark>
    <v-toolbar class="fix-high" fixed app>
      <nuxt-link to="/">
        <img src="/favicon.ico" alt="Talang" width="36px" height="36px" />
      </nuxt-link>
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <nuxt-link to="/about">About</nuxt-link>
    </v-toolbar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>
    <v-footer class="pad-footer" app>
      <span>Talon.One | Talang Land</span>
      <v-spacer />
      <span>&copy; 2019</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      title: 'Talang Land',
    }
  },
}
</script>

<style>
.fix-high {
  /* this here comes to keep toolbar higher than codemirror inputs */
  z-index: 5 !important;
}
.pad-footer {
  padding: 0 2rem;
}
</style>
