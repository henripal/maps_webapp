<template>
  <b-navbar type="dark" variant="primary" toggleable="md">
    <b-navbar-brand to="/">
     <img
     src="https://cdn4.iconfinder.com/data/icons/glyphlibrary-one/100/satellite-gps-512.png"
     class="d-inline-block align-top"
     alt="BV"
     height="30"
     >
     <img
     src="https://cdn0.iconfinder.com/data/icons/pixel-perfect-at-24px-volume-6/24/2175-512.png"
     class="d-inline-block align-top"
     alt="BV"
     height="30"
     >
    </b-navbar-brand>
      <b-navbar-nav class="ml-auto">


    </b-navbar-nav>
  </b-navbar>
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
