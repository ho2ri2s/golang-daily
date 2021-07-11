<template lang="pug">
h1 Diaries
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
}
</script>
