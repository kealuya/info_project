<template>
  <div>
    <van-sticky :offset-bottom="50">
      <navbar-component :title="'报销动态'"></navbar-component>
      <van-row align="center" style="width: 100%;height: 60px;background: #ffffff">
        <van-col span="16">
          <van-search
              v-model="search"
              shape="round"
              @search="onSearch"
              placeholder="请输入报销事由"
          />
        </van-col>
        <van-col span="8">
          <div class="stickyPicker">
            <div class="pickerCard" @click="showPicker=true">
              <van-row justify="center" align="center">
                <van-col>{{ typeKey }}</van-col>
                <van-col offset="1">
                  <van-icon :name="xiaLaIcon" size="13"/>
                </van-col>
              </van-row>
            </div>
          </div>
        </van-col>
      </van-row>
    <div>
      <van-dropdown-menu>
        <van-dropdown-item
            :title="approvalTitle" v-model="approvalStatus" :options="option1" @change="stateApprovalLoad"/>
        <van-dropdown-item
            :title="reportingTitle" v-model="reportingStatus" :options="option2" @change="stateReportingLoad"/>
        <van-dropdown-item
            :title="submitTitle" v-model="submitStatus" :options="option3" @change="stateSubmitLoad"/>
      </van-dropdown-menu>
    </div>
    </van-sticky>
    <van-popup v-model:show="showPicker" round position="bottom">
      <van-picker
          :columns="columns"
          @cancel="showPicker = false"
          @confirm="onConfirm"
      />
    </van-popup>
    <float-button @click="showTab"/>
    <van-pull-refresh v-model="refreshing" @refresh="onRefresh" style="min-height: 80vh">
      <template #success>
        <span v-if="refreshingOnly===true"><van-icon name="success"/> 刷新成功</span>
      </template>
      <div v-if="loading===false && list.length===0" style="min-height: 80vh">
        <van-row justify="center" align="center" class="tc">
          <div>
            <img  alt="暂无单据"
                style="height:50vh;width:80vw;object-fit: contain" src="../../assets/img/zanwupiaoju.png"
                v-cloak/>
          </div>
        </van-row>
        <div style="color: rgba(23,26,29,0.40);text-align: center">暂无报销单</div>
        <van-row justify="center" align="center" class="tc mt5">
        </van-row>

      </div>
      <van-list
          v-model="loading"
          v-else
          :immediate-check="false"
          :finished="finished"
          loading-text="正在加载中请稍后"
          finished-text="没有更多了"
          @load="onLoad"
          :offset="10"
      >
        <div class="ticket_content" v-for="(item) in list" :key="item" >
            <reimburse-card :get-data="item"/>
        </div>
      </van-list>
    </van-pull-refresh>
  </div>
  <van-action-sheet v-model:show="tabButtonShow" position="bottom">
    <reimburse-type-btn/>
  </van-action-sheet>
  <div v-show="loadingComponent">
    <loading-component :title="'加载中'"/>
  </div>
</template>

<script setup lang="ts">
import {inject, onMounted, ref} from "vue"
import LoadingComponent from '../../components/loading/index.vue'
import NavbarComponent from '../../components/navBar/index.vue'
import FloatButton from './component/floatButton/index.vue'
import ReimburseTypeBtn from './component/reimbuseTypebtn/index.vue'
import xiaLaIcon from '../../assets/icon/xiala.png'
import {bxList} from "../../services/reimburse";
import ReimburseCard from './component/reimburseCard/index.vue'
import router from "../../router";

