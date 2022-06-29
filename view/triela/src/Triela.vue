<template>
  <div class="bg">
    <div class="centering-box">
      <div class="glass">
        <transition name="slide">
          <div>
            <router-view></router-view>
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "triela-vue"
}
</script>

<style>
  @import url('https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.css');
  @import url('https://fonts.googleapis.com/css2?family=Zen+Kaku+Gothic+New:wght@300&display=swap');

  * {
    padding: 0;
    margin: 0;
    box-sizing: border-box;
  }

  html,
  body,
  div.background {
    height: 100%;
  }

  body {
    min-height: 100%;
    background: lightblue !important;
    font-family: 'Zen Kaku Gothic New', sans-serif !important;
  }

  .centering-box{
    margin: auto;
  }

  .glass {
    width: 100%;
    max-width: 90%;
    padding: 30px;
    background-color: white;
    box-shadow: 0 0 8px gray;
    border-radius: 15px;
    margin: auto;
    margin-top: 2%;
  }

  .glass h1 {
    text-align: center;
    font-size: 50px;
    font-weight: bold;
    letter-spacing: 0.06em;
  }

  .glass h2 {
    text-align: center;
    font-size: 40px;
    font-weight: bold;
    letter-spacing: 0.06em;
  }

  .glass p {
    margin-top: 20px;
    font-size: 20px;
  }

  .margin-bottom{
    margin-bottom: 30px;
  }

  .maintext{
    resize: none;
    width: 100%;
    height: 80vh;
  }

  .form-chain{
    background: white;
  }

  .centering-text{
    text-align: center;
  }

  .width-90{
    width: 90%;
  }

  .selector{
    text-decoration: none;
    color: black;
  }

  .select{
    border-radius: 30px;
    border: 1px solid black;
    padding: 10px;
  }

  .selector:hover{
    text-decoration: none;
    color: black;
  }

  .select:hover{
    cursor: pointer;
    opacity: 0.5;
  }

  .righting-text{
    text-align: right;
  }

  .right-enter-active, .right-leave-active {
    transform: translate(0px, 0px);
    transition: transform 225ms cubic-bezier(0, 0, 0.2, 1) 0ms;
  }

  .right-enter, .right-leave-to {
    transform: translateX(100vw) translateX(0px);
  }
</style>