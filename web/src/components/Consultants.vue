<template>
    <div class="main-layout">
       <homenav></homenav>
       <img src="../assets/images/konsultanti_mock.png" width="90%" alt="">
      </div>
</template>

<script>
import homenav from "./partials/homeNav.vue";
export default {
  name: 'UserLogin',
  components: {
    homenav,
  }
}
</script>

<style lang="scss" scoped>

.max-width{
    width:1240px;
    margin:0 auto;
}
</style>
