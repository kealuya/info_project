<template>
    <div class="m10">
      <van-row justify="space-between" align="center">
        <van-col>
          <span>累计申请</span>
        </van-col>
        <van-col>
          <div>
            <van-row align="center" justify="center">
              <span class="font_gray f14" @click="showPicker=true">{{ fieldValue }}</span>
              <van-col>
                <van-icon class="ml5"  size="10" :name="rightArrowIcon" />
              </van-col>
            </van-row>
          </div>
        </van-col>
      </van-row>
      <div class="mt10 tc mb10">
        <span class="f14" v-if="autoDateShow">{{ startDate }} 至 {{ endDate }}</span>
        <div v-else>
          <van-row justify="center" align="center">
            <van-col span="16">
              <div class="calendarCard">
                <van-row justify="center" align="center">
                  <van-col span="4">
                    <van-image :src="calendarIcon" width="20" height="20"/>
                  </van-col>
                  <van-col span="20">
                    <span class="f14" @click="show=true">{{ startDate }} 至 {{ endDate }}</span>
                  </van-col>
                </van-row>
                <van-calendar type="range" v-model:show="show" @confirm="notAutoDateConfirm"/>
              </div>
            </van-col>
          </van-row>
        </div>
      </div>
      <div id="main" class="testMain"></div>
      <van-popup v-model:show="showPicker" round position="bottom">
        <van-picker
            :columns="columns"
            @cancel="showPicker = false"
            @confirm="onConfirm"
        />
      </van-popup>
    </div>
</template>

