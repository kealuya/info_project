<template>
  <div>
    <van-row justify="center" align="center">
       <div class="toggleCard">
        <van-row>
          <van-col>
            <div class="f14 typeCard"
                :class="dataType=='amount'?'checkAmountCard':'amountCard'"
                @click="toggleChartsCode('amount')">查看金额</div>
          </van-col>
          <van-col>
            <div class="f14 typeCard"
                 :class="dataType=='total'?'checkTotalCard':'totalCard'"
                 @click="toggleChartsCode('total')">查看次数</div>
          </van-col>
        </van-row>
       </div>
    </van-row>
    <div class="m10">
      <van-row justify="space-between" align="center">
        <van-col>
          <span v-if="dataType=='amount'">累计出行金额</span>
          <span v-if="dataType=='total'">累计出行次数</span>
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
      <div  v-show="dataType=='amount'" id="main" class="testMain"></div>
      <div  v-show="dataType=='total'" id="pieMain" class="testMain"></div>
      <div  v-show="dataType=='amount'">
       <div class="numCard">
         <van-row justify="space-between" align="center">
           <van-col>
             <van-row align="center">
               <van-image :src="chuxingIcon" width="30" height="30"/>
               <span class="f14 bold ml5">总出行数</span>
             </van-row>
           </van-col>
           <van-col>
             <span class="f14 bold mr5">{{ travelInfo.moneyTotal }} 元</span>
           </van-col>
         </van-row>
       </div>
       <div class="numCard">
         <van-row justify="space-between" align="center">
           <van-col>
             <van-row align="center">
               <van-image :src="feijiIcon" width="30" height="30"/>
               <span class="f14 bold ml5">飞机出行</span>
             </van-row>
           </van-col>
           <van-col>
             <span class="f14 bold mr5">{{ travelInfo.flightMoney }} 元</span>
           </van-col>
         </van-row>
       </div>
       <div class="numCard">
         <van-row justify="space-between" align="center">
           <van-col>
             <van-row align="center">
               <van-image :src="huocheIcon" width="30" height="30"/>
               <span class="f14 bold ml5">火车出行</span>
             </van-row>
           </van-col>
           <van-col>
             <span class="f14 bold mr5">{{ travelInfo.trainMoney }} 元</span>
           </van-col>
         </van-row>
       </div>
       <div class="numCard">
         <van-row justify="space-between" align="center">
           <van-col>
             <van-row align="center">
               <van-image :src="jiudianIcon" width="30" height="30"/>
               <span class="f14 bold ml5">酒店预定</span>
             </van-row>
           </van-col>
           <van-col>
             <span class="f14 bold mr5">{{ travelInfo.hotelMoney }} 元</span>
           </van-col>
         </van-row>
       </div>
      </div>
      <div  v-show="dataType=='total'">
        <div class="numCard">
          <van-row justify="space-between" align="center">
            <van-col>
              <van-row align="center">
                <van-image :src="chuxingIcon" width="30" height="30"/>
                <span class="f14 bold ml5">总出行数</span>
              </van-row>
            </van-col>
            <van-col>
              <span class="f14 bold mr5">{{ travelInfo.totalTicket }} 次</span>
            </van-col>
          </van-row>
        </div>
        <div class="numCard">
          <van-row justify="space-between" align="center">
            <van-col>
              <van-row align="center">
                <van-image :src="feijiIcon" width="30" height="30"/>
                <span class="f14 bold ml5">飞机出行</span>
              </van-row>
            </van-col>
            <van-col>
              <span class="f14 bold mr5">{{ travelInfo.flightNum }} 次</span>
            </van-col>
          </van-row>
        </div>
        <div class="numCard">
          <van-row justify="space-between" align="center">
            <van-col>
              <van-row align="center">
                <van-image :src="huocheIcon" width="30" height="30"/>
                <span class="f14 bold ml5">火车出行</span>
              </van-row>
            </van-col>
            <van-col>
              <span class="f14 bold mr5">{{ travelInfo.trainNum }} 次</span>
            </van-col>
          </van-row>
        </div>
        <div class="numCard">
          <van-row justify="space-between" align="center">
            <van-col>
              <van-row align="center">
                <van-image :src="jiudianIcon" width="30" height="30"/>
                <span class="f14 bold ml5">酒店预定</span>
              </van-row>
            </van-col>
            <van-col>
              <span class="f14 bold mr5">{{ travelInfo.hotelNum }} 次</span>
            </van-col>
          </van-row>
        </div>
      </div>
      <van-popup v-model:show="showPicker" round position="bottom">
        <van-picker
            :columns="columns"
            @cancel="showPicker = false"
            @confirm="onConfirm"
        />
      </van-popup>
    </div>
  </div>
</template>

