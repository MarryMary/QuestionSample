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
      <button type="button" class="btn btn-danger width-90" @click="openModal">問題の削除</button>
    </div>
    <div id="overlay" v-show="showContent">
      <div id="content">
        <button type="button" class="btn btn-primary" @click="closeModal">X</button>
        <div class="text-center margin-top">
          <h1>問題を削除しますか？</h1>
          <p>
            問題を削除すると復旧できません。<br>
            問題を訂正する場合は修正ボタンで修正して下さい。<br>
            問題を完全に削除する場合は続けて下さい。
          </p>
          <div class="text-center margin-bottom">
            <button type="button" class="btn btn-danger width-90" @click="closeModal">問題の削除</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "QDetail",
  data(){
    return {
      showContent: false,
      dataid: null,
      tag: '',
      type: '',
      answer: '',
      score: ''
    }
  },
  methods: {
    preview(){
      this.$router.push(`/preview/Q/${ this.$route.params.ExName }/${ this.$route.params.ExType }/${ this.$route.params.QName }`)
    },

    openModal: function(){
      this.showContent = true
    },

    closeModal: function(){
      this.showContent = false
    }
  },
  mounted: function() {
    axios.post(
        'http://localhost:8080/FindQuestion',
        {
          ExName: this.$route.params.ExName,
          ExType: this.$route.params.ExType,
          QName: this.$route.params.QName
        },
        {
          withCredentials: true
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

#content{
  z-index:2;
  width:50%;
  padding: 1em;
  background:#fff;
}

#overlay{
  z-index:1;
  position:fixed;
  top:0;
  left:0;
  width:100%;
  height:100%;
  background-color:rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;

}
</style>