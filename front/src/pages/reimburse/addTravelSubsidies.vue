<template>
    <div>
        <div>
            <NavBar :leftText="'返回'" :title="'差旅补助报销单'" :leftArrow="true" @leftEvent="back()"></NavBar>
        </div>
        <div>
            <van-notice-bar color="#1989fa" :scrollable="false" wrapable background="#ecf9ff" left-icon="info-o" class="f10">
                提交时需要出行明细或者发票信息其中之一即可
            </van-notice-bar>
            <div>
                <van-form @submit="checkMaxAmountMth">
                    <div :style="{height:realHeight,overflowY:'scroll' }">
                        <van-cell-group class="pb10 mt10">
                            <div class="ml10 pt10 f14">
                                <van-row justify="space-between">
                                    <van-col span="17">{{travelNum}}</van-col>
                                    <van-col span="7">{{createDate}}</van-col>
                                </van-row>
                                <van-row justify="space-between" class="pt10">
                                    <van-col span="17">{{userInfoData.userInfo.userName}}</van-col>
                                    <van-col span="7">{{userInfoData.userInfo.deptName}}</van-col>
                                </van-row>
                            </div>
                        </van-cell-group>
                        <van-cell-group class="pb10 mt10">
                            <van-field
                                    v-model="travelTime"
                                    name="travelTime"
                                    label="行程时间"
                                    placeholder="出发时间—结束时间"
                                    :rules="[{ required: true, message: '请选择行程时间' }]"
                                    class="mt10 mb10"
                                    label-width="7.5em"
                                    label-align="top"
                                    required="true"
                                    is-link
                                    readonly
                                    @click="showCalendarMth"
                            />
                            <van-field
                                    v-model="reason"
                                    name="reason"
                                    label="报销事由"
                                    placeholder="请输入报销事由"
                                    :rules="[{ required: true, message: '请输入报销事由' }]"
                                    class="mt10 mb10 "
                                    label-width="7.5em"
                                    label-align="top"
                                    required="true"
                                    maxlength="25"

                            />
                            <van-field
                                    v-model="reimbursementAmount"
                                    name="reimbursementAmount"
                                    label="报销总额(根据补助金额和票据进行自动计算)"
                                    placeholder="0.00"
                                    :rules="[{ required: true, message: '请输入报销总额' }]"
                                    class="mt10 mb10 "
                                    label-width="20em"
                                    maxlength="10"
                                    type="number"
                                    label-align="top"
                                    required="true"
                            />
                        </van-cell-group>
                        <van-field
                                v-model="costCenter"
                                name="costCenter"
                                label="成本中心"
                                placeholder="请选择成本中心"
                                :rules="[{ required: isCost.value, message: '请选择成本中心' }]"
                                class="mt10 mb10 "
                                readonly
                                isLink
                                label-width="7.5em"
                                label-align="top"
                                :required="isCost"
                                @click="selectCostCenterOpen"
                        />
                        <div v-if="!isDeptStatus && isCost && costCenter !=''">
                            <van-field
                                    v-model="costApprovel"
                                    name="costApprovel"
                                    label="审批流"
                                    placeholder="请选择审批流"
                                    :rules="[{ required: true, message: '请选择审批流' }]"
                                    class="mt10 mb10 "
                                    readonly
                                    isLink
                                    label-width="7.5em"
                                    label-align="top"
                                    required="true"
                                    :right-icon="costApprovel==''?'':'info-o'"
                                    @click-right-icon.stop="queryFlowDetail('cost')"
                                    @click="showApproveSelectPopup = true"
                            />
                        </div>
                        <div v-if="isDeptStatus">
                            <van-field
                                    v-model="deptApprovel"
                                    name="deptApprovel"
                                    label="部门审批流"
                                    placeholder="请选择部门审批流"
                                    :rules="[{ required: true, message: '请选择部门审批流' }]"
                                    class="mt10 mb10 "
                                    readonly
                                    isLink
                                    label-width="7.5em"
                                    label-align="top"
                                    required="true"
                                    :right-icon="deptApprovel==''?'':'info-o'"
                                    @click-right-icon.stop="queryFlowDetail('dept')"
                                    @click="showApproveSelectPopupDept = true"
                            />
                        </div>
                        <van-field
                                v-model="paymentInformation"
                                name="paymentInformation"
                                label="收款信息(账户信息)"
                                placeholder="请选择银行卡号"
                                :rules="[{ required: true, message: '请选择银行卡号' }]"
                                class="mt10 mb10"
                                readonly
                                is-link="true"
                                label-width="9.5em"
                                label-align="top"
                                required="true"
                                @click="selectBankAccount"
                        />

                        <div style="background: #FFFFFF;margin-bottom: 10px">
                            <van-row class="pt10 pl20">
                                <van-col span="8" style="font-size: 14px">
                                    <a>出行明细</a>
                                </van-col>
                                <van-col span="14">
                                </van-col>
                                <van-col span="2">
                                    <van-icon name="add-o" @click="showTravelDetailDialog"/>
                                </van-col>
                            </van-row>
                            <van-row class="f14 pl20 pt10">
                                <div v-if="travelDetails.length>0">
                                    <div v-for="item in travelDetails"
                                         :style="{width:realWidth+'px', border:'#4BA3FB solid 1px',borderRadius:'5px',minHeight:'30px',marginBottom: '10px'}">
                                        <div class="pt10 pl20">
                                            <van-row>
                                                <van-col span="2">
                                                    <van-icon v-if="item.type=='H'" :name="jiudian"></van-icon>
                                                    <van-icon v-if="item.type=='F'" :name="feiji"></van-icon>
                                                    <van-icon v-if="item.type=='T'" :name="huoche"></van-icon>
                                                </van-col>
                                                <van-col span="10" class="fc fw600">
                                                    {{item.passengernames}}
                                                </van-col>
                                                <van-col span="3"></van-col>
                                                <van-col span="7" class="money f16" align="right">
                                                    ￥{{item.totalamount}}
                                                </van-col>
                                            </van-row>
                                            <van-row class="pt10 f14" v-if="item.type == 'H'">
                                                {{item.endcityname}}
                                            </van-row>
                                            <van-row class="pt10 f14" v-if="item.type == 'H'">
                                                {{item.startdate}} 至 {{item.enddate}}
                                            </van-row>
                                            <van-row v-if="item.type == 'F'||item.type == 'T'">
                                                <van-steps class="stepsCss" direction="vertical" active-color="#969799"
                                                           :style="{width:realWidth-140+'px'}">
                                                    <van-step>
                                                        <van-row>
                                                            <van-col span="8">{{item.startcityname}}</van-col>
                                                            <van-col span="16" class="fc">{{item.startdate}}</van-col>
                                                        </van-row>
                                                    </van-step>
                                                    <van-step>
                                                        <van-row>
                                                            <van-col span="8">{{item.endcityname}}</van-col>
                                                            <van-col span="16" class="fc">{{item.enddate}}</van-col>
                                                        </van-row>
                                                    </van-step>
                                                </van-steps>
                                            </van-row>

                                            <van-row align="left" class="pt10 pb10" v-if="item.type == 'H'">
                                                <van-col span="4">
                                                    <van-tag text-color="#969799" color="rgb(150 151 153 / 16%)"
                                                             v-if="item.tickettype=='CONFIRM'"> 已确认
                                                    </van-tag>
                                                    <van-tag text-color="#969799" color="rgb(150 151 153 / 16%)"
                                                             v-if="item.tickettype=='CHECK_OUT'"> 已离店
                                                    </van-tag>
                                                    <van-tag text-color="#969799" color="rgb(150 151 153 / 16%)"
                                                             v-if="item.tickettype=='CANCELED'"> 已取消
                                                    </van-tag>
                                                    <van-tag text-color="#969799" color="rgb(150 151 153 / 16%)"
                                                             v-if="item.tickettype=='CHECK_IN'"> 已入住
                                                    </van-tag>
                                                </van-col>
                                                <van-col span="6">
                                                    <van-tag text-color="#969799" color="rgb(150 151 153 / 16%)">
                                                        {{item.startcityname}}
                                                    </van-tag>
                                                </van-col>
                                                <van-col span="8">
                                                    <van-tag text-color="#969799" color="rgb(150 151 153 / 16%)">
                                                        {{item.orderno}}
                                                    </van-tag>
                                                </van-col>
                                            </van-row>

                                            <van-row align="left" class="pt10 pb10" v-if="item.type == 'F'||item.type == 'T'">
                                                <van-col span="5">
                                                    <van-tag v-if="item.tickettype == 'Y'" text-color="#969799"
                                                             color="rgb(150 151 153 / 16%)">原单
                                                    </van-tag>
                                                    <van-tag v-if="item.tickettype == 'T'" text-color="#969799"
                                                             color="rgb(150 151 153 / 16%)">退票单
                                                    </van-tag>
                                                </van-col>
                                                <van-col span="12">
                                                    <van-tag text-color="#969799" color="rgb(150 151 153 / 16%)">
                                                        {{item.orderno}}
                                                    </van-tag>
                                                </van-col>
                                            </van-row>
                                        </div>
                                    </div>
                                </div>
                            </van-row>
                        </div>

                        <div style="background: #FFFFFF;margin-bottom: 10px">
                            <van-row class="pt10 pl20">
                                <van-col span="8" style="font-size: 14px">
                                    <a>发票明细</a>
                                </van-col>
                                <van-col span="14">
                                </van-col>
                                <van-col span="2">
                                    <van-icon name="add-o" @click="selectInvoiceMth"/>
                                </van-col>
                            </van-row>
                            <van-row class="f14 pl20 pt10">
                                <div v-if="invoiceInformation.length>0">
                                    <div v-for="(item,index) in invoiceInformation"
                                         :style="{width:realWidth+'px',minHeight: '30px',background:'rgb(157 194 232 / 30%)',borderRadius:'3px',marginBottom:'5px'}">
                                        <van-row justify="space-between" class="pl10 pt10 pb10">
                                            <van-col span="9" class="fw600 fc">{{ item.invName }}</van-col>
                                            <van-col span="1">
                                                <van-icon v-if="item.btnStatus" name="arrow-down" @click="showInvoice(item,index)"/>
                                                <van-icon v-if="!item.btnStatus" name="arrow" @click="showInvoice(item,index)"/>
                                            </van-col>
                                            <van-col span="3"></van-col>
                                            <van-col span="7" align="right" class="fw600 f14" style="color: #0080FF">￥{{ item.zje }}</van-col>
                                            <van-col span="2" align="center">
                                                <van-icon :name="invoiceCloseIcon" @click="removeInvoice(item,index)"/>
                                            </van-col>
                                            <van-col span="1">
                                            </van-col>
                                        </van-row>
                                        <div class="pl10 pt10 pb10 pr5" v-if="item.btnStatus">
                                            <van-row justify="space-between" class="pr5 pt10">
                                                <van-col>发票代码</van-col>
                                                <van-col>{{ item.fpdm }}</van-col>
                                            </van-row>
                                            <van-row justify="space-between" class="pr5 pt10">
                                                <van-col>发票号码</van-col>
                                                <van-col>{{ item.fphm }}</van-col>
                                            </van-row>
                                            <van-row justify="space-between" class="pr5 pt10">
                                                <van-col>票据信息</van-col>
                                                <van-col>{{ item.invName }}</van-col>
                                            </van-row>
                                            <van-row justify="space-between" class="pr5 pt10">
                                                <van-col>开票时间</van-col>
                                                <van-col>{{ item.cjrq }}</van-col>
                                            </van-row>
                                        </div>
                                    </div>
                                </div>
                            </van-row>
                        </div>
                        <div style="background: #FFFFFF;margin-bottom: 20px">
                            <van-row class="f12 pl20 pt10">
                                <van-col span="9" class="fw600">补助总金额：￥{{subsidyAmount}}</van-col>
                                <van-col span="3"></van-col>
                                <van-col span="12">
                                    <van-checkbox v-model="isSpecialAreas" icon-size="15px"
                                                  @change="changeSpecialAreasStatus">是否新疆、青海、西藏出差
                                    </van-checkbox>
                                </van-col>
                            </van-row>
                            <van-row class="f14 pl20 pt10 pb10" v-for="i in needTypeList">
                                <div :style="{width:realWidth+'px',border:' #4BA3FB solid 1px',borderRadius:'5px',paddingBottom:'20px'}">
                                    <van-row class="pl20 pt10">
                                        <van-col span="8" style="font-weight: 600">{{i.typeName}}</van-col>
                                        <van-col span="10"></van-col>
                                        <van-col span="6">￥{{i.moneyTotal}}</van-col>
                                    </van-row>
                                    <van-row class="pl20 pt10 " justify="left">
                                        <van-tag type="warning" size="medium">￥{{isSpecialAreas?i.needMoneySf:i.needMoney}}/天</van-tag>
                                        <van-tag type="primary" size="medium" class="ml20">出差{{travelDays}}天</van-tag>
                                    </van-row>
                                </div>
                            </van-row>
                        </div>
                    </div>
                    <div class="bottom-wrap" :style="{position: 'relative',bottom:'0px'}">
                        <van-row justify="center">
                            <div style="margin: 16px;">
                                <van-button block type="primary" style="width:200px" native-type="submit" :loading="submitBtnDisable"> 提交 </van-button>
                            </div>
                        </van-row>
                    </div>
                </van-form>
            </div>
        </div>

        <!-- 选择成本中心 -->
        <van-popup v-model:show="selectCostCenterPicker" style="width:100vw; height:100vh" position="bottom">
            <select-cost-center-picker @select-over="selectCostCenter" :data="costCenterData"/>
        </van-popup>
        <!--审批流--->
        <van-action-sheet v-model:show="showApproveSelectPopup" title="请选择审批流" :actions="approvelList" @select="onSelect" style="min-height: 80px">
            <div v-if="!isDeptStatus && approvelList.length<1" style="text-align: center;color: #0080FF;padding: 20px">
                <span>该成本中心未设置审批流,请重新选择成本中心或联系管理员!</span>
            </div>
        </van-action-sheet>
        <van-action-sheet v-model:show="showApproveSelectPopupDept" :actions="approvelList" @select="onSelectDept" style="min-height: 80px">
            <div v-if="isDeptStatus && approvelList.length<1" style="text-align: center;color: #0080FF;padding: 20px">
                <span>该部门未设置审批流,请重新选择成本中心或联系管理员!</span>
            </div>
        </van-action-sheet>
    </div>

    <!--收款账户-->
    <van-popup v-model:show="bankSelectStatus" style="width:100vw;height:100vh" position="bottom">
        <BankCardSelector :isComponentsStatus="true" :selectBankMesssage="selectBankMesssage" @selectBankCardInformation="selectBankCardMth"/>
    </van-popup>
    <!-- 审批显示 -->
    <van-popup v-model:show="seeApproveInfoPicker" style="width:100vw;height:100vh" position="bottom">
        <approval-info @select-over="seeApproveInfoPicker=false" :approve-info-list="approveInfoList" :code="approveNum"/>
    </van-popup>
    <!--  选择发票-->
    <van-popup v-model:show="selectInvoice" style="width:100vw;height:100vh" position="bottom">
        <select-invoice-picker @select-over="selectInvoiceEvent" ref="selectInvoiceOver" :select-list="invoiceInformation"/>
    </van-popup>
    <!--出行明细-->
    <van-popup v-model:show="selectTicketing" v-if="selectTicketing" style="width:100vw;height:100vh" position="bottom">
        <TicketingListSelector ref="ticketSelectorRef" v-if="selectTicketing" @select-over="selectTicketingEvent" :checked="journeyDetailsPrimaryKey" @close="closeSelector"/>
    </van-popup>
    <!--时间选择-->
    <van-calendar v-model:show="showCalander" type="range" :min-date="minDateTime" :max-date="maxDateTime" :default-date="currentCalender" @confirm="onConfirmCalander"/>
