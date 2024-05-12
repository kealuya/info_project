<template>
    <van-nav-bar
            title="文档上传"
            left-text="返回"
            left-arrow
            @click-left="onClickLeft"
    />
    <div class="header">
        <div class="header_title">
            <van-image :src="huiyijilu" width="25" height="25"></van-image>
            <span>文档上传</span>
        </div>
        <div class="f-z-14-c">会议ID</div>
        <div class="f-z-16">{{ applyId }}</div>
        <!--    <div style="display:flex;justify-content: space-between;align-items: center">-->
        <div>
            <div class="item">
                <div class="f-z-14-c m-t-10"><span style="color: red">*</span>会议标题</div>
                <van-field v-model="reasonForBorrowing" name="reasonForBorrowing"  clearable placeholder="请输入会议标题" :rules="[{ required: true, message: '请填写会议标题' }]"/>
            </div>
            <!--         //会议类型-->
<!--            <div class="item">-->
<!--            <div class="f-z-14-c m-t-10">会议类型</div>-->
<!--            <van-field v-model="meetingType" name="meetingType" placeholder="请输入会议的类型"/>-->
<!--            </div>-->

            <div class="item">
            <div class="f-z-14-c m-t-10">会议时间</div>
                <van-field
                    v-model="openCardDate"
                    name="openCardDate"
                    placeholder="请选择开始时间"
                    is-link
                    required
                    @click="showDatePicker = true"
                    readonly
                    :rules="[{ required: true, message: '请选择开始时间' }]"
                />
                <van-popup v-model:show="showDatePicker" position="bottom">
                    <van-calendar  v-model:show="showDatePicker"
                                   @confirm="onConfirm"/>
                </van-popup>
<!--                <van-cell :title="title" :value="date" @click="show1 = true" />-->
<!--                <van-calendar v-model:show="show1" @confirm="onConfirm" />-->
            </div>
<!--            省市区-->
         <div class="item">
             <div class="f-z-14-c m-t-10">会议城市</div>
             <van-field
                 v-model="fieldValue"
                 is-link
                 readonly
                 placeholder="请选择所在地区"
                 @click="showArea = true"
             />
             <van-popup v-model:show="showArea" round position="bottom">
                 <van-cascader
                     v-model="cascaderValue"
                     title="请选择所在地区"
                     :options="options"
                     @close="showArea = false"
                     @finish="onFinish"
                 />
             </van-popup>
         </div>
<!--            //详细地址-->
            <div class="item">
                <div class="f-z-14-c m-t-10">会议详细地址</div>
                <van-field v-model="createMeetingParams.meetingAddress" name="meetingType" placeholder="请输入会议详细地址"/>
            </div>
<!--            //meetingPeople-->
            <div class="item">
                <div class="f-z-14-c m-t-10">参会人</div>
                <van-field v-model="createMeetingParams.meetingPeople" name="meetingType" placeholder="请输入会议参会人"/>
            </div>
            <!--     </div>-->
            <div class="f-z-14-c m-t-10"><span style="color: red">*</span>附件上传 <span class="f-z-12-c">(文件格式仅限.doc, .docx, .txt, .xls, .xlsx, .pdf)</span></div>
            <div class="f-z-16">
                <van-uploader v-model="fileList" :before-read="beforeRead" :after-read="afterRead" :max-count="1" accept=".doc,.docx,.pdf,.ppt,.pptx,.xlsx,.xls">
                    <van-button icon="plus" type="primary" size="small">上传文件</van-button>
                </van-uploader>
            </div>
        </div>
    </div>


    <div class="cardTable">
        <van-button type="primary" block style="width: 92%; margin: 20px" :loading="isLoading" @click="handelMetting"  loading-text="文档记录"
        >文档记录
        </van-button
        >
    </div>
<!--    <van-dialog-->
<!--            v-model:show="show"-->
<!--            :show-cancel-button="false"-->
<!--            :showConfirmButton="false"-->
<!--    >-->
<!--        <van-form @submit="onSubmit" style="padding: 10px" ref="formRef">-->
<!--            <van-cell-group inset>-->
<!--                <van-field-->
<!--                        v-model="reasonForBorrowing"-->
<!--                        name="会议标题"-->
<!--                        label="会议标题"-->
<!--                        placeholder="会议标题"-->
<!--                        :rules="[{ required: true, message: '请填写会议标题' }]"-->
<!--                />-->
<!--            </van-cell-group>-->
<!--            <div style="margin: 16px 0;display:flex">-->
<!--                <van-button block style="margin-right: 10px" @click="cancel" size="small">-->
<!--                    取消-->
<!--                </van-button>-->
<!--                <van-button block type="primary" native-type="submit" size="small">-->
<!--                    提交-->
<!--                </van-button>-->
<!--            </div>-->
<!--        </van-form>-->
<!--    </van-dialog>-->
</template>

