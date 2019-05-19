<template>
	<div class="calls-wrapper">
        <div class="max-width">
           <div class="card">
                    <div class="card-body">
                        <h2>Poziv</h2>
                        <div class="img">
                            <img src="../assets/images/10.jpg" alt="">
                        </div>
                        <h5 class="card-title">{{singleCall.title}}</h5>
                        <div class="titles" style="display:flex; text-align:center; align-items:center; justify-content:space-between; width:50%; margin:0 auto;">
                         <h6 class="card-subtitle mb-2 text-muted" style="text-transform:capitalize"><span >Program:</span> {{singleCall.program}} </h6>
                        <h6 class="card-subtitle mb-2 text-muted"> <span> Datum:</span>{{formatTime}}</h6>
                        </div>
                        <p class="card-text" v-for="desc in singleCall.description" :key="desc.id">{{desc}}</p>
                        <div class="links">
                            <a href="/">Aplikacioni formular</a>
                            <a href="#">Pravila</a>
                            <a href="#">Smernice</a>
                        </div>
                        <div class="form">
                        <p class="addition"> Klikom na sledece dugme otvara ti se alat koji ce ti pomoci da procenis spremnost svog preduzeca da apliciras na ovaj konkurs,
                        pruzi ti smernice za unapredjenje i alternative za pribavljanje sredstava.</p>
                        <button ><a href="/form">Proceni svoju spremnost</a></button>
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
    display: flex;
   h2{
       text-align: left;
       font-weight: 900;
        padding: 20px 40px 20px 20px;
   }
    .card{
        text-align:left;
        .img{
            width:100%;
            height:500px;
            overflow: hidden;
            display: flex;
            justify-content: center;
            align-items: center;
            img{
                width:100%;
                height:auto;
            }
        }
        h5{
            font-size: 35px;
            margin:50px 0px;
        }
        .titles{
            width:100% !important;
            padding-bottom:30px;
        }
        .card-body{
            padding: 1.25rem 1.25rem 4rem 1.25rem;
        }
    }
    .links{
        margin:10px 10px 10px 0px;
        a{
            margin:20px;
            color:#3483ca;
            &:first-of-type{
                margin-left:0px;
            }
        }
    }
    .form{
        background-color: #f4f4f4;
        padding:20px;
        button{
            background-color: #8cc63f;
            border:none;
            color:white;
            font-weight: 900;
            padding:8px 24px;
            a{
                color:white;
            }
        }
    }
}

</style>
