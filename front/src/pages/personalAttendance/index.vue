<template>
  <div>
    <nav-bar title="个人考勤统计"/>
    <div class="head">
      <div class="headSwitch">
        <van-row justify="space-around" align="center">
          <van-col span="12">
            <div :class="attendanceCode=='work'?'checkedCard':''" @click="toggleCode('work')">加班统计</div>
          </van-col>
          <van-col span="12">
            <div :class="attendanceCode=='leave'?'checkedCard':''" @click="toggleCode('leave')">请假统计</div>
          </van-col>
        </van-row>
      </div>
      <div style="padding: 8px">
        <van-row justify="space-around" align="center">
          <van-col span="6">
            <div class="arrowCard" :class="dateString==moment(minDate).format('yyyy-MM')?'arrowNoCard':''"
                 @click="lastMouth">
              <van-icon name="arrow-left"/>
              <span class="f14">上个月</span>
            </div>
          </van-col>
          <van-col span="10">
            <div class="calendarCard">
              <van-row justify="center" align="center">
                <van-col span="4">
                  <van-image :src="calendarIcon" width="20" height="20"/>
                </van-col>
                <van-col span="10">
                  <span class="f14" @click="show=true">{{ dateString }}</span>
                </van-col>
              </van-row>
              <van-popup v-model:show="show" position="bottom">
                <van-date-picker v-model="currentDate"
                                 title="选择年月"
                                 :min-date="minDate"
                                 :max-date="maxDate"
                                 :formatter="formatter"
                                 :columns-type="columnsType"
                                 @confirm="onConfirm" @cancel="show = false"/>
              </van-popup>
            </div>

          </van-col>
          <van-col span="6">
            <div class="arrowCard" :class="dateString==moment(maxDate).format('yyyy-MM')?'arrowNoCard':''"
                 @click="nextMouth">
              <span class="f14">下个月</span>
              <van-icon name="arrow"/>
            </div>
          </van-col>
        </van-row>
      </div>
      <div>
        <van-row justify="space-around">
          <van-col span="12">
            <div class="check">
              <div class="pt10 pl15">
                <div v-if="attendanceCode=='work'">
                  <van-row align="center">
                    <van-image :src="overWorkIcon1" width="20" height="20"/>
                    <span class="bold ml5">加班天数</span>
                  </van-row>
                </div>
                <div v-else>
                  <van-row align="center">
                    <van-image :src="overWorkIcon1" width="20" height="20"/>
                    <span class="bold ml5">请假天数</span>
                  </van-row>
                </div>

              </div>
              <div class=" pl15" style="padding-top: 8px">
                <span class="f26 bold ml20">
                  {{ attendFate }}
                </span>
                <span class="f12 font_gray ml5">天</span>
              </div>
            </div>
          </van-col>
          <van-col span="12">
            <div class="num">
              <div class="pt10 pl15">
                <div v-if="attendanceCode=='work'">
                  <van-row align="center">
                    <van-image :src="overWorkIcon2" width="20" height="20"/>
                    <span class="ml5 bold">加班次数</span>
                  </van-row>
                </div>
                <div v-else>
                  <van-row align="center">
                    <van-image :src="overWorkIcon2" width="20" height="20"/>
                    <span class="ml5 bold">请假次数</span>
                  </van-row>
                </div>
              </div>
              <div class=" pl15" style="padding-top: 8px">
                <span class="f26 bold ml20">
                  {{ attendDegree }}
                </span>
                <span class="f12 font_gray ml5">次</span>
              </div>
            </div>
          </van-col>
        </van-row>
      </div>
    </div>
    <div class="mt10">
      <van-row>
        <van-col>
          <span class="ml10 bold f16">
            <span v-if="attendanceCode=='work'">加班</span>
            <span v-if="attendanceCode=='leave'">请假</span>列表
          </span>
        </van-col>
        <van-col><span class="f12 font_gray ml10">{{ dateString }}-01 至 {{ getEndDate() }}</span></van-col>
      </van-row>
      <div v-if="attendanceCardList.length==0">
        <van-row justify="center" align="center" class="tc">
          <div>
            <img style="height:50vh;width:80vw;object-fit: contain" src="../../assets/img/zanwupiaoju.png" v-cloak/>
          </div>
          <div style="color: rgba(23,26,29,0.40)">未查询到当月的申请信息</div>
        </van-row>
      </div>
      <div v-for="item in attendanceCardList">
        <attendance-card :get-data="item" :code="attendanceCode"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import NavBar from "../../components/navBar/index.vue";
