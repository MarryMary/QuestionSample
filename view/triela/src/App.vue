<template>
  <div class="sidebar">
    <div class="user-data">
      <img src="https://avatars.githubusercontent.com/u/92404990?v=4" width="50" height="50" alt="profile">
      <div class="flex-vertical">
        <h3>田中琴葉</h3>
        <p>学生</p>
        <a href="#">ログアウト</a>
      </div>
    </div>
    <hr>
    <div class="selecter">
      <ul>
        <li><font-awesome-icon icon="home"/>ホーム</li>
        <hr>
        <li><font-awesome-icon icon="book"/>テスト管理</li>
        <hr>
        <li><font-awesome-icon icon="pen"/>問題管理</li>
      </ul>
    </div>
  </div>
  <div class="content">
    <router-view/>
  </div>
</template>

<style>
*{
  margin: 0;
  padding: 0;
}

.sidebar{
  box-shadow: 5px 0 3px lightgray;
  width: 25vh;
  height: 100vh;
  float: left;
  margin-right: 10px;
}

.user-data{
  padding: 5%;
  display: flex;
}

.user-data img{
  margin-right: 5%;
}

.flex-vertical{
  display:flex;
  flex-direction: column;
}

.selecter{
  width: 80%;
  margin: auto;
  padding-top: 5%;
  text-align: center;
}

.box-centering{
  width: 70%;
  margin: auto;
}

.selecter ul{
  list-style: none;
  border: 1px solid lightgray;
  margin: 0;
  padding: 0;
}

.selecter li{
  border-bottom: 1px lightgray;
  padding: 10px;
}

.selecter li:hover{
  opacity: 0.5;
  cursor: pointer;
}

ul hr{
  margin: 0;
}

.content{
  padding: 10px;
  margin: 0;
}

.left-text{
  text-align: left;
}

.right-text{
  text-align: right;
}

.center-text{
  text-align: center;
}

.content{
  margin-left: 25vh
}

.wizard{
  width: 80%;
  margin: auto;
  padding: 2%;
  box-shadow: 0 0 8px gray;
}

.margin-top{
  margin-top: 40px;
}

.float-right{
  float: right;
}

.float-left{
  float: left;
}
</style>
