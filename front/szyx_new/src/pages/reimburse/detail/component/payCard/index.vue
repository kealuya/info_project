<template>
<div>
  <div class="f18">收款信息</div>
  <div class="mt10 mb10">
    <div class="mt5">
      <van-row>
        <van-col>
          <div class="graytxt"> 账户类型：</div>
        </van-col>
        <van-col> {{data.accountType=='personalAccount'?'个人账户':'对公账户'}} </van-col>
      </van-row>
    </div>
    <div class="mt5">
      <van-row>
        <van-col>
          <div class="graytxt"> 开户名称：</div>
        </van-col>
        <van-col>
          {{data.bankName}}
        </van-col>
      </van-row>
    </div>
    <div class="mt5">
      <van-row>
        <van-col>
          <div class="graytxt"> 银行卡号：</div>
        </van-col>
        <van-col>
          {{data.bankNum}}
        </van-col>
      </van-row>
    </div>
    <div class="mt5">
      <van-row>
        <van-col>
          <div class="graytxt"> 收款名称：</div>
        </van-col>
        <van-col>
          {{data.payeeName}}
        </van-col>
      </van-row>
    </div>
    <div class="mt5">
      <van-row>
        <van-col>
          <div class="graytxt"> 开户网点：</div>
        </van-col>
        <van-col>
          {{data.accountNetwork}}
        </van-col>
      </van-row>
    </div>
  </div>
</div>
</template>

<script lang="ts" setup>
interface Props{
  data?:object;

}
defineProps<Props>()
</script>

<style scoped>

</style>
