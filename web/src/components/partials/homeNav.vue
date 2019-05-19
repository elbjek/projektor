<template>
    <div class="navigation-wrapper">
       <div class="navigation max-width">
        <ul class="right">
            <li><a class="mail first" href="">Pocetna</a></li>
            <li><a class="mail" href="#">Konsultanti</a></li>
            <li><a class="mail" href="#">Dogadjaji</a></li>
            <li><a class="mail" href="#">Znanje</a></li>
            <li><a class="mail" href="/calls">Pozivi</a></li>
        </ul>
       </div> 
    </div>
</template>
<script>

export default {
  name: 'homenav',
    data(){
        return {
            userData:''
        }
    },
    mounted(){
        this.getUser()
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
                console.log(this.userData)
            }, error => {
                console.error(error);
            });;
      },
      logout(){
          console.log()
          localStorage.removeItem('token');
          this.$router.push('/');
      }
    }
}
</script>

<style lang="scss" scoped>
.max-width{
    width:1240px;
    margin:0 auto;
}
.max-width{
    width:1240px;
    margin:0 auto;
}
.navigation-wrapper{
    background:white;
    display:flex;
    padding:20px;
}
.navigation{
    display:flex;
    justify-content: space-between;
    img{
        height:20px;
    }
    ul{
        padding:0px;
        margin:0px;
        li{
            list-style-type: none;
            a{
                color:white;
                &:hover{
                    border-bottom:none;
                    text-decoration: none;
                }
            }
        }
    }
    .logo{
        a{
            font-weight:900;
            font-size:18px;
        }
    }
    .right{
        display: flex;
        color:#048BA8;
        .username{
            img{
                margin-right:20px;
            }
        }
        .first{
            padding:20px;
            border-bottom:2px solid #048BA8;
        }
        .username,.mail{
            margin:20px;
            color:#048BA8;
            &:hover{
                font-weight: 900;
                transition: 1s ease-in-out;
                border-bottom:2px solid #048BA8;
            }
        }
        ul{
            li{
                margin:0px 20px;
                a{
                    text-decoration: none;
                    margin:20px;
                    font-size:14px;
                    font-weight: 900;
                    color:#048BA8
                }
            }
        }
    }
}

</style>
