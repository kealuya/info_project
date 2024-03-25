<template>
  <div style="padding: 20px">
    <van-row gutter="20" align="center">
      <van-col span="8" @click="addDailyBx">
        <van-row align="center" justify="center">
          <van-image :src="daiyReimburse" width="60"  height="60" />
          <div class="f14 mt5" >日常报销</div>
        </van-row>
      </van-col>
      <van-col span="8">
        <van-row align="center" justify="center"  @click="showTravelBxStatus">
          <van-image :src="travelBx" width="60"  height="60" />
          <div class="f14 mt5" >差旅补助报销</div>
        </van-row>
      </van-col>
      <van-col span="8">
        <van-row align="center" justify="center" @click="showLoanStatus">
          <van-image :src="loanBx" width="60"  height="60" />
          <div class="f14 mt5">借款申请</div>
        </van-row>
      </van-col>
    </van-row>
  </div>
</template>

<script lang="ts" setup>
import daiyReimburse from '../../../../assets/icon/apply_icon_05.png'
import travelBx from '../../../../assets/icon/invoice_icon_14.png'
import loanBx from '../../../../assets/icon/apply_icon_06.png'
import {useRouter} from "vue-router";
const router = useRouter();

function addDailyBx() {
  // 因为日常报销存在
  const data :any ={
    'bxType':'rcbx',
    'approveNo':'',
    'isBx':'-'
  };
  sessionStorage.setItem('reimburseDetail', JSON.stringify(data));
  sessionStorage.setItem('billList', JSON.stringify({'billId':''}));
  sessionStorage.setItem('invoiceList', JSON.stringify([]));
  router.push({
    name: 'addDailyBx',
  })
}

function showLoanStatus() {
  router.push({
    name: 'addLoan', params: {
      selectAccountInformation: JSON.stringify("")
    }
  })
}

function showTravelBxStatus() {
  router.push({
    name: 'addTravelSubsidies'
  })
}


</script>

<style scoped>

</style>
