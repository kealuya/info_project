<template>
  <div class="reimburseCard" @click="reimburiseDetail(getData)">
    <van-row justify="space-between" align="center">
      <van-col span="17">
        <van-row  align="center">
          <van-col span="8">
            <div class="travelType" v-if="getData.bxType=='rcbx'" >日常报销</div>
            <div class="bxType"  v-if="getData.bxType=='travelbx'">差旅补助</div>
            <div class="loanType"  v-if="getData.bxType=='jkbx'">借款申请</div>
          </van-col>
          <van-col span="16" class="textEllipsis">
            <span class="bold ml5 ">{{getData.applyReason}}</span>
          </van-col>
        </van-row>
      </van-col>
      <van-col span="7" style="text-align: right">
        <span class="mr5 red bold f18 tr">¥ {{moneyFormat(parseFloat(getData.bxAmount).toFixed(2))}}</span></van-col>
    </van-row>
    <van-row justify="space-between" align="center">
      <van-col>
          <div class="ml5">
            <div class="graytxt f12">业务编号：{{getData.applyNo}}</div>
            <div class="graytxt f12">创建时间：{{getData.createTime}}</div>
          </div>
      </van-col>
      <van-col>
        <div class="notSubmitedTip" v-if="getData.approvalState=='0'">待审核</div>
        <div class="sucesssTip" v-if="getData.approvalState=='1'">已通过</div>
        <div class="waitTip" v-if="getData.approvalState=='-'">未提交</div>
        <div class="errorTip" v-if="getData.approvalState=='2'">已拒绝</div>
        <div class="notTip" v-if="getData.approvalState==null">未知状态</div>
      </van-col>
    </van-row>
  </div>
</template>

<script lang="ts" setup>
import {useRouter} from "vue-router";
const router = useRouter();
interface Props{
  getData?:Object;
}
defineProps<Props>()
// 规范金额显示添加千分位
const moneyFormat=(item:any)=>{
  return item.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",")
}
// 根据类型跳转到对应的报销单当中
const reimburiseDetail=(item:any)=>{
  // 当报销单类型是日常报销时
  if(item.bxType=='rcbx'){
    const data :any ={
      'bxType':item.bxType,
      'approveNo':item.applyNo,
      'isBx':item.isBx
    };
    sessionStorage.setItem('reimburseDetail', JSON.stringify(data))
    if(item.approvalState=='-'){
      router.push('/addDailyBx');
    }else{
      router.push('/reimburse/detail');
    }
  }else if(item.bxType=='travelbx'){ // 当报销单类型是差旅补助报销时
    router.push({
      name: 'travelSubsidiesDetail', params: {
        applyNo: item.applyNo
      }
    })

  }else{ // 当报销类型是借款申请时
    router.push({
      name: 'loanDetail', params: {
        loanId: item.applyNo
      }
    })
  }
}

</script>

<style scoped>
.bxType{
  margin: 5px;
  padding: 8px;
  font-size: 12px;
  color:#ffffff;
  background: rgba(0, 128, 255, 0.7);
  border-radius: 10px;
}
.travelType{
  margin: 5px;
  padding: 8px;
  font-size: 12px;
  color:#ffffff;
  background: rgba(250, 97, 49, 0.7);
  border-radius: 10px;
}
.loanType{
  margin: 5px;
  padding: 8px;
  font-size: 12px;
  color:#ffffff;
  background: rgba(245, 166, 35, 0.7);
  border-radius: 10px;
}
.reimburseCard{
  margin: 10px;
  padding:5px;
  border-radius: 10px;
  background: #ffffff;
}
.notSubmitedTip{
  width:80px;
  height: 25px;
  text-align: center;
  padding-top: 2px;
  border-radius: 10px;
  background: rgba(251, 192, 45, 0.15);
  border: 1px solid rgba(251, 192, 45, 0.6);
  color:#FBC02D;
}
.waitTip{
  width:80px;
  height: 25px;
  text-align: center;
  border-radius: 10px;
  padding-top: 2px;
  background: rgba(182, 219, 255, 0.2);
  color:#0080FF;
  border:1px solid rgba(182, 219, 255, 0.6);
}
.sucesssTip{
  width:80px;
  height: 25px;
  text-align: center;
  padding-top: 2px;
  color:#46D2B1;
  background: rgba(70, 210, 177, 0.2);
  border:1px solid rgba(70, 210, 177, 0.6);
  border-radius: 10px;
}
.errorTip{
  width:80px;
  height: 25px;
  text-align: center;
  padding-top: 2px;
  border-radius: 10px;
  border:1px solid rgba(221, 106, 164, 0.6);
  background: rgba(221, 106, 164, 0.1);
  color:#DD6AA4;
}
.notTip{
  width:80px;
  height: 25px;
  padding-top: 2px;
  text-align: center;
  border-radius: 10px;
  background: #696969;
  color:#C0C0C0;
}
.textEllipsis{
  white-space: nowrap;/*设置不换行*/
  overflow: hidden; /*设置隐藏*/
  text-overflow:ellipsis;
}
</style>
