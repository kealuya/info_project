<template>
  <div>
    <div>
      <van-row align="center">
        <van-col span="6" :offset="2" class="graytxt">发票代码：</van-col>
        <van-col span="6" class="detail">{{ getData.fpdm??'----' }}</van-col>
      </van-row>
    </div>
    <div class="mt10">
      <van-row align="center">
        <van-col span="6" :offset="2" class="graytxt">发票号码：</van-col>
        <van-col span="6" class="detail">{{ getData.fphm??'----' }}</van-col>
      </van-row>
    </div>
    <div class="mt5 f14">
      <div class="line2"></div>
    </div>
    <div class="mt10">
      <van-row align="center">
        <van-col span="6" :offset="2" class="graytxt">创建时间：</van-col>
        <van-col span="10" class="detail">{{getData.cjrq.substr(0,10)??'----'}}</van-col>
      </van-row>
    </div>
    <van-row class="mt10">
      <van-col span="6" :offset="2" class="graytxt">发票金额：</van-col>
      <van-col span="6" class="detail" style="color: #C6666A">￥{{ saveMoneyTwoNum(getData.zje) }}</van-col>
    </van-row>
  </div>
</template>

<script setup lang="ts">
interface Props {
  getData?: Object;
}

defineProps<Props>();

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
