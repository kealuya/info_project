<template>
  <div>
<!--    <van-sticky>-->
<!--      <van-nav-bar-->
<!--          title="快递寄送"-->
<!--          left-arrow-->
<!--          @click-left="onClickLeft"-->
<!--      />-->
<!--    </van-sticky>-->
    <NavBar title="快递寄送"/>
    <div>
      <van-form @submit="onSubmit">
        <div class="formCard">
          <div class="pl5 f16"><span class="require">*</span>申请事由</div>
          <van-field
              v-model="approveReason"
              name="approveReason"
              placeholder="请输入申请事由"
              rows="3"
              type="textarea"
              autosize
              required
              :rules="[{ required: true, message: '请输入申请事由' }]"
          />
        </div>
        <div class="formCard">
          <div class="pl5 f16"><span class="require">*</span> 支付方式</div>
          <van-field
              v-model="useStyle"
              name="workType"
              placeholder="请选择支付方式"
              is-link
              required
              @click="showPicker = true"
              readonly
              :rules="[{ required: true, message: '请选择领用用途' }]"
          />
          <van-popup v-model:show="showPicker" position="bottom">
            <van-picker
                :columns="useClassList"
                @confirm="onConfirm"
                @cancel="showPicker = false"
            />
          </van-popup>
        </div>
        <div class="formCard">
          <div class="pl5 f16"><span class="require">*</span>预计寄送时间</div>
          <van-field
              v-model="outWorkDate"
              name="openCardDate"
              placeholder="请选择预计寄送时间"
              is-link
              required
              @click="showDatePicker = true"
              readonly
              :rules="[{ required: true, message: '请选择预计使用时间' }]"
          />
          <van-popup v-model:show="showDatePicker" position="bottom">
            <van-calendar  v-model:show="showDatePicker"
                           :min-date="minDate"
                           type="range"
                           @confirm="onConfirmDate"/>
          </van-popup>
        </div>
        <div class="formCard">
          <div class="pl5 f16"><span class="require">*</span>预算金额</div>
          <van-field
              v-model="goodsPrice"
              name="goodsPrice"
              placeholder="请输入预算金额"
              type="number"
              required
              :rules="[{ required: true, message: '请输入预算金额' },{ pattern, message: '金额格式输入错误请重新输入' }]"
          />
        </div>
        <div class="formCard">
          <div class="pl5 f16">
            <van-row align="center">
              <van-col span="21">
                <span class="require">*</span>审批信息
                <span class="f12 graytxt" v-if="approveList.length!=0">免审需选择自己组内的负责人进行抄送告知</span>
                <span class="f12 graytxt" v-else>该审批采用自选审批人模式</span>
              </van-col>
              <van-col span="1" v-if="approveList.length==0">
                <van-image :src="addPerson" @click="addFlow" width="25" height="25"/>
              </van-col>
            </van-row>
          </div>
          <div v-if="approveList.length!=0">
            <van-field
                v-model="flowName"
                name="flowName"
                placeholder="请选择抄送流程信息"
                is-link
                required
                @click="showApprPicker = true"
                readonly
                :rules="[{ required: true, message: '请选择审批流程' }]"
            />
            <van-popup v-model:show="showApprPicker" position="bottom">
              <van-picker
                  :columns="approveList"
                  @confirm="onApprovalConfirm"
                  :columns-field-names="customFieldName"
                  @cancel="showApprPicker = false"
              />
            </van-popup>
          </div>
          <div v-else class="mt10 mb10">
            <van-row class="pt10 pb10" v-for="(item,index) in resultAps" align="center">
              <van-col span="18">
                <div class="number">{{ index + 1 }}</div>
                <span class="f14 graytxt"><span
                    v-if="item.toUserName != ''">{{
                    item.isAllPass ? item.isAllPass == '1' ? '或审：' : '会审：' : '单审：'
                  }}</span>{{
                    item.toUserName == '' ? '请选择审批人' : item.toUserName
                  }}</span>
              </van-col>
              <van-col span="5" v-if="index!=0">
                <van-image :src="deletePerson" @click="deleteFlow(index)" width="25" height="25"/>
                <van-image :src="selectPerson" @click="addCcUser(index)" class="ml20" width="25" height="25"/>
              </van-col>
              <van-col span="2" v-else>
                <van-image :src="selectPerson" @click="addCcUser(index)" style="margin-left: 45px" width="25"
                           height="25"/>
              </van-col>
            </van-row>
            <van-row class="pt10 pb10">
              <van-col span="21">
                <div class="number">抄</div>
                <span class="f14 graytxt">{{ccUser==''?'请选择抄送人':`抄送： ${ccUser}`}}</span>
              </van-col>
              <van-col span="1">
                <van-image :src="selectPerson" @click="addCcPerson()" width="25" height="25"/>
              </van-col>
            </van-row>
          </div>

        </div>
        <!--     设定审批流回显   -->
        <div class="formCard" v-if="toUserList.length!=0">
          <div class="pl5 f16 mb5">
            <van-row align="center">
              <van-col span="21" class="f14">
                审批信息
              </van-col>
            </van-row>
            <div class="mt10 mb10 pb10">
              <van-row v-for="item in toUserList">
                <van-col :span="1">
                  <div class="tc" style="flex: 1; text-align: center">
                    <div
                        class="ml10"
                        style="width: 15px;height: 15px;border: 1px solid #f1f2f3;  border-radius: 20px;background: #c6c6c6"></div>
                    <div v-if="toUserList.length>1&&item.no!=toUserList.length">
                      <div :id="'bottom'+item.no " :style="'height:'+item.children.length*26+'px'"
                           style="width:30px; border-left: 2px solid #c6c6c6;margin-left: 18px; "></div>
                    </div>

                  </div>
                </van-col>
                <van-col :span="20" :offset="1">
                  <div class="f14 ml5">
                    第{{ item.no }}级
                    <span>{{ item.isAllPass ? item.isAllPass == '1' ? '或审' : '会审' : '单审' }}</span>
                  </div>
                  <div v-for="key in item.children">
                    <span class="f14 ml5">{{ key.toUserName }}</span>
                    <span class="f14 ml5">({{ key.toRole }})</span>
                  </div>
                </van-col>
              </van-row>
            </div>
          </div>
          <div class="pl5 f16 mb10" v-if="ccUser.value!=null">
            <van-row align="center" style="padding-bottom: 10px">
              <van-col span="23" class="f14">
                <div>抄送信息</div>
                <span class="graytxt ml15 mt10  f12" style="font-weight: bolder">抄送: {{ ccUser }}</span>
              </van-col>
            </van-row>
          </div>
        </div>
        <div style="margin: 16px;">
          <van-button block type="primary" native-type="submit">
            提交
          </van-button>
        </div>
      </van-form>
    </div>
  </div>
  <div v-show="loading">
    <loading-component :title="'提交中'"/>
  </div>
  <van-popup v-model:show="selectFlowPeoplePicker" style="width:100vw;height:100vh" position="bottom">
    <select-peron-picker @select-over="selectApprove" ref="selectFlowOver"   :list="selectPersonList"/>
  </van-popup>
  <van-popup v-model:show="selectCopyPeoplePicker" style="width:100vw;height:100vh" position="bottom">
    <select-copy-people-picker  @select-over="selectCC" :list="selectCopyPersonList"/>
  </van-popup>
