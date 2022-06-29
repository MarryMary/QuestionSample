<template>
  <h1>Create successful!</h1>
  <h2>作成が完了しました。</h2>
  <hr>
  <div class="centering-text">
    <p>エクササイズが正常に作成されました。<br>
      引き続き問題管理画面から問題を作成してください。
    </p>
    <router-link to="/exwizard/start">問題作成ウィザードの開始</router-link>
  </div>
</template>

<script>
export default {
  name: "SuccessFinished.vue"
}
</script>

<style scoped>

</style>