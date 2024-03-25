<template>
<div>
  <div class="formCard">
    <div class="pl15">
      <van-row align="center">
        <van-col :span="3">
          <application-icon :type="getData.approveType" />
        </van-col>
        <van-col :span="18">
          {{getData.approveName}}
        </van-col>
      </van-row>
      <van-row class="mb5">
        <van-col :span="16">
          <div class="f14 graytxt mt5">申请单号</div>
          <div class="f16 ">{{getData.approveNo}}</div>
          <div class="f14 graytxt mt5">申请时间</div>
          <div class="f16 mb10">{{getData.createTime}}</div>
        </van-col>
        <van-col :span="7">
          <div class="tr">
            <approval-status  :type="getData.approvalState" />
          </div>
        </van-col>
      </van-row>
    </div>
  </div>
</div>
</template>

<script lang="ts" setup>
import  ApplicationIcon from '../../../../../components/applicationIcon/index.vue'
import  ApprovalStatus from '../../../../../components/approvalStatus/index.vue'
interface Props{
  getData?:Object;

}
defineProps<Props>()
</script>

<style scoped>
.formCard {
  margin: 10px 5px;
  padding-top: 10px;
  padding-bottom: 5px;
  background: white;
  border-radius: 10px;
}
</style>