</template>

<script lang="ts" setup>
import {ref, inject, onMounted} from "vue";
import {createLogApply,} from "../../../../services/apply";
import {getApprovalFlowHistory,getApprovalFlowList} from '../../../../services/current';
import selectPerson from '../../../../assets/icon/selectPerson.png';
import addPerson from '../../../../assets/icon/addPerson.png'
import deletePerson from '../../../../assets/icon/deletePerson.png'
import SelectPeronPicker from '../../../../components/selectComponent/selectPersonPicker/index.vue' // 自定义组件引入需要大写名称 （此处是选择）
import SelectCopyPeoplePicker from "../../../../components/selectComponent/selectCopyPeoplePicker/index.vue" // 此处是选择抄送人组件列表
import NavBar from "../../../../components/navBar/index.vue";
import {orderNumberM} from '../ts/orderNumberM';
import {showSuccessToast} from "vant";
import router from "../../../../router";
import {showFailToast} from "vant/lib/toast/function-call";
import moment from "moment/moment";
import LoadingComponent from '../../../../components/loading/index.vue' // 加载组件
const onClickLeft = () => history.back();
const useStyle = ref(''); //使用方式
const loading=ref(false); //loading加载
const ywNo=ref('');
const showDatePicker= ref(false); // 控制补卡时间
const personType=ref(''); // 审批人or抄送人  approve 审批人
const pattern = /^[0-9]+(\.[0-9]{1,2}|)$/; // 保留两位小数的金额校验
const approveReason = ref(''); //申请事由
const flowId = ref(''); //审批flowId
const flowName = ref(''); //审批名字
const showPicker = ref(false); // 控制请假类型
const showApprPicker = ref(false); // 控制审批流程
const selectFlowPeoplePicker = ref(false); // 控制选择审批人
const selectCopyPeoplePicker =ref(false); // 控制选择抄送人
const nowDate = new Date();
let minDate = new Date(nowDate.getFullYear(), nowDate.getMonth()-6, nowDate.getDay());
const useClassList = [
  {text: '因公月结', value: '因公月结'},
];
const toUserList = ref([]); // 获取到的审批人
const ApproveUserList = ref([]); // 获取到的审批人
const ccUser = ref(''); // 获取到的抄送人
let userInfoData: any = inject("userInfo"); //获取审批数组
const selectFlowOver =ref();
const selectkey = ref();
let selectPersonList= []; // 用于存放审批人列表
let selectCopyPersonList=[]; // 用于存放抄送人列表
const customFieldName = {
  text: 'flow_name',
  value: 'flow_id',
};

