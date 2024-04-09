
<template>
    <van-image :src="bgsy" width="100%" height="180"></van-image>
  <van-row  justify="space-around" :getter="10" align="center" class="mt10">
    <van-col span="11">
      <div class="mainButton_zsk f_white" @click="hangyezhishiku">
        <van-row justify="space-between" align="center" >
          <van-col>
            <p class="f16">行业知识库</p>
            <p class="f12">资源共享 轻松查找共享信息</p>
          </van-col>
<!--          <van-col>-->
<!--            <van-icon :name="shenqingdanIcon" size="3em"   />-->
<!--          </van-col>-->
        </van-row>
      </div>
    </van-col>
    <van-col span="11">
      <div class="mainButton f_white">
        <van-row justify="space-between" align="center" @click="taskProcessingHandle">
          <van-col>
            <p class="f16">任务列表</p>
            <p class="f12">最新发布 参与流程</p>
          </van-col>
<!--          <van-col>-->
<!--            <van-icon :name="shenqingdanIcon" size="3em"   />-->
<!--          </van-col>-->
        </van-row>
      </div>
      <div class="mainButton_sec  f_white">
        <div>
          <van-row justify="space-between" align="center" @click="applicationHandle">
            <van-col>
              <p class="f16">价值申请</p>
              <p class="f12">一键申请  随时随地 </p>
            </van-col>
<!--            <van-col>-->
<!--              <van-icon :name="danjuIcon" size="3em"  />-->
<!--            </van-col>-->
          </van-row>
        </div>
      </div>
    </van-col>
  </van-row>
  <div class="mt5">
    <van-row class="mt10 card_padding" align="center">
      <span class="title">任务发布</span> <span style="color: #bbc2cc; font-size: 12px;margin: 5px 10px 0">更多海量任务、科研成果共享完成</span>
    </van-row>
    <div>
    </div>
  </div>

  <!--      任务二-->
  <van-row>
    <div class="task2  f_white">
      <div class="task_card">
        <div class="task_title">
          了解特定药品在治疗某些疾病时的疗效和患者反馈。
        </div>
        <div class="content-font">
          与医生、患者或相关人员进行访谈,收集某种药品在临床使用中的疗效数据、不良反应情况、用药体验等。可以录音或记录文字资料。
        </div>
      </div>
      <div class="btn"  @click="toDetail">
        <van-icon name="play-circle-o" /> 任务详情
      </div>
    </div>
  </van-row>
<!--  <van-row>-->
<!--    <div class="task f_white">-->
<!--      <div class="task_card">-->
<!--        <div class="task_title">-->
<!--          高血压患者心血管事件风险评估-->
<!--        </div>-->
<!--        <div class="content-font">-->
<!--          纳入1000名已诊断高血压的患者，收集其血清样本，并运用生物化学方法进行分析。随访3年，观察心血管事件发生情况-->
<!--        </div>-->
<!--      </div>-->
<!--&lt;!&ndash;      <van-button icon="play-circle-o" type="primary" size="small">立即参与</van-button>&ndash;&gt;-->
<!--      <div class="btn"  @click="toDetail">-->
<!--        <van-icon name="play-circle-o" /> 立即参与-->
<!--      </div>-->
<!--    </div>-->
<!--  </van-row>-->
  <van-row>
    <div class="task3  f_white">
      <div class="task_card">
        <div class="task_title">
          掌握某些常见疾病的诊疗方案和用药指南。
        </div>
        <div class="content-font">
          与医院专家访谈,或查阅相关文献资料,了解针对某类疾病的诊断流程、治疗方案、推荐用药等,形成诊疗方案概要。
        </div>
      </div>
      <div class="btn" @click="toDetail2">
        <van-icon name="play-circle-o" /> 任务详情
      </div>
    </div>
  </van-row>
  <van-row>
    <div class="task4  f_white" >
      <div class="task_card">
        <div class="task_title">
          及时获取同行业新上市药品的信息。
        </div>
        <div class="content-font">
          关注行业内新药上市动态,及时采集新药的适应症、功能主治、使用说明、销售情况等资料。
        </div>
      </div>
      <div class="btn" @click="toDetail3">
        <van-icon name="play-circle-o" /> 任务详情
      </div>
    </div>
  </van-row>
  <van-row>
    <div class="task5  f_white">
      <div class="task_card">
        <div class="task_title">
          了解竞品在临床和市场的使用状况。
        </div>
        <div class="content-font">
          通过走访医院、药房等渠道,调查竞品产品的医生开具率、患者使用反馈、市场占有率等情况,并收集相关资料
        </div>
      </div>
      <div class="btn" @click="toDetail4">
        <van-icon name="play-circle-o" /> 任务详情
      </div>
    </div>
  </van-row>
  <div class="box"></div>
</template>

<script setup lang="ts">
  import bgsy from '../../assets/img/bgsy.png';
