<template>
  <div class="approvalCard">
    <van-row justify="space-between" align="center">
      <van-row align="center">
        <van-col v-if="getData.deptName!=''">
          {{ getData.deptName }}
        </van-col>
        <van-col v-if="getData.xmmc!=''">
          {{ getData.xmmc }}
        </van-col>
      </van-row>
    </van-row>
    <div class=" mt5 mb5">
      <van-row v-if="getData.deptId">
        <van-col>
          <span class="f14 graytxt">部门编号：{{ getData.deptId }} </span>
        </van-col>
      </van-row>
      <van-row v-if="getData.xmbh">
        <van-col>
          <span class="f14 graytxt">项目编号：{{ getData.xmbh }} </span>
        </van-col>
      </van-row>
      <van-row>
        <van-col>
          <span class="f14 graytxt">
            经费额度：
            <span class="orange bold">
              ¥ {{ Number(getData.deptBudGetingAmount??getData.xmje).toFixed(2) }}
            </span>
          </span>
        </van-col>
      </van-row>
      <van-row>
        <van-col>
          <span class="f14 graytxt">
            经费余额：
            <span style="color:#0088ff" class="bold">
              ¥ {{ Number(getData.deptBudGetingYe??getData.xmye).toFixed(2) }}
            </span>
          </span>
        </van-col>
      </van-row>
      <van-row v-if="getData.subjectNameList!=''">
        <van-col>
          <span class="f14 graytxt">消费科目：{{ getData.subjectNameList }}</span>
        </van-col>
      </van-row>
      <van-row align="center">
        <van-col span="6"> <div class="typeTips">{{getData.budGetingYear}}</div></van-col>
        <van-col span="8" offset="1">
          <div class="deptTips" v-if="getData.deptId">部门成本中心</div>
          <div class="deptTips" v-if="getData.xmbh">项目成本中心</div>
        </van-col>
      </van-row>
    </div>
    <van-row>
      <span  style="color: #c9c9c9;font-size: 12px">{{ getData.creationTime??getData.creationtime }}</span>
    </van-row>
  </div>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";

interface Props {
  getData?: Object;
}
defineProps<Props>()
const router = useRouter();

</script>

<style scoped>
.approvalCard {
  margin: 10px;
  padding: 10px;
  border-radius: 10px;
  background: #ffffff;
}

.typeTips{
  padding: 5px;
  margin-top: 10px;
  border: 1px solid #0088ff;
  background: rgba(0,136,255,0.2);
  color:#0088ff;
  font-size:12px;
  text-align: center;
  border-radius: 10px;
}
.deptTips{
  padding: 5px;
  margin-top: 10px;
  border: 1px solid #46D2B1;
  background: rgba(70,210,177,0.2);
  color:#46D2B1;
  font-size:12px;
  text-align: center;
  border-radius: 10px;
}
</style>
