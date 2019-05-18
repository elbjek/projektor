<template>
	<div class="main-layout">
		<div class="">
			<div class="hero-wrapper">
				<div class="max-width">
					<div class="hero">
					<h1>Sve o EU fondovima na jednom mestu</h1>
					<p>Lorem ipsum, dolor sit amet consectetur adipisicing elit.</p>
					<a class="btn" href="/">Button</a>
					<a class="btn" href="/">Button 2</a>
					</div>
				</div>
			</div>
			<div class="data-wrapper">
				<div class="wrapper max-width">
				<div class="block">	
					<h3>Pretrazite bazu aktuelnih EU projekata.</h3>
					<a class="btn orange" href="#">Button</a>
				</div>
				<div class="block">
					<h3>Pretrazite bazu aktuelnih EU projekata.</h3>
					<a class="btn blue" href="#">Button</a>
				</div>
				<div class="block">
					<h3>Pretrazite bazu aktuelnih EU projekata.</h3>
					<a class="btn purple" href="#">Button</a>
				</div>
				</div>
				
			</div>
		</div>
	</div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'LandingPage',
  components: {
    
  },
  functions:{
  }
}
</script>

<style lang="scss" scoped>

	.main-layout{
		.max-width{
			width:1420px;
			margin:0 auto;
		}
		.hero-wrapper{
			background-color:#999999;
		}
		.hero{
			height:400px;
			display: flex;
			flex-direction: column;
			align-items: flex-start;
			justify-content: center;
			h1{
				color:white;
				font-weight: 900;
			}
			p{
				font-style: italic;
				color:white;
			}		   
		}
		.block{
			display: flex;
			justify-content: center;
			align-items: flex-start;
			min-height:300px;
			flex-direction: column;
			padding:20px;
			&:nth-child(even){
				align-items:flex-end;
			}
		}
		.wrapper{
			background-color:white;
		}

		h3{
			font-weight: 900;
			font-size: 35px;
			padding:20px;
		}
		.btn{
            padding:8px 14px;
			border-radius:20px;
		}
		.orange{
			background-color:#f2a755;
		}
		.purple{
			background-color:#ef84e0;
		}	
		.blue{
			background-color:lightblue;
		}
	}

</style>
