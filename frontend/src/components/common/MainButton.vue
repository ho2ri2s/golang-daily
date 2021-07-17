<template lang="pug">
.main__button
  button(type=buttonType, @click="onClick") {{ value }}
</template>

<script>
export default {
  name: "MainButton",
  props: {
    buttonType: { type: String, required: false },
    value: { type: String, required: true },
  },
  methods: {
    /**
     * 親にクリックイベントを渡す
     */
    onClick: function () {
      this.$emit("click");
    },
  },
};
</script>

<style>
.main__button button {
  color: #ffffff;
  font-weight: bold;
  background-color: #c7b379;
  padding: 0.5rem 1rem;
  margin: 2rem;
  border-radius: 0.3rem;
}

.main__button button:hover {
  opacity: 0.3;
  transition: 0.3s;
}
</style>