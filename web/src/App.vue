<template>
  <div id="app">
    <navigation v-if="!hasToken"></navigation>
    <loggednav v-if="hasToken"></loggednav>
    <router-view></router-view>
    <MyFooter></MyFooter>
  </div>
</template>

<script>
import navigation from'./components/shared/Navigation.vue'
import MyFooter from './components/shared/Footer.vue'
import loggednav from './components/LoggedNav.vue'
    /*es-lint-disable*/

export default {
  name: 'app',
  components: {
   navigation,
   MyFooter,
   loggednav
  },
  data(){
    return{
      token:''
    }
  },
  mounted(){
    this.getUser()
  },
  computed:{
    hasToken(){
      var div = document.getElementById('navigationbar')
      if(localStorage.getItem('token')){
        this.token - localStorage.getItem('token');
        return true;
      } else{
        return false
      };
    }
  },
  methods:{
      getUser(){
            var config = {
                 headers: {'Access-Control-Allow-Origin': '*'}
            };
            var token = localStorage.getItem('token');
            var userId = this.$router.currentRoute.path.slice(6);
            this.axios(
                {method: "GET", "url": "http://localhost:8081/users/v1/user/",
                "headers": { "Authorization": 'Bearer ' + token},
                config
                })
                .then(result => {
                this.userData = result.data;
            }, error => {
                console.error(error);
            });;
      },
  }
}
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}
</style>
