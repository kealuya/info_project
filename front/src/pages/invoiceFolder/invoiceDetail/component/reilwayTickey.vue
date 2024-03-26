<template>
  <div>
    <div>
      <van-row align="center">
        <van-col span="6" :offset="2" class="graytxt">车次：</van-col>
        <van-col span="6" class="detail">{{ getData.fpdm}}</van-col>
      </van-row>
    </div>
    <div class="mt5 f14">
      <div class="line2"></div>
    </div>
    <div class="mt10">
      <van-row align="center">
        <van-col span="6" :offset="2" class="graytxt">出行人：</van-col>
        <van-col span="6" class="detail">{{ getData.xm??'----' }}</van-col>
      </van-row>
    </div>
    <div class="mt10">
      <van-row align="center">
        <van-col span="6" :offset="2" class="graytxt">出发时间：</van-col>
        <van-col span="12" class="detail">{{ trainFormat(getData.cfsj) ??'----'}}</van-col>
      </van-row>
    </div>
    <div class="mt5 f14">
      <div class="line2"></div>
    </div>
    <div class="mt10">
      <van-row align="center">
        <van-col span="6" :offset="2" class="graytxt">出发地点：</van-col>
        <van-col span="6" class="detail">{{ getData.cfdd ??'----'}}</van-col>
      </van-row>
    </div>
    <div class="mt10">
      <van-row align="center">
        <van-col span="6" :offset="2" class="graytxt">到达地点：</van-col>
        <van-col span="12" class="detail">{{ getData.dddd ??'----'}}</van-col>
      </van-row>
    </div>
    <div class="mt5 f14">
      <div class="line2"></div>
    </div>
    <van-row class="mt10">
      <van-col span="6" :offset="2" class="graytxt">发票金额：</van-col>
      <van-col span="6" class="detail" style="color: #C6666A">￥{{ saveMoneyTwoNum(getData.zje) }}</van-col>
    </van-row>
  </div>
</template>

<script setup lang="ts">
import moment from "moment";

interface Props {
  getData?: Object;
}
defineProps<Props>();
function trainFormat(date:any){
  return moment(date).format('YYYY-MM-DD');
}
function saveMoneyTwoNum(num: any) {
  if ((num + '').indexOf('.') != -1) {
    if ((num + '').split(".")[1].length == 1) {
      num = num + '0'
    } else {
      num = num;
    }
  } else {
    num = Number(num).toFixed(2)
  }
  return num;
}
</script>

<style scoped>

</style>
