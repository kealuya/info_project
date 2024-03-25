<template>
  <van-sticky>
    <NavBar title="单据动态" />
    <van-dropdown-menu z-index="999"	>
      <van-dropdown-item teleport="body" disabled  >
        <template #title>
          <div @click="showTypePicker=true">
           <span :class="showTypePicker==true?'checked':'NoCheck'">  {{typeTitle}}</span>
          </div>
        </template>
      </van-dropdown-item>
      <van-dropdown-item     :title="stateTitle" v-model="value2" :options="option2" @change="stateLoad" />
      <van-dropdown-item disabled>
        <template #title >
           <div @click="showCreateCalendar=true">
             <span :class="showCreateCalendar==true?'checked':'NoCheck'">  {{timeTitle}}</span>
           </div>
        </template>
      </van-dropdown-item>
      <van-dropdown-item  :title="searchTitle"   disabled>
        <template #title>
          <span @click="showSearchPicker=true">
               <span :class="showSearchPicker==true?'checked':'NoCheck'">  {{searchTitle}}</span>
          </span>
        </template>
      </van-dropdown-item>
    </van-dropdown-menu>
  </van-sticky>
  <van-pull-refresh v-model="refreshing" @refresh="onRefresh" style="min-height: 80vh">
    <template #success>
      <span v-if="refreshingOnly==true"><van-icon name="success" /> 刷新成功</span>
    </template>
    <div v-if="loading==false && list.length==0" style="height: 500px">
      <van-row justify="center" align="center" class="tc">
        <div>
          <img style="height:50vh;width:80vw;object-fit: contain" src="../../assets/img/zanwupiaoju.png" v-cloak/>
        </div>
      </van-row>
      <div style="color: rgba(23,26,29,0.40);text-align: center">暂无申请单</div>
    </div>
    <van-list
        v-model="loading"
        v-else
        :finished="finished"
        loading-text="正在加载中请稍后"
        :immediate-check="false"
        finished-text="没有更多了"
        @load="onLoad"
        :offset="10"
        style="min-height: 80vh"
    >
      <div class="ticket_content" v-for="(item,index) in list" :key="item">
       <apply-card :get-data="item" />
      </div>
    </van-list>
  </van-pull-refresh>
  <van-calendar
      ref="myCalendar"
      v-model:show="showCreateCalendar"
        @select="checkedConfirm" :show-confirm="true"  :min-date="minDate" >
    <template  #footer>
     <div class="mb10 mt10">
       <van-row>
         <van-col span="6">
           <van-button type="default" @click="resetAll" >重置</van-button>
         </van-col>
         <van-col span="18">
           <van-button type="primary" block @click="onConfirmAll">确定</van-button>
         </van-col>
       </van-row>
     </div>
    </template>
  </van-calendar>
  <van-popup
      v-model:show="showTypePicker"
      round
      closeable
      position="bottom"
      :style="{ height: '50%' }"
  >
  <van-row justify="center" align="center">
    <van-col class="mt10">
      <span class="f18  tc" style="text-align: center">单据类型</span>
    </van-col>
  </van-row>
    <van-row class="mt10">
      <van-col :span="8" v-for="item in applicationList">
        <div class="selectCard"
            :class="applicationChecked==item.value?'applcationCheckCard':'applcationNoCheckCard'"
            @click="selcectApplicationCheckedCard(item)">{{item.text}}</div>
      </van-col>
    </van-row>
  </van-popup>
  <!-- 右侧弹出搜索 -->
  <van-popup
      v-model:show="showSearchPicker"
      position="right"
      :style="{ width: '70%', height: '100%' }"
  >
    <div class="mt10 ">
      <span class="f18 ml10  tc" style="text-align: center">单据搜索</span>
    </div>
    <div class="mt5">
      <span class="ml10 font_gray f14">申请事由</span>
    </div>
    <van-search v-model="value" placeholder="请输入申请事由" />
    <van-row justify="center" align="center">
      <van-col >
        <van-button type="default" @click="resSearch" size="small">重置</van-button>
      </van-col>
      <van-col offset="1">
        <van-button type="primary" @click="subSearch" size="small">提交</van-button>
      </van-col>
    </van-row>
  </van-popup>
  <div v-show="loadingComponent">
    <loading-component :title="'提交中'"/>
  </div>
