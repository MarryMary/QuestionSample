<template>
  <div class="wizard qwizard">
    <div class="center-text">
      <transition mode="out-in">
        <div v-if="page === 1">
          <h3>問題作成ウィザード</h3>
          <p>以下に必要事項を記載してください。</p>
          <div v-if="err">
            <p style="color: red;">{{ errmsg }}</p>
          </div>
          <hr>
          <p>タイトル</p>
          <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.main_title" placeholder="例）問1" required="required">
          <p>タグ</p>
          <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.tag" placeholder="例）Go言語 Map Struct Type Slice" required="required">
          <p>本文</p>
          <textarea class="form-control not-resize" v-model="dumper.main_text" placeholder="マークダウン形式で書いてみましょう！" required="required"></textarea>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="next">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 2">
          <h3>問題作成ウィザード</h3>
          <p>問題の回答方式はどちらにしますか？</p>
          <hr>
          <div class="form-check">
            <input class="form-check-input" type="radio" v-model="dumper.qtype" name="qtype" id="write" value="write">
            <label class="form-check-label" for="write">
              記述問題
            </label>
          </div>
          <div class="form-check">
            <input class="form-check-input" type="radio" v-model="dumper.qtype" name="qtype" id="select" value="select">
            <label class="form-check-label" for="select">
              選択問題
            </label>
          </div>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="next">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 3 && dumper.qtype === 'write'">
          <h3>問題作成ウィザード</h3>
          <p>回答記入時の採点ポイントとなる単語を空白区切りで入力して下さい。</p>
          <hr>
          <input type="text" class="form-control width-90" v-model="dumper.answer" placeholder="例）Map Slice Struct" required="required">
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="next">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 3 && dumper.qtype === 'select'">
          <h3>問題作成ウィザード</h3>
          <p>選択肢を入力して下さい。</p>
          <hr>
          <div v-for="(answers, index) in select" :key="index" class="box-centering">
            <input type="text" class="form-control width-90" v-model="answers.answer" placeholder="例）ア:Type Sample Struct" required="required">
          </div>
          <div class="center-text">
            <button type="button" class="btn btn-success width-90 margin-top" @click="addAnswer">回答の追加</button>
          </div>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="next">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 4 && dumper.qtype === 'select'">
          <h3>問題作成ウィザード</h3>
          <p>回答を入力して下さい。</p>
          <hr>
          <div v-for="(exanswer, index) in select" :key="index">
            <div class="form-check">
              <input class="form-check-input" type="radio" v-model="dumper.answer" name="answer" id="select" value="select" required="required">
              <label class="form-check-label" for="select">
                {{exanswer.answer}}
              </label>
            </div>
          </div>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="next">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 5 && dumper.qtype === 'select' || page === 4 && dumper.qtype === 'write'">
          <h3>問題作成ウィザード</h3>
          <p>この問題の配点を整数で入力して下さい。(小数は採点時に無視されます。)</p>
          <hr>
          <input type="text" class="form-control width-90" v-model="dumper.score" placeholder="例）10" required="required">
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="next">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 6 && dumper.qtype === 'select' || page === 5 && dumper.qtype === 'write'">
          <h3>問題作成ウィザード</h3>
          <p>以下の内容で登録します。</p>
          <hr>
          <p>タイトル</p>
          <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.main_title" disabled>
          <p>タグ</p>
          <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.tag" disabled>
          <p>問題種別</p>
          <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.qtype" disabled>
          <p>回答</p>
          <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.answer" disabled>
          <p>得点</p>
          <input class="form-control" v-model="dumper.main_text" disabled>
          <p>本文</p>
          <textarea class="form-control not-resize" v-model="dumper.main_text" disabled></textarea>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="next">作成</button>
          </div>
        </div>
        <div v-else-if="page === 7 && dumper.qtype === 'select' || page === 6 && dumper.qtype === 'write'">
          <h3>問題作成ウィザード</h3>
          <p>全て完了しました！</p>
          <hr>
          <p>
            問題がエクササイズに追加されました。<br>
            引き続き問題を追加する場合は<a @click="reload">こちら</a>
          </p>
          <div class="text-center">
            <button type="button" class="btn btn-primary width-90" @click="home">エクササイズ管理画面へ</button>
          </div>
        </div>
        <div v-else>
          <h3>Oops!</h3>
          <p>問題が発生しました。やり直して下さい。</p>
        </div>
      </transition>
    </div>
  </div>
</template>

<script>
import axios from "axios"
export default {
  data(){
    return {
      page: 1,
      finished: false,
      err: false,
      errmsg: '1つまたは複数の入力項目が欠損しています。',
      select:[
        {answer: ""}
      ],
      dumper: {
        main_title: '',
        tag: '',
        main_text: '',
        qtype: 'select',
        answer: '',
        score: '0'
      }
    }
  },
  name: "QCreate",
  methods: {
    next(){
      this.page++
    },

    back(){
      this.page--
    },

    addAnswer(){
      this.select.push({answer: ""})
    },

    btnClick() {
      const ref = this.$refs.input;
      ref.click();
    },

    reload() {
      this.$router.go({path: this.$router.currentRoute.path, force: true})
    },

    sendEnter(){
      axios.post(
          'http://localhost:8080/PushExercise',
          {
            MainTitle: this.dumper.main_title,
            Detail: this.dumper.detail,
            Tag: this.dumper.tag,
            Limit: this.dumper.time_limit
          }
      ).then(
          function(response){
            console.log(response)
          }
      ).catch(
          error => console.log(error)
      )

    }
  }
}
</script>

<style scoped>
.v-enter-to {
  transition: transform 0.1s;
  transform: translateX(0px);
}
.v-enter-from {
  transform: translateX(50px);
}

.v-leave-to {
  transition: transform 0.01s;
  transform: translateX(-50px);
}
.v-leave-from {
  transform: translateX(0px);
}

.not-resize{
  resize: none;
  height: 40vh;
  max-height: 40vh;
  overflow-y: scroll;
}

a{
  text-decoration: underline!important;
  color: blue!important;
}

a:hover{
  opacity: 0.5;
  cursor: pointer;
}
</style>