</template>

<script setup lang="ts">

    import router from "../../router";
    import {ref, inject, onMounted, watch} from 'vue'
    import moment from 'moment'
    import TicketingListSelector from './component/TicketingListSelector.vue'
    import SelectCostCenterPicker from '../../components/selectComponent/selectCostCenterPicker/index.vue'; // 引入选择成本中心的组件
    import BankCardSelector from '../manageAccounts/index.vue'
    import BankCardSelectorAdd from '../manageAccounts/add.vue'
    import ApprovalInfo from "../application/form/components/approveInfo/index.vue";
    import invoiceCloseIcon from '../../assets/icon/close_invoice.png'
    import jiudian from "../../assets/icon/jiudian.png"
    import feiji from "../../assets/icon/feiji.png"
    import huoche from "../../assets/icon/huoche.png"
    import {
        checkMaxAmount, getNeedTypeList, submitApproval
    } from '../../services/reimburse'; // 导入需要用到的接口
    import {getDefaultPay} from '../../services/manageAccounts'
    import {getApprovalFlowHistory, getApprovalFlowList, selectPerIDCard, getApprovalFlow} from "../../services/current"
    import {showFailToast, showSuccessToast, showToast} from "vant";

    let userInfoData: any = inject("userInfo");
    let realHeight = (document.documentElement.offsetHeight - 163) + 'px';
    let realWidth = (document.documentElement.offsetWidth - 40);
    const submitBtnDisable = ref(false)
    const travelTime = ref("")
    const reimbursementAmount:any = ref('')
    const reason = ref("")
    const costCenter = ref("")
    const paymentInformation = ref("")
    const costApprovel = ref("")
    const deptApprovel = ref("")
    let currentTime = moment(new Date()).format("YYYY-MM-DD")
    let minDate = new Date(currentTime)
    const selectCostCenterPicker = ref(false)
    const costCenterData = ref({
        'xmmc': '',
        'xmye': '',
        'xmbh': '',
        'budGetingId': '',
        'costType': ''
    }); // 将获取回来的成本中心数据存放在该对象中
    const flowList = ref([]); // 成本中心审批流数组
    const flowName = ref(''); // 审批流数组
    const flowId = ref('') // 审批流程
    const isDeptStatus = ref(false)
    const isCost = ref(false)//是否必填
    const showApproveSelectPopup = ref(false)//选择审批流弹出框
    const showApproveSelectPopupDept = ref(false)//选择审批流弹出框-----部门审批情况下
    const bankSelectStatus = ref(false)
    const approvelList = ref([])
    const selectBankMesssage = ref()
    const approvalFlowType = ref('')
    const deptSelector = ref()
    const seeApproveInfoPicker = ref(false)
    const approveInfoList = ref([])
    const approveCCInfo = ref([])
    const approveNum = ref([])
    const travelDetails = ref([])
    const invoiceInformation = ref([])
    const selectInvoice = ref(false)
    const selectTicketing = ref(false)
    const showCalander = ref(false)
    const needTypeList = ref([])
    const isSpecialAreas = ref(false)
    const travelDays = ref(0)
    const travelNum = ref('CL' + moment(new Date()).format('YYYYMMDDHHmmss') + Math.round(Math.random() * 10))
    const createDate = ref(moment((new Date())).format('YYYY-MM-DD'))
    const startTime = ref()
    const endTime = ref()
    const subsidyAmount = ref(0)
    let startDate = moment().subtract(3, "months").format("YYYY-MM-DD")
    let endDate = moment().add(3, "months").format("YYYY-MM-DD")
    const minDateTime = ref(new Date(startDate))
    const maxDateTime = ref(new Date(endDate))
    let currentCalender = new Date()
    const selectCostOver = ref()
    const journeyDetailsPrimaryKey = ref([])
    const ticketSelectorRef: any = ref<InstanceType<typeof TicketingListSelector>>(); // 实例化
    const selectInvoiceOver = ref()

    watch(
        () => reimbursementAmount.value,
        (newVal, oldVal) => {
          if(reimbursementAmount.value * 1!=0){
            reimbursementAmount.value = (reimbursementAmount.value * 1).toFixed(2) as never
          }
        },
        {
            immediate: true,
            deep: true
        }
    )

    onMounted(() => {
        getConfiguration()
        getDefaultPayMth()
        getNeedType()
    })

    //获取配置
    function getConfiguration() {
        let functionConfig = userInfoData.functionConfig;
        //判断是否成本中心审批
        if (functionConfig['fk005'].configValue == '1') {
            isDeptStatus.value = false
            isCost.value = true
            approvalFlowType.value = 'costCenterFlow'
        } else {
            isDeptStatus.value = true
            approvalFlowType.value = 'deptFlow'
            //部门审批情况下判断是否必填成本中心
            if (functionConfig['fk003'].configValue == '0') {
                isCost.value = true
            } else {
                isCost.value = false
            }
            let params = {
                applyTypeId: 'FK_CLBX_SQ',
                xmbh: '',
                approvalFlowType: approvalFlowType.value,
            }
            GetCostCenterFlow(params);
        }
    }

    //获取默认银行卡信息
    async function getDefaultPayMth() {
        let params = {corpCode: userInfoData.userInfo.corpCode, empCode: userInfoData.userInfo.empCode}
        let res: any = await getDefaultPay(params)
        if (res.code == 200 && res.data != null) {
            selectBankMesssage.value = res.data
            paymentInformation.value = res.data.payeeName + '  ' + res.data.bankcardnum
        }
    }

    //获取补助信息
    async function getNeedType() {
        let params = {
            corpCode: userInfoData.userInfo.corpCode
        }
        let res: any = await getNeedTypeList(params)
        if (res.code == 200) {
            res.data.map((item: any) => {
                item.moneyTotal = 0
            })
            needTypeList.value = res.data
        }
    }

    function showCalendarMth() {
        showCalander.value = true
    }

    //选择时间
    function onConfirmCalander(e: any) {
        let date = ''
        let days = 0
        if (e) {
            startTime.value = moment(e[0]).format("YYYY-MM-DD")
            endTime.value = moment(e[1]).format("YYYY-MM-DD")
            date = startTime.value + " 至 " + endTime.value
            days = (e[1].getTime() - e[0].getTime()) / 24 / 60 / 60 / 1000
            travelDays.value = days
        }
        getTotalReimbursementAmount()
        travelTime.value = date
        showCalander.value = false
    }
    //选择是否是新疆、青海、西藏出差
    function changeSpecialAreasStatus(e: any) {
        getTotalReimbursementAmount()
    }

    //关闭出行明细弹窗
    function closeSelector() {
        selectTicketing.value = false
    }

    //显示出行明细弹窗
    function showTravelDetailDialog() {
        selectTicketing.value = true
    }

    //选择行程明细
    function selectTicketingEvent(e: any) {
        travelDetails.value = e
        let list: any = []
        if (e.length > 0) {
            travelDetails.value = e
            e.map((item: any) => {
                list.push(item.id as never)
            })
        }
        journeyDetailsPrimaryKey.value = list
        getTotalReimbursementAmount()
        closeSelector()
    }


    function selectInvoiceMth() {
        selectInvoice.value = true
        setTimeout(()=> selectInvoiceOver.value.onRefresh('1'),500)
    }

    function selectInvoiceEvent(e: any) {
        let money = 0
        e.map((item: any) => {
            item.btnStatus = false
        })
        invoiceInformation.value = e
        getTotalReimbursementAmount()
        selectInvoice.value = false
    }

    function showInvoice(item: any, index: any) {
        invoiceInformation.value[index].btnStatus = !item.btnStatus
    }

    function removeInvoice(item: any, index: any) {
        invoiceInformation.value.splice(index, 1)
        getTotalReimbursementAmount()
    }

    const onConfirm = (date: any) => {
        showCalander.value = false;
    };

    //获取报销总额
    function getTotalReimbursementAmount() {
        let money = 0
        if (isSpecialAreas.value) {
            needTypeList.value.map((item: any) => {
                item.moneyTotal = item.needMoneySf * travelDays.value
                money = money + item.moneyTotal
            })
        } else {
            needTypeList.value.map((item: any) => {
                item.moneyTotal = item.needMoney * travelDays.value
                money = money + item.moneyTotal
            })
        }
        subsidyAmount.value = money
        if (invoiceInformation.value.length > 0) {
            invoiceInformation.value.map((item: any) => {
                money = money * 1 + item.zje * 1
            })
        }
        if (travelDetails.value.length > 0) {
            travelDetails.value.map((ite: any) => {
                money = money * 1 + ite.totalamount * 1
            })
        }
        // 不等于0的时候再赋值
        if(money!=0){
          reimbursementAmount.value = money
        }
    }

    // 打开成本中心组件
    const selectCostCenterOpen = () => {
        selectCostCenterPicker.value = true;
    }

    // 选择成本中心
    const selectCostCenter = (data: any) => {
        if (data.xmmc) {
            costCenterData.value = data;
            costCenter.value = `${data.budGetingYear}-${data.xmmc}   ¥ ${data.xmye}`;
            if (isDeptStatus.value == false) {
                let params = {
                    applyTypeId: '',
                    xmbh: data.xmbh,
                    approvalFlowType: approvalFlowType.value,
                }
                GetCostCenterFlow(params);
            }
        }
        selectCostCenterPicker.value = false;
    }

    // 获取审批流
    const GetCostCenterFlow = async (params: any) => {
        let payLoad = {
            applyTypeId: params.applyTypeId,
            xmbh: params.xmbh,
            deptId: userInfoData.userInfo.deptId,
            corpCode: userInfoData.userInfo.corpCode,
            approvalFlowType: params.approvalFlowType,
        }
        let result: any = await getApprovalFlow(payLoad);
        flowName.value = '';
        flowId.value = '';
        if (result.data[0].length != 0) {
            let dataList = result.data[0]
            let flowList: { name: any; value: any; }[] = []
            dataList.map((item: any) => {
                let obj = {
                    name: item.flowname,
                    value: item.flowid
                }
                flowList.push(obj)
            })
            approvelList.value = flowList as never
        } else {
            approvelList.value = []
        }
    }

    //获取审批流详情
    function queryFlowDetail(e: any) {
        getFlowHistory(flowId.value)
        seeApproveInfoPicker.value = true
    }

    // 根据审批流回显审批流人员信息
    const getFlowHistory = async (value: any) => {
        let payload = {
            "school": userInfoData.userInfo.corpCode,
            'flowId': value,
            "corpCode": userInfoData.userInfo.corpCode,
            'empCode': userInfoData.userInfo.empCode,
        }
        let result: any = await getApprovalFlowHistory(payload);
        if (result.success) {
            let flowHistory: any = JSON.parse(result.data);
            approveInfoList.value = flowHistory.data.toUserList;
            approveCCInfo.value = flowHistory.data.ccUserList;
            approveNum.value = flowHistory.data.isVote;
            const obj: any = {};
            for (let i = 0; i <= approveInfoList.value.length - 1; i++) {
                let no = approveInfoList.value[i]['no'];
                let isAllPass = approveInfoList.value[i]['isAllPass'];
                if (!obj[no]) {
                    obj[no] = {
                        no,
                        isAllPass,
                        children: []
                    }
                }
                if (no === obj[no].no) {
                    obj[no].children.push(approveInfoList.value[i])
                }
            }
            approveInfoList.value = [];
            Object.values(obj).forEach((item) => {
                return approveInfoList.value.push(item as never);
            })
            let userName: any = [];
            flowHistory.data.ccUserList.forEach((item: any) => userName.push(`${item.ccUserName}(${item.toRole})`));
        }
    }

    //选择成本中心审批流
    function onSelect(e: any) {
        costApprovel.value = e.name
        flowId.value = e.value
        showApproveSelectPopup.value = false
    }

    //选择部门审批流
    function onSelectDept(e: any) {
        deptApprovel.value = e.name
        flowId.value = e.value
        deptSelector.value = e
        showApproveSelectPopupDept.value = false
    }

    //显示银行卡弹窗
    function selectBankAccount() {
        bankSelectStatus.value = true
    }

    //选中银行卡信息
    function selectBankCardMth(e: any) {
        if (e) {
            selectBankMesssage.value = e
            paymentInformation.value = e.payeeName + '  ' + e.bankcardnum
        }
        bankSelectStatus.value = false
    }

    //提交
    async function onSubmit() {
        let needListType: any = [] //报销补助信息
        let orderList: any = [] //出行明细
        needTypeList.value.map((item: any) => {
            let obj = {
                type: item.type,
                needsstandard: isSpecialAreas ? item.needMoneySf : item.needMoney,
                days: travelDays.value,
                peoplenum: 1,
                needsmoney: item.moneyTotal
            }
            needListType.push(obj)
        })
        travelDetails.value.map((item: any) => {
            let i = {
                id: item.id,
                type: item.type
            }
            orderList.push(i)
        })
        //发票明细和出行明细至少选一个
        if (travelDetails.value.length < 1 && invoiceInformation.value.length < 1) {
            showToast("发票明细和出行明细至少选择一项")
            return
        }
        let invoiceList = []
        invoiceInformation.value.map((ite: any) => {
            let obj = {
                spendAmount: ite.zje,
                spendType: ite.invName,
                dzfph: ite.dzfph,
                fpdm: ite.fpdm,
                fphm: ite.fphm,
                fpygbh: ite.fpygbh,
                fjid: ite.fjid
            }
            invoiceList.push(obj);
        })
        let params = {
            toAccount: "test_auvgo8899",
            sysType: "用友",
            applyNo: travelNum.value,
            deptId: userInfoData.userInfo.deptId,
            applyReason: reason.value,
            empCode: userInfoData.userInfo.empCode,
            empName: userInfoData.userInfo.userName,
            deptName: userInfoData.userInfo.deptName,
            bxAmount: reimbursementAmount.value,
            corpCode: userInfoData.userInfo.corpCode,
            isRemote: isSpecialAreas.value,
            applicationInfo: {
                approveNo: "",
                startTime: startTime.value,
                returnTime: endTime.value,
                fromCity: "",
                toCity: "",
                togetherStaff: userInfoData.userInfo.userName
            },
            bxNeedsInfoList: needListType,
            payeeinfo: selectBankMesssage.value,
            travelinvoicedetails: invoiceList,
            orderList: orderList,
            xmmc: costCenterData.value.xmmc,
            xmbh: costCenterData.value.xmbh,
            flowId: flowId.value,
            approvalFlowType: approvalFlowType.value,
            isLeader: "true",
            costType: isDeptStatus.value ? 'department' : 'cost',
            budGetingId: costCenterData.value.budGetingId,
        }
        submitBtnDisable.value = true
        let res: any = await submitApproval(params)
        if (res.code == 200) {
            if (res.data.code == 200) {
                setTimeout(() => {
                    submitBtnDisable.value = false
                }, 300)
                showSuccessToast('提交成功')
                router.go(-1)
            } else {
                showFailToast(res.data.message)
                submitBtnDisable.value = false
            }
        } else {
            showFailToast(res.data.message)
            submitBtnDisable.value = false
        }
    }

    //校验是否超出预算
    async function checkMaxAmountMth() {
        let params = {
            bxType: "travel",
            corpCode: userInfoData.userInfo.corpCode,
            deptId: userInfoData.userInfo.deptId,
            empCode: userInfoData.userInfo.empCode,
            money: reimbursementAmount.value
        }
        let res: any = await checkMaxAmount(params)
        if (res.code == 200) {
            if (res.data.isControl == '1') {
                switch (res.data.tips) {
                    case"0": //提示信息
                        showFailToast(res.message)
                        setTimeout(() => {
                            //提交
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
    }

    function back() {
        router.go(-1)
    }

</script>
<style lang="less" scoped>

    .bottom-wrap {
        height: 75px;
        width: 100%;
        /*top:var(--realHeight);*/
        background: #FFFFFF;
    }

    .spanCss {
        color: red;
        position: absolute;
        left: 80px;
        z-index: 999;
        padding-top: 13px;
    }

    .fw600 {
        font-weight: 600;
    }

    .fc {
        color: #1a1a1a;
    }

    /deep/ .stepsCss .van-step--vertical:not(:last-child):after {
        border-bottom-width: 0px;
    }

    /deep/ .stepsCss .van-icon-checked:before {
        content: "";
        display: block;
        width: var(--van-step-circle-size);
        height: var(--van-step-circle-size);
        background-color: var(--van-step-circle-color);
        border-radius: 50%;
        margin-bottom: 2px;

    }
</style>
