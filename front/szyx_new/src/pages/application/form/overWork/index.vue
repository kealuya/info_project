<template>
  <div>
    <nav-bar title="加班申请" />
    <div>
      <van-form @submit="onSubmit">
        <div class="formCard">
          <div class="pl5 f16"><span class="require">*</span> 加班类型</div>
          <van-field
              v-model="workType"
              name="workType"
              placeholder="请选择加班类型"
              is-link
              required
              @click="showPicker = true"
              readonly
              :rules="[{ required: true, message: '请选择加班类型' }]"
          />
          <van-popup v-model:show="showPicker" position="bottom">
            <van-picker
                :columns="WorkList"
                @confirm="onConfirm"
                @cancel="showPicker = false"
            />
          </van-popup>
        </div>
        <div v-if="workDay==false">
          <div class="formCard">
            <div class="pl5 f16"><span class="require">*</span>开始结束时间</div>
            <van-field
                v-model="overWorkDate"
                name="overWorkDate"
                placeholder="请选择开始结束时间"
                is-link
                required
                @click="showDatePicker = true"
                readonly
                :rules="[{ required: true, message: '请选择开始结束时间' }]"
            />
            <van-popup v-model:show="showDatePicker" position="bottom">
              <van-calendar v-model:show="showDatePicker" type="range" @confirm="onConfirmDate"/>
            </van-popup>
          </div>
        </div>
        <div v-else>
          <div class="formCard">
            <div class="pl5 f16"><span class="require">*</span>开始时间</div>
            <van-field
                v-model="overWorkStartDate"
                name="overWorkDate"
                placeholder="请选择开始时间"
                is-link
                required
                @click="showStartDatePicker = true"
                readonly
                :rules="[{ required: true, message: '请选择开始时间' }]"
            />
            <van-popup v-model:show="showStartDatePicker" position="bottom">
              <van-calendar :min-date="minDate"  v-model:show="showStartDatePicker"
                            @confirm="onStartConfirmDate"/>
            </van-popup>
          </div>
          <div class="formCard">
            <div class="pl5 f16"><span class="require">*</span>结束时间</div>
            <van-field
                v-model="overWorkEndDate"
                name="overWorkDate"
                placeholder="请选择结束时间"
                readonly
                required
                :rules="[{ required: true, message: '请选择结束时间' }]"
            />
<!--      @click="showEndDatePicker = true"
                readonly   is-link    -->
<!--            <van-popup v-model:show="showEndDatePicker" position="bottom">-->
<!--              <van-calendar v-model:show="showEndDatePicker" @confirm="onEndConfirmDate"/>-->
<!--            </van-popup>-->
          </div>
        </div>
        <div class="formCard">
          <div class="pl5 f16"><span class="require">*</span>加班天数 (天)</div>
          <van-field
              v-model="overWorkDate"
              name="overWorkDate"
              required
              readonly
              :rules="[{ required: true, message: '请选择' }]"
          >
            <template #input>
              <van-stepper v-model="overWorkDays" min="0" :max="overWorkSize" step="0.5" readonly/>
            </template>
          </van-field>
        </div>
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
          <div class="pl5 f16">
            <van-row align="center">
              <van-col span="21">
                <span class="require">*</span>抄送信息
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
                <span class="f14 graytxt">请选择抄送人</span>
              </van-col>
              <van-col span="1">
                <van-image :src="selectPerson" @click="addCcUser()" width="25" height="25"/>
              </van-col>
            </van-row>
          </div>

        </div>
        <div class="formCard" v-if="toUserList.length!=0">
          <div class="pl5 f16 mb5">
            <van-row align="center">
              <van-col span="21" class="f14">
                审批信息
              </van-col>
            </van-row>
            <div class="mt10 mb10">
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
          <div class="pl5 f16 mb10" v-if="ccUser.value!=' '">
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
    <select-peron-picker @select-over="selectApprove" ref="selectFlowOver"    :list="selectPersonList"/>
  </van-popup>
  <van-popup v-model:show="selectCopyPeoplePicker" style="width:100vw;height:100vh" position="bottom">
    <select-copy-people-picker @select-over="selectCC" ref="selectCCOver"    :list="selectCopyPersonList"/>
  </van-popup>
</template>

