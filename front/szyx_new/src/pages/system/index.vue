<template>
  <van-tabs v-model:active="active" sticky>
    <van-tab title="资产">
      <div style="width: 100%;height: 750px;">
      <iframe :src="ziChanUrl"  width="100%" height="100%" style="border: none">
      </iframe>
      </div>
    </van-tab>
    <van-tab title="合同">
      <div style="width: 100%;height: 750px;">
      <iframe src="https://zc.qcykj.com.cn/xx_asset_h5/#/contractQuery"  width="100%" height="100%" style="border: none" >
      </iframe>
      </div>
    </van-tab>
    <van-tab title="人事">
      <div style="width: 100%;height: 750px;">
      <iframe src="https://zc.qcykj.com.cn/xx_asset_h5/#/personnelQuery"  width="100%" height="100%" style="border: none;" >
      </iframe>
      </div>
    </van-tab>
    <van-tab title="库存">
      <div style="width: 100vw;height: 750px;">
      <iframe src="https://tech.szhtkj.com.cn/stock/index.html#/home"  width="100%" height="100%"  style="border: none" scrolling="no">
      </iframe>
      </div>
    </van-tab>
  </van-tabs>
</template>

<script setup lang="ts">
import {ref, inject} from "vue";
const active = ref(0);
let userInfoData: any = inject("userInfo");
const ziChanUrl = `https://zc.qcykj.com.cn/xx_asset_h5/#/home?userId=${userInfoData.userInfo.empCode}&client=yyk`
</script>

<style lang="less" scoped>
</style>
