<template>
  <div class="weatherComponent">
    <van-row justify="end">
      <van-col span="12">
        <div v-if="weatherInfo" class="weather">
          <div class="weatherCard">
            <span>{{ weatherInfo.city }}  {{ weatherInfo.weather }} {{ weatherInfo.temperature }}℃</span></div>
        </div>
      </van-col>
    </van-row>
    <div class="dateBottom ml10 mb10">
      <van-row justify="space-between" align="end">
        <van-col span="12">
          <div>
            <span class="f26 f_white">{{ nowTime }}</span>
            <span class="f14 f_white">{{ weekTime }}</span>
          </div>
          <div>
            <span class="f12 f_white">{{ lizhiText(parseInt(Math.random() * (0 - 8) + 8)) }}</span>
          </div>
        </van-col>
        <van-col span="8">
          <van-row justify="space-between" align="center">
            <van-col span="12">
              <div>
                <span class="f26 f_white">{{ accountBook > 99 ? '99+' : accountBook }}</span>
              </div>
              <div class="f12 f_white" @click="showLedgerListEvent">账本 ></div>
            </van-col>
            <van-col span="12">
              <div class="f26 f_white">{{ msg > 99 ? '99+' : msg }}</div>
              <div class="f12 f_white" @click="gotoExamine">审批 ></div>
            </van-col>
          </van-row>
        </van-col>
      </van-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {useRouter} from "vue-router";
import {inject} from 'vue';

let userInfo: any = inject("userInfo");
const router = useRouter();

interface Props {
  weatherInfo?: any;
  nowTime?: string;
  weekTime?: string;
  accountBook?: any;
  msg?: any;
}

defineProps<Props>()

// 励志言论
const lizhiText = (num: any) => {
  var text = '';
  switch (num) {
    case 0:
      text = '行动是治愈恐惧的良药。';
      break;
    case 1:
      text = '自信的生命最美丽！';
      break;
    case 2:
      text = '心有多大，舞台就有多大';
      break;
    case 3:
      text = '良好的习惯是成功的保证';
      break;
    case 4:
      text = '彩虹风雨后，成功细节中';
      break;
    case 5:
      text = '绳锯木断，水滴石穿';
      break;
    case 6:
      text = '要成就大事，先做好小事';
      break;
    case 7:
      text = '不怕路远，就怕志短';
      break;
    case 8:
      text = '差之毫厘，谬以千里。';
      break;
    case 9:
      text = '前进不止，奋斗不息。';
      break;
    case 10:
      text = '有志者，事竟成;有心人，天不负。';
      break;
    case 11:
      text = '刻意追求理性真的有用，值得你付出一生的努力。';
      break;
    case 12:
      text = '把问题彻底想明白，问题就解决了一半。';
      break;
  }
  return text;
}

function gotoExamine() {
  let url = `http://122.9.41.215/fk_move/MapprList?userid=${userInfo.userInfo.empCode}&username=${userInfo.userInfo.userName}&school=${userInfo.userInfo.corpCode}&type=false&isDark=false#/`
  router.push({
    name: 'approvalDetailsIframe',
    params: {
      url: url,
      realHeight: document.documentElement.offsetHeight,
      titleName: '审批'
    }
  })
}

// 跳转至账本
function showLedgerListEvent() {
  router.push({name: 'accountBook'})
}
</script>

<style scoped>
.weatherComponent {
  display: flex;
  height: 100%;
  flex-direction: column;
  justify-content: space-between;
}

.weather {
  text-align: right;
  margin-top: 10px;
  margin-right: 10px;
  color: #ffffff;
  font-size: 16px;
}

.weatherCard {
  margin: 10px 5px;
  padding: 8px;
  background: rgba(245, 245, 245, 0.5);
  text-align: center;
  border-radius: 12px;
}

.dateBottom {
  width: 96vw;
}
</style>
