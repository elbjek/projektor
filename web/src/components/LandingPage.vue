<template>
	<div class="main-layout">
		<div class="">
			<div class="hero-wrapper">
				<div class="max-width">
					<div class="hero">
					<h1>Sve o EU fondovima na jednom mestu</h1>
					<p>Lorem ipsum, dolor sit amet consectetur adipisicing elit.</p>
					<div class="buttons" style="display:flex">
							<a class="btn orange" href="/">Button</a>
							<a class="btn white" href="/">Button 2</a>
					</div>
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
			display: flex;
			flex-direction: column;
			align-items: flex-start;
			justify-content: center;
			padding:150px 0px;
			h1{
				color:white;
				font-weight: 900;
				width: 50%;
   			text-align: left;
				font-size: 4em;
			}
			p{
				font-style: italic;
				color:white;
			}		   
		}
		.buttons{
			.btn:last-child{
				margin-left:15px;
			}
		}
		.block{
			display: flex;
			justify-content: center;
			align-items: flex-start;
			flex-direction: column;
			padding:100px 0px;
			&:nth-child(even){
				padding:0px;
				align-items:flex-end;
			}
		}
		.data-wrapper{
			display: flex;
		}
		.wrapper{
			background-color:white;
		}

		h3{
			font-weight: 900;
			font-size: 35px;
			padding:20px 0px;
		}
		.btn{
      padding:8px 24px;
			border-radius:20px;
			font-weight: bolder;
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
		.white{
			background-color:white;

		}
	}

</style>
