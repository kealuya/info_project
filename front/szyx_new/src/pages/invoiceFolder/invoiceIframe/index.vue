<template>
  <div style="width: 100%;height: 820px">
    <iframe :src="iframeUrl" width="100%" height="100%" style="border: none"></iframe>
  </div>
</template>

<script setup lang="ts">
import {onBeforeMount} from "vue";
import {useRoute} from 'vue-router'
const route = useRoute();
let iframeUrl='';
onBeforeMount(async ()=>{
  iframeUrl=sessionStorage.getItem('url') as string
})

</script>

<style scoped>

</style>
