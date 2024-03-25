<template>
  <div class="approvalCard" @click="goDetail(getData)">
    <van-row justify="space-between" align="center">
     <van-row  align="center">
       <van-col >
         <van-image width="30" height="30" v-if="getData.approveType=='travel'" :src="travelBxIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='leave'" :src="leaveIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='work'" :src="workIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='reimbursement'" :src="costIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='seal'" :src="userSealIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='car'" :src="userCarIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='shop'" :src="subscriptionIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='sign'" :src="DSQpIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='card'" :src="cardIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='itemCollection'" :src="itemCollectionIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='fund'" :src="jiJinBaoXiaoIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='train'" :src="trainIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='entertain'" :src="entertainIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='meeting'" :src="meetingIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='outWork'" :src="outWorkIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='log'" :src="logIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='rcbx'" :src="rcbxIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='jksq'" :src="jksqIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='clbx'" :src="clbxIcon"/>
         <van-image width="30" height="30" v-if="getData.approveType=='otherIcon'" :src="otherIcon"/>
       </van-col>
       <van-col>
         <span class="f14 ml10">{{getApprovalTypeList(getData.approveType)}}</span>
       </van-col>
     </van-row>
      <van-col>
        <div  class="waitTip" v-if="getData.approvalState=='0'">等待审批</div>
        <div  class="seccessTip" v-else-if="getData.approvalState=='1'">审批通过</div>
        <div  class="dangerTip" v-else-if="getData.approvalState=='2'">审批拒绝</div>
        <div  class="warnTip" v-else-if="getData.approvalState=='3'">已撤销</div>
        <div  class="seccessTip" v-else-if="getData.approvalState=='4'">免审</div>
        <div  class="dangerTip" v-else>审批拒绝</div>
      </van-col>
    </van-row>
    <div class="ml10 mt5 mb5">
      <van-row>
        <van-col  class="textEllipsis"><span class="f14 graytxt">申请事由：{{getData.approveReason}}</span></van-col>
      </van-row>
      <van-row>
        <van-col><span class="f14 graytxt">申请人：{{getData.empName}}</span></van-col>
      </van-row>
      <van-row>
        <van-col><span class="f14 graytxt">申请单号：{{getData.approveNo}}</span></van-col>
      </van-row>
    </div>
    <van-row>
      <span class="ml5" style="color: #c9c9c9;font-size: 12px">{{getData.createTime}}</span>
    </van-row>
  </div>
</template>

<script setup lang="ts">
import {applyType} from "./ts/applyType";
import  trainIcon from '../../../assets/icon/application_icon_01.png';
import  rcbxIcon from '../../../assets/icon/application_icon_11.png';
import  jksqIcon from '../../../assets/icon/application_icon_13.png';
import  clbxIcon from '../../../assets/icon/application_icon_12.png';
import  otherIcon from '../../../assets/icon/application_icon_12.png';
import  travelBxIcon from '../../../assets/icon/application_icon_13.png';
import  leaveIcon from '../../../assets/icon/application_icon_01.png';
import  workIcon from '../../../assets/icon/application_icon_02.png';
import  costIcon from '../../../assets/icon/application_icon_14.png';
import  userSealIcon from '../../../assets/icon/application_icon_07.png';
import  userCarIcon from '../../../assets/icon/application_icon_15.png';
import  subscriptionIcon from '../../../assets/icon/application_icon_06.png';
import  DSQpIcon from '../../../assets/icon/application_icon_11.png';
import  cardIcon from '../../../assets/icon/application_icon_04.png';
import  itemCollectionIcon from '../../../assets/icon/application_icon_05.png';
import  jiJinBaoXiaoIcon from '../../../assets/icon/application_icon_10.png';
import  entertainIcon from '../../../assets/icon/application_icon_08.png';
import  meetingIcon from '../../../assets/icon/application_icon_09.png';
import  outWorkIcon from '../../../assets/icon/application_icon_03.png';
import  logIcon from '../../../assets/icon/application_icon_05.png';
import {useRouter} from "vue-router";
interface Props{
  getData:{
    [propName: string]: string
  };

}
defineProps<Props>()
const router = useRouter();
const goDetail=(item:any)=>{
  // 根据单据类型跳转至不同的页面
  if(item.approveType=="rcbx"){
    const data :any ={
      'bxType':'rcbx',
      'approveNo':item.approveNo,
      'isBx':item.isBx
    };
    sessionStorage.setItem('reimburseDetail', JSON.stringify(data))
    if(item.approvalState=='-'){
      router.push('/addDailyBx');
    }else{
      router.push('/reimburse/detail');
    }
  }else if(item.approveType=="clbx"){
    router.push({
      name:'travelSubsidiesDetail',
      params:{
        applyNo:item.approveNo
      }
    })
  }else if(item.approveType == "jksq"){
    router.push({
      name:'loanDetail',
      params:{
        loanId:item.approveNo
      }
    })

  }else{
    const data :any ={
      'approveType':item.approveType,
      'approveNo':item.approveNo,
    };
    sessionStorage.setItem('applicationDetail', JSON.stringify(data))
    router.push('/applicationDetail')
  }

}
var approvalTypeTitle='';
function getApprovalTypeList (data: any) {
  let approvalType=applyType.find((e)=>data==e.value);
  approvalTypeTitle=approvalType==undefined?'':approvalType.text;
  return approvalTypeTitle;
}
</script>

<style scoped>
.textEllipsis{
  white-space: nowrap;/*设置不换行*/
  overflow: hidden; /*设置隐藏*/
  text-overflow:ellipsis;
}
.approvalCard{
  margin: 10px;
  padding: 10px;
  border-radius: 10px;
  background: #ffffff;
}
.waitTip{
  margin: 5px 5px;
  padding: 8px 10px;
  border-radius: 8px;
  background: #B6DBFF;
  color: #0080ff;
  font-size: 14px;
}
.seccessTip{
  margin: 10px 5px;
  padding: 8px 10px;
  border-radius: 12px;
  background: #ECFBF7;
  color: #46D2B1;
  font-size: 14px;
}
.dangerTip{
  margin: 10px 5px;
  padding: 8px 10px;
  border-radius: 12px;
  background: #F6DCEA;
  color: #DD6AA4;
  font-size: 14px;
}
.warnTip{
  margin: 5px 5px;
  padding: 8px 10px;
  border-radius: 8px;
  background: rgba(255,146,0,0.48);
  color: #FF9200;
  font-size: 14px;
}
</style>
