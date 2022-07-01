<template>
  <div class="wizard">
    <h1>{{ main_title }}</h1>
    <p>エクササイズ情報の編集や、問題の作成ができます。</p>
    <hr>
    <h2>情報</h2>
    <hr>
    <div class="text-center">
      <h3>説明：{{ detail }}</h3>
      <h3>タグ：{{ tag }}</h3>
      <h3>制限時間：{{time_limit}}</h3>
    </div>
    <div class="text-center margin-bottom">
      <button type="button" class="btn btn-primary width-90" @click="home">エクササイズの情報修正</button>
    </div>
    <h2>問題</h2>
    <hr>
    <div class="text-center">
      <button type="button" class="btn btn-success width-90" @click="home">問題の追加</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "ExerciseDetail.vue",
  data(){
    return {
      dataid: null,
      main_title: '',
      detail: '',
      tag: '',
      time_limit: '01:00:00'
    }
  },
  mounted: function() {
    axios.post(
        'http://localhost:8080/FindExercise',
        {
          ExId: this.$route.params.id
        }
    ).then(
        function(response){
          var exercise = response.data.DATA
          this.main_title = exercise.ExName
          this.detail = exercise.ExDetail
          this.tag = exercise.ExTags
          this.time_limit = exercise.ExTimeUp
          this.dataid = this.$route.params.id
        }.bind(this)
    ).catch(
        error => console.log(error)
    )
  }
}
</script>

<style scoped>

</style>