<template>
	<div class="calls-wrapper">
        <div class="max-width">
            <h2> Pozivi</h2>
            <div class="calls">
                <div class="card" v-for="call in allCalls" :key="call.id">
                    <div class="card-body">
                        <h5 class="card-title">{{call.title}}</h5>
                        <div class="titles" style="display:flex; text-align:center; align-items:center; justify-content:space-between; width:50%; margin:0 auto;">
                         <h6 class="card-subtitle mb-2 text-muted" style="text-transform:capitalize"><span >Program:</span> {{call.program}} </h6>
                        <h6 class="card-subtitle mb-2 text-muted"> <span> Datum:</span>{{formatTime}}</h6>
                        </div>
                        <p class="card-text">{{call.description}}</p>
                        <a  v-bind:href="'/calls/'+ call.id" @click="storeId(call.id)" class="card-link">Jos</a>
                    </div>
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
          allCalls:{},
      }
  },
  mounted(){
    this.getCalls()
  },
  computed:{
      formatTime() {
        var moment = require('moment');
         moment().format()
         var formatTime = moment(this.allCalls.deadline).format('MMMM Do YYYY')
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
                {method: "GET", "url": "http://localhost:8081/calls/v1/calls",
                "headers": { "Authorization": 'Bearer ' + token},
                config
                })
                .then(result => {
                this.allCalls = result.data.calls;
                console.log(this.allCalls)
            }, error => {
                console.error(error);
            });;
      }, storeId(id) {
          var data = {id:id}
        
        //   this.$root.$emit('getid', data)
            var test = localStorage.setItem('id',id)
            console.log(localStorage.getItem('id'))
      }
     
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
        padding: 40px 40px 20px 20px;
       text-align: left;
       font-weight: 900;
   }
   .calls{
       display: grid;
       grid-template-columns: 50% 50%;
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
