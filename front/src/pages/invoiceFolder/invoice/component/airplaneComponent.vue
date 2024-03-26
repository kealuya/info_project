<template>
  <div class="card">
    <div  @click="checkInvoice(getData)">
      <van-row justify="space-between" align="center">
        <van-col span="5" style="color:rgba(23,26,29,0.40);">
          <div class="airtip"  :class="{'zt1':getData.zt!='1'} " style="">
            <span style="margin-left: 10px">{{getInvoiceTypeList(getData.type)}}</span>
          </div>
        </van-col>
        <van-col span="2">
          <van-checkbox v-if="getData.zt!='1'"  label-disabled v-model="getData.checked" @click="checkInvoice(getData)"></van-checkbox>
        </van-col>
      </van-row>
      <van-row  justify="space-around" align="center">
        <van-col span="7" class="f14 ml5 bold title" style="font-size: 15px;">
          {{getData.cfdd}}
        </van-col>
        <van-col span="7" class="f12 tr">
          <div class="money f16 tc" style="position: relative;top:8px;">¥ {{saveMoneyTwoNum(getData.zje)}}</div>
          <div style="position: relative;top:0px;left: 0px">
            <img :src="arrowRight" width="90" height="10" alt="">
          </div>
          <div class="ml10">
            <van-row>
              <van-col span="5" class="f12 tr mr5">
                <img :src="plane" width="20" height="20" alt="" >
              </van-col>
              <van-col span="7" class="f10 tr mr5 graytxt">
                {{getData.cc}}
              </van-col>
            </van-row>
          </div>
        </van-col>
        <van-col span="7" class="f16 tr  mr5 bold title" >  {{getData.dddd}}</van-col>
      </van-row>
    </div>
    <div class="mt5 f14">
      <div class="line"></div>
    </div>
    <div @click="goInvoiceDetail(getData)">
      <van-row  justify="space-between" align="center" class="mt10">
        <van-col span="4" class="f14 bold">{{getData.xm}}</van-col>
        <van-col span="18" class="f14 tr bold "> {{getData.fphm}}</van-col>
      </van-row>
      <van-row  justify="space-between"  class="mt5">
        <van-col span="12" class="f10   graytxt">起飞时间: {{trainFormat(getData.cfsj)}}</van-col>
        <van-col span="12" class="f14 pl15  tr graytxt">创建时间: {{trainFormat(getData.cjrq)}}</van-col>
      </van-row>
    </div>
  </div>
</template>

<script  setup lang="ts">
import arrowRight from '../../../../assets/img/icon_arrow_right.png';
import plane from '../../../../assets/img/planeNumber.png';
import {invoiceTypeList} from "./ts/invoiceTypeList";
import moment from "moment";
import {useRoute, useRouter} from "vue-router";
import {  defineEmits} from "vue";
const emit=defineEmits(['selectInvoice']); // emit事件方法
const router = useRouter();
interface Props{
  getData?:Object;

}
defineProps<Props>()
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
function trainFormat(date:any){
  return date.substr(0,10);
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
.zt1{
  color: #4BA3FB;
}
.airtip{
  padding: 3px 0px 3px 5px;
  height: 20px;
  font-size: 14px;
  font-weight: bold;
  width:90px;
  text-align: left;
  line-height: 20px;
  background: url('../../../../assets/img/fenlei_bg.png') no-repeat center ;
  background-size: 90px 40px;
  /*word-break: keep-all;*/
  position: relative;
  left: -13px;
  /*top: 10px;*/
}
</style>