import AttendanceCard from './component/attendanceCard/index.vue'
import calendarIcon from '../../assets/icon/kaoqin_icon_05.png'
import overWorkIcon1 from '../../assets/icon/kaoqin_icon_01.png'
import overWorkIcon2 from '../../assets/icon/kaoqin_icon_02.png'
import moment from "moment";
import {inject, onMounted, ref} from "vue";
import {countWorkOvertimeOrAskForLeave} from "../../services/statement";
import {showDialog} from "vant";
const attendanceCode = ref('work'); //控制组件显隐的code
const attendFate = ref(''); // 考勤天数
const attendDegree = ref('');// 考勤次数
const show = ref(false); //控制选择器的显隐
const dateString = ref(''); //选择器的点击的值
const nowDate = new Date(); // 获取当前时间
const currentDate = ref(['2023', '08']); // 初始化时间选择器的选择的值
const columnsType = ['year', 'month']; // 规范选择的格式
const minDate = new Date(2023, 0, 1); //最小时间
const maxDate = new Date();
let userInfoData: any = inject("userInfo"); //获取审批数组
const attendanceCardList = ref([]); // 用于存放数组列表
// 格式化选项
const formatter = (type: any, option: any) => {
  if (type === 'year') {
    option.text += '年';
  }
  if (type === 'month') {
    option.text += '月';
  }
  return option;
};

// 访问接口所得到的值
let payload = {
  "corpCode": userInfoData.userInfo.corpCode,
  'empCode': userInfoData.userInfo.empCode,
  'startTime': '',
  'endTime': '',
  'dataType': 'work'
}
// 点击调整时间去上个月
const lastMouth = () => {
  let minMouth = moment(minDate,moment.RFC_2822).format("YYYY-MM");
  if (minMouth == dateString.value) {
    showDialog({
      message: '很抱歉，无法获取到更早的数据了',
    }).then(() => {
      // on close
    });
  } else {
    let nowDate = new Date(dateString.value);
    let lastDate = nowDate.setMonth(nowDate.getMonth() - 1);
    let aDate = moment(lastDate).format('YYYY-MM');
    currentDate.value=[moment(lastDate,moment.ISO_8601).format('YYYY'),moment(lastDate,moment.ISO_8601).format('MM')]
    dateString.value = aDate;
    getOverWork();
  }
}
// 点击调整时间去下个月
const nextMouth = () => {
  let maxMouth = moment(maxDate).format("YYYY-MM");
  if (maxMouth == dateString.value) {
    showDialog({
      message: '已经是当前最新月份了',
    }).then(() => {
      // on close
    });
  } else {
    let nowDate = new Date(dateString.value);
    let lastDate = nowDate.setMonth(nowDate.getMonth() + 1);
    let aDate = moment(lastDate).format('YYYY-MM');
    currentDate.value=[moment(lastDate,moment.ISO_8601).format('YYYY'),moment(lastDate).format('MM')]
    dateString.value = aDate;
    getOverWork();
  }
}
// 切换申请单类型
const toggleCode = (code: any) => {
  attendanceCode.value = code;
  payload.dataType = code;
  dateString.value = moment(nowDate).format("YYYY-MM");
  getOverWork();
}
// 选择时间
const onConfirm = (selectedValues: any) => {
  dateString.value = selectedValues.selectedValues.join('-');
  getOverWork();
  show.value = false;
}
// 众所周知，一个月的开始是不会变的都是从1号开始，但是一个月的结束却有四种可能，
// 所以这个方法就根据老谚语来处理月份的结束日期，此处用于展示以及调用接口
const getEndDate = () => {
  let nowDate = new Date(dateString.value);
  nowDate.setMonth(nowDate.getMonth() + 1);
  nowDate.setDate(0);
  let lastDay = nowDate.toLocaleDateString()
  return moment(lastDay, 'YYYY-MM-DD').format('YYYY-MM-DD')
}
// 调用接口获取加班统计的数据
const getOverWork = async () => {
  payload.startTime = `${dateString.value}-01`;
  payload.endTime = getEndDate();
  let result: any = await countWorkOvertimeOrAskForLeave(payload);
  attendFate.value = result.data.fate;
  attendDegree.value = result.data.degree;
  attendanceCardList.value = result.data.list;
}
// 初始化需要处理的信息
onMounted(() => {
  dateString.value = moment(nowDate,moment.ISO_8601).format("YYYY-MM");
  getOverWork();
});

</script>

<style scoped>
.head {
  width: 100%;
  height: 200px;
  padding-top: 2px;
  background: #ffffff;
}

.headSwitch {
  margin: 10px;
  background: #F9F9F9;
  padding: 5px;
  text-align: center;
  font-size: 14px;
  border-radius: 8px;
}

.checkedCard {
  color: #0088ff;
  padding: 5px;
  font-size: 14px;
  width: 100%;
  border-radius: 8px;
  background: #ffffff;
}

.arrowNoCard {
  background: #F6F8FA;
}

.arrowCard {
  padding: 5px;
  text-align: center;
  border-radius: 8px;
  border: 1px solid #999999;
}

.calendarCard {
  /*padding:5px;*/
  padding-top: 4px;
  padding-bottom: 4px;
  text-align: center;
  border-radius: 8px;
  border: 1px solid #999999;
}

.check {
  background: url("../../assets/img/kaoqin_bg_01.png") no-repeat;
  margin-left: 5px;
  margin-right: 5px;
  background-size: 100% 80px;
  height: 80px;
}

.num {
  background: url("../../assets/img/kaoqin_bg_02.png") no-repeat;
  margin-left: 5px;
  margin-right: 5px;
  background-size: 100% 80px;
  height: 80px;
}
</style>
