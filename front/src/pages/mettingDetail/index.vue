<template>
  <van-nav-bar
    title="会议详情"
    left-text="返回"
    left-arrow
    @click-left="onClickLeft"
  />
  <div class="card-metting">
    <div class="name">L.的快速会议</div>
    <div class="time">03月29日16：55-17：55（GMT+8）</div>
    <div>会议号：251 963 045</div>
    <div>创建人：L.</div>
    <div class="metting-time">
      <div class="metting_left">
        <van-image :src="startTime" width="30px" height="30px"></van-image>
        <div>03/29 16:55</div>
        <div class="c-gray">最近入会</div>
      </div>
      <div class="metting_right">
        <van-image :src="duration" width="30px" height="30px"></van-image>
        <div>00：01：22</div>
        <div class="c-gray">参会时长</div>
      </div>
    </div>
  </div>
  <div class="list" v-for="item in list">
    <div class="left">
      <div class="time">{{ item.time }}</div>
      <div>{{ item.date }}</div>
    </div>
    <div class="right" style="line-height: 20px">
      <div>{{ item.long }}</div>
      <van-image
        :src="start"
        width="20px"
        height="20px"
        style="vertical-align: middle; margin-left: 10px"
      ></van-image>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import startTime from "../../assets/img/start_time.png";
import start from "../../assets/img/start.png";
// import start from '../../assets/start.png';
import duration from "../../assets/img/duration.png";
interface ListItem {
  time: string;
  date: string;
  long: string;
}

const list = ref<Array<ListItem>>([
  {
    time: "20240329_170029",
    date: "2024/3/29",
    long: "00:00:14",
  },
]);
const onClickLeft = () => history.back();
</script>
<style lang="less" scoped>
.card-metting {
  background-color: #fff;
  border-radius: 8px;
  margin: 20px;
  padding: 20px;
  font-size: 14px;
  div {
    margin-bottom: 10px;
  }
  .metting-time {
    display: flex;
    justify-content: space-around;
    margin-top: 20px;
    .metting_left {
      text-align: center;
      .c-gray {
        color: #7b7b7b;
      }
    }
    .metting_right {
      text-align: center;
      .c-gray {
        color: #7b7b7b;
      }
    }
  }
}
.list {
  display: flex;
  justify-content: space-between;
  padding: 20px;
  margin: 20px;
  background-color: #fff;
  border-radius: 8px;
  align-items: center;
  font-size: 12px;
  color: #7b7b7b;
  .right {
    display: flex;
    text-align: center;
  }
  .time {
    font-size: 14px;
    font-weight: 600;
    margin-bottom: 10px;
    color: #000
  }
}
.name {
  font-size: 18px;
  font-weight: 600;
}
</style>
