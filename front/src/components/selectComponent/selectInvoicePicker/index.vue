<!--
  选择发票信息组件（单选）
  使用方法如下： 可参考日常报销
-->
<template>
  <div class="selectPersonPicker">
    <van-sticky>
      <van-nav-bar left-arrow>
        <template #left>
          <van-icon name="arrow-left" color="#ffffff" @click="escOver" size="20"/>
        </template>
        <template #title>
          <h4 class="f_white">选择发票</h4>
        </template>
        <template #right>
          <van-icon :name="ScreenIcon" @click="showFilter=true" size="20"/>
        </template>
      </van-nav-bar>
    </van-sticky>
    <div style="margin-bottom: 60px">
      <van-pull-refresh v-model="refreshing" @refresh="onRefresh" style="background:#f1f2f3; min-height: 100vh;">
        <template #success>
          <span v-if="refreshingOnly==true"><van-icon name="success"/> 刷新成功</span>
        </template>
        <div v-if="loading==false && list.length==0" style="height: 500px">
          <van-row justify="center" align="center" class="tc">
            <div>
              <img style="height:50vh;width:80vw;object-fit: contain" src="../../../assets/img/zanwupiaoju.png"
                   v-cloak/>
            </div>
          </van-row>
          <div style="text-align:center;color: rgba(23,26,29,0.40)">暂无可以用的发票信息</div>
          <van-row justify="center" align="center" class="tc mt5">
            <div>
              <van-button square plain type="primary" @click="onRefresh" style="padding: 5px 5px;border-radius: 8px">刷新
              </van-button>
            </div>
          </van-row>

        </div>
        <van-list
            v-model="loading"
            v-else
            :finished="finished"
            :immediate-check="false"
            loading-text="正在加载中请稍后"
            finished-text="没有更多了"
            @load="onLoad"
            :offset="10"
        >
          <div class="ticket_content" v-for="(item,index) in list" :key="item">
            <airplane-component v-if="item.type=='5'" :get-data="item" @click="selectInvoiceOverList(item,index)"/>
            <train-component v-else-if="item.type=='8'" :get-data="item" @click="selectInvoiceOverList(item,index)"/>
            <invoice-component v-else :get-data="item" @click="selectInvoiceOverList(item,index)"/>
          </div>
        </van-list>
      </van-pull-refresh>
      <div>
        <!--  浮动添加按钮      -->
        <van-sticky position="bottom">
          <div class="sticky" @click="onClick">
            <van-icon name="records" color="#0088ff" size="30"/>
          </div>
        </van-sticky>

      </div>
    </div>
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
    <div class="bottomGroup">
      <van-button type="primary" style="width:95vw;height:40px" @click="costCenterOK" size="normal">
        确定（{{ selectInvoiceList.length }}）
      </van-button>
    </div>
    <load-component v-if="loading" />
  </div>
</template>

<script lang="ts" setup>
import {useRouter} from "vue-router";
import {inject, ref, defineEmits} from "vue";
import {getInvoiceWbList} from "../../../services/invoice";
import ScreenIcon from '../../../assets/icon/screen.png'
import AirplaneComponent from "./components/airplaneComponent.vue";
import TrainComponent from "./components/trainComponent.vue";
import InvoiceComponent from "./components/invoiceComponent.vue";
import Navbar from  '../../../components/NavBar.vue'
import {invoiceTypeList} from "../../../pages/invoiceFolder/invoice/component/ts/invoiceTypeList";
import LoadComponent from '../../loading/index.vue'
const router = useRouter(); // 用于监控路由在此处声明一下路由参数
let userInfoData: any = inject("userInfo"); //获取审批数组
let list = ref<any>([]);//加载loading
let loading = ref(true); //加载效果
let refreshingOnly = ref(false); //加载效果只有下拉变化
let finished = ref(false); // 结束
const refreshing = ref(false); // 重新加载还是追加
const emit = defineEmits(['selectOver']) //emit 方法声明

// 提交获取账本列表的数组
let payload = {
  "empCode": userInfoData.userInfo.empCode,
  "companyCode": userInfoData.userInfo.corpCode,
  "corpCode": userInfoData.userInfo.corpCode,
  "page": '1',
  "keyWord": "",
  "equipmentType": "",
  "row": '10',
  "zt": "0",
  "type": "",
}
// 用于筛选栏显隐
const showFilter = ref(false);
// 发票类型
let fpFilterList = invoiceTypeList;
//数组下标
let chooseIndex = ref(0);
//点击条件进行筛选
const chooseFplx = async (type: any, index: any) => {
  chooseIndex.value = index
  //改变所需要的type值
  showFilter.value = false;
  payload.type = type.toString();
  onRefresh('1');
}

