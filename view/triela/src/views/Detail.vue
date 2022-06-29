<template>
  <QuestionWizard CardMessage="以下の内容を埋めてください。"/>
  <p>主題</p>
  <input type="text" name="text" class="form-control margin-bottom" placeholder="例）応用技術者認定試験 午前">
  <p>詳細情報</p>
  <input type="text" name="text" class="form-control margin-bottom" placeholder="例）2018(平成30年)の秋に開催された午前問題です。">
  <p>主タグ</p>
  <input type="text" name="text" class="form-control margin-bottom" placeholder="例）応用技術者 応用 技術者 午前">
  <p>制限時間(時間：分：秒)</p>
  <input type="time" step="1" class="form-control centering-text margin-bottom" value="01:00:00">
  <div class="righting-text">
    <button type="button" class="btn btn-primary" v-on:click="next">確認と作成</button>
  </div>
</template>

<script>
import QuestionWizard from "@/components/QuestionWizard";

export default {
  name: "DetailVue",
  components: {
    QuestionWizard
  },
  methods: {
    next() {
      this.$router.push('/exwizard/check')
    }
  }
}
</script>

<style scoped>

</style>