
<template>
  <van-nav-bar
      title="任务列表"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
  <van-tabs v-model:active="active" @change="changeList">
    <van-tab title="我的任务">
      <van-list
          v-model="loading"
          :finished="finished"
          :immediate-check="false"
          loading-text="正在加载中请稍后"
          finished-text="没有更多了"
      >
      <div class="list" >
        <div class="list_cardtask" v-for="item in rwList" @click="businessDetailHandle">
          <div class="m-b-10 icon">
            <van-image
                width="20"
                height="20"
                :src="businessIcon"
            />
            <div class="list_title">{{item.taskTitle}}</div>

          </div>
          <van-divider />
          <div class="m-b-10 f-z-12"><span style="color: #4BA3FB" >任务目标：</span>{{item.taskTarget}}</div>
          <div class="m-b-10 f-z-12" style="display: flex;justify-content: space-between">
            <div>
              <span style="color: #4BA3FB" >发布时间：</span>{{item.createTime}}
            </div>
            <van-tag type="primary">待处理</van-tag></div>

<!--          <div class="m-b-10 f-z-12"><span style="font-size: 12px;color: #4BA3FB">发布时间：<span>{{item.time}}</span></span></div>-->
        </div>
      </div>
      </van-list>
      <div class="box"></div>
    </van-tab>
    <van-tab title="已完成">
      <div class="list">
        <div class="list_card" v-for="item in rwList1" @click="businessDetailHandle">
          <div class="flex-space m-b-10">
            <van-image
                width="20"
                height="20"
                :src="businessIcon"
            />
            <div class="list_title">{{item.name}}</div>
            <van-image  width="72"

                        height="60"
                        :src="yiwancheng" style="position:absolute;right:8vw;top:4vw"></van-image>
          </div>
          <van-divider />
          <div class="m-b-10 f-z-12">任务目标：{{item.rwmb}}</div>
          <div class="m-b-10 f-z-12"><span class="f-w">创建时间：</span>{{item.createtime}}</div>

          <div class="m-b-10 f-z-12"><span class="f-w">完成时间：</span>{{item.time}}</div>
<!--          <van-image  width="70"-->
<!--                      v-if="item.isShow"-->
<!--                      height="70"-->
<!--                      :src="daiwancheng" style="position:absolute;right:8vw;top:4vw"></van-image>-->
        </div>
      </div>
    </van-tab>
  </van-tabs>
</template>

