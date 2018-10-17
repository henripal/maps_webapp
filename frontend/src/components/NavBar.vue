<template>
  <b-navbar type="dark" variant="primary" toggleable="md">
    <b-navbar-brand to="/">NewCo</b-navbar-brand>
      <b-navbar-nav>
        <template v-if="!loggedIn">
          <ModalSignup/>
          <ModalSignin/>
        </template>
        <template v-else>
          <UserMenu/>
        </template>

        <b-nav-item to="/about">About</b-nav-item>

    </b-navbar-nav>
  </b-navbar>
  <!-- <div class="navbar">
    <p>this is a navbar</p>
    <router-link to="/">Home</router-link>
    <router-link to="/about">About</router-link>
  </div> -->
</template>

<script>
import ModalSignup from '../components/ModalSignup.vue'
import ModalSignin from '../components/ModalSignin.vue'
import UserMenu from '../components/UserMenu.vue'

export default {
  name: 'NavBar',
  data() {
    return {
    }
  },
  components: {
    ModalSignup,
    ModalSignin,
    UserMenu,
  },
  props: {
  },
  methods: {
  },
  computed: {
    loggedIn () {
      return this.$store.getters.loggedIn
    },
  },
  mounted () {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.title-area {
  float: left !important;
}
.menu-items {
  float: right !important;
}
</style>
