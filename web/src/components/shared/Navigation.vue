<template>
    <div class="navigation-wrapper">
       <div class="navigation max-width">
            <ul class="logo">
            <li><a href="/">Projektor</a></li>
        </ul>
        <ul class="right">
            <li><a class="login" href="/login">Uloguj se</a></li>
            <!-- <li><a class="login" v-bind="show" @click="openModal()" href="#">Uloguj se</a></li> -->
            <login v-if="show"></login>
            <li><a class="reg" href="">Registruj se</a></li>            
        </ul>
       </div> 
    </div>
</template>
<script>
import login from '../Login.vue'

export default {
  name: 'navigation',
  components: {
    login
  },
    data(){
        return {
            show:false
        }
    }
}
</script>

<style lang="scss" scoped>
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
        ul{
            li{
                a{
                    text-decoration: none;
                    margin:20px;
                    font-size:14px;
                    font-weight: 900;
                    color:white
                }
            }
        }
        .login{
            border-color:white;
        }
        .login, .reg{
            padding:8px 24px;
            border:3px solid white;
            border-radius:20px;
        }
        .reg{
            background:white;
            margin-left:20px;
            color:#384b60;
        }
    }
}

</style>