interface Props {
};
const props = defineProps<{
  selectList:[]
}>()
// 用于存放选中的账单列表
let selectInvoiceList = ref([]);
// 分页加载
const onLoad = async (code: any) => {
  if (code == '1') {
    payload.page = '1';
  }
  if (refreshing.value) {
    return;
  }
  if (finished.value == true) {
    return
  }
  await getCostCenterList(payload, 'loading')
};
// 首次进入刷新
const refresh = async () => {
  if (refreshing.value) {
  }
  if (finished.value == true) {
    return
  }
  payload.page = '1';
  loading.value=true;
  await getCostCenterList(payload, 'refresh')
}
// 刷新处理
const onRefresh = (code: any) => {
  payload.page = '1';
  refreshing.value = true;
  if (code != '1') {
    refreshingOnly.value = true;
  } else {
    refreshingOnly.value = false;
  }
  finished.value = false;
  loading.value = true;
  refresh();
};
// 选中后提交
const costCenterOK = () => {
  emit('selectOver', selectInvoiceList.value);
  selectInvoiceList.value=[];
}
// 点击跳转到添加账本页
const onClick = () => {
  router.push({
    name: 'addList',
  })
};
// 选择发票
const selectInvoiceOverList = (data: any, index: any) => {
  /**
   *  选择发票组件
   *  */
      // 此处是获取数据以及其数组下标
  let indexNum = selectInvoiceList.value.findIndex((item: { fpdm: any; }) => {
        return data.fpdm == item.fpdm;
      });
  if (indexNum != -1) {
    list.value[index].checked = false;
    selectInvoiceList.value.splice(indexNum, 1);
  } else {
    list.value[index].checked = true;
    selectInvoiceList.value.push(data as never);
  }
};
// 调用发票列表
const getCostCenterList = async (payload: any, code: any) => {
  let result: any = await getInvoiceWbList(payload);
  if (code == 'refresh') {
    list.value = [];
  }
  list.value.push(...result.data.data.data);
  listFormat();
  if (list.value.length === result.data.data.count) {
    finished.value = true;
    loading.value = false;
  } else {
    let page = Number(payload.page) + 1;
    payload.page = page.toString();
    loading.value = false;
  }
  refreshing.value = false;
}
const escOver=()=>{
  selectInvoiceList.value=[];
  emit('selectOver',[])
}
// 将发票列表赋初值
const listFormat=()=>{
  if(props.selectList.length!=0){
    selectInvoiceList.value=props.selectList;
    let invoiceList:any=[];
    list.value.forEach((item:any)=>{
      let key=inList(item.fpdm);
      return invoiceList.push({...item,checked:key});
    })
    list.value=invoiceList;
  }

}
// 判断是否存在于数组
const inList=(key:any)=>{
  let label=selectInvoiceList.value.find((item: { fpdm: any; }) => {
    return key==item.fpdm;
  })
  if(label==undefined){
    return false
  }else{
    return  true
  }
}
// 放出外部可调用的方法
defineExpose({
  onLoad, onRefresh, refresh, payload, list
})
</script>

<style scoped>
.van-nav-bar {
  background: #0088ff;
}

.selectPersonPicker {
  /*height: 100%;*/
  display: flex;
  flex-direction: column;
}

.sticky {
  position: fixed;
  right: 10px;
  bottom: 100px;
  width: 40px;
  height: 30px;
  text-align: center;
  padding: 8px;
  border: 1px solid #0088ff;
  border-radius: 20px;
  background: #FFFFFF;
}

.fiterFp {
  background: rgba(204, 204, 204, 0.15);
  text-align: center;
  height: 40px;
  line-height: 40px;
  border-radius: 7px 7px 7px 7px;
  font-size: 14px;
  padding: 0;
  word-break: keep-all
}

.active {
  text-align: center;
  border-radius: 7px 7px 7px 7px;
  font-size: 15px;
  background: rgba(0, 128, 255, 0.5);
  color: #fff;
  box-shadow: 0.5px 3.5px 7px rgba(0, 128, 255, 0.5);
}

.bottomGroup {
  width: 100%;
  height: 40px;
  background: #ffffff;
  position: fixed;
  bottom: 0px;
  padding: 10px;
}
</style>
