<template>
  <router-view />
  <van-tabbar v-model="active"  route safe-area-inset-bottom v-if="isShowTabBar">
    <!--    <van-tabbar-item replace icon="home-o" :to="{name:'home'}"  >首页</van-tabbar-item>-->
    <!--    <van-tabbar-item icon="records" name="book" :to="{name:'book'}" replace>记账本<
    /van-tabbar-item>-->
    <van-tabbar-item name="news" icon="home-o"  :to="{name: 'homeNew'}">首页</van-tabbar-item>
    <van-tabbar-item name="news" icon="apps-o"  :to="{name: 'business'}">业务</van-tabbar-item>
        <van-tabbar-item replace icon="user-o" :to="{name:'personal'}"  >我的</van-tabbar-item>
  </van-tabbar>
</template>

<script setup lang="ts">
import {ref, onBeforeMount, inject,watch} from "vue";
import { useRoute, useRouter } from 'vue-router';
const active = ref('homeMetting');
const route = useRoute();
const router = useRouter();
const isShowTabBar=ref<boolean>(false)
const noSHowTabBarPathList=['/finishDetail','/speech']    //在这里面加 你要隐藏哪一页 的 路径
watch(() => route.path, (value) => {
  isShowTabBar.value=!noSHowTabBarPathList.includes(value)
},{
  immediate:true,
  deep:true
});

</script>

<style lang="less" scoped>
</style>