<script setup lang="ts">
import {showFailToast, showSuccessToast, showToast} from 'vant';
import huiyijilu from '../../assets/img/huiyijilu.png'
import {showConfirmDialog} from 'vant';
import {ref, inject, onMounted} from "vue";
import {useRouter} from "vue-router";
import {createMeeting, uploadMeetingFile} from '../../services/task-processing/index'
import { useCascaderAreaData } from '@vant/area-data';
import zanwupiaoju from "../../assets/img/zanwupiaoju.png";
interface createMeetingType{
    meetingId:  number |undefined, //会议ID
    meetingTitle: string | undefined, //会议标题
    meetingType: string, //会议类型   audio：音频会议    document：文档会议
    meetingTime: any, //会议时间  年-月-日  时：分：秒
    meetingCity: string, //会议城市
    meetingAddress: string, //会议地址
    meetingPeople: string, //参会人
    corpCode: string |null, //企业code
    creater: string, //创建人id
}
const value = ref<string>('')
const show = ref(false);
const username = ref("");
const meetingType = ref<string | undefined>()
const indexValue = ref()
const fileList = ref<never[]>([]);
const formRef = ref<any>(null); // 定义表单引用
const applyId = ref<number | undefined>()
const openCardDate = ref<string>('')
const date = ref('');
const showDatePicker = ref(false);
const isLoading = ref<boolean>(false)
const formatDate = (date) => {
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();

    // 如果月份或者日期小于 10，则在前面加上 '0'，否则保持原样
    const formattedMonth = month < 10 ? `0${month}` : month;
    const formattedDay = day < 10 ? `0${day}` : day;

    return `${year}/${formattedMonth}/${formattedDay}`;
};
const onConfirm = (value) => {
    showDatePicker.value = false;
    date.value = formatDate(value);
    openCardDate.value = date.value
};
// const meetingPeople = ref<string|undefined>()  //参会人
// const currentDate = ref<string>();
// const minDate =  ref<string | undefined>()
// const maxDate = ref<string | undefined>()
let params = ref<any>(
    {}
)
let createMeetingParams =ref<createMeetingType>({
    meetingId: 0, //会议ID
    meetingTitle: "", //会议标题
    meetingType: "document", //会议类型   audio：音频会议    document：文档会议
    meetingTime: "", //会议时间  年-月-日  时：分：秒
    meetingCity: "", //会议城市
    meetingAddress: "", //会议地址
    meetingPeople: "", //参会人
    corpCode: "", //企业code
    creater: "", //创建人id
})
const showArea = ref(false);
const fieldValue = ref('');
const cascaderValue = ref('');
const options = useCascaderAreaData();
const onFinish = ({ selectedOptions }) => {
    showArea.value = false;
    fieldValue.value = selectedOptions.map((option) => option.text).join('/');
};
interface ListItem {
    time: string;
    date: string;
    long: string;
}

//上传的 前置 处理
const allowedExtensions = ['.doc', '.docx', '.txt', '.xls', '.xlsx', '.pdf'];

const beforeRead = (file: File) => {
    console.log('file', file)
    const extension = file.name.split('.').pop()?.toLowerCase(); // 获取文件扩展名并转换为小写
    if (!extension || !allowedExtensions.includes(`.${extension}`)) {
        showToast('请上传指定格式的文件');
        return false;
    }
    return true;
};
//文件读取完成后的回调函数
const afterRead = () => {
    console.log(fileList.value)
    params.value.meetingId = String(applyId.value)
    params.value.file = fileList.value[0].file
    // console.log('params',params.value)
}
//生成随机数的函数
const generateTwoDigitNumber = () => {
    const randomNumber = Math.floor(Math.random() * 900) + 100; // 生成100到999之间的随机数
    return randomNumber.toString(); // 将随机数转换为字符串并返回
}

