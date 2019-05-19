<template>
    <div class="navigation-wrapper">
       <div class="navigation max-width">
            <ul class="logo">
            <li><a href="/home"><img src="../assets/logo.png"></a></li>
        </ul>
        <ul class="right">
            <li><a class="mail" href=""><img src="../assets/icons/mail.svg" alt=""></a></li>
            <li><a class="username"  v-bind:href="'/user/'+ userData.ID"><img src="../assets/icons/user.svg" alt="">{{userData.firstname}} {{userData.lastname}}</a></li>
            <li><a class="logout" href="" @click="logout()">Odjavi se</a></li>            
        </ul>
       </div> 
    </div>
</template>
<script>

export default {
  name: 'loggednav',
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
            }, error => {
                console.error(error);
            });;
      },
      logout(){
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
    background:#048BA8;
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
                font-weight: 900;
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
        .username{
            img{
                margin-right:20px;
            }
        }
        .logout{
        }
        .username,.mail{
            margin:20px;
        }
        ul{
            li{
                margin:0px 20px;
                a{
                    text-decoration: none;
                    margin:20px;
                    font-size:14px;
                    font-weight: 900;
                    color:white
                }
            }
        }
    }
}

</style>
