<template>
  <h1>問題画面プレビューモード</h1>
  <h2>本番ではサイドバーやこのテキストは表示されません。</h2>
  <div class="wizard">
    <div class="row">
      <div class="col-sm-6">
        <div class="iframe-fix">
          <vue-mathjax :formula="main" :safe="false"></vue-mathjax>
        </div>
        <div class="box-centering margin-top">
          <table>
            <tr>
              <th>選択肢</th>
              <th>回答</th>
            </tr>
            <tr v-for="(selecter, index) in choice" :key="index">
              <td v-html="choice[index]"></td>
              <td v-html="choicetext[index]"></td>
            </tr>
          </table>
        </div>
      </div>
      <div class="col-sm-5">
        <div class="select">
          <div class="text-centering">
            <h2>解答欄</h2>
          </div>
          <hr>
          <div class="box-centering margin-bottom">
            <div class="row" v-if="type === 'select'" >
              <div v-for="(choicebtn, index) in choice" :key="index" class="col-sm-6">
                <button type="button" class="select">{{ choicebtn }}</button>
              </div>
            </div>
            <div v-else>
              <textarea class="answer-textarea" v-model="data.answer"></textarea>
            </div>
          </div>
        </div>
        <div class="input-group mb-3 margin-top">
          <span class="input-group-text button">前へ</span>
          <input type="text" id="status" class="form-control text-centering" value="00/00問 残り 01:00:00(PREVIEW SAMPLE)">
          <span class="input-group-text button">次へ</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { marked } from 'marked'

export default {
  name: "QuestionTemplate",
  data(){
    return {
      dataid: null,
      tag: '',
      type: '',
      answer: '',
      choice: null,
      choicetext: null,
      score: '',
      main: '',
      data: {
        "answer": '',
      }
    }
  },
  methods: {
    addQuestion() {
      this.$router.push(`/wizard/Question/${this.$route.params.id}`)
    }
  },
  mounted: function() {
    axios.post(
        'http://localhost:8080/FindQuestion',
        {
          ExName: this.$route.params.ExName,
          ExType: this.$route.params.ExType,
          QName: this.$route.params.QName
        }
    ).then(
        function(response){
          var exercise = response.data.DATA
          this.main = marked(exercise.QMain)
          this.main_title = exercise.QName
          this.tag = exercise.QTags
          this.type = exercise.AnswerArea
          this.answer = exercise.QAnswer
          this.score = exercise.QScore

          var select = []
          var choice = []

          if(this.type === "select"){
            var choice_csv = exercise.QChoice.split(",")
            choice_csv.forEach(function(choices){
              var choice_select = choices.split(":")
              choice.push(choice_select[0])
              select.push(choice_select[1])
            })
          }
          this.choice = choice
          this.choicetext = select
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
.container{
  margin-top: 5%;
}

.width-base{
  width: 90%;
  margin: auto;
  margin-top: 10px;
}
.text-centering{
  text-align: center;
}

.box-centering{
  margin: auto;
}

.select{
  width: 100%;
  border: 1px solid lightgray;
  padding: 10px;
}

.disp-block{
  display: block;
}

.float-left{
  float: left;
  overflow: hidden;
}

.select button{
  width: 100%;
  margin-top: 10px;
}

.select button:hover{
  opacity: 0.5;
}

.margin-bottom{
  margin-bottom: 10px;
}

.question-window{
  width: 40%;
  height: 300px;
}

.manager{
  width: 100%;
  margin: auto;
}

.margin-top{
  margin-top: 10px;
}

.manager button{
  display: block;
  margin: 0;
  border-radius: 0;
  height: 36px;
  float:left;
}

.manager input{
  display: block;
  margin: 0;
  border-radius: 0;
  height: 30px;
  float: left;
}

.button:hover{
  cursor: pointer;
  opacity: 0.5;
}

.iframe-fix{
  border: 1px solid lightgray;
  width: 100%;
  height: 40vh;
  max-height: 40vh;
  overflow: scroll;
  padding: 10px;
}

table {
  border: solid 1px #ccc;
  border-collapse:collapse;
}
th {
  padding: 5px;
  border: solid 1px #ccc;
  background-color: #eee;
}
td {
  padding: 5px;
  border: solid 1px #ccc;
}
</style>