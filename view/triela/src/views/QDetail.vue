<template>
  <div class="wizard">
    <h1>問題：{{ main_title }}</h1>
    <p>問題の修正や削除ができます。</p>
    <hr>
    <h2>情報</h2>
    <hr>
    <div class="text-center">
      <p>タグ：<br>{{ tag }}</p>
      <p>問題タイプ：<br>{{type}}</p>
      <p>回答（または記述ポイント）：<br>{{answer}}</p>
      <p>得点：<br>{{score}}</p>
      <p>問題プレビュー：<br><button type="button" class="btn btn-warning" @click="preview">回答画面をプレビューする</button></p>
    </div>
    <div class="text-center margin-bottom">
      <button type="button" class="btn btn-primary width-90" @click="fix_q">問題の修正</button>
    </div>
    <div class="text-center margin-bottom">
      <button type="button" class="btn btn-danger width-90" @click="del_q">問題の削除</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "QDetail",
  data(){
    return {
      dataid: null,
      tag: '',
      type: '',
      answer: '',
      score: ''
    }
  },
  methods: {
    preview(){
      this.$router.push(`/preview/Q/${ this.$route.params.id }`)
    }
  },
  mounted: function() {
    axios.post(
        'http://localhost:8080/FindQuestion',
        {
          QId: this.$route.params.id
        }
    ).then(
        function(response){
          var exercise = response.data.DATA
          this.main_title = exercise.QName
          this.tag = exercise.QTags
          this.type = exercise.AnswerArea
          this.answer = exercise.QAnswer
          this.score = exercise.QScore
          this.dataid = this.$route.params.id
        }.bind(this)
    ).catch(
        error => console.log(error)
    )

    axios.get(
        'http://localhost:8080/GetAllQuestion',
    ).then(
        function(response){
          this.questions = response.data.DATA
        }.bind(this)
    ).catch(
        error => console.log(error)
    )
  }
}
</script>

<style scoped>
p{
  font-size: 20px;
}
</style>