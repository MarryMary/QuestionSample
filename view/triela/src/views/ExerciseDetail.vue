<template>
  <div class="wizard">
    <h1>名称：{{ main_title }}</h1>
    <p>エクササイズ情報の編集や、問題の作成ができます。</p>
    <hr>
    <h2>情報</h2>
    <hr>
    <div class="text-center">
      <p>説明：<br>{{ detail }}</p>
      <p>タグ：<br>{{ tag }}</p>
      <p>制限時間：<br>{{time_limit}}</p>
    </div>
    <div class="text-center margin-bottom">
      <button type="button" class="btn btn-primary width-90" @click="home">エクササイズの情報修正</button>
    </div>
    <h2>問題</h2>
    <hr>
    <div class="text-center">
      <button type="button" class="btn btn-success width-90" @click="addQuestion">問題の追加</button>
    </div>
    <div class="exercises" v-for="(question, index) in questions" :key="index">
      <router-link :to="`/detail/Q/${ question.QId }`">
        <div class="exercise">
          <h5>{{ question.QName }}</h5>
          <p>タグ：{{ question.QTags }}</p>
        </div>
      </router-link>
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
      time_limit: '01:00:00',
      questions: null
    }
  },
  methods: {
    addQuestion(){
      this.$router.push(`/wizard/Question/${ this.$route.params.id }`)
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
a{
  text-decoration: none;
  color: black;
}

p{
  font-size: 20px;
}
</style>