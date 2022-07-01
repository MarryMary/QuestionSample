<template>
  <div class="wizard">
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
          <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.main_title" placeholder="例）問1">
          <p>タグ</p>
          <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.tag" placeholder="例）Go言語 Map Struct Type Slice">
          <p>本文</p>
          <textarea class="form-control not-resize" placeholder="マークダウン形式で書いてみましょう！"></textarea>
          <div class="right-text">
            <button type="button" class="btn btn-success margin-top" @click="next">次へ</button>
          </div>
        </div>
        <div v-else-if="page === 2">
          <h3>問題作成ウィザード</h3>
          <p>問題の回答方式はどちらにしますか？</p>
          <hr>
          <div class="form-check">
            <input class="form-check-input" type="radio" name="flexRadioDefault" id="flexRadioDefault1">
            <label class="form-check-label" for="flexRadioDefault1">
              記述問題
            </label>
          </div>
          <div class="form-check">
            <input class="form-check-input" type="radio" name="flexRadioDefault" id="flexRadioDefault2" checked>
            <label class="form-check-label" for="flexRadioDefault2">
              回答問題
            </label>
          </div>
        </div>
        <div v-else-if="page === 3">
          <h3>問題作成ウィザード</h3>
          <p>選択肢を入力して下さい。</p>
          <hr>
          <div class="form-check">
            <input class="form-check-input" type="radio" name="flexRadioDefault" id="flexRadioDefault1">
            <label class="form-check-label" for="flexRadioDefault1">
              記述問題
            </label>
          </div>
          <div class="form-check">
            <input class="form-check-input" type="radio" name="flexRadioDefault" id="flexRadioDefault2" checked>
            <label class="form-check-label" for="flexRadioDefault2">
              回答問題
            </label>
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
      dumper: {
        main_title: '',
        detail: '',
        tag: '',
        time_limit: '01:00:00'
      }
    }
  },
  name: "QCreate",
  methods: {
    next(){
      if(this.dumper.main_title.trim() !== '' && this.dumper.detail.trim() !== '' && this.dumper.tag.trim() !== '' && this.dumper.time_limit.trim() !== '') {
        this.err = false
        this.page++
      }else{
        this.err = true
      }
    },

    back(){
      this.page--
    },

    dataImport() {
      const ref = this.$refs.input;
      const file = ref.files[0];
      console.log(file);
    },

    btnClick() {
      const ref = this.$refs.input;
      ref.click();
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
}
</style>