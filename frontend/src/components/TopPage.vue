<template lang="pug">
.main
  h1 My Diary
  .main__button
    button Let's go!
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
