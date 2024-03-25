<template>
  <div class="m10">
    <van-row justify="space-between" align="center">
      <van-col>
        <span class="ml5 f16 bold">数据总览</span>
      </van-col>
      <van-col>
        <van-row align="center">
          <span class="font_gray f14" @click="showPicker=true">{{ fieldValue }}</span>
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
    <div v-if="countCorprateNum.totalTicket!==0">
      <div id="main" class="testMain"></div>
      <div>
        <div class="numCard">
          <van-row justify="space-between" align="center">
            <van-col>
              <van-row align="center">
                <van-image :src="chuxingIcon" width="30" height="30"/>
                <span class="f14 bold ml5">总出行数</span>
              </van-row>
            </van-col>
            <van-col>
              <span class="f14 bold mr5">{{ countCorprateNum.totalTicket }}次</span>
            </van-col>
          </van-row>
        </div>
        <div class="numCard">
          <van-row justify="space-between" align="center">
            <van-col>
              <van-row align="center">
                <van-image :src="feijiIcon" width="30" height="30"/>
                <div>
                  <div class="f14 bold ml5">飞机出行</div>
                  <div class="f12 ml5 graytxt">
                    在时间期限内该方式使用总数占比{{ numPercentage(countCorprateNum.flightNum).toFixed(2) }}%
                  </div>
                </div>
              </van-row>
            </van-col>
            <van-col>
              <span class="f14 bold mr5">{{ countCorprateNum.flightNum }}次</span>
            </van-col>
          </van-row>
        </div>
        <div class="numCard">
          <van-row justify="space-between" align="center">
            <van-col>
              <van-row align="center">
                <van-image :src="huocheIcon" width="30" height="30"/>
                <div>
                  <div class="f14 bold ml5">火车出行</div>
                  <div class="f12 ml5 graytxt">
                    在时间期限内该方式使用总数占比{{ numPercentage(countCorprateNum.trainNum).toFixed(2) }}%
                  </div>
                </div>
              </van-row>
            </van-col>
            <van-col>
              <span class="f14 bold mr5">{{ countCorprateNum.trainNum }}次</span>
            </van-col>
          </van-row>
        </div>
        <div class="numCard">
          <van-row justify="space-between" align="center">
            <van-col>
              <van-row align="center">
                <van-image :src="jiudianIcon" width="30" height="30"/>
                <div>
                  <div class="f14 bold ml5">酒店预定</div>
                  <div class="f12 ml5 graytxt">
                    在时间期限内该方式使用总数占比{{ numPercentage(countCorprateNum.hotelNum).toFixed(2) }}%
                  </div>
                </div>
              </van-row>
            </van-col>
            <van-col>
              <span class="f14 bold mr5">{{ countCorprateNum.hotelNum }}次</span>
            </van-col>
          </van-row>
        </div>
      </div>
    </div>
    <div v-else>
      <van-row justify="center" align="center" class="tc">
        <div>
          <img style="height:30vh;width:80vw;object-fit: contain" src="../../../../assets/img/zanwupiaoju.png"
               alt="该时间段没有查询到出行相关的数据" v-cloak/>
        </div>
        <div style="color: rgba(23,26,29,0.40); font-size: 14px">该时间段没有查询到出行相关的数据～</div>
      </van-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
import chuxingIcon from '../../../../assets/icon/chuxing.png'
import feijiIcon from '../../../../assets/icon/feiji.png'
import huocheIcon from '../../../../assets/icon/huoche.png'
import jiudianIcon from '../../../../assets/icon/jiudian.png'
import xiaLaIcon from '../../../../assets/icon/xiala.png'
import calendarIcon from '../../../../assets/icon/kaoqin_icon_05.png'
import moment from "moment/moment"; // 导入时间格式处理组件
import {inject, onMounted, ref} from "vue";
import {countCorporateTravelSpending} from "../../../../services/statement"; // 引入所需要引用的接口
import * as echarts from "echarts"; //引用echarts图表
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
const show = ref(false); //控制选择器的显隐
const startDate = ref('2023-01-01'); // 数据起始时间
const endDate = ref('2023-12-31'); // 数据截止时间
const fieldValue = ref('本年度'); //  选项信息的初始值
const showPicker = ref(false); // 控制日期时间的选项
const pieChats = ref([]); // 饼图数据存放的数组
const autoDateShow = ref(true); // 自定义时间or 规范设定的时间
const countCorprateNum = ref({
  totalTicket: ''
}); // 用于接收接口中用于展示数据
let userInfoData: any = inject("userInfo"); //获取用户信息
// 访问接口所需要传递过去的值
let payload = {
  "corpCode": userInfoData.userInfo.corpCode,
  'startTime': '',
  'endTime': '',
}
//用于
const notAutoDateConfirm = (values: any) => {
  const [start, end] = values;
  startDate.value = moment(start).format("yyyy-MM-DD");
  endDate.value = moment(end).format("yyyy-MM-DD");
  getCountCorporate();
  show.value = false;
};
// 时间选择器更改时间选择期限
const onConfirm = (value: any) => {
  showPicker.value = false;
  fieldValue.value = value.selectedOptions[0].text;
  // 判断是否为自定义，不是自定义直接调用接口获取对应的数据信息
  if (value.selectedOptions[0].auto == '0') {
    autoDateShow.value=true;
    startDate.value = value.selectedOptions[0].startTime;
    endDate.value = value.selectedOptions[0].endTime;
    getCountCorporate();
  } else {
    autoDateShow.value=false;
  }
};
// 调用接口获取次数统计的数据
const getCountCorporate = async () => {
  countCorprateNum.value = {
    totalTicket: ''
  };
  pieChats.value = [];
  // 传递前赋值将时间格式赋值过去
  payload.startTime = startDate.value;
  payload.endTime = endDate.value;
  let result: any = await countCorporateTravelSpending(payload);
  // 赋值给对应的对象用于存储展示
  countCorprateNum.value = result.data;
  // 根据判断的数值追加到饼图数组中
  if (result.data.trainNum != '0') {
    pieChats.value.push({value: result.data.trainNum, name: '火车'} as never);
  }
  if (result.data.flightNum != '0') {
    pieChats.value.push({value: result.data.flightNum, name: '飞机'} as never);
  }
  if (result.data.hotelNum != '0') {
    pieChats.value.push({value: result.data.hotelNum, name: '酒店'} as never);
  }
  init();
};

// 绘制饼图
function init() {
  // 基于准备好的dom，初始化echarts实例
  let Chart = echarts.init(document.getElementById("main"));
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

// 计算次数的百分比
const numPercentage = (num: any) => {
  // 获取整体数组
  let countNum = countCorprateNum.value;
  let countN: any = countNum.totalTicket;
  return (num / countN) * 100;

}
onMounted(() => {
  getCountCorporate();
});
</script>

<style scoped>

.testMain {
  width: 80vw;
  height: 40vh;
}

.calendarCard {
  /*padding:5px;*/
  padding-top: 4px;
  padding-bottom: 4px;
  text-align: center;
  border-radius: 8px;
  border: 1px solid #999999;
}

.numCard {
  margin: 10px 5px;
  padding: 10px;
  background: #ffffff;
  border-radius: 5px;
}
</style>