let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
// 提交参数
let payload = {
  'corpCode': userInfoData.userInfo.corpCode,
  'empCode': userInfoData.userInfo.empCode,
  "approveState": "",
  "bxType": "",
  "isHk": "",
  "keyWord": "",
  "submitStatus": "",
  "bxState": "",
  "pageSize": 10, //  显示的个数
  "currentPage": 1, //  页码
} //提交动态
let list = ref<any>([]);//加载loading
let loading = ref(true); //加载效果]
let refreshingOnly = ref(false); //加载效果只有下拉变化
let finished = ref(false); // 结束
const refreshing = ref(false); // 重新加载还是追加
const tabButtonShow = ref(false) // 控制显示下方弹出层显隐
const search = ref(''); // 输入搜索项
const typeKey = ref('全部'); // 右侧筛选框
const loadingComponent=ref(false); //loading组件控制显示隐藏
//  审批状态文案
const approvalTitle = ref('审批状态');
// 审批状态
const approvalStatus = ref('');
// 审批状态列表
const option1 = [
  {text: '全部', value: ''},
  {text: '未审批', value: '0'},
  {text: '已通过', value: '1'},
  {text: '已拒绝', value: '2'},
];
// 控制状态显隐
const showPicker=ref(false);
// 报销单类型
const columns = [
  { text: '全部', value: '' },
  { text: '日常报销', value: 'rcbx' },
  { text: '差旅补助报销', value: 'travelbx' },
  { text: '借款申请', value: 'jkbx' },
];
// 搜索
const onSearch = (val:any) => {
  payload.keyWord=val;
  onRefresh('1');
};
// 日常报销
const onConfirm=(value: any)=>{
  showPicker.value = false;
  typeKey.value = value.selectedOptions[0].text;
  payload.bxType = value.selectedOptions[0].value;
  onRefresh('1');
}
// 审批状态更新方法
const stateApprovalLoad = (item: any) => {
  approvalTitle.value = findMenu(option1, item);
  approvalStatus.value = item;
  payload.approveState=item;
  onRefresh('1');
}
// 查询对应信息
const findMenu = (list: any, value: any) => {
  let title = list.find((item: any) => item.value == value);
  return title.text;
}
// 追加onload
const onLoad = async () => {
  if (refreshing.value) {
    return;
  }
  if(finished.value==true){
    return
  }
  await getDocList(payload,'loading')
};
// 填报状态文案
const reportingTitle = ref('填报状态')
// 填报状态
const reportingStatus = ref('');
// 审批状态列表
const option2 = [
  {text: '全部', value: ''},
  {text: '已报销', value: 'true'},
  {text: '已填报', value: 'false'},
];
// 报销状态
const stateReportingLoad = (item: any) => {
  reportingTitle.value = findMenu(option2, item);
  reportingStatus.value = item;
  payload.bxState=item;
  onRefresh('1');
}
// 提交状态文案
const submitTitle = ref('提交状态')
// 提交状态
const submitStatus = ref('');
// 提交状态
const option3 = [
  {text: '全部', value: ''},
  {text: '已提交', value: '已提交'},
  {text: '待提交', value: '-'},
]
const stateSubmitLoad = (item: any) => {
  submitTitle.value = findMenu(option3, item);
  submitStatus.value = item;
  payload.submitStatus=item;
  onRefresh('1');
}
// 用于下方弹出层显隐
const showTab = () => {
  tabButtonShow.value = true
}
// 处理单据动态的接口
const getDocList=async(payload:any,code:any)=>{
  let result: any = await bxList(payload);
  if(code=='refresh'){
    list.value = [];
  }
  list.value.push(...result.data.records);
  if (list.value.length.toString() === result.data.total.toString()){
    loading.value = false;
    finished.value = true;
  } else {
    payload.currentPage = payload.currentPage + 1;
    loading.value = false;
  }
  loadingComponent.value=false;
  refreshing.value = false;
}
// 刷新
const refresh = async () => {
  loadingComponent.value = true;
  if (refreshing.value) {
  }
  if (finished.value == true) {
    return
  }
  payload.currentPage = 1;
  await getDocList(payload, 'refresh')
}
// 首次刷新
const onRefresh = (code: any) => {
  payload.currentPage = 1;
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

<style lang="less" scoped>
.stickyPicker {
  padding-top: 10px;
  padding-bottom: 10px;
  background: #ffffff;
}
.pickerCard {
  padding: 5px;
  margin-right: 10px;
  background: rgba(23, 26, 29, 0.1);
  border-radius: 20px;
  color: #999999;
  font-size: 14px;
  font-weight: bold;
}
</style>
