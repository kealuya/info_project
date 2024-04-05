<template>
  <div>
    <navbar title="日常报销单"/>
    <div class="formCard">
      <van-row justify="space-between" align="center">
        <van-col>
          <div class="graytxt f14">{{ applyNo }}</div>
        </van-col>
        <van-col>
          <div class="graytxt f14">{{ nowDate }}</div>
        </van-col>
      </van-row>
      <div class="mt10">
        <van-row justify="space-between" align="center">
          <van-col>
            <span class="f14">{{ userInfoData.userInfo.userName }}</span>
          </van-col>
          <van-col>
            <span class="f14">{{ userInfoData.userInfo.deptName }}</span>
          </van-col>
        </van-row>
      </div>
      <div class="mt10">
        <van-row justify="space-between" align="center">
          <van-col>
            <span class="f14 bold">报销类型</span>
          </van-col>
          <van-col>
            <van-radio-group v-model="checked" direction="horizontal" @change="reimburseType">
              <van-radio name="1">冲借款</van-radio>
              <van-radio name="2">日常</van-radio>
            </van-radio-group>
          </van-col>
        </van-row>
      </div>
    </div>
    <!--  选择费用  -->
    <div class="formCard">
      <van-row justify="space-between" align="center" @click="openBill">
        <van-col>
          <div class="f14 bold">费用信息 <span class="red bold">*</span></div>
        </van-col>
        <van-col>
          <van-icon name="add-o"/>
        </van-col>
      </van-row>
      <div v-for="(item,index) in accountBookList">
        <div class="accountBookCard">
          <div>
            <van-icon class="closeIcon" size="20" :name="closeIcon" @click="deleteAccountBook(item,index)"/>
            <div class="ml10  mb5">
              <div class="pt10">
                <van-row justify="space-between" align="center">
                  <van-col class="bold">{{ item.reason }}</van-col>
                  <van-col class="mr10 bold orange">¥ {{ item.money }}</van-col>
                </van-row>
              </div>
              <div class="f14 mt5">
                <van-row align="center">
                  <van-col>
                    <van-row align="center">
                      <van-icon :name="vectorIcon"/>
                      <span class="ml5"> {{ item.expenseTypeName }}</span>
                    </van-row>
                  </van-col>
                  <van-col offset="1">
                    <van-row align="center">
                      <van-icon :name="timeIcon"/>
                      <span class="ml5">{{ item.createTime }}</span>
                    </van-row>
                  </van-col>
                </van-row>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!--  选择发票  -->
    <div class="formCard">
      <van-row justify="space-between" align="center" @click="openInvoice">
        <van-col>
          <div class="f14 bold">发票信息 <span class="red bold">*</span></div>
        </van-col>
        <van-col>
          <van-icon name="add-o"/>
        </van-col>
      </van-row>
      <div v-for="(item,index) in invoiceList">
        <div class="invoiceInfoCard">
          <div class="invoiceCard">
            <van-row justify="space-between" align="center">
              <van-col class="title f16" style="line-height: 25px;" span="15" @click="getInvoiceDetail(item)">
                <van-row align="center">
                  <van-col span="18" class="title f16" style="line-height: 25px;">
                    <span class="f14 title">{{ item.invName }}</span>
                    <van-icon class="ml5" :name="xialaIcon" size="10"/>
                  </van-col>
                </van-row>
              </van-col>
              <van-col>
                <van-row align="center">
                  <span class="font_blue bold">¥{{ item.zje }}</span>
                  <div class="ml5">
                    <van-icon :name="invoiceCloseIcon" size="20" @click="deleteInvoice(item, index)"/>
                  </div>
                </van-row>
              </van-col>
            </van-row>
          </div>
          <div class="p10" v-if="item.show">

            <van-row justify="space-between" v-if="item.invName=='火车票'||item.invName=='飞机票'">
              <van-col>
                <span class="f14 grayTxt">出发地点</span>
              </van-col>
              <van-col>
                <span>{{ dataInfo.cfdd }}</span>
              </van-col>
            </van-row>
            <van-row justify="space-between" v-if="item.invName=='火车票'||item.invName=='飞机票'">
              <van-col>
                <span class="f14 graytxt">出发时间</span>
              </van-col>
              <van-col>{{ dataInfo.kprq }}</van-col>
            </van-row>
            <van-row justify="space-between" v-if="item.invName=='火车票'||item.invName=='飞机票'">
              <van-col>
                <span class="f14 graytxt">到达地点</span>
              </van-col>
              <van-col>{{ dataInfo.dddd }}</van-col>
            </van-row>
            <van-row justify="space-between">
              <van-col>
                <span class="f14 graytxt">发票代码</span>
              </van-col>
              <van-col>{{ item.fpdm }}</van-col>
            </van-row>
            <van-row justify="space-between" v-if="item.invName!='火车票'&&item.invName!='飞机票'">
              <van-col>
                <span class="f14 graytxt">发票号码</span>
              </van-col>
              <van-col>{{ item.fphm }}</van-col>
            </van-row>
            <van-row justify="space-between">
              <van-col>
                <span class="f14 graytxt">金额</span>
              </van-col>
              <van-col>¥{{ item.zje }}</van-col>
            </van-row>
            <!--       <van-row justify="space-between">-->
            <!--         <van-col class="graytxt">发生日期</van-col>-->
            <!--         <van-col class="graytxt">2022-07-08</van-col>-->
            <!--       </van-row>-->
          </div>
        </div>
      </div>
    </div>
    <!--  选择成本中心  -->
    <div class="formCard">
      <van-row>
        <van-col class="bold">成本中心 <span class="red" v-if="costCenterRoot=='0'">*</span></van-col>
      </van-row>
      <div class="mt5" @click="selectCostCenterOpen">
        <van-row justify="space-between" align="center">
          <van-col>
            <span class="graytxt f14">  {{ costCenterTitle }} </span>
          </van-col>
          <van-col>
            <van-icon name="arrow"/>
          </van-col>
        </van-row>
      </div>
    </div>
    <!--  选择部门领导审批流程  -->
    <div class="formCard" v-if="reimburseRoot=='0'">
      <van-row justify="space-between" align="center">
        <van-col class="bold">部门领导审批流程 <span class="red">*</span></van-col>
        <van-col class="mr10" v-if="flowName!=''" @click="seeApproveInfoPicker=true">
          <span class="f12">查看审批流程信息</span>
          <van-icon name="info-o" class="ml5" size="15"/>
        </van-col>
      </van-row>
      <div class="mt5" @click="selectApprovePicker=true">
        <van-row justify="space-between" align="center">
          <van-col>
            <span class="graytxt f14">  {{ departFlowId }} </span>
          </van-col>
          <van-col>
            <van-icon name="arrow"/>
          </van-col>
        </van-row>
      </div>
      <van-popup v-model:show="selectApprovePicker" position="bottom">
        <van-picker
            :columns="approveList"
            @confirm="onApproveConfirm"
            :columns-field-names="customApproveFieldName"
            @cancel="selectApprovePicker = false"
        />
      </van-popup>
    </div>
    <!--  选择申请单  -->
    <div class="formCard" v-if="approvalRoot=='0'">
      <van-row>
        <van-col class="bold">选择申请单</van-col>
      </van-row>
      <div class="mt5" @click="selectApplyPicker=true">
        <van-row justify="space-between" align="center">
          <van-col>
            <span class="graytxt f14">  {{ applyTitle }} </span>
          </van-col>
          <van-col>
            <van-icon name="arrow"/>
          </van-col>
        </van-row>
      </div>
    </div>
    <!--  选择成本中心审批流程  -->
    <div class="formCard" v-if="reimburseRoot=='1'&& costCenterTitle!='请选择成本中心'">
      <van-row justify="space-between" align="center">
        <van-col class="bold">成本中心审批流程 <span class="red">*</span></van-col>
        <van-col class="mr10" v-if="flowName!=''" @click="seeApproveInfoPicker=true">
          <span class="f12">查看审批流程信息</span>
          <van-icon name="info-o" class="ml5" size="15"/>
        </van-col>
      </van-row>
      <div class="mt5" @click="selectCostCenterFlowPicker=true">
        <van-row justify="space-between" align="center">
          <van-col>
            <span class="graytxt f14">  {{ costCenterFlowId }} </span>
          </van-col>
          <van-col>
            <van-icon name="arrow"/>
          </van-col>
        </van-row>
      </div>
      <van-popup v-model:show="selectCostCenterFlowPicker" position="bottom">
        <van-picker
            :columns="flowList"
            @confirm="onConfirm"
            :columns-field-names="customFieldName"
            @cancel="selectCostCenterFlowPicker = false"
        />
      </van-popup>
    </div>
    <!--  选择借款单  -->
    <div class="formCard" v-if="checked=='1'">
      <van-row>
        <van-col class="bold">选择借款单<span class="red">*</span></van-col>
      </van-row>
      <div class="mt5" @click="selectLoanPicker=true">
        <van-row justify="space-between" align="center">
          <van-col>
            <span class="graytxt f14">  {{ loanTitle }} </span>
          </van-col>
          <van-col>
            <van-icon name="arrow"/>
          </van-col>
        </van-row>
      </div>
    </div>
    <van-form>
      <van-field label-align="top"
                 :rules="[{ required: true, message: '请输入报销总额' },{ pattern, message: '金额格式输入错误请重新输入' }]">
        <template #label>
          <span class="bold f16">报销总额</span>
          <span class="graytxt f12 ml5">（根据发票金额进行自动计算） <span class="red">*</span></span>
        </template>
        <template #input>
          <span>¥</span>
          <input v-model="bxAmount" type="number" style="border: none" placeholder="0.00"/>
        </template>
      </van-field>
      <van-field
          v-model="applyReason"
          placeholder="请输入报销事由"
          label-align="top"
          :rules="[{ required: true, message: '请填写报销事由' }]"
      >
        <template #label>
          <span class="bold f16">报销事由</span>
          <span class="red">*</span>
        </template>
      </van-field>
      <van-field
          v-model="bz"
          placeholder="请输入备注"
          label-align="top"
      >
        <template #label>
          <span class="bold f16">备注</span>
        </template>
      </van-field>
      <!--  选择收款信息  -->
      <div class="formCard">
        <van-row>
          <van-col class="bold">收款信息（账户信息） <span class="red" v-if="costCenterRoot=='0'">*</span></van-col>
        </van-row>
        <div class="mt5" @click="bankSelectStatus=true">
          <van-row justify="space-between" align="center">
            <van-col>
              <span class="graytxt f14">{{ payeeTitle }}</span>
            </van-col>
            <van-col>
              <van-icon name="arrow"/>
            </van-col>
          </van-row>
        </div>
      </div>
      <div style="height: 60px"></div>
      <div class="bottomGroup">
        <van-row justify="space-between" align="center">
          <van-col span="4" v-if="saveInfoKey">
            <van-row justify="center" align="center">
              <div class="tc mt5" @click="delDailyBx">
                <van-icon :name="DeleteDomIcon" size="20" class="tc"/>
                <div class="red f12">删除草稿</div>
              </div>
            </van-row>
          </van-col>
          <van-col span="4">
            <van-row justify="center" align="center">
              <div class="tc mt5" @click="editReimburseSave">
                <van-icon :name="SaveDomIcon" size="20" class="tc"/>
                <div class="graytxt f12">保存草稿</div>
              </div>
            </van-row>
          </van-col>
          <van-col :span="saveInfoKey?16:20">
            <div style="padding: 10px 5px">
              <van-button type="primary" size="large" @click="submitVerification">提交</van-button>
            </div>
          </van-col>
        </van-row>
      </div>
    </van-form>
    <!-- 选择借款单   -->
    <van-popup v-model:show="selectLoanPicker" style="width:100vw;height:100vh" position="bottom">
      <select-loan-picker @select-over="selectLoan" ref="selectLoanOver" :data="loanInfo"/>
    </van-popup>
    <!--  选择申请单  -->
    <van-popup v-model:show="selectApplyPicker" style="width:100vw;height:100vh" position="bottom">
      <select-apply-picker @select-over="selectApply" ref="selectApplyOver" :data="applyInfo"/>
    </van-popup>
    <!--  选择费用-->
    <van-popup v-model:show="selectAccountBookPicker" style="width:100vw;height:100vh" position="bottom">
      <select-cost-picker @select-over="selectAccountBook" ref="selectCostOver" :select-list="accountBookList"
      />
    </van-popup>
    <!--  选择发票-->
    <van-popup v-model:show="selectInvoicePicker" style="width:100vw;height:100vh" position="bottom">
      <select-invoice-picker @select-over="selectInvoice" ref="selectInvoiceOver" :select-list="invoiceList"
      />
    </van-popup>
    <!-- 选择成本中心 -->
    <van-popup v-model:show="selectCostCenterPicker" style="width:100vw; height:100vh" position="bottom">
      <select-cost-center-picker @select-over="selectCostCenter" ref="selectCostCenterOver" :data="costCenterData"/>
    </van-popup>
    <!-- 加载中  -->
    <div v-show="loading">
      <loading-component :title="'提交中'"/>
    </div>
    <!-- 审批显示 -->
    <van-popup v-model:show="seeApproveInfoPicker" style="width:100vw;height:100vh" position="bottom">
      <approval-info @select-over="seeApproveInfoPicker=false" :approve-info-list="getApproveInfoList"
                     :code="getApproveNum"/>
    </van-popup>
    <!-- 选择收款账户 -->
    <van-popup v-model:show="bankSelectStatus" style="width:100vw;height:100vh" position="bottom">
      <BankCardSelector :isComponentsStatus="true" :selectBankMesssage="selectBankMesssage"
                        @selectBankCardInformation="selectBankCardMth"/>
    </van-popup>
  </div>
