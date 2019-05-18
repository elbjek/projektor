<template>
    <div class="form-wrapper">
      <form amethod="POST" action="/login">
      <input type="text" v-bind="input.username" placeholder="username" >
      <input type="text" v-bind="input.password" placeholder="password" >
      <!-- value="123qwe" -->
      <button type="submit" @:click="formSubmit()">Submit</button>
      </form>
      </div>
</template>

<script>

export default {
  name: 'UserLogin',
  components: {
    
  },
  data() {
    return{
        input:{
            username:'',
            password:''
        },
        response:''
    }
  },
  functions:{
            formSubmit() {  
                this.axios({ method: "POST", "url": "localhost:80/api/login", 
                "data": this.input, "headers": { "content-type": "application/json" } }).then(result => {
                    this.response = result.data;
                    // console.log(result)
                }, error => {
                    console.error(error);
                });
            }
  }
}
</script>

<style lang="scss" scoped>

	.form-wrapper{
		min-height:90vh;
		background-color:lightblue;
	}

</style>
