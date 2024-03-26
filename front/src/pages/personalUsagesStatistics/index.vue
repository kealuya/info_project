<template>
  <nav-bar title="个人使用统计" />
  <div>
    <van-row justify="space-around">
      <van-col>
        <div :class="eChat=='trip'?'checkCard':'noCheckCard'" @click="chatToggle('trip')">
          <p class="f16 bold">{{tripAmount}}元/{{tripNumber}}次</p>
          <p class="f14">累计出行</p>
        </div>
      </van-col>
      <van-col>
        <div  :class="eChat=='app'?'checkCard':'noCheckCard'" @click="chatToggle('app')">
          <p class="f16 bold">{{appNumber}}单</p>
          <p class="f14">累计申请</p>
        </div>
      </van-col>
      <van-col>
        <div :class="eChat=='bxMoney'?'checkCard':'noCheckCard'" @click="chatToggle('bxMoney')">
          <p class="f16 bold">{{parseInt(bxMoneyAmount)}}元</p>
          <p class="f14">累计报销</p>
        </div>
      </van-col>
    </van-row>
    <div class="mt10" v-if="eChat=='trip'">
      <travel-chat-component ref="travelChat" @over-data="overTrip" />
    </div>
    <div class="mt10" v-else-if="eChat=='app'">
      <!--  展示累计申请柱状图    -->
      <appr-chat-component ref="apprChat" @over-data="overCount" />
    </div>
    <div v-else>
      <!--   展示累计报销折线图   -->
      <bx-amount-chat-component ref="bxAmountChat" @over-data="overMoney" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import NavBar from '../../components/navBar/index.vue';
import ApprChatComponent from './component/apprChatComponent/index.vue';
import TravelChatComponent from './component/travelChatComponent/index.vue';
import BxAmountChatComponent from  './component/bxAmountChatComponent/index.vue'
import {ref, onBeforeMount, onUnmounted} from "vue";
let eChat=ref('trip');
let tripNumber=ref('0'); // 累计出行次数
let tripAmount=ref('0'); // 累计出行金额
let appNumber=ref('0'); // 累计申请次数
let bxMoneyAmount=ref('0'); // 报销金额
let apprChat=ref(); // 注册自组件的名
// 切换图表展示
const chatToggle=(code:any)=>{
  eChat.value=code;
}
// 将子组件-累计申请柱状图从接口获取的总数赋值至上方展示
const overTrip = (money: any,  count:any) => {
  tripNumber.value=count;
  tripAmount.value=money;
}
// 将子组件-累计申请柱状图从接口获取的总数赋值至上方展示
const overCount = (count: any) => {
  appNumber.value=count
}
// 将子组件-累计报销金额年单位折线图从接口获取的总数赋值至上方展示
const overMoney = (money: any) => {
  bxMoneyAmount.value=money
}
onBeforeMount(()=>{
  let personUserChat:any=sessionStorage.getItem('personUserChat');
  let chatInfo=JSON.parse(personUserChat);
  eChat.value=chatInfo.code;
  tripAmount.value=chatInfo.travelMoney;
  tripNumber.value=chatInfo.tripNumber;
  bxMoneyAmount.value=chatInfo.bxMoneyAmount;
  appNumber.value=chatInfo.appNumber;
})
onUnmounted(()=>{
  sessionStorage.removeItem('personUserChat');
})
</script>

<style scoped>
  .checkCard{
    padding:10px 20px;
    margin:10px;
    color:#ffffff;
    text-align:center;
    background: #0088ff;
    border-radius: 10px;
  }
  .noCheckCard{
    padding:10px 20px;
    margin:10px;
    color:#ffffff;
    text-align:center;
    background: rgba(0,136,255,0.2);
    border-radius: 10px;
  }
</style>