</template>

<script lang="ts" setup>
import Navbar from '../../../components/navBar/index.vue';
import SelectCostPicker from '../../../components/selectComponent/selectCostPicker/index.vue';
import SelectInvoicePicker from '../../../components/selectComponent/selectInvoicePicker/index.vue';
import SelectCostCenterPicker from '../../../components/selectComponent/selectCostCenterPicker/index.vue';
import SelectApplyPicker from '../../../components/selectComponent/selectApplyPicker/index.vue';
import LoadingComponent from '../../../components/loading/index.vue' // 加载组件引入
import ApprovalInfo from '../../application/form/components/approveInfo/index.vue'
import vectorIcon from '../../../assets/icon/Vector.png';
import timeIcon from '../../../assets/icon/baoxiao_icon_02.png';
import closeIcon from '../../../assets/img/close_icon.png';
import invoiceCloseIcon from '../../../assets/icon/close_invoice.png';
import SaveDomIcon from '../../../assets/icon/saveDom.png';
import DeleteDomIcon from '../../../assets/icon/deleteDom.png';
import xialaIcon from '../../../assets/icon/xiala.png';
import SelectLoanPicker from '../../../components/selectComponent/selectLoanPicker/index.vue'
import BankCardSelector from '../../manageAccounts/index.vue'
import {ref, inject, onActivated, watch, watchEffect, onDeactivated} from "vue";
import {applyNum} from "./ts/applyNo";
import moment from "moment/moment";
import {onBeforeRouteLeave, useRouter} from 'vue-router';
import {getApprovalFlow, getApprovalFlowHistory, getApprovalFlowList} from "../../../services/current";
import {
  getDefaultPay,
  editSave,
  checkMaxAmount,
  sendDailyBx,
  getDailyBxDetil,
  deleteEdit
} from "../../../services/reimburse";
import {showFailToast} from "vant/lib/toast/function-call";
import {showConfirmDialog, showSuccessToast} from "vant";
import {updateAgentState} from "../../../services/myDelegation";
import {getInvoiceWbById} from "../../../services/invoice";

