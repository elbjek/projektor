<template>
	<div class="footer-wrapper">
			<div class="footer">
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
    
  }
}
</script>

<style lang="scss" scoped>
    .footer-wrapper{
        .footer{
            ul{
                margin:0;
                padding:0;
                background-color:pink;
                li{
                    list-style-type: none;
                    display: inline;
                    a{
                        padding:20px;
                        text-decoration: none;
                    }
                }
            }
        }
    }
</style>
