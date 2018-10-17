
<template>
    <b-nav-item>{{ firstName  }}</b-nav-item>
</template>

<script>
export default {
  name: 'UserMenu',
  data() {
    return {
    }
  },
  methods: {
  },
  computed: {
    firstName () {
      return this.$store.getters.firstName
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