const checked = ref('2'); // 选择报销类型 1：冲借款 2：日常报销
let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
let applyNo = ref(''); // 申请单号
const router = useRouter(); // 获得路由
const selectCostOver = ref(); // 获取组件实例
const selectInvoiceOver = ref(); // 获取组件实例
const payeeTitle = ref('选择银行卡号'); // 银行卡号
const applyTitle = ref('请选择前置申请单'); // 选择前置申请单
const applyReason = ref(''); // 报销事由
const bxAmount = ref(''); // 报销总额
const bz = ref(''); // 备注
const applyInfo = ref({}); // 申请单信息
const loanInfo = ref({}); // 借款单信息
const loading = ref(false); // loading组件
const saveInfoKey = ref(false); // 控制是否展示删除按钮
const selectBankMesssage = ref() //
// 提交参数
let upData = ref({
  'applyNo': "",
  'approveNo': '',
  "toAccount": "test_auvgo8899",
  "sysType": "用友",
  "applyReason": "",
  "empCode": "",
  "empName": "",
  "deptId": "",
  "deptName": "",
  "corpCode": "",
  "bz": "",
  "bxAmount": "0.00",
  "billDtoList": [],
  "dailyBxDetailList": [],
  "payeeinfo": {
    "empcode": "",
    "payeeName": "",
    "bankcardnum": "",
    "bankaccountname": "",
    "region": "",
    "city": "",
    "accountnetwork": "",
    "corpCode": "",
    "type": ""
  },
  'flowId': '',
  'budGetingId': '',
  'costType': '',
  'loanApplyNo': '',
  'xmbh': "",
  'xmmc': "",
  'xmye': "",
  "approvalFlowType": "costCenterFlow", //审批（成本中心）
});
const pattern = /^[0-9]+(\.[0-9]{1,2}|)$/; // 保留两位小数的金额校验
//日常报销是否关联前置申请单 （0：关联，1：不关联）
let approvalRoot = ref('');
// 控制日常报销成本中心是否必填 (0：必填 1：不必填)
let costCenterRoot = ref('');
// 日常报销审批方式  (0：部门领导审批 1：成本中心审批)
let reimburseRoot = ref('');
// 当前时间
let nowDate = moment().format("yyyy-MM-DD");
// 控制账本显隐
let selectAccountBookPicker = ref(false);
// 控制发票显隐
let selectInvoicePicker = ref(false);
// 控制成本中心显隐
let selectCostCenterPicker = ref(false);
// 控制审批详情显隐
let seeApproveInfoPicker = ref(false);
// 控制审批选择栏显隐
let selectApprovePicker = ref(false);
// 控制成本中心审批流显隐
let selectCostCenterFlowPicker = ref(false);
// 选择申请单组件显隐
let selectApplyPicker = ref(false);
// 选择借款单组件显隐
let selectLoanPicker = ref(false);
// 选择账户信息
let bankSelectStatus = ref(false);
// 用于存放账本数组
let accountBookList = ref([]);
// 用于选择并放回去的账本数组
let selectAccountBookList = ref([]);
// 用于存放发票数组
let invoiceList = ref([]);
// 用于选择并放回去的发票数组
let selectInvoiceList = ref([]);
// 成本中心
let costCenterTitle = ref('请选择成本中心');
// 部门领导审批
let departFlowId = ref('请选择部门领导审批流');
// 成本中心审批
let costCenterFlowId = ref('请选择成本中心审批流');
// 审批详情数组
let getApproveInfoList = ref([]);
// 审批抄送数组
let getapproveCCInfo = ref([]);
// 审批进度
let getApproveNum = ref('');
// 成本中心传递内容
let costCenterData = ref({});
// 选择部门审批数组
let approveList = ref([]);
// 审批流程名称
let flowName = ref('');
// 审批流程id
let flowId = ref('');
// 成本中心数组
let flowList = ref([]);
// 借款单选择文案
let loanTitle = ref('请选择借款申请单');
// 用于规范选择框类型(成本中心审批流)
const customFieldName = {
  text: 'flowname',
  value: 'flowid',
};
// 用于规范选择框类型(部门领导审批流)
const customApproveFieldName = {
  text: 'flow_name',
  value: 'flow_id',
};
// 选择部门领导审批流
const onApproveConfirm = (value: any) => {
  flowName.value = value.selectedOptions[0].flow_name;
  flowId.value = value.selectedOptions[0].flow_id;
  departFlowId.value = `${flowName.value}`;
  upData.value.flowId = flowId.value;
  getFlowHistory(flowId.value);
  selectApprovePicker.value = false;
};
// 点击打开费用数组
const openBill = () => {
  selectAccountBookPicker.value = true;
  setTimeout(() => selectCostOver.value.refresh(), 500)
}
// 点击打开发票数组
const openInvoice = () => {
  selectInvoicePicker.value = true
  setTimeout(() => selectInvoiceOver.value.refresh(), 500)
}
// 点击打开成本中心选择器
const selectCostCenterOpen = () => {
  costCenterTitle.value = costCenterTitle.value;
  selectCostCenterPicker.value = true;
}
// 用于回显账户信息
const selectBankCardMth = (e: any) => {
  if (e) {
    selectBankMesssage.value = e
    let payeeinfo = {
      'empcode': e.empcode,
      "payeeName": e.payeeName,
      "bankcardnum": e.bankcardnum,
      "bankaccountname": e.bankaccountname,
      "region": e.region,
      "city": e.city,
      "accountnetwork": e.accountnetwork,
      "corpCode": e.corpCode,
      "id": e.id,
      "type": e.type,
    };
    payeeTitle.value = `${e.bankcardnum} (${e.bankaccountname})`
    upData.value.payeeinfo = payeeinfo;
  }
  bankSelectStatus.value = false
}
// 选择成本中心
const selectCostCenter = (data: any) => {
  if (data.xmmc) {
    costCenterData.value = data;
    costCenterTitle.value = `${data.budGetingYear}-${data.xmmc}   ¥ ${data.xmye}`;
    upData.value.xmbh = data.xmbh;
    upData.value.xmmc = data.xmmc;
    upData.value.xmye = data.xmye;
    upData.value.costType = data.costType;
    upData.value.budGetingId = data.budGetingId;
    GetCostCenterFlow(data.xmbh);
  }
  selectCostCenterPicker.value = false;
}
// 根据获取到成本中心选择可选择的审批流
const GetCostCenterFlow = async (key: any) => {
  let payLoad = {
    "applyTypeId": "",
    "xmbh": key,
    "deptId": userInfoData.userInfo.deptId,
    "corpCode": userInfoData.userInfo.corpCode,
    "approvalFlowType": "costCenterFlow",
    "equipmentType": ""
  }
  let result: any = await getApprovalFlow(payLoad);
  if (result.data[0].length != 0) {
    flowList.value = result.data;
    saveInfoApprove();
  } else {
    showFailToast('该成本中心下没有相关审批流，请联系管理员');
  }
}
// 找到保存的审批流程名称和相关信息
const saveInfoApprove = async () => {
  let flow = flowList.value;
  flow = flow[0];
  let flowData: any = flow.find((e: any) => {
    return e.flowid == flowId.value;
  })
  if (flowData != undefined) {
    flowName.value = flowData.flowname;
    flowId.value = flowData.flowid;
    costCenterFlowId.value = `${flowName.value}`;
    getFlowHistory(flowId.value);
  }
};

