<template>
  <van-tabs v-model:active="active" sticky>
    <van-tab title="语音识别">
<!--      <div style="width: 100%;height: 750px;">-->
<!--      <iframe :src="ziChanUrl"  width="100%" height="100%" style="border: none">-->
<!--      </iframe>-->
<!--      </div>-->
      <div>浏览器录音听写：<button id="btn_control">开始录音</button></div>
    </van-tab>
  </van-tabs>
</template>

<script setup lang="ts" >
import {ref, inject} from "vue";
import RecorderManager from "./kdxf";







const active = ref(0);
let userInfoData: any = inject("userInfo");
const ziChanUrl = `https://zc.qcykj.com.cn/xx_asset_h5/#/home?userId=${userInfoData.userInfo.empCode}&client=yyk`



var APPID = "";
var API_KEY = "";




</script>




