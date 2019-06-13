<template>
  <v-app dark>
    <v-toolbar class="fix-high" fixed app>
      <nuxt-link to="/" title="Home">
        <img src="/favicon.ico" alt="Talang" width="36px" height="36px" />
      </nuxt-link>
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-btn
        href="https://github.com/talon-one/talang/"
        title="https://github.com/talon-one/talang/"
        target="_blank"
        flat
      >
        GitHub
        <v-icon right>open_in_new</v-icon>
      </v-btn>
      <v-btn to="/about" title="About" nuxt flat>About</v-btn>
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
