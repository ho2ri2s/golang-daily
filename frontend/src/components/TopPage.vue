<template lang="pug">
.main
  h1 My Diary
  MainButton(value="Let's go!", @click="onClickButton")
  img(alt="logo", src="../assets/Notes-bro.svg")
</template>

<script>
import MainButton from "./common/MainButton.vue";
export default {
  name: "TopPage",
  components: {
    MainButton,
  },
  methods: {
    onClickButton: function () {
      this.$router.push("/diaries");
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
img {
  width: 30rem;
}
</style>
