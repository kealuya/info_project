<template>
  <div>
    <div>
      <van-row align="center">
        <van-col span="6">
          <div class="graytxt"> 报销类型：</div>
        </van-col>
        <van-col span="18"> {{ type }}</van-col>
      </van-row>
    </div>
    <div class="mt5">
      <van-row align="center">
        <van-col span="6">
          <div class="graytxt"> 报销状态：</div>
        </van-col>
        <van-col span="18"> {{ data.isBx = 'false' ? '已填报' : '已报销' }}</van-col>
      </van-row>
    </div>
    <div class="mt5">
      <van-row align="center">
        <van-col span="6">
          <div class="graytxt"> 申请时间：</div>
        </van-col>
        <van-col span="18"> {{ data.createtime }}</van-col>
      </van-row>
    </div>
    <div class="mt5">
      <van-row align="center">
        <van-col span="6">
          <div class="graytxt"> 申请人员：</div>
        </van-col>
        <van-col span="18"> {{ data.empname }}</van-col>
      </van-row>
    </div>
    <div class="mt5">
      <van-row align="center">
        <van-col span="6">
          <div class="graytxt"> 申请部门：</div>
        </van-col>
        <van-col span="18"> {{ data.deptName }}</van-col>
      </van-row>
    </div>
    <div class="mt5">
      <van-row align="center">
        <van-col span="6">
          <div class="graytxt"> 成本中心：</div>
        </van-col>
        <van-col span="18"> {{ data.xmmc }}（{{ data.xmbh }}）</van-col>
      </van-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref} from "vue";
const props=defineProps<{
  data?:object;
  type:string;
}>()
</script>

<style scoped>

</style>
