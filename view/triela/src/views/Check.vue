<template>
  <QuestionWizard CardMessage="以下の内容で登録します。"/>
  <p>主題</p>
  <input type="text" name="text" class="form-control margin-bottom" value="応用技術者認定試験 午前" disabled>
  <p>詳細情報</p>
  <input type="text" name="text" class="form-control margin-bottom" value="2018(平成30年)の秋に開催された午前問題です。" disabled>
  <p>主タグ</p>
  <input type="text" name="text" class="form-control margin-bottom" value="応用技術者 応用 技術者 午前" disabled>
  <p>制限時間</p>
  <input type="time" step="1" class="form-control centering-text margin-bottom" value="01:00:00" disabled>
  <div class="righting-text">
    <button type="button" class="btn btn-primary" v-on:click="next">作成して完了する</button>
  </div>
</template>

<script>
import QuestionWizard from "@/components/QuestionWizard";

export default {
  name: "CheckDetail",
  components: {
    QuestionWizard
  },
  methods: {
    next() {
      this.$router.push('/exwizard/SuccessFinished')
    }
  }
}
</script>

<style scoped>

</style>