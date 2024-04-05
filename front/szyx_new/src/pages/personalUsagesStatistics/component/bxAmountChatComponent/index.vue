<template>
  <div class="m10">
    <van-row justify="space-between" align="center">
      <van-col>
        <span class="ml5 f16 bold">数据总览</span>
      </van-col>
      <van-col>
        <van-row align="center">
          <span class="font_gray f14" @click="showPicker=true">{{ fieldValue }}年</span>
          <van-icon class="ml5 mr10" :name="xiaLaIcon" color="#999999"/>
        </van-row>
      </van-col>
    </van-row>
    <van-popup v-model:show="showPicker" round position="bottom">
      <van-picker
          :columns="columns"
          @cancel="showPicker = false"
          @confirm="onConfirm"
      />
    </van-popup>
    <div>
      <div id="main" class="testMain"></div>
    </div>
    <div class="">
      <van-row justify="center" align="center">
        <div class="toggleCard">
          <van-row>
            <van-col>
              <div class="f14 typeCard" :class="dataType=='dailyBx'?'checkDaiyCard':'daiyCard'"
                   @click="toggleChartsCode('dailyBx')">日常报销
              </div>
            </van-col>
            <van-col>
              <div class="f14 typeCard" :class="dataType=='travelBx'?'checkTravelCard':'noCard'"
                   @click="toggleChartsCode('travelBx')">差旅补助报销
              </div>
            </van-col>
            <van-col>
              <div class="f14 typeCard" :class="dataType=='loan'?'checkLoanCard':'loanCard'"
                   @click="toggleChartsCode('loan')">借款申请
              </div>
            </van-col>
          </van-row>
        </div>
      </van-row>
    </div>
    <div class="mt10">
      <van-row align="center">
        <van-col span="8" v-for="item in chartCardList">
          <div class="echartCard tc">
            <div class="bold f16">¥ {{item.money}}</div>
            <div class="font_gray f12 mt5">{{item.createDt}}</div>
          </div>
        </van-col>
      </van-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
import xiaLaIcon from '../../../../assets/icon/xiala.png'
import {onMounted, ref, inject, defineEmits} from "vue";
import moment from "moment";
import {countTravelSpendingStatistics} from "../../../../services/statement";
import * as echarts from "echarts";
const nowDate = new Date(); // 当前时间
const startYear = ref('2021'); // 系统起始时间
const fieldValue = ref('2023'); //  选项信息的初始值
const columns = ref([]); // 过往年度数据列表
const showPicker = ref(false); // 控制年份的选项
const dataType = ref('dailyBx'); // 折线图数据种类类型切换
let userInfoData: any = inject("userInfo"); //获取用户信息
const chartCardList = ref([]); // 存储卡片数组
const xAxisList=ref([]); // 下方x轴
const yAxisList=ref([]); // 折线图数据
const emit=defineEmits(['overData']) // 触发更新父组件的值
// 获取数据的提交参数
const revokeInfo = {
  "corpCode": userInfoData.userInfo.corpCode,
  'startTime': fieldValue.value,
  'dataType': dataType.value,
}
// 时间选择器更改年度选择期限
const onConfirm = (value: any) => {
  showPicker.value = false;
  fieldValue.value = `${value.selectedOptions[0].text}`;
  getCountCorporate();
  // 判断是否为自定义，不是自定义直接调用接口获取对应的数据信息
};
// 填充当前年份至系统起始年份
const formatYear = () => {
  let start = parseInt(startYear.value);
  let end = parseInt(fieldValue.value);
  for (let i = start; i <= end; i++) {
    columns.value.unshift({text: i, value: i} as never)
  }
}
// 调用接口获取次数统计的数据
const getCountCorporate = async () => {
  xAxisList.value=[];
  yAxisList.value=[];
  // 传递前赋值将时间格式赋值过去
  revokeInfo.startTime = fieldValue.value;
  revokeInfo.dataType = dataType.value;
  let result: any = await countTravelSpendingStatistics(revokeInfo);
  // 赋值给对应的对象用于存储展示
  chartCardList.value = result.data.totalByTime;
  let dataList=result.data.totalByTime;
  for(let item in dataList){
    xAxisList.value.push((dataList[item].createDt)as never);
    yAxisList.value.push((dataList[item].money)as never)
  }
  emit('overData',result.data.amount);
  // 根据判断的数值追加到折线图数组中
  init()
};
// 数组处理
// 切换图表的数据折线图
const toggleChartsCode = (data: any) => {
  dataType.value=data;
  getCountCorporate();
}

// 绘制饼图
function init() {
  // 基于准备好的dom，初始化echarts实例
  let Chart = echarts.init(document.getElementById("main"));
  // 绘制图表
  let options = {
    grid: {
      top: '20%',
      bottom: '5%',
      left: '2%',
      right: '2%',
      containLabel: true
    },
    tooltip: {
      trigger: 'item',
    },
    xAxis: {
      type: 'category',
      data: xAxisList.value,
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        data: yAxisList.value,
        type: 'line'
      }
    ]
  };

  // 渲染图表
  Chart.setOption(options);
}

onMounted(() => {
  setTimeout(()=>{
    fieldValue.value = moment(nowDate).format("yyyy");
    formatYear();
    getCountCorporate();
  },800)
});
</script>

<style scoped>
.testMain {
  width: 90vw;
  height: 40vh;
}

.toggleCard {
  text-align: center;
  padding: 0px 0px;
  background: #ffffff;
  border-radius: 5px;
}

.typeCard {
  text-align: center;
  padding: 5px 10px;
}

.checkDaiyCard {
  background: rgba(250, 97, 49, 0.6);
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
  color: #ffffff;
}
.daiyCard{
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
  background: #ffffff;
  color: #333333;
}
.checkTravelCard {
  color: #ffffff;
  background: rgba(0, 128, 255, 0.6);
}
.checkLoanCard {
  color: #ffffff;
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px;
  background: rgba(245, 166, 35, 0.6);
}
.loanCard{
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px;
  background: #ffffff;
  color: #333333;
}
.noCard {
  background: #ffffff;
  color: #333333;
}
.echartCard{
  background: #FFFFFF;
  color: #333333;
  border-radius: 10px;
  margin: 5px;
  padding:10px
}
</style>
