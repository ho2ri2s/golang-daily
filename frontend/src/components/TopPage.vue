<template lang="pug">
.main
  h1 My Diary
  .main__button
    button Let's go
  img(alt="logo" src="../assets/Notes-bro.svg")
</template>

<script>
export default {
  name: 'TopPage',
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}

img {
  width: 50rem;
}
</style>
