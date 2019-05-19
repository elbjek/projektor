<template>
	<div class="calls-wrapper">
        <div class="max-width">
           <h2>Poziv</h2>
           <div class="card">
                    <div class="card-body">
                        <h5 class="card-title">{{singleCall.title}}</h5>
                        <div class="titles" style="display:flex; text-align:center; align-items:center; justify-content:space-between; width:50%; margin:0 auto;">
                         <h6 class="card-subtitle mb-2 text-muted" style="text-transform:capitalize"><span >Program:</span> {{singleCall.program}} </h6>
                        <h6 class="card-subtitle mb-2 text-muted"> <span> Datum:</span>{{formatTime}}</h6>
                        </div>
                        <p class="card-text">{{singleCall.description}}</p>
                    </div>
                </div>     
        </div>
        
	</div>
</template>

<script>

export default {
  name: 'Profile',
  components:{
  },
  data(){
      return{
          singleCall:{},
          callid:''
      }
  },
  beforeMount(){
      this.callid= localStorage.getItem('id');
    console.log(this.callid)
  },
   mounted(){
      this.getCalls()
  },
    computed:{
      formatTime() {
        var moment = require('moment');
         moment().format()
         var formatTime = moment(this.singleCall.deadline).format('MMMM Do YYYY')
         return formatTime
      }
  },
  computed:{
      formatTime() {
        var moment = require('moment');
         moment().format()
         var formatTime = moment(this.singleCall.deadline).format('MMMM Do YYYY')
         return formatTime
      }
  },
  methods:{
      getCalls(){
            var config = {
                 headers: {'Access-Control-Allow-Origin': '*'}
            };
            var token = localStorage.getItem('token');
            this.axios(
                {method: "GET", "url": "http://localhost:8081/calls/v1/call/" + this.callid,
                "headers": { "Authorization": 'Bearer ' + token},
                config
                })
                .then(result => {
                this.singleCall = result.data;
                console.log(this.singleCall)
            }, error => {
                console.error(error);
            });;
      },
     
  }
}

</script>

<style lang="scss" scoped>
.max-width{
    width:1240px;
    margin:0 auto;
}
.calls-wrapper{
    min-height:100vh;
    background-color:#f4f4f4;
   h2{
       padding:80px 80px 40px 0px;
       text-align: left;
       font-weight: 900;
   }
   .calls{
       display: grid;
       grid-template-columns: 100%;
       grid-gap: 20px;
        p{
            display: -webkit-box;
            -webkit-line-clamp: 3;
            -webkit-box-orient: vertical; 
            overflow: hidden;
        }
   }
}

</style>
