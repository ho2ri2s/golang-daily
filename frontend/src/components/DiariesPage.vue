<template lang="pug">
.diaries
  h1 Diaries
  .diaries__list
    .diaries__list__item(v-for="listItem in this.diaryList")
      diary-item(:listItem="listItem")
  .diaries__btn
    MainButton(value="New Diary", @click="onClickNewDiaryButton")
</template>

<script>
import axios from "axios";
import MainButton from "./common/MainButton.vue";
import DiaryItem from "./common/Diary.vue";

export default {
  name: "Diaries",
  components: {
    MainButton,
    DiaryItem,
  },
  data () {
    return {
      diaryList: [],
    }
  },
  created() {
    axios.get("/diaries").then((response) => {
      console.log(response.data);
      this.diaryList= response.data;
    });
  },
  methods: {
    onClickNewDiaryButton: function () {
      this.$router.push("/diaries/new");
    },
  },
};
</script>

<style>
.diaries__list {
  display: flex;
  flex-wrap: wrap;
}


.diaries__list__item {
  width: calc(16rem - 2rem);
  margin: 1rem 1rem;
  border: solid 0.1rem;
  border-radius: 0.3rem;
}
.diaries__list__item:hover{
  cursor: pointer;
  opacity: 0.6;
  transition: 0.5s;
  transform: translateY(-30px);
}
</style>
