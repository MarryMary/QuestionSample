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
          <div class="row">
            <div v-bind:class="{'col-sm-6 text-left': preview_visible, 'col-sm-12 text-left': !preview_visible}">
              <textarea class="form-control not-resize" v-model="dumper.main_text" placeholder="マークダウン形式で書いてみましょう！" required="required" @input="Change"></textarea>
            </div>
            <div class="col-sm-6 border text-left" v-show="preview_visible">
              <vue-mathjax :formula="formula" :safe="false"/>
            </div>
            <div class="left-text margin-top">
              <button type="button" class="btn btn-warning" @click="Visible"><font-awesome-icon icon="fa-solid fa-eye" />問題文プレビュー表示</button>
            </div>
          </div>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="ValidSettings">次へ</button>
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
          <div class="left-text float-left">
            <button type="button" class="btn btn-success margin-top" @click="back">修正</button>
          </div>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="ValidType">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 3 && dumper.qtype === 'write'">
          <h3>問題作成ウィザード</h3>
          <p>回答記入時の採点ポイントとなる単語を空白区切りで入力して下さい。</p>
          <hr>
          <input type="text" class="form-control width-90" v-model="dumper.answer" placeholder="例）Map Slice Struct" required="required">
          <div class="left-text float-left">
            <button type="button" class="btn btn-success margin-top" @click="back">修正</button>
          </div>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="ValidPoint">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 3 && dumper.qtype === 'select'">
          <h3>問題作成ウィザード</h3>
          <p>選択肢を入力して下さい。</p>
          <hr>
          <div v-for="(answers, index) in select" :key="index" class="box-centering">
            <div class="input-group">
              <input type="text" placeholder="ア" class="form-control" v-model="answers.select">
              <input type="text" placeholder="Type Sample Struct" v-model="answers.answer" class="form-control">
            </div>
          </div>
          <div class="center-text">
            <button type="button" class="btn btn-success width-90 margin-top" @click="addAnswer">回答の追加</button>
          </div>
          <div class="left-text float-left">
            <button type="button" class="btn btn-success margin-top" @click="back">修正</button>
          </div>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="ValidChoose">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 4 && dumper.qtype === 'select'">
          <h3>問題作成ウィザード</h3>
          <p>回答を選択して下さい。</p>
          <hr>
          <div v-for="(exanswer, index) in select" :key="index">
            <div class="form-check">
              <input class="form-check-input" type="radio" v-model="dumper.answer" name="answer" v-bind:value="exanswer.select" v-bind:id="index" required="required" checked>
              <label class="form-check-label" v-bind:for="index">
                {{exanswer.answer}}
              </label>
            </div>
          </div>
          <div class="left-text float-left">
            <button type="button" class="btn btn-success margin-top" @click="back">修正</button>
          </div>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="ValidAnswer">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 5 && dumper.qtype === 'select' || page === 4 && dumper.qtype === 'write'">
          <h3>問題作成ウィザード</h3>
          <p>この問題の配点を整数で入力して下さい。(小数は採点時に無視されます。)</p>
          <hr>
          <div class="box-centering">
            <input type="text" class="form-control width-90 center-text" v-model="dumper.score" placeholder="例）10" required="required">
          </div>
          <div class="left-text float-left">
            <button type="button" class="btn btn-success margin-top" @click="back">修正</button>
          </div>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="ValidScore">次へ</button>
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
          <input class="form-control" v-model="dumper.score" disabled>
          <p>本文</p>
          <div class="border text-left" style="padding: 10px">
            <vue-mathjax :formula="formula" :safe="false"></vue-mathjax>
          </div>
          <div class="left-text float-left">
            <button type="button" class="btn btn-success margin-top" @click="back">修正</button>
          </div>
          <div class="right-text">
            <button type="button" class="btn btn-primary margin-top" @click="sendEnter">作成</button>
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
            <button type="button" class="btn btn-primary width-90">エクササイズ管理画面へ</button>
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

<script type="module">
import axios from "axios"
import { marked } from 'marked'

