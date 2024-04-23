<template>
  <el-button @click="toggleDarkMode">Toggle Dark Mode</el-button>
</template>

<script>
export default {
  data() {
    return {
      darkMode: false
    };
  },
  methods: {
    toggleDarkMode() {
      this.darkMode = !this.darkMode;
      document.body.classList.toggle('dark-mode', this.darkMode);
    }
  }
};
</script>

<style>
body {
  transition: background-color 0.5s ease, color 0.5s ease;
}

.dark-mode {
  background: #333;
  /* color: #fff; */
}
</style>
