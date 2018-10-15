<template>
  <div class="home">
    <HelloWorld msg="Welcome to Your BIDOUUU App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'home',
  components: {
    HelloWorld
  }, 
  mounted() {
    this.$store.dispatch("getData")
  }
}
</script>