// 选择成本中心审批流
const onConfirm = (value: any) => {
  flowName.value = value.selectedOptions[0].flowname;
  flowId.value = value.selectedOptions[0].flowid;
  costCenterFlowId.value = flowName.value;
  getFlowHistory(flowId.value);
  selectCostCenterFlowPicker.value = false;
}
// 删除发票
const deleteInvoice = (item: any, index: any) => {
  invoiceList.value.splice(index, 1);
  let list = invoiceList.value;
  let aMount = 0;
  for (let i in list) {
    let da: { zje: string } = list[i];
    aMount = aMount + parseFloat(da.zje);
  }
  bxAmount.value = aMount.toString();
}
const unwatch = watchEffect(() => {
})
// 用于监听跳转页面的数据
watch(
    () => router.currentRoute.value.name,
    (value, oldValue) => {
      // 如果来自于账本 刷新账本列表
      if (oldValue == 'book') {
        console.log(selectCostOver.value)
        if (selectCostOver.value != undefined) {
          selectCostOver.value.payload.page = 1;
          selectCostOver.value.refresh()
        }
        // 如果来自于发票 刷新发票列表
      } else if (oldValue == 'addList') {
        if (selectInvoiceOver.value != undefined) {
          selectInvoiceOver.value.payload.page = '1';
          selectInvoiceOver.value.refresh()
        }
      }
      // 如果来自于报销列表 将原来有数据重制
      else if (oldValue == 'reimburse') {
        resetReimburse();
      } else {
        unwatch();
      }
    },
    {immediate: true}
);
// 重制报销单
const resetReimburse = () => {
  upData.value = {
    'applyNo': "",
    'approveNo': '',
    "toAccount": "test_auvgo8899",
    "sysType": "用友",
    "applyReason": "",
    "empCode": "",
    "empName": "",
    "deptId": "",
    "deptName": "",
    "corpCode": "",
    "bz": "",
    "bxAmount": "0.00",
    "billDtoList": [],
    "dailyBxDetailList": [],
    "payeeinfo": {
      "empcode": "",
      "payeeName": "",
      "bankcardnum": "",
      "bankaccountname": "",
      "region": "",
      "city": "",
      "accountnetwork": "",
      "corpCode": "",
      "type": ""
    },
    'flowId': '',
    'budGetingId': '',
    'costType': '',
    'loanApplyNo': '',
    'xmbh': "",
    'xmmc': "",
    'xmye': "",
    "approvalFlowType": "costCenterFlow", //审批（成本中心）
  };
  bxAmount.value = '';
  applyNo.value = applyNum(checked.value);
  accountBookList.value = [];
  invoiceList.value = [];
  costCenterData.value = [];
  applyReason.value = '';
  costCenterFlowId.value = '请选择成本中心审批流';
  costCenterTitle.value = '请选择成本中心';
  departFlowId.value = '请选择部门领导审批流';
  loanTitle.value = '请选择借款单';
  flowName.value = '';
  flowId.value = '';
  flowList.value = [];
  // defaultPayeeCode();
}
// 删除费用
const deleteAccountBook = (item: any, index: any) => {
  accountBookList.value.splice(index, 1);
};
// 接收传递来的借款单
const selectLoan = (item: any) => {
  if (item.length != 0) {
    loanTitle.value = item.applyreason;
    loanInfo.value = item;
    upData.value.loanApplyNo = item.applyno;
  }
  selectLoanPicker.value = false;
}
// 接收传递来的申请单
const selectApply = (item: any) => {
  if (item.length != 0) {
    applyTitle.value = item.approveReason;
    upData.value.approveNo = item.applyNo;
  }
  selectApplyPicker.value = false;
}
// 接收传递来的费用
const selectAccountBook = (list: any) => {
  if (list.length != 0) {
    accountBookList.value = [];
    accountBookList.value = list;
  }
  selectAccountBookPicker.value = false;
}
// 接收传递来的发票
const selectInvoice = (list: any) => {
  if (list.length != 0) {
    invoiceList.value = [];
    invoiceList.value = list;
    let aMount = 0;
    for (let i in list) {
      aMount = aMount + parseFloat(list[i].zje)
    }
    bxAmount.value = aMount.toString();
  }
  selectInvoicePicker.value = false;
}
// 切换类型变换对应的报销单号
const reimburseType = () => {
  applyNo.value = applyNum(checked.value);
}
// 根据系统设定的权限控制审批模式处理
const ywConfigRoot = () => {
  // 判断
  let rootList = userInfoData.functionConfig;
  for (var item in rootList) {
    if (rootList[item].ywTypeID == 'fk002') {
      //日常报销是否关联前置申请单 （0：关联，1：不关联）
      approvalRoot.value = rootList[item].configValue;
    }
    if (rootList[item].ywTypeID == 'fk001') {
      // 控制日常报销成本中心是否必填 (0：必填 1：不必填)
      costCenterRoot.value = rootList[item].configValue;
    }
    if (rootList[item].ywTypeID == 'fk004') {
      // 日常报销审批方式  (0：部门领导审批 1：成本中心审批)
      reimburseRoot.value = rootList[item].configValue;
    }
  }
  if (reimburseRoot.value != '0') {
    deptApproveFlow()
  }
}
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
  getApproveInfoList.value = flowHistory.data.toUserList;
  getapproveCCInfo.value = flowHistory.data.ccUserList;
  getApproveNum.value = flowHistory.data.isVote;
  const object: any = {};
  for (let i = 0; i <= getApproveInfoList.value.length - 1; i++) {
    let no = getApproveInfoList.value[i]['no'];
    let isAllPass = getApproveInfoList.value[i]['isAllPass'];
    if (!object[no]) {
      object[no] = {
        no,
        isAllPass,
        children: []
      }
    }
    if (no === object[no].no) {
      object[no].children.push(getApproveInfoList.value[i])
    }
  }
  getApproveInfoList.value = [];
  Object.values(object).forEach((item) => {
    return getApproveInfoList.value.push(item as never);
  })
  let userName: any = [];
  flowHistory.data.ccUserList.forEach((item: any) => userName.push(`${item.ccUserName}(${item.toRole})`));
}
// 获取用户默认收款账号
const defaultPayeeCode = async () => {
  let payload = {
    "corpCode": userInfoData.userInfo.corpCode,
    "empCode": userInfoData.userInfo.empCode,
    "equipmentType": ""
  };
  let result: any = await getDefaultPay(payload);
  selectBankMesssage.value = result.data;
  upData.value.payeeinfo = selectBankMesssage.value;
  payeeTitle.value = `${result.data.bankcardnum} (${result.data.bankaccountname})`
}
// 在处于部门审批流程的时候
const deptApproveFlow = async () => {
  let payload = {
    "corpCode": userInfoData.userInfo.corpCode,
    'deptId': userInfoData.userInfo.deptId,
    "typeId": "FK_RCBX_SQ",
    'isOpen': '1',
    'empCode': userInfoData.userInfo.empCode,
    'equipmentType': "pc"
  }
  let result: any = await getApprovalFlowList(payload);
  approveList.value = result.data;
}
// 点击保存日常报销
const editReimburseSave = async () => {
  // 赋值报销事由
  upData.value.applyReason = applyReason.value;
  upData.value.applyNo = applyNo.value;
  if (upData.value.applyReason == '') {
    showSuccessToast('报销事由不能为空');
    return;
  }
  // 添加账本费用至提交提交费用数组
  let billList = accountBookList.value;
  for (var item in billList) {
    let billInfo: {
      billId: any,
      empCode: any,
      expenseType: any,
      expenseTypeName: any,
      money: any,
      expenseOwner: any
    } = billList[item];
    upData.value.billDtoList.push({
      'billId': billInfo.billId,
      'empCode': billInfo.empCode,
      'expenseType': billInfo.expenseType,
      'expenseTypeName': billInfo.expenseTypeName,
      'money': billInfo.money,
      'expenseOwner': billInfo.expenseOwner,
    } as never);
  }
  // 添加账本费用至提交提交费用数组
  let invoiceBxList = invoiceList.value;
  // 添加发票数据至提交发票数组
  for (var item in invoiceBxList) {
    let invoiceInfo: {
      zje: any,
      invName: any,
      dzfph: any,
      fpdm: any,
      fphm: any,
      fpygbh: any,
      fjid: any
    } = invoiceBxList[item];
    upData.value.dailyBxDetailList.push({
      'spendAmount': invoiceInfo.zje,
      'spendType': invoiceInfo.invName,
      'dzfph': invoiceInfo.dzfph,
      'fpdm': invoiceInfo.fpdm,
      'fphm': invoiceInfo.fphm,
      'fpygbh': invoiceInfo.fpygbh,
      'fjid': invoiceInfo.fjid,
    } as never);
  }
  upData.value.payeeinfo = selectBankMesssage.value;
  // 个人信息
  upData.value.deptId = userInfoData.userInfo.deptId;
  upData.value.deptName = userInfoData.userInfo.deptName;
  upData.value.empCode = userInfoData.userInfo.empCode;
  upData.value.empName = userInfoData.userInfo.userName;
  upData.value.bxAmount = bxAmount.value;
  upData.value.flowId = flowId.value;
  loading.value = true;
  let result: any = await editSave(upData.value);
  if (result.code == 200) {
    loading.value = false;
    resetReimburse();
    showSuccessToast(result.data);
    router.back();
  } else {
    loading.value = false;
    showFailToast(result.data.message);
  }
}
// 提交校验
const submitVerification = () => {
  // 判断费用信息是否为空
  if (accountBookList.value.length == 0) {
    showFailToast('账本费用信息不能为空');
    return false;
  }
  // 判断发票数组信息是否为空
  if (invoiceList.value.length == 0) {
    showFailToast('发票信息不能为空');
    return false;
  }
  // 判断成本中心
  if (costCenterRoot.value == '0' && upData.value.xmbh == '') {
    showFailToast('成本中心不能为空');
    return false;
  }
  // 判断审批流是否为空
  if (flowId.value == '') {
    if (reimburseRoot.value == '0') {
      showFailToast('成本中心审批流不能为空');
      return false;
    } else {
      showFailToast('部门领导审批流不能为空');
      return false;
    }
  }
  // 判断报销金额
  if (bxAmount.value == '0.00') {
    showFailToast('报销总额不能为0');
    return false;
  }
  // 报销事由是否为空
  if (applyReason.value == '') {
    showFailToast('报销金额不能为空');
    return false;
  }
  //  收款账户是否为空
  if (upData.value.payeeinfo.bankcardnum == '') {
    showFailToast('收款人账户信息不能为空');
    return false;
  }
  checkAmount()
}
// 日常报销的校验金额控制
const checkAmount = async () => {
  let resData = {
    bxType: 'daily',
    corpCode: userInfoData.userInfo.corpCode,
    deptId: userInfoData.userInfo.deptId,
    empCode: userInfoData.userInfo.empCode,
    money: bxAmount.value
  }
  let result: any = await checkMaxAmount(resData);
  if (result.data.isControl == '1') {
    switch (result.data.tips) {
      case"0": //提示信息
        showFailToast(result.message)
        setTimeout(() => {
          onSubmit()
        }, 200)
        break;
      case"1": //禁止提交
        break;

      case"2": //不控制
        break;
    }

  } else {
    onSubmit()
  }
}
// 保存的数据回显
const getReimburseDetailInfo = () => {
  let reimburseDetail: any = sessionStorage.getItem('reimburseDetail');
  reimburseDetail = JSON.parse(reimburseDetail);
  //  查看是否有申请单号
  if (reimburseDetail.approveNo != '') {
    saveInfoKey.value = true;
    let approveNo = reimburseDetail.approveNo;
    applyNo.value = approveNo;
    ywConfigRoot();
    applyType(approveNo);
    getDailyBx(approveNo);
  } else {
    defaultPayeeCode();
    ywConfigRoot();
    accountBookList.value = [];
    invoiceList.value = [];
    let billList: any = sessionStorage.getItem('billList');
    let invoiceListStores: any = sessionStorage.getItem('invoiceList');
    billList = JSON.parse(billList);
    invoiceListStores = JSON.parse(invoiceListStores);
    if (invoiceListStores.length != 0) {
      invoiceList.value = invoiceListStores;
      let aMount = 0;
      for (let i in invoiceListStores) {
        aMount = aMount + parseFloat(invoiceListStores[i].zje)
      }
      bxAmount.value = aMount.toString();
    }
    if (billList.billId != '') {
      accountBookList.value.push(billList as never);
    }
    applyNo.value = applyNum(checked.value);
    saveInfoKey.value = false
  }
}
// 从账本中进入提交报销申请
// 根据申请单号以及类型判断该报销单类型
const applyType = (item: any) => {
  var str = item.substring(0, 2);
  if (str == 'HK') {
    checked.value = '1'
  } else {
    checked.value = '2'
  }
}
// 请求参数
const resData = {
  "dzfph": "",
  "corpCode": userInfoData.userInfo.corpCode,
  "empCode": userInfoData.userInfo.empCode,
  "equipmentType": ""
}
// 接收接口返回的发票详情
const dataInfo = ref({fpdm: ''});
// 获取发票详情
const getInvoiceDetail = async (item: any) => {
  resData.dzfph = item.dzfph;
  let result: any = await getInvoiceWbById(resData);
  dataInfo.value = result.data.data[0]
  item.show = !item.show
};
// 删除报销单
const delDailyBx = async (item: any) => {
  showConfirmDialog({
    title: '您是否删除报销单'
  }).then(async () => {
    let data = {"applyNo": applyNo.value, "corpCode": userInfoData.userInfo.corpCode, "equipmentType": ""}
    let result: any = await deleteEdit(data);
    if (result.code == 200) {
      showSuccessToast(result.data);
      resetReimburse();
      router.back();
    } else {
      loading.value = false;
      showFailToast(result.data.message);
    }
  }).catch(() => {
    item.status = !item.status
  });

};
// 获取保存的申请单详情
//获取申请单详情
const getDailyBx = async (item: any) => {
  let payload = {
    "empCode": userInfoData.userInfo.empCode,
    "corpCode": userInfoData.userInfo.corpCode,
    "applyNo": item,
    "equipmentType": ""
  }; // 提交信息
  let result: any = await getDailyBxDetil(payload);
  ;
  // 添加账本费用至提交提交费用数组
  let billList = result.data.billList;
  for (var key in billList) {
    let billInfo: any = billList[key];
    accountBookList.value.push({
      'billId': billInfo.billId,
      'empCode': billInfo.empCode,
      'expenseType': billInfo.expenseType,
      'expenseTypeName': billInfo.expenseTypeName,
      'reason': billInfo.reason,
      'createTime': billInfo.createTime,
      'money': billInfo.money,
      'expenseOwner': billInfo.expenseOwner,
    } as never);
  }
  // 添加账本费用至提交提交费用数组
  let invoiceBxList = result.data.spendDetailList;
  // 添加发票数据至提交发票数组
  for (var info in invoiceBxList) {
    let invoiceInfo: any = invoiceBxList[info];
    invoiceList.value.push({
      'zje': invoiceInfo.spendAmount,
      'invName': invoiceInfo.spendType,
      'dzfph': invoiceInfo.dzfph,
      'fpdm': invoiceInfo.fpdm,
      'fphm': invoiceInfo.fphm,
      'fpygbh': invoiceInfo.fpygbh,
      'fjid': invoiceInfo.fjid,
    } as never);
  }
  // 成本中心赋值
  if (result.data.costcenter) {
    costCenterData.value = result.data.costcenter;
    upData.value.xmbh = result.data.costcenter.xmbh;
    upData.value.xmmc = result.data.costcenter.xmmc;
    upData.value.xmye = result.data.costcenter.xmye;
    upData.value.costType = result.data.costcenter.costType;
    upData.value.budGetingId = result.data.costcenter.budGetingId;
    GetCostCenterFlow(result.data.costcenter.xmbh);
    flowName.value = result.data.costcenter.xmmc;
    flowId.value = result.data.dailyDetailDto.flowId;
    costCenterTitle.value = `${result.data.costcenter.budGetingYear}-${result.data.costcenter.xmmc}   ¥ ${result.data.costcenter.xmye}`;
  }
  if (result.data.loanDetail) {
    loanTitle.value = result.data.loanDetail.applyreason;
    upData.value.loanApplyNo = result.data.loanDetail;
    loanInfo.value = result.data.loanDetail;
  }
  upData.value.flowId = flowId.value;
  bxAmount.value = result.data.dailyDetailDto.bxamount;
  applyReason.value = result.data.dailyDetailDto.applyreason;
  bz.value = result.data.dailyDetailDto.bz;
  // 收款账户
  let payeeinfo = {
    'empcode': result.data.dailyDetailDto.empcode,
    "payeeName": result.data.dailyDetailDto.payeeName,
    "bankcardnum": result.data.dailyDetailDto.bankNum,
    "bankaccountname": result.data.dailyDetailDto.bankName,
    "region": result.data.dailyDetailDto.region,
    "city": result.data.dailyDetailDto.city,
    "accountnetwork": result.data.dailyDetailDto.accountNetwork,
    "corpCode": result.data.dailyDetailDto.corpCode,
    "type": result.data.dailyDetailDto.accountType,
    "id": result.data.dailyDetailDto.id,
  };
  selectBankMesssage.value = payeeinfo;
  payeeTitle.value = `${result.data.dailyDetailDto.bankNum} (${result.data.dailyDetailDto.bankName})`
}
// 提交日常报销
const onSubmit = async () => {
  // 赋值报销事由
  upData.value.applyReason = applyReason.value;
  upData.value.applyNo = applyNo.value;
  // 添加账本费用至提交提交费用数组
  let billList = accountBookList.value;
  for (var item in billList) {
    let billInfo: {
      billId: any,
      empCode: any,
      expenseType: any,
      expenseTypeName: any,
      money: any,
      expenseOwner: any
    } = billList[item];
    upData.value.billDtoList.push({
      'billId': billInfo.billId,
      'empCode': billInfo.empCode,
      'expenseType': billInfo.expenseType,
      'expenseTypeName': billInfo.expenseTypeName,
      'money': billInfo.money,
      'expenseOwner': billInfo.expenseOwner,
    } as never);
  }
  let invoiceBxList = invoiceList.value;
  // 添加发票数据至提交发票数组
  for (var item in invoiceBxList) {
    let invoiceInfo: {
      zje: any,
      invName: any,
      dzfph: any,
      fpdm: any,
      fphm: any,
      fpygbh: any,
      fjid: any
    } = invoiceBxList[item];
    upData.value.dailyBxDetailList.push({
      'spendAmount': invoiceInfo.zje,
      'spendType': invoiceInfo.invName,
      'dzfph': invoiceInfo.dzfph,
      'fpdm': invoiceInfo.fpdm,
      'fphm': invoiceInfo.fphm,
      'fpygbh': invoiceInfo.fpygbh,
      'fjid': invoiceInfo.fjid,
    } as never);
  }
  // 添加金额
  upData.value.deptId = userInfoData.userInfo.deptId;
  upData.value.deptName = userInfoData.userInfo.deptName;
  upData.value.empCode = userInfoData.userInfo.empCode;
  upData.value.empName = userInfoData.userInfo.userName;
  upData.value.bxAmount = bxAmount.value;
  upData.value.flowId = flowId.value;
  upData.value.payeeinfo = selectBankMesssage.value;
  loading.value = true;
  let result: any = await sendDailyBx(upData.value);
  if (result.code == 200) {
    loading.value = false;
    showSuccessToast(result.data);
    resetReimburse();
    router.back();
  } else {
    loading.value = false;
    showFailToast(result.data.message);
  }
}
onDeactivated(() => {
  // 在从 DOM 上移除、进入缓存
  // 此处用于关闭监听
  watchEffect(() => {
  })
  sessionStorage.removeItem('reimburseDetail');
  sessionStorage.removeItem('bilList');
  sessionStorage.removeItem('invoiceList');
  // 以及组件卸载时调用
})
// 活动时调用
onActivated(() => {
  getReimburseDetailInfo();
});
</script>

<style scoped>
.invoiceCard {
  margin-top: 5px;
  padding: 10px;
  background: #EEF7FF;
  border-radius: 10px;
}

.invoiceInfoCard {
  margin-left: 5px;
  margin-right: 5px;
  margin-bottom: 5px;
  background: #F7FBFF;
  border-bottom-right-radius: 10px;
  border-bottom-left-radius: 10px;
}

.formCard {
  margin: 10px 5px;
  padding: 10px;
  background: white;
  border-radius: 8px;
}

.accountBookCard {
  margin: 10px;
  border: 1px solid rgba(0, 128, 255, 0.4);
  border-radius: 6px;
}

.invoiceCard {
  padding: 10px;
  margin-top: 10px;
  background: rgba(0, 128, 255, 0.2);
  border-radius: 6px;
}

.bottomGroup {
  width: 100%;
  height: 70px;
  background: #ffffff;
  position: fixed;
  bottom: 0px;
}

.closeIcon {
  float: right;
  top: 0px;
  right: 0px;
}
</style>
