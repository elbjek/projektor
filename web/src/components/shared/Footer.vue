<template>
	<div class="footer-wrapper">
			<div class="footer max-width">
                <ul>
                    <li><a href="">Uslovi Koriscenja</a></li>
                    <li><a href="">Nesto drugo</a></li>                    
                </ul>
            </div>
	</div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'footer',
  components: {
    
  },

}
</script>

<style lang="scss" scoped>
    .footer-wrapper{
        background-color:#2d2d2d;
        .footer{
            padding:20px;
            ul{
                margin:0;
                padding:0;
                li{
                    list-style-type: none;
                    display: inline;
                    a{
                        padding:20px;
                        text-decoration: none;
                        color:white;
                        &:hover{
                            text-decoration: underline;
                        }
                    }
                }
            }
        }
    }
</style>
