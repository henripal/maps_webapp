<template>
  <b-navbar type="dark" variant="primary">
    <b-navbar-brand to="/">NewCo</b-navbar-brand>
      <b-navbar-nav>
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

export default {
  name: 'NavBar',
  data() {
    return {
    }
  },
  props: {
  },
  methods: {
  },
  computed: {
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
