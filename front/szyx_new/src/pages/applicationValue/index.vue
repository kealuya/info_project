<template>
  <van-nav-bar
      title="价值申请"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
  <van-tabs v-model:active="active">
    <van-tab title="待申请列表">
      <!--      <div class="box"></div>-->
      <!--  !&#45;&#45;   //列表部分&ndash;&gt;-->
      <div class="list">
        <div class="list_card" v-for="item in jzList" :key="item.id">
          <div class="flex-space m-b-10">
            <div style="display:flex">
              <van-image :src="Money" width="20" height="20"></van-image>
              <div class="list_title">{{item.name}}</div>
            </div>
            <van-checkbox v-model="item.isCheck"></van-checkbox>
            <!--            <van-image-->
            <!--                width="20"-->
            <!--                height="20"-->
            <!--                :src="businessIcon"-->
            <!--            />-->
          </div>
          <van-divider />
          <div>
            <div class="f-z-12 m-b-10" >任务评分：{{item.number}}</div>
            <div class="flex-space f-z-12 m-b-10">
              <div>预计可申请金额：{{item.sqje}}</div>
            </div>
            <div class="flex-space f-z-12">
<!--              <div></div>-->
              <div>{{item.time}}</div>
<!--              <div>已完成</div>-->

            </div>
          </div>
        </div>
        <div class="footer_btn">
          <van-button block type="primary" @click="handleSubmit">
            提交申请
          </van-button>
        </div>
      </div>
      <!--      <van-row  justify="space-around" :getter="20" align="center" class="mt10">-->
      <!--        <van-col span="11">-->
      <!--          <div class="mainButton f_white">-->
      <!--            <div>-->
      <!--              <van-row justify="space-between" align="center">-->
      <!--                <van-col>-->
      <!--                  <p class="f16">音频记录</p>-->
      <!--                  &lt;!&ndash;                <p class="f12">快速申请 规范流程</p>&ndash;&gt;-->
      <!--                </van-col>-->
      <!--                <van-col>-->
      <!--                  &lt;!&ndash;              <van-icon :name="shenqingdanIcon" size="3em"   />&ndash;&gt;-->
      <!--                </van-col>-->
      <!--              </van-row>-->
      <!--            </div>-->
      <!--          </div>-->
      <!--        </van-col>-->
      <!--        <van-col span="11">-->
      <!--          <div class="mainButton_sec  f_white">-->
      <!--            <div>-->
      <!--              <van-row justify="space-between" align="center">-->
      <!--                <van-col>-->
      <!--                  <p class="f16">文档上传</p>-->
      <!--                  &lt;!&ndash;                <p class="f12">移动报销 随时随地</p>&ndash;&gt;-->
      <!--                </van-col>-->
      <!--                <van-col>-->
      <!--                  &lt;!&ndash;              <van-icon :name="danjuIcon" size="3em"  />&ndash;&gt;-->
      <!--                </van-col>-->
      <!--              </van-row>-->
      <!--            </div>-->
      <!--          </div>-->
      <!--        </van-col>-->
      <!--      </van-row>-->
    </van-tab>
    <van-tab title="申请记录">
      <!--      <div class="box"></div>-->
      <!--  !&#45;&#45;   //列表部分&ndash;&gt;-->
      <div class="list">
        <div class="list_card" v-for="item in jzList1" @click="handleJump">
          <div style="display:flex">
            <van-image :src="Money" width="20" height="20"></van-image>
            <div class="list_title">{{item.name}}</div>
          </div>
<!--            <van-image-->
<!--                width="20"-->
<!--                height="20"-->
<!--                :src="businessIcon"-->
<!--            />-->
<!--          </div>-->
          <van-divider />
          <div class="f-z-12 m-b-10" >任务评分:{{item.number}}</div>
          <div class="f-z-12 m-b-10">预计可申请金额:{{item.sqje}}</div>
          <div class="flex-space f-z-12">
