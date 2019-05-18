<template>
    <div class="form-wrapper">
      <form action="post">
              <input type="text" v-bind="username" placeholder="username" value="gliga">

      <input type="text" v-bind="password" placeholder="password" value="123qwe">
      <button type="submit" v-on:click="submitForm">Submit</button>
      </form>
      </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'UserLogin',
  components: {
    
  },
  data() {
    return{
      username:'',
      password:''
    }
  },
  functions:{
    submitForm(){
      console.log("hi")
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