<script lang="ts" setup>
import rightArrowIcon from '../../../../assets/icon/xiala.png';
import calendarIcon from '../../../../assets/icon/kaoqin_icon_05.png'
import chuxingIcon from '../../../../assets/icon/chuxing.png'
import feijiIcon from '../../../../assets/icon/feiji.png'
import huocheIcon from '../../../../assets/icon/huoche.png'
import jiudianIcon from '../../../../assets/icon/jiudian.png'
import {onMounted, ref, inject, defineEmits, onBeforeMount} from "vue";
import moment from "moment";
import {countTravelSpendingStatistic} from "../../../../services/statement";
import * as echarts from "echarts";
const nowDate = new Date(); // 当前时间
const autoDateShow = ref(true); // 自定义时间or 规范设定的时间
const show = ref(false); //控制选择器的显隐
const xChartList=ref([]); // 柱状图的x轴的数据类别
const yChartList=ref([]); // 柱状图的y轴的数据
const pieChats = ref([]); // 饼图数据存放的数组
const travelInfo=ref({}); // 出行信息
const startDate = ref(moment(nowDate).subtract(1, 'years').format("yyyy-MM-DD")); // 数据起始时间
const endDate = ref(moment(nowDate).format("yyyy-MM-DD")); // 数据截止时间
const fieldValue = ref('近一年'); //  选项信息的初始值
const showPicker = ref(false); // 控制年份的选项
const dataType = ref('amount'); // 累计出行数据种类类型切换
let userInfoData: any = inject("userInfo"); //获取用户信息
const emit=defineEmits(['overData']) // 触发更新父组件的值
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
      formatter: '{b} : {c}元'
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
            formatter: "{c}"+"元",
            show: true,
            position: "top",
            textStyle: {
              fontWeight: "bolder",
              fontSize: "12",
              color: "#333"
            }
          }
        }
      }
    ]
  };

  // 渲染图表
  Chart.setOption(options);
}

// 绘制饼图
function initPie() {
  // 基于准备好的dom，初始化echarts实例
  let Chart = echarts.init(document.getElementById("pieMain"));
  // 绘制图表
  let options = {
    color: [
      '#7962F9',
      '#FF0037',
      '#0088ff'
    ],
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b} : {c}次 ({d}%)'
    },
    legend: {
      left: 'center'
    },
    series: [
      {
        name: '出行数据统计',
        type: 'pie',
        radius: '60%',
        label: {
          formatter: '{b}：{c}次',
          position: 'inner',
        },
        avoidLabelOverlap: false,
        labelLine: {
          show: false
        },
        data: pieChats.value,
      }
    ]
  };

  // 渲染图表
  Chart.setOption(options);
}
// 切换展示的图表
const toggleChartsCode = (data: any) => {
  dataType.value=data;
}
// 时间选择器更改时间选择期限
const onConfirm = (value: any) => {
  showPicker.value = false;
  fieldValue.value = value.selectedOptions[0].text;
  // 判断是否为自定义，不是自定义直接调用接口获取对应的数据信息
  if (value.selectedOptions[0].auto == '0') {
    autoDateShow.value=true;
    startDate.value = value.selectedOptions[0].startTime;
    endDate.value = value.selectedOptions[0].endTime;
    countTravel();
  } else {
    autoDateShow.value=false;
  }
};
//用于 自选的申请图表
const notAutoDateConfirm = (values: any) => {
  const [start, end] = values;
  startDate.value = moment(start).format("yyyy-MM-DD");
  endDate.value = moment(end).format("yyyy-MM-DD");
  countTravel();
  show.value = false;
};
// 获取累计出行的次数和钱数
const countTravel=async()=>{
  let revokeInfo={
    "empCode":userInfoData.userInfo.empCode,
    'corpCode':userInfoData.userInfo.corpCode,
    'startTime':`${startDate.value} 00:00:00`,
    'endTime':`${endDate.value} 23:59:59`,
  }
  let result: any = await countTravelSpendingStatistic(revokeInfo);
  travelInfo.value=result.data;
  xChartList.value=['飞机'as never,'火车' as never,'酒店' as never,];
  yChartList.value=[result.data.flightMoney as never,result.data.trainMoney as never, result.data.hotelMoney as never];
  pieChats.value=[];
  if (result.data.trainNum != '0') {
    pieChats.value.push({value: result.data.trainNum, name: '火车'} as never);
  }
  if (result.data.flightNum != '0') {
    pieChats.value.push({value: result.data.flightNum, name: '飞机'} as never);
  }
  if (result.data.hotelNum != '0') {
    pieChats.value.push({value: result.data.hotelNum, name: '酒店'} as never);
  }
  emit('overData',result.data.moneyTotal,result.data.totalTicket);
  init();
  initPie();
}
onBeforeMount(()=>{
  setTimeout(()=>countTravel(),800);
})
</script>

<style scoped>
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
.calendarCard {
  /*padding:5px;*/
  padding-top: 4px;
  padding-bottom: 4px;
  text-align: center;
  border-radius: 8px;
  border: 1px solid #999999;
}
.checkAmountCard{
  background: rgba(250, 97, 49, 0.6);
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
  color: #ffffff;
}
.amountCard{
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
  background: #ffffff;
  color: #333333;
}
.checkTotalCard{
  color: #ffffff;
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px;
  background: rgba(245, 166, 35, 0.6);
}
.totalCard{
  color: #333333;
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px;
  background: #ffffff;
}
.testMain {
  width: 95vw;
  height: 40vh;
}
.numCard {
  margin: 10px 5px;
  padding: 10px;
  background: #ffffff;
  border-radius: 5px;
}
</style>