const approveList = ref([]); // 审批流程列表
const resultAps = ref<any>([]); // 审批列表
const copyPush = ref([]); // 抄送列表
const fileList = ref([]); // 图片上传
const startTime=ref(''); // 开始时间
const endTime=ref(''); // 结束时间
const outWorkDate=ref(''); //开始结束时间展示
const goodsPrice=ref(''); //开始结束时间展示
const onConfirmDate = (values: any) => {
  const [start, end] = values;
  showDatePicker.value = false;
  startTime.value = moment(start).format("yyyy-MM-DD");
  endTime.value =  moment(end).format("yyyy-MM-DD");
  outWorkDate.value = `${startTime.value} - ${endTime.value}`;
};
// 选择审批人
const selectApprove=(list:any)=>{
  if(list.length!=0){
    let toSelectUserList: any = [];
    list.forEach((item: {
      empCode: any; empName: any;
    }) => {
      return toSelectUserList.push({...item, toUserName: item.empName, toUserId: item.empCode});
    })
    let userName: any = [];
    list.forEach((item: any) => userName.push(item.empName));
    let data = {
      'no': selectkey,
      'isAllPass': list[0].isAllPass,
      'toUserName': userName.length <= 3 ? userName.join(' , ') : `${userName.splice(2, userName.length).join()}等${list.length}人`,
      'list': toSelectUserList,
    }
    resultAps.value[selectkey.value] = data;
  }
  selectFlowPeoplePicker.value = false;
}
// 选择抄送人
const selectCC = (list: any) => {
  if(list.length!=0){
    ccUser.value='';
    let toSelectCCUserList: any = [];
    list.forEach((item: {
      corpCode: any;
      jobName: any;
      empCode: any; empName: any;
    }) => {
      let data={
        "ywNo": ywNo.value,
        "ccUserId": item.empCode,
        "ccUser": item.empName,
        "toRole": item.jobName,
        "school": item.corpCode,
        "type": "fk"
      };
      ccUser.value=ccUser.value+item.empName+'，';
      return toSelectCCUserList.push(data);
    })
    ccUser.value=ccUser.value.substring(0,ccUser.value.length-1);
    copyPush.value=toSelectCCUserList;
  }
  selectCopyPeoplePicker.value = false;
}
// 整理提交的数组
const concatList=()=>{
  // approveList
  let apprlist: any=[];
  for (var i = 0;i<resultAps.value.length;i++){
    for(var j=0;j<resultAps.value[i].list.length;j++){
      let data={
        "no": i+1,
        "toUserId": resultAps.value[i].list[j].empCode,
        "toUserName": resultAps.value[i].list[j].empName,
        "isAllPass": resultAps.value[i].list[j].isAllPass,
        "currentNo": j,
        "toRole": resultAps.value[i].list[j].jobName
      }
      apprlist.push(data);
    }
  }
  ApproveUserList.value=apprlist
}
// 选择使用方式
const onConfirm = (value: any) => {
  useStyle.value = value.selectedValues[0];
  showPicker.value = false;
};
const addCcUser = (index: any) => {
  selectkey.value = index;
  selectPersonList=[];
  if(resultAps.value[index].list){
    selectPersonList=resultAps.value[index].list;
  }
  selectFlowPeoplePicker.value = true;
}
const addCcPerson = (index: any) => {
  personType.value='cc'
  selectkey.value = copyPush.value;
  selectCopyPeoplePicker.value = true;
}
// 添加审批流程
const addFlow = () => {
  let approvalPeople = {
    'no': resultAps + 1,
    'toUserId': '',
    'toUserName': '',
    'currentNo': '',
    'toRole': '',
  }
  resultAps.value.push(approvalPeople);
};
const deleteFlow = (index: any) => {
  resultAps.value.splice(index, 1)
};
//  获取审批列表
const getDocList = async () => {
  let payload = {
    "corpCode": userInfoData.userInfo.corpCode,
    'deptId': userInfoData.userInfo.deptId,
    "typeId": "FK_WL_SQ",
    'isOpen': '1',
    'empCode': userInfoData.userInfo.empCode,
    'equipmentType': ""
  }
  let result: any = await getApprovalFlowList(payload);
  approveList.value = result.data;
}
// 设定审批流的情况下
const onApprovalConfirm = (value: any) => {
  flowId.value = value.selectedOptions[0].flow_id;
  flowName.value = value.selectedOptions[0].flow_name;
  getFlowHistory(value.selectedValues[0])
  showApprPicker.value = false;
};
// 根据审批流回显审批流人员信息
const getFlowHistory = async (value: any) => {
  let payload = {
    "school": userInfoData.userInfo.corpCode,
    'flowId': value,
    "corpCode": userInfoData.userInfo.corpCode,
    'empCode': userInfoData.userInfo.empCode,
    'equipmentType': ""
  }
  let result: any = await getApprovalFlowHistory(payload);
  let flowHistory: any = JSON.parse(result.data);
  toUserList.value = flowHistory.data.toUserList;
  const object: any = {};
  // ccUserList.value=flowHistory.data.ccUserList;
  for (let i = 0; i <= toUserList.value.length - 1; i++) {
    let no = toUserList.value[i]['no'];
    let isAllPass = toUserList.value[i]['isAllPass'];
    if (!object[no]) {
      object[no] = {
        no,
        isAllPass,
        children: []
      }
    }
    if (no === object[no].no) {
      object[no].children.push(toUserList.value[i])
    }
  }
  toUserList.value = [];
  Object.values(object).forEach((item) => {
    return toUserList.value.push(item as never);
  })
  let userName: any = [];
  flowHistory.data.ccUserList.forEach((item: any) => userName.push(`${item.ccUserName}(${item.toRole})`));
  ccUser.value = userName.join(',');
  ccUser.value=ccUser.value.substring(0,ccUser.value.length-1);
}

