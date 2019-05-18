<template>
  <div id="app">
    <navigation></navigation>
    <router-view></router-view>
    <MyFooter></MyFooter>
  </div>
</template>

<script>
import navigation from'./components/shared/Navigation.vue'
import MyFooter from './components/shared/Footer.vue'
    /*es-lint-disable*/

export default {
  name: 'app',
  components: {
   navigation,
   MyFooter
  }
}
</script>

<style>
#app {
  /* font-family: 'Avenir', Helvetica, Arial, sans-serif; */
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}
</style>
