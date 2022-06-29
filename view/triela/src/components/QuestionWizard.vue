<template>
  <div class="righting-text">
    <button type="button" class="btn btn-danger">ウィザードの終了</button>
  </div>
  <h1>問題作成ウィザード</h1>
  <p>{{ CardMessage }}</p>
  <hr>
</template>

<script>
export default {
  name: "QuestionWizard.vue",
  props: {
    CardMessage: String,
  }
}
</script>

<style scoped>

</style>