// 提交用章申请
const onSubmit = async (values: any) => {
  if(approveList.value.length == 0){
    concatList();
  }
  const logisticsApplyDetail: any = {
    "approveNo":ywNo.value,
    "empName": userInfoData.userInfo.userName,
    "empCode": userInfoData.userInfo.empCode,
    "corpCode": userInfoData.userInfo.corpCode,
    "deptName": userInfoData.userInfo.deptName,
    "deptId": userInfoData.userInfo.deptId,
    "flowId": flowId.value,
    "startTime": startTime.value,
    "endTime":endTime.value,
    "dateRange":[startTime.value,endTime.value],
    "reasons": approveReason.value,
    "contactPhone": "",
    "amount": goodsPrice.value,
    "paymentMethod": useStyle.value,
    "contactName": "",
    "contactAddresst": "",
    "arrivalPlace": "",
    "packageInfos": "",
    "useStyle": "",
  };
  const rowShowJson: any = [
    {'label': '申请人员', 'value': userInfoData.userInfo.userName, 'isImportant': '0'},
    {'label': '申请部门', 'value': userInfoData.userInfo.deptName, 'isImportant': '0'},
    {'label': '申请事由', 'value': approveReason.value, 'isImportant': '1'},
    {'label': '支付方式', 'value': useStyle.value, 'isImportant': '1'},
    {'label': '申请事由', 'value': approveReason.value, 'isImportant': '1'},
    {'label': '预计寄送时间', 'value': `${startTime.value} 至 ${ endTime.value}`, 'isImportant': '1'},
  ];
  const postData: any = {
    'approveNo':ywNo.value,
    'creater': userInfoData.userInfo.userName,
    'empCode': userInfoData.userInfo.empCode,
    'empName': userInfoData.userInfo.userName,
    'deptName': userInfoData.userInfo.deptName,
    'approveReason': approveReason.value,
    "bz1": "",
    "totalMinutes": "",
    "approvalName": userInfoData.userInfo.userName,
    "approvalEmpCode": userInfoData.userInfo.empCode,
    "timeStatus": false,
    "approveType": "快递寄送",
    "logisticsApplyDetail": logisticsApplyDetail,
    "approvalMode": approveList.value.length != 0 ? '0' : '1',
    "approvalFlowId": flowId.value,
    "agentPeopleDetail": ApproveUserList.value,
    "rowShowJson": JSON.stringify(rowShowJson),
    "ccUserList": copyPush.value,
    "corpCode": userInfoData.userInfo.corpCode,
    "equipmentType": ""
  }
  loading.value=true;
  let result: any = await createLogApply(postData);
  if(result.success==true){
    loading.value=false;
    showSuccessToast(result.msg);
    router.back();
  }else{
    loading.value=false;
    showFailToast(result.msg);
  }
};

onMounted(() => {
  ywNo.value=orderNumberM();
  getDocList();
  let approvalPeople = {
    'no': 1,
    'toUserId': '',
    'toUserName': '',
    'currentNo': '',
    'toRole': '',
  }
  resultAps.value.push(approvalPeople);
});
</script>

<style scoped>
.formCard {
  margin-top: 10px;
  padding-top: 10px;
  background: white;
}

.require {
  color: red;
  font-size: 16px;
}

.number {
  display: inline-block;
  width: 20px;
  height: 20px;
  border: 2px solid #333333;
  border-radius: 15px;
  margin-left: 10px;
  margin-right: 10px;
  text-align: center;
}
</style>
