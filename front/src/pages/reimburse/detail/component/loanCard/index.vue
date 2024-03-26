<template>
  <div>
    <div class="f18">请款申请</div>
    <van-row align="center" justify="space-between">
      <van-col>
        {{data.xmmc}}
      </van-col>
      <van-col>
        <div class="orange">¥{{data.loanamount}}</div>
      </van-col>
    </van-row>
    <div class="mt5 f14">
      <van-row align="center">
        <div>借款事由：{{data.applyreason}}</div>
      </van-row>
    </div>
    <div class="mt5 f14">
      <van-row align="center">
        <div>预计还款时间：{{data.repaymenttime}}</div>
      </van-row>
    </div>
    <div class="mt5 f14">
      <van-row align="center">
        <div>创建时间：{{data.createtime}}</div>
      </van-row>
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