</template>
<script lang="ts" setup>
import NavBar from '../../components/navBar/index.vue';
import {ref, inject, onMounted, defineEmits} from "vue";
import {getAllApplyList, getDocStateList} from "../../services/home";
import ApplyCard from "./component/applyCard.vue";
import LoadingComponent from '../../components/loading/index.vue'
import moment from "moment/moment";
const value2 = ref('a'); // 初始下拉动态复制
const typeTitle=ref('类型'); //  初始下拉类型标题赋值
const timeTitle=ref('创建时间'); //  初始下拉创建时间标题赋值
const stateTitle=ref('状态'); //  初始下拉状态标题赋值
const searchTitle=ref('筛选'); //  初始下拉筛选标题赋值
const loadingComponent=ref(false); //loading组件控制显示隐藏
const value=ref(''); // 搜索框的值
const showTypePicker=ref(false); // 控制弹出框显隐
const showSearchPicker=ref(false); // 控制弹出层
const myCalendar=ref();
const nowDate = new Date(); //获取当前系统时间
let minDate = new Date(nowDate.getFullYear(), nowDate.getMonth()-6, nowDate.getDay()); //最小时间是当前日期的前半年
const showCreateCalendar=ref(false); // 创建日历组件判断
const emit=defineEmits(['confirm']); // emit事件方法
const option2 = [
  { text: '全部', value: '' },
  { text: '待审批', value: '0' },
  { text: '已通过', value: '1' },
  { text: '已拒绝', value: '2' },
  { text: '已撤销', value: '3' },
  { text: '免审', value: '4' },
]; // 存放状态的数组
const applicationList=[
  { text:'全部',value:'' }, { text:'差旅申请',value:'travel' }, { text:'请假申请',value:'leave' },
  { text:'加班申请',value:'work' }, { text:'费用申请',value:'reimbursement' }, { text:'用章申请',value:'seal' },
  { text:'用车申请',value:'car' }, { text:'物品申购',value:'shop' }, { text:'董事签批',value:'sign' },
  { text:'考勤申请',value:'card' }, { text:'物品领用',value:'itemCollection' }, { text:'基金报销',value:'fund' },
  { text:'培训申请',value:'train' }, { text:'客户宴请',value:'entertain' }, { text:'会议记录',value:'meeting' },
  { text:'外出办公',value:'outWork' }, { text:'快递物流',value:'log' }, { text:'日常报销',value:'rcbx' },
  { text:'借款申请',value:'jksq' }, { text:'差旅补助',value:'clbx' }, { text:'其他申请',value:'other' },
]; //存放类型的数组
let list = ref<any>([]);//加载loading
let loading = ref(true); //加载效果]
let refreshingOnly = ref(false); //加载效果只有下拉变化
let finished = ref(false); // 结束
const refreshing = ref(false); // 重新加载还是追加
let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
let applicationChecked= ref(""); //选中的数组元素
let selectDate=ref(''); // 用于储存选中的时间
// 提交参数
let payload = {
  // 'companyCode': userInfoData.userInfo.corpCode,
  'page': 1,
  'pageSize': 10,
  'corpCode':  userInfoData.userInfo.corpCode,
  'empCode':userInfoData.userInfo.empCode,
  'approveType': [],
  'createTime': '',
  'keyWord': '',
  'approveReason': '',
  "approvalState": "",
  // "state": "全部",
  // 'state':''
} //提交动态
// 重置搜索框搜索单据动态
const resSearch=()=>{
  value.value="";
  payload.approveReason="";
};
// 提交搜索框搜索单据动态
const subSearch=()=>{
  payload.approveReason=value.value;
  if(value.value==''){
    searchTitle.value='筛选中';
  }else{
    searchTitle.value='筛选';
  }
  showSearchPicker.value=false;
  onRefresh('1');
}
// 通过点击按钮触发日历组件提交功能
const checkedConfirm=(values:any) =>{
  selectDate.value=values;
}
const onConfirmAll =()=>{
  showCreateCalendar.value = false;
  payload.createTime=moment(selectDate.value).format("yyyy-MM-DD");
  timeTitle.value =moment(selectDate.value).format("MM-DD");
  onRefresh('1');
}
// 清空重置时间筛选
const resetAll=()=>{
  timeTitle.value='创建时间';
  if(payload.createTime!=''){
    payload.createTime='';
    onRefresh('1');
  }
  showCreateCalendar.value = false;
}
// 根据单据类型筛选对应的申请单
const selcectApplicationCheckedCard=(item:any)=>{
  applicationChecked.value=item.value;
  payload.approveType=[];
  if(item.text!='全部'){
    payload.approveType.push(item.value as never);
  }
  typeTitle.value=item.text;
  showTypePicker.value=false
  onRefresh('1');
}
// 处理单据动态的接口
const getDocList=async(payload:any,code:any)=>{
  let result: any = await getDocStateList(payload);
  if(code=='refresh'){
    list.value = [];
  }
  list.value.push(...result.data.list);
  loadingComponent.value=false;
  if (list.value.length.toString() === result.data.total){
    loading.value = false;
    finished.value = true;
  } else {
    payload.page = payload.page + 1;
    // payload.page = page.toString();
    loading.value = false;
  }
  refreshing.value = false;
}
// 追加onload
const onLoad = async (code:any) => {
  // if(code=='1'){
  //   payload.page=payload.page + 1;
  // }
  if (refreshing.value) {
    return;
  }
  if(finished.value==true){
    return
  }
  await getDocList(payload,'loading')
};
// 根据条件更新
const stateLoad=(item:any)=>{
  stateTitle.value = findMenu(option2, item);
  payload.approvalState=item;
  onRefresh('1');
}
// 查询对应信息
const findMenu = (list: any, value: any) => {
  let title = list.find((item: any) => item.value == value);
  return title.text;
}
// 刷新
const refresh=async()=>{
  loadingComponent.value=true;
  if (refreshing.value) {
  }
  if(finished.value==true){
    return
  }
  await getDocList(payload,'refresh')
}
// 首次刷新
const onRefresh = (code:any) => {
  payload.page = 1;
  if(code!='1'){
    refreshing.value = true;
    refreshingOnly.value = true;
  }else{
    refreshing.value = false;
    refreshingOnly.value = false;
  }
  finished.value = false;
  loading.value = true;
  refresh();
};
onMounted(() => {
  refresh();
});
</script>

<style scoped>
.selectCard{
  margin: 5px;
  padding: 10px;
  color: #1a1a1a;
  border-radius: 10px;
  font-size: 12px;
  text-align: center;
}
  .NoCheck{
    color:#333333;
  }
  .checked{
    color:#1989fa;
  }
  .applcationNoCheckCard{
    background: rgba(207,208,217,0.2);
  }
  .applcationCheckCard{
    background: rgba(0,128,255,1);
    color: #ffffff;
  }
</style>
