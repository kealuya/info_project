<template>
  <div>
    <div>
      <img src="../../../../assets/img/yiyanzhen.png" class="yzImg" v-if="getData.isyz=='1'"/>
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
      <van-row class="mt10">
        <van-col span="6"  :offset="2" class="graytxt">开票日期：</van-col>
        <van-col span="10" class="detail">{{getData.kprq? getData.kprq:'---'}}</van-col>
      </van-row>
      <van-row class="mt10">
        <van-col span="6" :offset="2" class="graytxt">发票金额：</van-col>
        <van-col span="6" class="detail" style="color: #C6666A">￥{{ saveMoneyTwoNum(getData.zje) }}</van-col>
      </van-row>
    </div>
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
.yzImg{
  width:130px;
  height:100px;
  position: absolute;
  top: 60px;
  right:20px;
}
</style>
