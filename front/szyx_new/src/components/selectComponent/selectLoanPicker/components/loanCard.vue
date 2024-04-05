<template>
  <div class="costCenterCard" >
    <van-row justify="space-between" align="center">
      <van-row  align="center">
        <van-col >
          <van-image width="30" height="30" :src="jksqIcon"/>
        </van-col>
        <van-col>
          <span class="f14 ml10">借款申请</span>
        </van-col>
      </van-row>
      <van-col>
        <van-icon :name="CheckedIcon" v-if="checkKey==getData.applyno" />
      </van-col>
    </van-row>
    <div class="ml10 mt5 mb5">
      <van-row>
        <van-col><span class="f14 graytxt">申请事由：{{getData.applyreason}}</span></van-col>
      </van-row>
      <van-row>
        <van-col><span class="f14 graytxt">申请人：{{getData.payeeName}}</span></van-col>
      </van-row>
      <van-row>
        <van-col><span class="f14 graytxt">申请单号：{{getData.applyno}}</span></van-col>
      </van-row>
    </div>
    <van-row>
      <span class="ml5" style="color: #c9c9c9;font-size: 12px">{{getData.createtime}}</span>
    </van-row>
  </div>
</template>

<script lang="ts" setup>
import iconSelectCost from '../../../../assets/icon/icon_loan_Checked.png'
import  trainIcon from '../../../../assets/icon/application_icon_01.png';
import  rcbxIcon from '../../../../assets/icon/application_icon_11.png';
import  jksqIcon from '../../../../assets/icon/application_icon_13.png';
import  clbxIcon from '../../../../assets/icon/application_icon_12.png';
import  otherIcon from '../../../../assets/icon/application_icon_12.png';
import  travelBxIcon from '../../../../assets/icon/application_icon_13.png';
import  leaveIcon from '../../../../assets/icon/application_icon_01.png';
import  workIcon from '../../../../assets/icon/application_icon_02.png';
import  costIcon from '../../../../assets/icon/application_icon_14.png';
import  userSealIcon from '../../../../assets/icon/application_icon_07.png';
import  userCarIcon from '../../../../assets/icon/application_icon_15.png';
import  subscriptionIcon from '../../../../assets/icon/application_icon_06.png';
import  DSQpIcon from '../../../../assets/icon/application_icon_11.png';
import  cardIcon from '../../../../assets/icon/application_icon_04.png';
import  itemCollectionIcon from '../../../../assets/icon/application_icon_05.png';
import  jiJinBaoXiaoIcon from '../../../../assets/icon/application_icon_10.png';
import  entertainIcon from '../../../../assets/icon/application_icon_08.png';
import  meetingIcon from '../../../../assets/icon/application_icon_09.png';
import  outWorkIcon from '../../../../assets/icon/application_icon_03.png';
import CheckedIcon from '../../../../assets/icon/checked.png'
import  logIcon from '../../../../assets/icon/application_icon_05.png';
import {applyType} from "../../../../pages/home/component/ts/applyType";
interface Props{};
const props=defineProps<{
  getData:{};
  checkKey:String;
}>()
var approvalTypeTitle='';
function getApprovalTypeList (data: any) {
  let approvalType=applyType.find((e)=>data==e.value);
  approvalTypeTitle=approvalType==undefined?'':approvalType.text;
  return approvalTypeTitle;
}
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
