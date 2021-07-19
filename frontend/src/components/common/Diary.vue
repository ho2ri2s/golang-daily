<template lang="pug">
h1.diary__item(v-on:click="onClickIndexButton")
  img(alt="logo", src="../../assets/Notes-bro.svg")
  h2 {{listItem.title}}

</template>

<script>
export default {
  name: "DiaryItem",
  props: {
    listItem: { type: Object, required: false },
  },
  created() {
    const id = "/diaries/" + String(this.listItem.ID);
    console.log(id);
  },
  method: {
    onClickIndexButton: function () {
      const id = "/diaries/" + String(this.listItem.ID);
      console.log(id);
      this.$router.push(id);
    },
  }
};
</script>

<style>
.diary__item{
  display: flex;
  flex-direction: column;
  z-index: +1000;
}
.diary__item img {
  border-bottom: solid 0.1rem;
}
</style>