<template>
  <v-app>
    <v-app-bar class="fix-high" absolute app hide-on-scroll>
      <nuxt-link class="title-link-anchor" to="/" title="Home">
        <img src="/favicon.ico" alt="Talang" width="36px" height="36px" />
        <v-toolbar-title class="title-link" v-text="title" />
      </nuxt-link>
      <v-spacer />
      <v-btn
        href="https://github.com/talon-one/talang/"
        title="https://github.com/talon-one/talang/"
        target="_blank"
        text
      >
        GitHub
        <v-icon right>mdi-open-in-new</v-icon>
      </v-btn>
      <v-btn to="/about" title="About" nuxt text>About</v-btn>
    </v-app-bar>
    <v-content>
      <v-container class="px-6 py-6">
        <nuxt />
      </v-container>
    </v-content>
    <v-footer class="pad-footer">
      <span>Talon.One | Talang Land</span>
      <v-spacer />
      <span>&copy; {{ year }}</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      title: 'Talang Land',
      year: new Date().getFullYear(),
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
  padding: 0.5rem 2rem;
}
.title-link-anchor {
  display: flex;
  text-decoration: none;
}
.title-link {
  margin-left: 1rem;
  color: #fff !important;
  font-weight: bold;
}
@media only screen and (max-width: 410px) {
  .title-link {
    display: none;
  }
}
</style>
