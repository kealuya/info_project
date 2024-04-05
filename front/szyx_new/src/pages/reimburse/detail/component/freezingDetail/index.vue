<template>
  <div class="freezingCard">
    <van-row justify="space-between" align="center">
      <van-col>
        <div class="bold f16"> {{data.costName}}</div>
      </van-col>
      <van-col>
        <div v-if="data.type=='0'" class="errorTip">冻结经费</div>
        <div v-else class="sucesssTip">解冻经费</div>
      </van-col>
    </van-row>
    <van-row align="center">
      <van-icon :name="jineIcon" size="15" />
      <span class="ml5">冻结金额:</span>
      <span class="ml5 orange bold f14">¥ {{data.je}}</span>
    </van-row>
    <van-row align="center" class="mt5">
      <van-icon :name="vectorIcon" size="15" />
      <span class="ml5">操作时间:</span>
      <span class="ml5">{{data.creationTime}}</span>
    </van-row>
  </div>
</template>

<script lang="ts" setup>
import jineIcon from '../../../../../assets/icon/jine_icon.png';
import vectorIcon from '../../../../../assets/icon/Vector.png';
interface Props{
  data?:object;
}
defineProps<Props>()
</script>

<style scoped>
.freezingCard{
  margin: 10px 0px;
  padding: 10px;
  border:1px solid #0088ff;
  border-radius: 8px;
}
.sucesssTip{
  margin: 5px;
  padding: 5px;
  border: #0088ff solid 1px;
  border-radius: 10px;
  font-size: 12px;
  background: rgba(0, 128, 255, 0.3);
  color:#0088ff;
}
.errorTip{
  margin: 5px;
  padding: 5px;
  border: tomato solid 1px;
  font-size: 12px;
  border-radius: 10px;
  background: rgba(250, 97, 49, 0.3);
  color:tomato;
}
</style>
