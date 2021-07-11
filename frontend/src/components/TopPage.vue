<template lang="pug">
.main
  h1 My Diary
  .main__button
    button(v-on:click="onClickButton") Let's go!
  img(alt="logo" src="../assets/Notes-bro.svg")
</template>

<script>
export default {
  name: 'TopPage',
  methods: {
    onClickButton: function() {
      this.$router.push("/diaries")
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
img {
  width: 30rem;
}

.main__button button {
  color: #ffffff;
  font-weight: bold;
  background-color: #C7B379;
  padding: 0.5rem 1.0rem;
  margin: 2.0rem;
  border-radius: 0.3rem;
}

.main__button button:hover {
  opacity: 0.3;
  transition: 0.3s;
}
</style>
