<template>
  <div class="card">
    <div @click="checkInvoice(getData)">
      <van-row justify="start" align="center">
        <van-col span="5" style="color:rgba(23,26,29,0.40);">
          <div class="invoiceTip"  :class="{'zt1':getData.zt!='1'} " style="">
            <span style="margin-left: 10px">{{getInvoiceTypeList(getData.type)}}</span>
          </div>
        </van-col>
        <van-col span="17" class="title f16 pl25" style="line-height: 25px;">{{getData.invName}}</van-col>
        <van-col span="2">
          <van-checkbox v-if="getData.zt!='1'"  label-disabled v-model="getData.checked" @click="checkInvoice(getData)"></van-checkbox>
        </van-col>
      </van-row>
    </div>

    <div class="mt5 f14">
      <div class="line"></div>
    </div>
    <div  @click="goInvoiceDetail(getData)">
      <van-row  justify="space-between" align="center" class="mt5">
        <van-col span="6" class="f14 graytxt">发票号码</van-col>
        <van-col span="14" class="f14 tr mr5">  {{getData.fphm}}</van-col>
      </van-row>
      <van-row  justify="space-between" align="center" class="mt5">
        <van-col span="6" class="f14 graytxt">发票代码</van-col>
        <van-col span="14" class="f14 tr mr5">  {{getData.fpdm}}</van-col>
      </van-row>
      <van-row  justify="space-between" align="center" class="mt5 f14" >
        <van-col span="6" class="f14 graytxt" >创建时间</van-col>
        <van-col span="14" class="f14 tr mr5">{{moment(getData.cjrq).format("YYYY-MM-DD")}}</van-col>
      </van-row>
      <van-row  justify="space-between" align="center" class="mt5">
        <van-col span="6" class="f14 graytxt">发票金额</van-col>
        <van-col span="14" class="tr mr5 money bold" style="font-size: 16px">¥ {{saveMoneyTwoNum(getData.zje)}}</van-col>
      </van-row>
      <img src="../../../../assets/img/yiyanzhen.png" class="yzImg" v-if="getData.isyz=='1'"/>
    </div>
  </div>
</template>

<script  setup lang="ts">
import {invoiceTypeList}from './ts/invoiceTypeList'
import {useRoute, useRouter} from "vue-router";
import moment from "moment";
import {  defineEmits} from "vue";
const emit=defineEmits(['selectInvoice']); // emit事件方法
const router = useRouter();
interface Props{
  getData?:Object;
}
defineProps<Props>();
var invoiceTitle='';
function getInvoiceTypeList (data: any) {
  let invoiceType=invoiceTypeList.find((e)=>data==e.value);
    invoiceTitle=invoiceType==undefined?'':invoiceType.title;
    return invoiceTitle;
}
function saveMoneyTwoNum(num:any){
  if((num+'').indexOf('.') != -1){
    if((num+'').split(".")[1].length == 1){
      num = num+'0'
    }else{
      num = num;
    }
  }else{
    num = Number(num).toFixed(2)
  }
  return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}
//点击跳转到发票详情页
const goInvoiceDetail = (item: any) => {
  sessionStorage.setItem('invoiceDetail', JSON.stringify(item))
  router.push('/detail');
};
//选中票据传递给上一级父组件
const checkInvoice=(data:any)=>{
  data.checked=!data.checked;
  emit('selectInvoice',data);
}
</script>

<style scoped>
  .card{
    position: relative;
  }
  .zt1{
    color: #4BA3FB;
  }
  .invoiceTip{
    padding: 3px 0px 3px 5px;
    height: 20px;
    font-size: 14px;
    font-weight: bold;
    width:90px;
    text-align: left;
    line-height: 20px;
    background: url('../../../../assets/img/fenlei_bg.png') no-repeat center ;
    background-size: 90px 40px;
    word-break: keep-all;
    position: absolute;
    left: -3px;
    top: 10px;
  }
  .yzImg{
    width:95px;
    height:75px;
    position: absolute;
    bottom: 10px;
    right:62px;
  }
</style>