export default {
  data(){
    return {
      page: 1,
      finished: false,
      err: false,
      errmsg: '1つまたは複数の入力項目が欠損しています。',
      formula: '',
      preview_visible: false,
      select:[
        {answer: "", select: ""},
        {answer: "", select: ""}
      ],
      dumper: {
        main_title: '',
        tag: '',
        main_text: '',
        qtype: 'select',
        choice: '',
        answer: '',
        score: '0'
      }
    }
  },
  name: "QCreate",
  methods: {
    ValidSettings(){
      if(this.dumper.main_title.trim() !== "" && this.dumper.tag.trim() !== "" && this.dumper.main_text.trim() !== ""){
        this.err = false
        this.next()
      }else{
        this.err = true
        if(this.dumper.main_title.trim() === "") {
          alert("タイトルが無効です。空白のみや未入力にはできません。")
        }else if(this.dumper.tag.trim() === ""){
          alert("タグが無効です。空白のみや未入力にはできません。")
        }else if(this.dumper.main_text.trim() === ""){
          alert("本文が無効です。空白のみや未入力にはできません。")
        }
      }
    },

    Visible(){
      (this.preview_visible ? this.preview_visible = false : this.preview_visible = true)
    },

    Change(){
      this.formula = marked(this.htmlspecialchars(this.dumper.main_text))
    },

    ValidType(){
      if(this.dumper.qtype.trim() !== ""){
        this.err = false
        this.next()
      }else{
        this.err = true
        alert("タイプが選択されていません。タイプは必ず選択する必要があります。")
      }
    },

    ValidPoint(){
      if(this.dumper.answer.trim() !== ""){
        this.err = false
        this.next()
      }else{
        this.err = true
        alert("採点のポイントが無効です。空白のみや未入力にはできません。")
      }
    },

    ValidChoose(){
      var choice = []
      this.select.forEach((elem, index) => {
        if(elem.answer.trim() === "" && elem.select.trim() === ""){
          this.select.splice(index, 1)
        }else{
          var insert = String(elem.select.replace("：", ":").replace(":", "&#058;")).replace(",", "&#044;")+':'+String(elem.answer.replace("：", ":").replace(":", "&#058;").replace(",", "&#044;"))
          choice.push(insert)
        }
      })
      if(this.select.length === 0){
        this.addAnswer()
        this.err = true
        alert("選択肢は2つ以上入力する必要があります。")
      }else{
        this.dumper.choice = choice.join(",")
        this.err = false
        this.next()
      }
    },

    ValidAnswer(){
      if(this.dumper.answer.trim() !== ""){
        this.err = false
        this.next()
      }else{
        this.err = true
        alert("回答は必ず設定する必要があります。")
      }
    },

    ValidScore(){
      if(this.dumper.score.trim() !== "" && this.ctype_digit(this.dumper.score)){
        this.err = false
        this.next()
      }else{
        this.err = true
        if(this.ctype_digit(this.dumper.score)){
          alert("入力は必ず数値（整数）である必要があります。小数は自動的に四捨五入されます。")
        }else{
          alert("得点が無効です。空白のみや未入力にはできません。")
        }
      }
    },

    ctype_digit(n){
      const v = n - 0;
      return !!(v || v === 0);

    },

    next(){
      this.page++
    },

    back(){
      this.page--
    },

    delform(){
      if(this.select.length < 2){
        alert("これ以上回答を減らすことはできません。")
      }else {
        this.select.pop()
      }
    },

    addAnswer(){
      this.select.push({answer: "", select: ""})
    },

    btnClick() {
      const ref = this.$refs.input;
      ref.click();
    },

    reload() {
      this.$router.go({path: this.$router.currentRoute.path, force: true})
    },

    htmlspecialchars(unsafeText){
      if(typeof unsafeText !== 'string'){
        return unsafeText;
      }
      return unsafeText.replace(
          /[&'`"<>]/g,
          function(match) {
            return {
              '&': '&amp;',
              "'": '&#x27;',
              '`': '&#x60;',
              '"': '&quot;',
              '<': '&lt;',
              '>': '&gt;',
            }[match]
          }
      );
    },

    sendEnter(){
      axios.post(
          'http://localhost:8080/PushQuestion',
          {
            ExName: this.$route.params.ExName,
            ExType: this.$route.params.ExType,
            MainTitle: this.dumper.main_title,
            Tag: this.dumper.tag,
            MainText: this.dumper.main_text,
            QType: this.dumper.qtype,
            Answer: this.dumper.answer,
            Choice: this.dumper.choice,
            Score: this.dumper.score
          },
          {
            withCredentials: true
          }
      ).then(
          function(response){
            if(response.status === 200){
              this.next()
            }else{
              alert("サーバーで問題が発生しました。再度お試し下さい。")
            }
          }.bind(this)
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