<script lang="ts" setup>
import rightArrowIcon from '../../../../assets/icon/xiala.png';
import calendarIcon from '../../../../assets/icon/kaoqin_icon_05.png'
import {inject, onBeforeMount,ref,defineEmits} from "vue";
import moment from "moment/moment"; // 导入时间格式处理组件
import * as echarts from "echarts";
import {countIndividualNumber} from "../../../../services/statement";
import {applyType} from "../../../home/component/ts/applyType"; //引用echarts图表
const nowDate = new Date(); // 当前时间
const year = nowDate.getFullYear(); // 当前年份
const month = nowDate.getMonth() + 1; // 月份从0开始计算，需要加1
// 选择时间的范围适用于调用选择接口
// 第一次发现moment格式化还可以这样写 moment(nowDate).format("yyyy-01-01") 获取的是当前年份
// 可以利用add和subtract 进行日期时间的相加减 后方传递个格式即可 这样可以省去搞一些处理时间的方法
// 为了能充分的懒，所以就这样取得当前的时间的方法（暂时觉得这个方法比较省时间，如果有更好的方法会进行修改）
// 以下是时间处理数组
const columns = [
  {
    text: '近30天',
    value: '0',
    startTime: moment(nowDate).subtract(1, 'months').format("yyyy-MM-DD"),
    endTime: moment(nowDate).format("yyyy-MM-DD"),
    auto: '0' // 用于区分自定义还是设定的时间选项
  },
  {
    text: '当前月',
    value: '2',
    startTime: moment(nowDate).format("yyyy-MM-01"),
    endTime: moment(new Date(year, month, 0)).format("yyyy-MM-DD"),
    auto: '0' // 用于区分自定义还是设定的时间选项
  },
  {
    text: '近三月',
    value: '3',
    startTime: moment(nowDate).subtract(3, 'months').format("yyyy-MM-DD"),
    endTime: moment(nowDate).format("yyyy-MM-DD"),
    auto: '0' // 用于区分自定义还是设定的时间选项
  },
  {
    text: '近半年',
    value: '4',
    startTime: moment(nowDate).subtract(6, 'months').format("yyyy-MM-DD"),
    endTime: moment(nowDate).format("yyyy-MM-DD"),
    auto: '0' // 用于区分自定义还是设定的时间选项
  },
  {
    text: '本年度',
    value: '5',
    startTime: moment(nowDate).format("yyyy-01-01"),
    endTime: moment(nowDate).format("yyyy-12-31"),
    auto: '0' // 用于区分自定义还是设定的时间选项
  },
  {
    text: '近一年',
    value: '6',
    startTime: moment(nowDate).subtract(1, 'years').format("yyyy-MM-DD"),
    endTime: moment(nowDate).format("yyyy-MM-DD"),
    auto: '0' // 用于区分自定义还是设定的时间选项
  },
  {
    text: '自定义',
    value: '7',
    startTime: '',
    endTime: '',
    auto: '1' // 用于区分自定义还是设定的时间选项
  },
];
const startDate = ref(moment(nowDate).subtract(1, 'years').format("yyyy-MM-DD")); // 数据起始时间
const endDate = ref(moment(nowDate).format("yyyy-MM-DD")); // 数据截止时间
const fieldValue = ref('近一年'); //  选项信息的初始值
const showPicker = ref(false); // 控制日期时间的选项
const autoDateShow = ref(true); // 自定义时间or 规范设定的时间
const show = ref(false); //控制选择器的显隐
const xChartList=ref([]); // 柱状图的x轴的数据类别
const yChartList=ref([]); // 柱状图的y轴的数据
let userInfoData: any = inject("userInfo"); //获取用户信息
const applyCount=ref(0);
const emit=defineEmits(['overData']) // 触发子组件方法
// 时间选择器更改时间选择期限
const onConfirm = (value: any) => {
  showPicker.value = false;
  fieldValue.value = value.selectedOptions[0].text;
  // 判断是否为自定义，不是自定义直接调用接口获取对应的数据信息
  if (value.selectedOptions[0].auto == '0') {
    autoDateShow.value=true;
    startDate.value = value.selectedOptions[0].startTime;
    endDate.value = value.selectedOptions[0].endTime;
    countNumber();
  } else {
    autoDateShow.value=false;
  }
};
//用于 自选的申请图表
const notAutoDateConfirm = (values: any) => {
  const [start, end] = values;
  startDate.value = moment(start).format("yyyy-MM-DD");
  endDate.value = moment(end).format("yyyy-MM-DD");
  countNumber();
  show.value = false;
};
// 获取个人申请单数目
const countNumber=async()=>{
  let revokeInfo={
    "empCode":userInfoData.userInfo.empCode,
    'corpCode':userInfoData.userInfo.corpCode,
    'startTime':`${startDate.value} 00:00:00`,
    'endTime':`${endDate.value} 23:59:59`,
  }
  let result: any = await countIndividualNumber(revokeInfo);
  applyCount.value=result.data.totalNum;
  const list= dataFormat(result.data.approveNum);
  dataTwoFormat(list);
  emit('overData',applyCount.value);
}
// 处理申请单数据展示
const dataFormat=(dataList:any)=>{
  let list=dataList;
  for(let i=0;i<=list.length;i++){
    for(let j=0;j<list.length-1-i;j++){
      if (list[j].approveNumAmount < list[j+1].approveNumAmount) {        //相邻元素两两对比
        var temp = list[j+1];        //元素交换
        list[j+1] = list[j];
        list[j] = temp;
      }
    }
  }
  return list;
}
// 处理图标展示数据
const dataTwoFormat=(list:any)=>{
  xChartList.value=[];
  yChartList.value=[];
  let amount=applyCount.value;
  for(let i=0;i<4;i++){
    xChartList.value.push(getApprovalTypeList(list[i].approveType) as never);
    yChartList.value.push(list[i].approveNumAmount as never);
    amount=amount-list[i].approveNumAmount;
  }
  xChartList.value.push('其他' as never);
  yChartList.value.push(amount as never);
  init();
}
// 绘制柱状图
function init() {
  // 基于准备好的dom，初始化echarts实例
  let Chart = echarts.init(document.getElementById("main"));
  // 绘制图表
  let options = {
    xAxis: {
      type: 'category',
      data: xChartList.value,
    },
    grid: {
      top: '20%',
      bottom: '5%',
      left: '2%',
      right: '4%',
      containLabel: true
    },
    position: "insideRight",
    tooltip: {
      trigger: 'axis',
      formatter: '{b} : {c}次'
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {
        data: yChartList.value,
        type: 'bar',
        itemStyle: {
          label: {
            formatter: "{c}"+"次",
            show: true,
            position: "top",
            textStyle: {
              fontWeight: "bolder",
              fontSize: "12",
              color: "#333"
            }
          }}
      }
    ]
  };

  // 渲染图表
  Chart.setOption(options);
}
function getApprovalTypeList (data: any) {
  let approvalType=applyType.find((e)=>data==e.value);
  let approvalTypeTitle=approvalType==undefined?'':approvalType.text;
  return approvalTypeTitle;
}
onBeforeMount(()=>{
  setTimeout(()=>countNumber(),800)
})
defineExpose({
  countNumber, applyCount
})
</script>

<style scoped>

.calendarCard {
  /*padding:5px;*/
  padding-top: 4px;
  padding-bottom: 4px;
  text-align: center;
  border-radius: 8px;
  border: 1px solid #999999;
}
.testMain {
  width: 95vw;
  height: 40vh;
}

</style>
