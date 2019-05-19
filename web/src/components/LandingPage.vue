<template>
	<div class="main-layout">
		<div class="">
			<div class="hero-wrapper">
				<div class="max-width">
					<div class="hero">
					<h1>Lorem ipsum dolor sit amet.</h1>
					<p>Lorem ipsum, dolor sit amet consectetur adipisicing elit.</p>
					<div class="buttons" style="display:flex">
							<a class="btn white" href="/">Registruj se</a>
					</div>
					</div>
				</div>
			</div>
			<div class="data-wrapper">
				<h2>Šta nudimo?</h2>
				<div class="wrapper max-width">
					<div class="card" style="width: 18rem;">
					<div class="card-body">
						<div class="img">
							<img src="../assets/icons/customer-service.svg" alt="">
						</div>
						<h5 class="card-title">Lorem Ipsum</h5>
							<p class="card-text"> Lorem ipsum, dolor sit amet consectetur adipisicing elit.</p>				
						<a href="#" class="card-link">Another link</a>
					</div>
					</div>
					<div class="card" style="width: 18rem;">
					<div class="card-body">
						<div class="img">
							<img src="../assets/icons/database.svg" alt="">
						</div>
						<h5 class="card-title">Lorem Ipsum</h5>
							<p class="card-text">Lorem ipsum, dolor sit amet consectetur adipisicing elit.</p>

						<a href="#" class="card-link">Another link</a>
					</div>
					</div>
					<div class="card" style="width: 18rem;">
					<div class="card-body">
						<div class="img">
							<img src="../assets/icons/database.svg" alt="">
						</div>					
						<h5 class="card-title">Lorem Ipsum</h5>
						<p class="card-text">Lorem ipsum, dolor sit amet consectetur adipisicing elit.</p>
						<a href="#" class="card-link">Another link</a>
					</div>
					</div>
				</div>
			</div>
				<div class="card person">
					<div class="bg-image">
						<h3 class="card-title">Jovan Popovic iz Kraljeva</h3>
						<p class="card-text">Projektor mi je pomogao da dodjem do sredstava iz IPARD fonda i unapredim proizvodnju.</p>
						<button class="btn">Registruj se</button>
					</div>
				</div>
		</div>
	</div>
</template>

<script>
export default {
  name: 'LandingPage'
}
</script>

<style lang="scss" scoped>

	.main-layout{
		.max-width{
			width:1240px;
			margin:0 auto;
		}
		.hero-wrapper{
			background-color:#048BA8;
		}
		.hero{
			display: flex;
			flex-direction: column;
			align-items: flex-start;
			justify-content: center;
			padding:150px 10px;
			@media screen and (max-width:1240px){
				padding:150px 40px;
			}
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
			.white{
				border:2px solid white;
				color:white;
			}   
		}
		.buttons{
			.btn:last-child{
				margin-left:15px;
			}
		}
		.block{
			padding: 100px 0px;
			background-color: white;
			border-radius: 20px;
		}
		}
		.data-wrapper{
			background-color: #f4f4f4;
			min-height: 600px;
			display: flex;
			padding: 80px 0px 100px 0px;
			flex-direction: column;
			justify-content: center;
			align-items: center;
				.wrapper{
					display: grid;
					justify-content: center;
					grid-template-columns: 33% 33% 33%;
					width:100%;
					@media screen and (max-width:630px){
						grid-template-columns: 1fr;
						justify-content: center;
						align-items: center;
					}
				}
				.card{
					border:none;
					border-radius:8px;
					justify-self: center;
				}
		.img{
			width:40px;
			text-align: center;
			margin:0 auto;
			// padding:20px;
		}
		}
		h5{
			margin:10px 0px;
		}	
		h2{
			font-weight:900;
			max-width:1440px;
			margin:0 auto;
			padding:20px 0px 40px 0px;
		}

		.person{
			height:500px;
			background-image:url('../assets/images/background.png');
			background-position: center;
			background-size: cover;
			border:none;
			.bg-image{
				display: flex;
				flex-direction: column;
				width: 50%;
				justify-self: flex-end;
				align-self: flex-end;
				justify-content: center;
				align-items: flex-start;
				height: 100%;
				padding-right:40px;
				h3{
					font-size:55px;
					font-weight: 900;
					color:white;
				}
				p{
					color: white;
					margin-bottom: 1em;
					font-size: 19px;
					word-break: unset;
					word-wrap: break-word;
					text-align: left;
					font-style: italic;
				}
				.btn{
					background-color:white;
					padding:8px 23px;
				}
			}
		}
		.news-wrap{
			display: grid;
			grid-template-columns: 33% 33% 33%;
			grid-row-gap: 50px;
			padding: 100px 0px;
			@media screen and (max-width){
				width:100%;
				grid-template-columns: 1fr;
			}
			.card{
				justify-self: center;
				border:none;
				width:200%;
				// background-color: #f4f4f4;
    			padding: 20px;
				border-radius: 8px;
				width:24em !important;
				img{
					min-height:200px;
				}
				.img{
					height: 200px;
					overflow: hidden;
					display: flex;
					align-items: center;
				}
				.card-body{
					// display: flex;
					// justify-content: center;
					padding:0px;
					.card-title{
						margin:0px;
						padding:10px;
						font-weight: 900;
						background-color: #f4f4f4;
					}
				}
			}
		}

</style>
