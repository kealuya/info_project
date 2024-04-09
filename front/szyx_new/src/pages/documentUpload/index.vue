<template>
  <van-nav-bar
      title="文档上传"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
  />
  <div class="header">

    <div class="f-z-14-c">会议ID</div>
    <div class="f-z-16">WD2024040812552563</div>
    <div style="display:flex;justify-content: space-between;align-items: center">
     <div>
       <div class="f-z-14-c m-t-10">会议标题</div>
       <div class="f-z-16"> <van-field v-model="reasonForBorrowing" name="reasonForBorrowing"  placeholder="20240408会议1" /></div>
     </div>

    </div>
  </div>
  <div class="header">
    <div class="header_title">
      <van-image :src="huiyijilu" width="25" height="25"></van-image>
      <span>文档上传</span>
    </div>

    <div style="display:flex;justify-content: space-between;align-items: center">
      <div>
        <div class="f-z-14-c m-t-10">会议标题</div>
        <div class="f-z-16"> <van-field   placeholder="请输入文档名称" /></div>
        <div class="f-z-14-c m-t-10">图片文件</div>
        <div class="f-z-16"> <van-uploader :after-read="afterRead" /></div>
        <div class="f-z-14-c m-t-10">附件上传</div>
        <div class="f-z-16">
          <van-uploader>
          <van-button icon="plus"  type="primary" size="small">上传文件</van-button>
        </van-uploader>
        </div>

      </div>

    </div>
  </div>


  <div class="cardTable" @click="handelMetting">
    <van-button type="primary" block style="width: 92%; margin: 20px"
    >会议结束</van-button
    >
  </div>
  <van-dialog
      v-model:show="show"
      :show-cancel-button="false"
      :showConfirmButton="false"
  >
    <van-form @submit="onSubmit" style="padding: 10px" ref="formRef">
      <van-cell-group inset>
        <van-field
            v-model="username"
            name="会议标题"
            label="会议标题"
            placeholder="会议标题"
            :rules="[{ required: true, message: '请填写会议标题' }]"
        />
      </van-cell-group>
      <div style="margin: 16px;display:flex">
        <van-button block style="margin-right: 10px" @click="cancel"  size="small">
          取消
        </van-button>
        <van-button block type="primary" native-type="submit"  size="small">
          提交
        </van-button>
      </div>
    </van-form>
  </van-dialog>
</template>

<script setup lang="ts">
import huiyijilu from '../../assets/img/huiyijilu.png'
import { showConfirmDialog } from 'vant';
import { ref, inject,onMounted } from "vue";
import { useRouter } from "vue-router";
import zanwupiaoju from "../../assets/img/zanwupiaoju.png";
const value = ref<string>('')
const isShow = ref<Boolean>(false);
const isHide = ref<Boolean>(true);
const show = ref(false);
const username = ref("");
const indexValue = ref()
const formRef = ref<any>(null); // 定义表单引用
interface ListItem {
  time: string;
  date: string;
  long: string;
}

const list = ref<Array<ListItem>>([]);
const deleteItem = (index:any)=>{
  indexValue.value = index
  showConfirmDialog({
    title: '提示',
    message:
        '你确定要删除吗？',
  })
      .then(() => {
        list.value.splice(indexValue.value,1)
      })
      .catch(() => {
        // on cancel
      });
}
const reasonForBorrowing = ref<string>('2024-04-10文档会议记录1')
const onSubmit = (values:any) => {
  const valid = formRef.value.validate();
  if (valid) {
    console.log('通过了')
    // 如果验证通过，则执行页面跳转
    // 请确保您已经设置好了Vue Router
    router.replace('/business');
  }
};
const cancel = ()=>{
  show.value = false
}

const active = ref(0);
let userInfoData: any = inject("userInfo");
const router = useRouter();
let rec: any, wave: any, recBlob: any;
const onClickLeft = () => {
  router.push('/business')
}
//点击跳转到会议列表页面
const handelMetting = () => {
  show.value = true;
  // console.log("点击了");
  // router.push({
  //   name: "metting",
  // });
};
</script>

<style lang="less" scoped>
.cardTable {
  width: 100%;
  // margin-top: 20px;
  text-align: center;
  position: fixed;
  bottom: 60px;
}
.uploader{
  //text-align:center;
  margin-left: 10px;
}
.cardTable {
  width: 100%;
  // margin-top: 20px;
  text-align: center;
  position: fixed;
  bottom: 60px;
}

.touxiangimg {
  width: 20px;
  height: 20px;
}

.container {
  position: relative;
  width: 100%;
  min-height: 750px;
}

.bottom_box {
  min-width: 100%;
  position: absolute;
  bottom: 0;
  text-align: center;
}
.header {
  padding:2vw 4vw;
  margin:10px;
  border-radius: 10px;
  background-color: #fff;
  // width: 96%;
  //display: flex;
  //justify-content: space-between;
  //align-items: center;
  //padding: 10px;
  .left {
    display: flex;
    .information {
      margin-left: 14px;
      font-size: 14px;
      .amount-12 {
        font-size: 12px;
        // color: #ccc;
        margin-top: 6px;
      }
    }
  }
  .header_title{
    display: flex;
    align-items: center;
    font-size: 1em;
    margin-bottom:2vw;
    span{
      margin-left: 10px;
    }
  }
}
.f-w-550 {
  font-weight: 550;
}
.f-z-14-c{
  color:rgba(23,26,29,0.40);
  font-size: 0.875em;
}
.m-t-10{
  margin-top: 2vw;
}
.f-z-16{
  margin-top: 5px;
  font-size: 16px;
}
.metting {
  display: flex;
  justify-content: space-around;
  margin: 10px 0;
  text-align: center;
  font-size: 12px;
  // justify-content: center;
}
.round-button {
  display: flex;
  justify-content: space-around;
  margin: 10px;
  text-align: center;
  font-size: 12px;
  position: fixed;
  bottom: 130px;
  width: 94%;
}
.footer {
  width: 100%;
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
  // color: #ccc;
  .right {
    display: flex;
    text-align: center;
  }
  .time {
    font-size: 14px;
    font-weight: 600;
    margin-bottom: 10px;
    color: #000;
  }
}
.scrollable-div {
  max-height: 28vh; /* 设置最大高度为300px */
  overflow-y: auto; /* 当内容超出高度时显示垂直滚动条 */
}
///* 小屏幕设备 */
//@media screen and (max-width: 100px) {
//  .scrollable-div {
//    height: 28vh; /* 设置列表高度为屏幕高度的60% */
//    overflow-y: auto;
//  }
//}
//
///* 大屏幕设备 */
//@media screen and (min-width: 101px) {
//  .scrollable-div {
//    height: 30vh; /* 设置列表高度为屏幕高度的80% */
//    overflow-y: auto;
//  }
//}
.empty{
  font-size: 12px;
  color: #999a9e;
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
::v-deep(.van-cell){
  padding:0;
}

</style>
