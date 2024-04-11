<template>
    <!-- 上半部分的header -->
    <div style="background:#0080FF;height: 30%;padding-bottom:13px">
        <van-row class="pt10 pl20">
            <van-col span="23" style="text-align: center">
                <span class="f_white" style="font-weight: bold">业务列表</span>
            </van-col>
        </van-row>
    </div>
  <div class="header">
   <van-dropdown-menu class="menuItems">
     <!--    <van-dropdown-item v-model="value1" :options="option1" />-->
     <van-dropdown-item v-model="value2" :options="option2" />
   </van-dropdown-menu>
   <div  @click="showDate = true" class="menuItems">{{timer?timer:'时间范围'}}</div>
   <van-calendar v-model:show="showDate" @confirm="onConfirm" type="range" :min-date="minDate" :max-date="maxDate"/>
 </div>
  <div class="box"></div>
<!--  !&#45;&#45;   //列表部分&ndash;&gt;-->
  <div class="list">
    <div class="list_card" v-for="item in ywList" @click="businessDetailHandle(item.id)">
      <div class="flex-space m-b-10">
        <div class="list_title">
            <van-tag type="primary" v-if="item.hyType">音频会议</van-tag>
            <van-tag type="success" v-else>文档记录</van-tag>
            {{item.name}}
        </div>
      </div>
      <van-divider />
      <div class="f-z-12 m-b-10" >会议时长：{{item.time}}</div>
      <div class="f-z-12 m-b-10">会议地址：{{item.address}}</div>
      <div class="f-z-12 m-b-10">结束时间：{{item.hyTime}}</div>
       <div class="f-z-12 m-b-10">
           <span v-if="item.hyType"  style="color: #0080FF">已关联任务：RW202404080001</span>
       </div>

    </div>
  </div>
  <van-row  justify="space-around" :getter="20" align="center" class="mt15 radio">
    <van-col span="11">
      <div class="mainButton f_white" @click="speechHandle">
        <div>
          <van-row justify="space-between" align="center">
            <van-col>
              <p class="f16">音频记录</p>
              <p class="f12">会议录制 准确还原</p>
            </van-col>
          </van-row>
        </div>
      </div>
    </van-col>
    <van-col span="11">
      <div class="mainButton_sec  f_white" @click="documentUpload">
        <div>
          <van-row justify="space-between" align="center">
            <van-col>
              <p class="f16">文档上传</p>
              <p class="f12">集中存储 轻松共享</p>
            </van-col>
          </van-row>
        </div>
      </div>
    </van-col>
  </van-row>
<!--  <div class="mb10"></div>-->
</template>
<script setup lang="ts">
import businessIcon from '../../assets/img/business_icon.png';
import {ref,computed} from 'vue';
// import userRouter from 'user'
import { useRouter } from "vue-router";
const router = useRouter()
const value2 = ref('a');
const timer = ref<string>()
const showDate = ref<boolean>(false)
//业务数据列表
    const ywList = ref([{
      id: '1',
      name: '聚醚酮酮数字化临床效果会议',
      time: '1小时53分钟',
      address: '天津市南开区红旗南路一中心总医院B座',
      hyTime: '2024-04-07 15:15:35',
      isGlrw: true,
      hyType: true,
      isCheck: 'fasle'
    }, {
      id: '2',
      name: '青霉素的抗生素治疗肺炎探讨会议',
      time: '2小时24分钟',
      address: '浙江省杭州市西湖区文三路138号东方通信大厦',
      hyTime: '2024-04-06 13:14:05',
      isGlrw: false,
        hyType: true,
      isCheck: 'fasle'
    }, {
      id: '3',
      name: '支原体抗生素大环内酯类的研究会议',
      time: '34分钟',
      address: '北京市海淀区北京第三人民医院文三路138号',
      hyTime: '2024-04-06 09:17:15',
      isGlrw: false,
        hyType: false,
      isCheck: 'fasle'
    }/*, {
      id: '4',
      name: '内镜黏膜术和内镜切除术比较会议',
      time: '4小时07分钟',
      address: '北京市朝阳区协和医院总院附属医院办公楼C座',
      hyTime: '2024-04-04 17:16:13',
      isGlrw: false,
        hyType: false,
      isCheck: 'fasle'
    } */ ])
const option2 = [
  { text: '会议类型', value: 'a' },
  { text: '音频记录', value: 'b' },
  { text: '文档记录', value: 'c' },
];
const thisYear = new Date().getFullYear();

const minDate = computed(() => {
  return new Date(thisYear, 0); // 今年一月
});

const maxDate = computed(() => {
  return new Date(thisYear + 1, 0); // 明年一月
});
const formatDate = (date: any) => `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`;  //格式化日期 （年-月-日）
const onConfirm = async(values: any) => {
  const [start, end] = values;
  // console.log(values,'zheli')
  showDate.value = false;
  let beginTime = formatDate(start);
  let endTime = formatDate(end);
  console.log(beginTime,endTime)
  timer.value = `${beginTime}~${endTime}`;

};
const businessDetailHandle = (id:number)=>{
  console.log('number',id)
  if(id ==3){
    router.replace('/documentationDetail')
  }else{
    router.replace('/businessDetail')
  }
}
const speechHandle = ()=>{
  router.replace('/speech')

}
const documentUpload=()=>{
  router.replace('/documentUpload')
}
</script>
<style lang="less" scoped>
.header {
  position: fixed;
  z-index: 99;
  width: 100%;
  height: 60px;
  background-color: #fff;
  // padding: 0 10px;
  font-size: 14px;
  display: flex;
  text-align: center;
  height: 50px;
  line-height: 60px;
  border-bottom: 1px solid #f0f0f0;
  //margin-bottom: 50px;
  .menuItems {
    font-size: 14px;
    width: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    // border: 1px solid red;
  }
}
.box{
  height:50px;
}
.list{
  max-height: 75vh;
  overflow: auto;
  .list_card{
    width:92vw;
    margin:2vw;
    padding:2vw;
    background-color: #FFFFFF;
    border-radius: 10px;
    .flex-space{
      display: flex;
      justify-content: space-between;
    }
    .list_title{
      font-size: 0.875em;
      font-weight: bold;
    }
    .f-z-12{
      font-size: 0.75em;
      color:#4B5563;
    }
    .m-b-10{
      margin-bottom: 2vw;
    }
  }
}
.mainButton {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/yinpin_bg.png') no-repeat center;
  background-size: 48vw 10vh;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}
.radio{
    position: fixed;
    bottom:10vh;
}
.mainButton_sec {
  margin-left: 2vw;
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/wendangshangchuan_bg.png') no-repeat center;
  background-size: 48vw 10vh;
  //background-image: linear-gradient(to right bottom, #33B2AD, #3BD7D3);
}
:deep(.van-dropdown-menu__bar){
  box-shadow: none;
}
:deep(.van-divider){
  margin: 2vw;
}
</style>
