<template>
  <div class="applyCard" @click="goDetail(getData,code)">
    <van-row align="center">
      <van-col span="3">
        <van-image :src="overWorkIcon3"
                   v-if="getData.Type=='工作日'||getData.Type=='工作日（至晚上21点）'||getData.Type=='工作日（至晚上22点）'"
                   width="30" height="30"/>
        <van-image :src="overWorkIcon3" v-else-if="getData.Type=='延时'" width="30" height="30"/>
        <van-image :src="overWorkIcon4" v-else-if="getData.Type=='休息日'" width="30" height="30"/>
        <van-image :src="overWorkIcon6" v-else-if="getData.Type=='法定节假日'" width="30" height="30"/>
        <van-image :src="leaveIcon1" v-else-if="getData.Type=='事假'" width="30" height="30"/>
        <van-image :src="leaveIcon2" v-else-if="getData.Type=='病假'" width="30" height="30"/>
        <van-image :src="leaveIcon3" v-else-if="getData.Type=='年假'" width="30" height="30"/>
        <van-image :src="leaveIcon5" v-else-if="getData.Type=='婚假'" width="30" height="30"/>
        <van-image :src="leaveIcon6" v-else-if="getData.Type=='产假'" width="30" height="30"/>
        <van-image :src="leaveIcon7" v-else-if="getData.Type=='陪产假'" width="30" height="30"/>
        <van-image :src="leaveIcon10" v-else width="30" height="30"/>

      </van-col>
      <van-col span="11">
        <span class="f14">{{ getData.approveReason }}</span>
      </van-col>
      <van-col span="10">
        <div class="skyTip" v-if="getData.Type=='工作日'||getData.Type=='工作日（至晚上21点）'||getData.Type=='工作日（至晚上22点）'||getData.Type=='延时'||getData.Type=='事假'">
          {{ getData.Type }}
        </div>
        <div class="tomatoTip" v-if="getData.Type=='休息日'||getData.Type=='病假'">
          {{ getData.Type }}
        </div>
        <div class="greenTip" v-if="getData.Type=='法定节假日'||getData.Type=='年假'">
          {{ getData.Type }}
        </div>
        <div class="roseTip" v-if="getData.Type=='婚假'">
          {{ getData.Type }}
        </div>
        <div class="violetTip" v-if="getData.Type=='产假'">
          {{ getData.Type }}
        </div>
        <div class="skyTip" v-if="getData.Type=='陪产假'">
          {{ getData.Type }}
        </div>
        <div class="indigoTip" v-if="getData.Type=='丧假'">
          {{ getData.Type }}
        </div>
        <!--        <div class="soilTip" v-else >{{getData.Type}}</div>-->
      </van-col>
    </van-row>
    <div class=" ml10 f14 pt10">
      <van-row>
        <van-col>
          <span class="graytxt"><span v-if="code=='work'">加班</span><span v-if="code=='leave'">请假</span>时长：</span>
        </van-col>
        <van-col>
          {{ getData.duration }} 天
        </van-col>
      </van-row>
      <van-row>
        <van-col>
          <span class="graytxt">开始时间：</span>
        </van-col>
        <van-col>
          {{ getData.startTime }}
        </van-col>
      </van-row>
      <van-row>
        <van-col>
          <span class="graytxt">结束时间：</span>
        </van-col>
        <van-col>
          {{ getData.endTime }}
        </van-col>
      </van-row>
      <van-row>
        <van-col>
          <span class="graytxt">申请单号：</span>
        </van-col>
        <van-col>
          {{ getData.approveNo }}
        </van-col>
      </van-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
import overWorkIcon3 from '../../../../assets/icon/kaoqin_icon_03.png'
import overWorkIcon4 from '../../../../assets/icon/kaoqin_icon_04.png'
import overWorkIcon6 from '../../../../assets/icon/kaoqin_icon_06.png'
import leaveIcon1 from '../../../../assets/icon/qingjia_icon_01.png'
import leaveIcon2 from '../../../../assets/icon/qingjia_icon_02.png'
import leaveIcon3 from '../../../../assets/icon/qingjia_icon_03.png'
import leaveIcon5 from '../../../../assets/icon/qingjia_icon_05.png'
import leaveIcon6 from '../../../../assets/icon/qingjia_icon_06.png'
import leaveIcon7 from '../../../../assets/icon/qingjia_icon_07.png'
import leaveIcon10 from '../../../../assets/icon/qingjia_icon_10.png'
import {useRouter} from "vue-router";

interface Props {
  getData?: Object;
  code?: String;
}

defineProps<Props>()
const router = useRouter();
const goDetail = (item: any, code: any) => {
  const data: any = {
    'approveType': code,
    'approveNo': item.approveNo,
  };
  sessionStorage.setItem('applicationDetail', JSON.stringify(data))
  router.push('/applicationDetail')
}
</script>

<style scoped>
.applyCard {
  margin: 10px;
  padding: 10px;
  border-radius: 10px;
  background: #ffffff;
}

.skyTip {
  padding: 8px;
  border-radius: 6px;
  color: #43A1FF;
  background: rgba(67, 161, 255, 0.3);
  text-align: center;
  font-size: 12px;
}

.tomatoTip {
  padding: 8px;
  border-radius: 6px;
  color: #FF6231;
  background: rgba(255, 61, 0, 0.3);
  text-align: center;
  font-size: 12px;
}

.greenTip {
  padding: 8px;
  border-radius: 6px;
  color: #46D2B1;
  background: rgba(70, 210, 177, 0.3);
  text-align: center;
  font-size: 12px;
}

.roseTip {
  padding: 8px;
  border-radius: 6px;
  color: #F52323;
  background: rgba(245, 35, 35, 0.3);
  text-align: center;
  font-size: 12px;
}

.violetTip {
  padding: 8px;
  border-radius: 6px;
  color: #868CFF;
  background: rgba(245, 35, 35, 0.3);
  text-align: center;
  font-size: 12px;
}

.indigoTip {
  padding: 8px;
  border-radius: 6px;
  color: #7B7DFF;
  background: rgba(134, 140, 255, 0.3);
  text-align: center;
  font-size: 12px;
}

.soilTip {
  padding: 8px;
  border-radius: 6px;
  color: #F5A623;
  background: rgba(245, 166, 35, 0.3);
  text-align: center;
  font-size: 12px;
}
</style>
