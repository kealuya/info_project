<template>
  <navbar title="发票夹" :icon="InvoiceBilling" @rightEvent="goInvoicing()" />
  <van-tabs v-model:active="active" sticky @click-tab="onClickTab" color="#ffffff" offset-top="45px"
            :border="true"
            style="margin-top: -1px"
            title-active-color="#ffffff"
            title-inactive-color="#ecf5ff"
            background="#0088ff" >
    <van-tab title="未使用" :name="0">
      <not-use-invoice ref="notuse"  @select-invoice="selectInvoiceList"  />
    </van-tab>
    <van-tab title="已使用" :name="1">
      <use-invoice ref="use"/>
    </van-tab>
  </van-tabs>
  <div class="positionBtn">
    <img :src="printInvoice" alt="" width="30" height="30" @click="linkToAdd()">
    <img :src="chooseIndex=='0'?filterInvoice:filterInvoice01" alt="" width="30" height="30" @click="showFilter=true"
         class="filter">
    <van-popup v-model:show="showFilter" position="bottom">
      <van-row justify="center" align="center" class="pt10 pb10">
        <van-col span="20" offset="1" class="f20 " style="text-align: center;">
          <van-row justify="space-between" align="center">
            <van-col span="20" offset="1" class="f20">
              发票种类
            </van-col>
            <van-col span="3" class="f20">
              <van-icon name="cross" @click="showFilter=false"/>
            </van-col>
          </van-row>
        </van-col>
      </van-row>
      <div class="ml10 mr10">
        <van-row justify="space-around" class="mb10" align="center">
          <van-col span="7" class="mt10 fiterFp " v-for="(item,index) in fpFilterList"
                   :class="{'active':chooseIndex==index} " :key="index" @click="chooseFplx(item.value,index)">
            <div>{{ item.title }}</div>
          </van-col>
        </van-row>
      </div>
      <div style="height: 10px"></div>
    </van-popup>
  </div>
  <div class="bottomBtn">
    <van-row justify="center" align="center">
      <van-col span="22">
        <van-button
            type="primary"
            @click="submitCheckInvoiceList"
            :disabled="checkInvoiceList.length==0"
            block>去报销({{ checkInvoiceList.length }})
        </van-button>
      </van-col>
    </van-row>
  </div>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import Navbar from '../../../components/navBar/index.vue'
import NotUseInvoice from "./component/notUseInvoice.vue";
import UseInvoice from "./component/useInvoice.vue";
import {useRouter} from "vue-router";
import printInvoice from '../../../assets/img/icon_printInvoice.png'
import InvoiceBilling from '../../../assets/icon/icon_invoice_billing.png'
import filterInvoice from '../../../assets/img/shaixuan_icon.png'
import filterInvoice01 from '../../../assets/img/shaixuan_icon_02.png'
import Invoicing from '../../../assets/icon/Invoicing.png'
import {invoiceTypeList} from './component/ts/invoiceTypeList'
//数组下标
let chooseIndex = ref(0);
//获取组件实例
let notuse = ref();
let use = ref();
const checkInvoiceList = ref([]);
let fpFilterList = invoiceTypeList;
// 进入发票抬头信息
const goInvoicing=()=>{
  router.push('invoiceBilling')
}
//点击条件进行筛选
const chooseFplx = async (type: any, index: any) => {
  chooseIndex.value = index
  //改变所需要的type值
  showFilter.value = false;
  notuse.value.payload.type = type.toString();
  await notuse.value.onRefresh('1');
  //当已使用组件挂载时，否则会有假报错
  if (use.value) {
    use.value.payload.type = type.toString()
    await use.value.onRefresh('1');
  }
}
const showFilter = ref(false)
const router = useRouter();
const active = ref(0);

const linkToAdd = () => {
  router.push('/addList');
  chooseIndex.value = 0;
  notuse.value.payload.type = '';
}
const onClickTab = async ({name, title}: { name: number, title: string }) => {
  // 测试方法
  //切换tabs值时我们需要将index变为0，并且获取全部的数据
  chooseIndex.value = 0
  if (notuse.value.list.length == 0) {
    notuse.value.payload.type = '';
    await notuse.value.onRefresh('1');
  }
  //当已使用组件挂载时，否则会有假报错
  if (use.value) {
    use.value.payload.type = '';
    await use.value.onRefresh('1');
  }
};
// 处理子组件选中的发票数组
const selectInvoiceList=(data:any)=>{
  let indexNum = checkInvoiceList.value.findIndex((item: { fpdm: any; }) => {
    return data.fpdm == item.fpdm;
  });;
  if(indexNum==-1){
    checkInvoiceList.value.push(data as never);
  }else{
    checkInvoiceList.value.splice(indexNum,1);
  }
}
// 携带发票信息进入日常报销页
function submitCheckInvoiceList() {
  // 因为日常报销存在
  const data :any ={
    'bxType':'rcbx',
    'approveNo':'',
    'isBx':'-'
  };
  sessionStorage.setItem('reimburseDetail', JSON.stringify(data))
  sessionStorage.setItem('invoiceList', JSON.stringify(checkInvoiceList.value))
  sessionStorage.setItem('billList', JSON.stringify({'billId':''}));
  router.push('addDailyBx');
}
</script>

<style lang="less" scoped>
.positionBtn {
  position: fixed;
  right: 10px;
  bottom: 100px;
  width: 30px;
  height: 30px;
  padding: 10px;
  border: 1.5px solid #0080ff;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.8);
}

.fiterFp {
  background: rgba(204, 204, 204, 0.15);
  text-align: center;
  height: 40px;
  line-height: 40px;
  border-radius: 7px 7px 7px 7px;
  font-size: 14px;
  //padding: 0;
  //word-break: keep-all
}

.active {
  text-align: center;
  border-radius: 7px 7px 7px 7px;
  font-size: 15px;
  background: rgba(0, 128, 255);
  color: #fff;
  box-shadow: 0.5px 3.5px 7px rgba(0, 128, 255, 0.5);
}

.invoicing {
  position: fixed;
  right: 10px;
  bottom: 240px;
  width: 30px;
  height: 30px;
  padding: 10px;
  border: 1.5px solid #0080ff;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.8);
  display: block;
  color: #4BA3FB;
  font-size: 38px;
  text-align: left;
}

.filter {
  position: fixed;
  right: 10px;
  bottom: 170px;
  width: 30px;
  height: 30px;
  padding: 10px;
  border: 1.5px solid #0080ff;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.8);
  display: block;
  color: #4BA3FB;
  font-size: 38px;
  text-align: left;
}
.van-tabs ::after {
  border: none;
}
.bottomBtn {
  position: fixed;
  bottom: 0px;
  width: 100%;
  padding: 10px 0px;
  background: #fff;
}
</style>
