<template>
  <QuestionWizard CardMessage="どちらの方法で問題作成を行いますか？"/>
  <div class="form-check">
    <input class="form-check-input" type="radio" name="flexRadioDefault" id="flexRadioDefault1">
    <label class="form-check-label" for="flexRadioDefault1">
      バンドル方式（JSON形式ファイルのアップロード）
    </label>
  </div>
  <div class="form-check">
    <input class="form-check-input" type="radio" name="flexRadioDefault" id="flexRadioDefault2" checked>
    <label class="form-check-label" for="flexRadioDefault2">
      インタラクティブ方式（対話式作成）
    </label>
  </div>
  <div class="righting-text">
    <button type="button" class="btn btn-primary" v-on:click="next">次へ進む</button>
  </div>
</template>

<script>
import QuestionWizard from "@/components/QuestionWizard";

export default{
  name: 'WhichMode',
  components: {
    QuestionWizard
  },
  methods: {
    next(){
      this.$router.push('/qwizard/detail')
    }
  }
}
</script>