<template>
  <div>
    <div>
      <div class="ml10">{{title}}</div>
      <div class="ml5 mt5 p5">
        <van-row justify="space-between" align="bottom">
          <van-col>
            <div class="mt5">
              <div class="graytxt">申请单号:</div>
              <div>{{data.applyno}}</div>
            </div>
            <div  class="mt5">
              <div class="graytxt">报销事由:</div>
              <div>{{data.applyreason}}</div>
            </div>
            <div  class="mt5">
              <div class="graytxt">报销总额:</div>
              <div class="f18 bold red">¥ {{data.bxamount}}</div>
            </div>
          </van-col>
          <van-col>
            <div class="tr mr10">
              <status-icon  :type="data.state" />
            </div>
          </van-col>
        </van-row>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import StatusIcon from '../statusIcon/index.vue';
interface Props{
  data?:object;
  title?:string;
}
defineProps<Props>()
// 根据类型返回
const titleKey=(code:any)=>{
  if(code=='rcbx'){
    return '日常报销单'
  }else if(code=='travelbx'){
    return '差旅补助报销单'
  }else{
    return '借款申请单'
  }
}
</script>

<style scoped>
</style>