<script setup lang="ts">
import {getTaskList, getTaskPoolList} from '../../services/index'
import {inject, onMounted, ref} from 'vue';
import  yiwancheng from '../../assets/img/yiwancheng.png'
import {useRouter} from "vue-router";
import businessIcon from "../../assets/img/business_icon.png";
const active = ref(0);
const router = useRouter()
const loading = ref(true);
const finished = ref(false);
//定义后端返给list数组的数据类型
interface  LISTTYPE{
  bz1: String,
  bz2: String,
  bz3: String,
  corpCode:String,
  corpName:String,
  createTime:String,
  creater:String,
  taskContent:String,
  taskId:String,
  taskImg:String,
  taskStatus:String,
  taskTarget:String,
  taskTitle:String,
  taskType:String
}
const rwList = ref<LISTTYPE[]>([]);
interface paramsTye{
  currentPage:number,
  pageSize:number,
  corpCode:string | null | undefined,
  flag:string,
  userId:string
}
let params = ref<paramsTye>(
    {
      currentPage:1,
      pageSize:10,
      corpCode:'',
      flag:'0',
      userId:''
    }
)
const changeList=(name,title)=>{

}
// //任务数据列表
// const rwList = ref(
//     [{
//       id: '1',
//       name: ' 互联网销售的目标设定、客户分析\n',
//       rwmb:'确定推广的具体目标，如增加网站流量、提高转化率。\n',
//       time: '2024-04-08 15:15:35',
//       createtime:'2024-04-08 15:15:35',
//       isShow:false
//     }, {
//       id: '2',
//       createtime:'2024-04-08 15:15:35',
//       name: '通过社交媒体提升品牌知名度',
//       rwmb:'确定目标受众特征，了解其兴趣、需求和行为习惯。 创作有吸引力、有价值的内容，包括文章、视频、图片等，以满足受众需求。\n',
//       time: '2024-04-07 09:15:35',
//       isShow:false
//     } , {
//       id: '3',
//       createtime:'2024-04-08 15:15:35',
//
//       name: '内容营销长期策略\n',
//       rwmb:'持续发布新内容可以吸引用户回访，保持用户对品牌的关注度。不断推出有价值的内容可以加强品牌在目标受众中的认知和印象。\n' +
//           '\n',
//       time: '2024-03-24 15:18:35',
//       isShow:true
//     }, {
//       id: '4',
//       createtime:'2024-04-08 15:15:35',
//
//       name: '企业收集客户资料，了解客户需求',
//       rwmb:'在网站上设置注册表单，让用户填写基本信息如姓名、电子邮件地址等。设计问卷并在网站、社交媒体或邮件中分享，收集客户反馈和信息。',
//       time: '2024-03-15 08:15:45',
//       isShow:false
//     } ]
// );
// const rwList1 = ref(
//     [
//       {
//         id: '2',
//         createtime:'2024-04-08 15:15:35',
//         name: '通过社交媒体提升品牌知名度',
//         rwmb:'确定目标受众特征，了解其兴趣、需求和行为习惯。 创作有吸引力、有价值的内容，包括文章、视频、图片等，以满足受众需求。\n',
//         time: '2024-04-07 09:15:35',
//         isShow:false
//       } , {
//       id: '3',
//       createtime:'2024-04-08 15:15:35',
//
//       name: '内容营销长期策略\n',
//       rwmb:'持续发布新内容可以吸引用户回访，保持用户对品牌的关注度。不断推出有价值的内容可以加强品牌在目标受众中的认知和印象。\n' +
//           '\n',
//       time: '2024-03-24 15:18:35',
//       isShow:true
//     },
//     ]
// );
// const toDetail=()=>{
//   router.replace('/taskDetail')
// }
const businessDetailHandle = ()=>{
  router.replace('/task')

}
const onClickLeft = () => {
  // history.go(-2)
  router.replace('/homeNew')
}
const getList =()=>{
  getTaskList(params.value).then((res:any)=>{
    console.log('res',res)
    rwList.value.push(...res.data.myTaskList);
  })
}
onMounted(async()=>{
  let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
  params.value.corpCode = await localStorage.getItem('corpCode') //从本地获取corpCode
  params.value.userId = userInfoData.userInfo.userId
  await getList()
})
</script>

<style lang="less" scoped>
.mainButton {
  border-radius: 10px;
  padding: 2vw;
  width:42vw;
  height: 50px;
  background: url('../../assets/img/shenqigndan_bg.png') no-repeat center;
  background-size: 100vw 10vh;
  //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
}
.task{
  //margin-top: 2vw;
  border-radius: 10px;
  //padding: 0 2vw;
  width:100vw;
  height: 160px;
  background: url('../../assets/img/task_bgc_01.png') no-repeat center;
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
    font-size: 0.875em;
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
  background: url('../../assets/img/task_bgc_02.png') no-repeat center;
  background-size: cover;
  //box-sizing: border-box;
  margin-left: 2vw;
  .task_card{
    width: 62vw;
    background-color: #e9f8f1;
    opacity: 0.5; /* 设置透明度为50% */
    color:#000000;
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
.list{
  //height: 100vh;
  overflow: auto;
  .list_card{
    position: relative;
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
      white-space: nowrap; /* 禁止文本换行 */
      overflow: hidden; /* 隐藏超出容器的部分文本 */
      text-overflow: ellipsis; /* 显示省略号 */
      width: 100vw; /* 设置容器宽度 */
      margin-left: 2vw;
    }
    .f-z-12{
      font-size: 0.75em;
      color:#4B5563;
    }
    .m-b-10{
      margin-bottom: 2vw;
    }
    .icon{
      display:flex;
      align-items: center;
    }
    .f-w{
      font-weight: 550;
      color:#000000;
    }
  }
  .list_cardtask{
    position: relative;
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
      white-space: nowrap; /* 禁止文本换行 */
      overflow: hidden; /* 隐藏超出容器的部分文本 */
      text-overflow: ellipsis; /* 显示省略号 */
      width: 100vw; /* 设置容器宽度 */
      margin-left: 2vw;
    }
    .f-z-12{
      font-size: 0.75em;
      color:#4B5563;
    }
    .m-b-10{
      margin-bottom: 2vw;
    }
    .icon{
      display:flex;
      align-items: center;
    }
    .f-w{
      font-weight: 550;
      color:#000000;
    }
  }
}


.box{
  height:50px;
}
:deep(.van-divider){
  margin: 10px;
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
