<template>
  <router-view />
  <van-tabbar v-model="active"  route   >
<!--    <van-tabbar-item replace icon="home-o" :to="{name:'home'}"  >首页</van-tabbar-item>-->
    <!--    <van-tabbar-item icon="records" name="book" :to="{name:'book'}" replace>记账本<
    /van-tabbar-item>-->
    <van-tabbar-item name="news" icon="home-o"  :to="{name: 'homeMetting'}">首页</van-tabbar-item>
    <van-tabbar-item name="news" icon="apps-o"  :to="{name: 'business'}">业务</van-tabbar-item>
<!--    <van-tabbar-item replace icon="apps-o" :to="{name:'business'}"  >业务</van-tabbar-item>-->
  </van-tabbar>
</template>

<script setup lang="ts">
import {ref, onBeforeMount, inject} from "vue";
import {getMessagePull} from "./services/message";
import {userInfoData} from './store'
import router from "./router/index";
const active = ref('homeMetting');
const onChange = (index:string) => {
  console.log('index',index)

}






</script>

<style lang="less" scoped>
</style>