<script lang="ts" setup>
import {ref, inject, onMounted} from "vue";
import {createWorkApply} from "../../../../services/apply";
import {getApprovalFlowHistory,getApprovalFlowList} from "../../../../services/current";
import selectPerson from '../../../../assets/icon/selectPerson.png';
import addPerson from '../../../../assets/icon/addPerson.png'
import deletePerson from '../../../../assets/icon/deletePerson.png'
import SelectPeronPicker from '../../../../components/selectComponent/selectPersonPicker/index.vue' // 自定义组件引入需要大写名称
import SelectCopyPeoplePicker from "../../../../components/selectComponent/selectCopyPeoplePicker/index.vue" // 选择抄送人
import NavBar from '../../../../components/navBar/index.vue'
import {orderNumberM} from '../ts/orderNumberM';
import {showSuccessToast} from "vant";
import router from "../../../../router";
import {showFailToast} from "vant/lib/toast/function-call";
import moment from "moment/moment";

const onClickLeft = () => history.back();
const workType = ref('');
const overWorkStartDate = ref('');// 请假开始时间
const overWorkEndDate = ref(''); //请假结束时间
const workDay = ref(true); //
const overWorkDate = ref(''); //请假时间展示
const overWorkDays = ref(''); //请假天数
const approveReason = ref(''); //申请事由
const loading=ref(false); // 加载组件控制
const flowId = ref(''); //申请事由
const flowName = ref(''); //申请事由
const overWorkSize = ref(0); //限制最大天数
const showPicker = ref(false); // 控制请假类型
const showDatePicker = ref(false); // 控制开始结束时间 （适用于开始结束日期// ）
const showStartDatePicker = ref(false); // 控制开始结束时间 （适用于工作日开始日期// ）
const showEndDatePicker = ref(false); // 控制开始结束时间 （适用于工作日结束日期// ）
const showApprPicker = ref(false); // 控制审批流程
const selectFlowPeoplePicker = ref(false); // 控制选择审批人
const selectCopyPeoplePicker=ref(false); // 控制选择抄送人
const toUserList = ref([]); // 获取到的审批人
const ccUser = ref(''); // 获取到的抄送人
let selectPersonList=[]; // 用于存放审批人列表
const selectCopyPersonList=ref([]); // 用于存放抄送人列表
const nowDate = new Date();
let minDate = new Date(nowDate.getFullYear(), nowDate.getMonth()-6, nowDate.getDay());
let userInfoData: any = inject("userInfo"); //获取审批数组
const selectkey = ref();
const customFieldName = {
  text: 'flow_name',
  value: 'flow_id',
};
const WorkList = [
  {text: '工作日', value: '工作日'},
  {text: '休息日', value: '休息日'},
  {text: '法定节假日', value: '法定节假日'},
  {text: '其他', value: '其他'},
];
const approveList = ref([]); // 审批流程列表
const resultAps = ref<any>([]); // 审批列表
const copyPush = ref([]); // 抄送列表
const fileList = ref([]); // 图片上传
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
  selectFlowPeoplePicker.value=false;
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
        "ywNo": orderNumberM(),
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
  selectCopyPeoplePicker.value=false;
}
const onConfirm = (value: any) => {
  workType.value = value.selectedValues[0];
  if (workType.value == '工作日') {
    workDay.value = true;
    mountedUp();
  } else {
    workDay.value = false;
    overWorkStartDate.value = '';
    overWorkEndDate.value = '';
    overWorkDate.value = '';
    overWorkDays.value = '0';
    overWorkSize.value = 0;
    approveReason.value='';
  }
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
    "typeId": "FK_JB_SQ",
    'isOpen': '1',
    'empCode': userInfoData.userInfo.empCode,
    'equipmentType': "pc"
  }
  let result: any = await getApprovalFlowList(payload);
  approveList.value = result.data;
}
const onApprovalConfirm = (value: any) => {
  flowId.value = value.selectedOptions[0].flow_id;
  flowName.value = value.selectedOptions[0].flow_name;
  getFlowHistory(value.selectedValues[0])
  showApprPicker.value = false;
};
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
  ccUser.value = userName.join(',')
  ccUser.value=ccUser.value.substring(0,ccUser.value.length-1);
}
const onConfirmDate = (values: any) => {
  const [start, end] = values;
  showDatePicker.value = false;
  overWorkDays.value = ((end - start) / (1000 * 3600 * 24)).toString();
  overWorkSize.value = (end - start) / (1000 * 3600 * 24);
  overWorkStartDate.value = moment(start).format("yyyy-MM-DD");
  overWorkEndDate.value = moment(end).format("yyyy-MM-DD");
  overWorkDate.value = `${overWorkStartDate.value} - ${overWorkEndDate.value}`;
};
const onStartConfirmDate = (values: any) => {
  overWorkStartDate.value = `${moment(values).format("yyyy-MM-DD")} 17:30:00`
  overWorkEndDate.value = `${moment(values).format("yyyy-MM-DD")} 21:00:00`
  overWorkDays.value='0.5';
  overWorkSize.value=0.5;
  showStartDatePicker.value = false;
};
const onEndConfirmDate = (values: any) => {
  overWorkEndDate.value = `${moment(values).format("yyyy-MM-DD")} 21:00:00`
  showEndDatePicker.value = false;
};
const onSubmit = async (values: any) => {
  const workDetail: any = {
    "approveNo": orderNumberM(),
    "empName": userInfoData.userInfo.userName,
    "empCode": userInfoData.userInfo.empCode,
    "corpCode": userInfoData.userInfo.corpCode,
    "deptName": userInfoData.userInfo.deptName,
    "deptId": userInfoData.userInfo.deptId,
    "workTime": overWorkDays.value.toString(),
    "workStartDate": overWorkStartDate.value,
    "flowId": flowId.value,
    "workEndDate": overWorkEndDate.value,
    "overtimeType": workType.value,
  };
  const rowShowJson: any = [
    {'label': '申请人员', 'value': userInfoData.userInfo.userName, 'isImportant': '0'},
    {'label': '申请部门', 'value': userInfoData.userInfo.deptName, 'isImportant': '0'},
    {'label': '申请事由', 'value': approveReason.value, 'isImportant': '1'},
    {'label': '请假类型', 'value': workType.value, 'isImportant': '1'},
    {'label': '开始日期', 'value': overWorkStartDate.value, 'isImportant': '1'},
    {'label': '结束日期', 'value': overWorkEndDate.value, 'isImportant': '1'},
    {'label': '休假天数', 'value': overWorkDays.value, 'isImportant': '1'},
  ];
  const postData: any = {
    'approveNo': orderNumberM(),
    'creater': userInfoData.userInfo.userName,
    'empCode': userInfoData.userInfo.empCode,
    'empName': userInfoData.userInfo.userName,
    'deptName': userInfoData.userInfo.deptName,
    'approveReason': approveReason.value,
    "bz1": "",
    "useStyle": "",
    "sealType": "",
    "totalMinutes": "",
    "approvalName": userInfoData.userInfo.userName,
    "approvalEmpCode": userInfoData.userInfo.empCode,
    "dateRange": [overWorkStartDate.value, overWorkEndDate.value],
    "approveType": "加班申请",
    "workDetail": workDetail,
    "approvalMode": approveList.value.length != 0 ? '0' : '1',
    "approvalFlowId": flowId.value,
    "agentPeopleDetail": [],
    "rowShowJson": JSON.stringify(rowShowJson),
    "ccUserList": [],
    "corpCode": userInfoData.userInfo.corpCode,
    "equipmentType": ""
  }
  loading.value=true
  let result: any = await createWorkApply(postData);
  if(result.success==true){
    loading.value=false;
    showSuccessToast(result.msg);
    router.back();
  }else{
    loading.value=false;
    showFailToast(result.msg);
  }
};
const mountedUp = () => {
  let corpCode = userInfoData.userInfo.corpCode
  if (corpCode == 'szht6666') {
    workType.value = '工作日';
    overWorkStartDate.value = `${moment(nowDate).format("yyyy-MM-DD")} 17:30:00`
    overWorkEndDate.value = `${moment(nowDate).format("yyyy-MM-DD")} 21:00:00`
    overWorkDays.value='0.5';
    overWorkSize.value=0.5;
    approveReason.value='晚上加班至九点'
  }
}
onMounted(() => {
  getDocList();
  let approvalPeople = {
    'no': 1,
    'toUserId': '',
    'toUserName': '',
    'currentNo': '',
    'toRole': '',
  }
  resultAps.value.push(approvalPeople);
  mountedUp();
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
