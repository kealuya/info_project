<template>
  <van-nav-bar
      title="完成任务"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"/>
    <div>
        <van-form>
            <div>
                <van-cell-group class="pb10 mt10">
                    <van-field v-model="reasonForBorrowing"
                            name="reasonForBorrowing"
                            label="任务名称"
                            placeholder="请选择任务名称"
                            class="mt10 mb10"
                            label-width="7.5em"
                            maxlength="20"
                            label-align="top"
                    />
                  <van-field
                      v-model="fieldValue"
                      is-link
                      readonly
                      label="任务类型"
                      placeholder="临床推广"
                      @click="showPicker = true"
                  />
                    <van-notice-bar color="#1989fa" :scrollable="false" wrapable background="#ecf9ff"
                                     class="f10">
                        * 完成以下任务内容：
                        <p>1.根据市场竞争提供科研推荐方案</p>
                        <p>2.制定临床推广策略，包括推广渠道和目标受众</p>
                        <p>3.提供科研会议成果资料，及上传会议内容</p>
                        <p>4.提供成果研究方案及配套针对性论文</p>
<!--                      //上传附件-->
                    </van-notice-bar>
<!--                    <van-field class="mt5" name="uploader" label="上传附件" label-width="6em">-->
<!--                        <template #input>-->
<!--                            <van-uploader v-model="fileList" :after-read="getFile" :before-delete="deleteFile"-->
<!--                                          max-count="4"/>-->
<!--                        </template>-->
<!--                    </van-field>-->
                  <div class="yw_content" @click="taskList">
                    业务内容<van-icon name="arrow" />
                  </div>
               <businessCard :isCollapse="isCollapse"></businessCard>
                </van-cell-group>
            </div>
        </van-form>
    </div>

    <div class="footer_btn">
        <van-button block type="primary" @click="showLoanStatus">
            提交任务
        </van-button>
    </div>
  <van-popup v-model:show="showPicker" round position="bottom">
    <van-picker
        :columns="columns"
        @cancel="showPicker = false"
        @confirm="onConfirm"
    />
  </van-popup>
  <div class="box"></div>
</template>
<script lang="ts" setup>
import businessCard from '../../components/businessCard/index.vue'
import { ref } from 'vue'
import { useRouter } from "vue-router";
// import { showSuccessToast, showFailToast } from 'vant';
import {showSuccessToast} from "vant";
const router = useRouter()
const isCollapse = ref<boolean>(true)
    const columns = [
      { text: '临床推广', value: '临床推广' },
      { text: '科研成果', value: '科研成果' } ,
      { text: '医药销售', value: '医药销售' }
    ];
    const fieldValue = ref('');
    const showPicker = ref(false);

    const onConfirm = ({ selectedOptions }) => {
      showPicker.value = false;
      fieldValue.value = selectedOptions[0].text;
    };

    const onClickLeft = () => {
  // history.go(-2)
  router.replace('/task')
}
const reasonForBorrowing = ref<string>('生物标志物作为高血压患者风险评估的有效指标')
const loanAmount = ref<string>('临床推广')

const taskProcessingHandle = ()=>{
  showSuccessToast('参与成功');

  router.replace('/homeNew')
}
const selectBankAccount = ()=>{
  router.replace('/taskList')
}

function showLoanStatus() {
  router.replace('/taskProcessing')
}
const taskList = ()=>{
  router.replace('/taskList')
}

</script>
<style lang="less" scoped>
.yw_content{
  align-items: center;
  padding:10px;
  display:flex;
  justify-content: space-between;
}
.task{
  //margin-top: 2vw;
  //border-radius: 10px;
  //padding: 0 2vw;
  width:96vw;
  margin: 2vw;
  height: 160px;
  background: url('../../assets/img/home_01.png') no-repeat center;
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
  width:95vw;
  padding:2vw;
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
