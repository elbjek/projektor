<template>
	<div class="main-layout">
        <homenav></homenav>
        <div class="home-content-wrap max-width">
                <div class="latestnews">
                    <div class="card">
                        <h5 class="card-title">Prethodne pretrage</h5>
                        <div class="card-body">
                            <ul>
                                <li>
                                    <a href="#">Lorem ipsum</a>
                                </li>
                              <li>
                                    <a href="#">Lorem ipsum</a>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            <div class="center-content">
                <div class="card">
                    <div class="img">
                        <img src="../assets/images/11.jpg"  class="card-img-top" alt="...">
                    </div>
                <div class="card-body">
                    <h5 class="card-title">Nova budžetska rešenja – odgovor na izazove u prostoru slobode, bezbednosti i pravde?</h5>
                    <p class="card-text">Lorem ipsum dolor sit, amet consectetur adipisicing elit. Possimus, quis ut. Cum dolorum est nesciunt suscipit 
                        excepturi. Quam iure ipsam, error ullam animi molestiae culpa iusto vero tenetur illum fugiat?...</p>
                    <a href="#" class="btn btn-primary">Vidi jos</a>
                </div>
                </div>
                 <div class="card">
                    <div class="img">
                        <img src="../assets/images/12.jpg"  class="card-img-top" alt="...">
                    </div>
                <div class="card-body">
                    <h5 class="card-title">Poziv za konkurs</h5>
                    <p class="card-text">Lorem ipsum dolor sit, amet consectetur adipisicing elit. Possimus, quis ut. Cum dolorum est nesciunt suscipit 
                    excepturi. Quam iure ipsam, error ullam animi molestiae culpa iusto vero tenetur illum fugiat?...</p>
                    <a href="#" class="btn btn-primary">Vidi jos</a>
                </div>
                </div>
                 <div class="card">
                    <div class="img">
                        <img src="../assets/images/9.jpg"  class="card-img-top" alt="...">
                    </div>
                <div class="card-body">
                    <h5 class="card-title">EU Summit 2019</h5>
                    <p class="card-text">Lorem ipsum dolor sit, amet consectetur adipisicing elit. Possimus, quis ut. Cum dolorum est nesciunt suscipit 
                    excepturi. Quam iure ipsam, error ullam animi molestiae culpa iusto vero tenetur illum fugiat?...</p>
                    <a href="#" class="btn btn-primary">Vidi jos</a>
                </div>
                </div>
            </div>
        </div>
	</div>
</template>

<script>
import homenav from './partials/homeNav.vue'
export default {
  name: 'Home',
  components:{
      homenav
  }
}
</script>

<style lang="scss" scoped>

.main-layout{
    .max-width{
    width:1240px;
    margin:0 auto;
    }
    .home-content-wrap{
        display: grid;
        grid-template-columns: 20% auto;
    }
    .card{
        padding:20px;
        border:none;
    }
    .latestnews{
        .card{
            text-align:left;
            border:none;
            .card-body{
                padding:0px;
            ul{
                padding:0px;
                margin:0px;
                li{
                    list-style-type: none;
                    text-align: left;
                }
            }
            }
        }
    }
    .center-content{
        .card{
            padding:20px;
            text-align: left;
            .img{
                height: 250px;
                overflow: hidden;
                display: flex;
                justify-content: center;
                align-items: center;                
            }
        }
        .btn-primary{
            background-color:#3483ca;
            color:white;
            font-weight: 900;
            padding:8px 24px;
        }
    }
}

</style>
