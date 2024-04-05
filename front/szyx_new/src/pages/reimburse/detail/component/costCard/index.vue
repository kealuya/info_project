<template>
  <div class="costCard">
    <van-row justify="space-between" align="center">
      <van-col>
        <div class="f16">{{data.reason}}</div>
      </van-col>
      <van-col>
        <div class="bold orange f18">¥{{data.money}}</div>
      </van-col>
    </van-row>
    <div class="mt5">
      <van-row align="center">
        <van-col>
          <van-row align="center">
            <van-icon :name="vectorIcon" />
            <div class="ml10">{{data.expenseTypeName}}</div>
          </van-row>
        </van-col>
        <van-col class="ml10">
          <van-row align="center">
            <van-icon :name="timeIcon" />
            <div class="ml10">{{data.createTime}}</div>
          </van-row>
        </van-col>
      </van-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
import vectorIcon from '../../../../../assets/icon/Vector.png';
import timeIcon from '../../../../../assets/icon/baoxiao_icon_02.png'
interface Props{
  data?:object;

}
defineProps<Props>()
</script>

<style scoped>
.costCard{
  margin: 10px 0px;
  padding: 10px;
  border:1px solid #0088ff;
  border-radius: 8px;
}
</style>
