<template>
  <div class="costCenterCard" >
    <van-row justify="space-between">
      <van-col>
        <van-row>
          <van-col>{{item.xmmc}}</van-col>
        </van-row>
        <van-row>
          <van-col>
            <span class="graytxt f12"> {{item.xmbh}}</span>
          </van-col>
        </van-row>
        <div class="mt5 mb5">
          <van-row>
            <van-col>
              <span class="graytxt f12">经费额度：</span>
            </van-col>
            <van-col>
              <span style="color:#F5A623" class="f12"> ¥ {{item.xmje}}</span>
            </van-col>
          </van-row>
          <van-row>
            <van-col>
              <span class="graytxt f12">经费余额：</span>
            </van-col>
            <van-col>
              <span style="color:#0080FF" class="f12"> ¥ {{item.xmye}}</span>
            </van-col>
          </van-row>
        </div>
        <div class="">
          <van-row justify="space-between">
            <van-row>
              <van-col >
                <div class='yearTip'>{{item.budGetingYear}}</div>
              </van-col>
              <van-col  class="ml10" v-if="item.costType=='department'">
                <div class="deptTip f12">部门成本中心</div>
              </van-col>
              <van-col class="ml10" v-else>
                <div class="projectTip f12">项目成本中心</div>
              </van-col>
            </van-row>
          </van-row>
        </div>
      </van-col>
      <van-col>
        <van-row align="center">
          <div class="checked" v-if="checkKey==item.xmbh">
            <van-image :src="iconSelectCost" width="30px" height="30px"  />
          </div>
        </van-row>
      </van-col>
    </van-row>
  </div>
</template>

<script lang="ts" setup>
import iconSelectCost from '../../../../assets/icon/icon_loan_Checked.png'
interface Props{};
const props=defineProps<{
  item:{
    xmmc:'',
    xmbh:'',
    budGetingYear:'',
    costType:''
  };
  checkKey:String;
}>()
</script>

<style scoped>
.costCenterCard{
  margin: 10px;
  padding: 10px;
  background: #ffffff;
  border-radius: 10px;
}
.yearTip{
  margin: 5px 0px;
  padding:5px 10px;
  text-align:center;
  font-size: 12px;
  border-radius: 5px;
  border: 1px solid #0088ff;
  color: #0088ff;
  background: #EFF8FF;
}
.deptTip{
  margin: 5px 0px;
  padding:5px 10px;
  text-align:center;
  font-size: 12px;
  border-radius: 5px;
  border: 1px solid #A2E8D7;
  color: #A2E8D7;
  background: #DAF6EF;
}
.projectTip{
  margin: 5px 0px;
  padding:5px 10px;
  text-align:center;
  font-size: 12px;
  border-radius: 5px;
  border: 1px solid #FCCD59;
  color: #FCCD59;
  background: #FEF6E0;
}
</style>
