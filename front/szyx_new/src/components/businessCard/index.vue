<template>
      <div class="list">
        <div class="list_card" v-for="item in rwList" :key="item.id">
          <div class="flex-space m-b-10">
            <div style="display: flex;align-items: center">
               <van-tag type="primary" v-if="item.hyType == 'mp3'">音频会议</van-tag>
                <van-tag type="success" v-else>文档记录</van-tag>
<!--              <van-image :src="word" width="30" height="30" v-if="item.hyType=='word'"></van-image>-->
<!--              <van-image :src="xmind" width="22" height="22" v-if="item.hyType=='xmind'"></van-image>-->
<!--              <van-image :src="mp3" width="25" height="25" v-if="item.hyType=='mp3'"></van-image>-->
              <div class="list_title">{{item.hyname}}</div>
            </div>
            <van-checkbox v-model="item.isCheck" v-if="!isCollapse"></van-checkbox>
            <!--            <van-image-->
            <!--                width="20"-->
            <!--                height="20"-->
            <!--                :src="businessIcon"-->
            <!--            />-->
          </div>
          <van-divider />
         <div style="display:flex;align-items: center;margin-bottom: 10px; margin-left: 10px;" v-for="i in item.nameList">
           <van-image :src="word" width="23" height="23" v-if="i.type=='word'"></van-image>
           <van-image :src="xmind" width="23" height="23" v-if="i.type=='xmind'"></van-image>
           <van-image :src="mp3" width="23" height="23" v-if="i.type=='mp3'"></van-image>
           <div class="list_title1">{{i.name}}</div>
         </div>
          <div>
<!--            <div class="f-z-12 m-b-10" >所属会议：{{item.sshy}}</div>-->
            <!--<div class="f-z-12 m-b-10" v-if="item.hyType!=='word'">会议地点：{{item.address}}</div>-->
            <div class="f-z-12 m-b-10" style="margin-left: 10px;">创建时间：{{item.hyTime}}</div>

          </div>
        </div>
<!--        <div class="footer_btn">-->
<!--          <van-button block type="primary" @click="handleSubmit">-->
<!--            选择-->
<!--          </van-button>-->
<!--        </div>-->
      </div>
</template>
<script setup lang="ts">
import hyTitle from '../../assets/img/hy_title.png'
import word from '../../assets/img/word.png'
import xmind from '../../assets/img/xmind.png'
import mp3 from '../../assets/img/mp3.png'

