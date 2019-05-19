<template>
    <div class="main-layout">
       <img src="" alt="">
      </div>
</template>

<script>

export default {
  name: 'UserLogin',
  components: {
    
  },
  data() {
    return{
        fields:{},
        username:'',
        password:'',
        response:'',
        message:''
    }
  },
  methods:{
    formSubmit(event) {
            var config = {
                 headers: {'Access-Control-Allow-Origin': '*'}
            };
            this.axios({ method: "POST", "url": "http://localhost:8081/login", 
            "data": this.fields, "headers": { "content-type": "application/json" }, config}).then(result => {
                this.response = result;
                localStorage.setItem(
                    result.data.name,
                    result.data.jwt
                )
                this.$router.push('/home')
            }, error => {
                console.error(error);
            });
        }
  }
}
</script>

<style lang="scss" scoped>

.max-width{
    width:1240px;
    margin:0 auto;
}
</style>