const list = ref<Array<ListItem>>([]);
const deleteItem = (index: any) => {
    indexValue.value = index
    showConfirmDialog({
        title: '提示',
        message:
            '你确定要删除吗？',
    })
        .then(() => {
            list.value.splice(indexValue.value, 1)
        })
        .catch(() => {
            // on cancel
        });
}
const reasonForBorrowing = ref<string>()
// const onSubmit = (values: any) => {
//     const valid = formRef.value.validate();
//     if (valid) {
//         //掉 会议上传的接口 未成功给提示 成功跳转到 会议列表页面
//         params.value.meetingTitle = reasonForBorrowing.value
//         uploadMeetingFile(params.value).then((res: any) => {
//             if (res.success) {
//                 createMeetingParams.value.meetingId = applyId.value
//                 createMeetingParams.value.meetingTitle = reasonForBorrowing.value
//                 createMeetingParams.value.meetingCity = fieldValue.value
//                 createMeetingParams.value.meetingTime = date.value
//                 console.log('createMeetingParams',createMeetingParams)
//                 //调取会议创建的接口
//                 createMeeting(createMeetingParams.value).then((res:any)=>{
//                     if(res.success){
//                         showSuccessToast(res.msg)
//                         router.replace('/business');
//                     }else{
//                         showFailToast(res.msg)
//                     }
//                 })
//             }else {
//                 showFailToast(res.msg)
//             }
//         })
//
//     }
// };
const cancel = () => {
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
    if(!reasonForBorrowing.value){
        showFailToast('请填写会议标题！')
    }else if(fileList.value.length ==0 ) {
        showFailToast('请先上传附件！')
    }else{
        params.value.meetingTitle = reasonForBorrowing.value
        isLoading.value = true
        // const uploadMeetingFileHandle = ()=>{
        //     uploadMeetingFile(params.value).then((res: any) => {
        //
        //     })
        // }
        uploadMeetingFile(params.value).then((res: any) => {
            if (res.success) {
                createMeetingParams.value.meetingId = applyId.value
                createMeetingParams.value.meetingTitle = reasonForBorrowing.value
                createMeetingParams.value.meetingCity = fieldValue.value
                createMeetingParams.value.meetingTime = date.value
                console.log('createMeetingParams',createMeetingParams)
                //调取会议创建的接口
                createMeeting(createMeetingParams.value).then((res:any)=>{
                    if(res.success){
                        isLoading.value = false
                        showSuccessToast(res.msg)
                        router.replace('/business');
                    }else{
                        showFailToast(res.msg)
                    }
                })
            }else {
                showFailToast(res.msg)
            }
        })
    }

};
onMounted(async () => {
    const currentDate = new Date();
    const year = currentDate.getFullYear();
    const month = String(currentDate.getMonth() + 1).padStart(2, '0');
    const day = String(currentDate.getDate()).padStart(2, '0');
    const hours = String(currentDate.getHours()).padStart(2, '0');
    const minutes = String(currentDate.getMinutes()).padStart(2, '0');
    const seconds = String(currentDate.getSeconds()).padStart(2, '0');
    const formattedDate = `${year}${month}${day}${hours}${minutes}${seconds}`;
    let number = generateTwoDigitNumber()
    applyId.value = `WD` + formattedDate + number
    // reasonForBorrowing.value = `${year}-${month}-${day}` + '文档会议记录'
    let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
    params.value.userId = userInfoData.userInfo.userId
    params.value.corpCode = await localStorage.getItem('corpCode') //从本地获取corpCode
    createMeetingParams.value.corpCode = await localStorage.getItem('corpCode') //从本地获取corpCode
    createMeetingParams.value.creater = userInfoData.userInfo.userId
    // minDate.value = currentDate; // 设置最大日期为当前日期加上五年
})
</script>

<style lang="less" scoped>
.cardTable {
  width: 100%;
  // margin-top: 20px;
  text-align: center;
  //position: fixed;
  //bottom: 60px;
}

.uploader {
  //text-align:center;
  margin-left: 10px;
}

//.cardTable {
//  width: 100%;
//  // margin-top: 20px;
//  text-align: center;
//  position: fixed;
//  bottom: 60px;
//}

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
  padding: 2vw 4vw;
  margin: 10px;
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

  .header_title {
    display: flex;
    align-items: center;
    font-size: 1em;
    margin-bottom: 2vw;

    span {
      margin-left: 10px;
    }
  }
}

.f-w-550 {
  font-weight: 550;
}
.item{
    padding: 2vw 0;
}
.f-z-14-c {
  //color:rgba(23,26,29,0.40);
  font-size: 0.875em;
  font-weight: 550;
}
.f-z-12-c{
    color:rgba(23,26,29,0.40);
    font-size: 0.875em;
}
.m-t-10 {
  margin: 2vw 0;
}

.f-z-16 {
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
.empty {
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

::v-deep(.van-cell) {
  padding: 0;
}
::v-deep(.van-field__label){
    margin: 0;
}
::v-deep(.van-cell-group--inset){
    margin: 0;
}

</style>
