<template lang="pug">
h1 Diaries
  MainButton(value="New Diary", @click="onClickNewDiaryButton")
</template>

<script>
import axios from "axios";
import MainButton from "./common/MainButton.vue";

export default {
  name: "Diaries",
  components: {
    MainButton,
  },
  created() {
    axios.get("/diaries").then((response) => {
      console.log(response.data);
    });
  },
  methods: {
    onClickNewDiaryButton: function () {
      this.$router.push("/diaries/new");
    },
  },
};
</script>
