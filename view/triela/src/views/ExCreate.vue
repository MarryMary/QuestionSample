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
            <p>年度</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.year" placeholder="オリジナル問題の場合は現在の年度 下2桁">
            <p>季節</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.season" placeholder="オリジナル問題や季節が存在しない場合は未入力">
            <p>ジャンル</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.Genre" placeholder="AZ, AP等">
            <small>命名ルールは<a href="#">こちら</a></small>
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
            <p>年度</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.year" placeholder="オリジナル問題の場合は現在の年度 下2桁" disabled>
            <p>季節</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.season" placeholder="オリジナル問題や季節が存在しない場合は未入力" disabled>
            <p>ジャンル</p>
            <input type="text" name="text" class="form-control center-text margin-bottom" v-model="dumper.Genre" placeholder="AZ, AP等" disabled>
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
          <div v-else-if="page === 3">
            <h3>エクササイズ作成ウィザード</h3>
            <p>全て完了しました！</p>
            <hr>
            <p>エクササイズが追加されました。</p>
            <div class="text-center">
              <button type="button" class="btn btn-primary width-90" @click="toExManage">戻る</button>
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
        year: '',
        season: '',
        Genre: '',
        tag: '',
        time_limit: '01:00:00'
      }
    }
  },
  name: "ExCreate",
  methods: {
    next(){
      if(this.dumper.main_title.trim() !== '' && this.dumper.year.trim() !== '' && this.dumper.season.trim() !== '' && this.dumper.Genre.trim() !== '' && this.dumper.tag.trim() !== '' && this.dumper.time_limit.trim() !== '') {
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
            Year: this.dumper.year,
            Season: this.dumper.season,
            Genre: this.dumper.Genre,
            Tag: this.dumper.tag,
            Limit: this.dumper.time_limit
          },
          {
            withCredentials: true
          }
      ).then(
          function(response){
            if(response.status === 200){
              this.err = false
              this.next()
            }else{
              alert("サーバーで問題が発生しました。再度お試し下さい。")
            }
          }.bind(this)
      ).catch(
          error => console.log(error)
      )
    },

    toExManage() {
      this.$router.push(`/manage/Ex`)
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