import { showSuccessToast, showFailToast, showToast } from 'vant';
import shenqingdanIcon from  '../../assets/img/shenqingdan.png';
import danjuIcon from  '../../assets/img/danju.png';
import tabbar from '../../Tabbar.vue'
import {useRouter} from "vue-router";
const router = useRouter()
const toDetail=()=>{
  router.replace('/taskDetail')
}
const toDetail2=()=>{
  router.replace('/taskDetail_two')
}
const toDetail3=()=>{
  router.replace('/taskDetail_three')
}
const toDetail4=()=>{
  router.replace('/taskDetail_four')
}
const taskProcessingHandle = ()=>{
  router.replace('/taskProcessing')

}
const hangyezhishiku = ()=>{
  showToast('暂未开放，敬请期待！');
}
const applicationHandle = ()=>{
  router.replace('/applicationValue')
}
</script>

<style lang="less" scoped>
.title_banner{
  font-size: 0.875em;
}
.headBg {
  width: 100vw;
  height: 22vh;
  background-color: #0080FF;
}
.card_padding{
  padding:2vw;
}
.title{
  color:#000000;
  font-size: 16px;
}
.mainButton {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/shenqingdan_bg.png') no-repeat center;
  background-size: 48vw 10vh;
  margin-bottom: 10px;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}

.mainButton_sec {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/tidanju_bg.png') no-repeat center;
  background-size: 48vw 10vh;
  //background-image: linear-gradient(to right bottom, #33B2AD, #3BD7D3);
}
.mainButton_zsk{
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 125px;
  background: url('../../assets/img/zhishiku_bg.png') no-repeat center;
  background-size: 48vw 20vh;
}
.box{
  height:80px;
}
.task_title{
  white-space: nowrap; /* 禁止文本换行 */
  overflow: hidden; /* 隐藏超出容器的部分文本 */
  text-overflow: ellipsis; /* 显示省略号 */
  width: 55vw; /* 设置容器宽度 */
  font-size: 1.1em;
  font-weight: 380;
}
.task{
  margin-top: 2vw;
  border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  height: 160px;
  background: url('../../assets/img/home_01.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  margin-left: 2vw;
  .task_card{
    width: 62vw;
    background-color: #000000;
    opacity: 0.5; /* 设置透明度为50% */
    color:#ffffff;
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
    font-size: 0.875em;
  }
}
.task2{
  margin-top: 4vw;
  border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  height: 160px;
  background: url('../../assets/img/task_bgc_01.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  margin-left: 2vw;
  .task_card{
    width: 58vw;
    background-color: #000000;
    opacity: 0.5; /* 设置透明度为50% */
    color:#ffffff;
    margin:2vw;
    padding:2vw;
    font-size: 0.875em;
    border-radius: 10px;
    .content-font{
      font-size: 0.75em;
      margin-top:1vw;
    }
  }
  .btn{
    border-radius: 10px;
    background-color: #1768ed;
    width:42vw;
    margin-left: 2vw;
    padding:1vw;
    text-align: center;
    font-size: 0.875em;
  }
}
.task3{
  margin-top: 4vw;
  border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  height: 160px;
  background: url('../../assets/img/home_01.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  margin-left: 2vw;
  .task_card{
    width: 62vw;
    background-color: #000000;
    opacity: 0.5; /* 设置透明度为50% */
    color:#ffffff;
    margin: 2vw;
    padding:2vw;
    font-size: 0.875em;
    border-radius: 10px;
    .content-font{
      font-size: 0.75em;
      margin-top:1vw;
    }
  }
  .btn{
    border-radius: 10px;
    background-color: #1768ed;
    width:42vw;
    margin-left: 2vw;
    padding:1vw;
    text-align: center;
    font-size: 0.875em;
  }
}
.task4{
  margin-top: 4vw;
  border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  height: 160px;
  background: url('../../assets/img/home_04.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  margin-left: 2vw;
  .task_card{
    width: 62vw;
    background-color: #000000;
    opacity: 0.5; /* 设置透明度为50% */
    color:#ffffff;
    margin: 2vw;
    padding:2vw;
    font-size: 0.875em;
    border-radius: 10px;
    .content-font{
      font-size: 0.75em;
      margin-top:1vw;
    }
  }
  .btn{
    border-radius: 10px;
    background-color: #1768ed;
    width:42vw;
    margin-left: 2vw;
    padding:1vw;
    text-align: center;
    font-size: 0.875em;
  }
}
.task5{
  margin-top: 4vw;
  border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  height: 160px;
  background: url('../../assets/img/home_05.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  margin-left: 2vw;
  .task_card{
    width: 62vw;
    background-color: #000000;
    opacity: 0.5; /* 设置透明度为50% */
    color:#ffffff;
    margin: 2vw;
    padding:2vw;
    font-size: 0.875em;
    border-radius: 10px;
    .content-font{
      font-size: 0.75em;
      margin-top:1vw;
    }
  }
  .btn{
    border-radius: 10px;
    background-color: #1768ed;
    width:42vw;
    margin-left: 2vw;
    padding:1vw;
    text-align: center;
    font-size: 0.875em;
  }
}

</style>
