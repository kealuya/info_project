<template>
  <van-config-provider>
    <van-sticky>
      <van-nav-bar :border="false"  @click-left="onClickLeft">
        <template #left>
          <van-icon name="arrow-left" color="#ffffff"  size="20" />
          <span class="f_white">返回</span>
        </template>
        <template #title>
          <h4 class="f_white">{{title}}</h4>
        </template>
        <template #right>
          <van-icon :name="icon" size="22" color="white" @click="rightEvent" />
        </template>
      </van-nav-bar>
    </van-sticky>
  </van-config-provider>
</template>

<script lang="ts" setup>
import router from "../../router";
import {defineProps, defineEmits, ref} from 'vue'
interface Props{
  title?:any;
  icon?:any;
  funcition?:any;
}
defineProps<Props>()
const emit = defineEmits(['leftEvent',"rightEvent"])
const onClickLeft = () => {
  emit('leftEvent',router.back())
};
const rightEvent = (() => {
  emit("rightEvent" as never)
})
</script>

<style scoped>
.van-nav-bar{
  background:#0088ff;
}
</style>
