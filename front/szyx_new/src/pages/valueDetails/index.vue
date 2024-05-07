<template>
  <van-nav-bar
      title="价值详情"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
    <div class="content">
      <div class="title_one m-b-10" >{{worthTitle}}</div>
      <div class="title_two f-z-14 f-w-550  m-b-10">
        <span>价值概况</span>
<!--        <span class="m-l">发布日期</span>-->
      </div>
      <div class="f-z-14  m-b-10">
        &nbsp;{{worthData}}
      </div>
      <div class="title_two f-z-14 f-w-550  m-b-10">
        <span>申请价值内容</span>
      </div>
      <div class="f-z-14 m-b-10">
        &nbsp;{{worthContent}}
      </div>
    </div>
    <div class="card_small">
      <div class="left">
        <van-image :src="pingfen" width="25" height="25"></van-image>
        <div>价值申请评分</div>
      </div>
      <div class="right"> <div>{{worthScore}}分</div></div>
    </div>
  <div class="card_small">
    <div class="left">
      <van-image :src="jine" width="25" height="25"></van-image>
      <div>价值申请金额</div>

    </div>
    <div class="right"> <div>{{money}}元</div></div>
  </div>
  <div style="wdith:96vw;margin:2vw;text-align: center;background-color:#fff;">
    <div><van-image :src="erweima" width="230" height="200"></van-image></div>
  </div>

    <div class="footer_btn">
<!--      <van-button block type="primary" @click="taskProcessingHandle">-->
<!--        立即参与-->
<!--      </van-button>-->
    </div>
  <div class="box"></div>
</template>
<script lang="ts" setup>
import erweima from '../../assets/img/erweima.jpg'
import pingfen from '../../assets/img/pf001.png'
import jine from '../../assets/img/jz001.png'
import {getWorthDetails} from '../../services/task-processing/index'
import {useRouter,useRoute} from "vue-router";
// import { showSuccessToast, showFailToast } from 'vant';
import {showSuccessToast} from "vant";
import {inject, onMounted, ref} from "vue";
const router = useRouter()
const route = useRoute()
let worthScore = ref<string>()
let money= ref<string>()
let worthData = ref<string>()
let worthContent = ref<string>()
let worthTitle = ref<string>()
interface paramsType{
  worthId: string, //价值ID，对应任务的ID
  corpName: string | undefined | null, //企业名称
  corpCode: string, //企业code
  userId: string, //用户ID
  userName: string, //用户姓名
  userMobile: string //用户手机号
}
let params =ref<paramsType>(
    {
      worthId: '', //价值ID，对应任务的ID
      corpName: "", //企业名称
      corpCode: '', //企业code
      userId: "", //用户ID
      userName: "", //用户姓名
      userMobile: "" //用户手机号
    }
)
const onClickLeft = () => {
  // history.go(-2)
  router.replace('/applicationValue')
}
const taskProcessingHandle = ()=>{
  showSuccessToast('参与成功');

  router.replace('/homeNew')
}
onMounted(()=>{
  params.value.worthId = route.query.id
  let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
  params.value.corpCode = localStorage.getItem('corpCode') //从本地获取corpCode
  //从本地获取corpCode
  params.value.userId = userInfoData.userInfo.userId
  params.value.userName = userInfoData.userInfo.userName
  params.value.userMobile = userInfoData.userInfo.userMobile
  params.value.corpName = userInfoData.userInfo.corpName
  console.log('params',params.value)
  getWorthDetails(params.value).then((res:any)=>{
  if(res.success){
    worthScore.value = res.data.worthScore
    money.value = res.data.money
    worthTitle.value = res.data.worthTitle
    worthData.value = res.data.worthData
    worthContent.value = res.data.worthContent
  }
  })
})
</script>
<style lang="less" scoped>
.card_small{
  width: 92vw;
  margin: 2vw;
  padding:2vw;
  background-color: #fff;
  border-radius: 6px;
  display:flex;
  justify-content: space-between;
  font-size:0.875em;
  font-weight: 550;
.left{
  display:flex;
  align-items: center;
  div{
    margin-left: 2vw;
  }
}
}
.task{
  //margin-top: 2vw;
  //border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  margin: 2vw;
  height: 160px;
  background: url('../../assets/img/home_05.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  //margin-left: 2vw;
  .task_card{
    width: 62vw;
    background-color: #e9f8f1;
    opacity: 0.5; /* 设置透明度为50% */
    color:#000000;
    margin: 2vw;
    padding:2vw;
    font-size: 1em;
    border-radius: 10px;
    .content-font{
      font-size: 0.75em;
      margin-top:1vw;
    }
  }
  .btn{
    border-radius: 10px;
    background-color: #253334;
    width:42vw;
    margin-left: 2vw;
    padding:1vw;
    text-align: center;
    font-size: 1em;
  }
}
.box{
  height:50px;
}
.content{
  margin:2vw;
  line-height:3vh;
.title_one{
  text-align:center;
  font-weight: 550;
  font-size: 0.875em;
}
 .f-z-14{
   font-size: 0.75em;
 }
  //.f-z-12{
  //  font-size: 0.75em;
  //}
  .f-w-550{
    font-weight: 550;
  }
  .m-l{
    margin-left: 20vw;
  }
  .m-b-10{
    margin-bottom:10px;
  }
}
.footer_btn{
  width:92vw;
  position: fixed;
  bottom:60px;
  padding:2vw;
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
<script setup lang="ts">
</script>
