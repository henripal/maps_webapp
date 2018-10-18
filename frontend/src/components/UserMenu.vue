
<template>
 <b-nav-item-dropdown :text="firstName">
    <b-dropdown-item @click="logOut">Logout</b-dropdown-item>
    </b-nav-item-dropdown>
</template>

<script>
export default {
  name: 'UserMenu',
  data() {
    return {
    }
  },
  methods: {
      logOut () {
          this.$store.dispatch("logOut")
      }
  },
  computed: {
    firstName () {
      return this.$store.getters.firstName
    }
  },
}
</script>

<style scoped>
</style>
