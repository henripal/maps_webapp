<template>

<b-jumbotron header="NewCo" lead="Making the world a better place through constructing elegant hierarchies for maximum code reuse and extensibility." >
  <p>Disrupt digital media with us:</p>
          <ModalSignup buttoncolor="primary"/>
</b-jumbotron>
</template>

<script>
import ModalSignup from '../components/ModalSignup.vue'

export default {
  name: 'home',
  components: {
    ModalSignup
  }, 
  mounted() {
  }
}
</script>

<style>
.jumbotron {
  height: 100vh !important; 
}
</style>