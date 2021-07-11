<template lang="pug">
h1 Create New Diary
form(action="/diaries/new" method="POST")
  dl
    dt: label(for="title") Title
    input(type="text", name="diary_title", id="title")
    dt: label(for="content") Content
    textarea(type="text", name="diary_content", id="content")
    .main__button
      button(type="submit") Create
</template>

<script>

export default {
  name: "New",
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

form {
  width: 40rem;
  margin-left: auto;
  margin-right: auto;
}
</style>