<template>
  <div class="wizard">
    <h3>エクササイズ管理</h3>
    <p>あなたが作成したエクササイズ問題を確認できます。</p>
    <hr>
    <div class="text-center">
      <button type="button" class="btn btn-primary width-90" @click="home">エクササイズの新規作成</button>
    </div>
    <div class="exercises" v-for="(exercise, index) in exercises" :key="index">
        <router-link :to="`/detail/Ex/${ exercise.ExName }/${ exercise.ExType }`">
          <div class="exercise">
            <h5>{{ exercise.ExName }}</h5>
            <p>タグ：{{ exercise.ExTags }}</p>
          </div>
        </router-link>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "ExerciseManagement",
  data(){
    return {
      exercises: null
    }
  },
  methods: {
    home(){
      this.$router.push('/wizard/Ex')
    }
  },
  mounted: function(){
    axios.get(
        "http://localhost:8080/GetAllExercise",
        {
          withCredentials: true
        }
    ).then(
        function(response){
          this.exercises = response.data.DATA
          console.log(this.exercises)
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
</style>