<template>
  <nav-bar title="企业数据统计"/>
  <div class="ml10 mr10 mb10">
    <van-row justify="center" align="center">
      <van-col span="8" v-for="item in echartList">
        <div :class="code===item.value?'checkedCard':'checkNoCard'" @click="toggleEchartsKey(item.value)">
          {{ item.title }}
        </div>
      </van-col>
    </van-row>
  </div>
  <div v-if="code==='goTrips'" >
    <go-trips-echarts-panel />
  </div>
  <div v-if="code==='travel'">
    <travel-amount-echarts-panel />
  </div>
  <div v-if="code=='enterprise'">
    <enterprise-echarts-panel />
  </div>
</template>

<script lang="ts" setup>
import NavBar from '../../components/navBar/index.vue'
import {ref} from "vue";
import GoTripsEchartsPanel from './component/goTripsEchartsPanel/index.vue'
import TravelAmountEchartsPanel from "./component/travelAmountEchartsPanel/index.vue";
import EnterpriseEchartsPanel from './component/enterpriseEchartsPanel/index.vue';
let code = ref('goTrips'); // 用于控制图表显示隐藏的标识
// 用于显示遍历的数组
let echartList = ref([
  {'title': '出行次数统计', 'value': 'goTrips'},
  {'title': '差旅费用统计', 'value': 'travel'},
  {'title': '企业费用统计', 'value': 'enterprise'},
]);
// 标识的图表显示的key
const toggleEchartsKey = (codeKey: any) => {
  code.value = codeKey;
}
</script>

<style scoped>
.checkedCard {
  margin: 5px;
  padding: 10px;
  font-size: 14px;
  color: #fff;
  background: #0088ff;
  border-radius: 5px;
  text-align: center;
}

.checkNoCard {
  margin: 5px;
  padding: 10px;
  font-size: 14px;
  color: #fff;
  background: rgba(0, 128, 255, 0.5);
  border-radius: 5px;
  text-align: center;
}
</style>
