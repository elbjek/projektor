<template>
    <div class="form-wrapper">
        <div class="form">
            <div class="form-header">
                <h2>Uloguj se</h2>
            </div>
            <form method="POST" action="/login"  v-on:submit.prevent="formSubmit">
                <input type="text" v-model="fields.username" value="gliga" placeholder="username" >
                <input type="password" v-model="fields.password" value="123qwe" placeholder="password" >
                <button class="btn">Submit</button>
            </form>
        </div>
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

	.form-wrapper{
		min-height:90vh;
		background-color:#f4f4f4;
        display:flex;
        align-items:center;
        justify-content: center;
        .form{
            display: -webkit-box;
            display: -ms-flexbox;
            display: flex;
            -webkit-box-orient: vertical;
            -webkit-box-direction: normal;
            -ms-flex-direction: column;
            flex-direction: column;
            background-color: white;
            padding: 40px;
            width: 40%;
            height: 500px;
            align-items: center;
            justify-content: center;
            .form-header{
                padding:20px;
                h2{
                font-weight: 900;                   
                }
            }
            form{
                display: flex;
                flex-direction: column;
                margin:0 auto;
                input{
                    margin:30px 0px;
                    border:none;
                    border-bottom:2px solid #2d2d2d;
                    padding:5px 5px 5px 15px;
                    &:hover{
                        border-bottom: 2px solid #f2a755;
                        transition: 0.4s ease-in-out;
                    }
                }
                button{
                    background-color: #3483ca;
                    border: none;
                    padding: 10px;
                    border-radius: 8px;
                    font-weight: 900;
                    margin-top: 30px;
                    width: 50%;
                    margin: 0 auto;
                    color: white;
                    &:focus{
                        outline: none;
                        border:none;
                    }
                }
            }
        }
	}

</style>
