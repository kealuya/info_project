<template>
  <van-sticky>
    <navbar title="成本中心" />
    <div class="search">
      <van-row justify="space-between" align="center">
        <van-col span="19">
          <van-search  @search="onSearch" v-model="value" shape="round" placeholder="请输入成本中心名称"   />
        </van-col>
        <van-col span="5">
          <div class="tc" @click="showPicker=true">
            <span class="f14 graytxt bold">{{typeTitle}}</span>
            <van-icon :name="xiaLaIcon" class="ml5" size="10"/>
          </div>
        </van-col>
      </van-row>
    </div>
  </van-sticky>
  <div class="mt10">
    <van-row align="center">
      <div @click="showYearPicker=true">
        <span style="color:#0088ff" class="f14 ml10">{{yearTitle}}年</span>
        <van-icon name="arrow-down" class="ml5" color="#c9c9c9" />
      </div>
    </van-row>
  </div>
  <div v-if="typeTitle=='部门'">
    <dept-cost-center-list ref="deptListRef" />
  </div>
  <div v-if="typeTitle=='项目'">
    <cost-center-list ref="costCenterListRef" />
  </div>
  <van-popup v-model:show="showPicker" round position="bottom">
    <van-picker
        :columns="costList"
        @cancel="showPicker = false"
        @confirm="onConfirm"
    />
  </van-popup>
  <van-popup v-model:show="showYearPicker" round position="bottom">
    <van-picker
        :columns="yearList"
        @cancel="showYearPicker = false"
        @confirm="onYearConfirm"
    />
  </van-popup>
</template>

<script lang="ts" setup>
import Navbar from '../../components/navBar/index.vue';
import xiaLaIcon from '../../assets/icon/xiala.png';
import DeptCostCenterList from './component/deptCostCenterlist/index.vue'; // 导入部门成本中心列表组件
import CostCenterList from './component/costCenterlist/index.vue';  // 导入项目成本中心列表组件
import { ref,onMounted} from "vue";
let typeTitle=ref('部门'); //切换成本中心类型
let showPicker=ref(false); // 切换成本中心类型的卡片
let showYearPicker=ref(false); // 切换成本中心年度的卡片
const deptListRef=ref(); //自定义组件-部门成本中心数组ref；
const costCenterListRef=ref(); //自定义组件-项目成本中心数组ref；
const nowDate = new Date(); // 获取当前时间
let nowYear = nowDate.getFullYear(); //获取当前年份
let costList=ref([
    {'text':"部门",'value':'dept'},
    {'text':"项目",'value':'cost'},
]); // 成本中心数组
let yearTitle=ref('2023'); // 区分年份
let yearList=ref([
])
let value=ref(''); // 搜索字段信息
let payload=ref(
    {
      "corpCode":"",
      "keyWord":"",
      "budGetingYear":"",
      "page":1,
      "pageSize":10,
      "equipmentType":""
    }); //提交参数
// 切换年度
// 输入名称和关键字搜索
const onSearch=(val:any) =>{
  // keyWord
  if(typeTitle.value=='部门'){
      deptListRef.value.payload.keyWord=val;
      deptListRef.value.refresh();
  }else{
      costCenterListRef.value.payload.keyWord=val;
      costCenterListRef.value.refresh();
  }
}
// 选择的年度
const onYearConfirm=(val:any) =>{
  if(typeTitle.value=='部门'){
      yearTitle.value=val.selectedOptions[0].text;
      deptListRef.value.payload.budGetingYear=val.selectedOptions[0].text;
      deptListRef.value.onRefresh();
      showYearPicker.value=false;
  }else{
    yearTitle.value=val.selectedOptions[0].text;
      costCenterListRef.value.payload.budGetingYear=val.selectedOptions[0].text;
      costCenterListRef.value.onRefresh();
      showYearPicker.value=false;
  }
}
// 切换成本中心类型
const  onConfirm=(value:any)=>{
  typeTitle.value=value.selectedOptions[0].text;
  showPicker.value=false;
} //切换预算编制
onMounted(()=>{
  yearTitle.value=`${nowYear}`;
  for(let i=0;i<5;i++){
    yearList.value.push({text:`${nowYear-i}`,value:`${nowYear-i}`,}as never)
  }
});
</script>

<style scoped>
.search{
  width: 100%;
  height: 55px;
  background: #ffffff;
}
</style>