<!--            <div>已关联任务100002</div>-->
            <div>2024.01.23</div>
          </div>
          <div class="list_icon" v-if="item.isShow"><van-image :src="waitshenpi" width="60" height="60"></van-image></div>
          <div class="list_icon" v-else><van-image :src="yitongyi" width="60" height="60"></van-image></div>
        </div>

      </div>
      <!--      <van-row  justify="space-around" :getter="20" align="center" class="mt10">-->
      <!--        <van-col span="11">-->
      <!--          <div class="mainButton f_white">-->
      <!--            <div>-->
      <!--              <van-row justify="space-between" align="center">-->
      <!--                <van-col>-->
      <!--                  <p class="f16">音频记录</p>-->
      <!--                  &lt;!&ndash;                <p class="f12">快速申请 规范流程</p>&ndash;&gt;-->
      <!--                </van-col>-->
      <!--                <van-col>-->
      <!--                  &lt;!&ndash;              <van-icon :name="shenqingdanIcon" size="3em"   />&ndash;&gt;-->
      <!--                </van-col>-->
      <!--              </van-row>-->
      <!--            </div>-->
      <!--          </div>-->
      <!--        </van-col>-->
      <!--        <van-col span="11">-->
      <!--          <div class="mainButton_sec  f_white">-->
      <!--            <div>-->
      <!--              <van-row justify="space-between" align="center">-->
      <!--                <van-col>-->
      <!--                  <p class="f16">文档上传</p>-->
      <!--                  &lt;!&ndash;                <p class="f12">移动报销 随时随地</p>&ndash;&gt;-->
      <!--                </van-col>-->
      <!--                <van-col>-->
      <!--                  &lt;!&ndash;              <van-icon :name="danjuIcon" size="3em"  />&ndash;&gt;-->
      <!--                </van-col>-->
      <!--              </van-row>-->
      <!--            </div>-->
      <!--          </div>-->
      <!--        </van-col>-->
      <!--      </van-row>-->
    </van-tab>
  </van-tabs>

  <!--  <div class="mb10"></div>-->
</template>
<script setup lang="ts">
import waitshenpi from '../../assets/img/waitshenpi_icon.png'
import yitongyi from '../../assets/img/yitongyi_icon.png'
import Money from '../../assets/img/money.png'
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
//业务数据列表
//价值申请数据列表
const handleJump = ()=>{
  router.replace('/valueDetails')
}
const jzList = ref([{
  id: '1',
  name: '肠道微生物群在抑郁症发病机制中的作用研究',
  number: '67分',
  sqje: '12,000',
  time: '2024-04-07 15:15:35',
  state: '待申请',
  isCheck:false
}, {
  id: '2',
  name: '利用机器学习算法改进手术风险评估模型的研究',
  number: '78.5分',
  sqje: '32,000',
  time: '2024-04-06 15:19:45',
  state: '待申请',
  isCheck:false

}, {
  id: '3',
  name: '利用CRISPR等技术进行基因编辑',
  number: '58.5分',
  sqje: '10,500',
  time: '2024-04-05 14:15:35',
  state: '已完成',
  isCheck:false

} ])
const jzList1 = ref([{
  id: '1',
  name: '抑郁症发病机制中和药物是否产生病理反应',
  number: '87分',
  sqje: '32,000',
  time: '2024-04-05 15:15:35',
  state: '待申请',
  isCheck:false,
  isShow:false
}, {
  id: '2',
  name: '心脏手术风险评估模型与临床实验的结合',
  number: '78.5分',
  sqje: '32,000',
  time: '2024-04-06 15:19:45',
  state: '待申请',
  isCheck:false,
  isShow: true

} ])
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
  router.replace('/homeNew')
}
const handleSubmit = ()=>{
  active.value = 1
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
  position: fixed;
  bottom:50px;
}
.box{
  height:50px;
}
.list{
  //max-height: 75vh;
  //overflow: auto;
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
      margin-left: 2vw;
    }
    .f-z-12{
      font-size: 0.75em;
      color:#4B5563;
    }
    .m-b-10{
      margin-bottom: 2vw;
    }
    .list_icon{
      position: absolute;
      right:30px;
      top:40px;
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
