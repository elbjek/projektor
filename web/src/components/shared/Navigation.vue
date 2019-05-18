<template>
    <div class="navigation-wrapper">
       <div class="navigation max-width">
            <ul class="logo">
            <li><a href="/">Projektor</a></li>
        </ul>
        <ul class="right">
            <li><a class="login" href="/login">Uloguj se</a></li>
            <li><a class="reg" href="">Registruj se</a></li>            
        </ul>
       </div> 
    </div>
</template>
<script>


export default {
  name: 'navigation',
  components: {
    
  },
    data(){
        return {
            
        }
    },
    functions:{
        // mounted () {
        // this.axios
        // .get('https://api.coindesk.com/v1/bpi/currentprice.json')
        // .then(response => (this.info = response))
        // }
    }
}
</script>

<style lang="scss" scoped>
.navigation-wrapper{
    background:#999999;
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
                color:#ffffff;
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
                }
            }
        }
        .login{
            border-color:white;
        }
        .login, .reg{
            padding:8px 24px;
            border:3px solid #f2a755;
            border-radius:20px;
        }
        .reg{
            background:#f2a755;
            color:white;
            margin-left:20px;
        }
    }
}
.max-width{
    width:1420px;
    margin:0 auto;
}
</style>