import businessIcon from '../../assets/img/business_icon.png';
import {ref,computed} from 'vue';
// import userRouter from 'user'
import { useRouter } from "vue-router";
const checked = ref(false);
const active = ref(0);
const router = useRouter()
const value2 = ref('a');
const timer = ref<string>()
const showDate = ref<boolean>(false)
const props = defineProps({
  isCollapse: {
    type: Boolean,
    default: false
  },
})
const rwList = ref(
    [{
      id: '1',
      hyname:' 数字化聚醚酮临床效果会议',
      nameList:[
        {
          name: '20240408-聚醚数字化临床效果会议文档.word',
          type:'word'
        },
        {
          name: '20240408-聚醚数字化临床效果思维导图.xmind',
          type:'xmind'
        },
        {
          name: '20240408-聚醚数字化临床效果会议录音.mp3',
          type:'mp3'
        }
      ],
      sshy: ' 聚醚酮酮数字化临床效果会议',
      hyTime: '2024-04-07 15:15:35',
      hyType: 'mp3',
      isCheck: false
    },
      {
        id: '4',
          nameList:[
              {
                  name: '20240408-内镜黏膜术和内镜切除术的效果.word',
                  type:'word'
              },
              {
                  name: '20240408-内镜黏膜术和内镜切除术的导图.xmind',
                  type:'xmind'
              }
          ],
        hyname:'内镜黏膜术和内镜切除术的效果比较',
        name: '内镜黏膜术和内镜切除术的效果比较.word',
        time: '4小时07分钟',
        hyType: 'word',
        address: '北京市朝阳区协和医院总院附属医院办公楼C座',
        hyTime: '2024-04-04 17:16:13',
        isGlrw: 'no',
        isCheck: false
      },
      {
        id: '5',
          nameList:[
              {
                  name: '20240408-乳牙早失间隙保持器的临床效果会议.word',
                  type:'word'
              }
          ],
        hyname:'乳牙早失间隙保持器的临床效果会议',
        name: '聚醚酮酮数字化乳牙早失间隙保持器的临床效果.xmind',
        time: '1小时53分钟',
        address: '天津市南开区红旗南路一中心总医院B座',
        hyTime: '2024-04-07 15:15:35',
        isGlrw: 'true',
        hyType: 'xmind',
        isCheck: false
      }
    ]
)
const rwList2 = ref(
    [ {
      id: '3',
      name: '青霉素的抗生素治疗肺炎探讨会议.mp3',
      sshy: ' 青霉素的抗生素治疗肺炎探讨会议',
      hyTime: '2024-03-31 18:14:35',
      hyType: 'mp3',
      isCheck: false
    } ]
)
//业务数据列表
const ywList = ref([{
  id: '1',
  name: '聚醚酮酮数字化乳牙早失间隙保持器的临床效果',
  time: '1小时53分钟',
  address: '天津市南开区红旗南路一中心总医院B座',
  hyTime: '2024-04-07 15:15:35',
  isGlrw: 'true',
  hyType: 'yes',
  isCheck: false
}, {
  id: '2',
  name: '青霉素的抗生素治疗肺炎支原体感染',
  time: '2小时24分钟',
  address: '浙江省杭州市西湖区文三路138号东方通信大厦',
  hyTime: '2024-04-06 13:14:05',
  isGlrw: 'yes',
  isCheck: false
}, {
  id: '3',
  name: '支原体抗生素大环内酯类的使用研究',
  time: '34分钟',
  address: '北京市海淀区北京第三人民医院文三路138号',
  hyTime: '2024-04-06 09:17:15',
  isGlrw: 'fasle',
  isCheck: false
}, {
  id: '4',
  name: '内镜黏膜术和内镜切除术的效果比较',
  time: '4小时07分钟',
  address: '北京市朝阳区协和医院总院附属医院办公楼C座',
  hyTime: '2024-04-04 17:16:13',
  isGlrw: 'no',
  isCheck: false
}  ]);
const option2 = [
  { text: '文档', value: 'a' },
  { text: '文档1', value: 'b' },
  { text: '文档2', value: 'c' },
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
const businessDetailHandle = ()=>{
  router.replace('/businessDetail')
}
const onClickLeft = () => {
  router.replace('/addTasks')
}
const handleSubmit = ()=>{
  router.replace('/addTasks')
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
.footer_btn{
  width:92vw;
  padding:2vw;
  margin: 2vw;
  //position: fixed;
  //bottom:60px;
}
.box{
  height:50px;
}
.list{
  //max-height: 75vh;
  //overflow: auto;
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
      font-weight: 550;
      margin-left: 2vw;
    }
    .list_title1{
      font-size: 0.75em;
      //font-weight: 550;
      margin-left: 2vw;
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
  background: url('../../assets/img/radio.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}

.mainButton_sec {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/radio.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #33B2AD, #3BD7D3);
}
:deep(.van-dropdown-menu__bar){
  box-shadow: none;
}
:deep(.van-divider){
  margin: 2vw;
}
::v-deep(.van-nav-bar__content) {
  background-color: #0088ff;
}

::v-deep(.van-nav-bar__title) {
  color: #ffffff;
}

::v-deep(.van-nav-bar__text) {
  color: #ffffff;
}

::v-deep(.van-icon-arrow-left:before) {
  color: #ffffff;
}
</style>
