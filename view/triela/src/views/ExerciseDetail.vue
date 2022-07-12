<template>
  <div class="wizard">
    <h1>名称：{{ main_title }}</h1>
    <p>エクササイズ情報の編集や、問題の作成ができます。</p>
    <hr>
    <h2>情報</h2>
    <hr>
    <div class="text-center">
      <p>年度：<br>{{ year }}</p>
      <p>時期：<br>{{ season }}</p>
      <p>問題ジャンル：<br>{{ genre }}</p>
      <p>タグ：<br>{{ tag }}</p>
      <p>制限時間：<br>{{time_limit}}</p>
    </div>
    <div class="text-center margin-bottom">
      <button type="button" class="btn btn-primary width-90" @click="FixExercise">エクササイズの情報修正</button>
    </div>
    <h2>問題</h2>
    <hr>
    <div class="text-center">
      <button type="button" class="btn btn-success width-90" @click="addQuestion">問題の追加</button>
    </div>
    <div class="exercises" v-for="(question, index) in questions" :key="index">
      <router-link :to="`/detail/Q/${ this.$route.params.ExName }/${ this.$route.params.ExType }/${ question.QName }`">
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
      main_title: '',
      year: '',
      season: '',
      genre: '',
      tag: '',
      time_limit: '01:00:00',
      questions: null
    }
  },
  methods: {
    addQuestion(){
      this.$router.push(`/wizard/Question/${ this.$route.params.ExName }/${ this.$route.params.ExType }`)
    },

    FixExercise(){
      this.$router.push(`/manage/fix/Ex/${ this.$route.params.ExName }/${ this.$route.params.ExType }`)
    }
  },
  mounted: function() {
    axios.post(
        'http://localhost:8080/FindExercise',
        {
          ExName: this.$route.params.ExName,
          ExType: this.$route.params.ExType,
        },
        {
          withCredentials: true
        }
    ).then(
        function(response){
          var exercise = response.data.DATA
          console.log(exercise)
          this.main_title = exercise.ExName
          this.year = exercise.ExYear
          this.season = exercise.ExSeason
          this.genre = exercise.ExType
          this.tag = exercise.ExTags
          this.time_limit = exercise.ExTimeLimit
        }.bind(this)
    ).catch(
        error => console.log(error)
    )

    axios.get(
        'http://localhost:8080/GetAllQuestion',
        {
          withCredentials: true
        }
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