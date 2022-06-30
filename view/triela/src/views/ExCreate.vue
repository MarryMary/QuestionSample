<template>
  <div class="wizard">
      <div class="center-text">
        <transition mode="out-in">
          <div v-if="page === 1">
            <h3>エクササイズ作成ウィザード</h3>
            <p>以下に必要事項を記載してください。</p>
            <div v-if="err">
              <p style="color: red;">{{ errmsg }}</p>
            </div>
            <hr>
            <p>タイトル</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.main_title" placeholder="例）応用技術者認定試験 午前">
            <p>詳細情報</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.detail" placeholder="例）2018(平成30年)の秋に開催された午前問題です。">
            <p>タグ</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.tag" placeholder="例）応用技術者 応用 技術者 午前">
            <p>制限時間(時間：分：秒)</p>
            <input type="time" step="1" class="form-control center-text margin-bottom" v-model="dumper.time_limit">
            <div class="right-text">
              <button type="button" class="btn btn-success margin-top" @click="next">確認と作成</button>
            </div>
          </div>
          <div v-else-if="page === 2">
            <h3>エクササイズ作成ウィザード</h3>
            <p>以下の内容でエクササイズを作成しても宜しいですか？</p>
            <hr>
            <p>タイトル</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.main_title" disabled>
            <p>詳細情報</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.detail" disabled>
            <p>タグ</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.tag" disabled>
            <p>制限時間(時間：分：秒)</p>
            <input type="time" step="1" class="form-control center-text margin-bottom" v-model="dumper.time_limit" disabled>
            <div class="left-text float-left">
              <button type="button" class="btn btn-primary margin-top" @click="back">修正</button>
            </div>
            <div class="right-text">
              <button type="button" class="btn btn-success margin-top" @click="sendEnter">作成する</button>
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
  name: "ExCreate",
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
</style>