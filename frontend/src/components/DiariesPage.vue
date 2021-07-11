<template lang="pug">
h1 Diaries
.main__button
  button(@click="onClickNewDiaryButton") New Diary
</template>

<script>
import axios from 'axios'

export default {
  name: 'Diaries',
  created() {
    axios.get('/diaries')
    .then((response) => {
      console.log(response.data);
    })
  },
  methods: {
    onClickNewDiaryButton: function() {
      this.$router.push('/diaries/new')
    }
  },
}
</script>

<style>
/* FIXME: 共通のCSSで定義したい */
.main__button button {
  color: #ffffff;
  font-weight: bold;
  background-color: #C7B379;
  padding: 0.5rem 1.0rem;
  margin: 2.0rem;
  border-radius: 0.3rem;
}
</style>