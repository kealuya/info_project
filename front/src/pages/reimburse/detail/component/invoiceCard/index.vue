<template>
  <div>
    <div class="invoiceInfoCard">
      <div class="invoiceCard" @click="getInvoiceDetail(data)">
        <van-row justify="space-between" align="center">
          <van-col>
            {{data.spendType}}
            <van-icon :name="xiaLaIcon" size="13"/>
          </van-col>
          <van-col>
            <div class="bold" style="color: #0088ff">¥ {{parseInt(data.spendAmount).toFixed(2)}}</div>
          </van-col>
        </van-row>
      </div>
     <div class="p10"  v-if="show">

       <van-row justify="space-between" v-if="data.spendType=='火车票'||data.spendType=='飞机票'">
         <van-col>
           <span class="f14 graytxt">出发地点</span>
         </van-col>
         <van-col>{{dataInfo.cfdd}}</van-col>
       </van-row>
       <van-row justify="space-between" v-if="data.spendType=='火车票'||data.spendType=='飞机票'">
         <van-col>
           <span class="f14 graytxt">出发时间</span>
         </van-col>
         <van-col>{{dataInfo.kprq}}</van-col>
       </van-row>
       <van-row justify="space-between" v-if="data.spendType=='火车票'||data.spendType=='飞机票'">
         <van-col>
           <span class="f14 graytxt">到达地点</span>
         </van-col>
         <van-col>{{dataInfo.dddd}}</van-col>
       </van-row>
        <van-row justify="space-between" >
         <van-col>
           <span class="f14 graytxt">发票代码</span>
         </van-col>
         <van-col>{{data.fpdm}}</van-col>
       </van-row>
       <van-row justify="space-between" v-if="data.spendType!='火车票'||data.spendType!='飞机票'">
         <van-col>
           <span class="f14 graytxt">发票号码</span>
         </van-col>
         <van-col>{{data.fphm}}</van-col>
       </van-row>
       <van-row justify="space-between">
         <van-col>
           <span class="f14 graytxt">金额</span>
         </van-col>
         <van-col>¥{{data.spendAmount}}</van-col>
       </van-row>
<!--       <van-row justify="space-between">-->
<!--         <van-col class="graytxt">发生日期</van-col>-->
<!--         <van-col class="graytxt">2022-07-08</van-col>-->
<!--       </van-row>-->
     </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import xiaLaIcon from '../../../../../assets/icon/xiala.png'
import {ref,inject} from "vue";
import {getInvoiceWbById} from "../../../../../services/invoice";
let userInfoData: any = inject("userInfo"); //获取审批数组
// 请求参数
const resData={
  "dzfph":"",
  "corpCode":userInfoData.userInfo.corpCode,
  "empCode":userInfoData.userInfo.empCode,
  "equipmentType":""
}
// 接收接口返回的发票详情
const dataInfo=ref({fpdm:''});
// 控制显示隐藏
let show = ref(false);
// 是否获取接口的数据
let dataTrue=ref(false);
// 获取发票详情
const getInvoiceDetail=async(item:any)=>{
  resData.dzfph=item.dzfph;
  let result: any = await getInvoiceWbById(resData);
  dataInfo.value=result.data.data[0]
  dataTrue.value=true;
  show.value=!show.value
};
interface Props{
  data?:object;
}
defineProps<Props>()
</script>

<style scoped>
.invoiceCard{
  margin-top: 5px;
  padding:10px;
  background: #EEF7FF;
  border-radius: 10px;
}
.invoiceInfoCard{
  margin-left: 5px;
  margin-right: 5px;
  margin-bottom: 5px;
  background: #F7FBFF;
  border-bottom-right-radius: 10px;
  border-bottom-left-radius: 10px;
}
</style>
