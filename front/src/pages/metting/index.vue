<template>
  <van-nav-bar
  title="会议列表"
  left-text="返回"
  left-arrow
  @click-left="onClickLeft"
/>
<!-- 搜索框 -->
<van-search v-model="value" placeholder="会议名称、会议号、发起人" />
<div @click="toDetail">
  <div class="data">
  <div>3月29日 周五</div>
  <div class="year">2024年</div>
</div>
<div class="list">
  <div class="list_left">
   <div>251 963 045</div>
   <div class="f-w-14">L.的快速会议</div>
   <div><span>时间：16：55</span><span style="margin-left: 20px;">发起人：L。</span></div>
  </div>
  <div class="list_right">
    <van-image :src="arrow" width="20px" height="20px"></van-image>
  </div>
</div>
</div>
<div @click="toDetail">
  <div class="data">
  <div>3月30日 周六</div>
  <div class="year">2024年</div>
</div>
<div class="list">
  <div class="list_left">
   <div>251 963 045</div>
   <div class="f-w-14">L.的快速会议</div>
   <div><span>时间：16：55</span><span style="margin-left: 20px;">发起人：L。</span></div>
  </div>
  <div class="list_right">
    <van-image :src="arrow" width="20px" height="20px"></van-image>
  </div>
</div>
</div>
<div @click="toDetail">
  <div class="data">
  <div>3月31日 周日</div>
  <div class="year">2024年</div>
</div>
<div class="list">
  <div class="list_left">
   <div>251 963 045</div>
   <div class="f-w-14">L.的快速会议</div>
   <div><span>时间：16：55</span><span style="margin-left: 20px;">发起人：L。</span></div>
  </div>
  <div class="list_right">
    <van-image :src="arrow" width="20px" height="20px"></van-image>
  </div>
</div>
</div>
</template>
<script setup lang="ts">
import arrow from '../../assets/img/arrow.png'
import { ref } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter()
const value = ref('')
// const list = ref([])
  const onClickLeft = () => history.back();
  const toDetail = ()=>{
    router.push({
      name:'mettingDetail'
    })
  }
</script> 
<style lang="less">
.data{
  display: flex;
  justify-content: space-between;
  padding: 20px;
  .year{
    color: #999a9e;
  }
}
.list{
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding:10px 20px;
  background-color: #fff;
  .list_left{
    font-size: 14px;
    color: #999a9e;
    div{
      margin-bottom: 10px;
    }
    .f-w-14{
      font-size: 16px;
      font-weight: 550;
      color: #000;
    }
  }
}
</style>