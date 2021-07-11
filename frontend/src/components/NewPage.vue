<template lang="pug">
h1 Create New Diary
form(action="/diaries/new", method="POST")
  dl
    dt: label(for="title") Title
    input#title(type="text", name="diary_title")
    dt: label(for="content") Content
    textarea#content(type="text", name="diary_content")
    MainButton(type="submit", value="Create")
</template>

<script>
import MainButton from "./common/MainButton.vue";

export default {
  name: "New",
  components: {
    MainButton,
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