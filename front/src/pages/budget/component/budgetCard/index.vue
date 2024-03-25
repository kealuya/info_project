<template>
  <div class="approvalCard" @click="goBudgetDetail(getData)">
    <van-row justify="space-between" align="center">
      <van-row  align="center">
        <van-col >
          {{getData.budGetingName}}
        </van-col>
      </van-row>
      <van-col>
        <div  class="waitTip" v-if="getData.approvalResult=='0'">等待审批</div>
        <div  class="seccessTip" v-if="getData.approvalResult=='1'">审批通过</div>
        <div  class="dangerTip" v-if="getData.approvalResult=='2'">审批拒绝</div>
      </van-col>
    </van-row>
    <div class=" mt5 mb5">
      <van-row>
        <van-col><span class="f14 graytxt">审定金额：<span class="orange">¥ {{Number(getData.budGetingAmount).toFixed(2)}}</span></span></van-col>
      </van-row>
      <van-row>
        <van-col><span class="f14 graytxt">预算编码：{{getData.budGetingId}}</span></van-col>
      </van-row>
      <van-row>
        <van-col><span class="f14 graytxt">负责人：{{getData.creation}}</span></van-col>
      </van-row>

    </div>
    <van-row>
      <span class="ml5" style="color: #c9c9c9;font-size: 12px">{{getData.creationTime}}</span>
    </van-row>
  </div>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
interface Props{
  getData?:Object;

}
defineProps<Props>()
const router = useRouter();
// 跳转至预算编制详情
const goBudgetDetail=(item:any)=>{
  const data :any ={
    'approveNo':item.budGetingId,
  };
  // 存储至预算编制
  sessionStorage.setItem('budgetDetail', JSON.stringify(data))
  router.push('/budgetDetail')
}
</script>

<style scoped>
.approvalCard{
  margin: 10px;
  padding: 10px;
  border-radius: 10px;
  background: #ffffff;
}
.waitTip{
  margin: 0px 5px;
  padding: 8px 10px;
  border-radius: 8px;
  background: #B6DBFF;
  color: #0080ff;
  font-size: 14px;
}
.seccessTip{
  margin: 5px 5px;
  padding: 8px 10px;
  border-radius: 12px;
  background: #ECFBF7;
  color: #46D2B1;
  font-size: 14px;
}
.dangerTip{
  margin: 10px 5px;
  padding: 8px 10px;
  border-radius: 12px;
  background: #F6DCEA;
  color: #DD6AA4;
  font-size: 14px;
}
</style>
