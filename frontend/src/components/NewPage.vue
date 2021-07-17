<template lang="pug">
h1 Create New Diary
form(@submit.prevent="submit")
  dl
    dt: label(for="title") Title
    input#title(type="text", v-model="diary_title")
    dt: label(for="content") Content
    textarea#content(type="text", v-model="diary_content")
    MainButton(type="submit", value="Create")
</template>

<script>
import axios from "axios";
import MainButton from "./common/MainButton.vue";

export default {
  name: "New",
  components: {
    MainButton,
  },
  data() {
    return {
      diary_title: '',
      diary_content: '',
    }
  },
  methods: {
    submit() {
      const params = new URLSearchParams()
      params.append('diary_title', this.diary_title)
      params.append('diary_content', this.diary_content)
      axios.post('/diaries/new', params)
      .then(respose => {
        console.log(respose)
        this.$router.push('/diaries')
      })
      .catch(error => {
        console.log(error)
      });
    },
  },
};
</script>

<style>
form {
  width: 40rem;
  margin-left: auto;
  margin-right: auto;
}
</style>