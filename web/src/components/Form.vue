<template>
    <div class="form-wrapper">
        <div class="form">
                <div class="card">
                    <div class="card-body">
                        <h4 class="card-title">Oblasti u kojima vam treba dodatno obrazovanje i obuka: </h4>
                        <div class="chbx" v-for="item in data" :key="item.id">
                          <label for="checkbox"><input type="checkbox" id="input" v-model="item.checked"> <span>{{item.description}}</span></label>

                            <p v-if="isChecked" class="assist">{{item.assist_text}}</p>
                        </div>
                            <button class="btn" href="#" v-on:click="check">Submit </button>
                            <a @click="alert" href="#">Posalji izvestaj na mail!.</a>
                    </div>
                </div>                
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
      data:'',
      isChecked:false
    }
  },
  mounted(){
    this.getData()
  },
  computed:{

  },
  methods:{
    alert(){
      alert("poslato!")
    },
    check(){
      this.data.forEach(element => {
        if(element.checked == false){
         return this.isChecked = true;
        }
    });

    },
    getData(event) {
            var config = {
                 headers: {'Access-Control-Allow-Origin': '*'}
            };
           var token = localStorage.getItem('token');
            this.axios({ method: "GET", "url": "http://localhost:8081/calls/v1/call/ceb1f316-52a4-40a5-bf3e-8d0f5e3b0d52",
             "headers": { "Authorization": 'Bearer ' + token}, config}).then(result => {
                this.data = result.data.checks;
                // console.log(result.data.checks)
            }, error => {
                console.error(error);
            });
        }
  }
}
</script>

<style lang="scss" scoped>

.btn{
  padding:8px 24px;
  background-color: #3483ca;
  color:white;
  font-weight: bolder;
  margin:10px;
}

    .assist{
     color:red;
     background-color: #f4f4f4;
     padding:10px;
      // display: none;
      &:empty{
        display: none;
      }
    }
    .show{
      display:block;
    }
	.form-wrapper{
		min-height:90vh;
		background-color:#f4f4f4;
    display:flex;
    align-items:center;
    justify-content: center;
  .form{
    width:1240px;
    margin:0 auto;
    .chbx{
      text-align: left;
    }
    label{
    text-align:left !important;
    font-size: 14px;
    span{
      margin:0px 20px;
    }
    }
  }
	}

</style>
