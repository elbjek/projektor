<template>
    <div class="form-wrapper">
        <div class="form">
          <h4 class="card-title">Proveri svoju spremnost:</h4>
          <form class="chbx" v-for="item in data" :key="item.id">
            <label for="">{{item.description}}</label>
              <input type="checkbox" :checked="item.checked" >
              <p v-if="item.checked">{{item.assist_text}}</p>
          </form>           
      </div>
    </div>
</template>

<script>

export default {
  name: 'Form',
  components: {
    
  },
  data() {
    return{
      data:''
    }
  },
  mounted(){
    this.formSubmit()
  },
  methods:{
    formSubmit(event) {
            var config = {
                 headers: {'Access-Control-Allow-Origin': '*'}
            };
                        var token = localStorage.getItem('token');
            this.axios({ method: "GET", "url": "http://localhost:8081/calls/v1/call/ceb1f316-52a4-40a5-bf3e-8d0f5e3b0d52",
             "headers": { "Authorization": 'Bearer ' + token}, config}).then(result => {
                this.data = result.data.checks;
                console.log(result.data.checks)
            }, error => {
                console.error(error);
            });
        }
  },
  isChecked(){
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
       
	}

</style>
