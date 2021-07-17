<template lang="pug">
.diary__item
  img(alt="logo", src="../../assets/Notes-bro.svg")
  h2 {{listItem.title}}

</template>

<script>
export default {
  name: "DiaryItem",
  props: {
    listItem: { type: Object, required: false },
  },
};
</script>

<style>
.diary__item img {
  border-bottom: solid 0.1rem;
}
</style>