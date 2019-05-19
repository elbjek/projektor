<template>
	<div class="bot-wrap">
        <div class="bot">
            <img src="../assets/icons/chat.svg" alt="">
        </div>
    </div>

</template>

<script>
export default {
  name: 'LandingPage'
}
</script>

<style lang="scss" scoped>
.bot{
background-color: #8cc63f;
height:80px;
width:80px;
padding:20px;
border-radius:50px;
position: fixed;
bottom:20px;
right:20px;
&:hover{
    transform:scale(1.1);
    cursor: pointer;
}
}

</style>
