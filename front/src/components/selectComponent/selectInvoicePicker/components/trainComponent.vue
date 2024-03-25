<template>
  <div class="card">
    <van-row justify="space-between" align="center">
      <van-col span="5" style="color:rgba(23,26,29,0.40);">
        <div class="tip"  :class="{'zt1':getData.zt!='1'} " style="">
          <span style="margin-left: 10px">{{getInvoiceTypeList(getData.type)}}</span>
        </div>
      </van-col>
      <van-col>
        <van-checkbox v-model="getData.checked"></van-checkbox>
      </van-col>
    </van-row>
    <van-row  justify="space-around" align="center" class="mb5">
      <van-col span="6" class="f15 ml5 bold title" style="font-size: 15px">
        {{getData.cfdd}}
      </van-col>
      <van-col span="6" class="f12 tr mr5">
        <div class="money f16" style="position: relative;top:8px;text-align: center">¥ {{saveMoneyTwoNum(getData.zje)}}</div>
        <div style="position: relative;top:0px;left: 5px">
          <img :src="arrowRight" width="80" height="10" alt="">
        </div>
        <div class="ml10">
          <van-row>
            <van-col span="5" class="f12 tr">
              <img :src="train" width="20" height="20" alt="" >
            </van-col>
            <van-col span="16" class="tr">
              <span style="font-size:12px" class="graytxt">{{getData.fpdm}}</span>
            </van-col>
          </van-row>
        </div>
      </van-col>
      <van-col span="6" class="f14 tr mr5 bold title" style="font-size: 15px">  {{getData.dddd}}</van-col>
    </van-row>
    <div class="mt10 f14">
      <div class="line"></div>
    </div>
    <van-row  justify="space-between" align="center" class="mt10">
      <van-col span="6" class="f14" style="font-weight: bold;">{{getData.xm}}</van-col>
      <van-col span="18" class=" tr mr5 f14" style="font-weight: bold;"> {{getData.pb}}</van-col>
    </van-row>
    <van-row  justify="space-between"  class="mt5">
      <van-col span="12" class=" graytxt f12" >  发车时间: {{trainFormat(getData.cfsj)}}</van-col>
      <van-col span="12" class=" graytxt f12 pl20" > 创建时间: {{getData.cjrq.substr(0,10)}}</van-col>
    </van-row>

  </div>
</template>

<script  setup lang="ts">
import arrowRight from '../../../../assets/img/icon_arrow_right.png';
import train from '../../../../assets/img/trainNumber.png';
import {invoiceTypeList} from "./ts/invoiceTypeList";
import moment from "moment";
interface Props{
  getData?:Object;

}
defineProps<Props>()
var invoiceTitle='';
function getInvoiceTypeList (data: any) {
  let invoiceType=invoiceTypeList.find((e:any)=>data==e.value);
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
  return num;
}
function trainFormat(date:any){
  return date.substr(0,10);
}
</script>

<style scoped>
.zt1{
  color: #4BA3FB;
}
.tip{
  padding: 3px 0px 3px 5px;
  height: 20px;
  font-size: 14px;
  font-weight: bold;
  width:90px;
  text-align: left;
  line-height: 20px;
  background: url('../../../../assets/img/fenlei_bg.png') no-repeat center ;
  background-size:90px 40px;
  /*word-break: keep-all;*/
  position: relative;
  left: -13px;
  /*top: 10px;*/
}
</style>
