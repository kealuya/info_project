<template>
  <RouterView v-slot="{ Component }">
    <keep-alive>
      <component v-if="$route.meta && $route.meta.keepAlive" :key="$route.name" :is="Component"/>
    </keep-alive>
    <component v-if="!($route.meta.keepAlive)" :key="$route.name" :is="Component"/>
  </RouterView>
</template>
<script setup lang="ts">
import {onBeforeUnmount} from "vue";

onBeforeUnmount(() => {
    window.localStorage.clear();
})
// onMounted(() => {
//   window.addEventListener('hashchange', () => {
//     let currentPath = window.location.hash.slice(1)
//     if (route.path !== currentPath) {
//       router.push(currentPath)
//     }
//   }, false)
// })
</script>
<style scoped>
</style>
