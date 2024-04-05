<template>
    <div>
        <div>
            <NavBar :leftText="'返回'" :title="'借款申请单'" :leftArrow="true" @leftEvent="back()"></NavBar>
        </div>
        <div>
            <van-notice-bar color="#1989fa" :scrollable="false" wrapable background="#ecf9ff"
                            left-icon="info-o" class="f10">
                借款申请是个人向公司提交的有具有公务性质的借款申请(例如：差旅费，采购费以及其他临时性借款)，不包括个人借款业务
            </van-notice-bar>
            <div>
                <van-form @submit="checkMaxAmountMth">
                    <div :style="{height:realHeight }">
                        <van-cell-group class="pb10 mt10">
                            <i class="spanCss">*</i>
                            <van-field
                                    v-model="reasonForBorrowing"
                                    name="reasonForBorrowing"
                                    label="借款事由"
                                    placeholder="请输入借款事由"
                                    :rules="[{ required: true, message: '请输入借款事由' }]"
                                    class="mt10 mb10"
                                    label-width="7.5em"
                                    maxlength="20"
                                    label-align="top"
                            />
                            <i class="spanCss">*</i>
                            <van-field
                                    v-model="loanAmount"
                                    name="loanAmount"
                                    label="借款金额"
                                    placeholder="请输入借款金额"
                                    :rules="[{ required: true, message: '请输入借款金额' }]"
                                    class="mt10 mb10 "
                                    label-width="7.5em"
                                    maxlength="10"
                                    type="number"
                                    label-align="top"
                            />
                            <i class="spanCss">*</i>
                            <van-field
                                    v-model="repaymentDate"
                                    name="repaymentDate"
                                    label="还款日期"
                                    placeholder="请选择还款日期"
                                    :rules="[{ required: true, message: '请选择还款日期' }]"
                                    class="mt10 mb10 "
                                    label-width="7.5em"
                                    label-align="top"
                                    readonly
                                    is-link
                                    @click="showCalendar=true"
                            />
                            <van-calendar v-model:show="showCalendar" :min-date="minDate" @confirm="onConfirm"/>
                            <van-field
                                    v-model="borrowingRemarks"
                                    name="borrowingRemarks"
                                    label="借款备注"
                                    placeholder="请输入借款备注"
                                    class="mt10 mb10 "
                                    label-width="7.5em"
                                    maxlength="25"
                                    label-align="top"

                            />
                        </van-cell-group>
                        <i class="spanCss mt5" style="margin-left: 60px">*</i>
                        <van-field
                                v-model="paymentInformation"
                                name="paymentInformation"
                                label="收款信息(账户信息)"
                                placeholder="请选择银行卡号"
                                :rules="[{ required: true, message: '请选择银行卡号' }]"
                                class="mt10 mb10 "
                                readonly
                                is-link="true"
                                label-width="9.5em"
                                label-align="top"
                                @click="selectBankAccount"

                        />
                        <i class="spanCss" v-if="isCost">*</i>
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
                                @click="selectCostCenterOpen"
                        />
                        <div v-if="!isDeptStatus && isCost && costCenter !=''">
                            <i class="spanCss">*</i>
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
                                    :right-icon="costApprovel==''?'':'info-o'"
                                    @click-right-icon.stop="queryFlowDetail('cost')"
                                    @click="showApproveSelectPopup = true"
                            />
                        </div>
                        <div v-if="isDeptStatus">
                            <i class="spanCss ml10">*</i>
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
                                    :right-icon="deptApprovel==''?'':'info-o'"
                                    @click-right-icon.stop="queryFlowDetail('dept')"
                                    @click="showApproveSelectPopupDept = true"
                            />
                        </div>
                    </div>
                    <div class="bottom-wrap" :style="{position: 'relative',bottom:'0px'}">
                        <van-row justify="center">
                            <div style="margin: 16px;">
                                <van-button block type="primary" style="width:200px" native-type="submit"
                                            :loading="submitBtnDisable"  loading-type="spinner">
                                    提交
                                </van-button>
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

        <van-action-sheet v-model:show="showApproveSelectPopup" :actions="approvelList" @select="onSelect"
                          style="min-height: 80px">
            <div v-if="!isDeptStatus && approvelList.length<1" style="text-align: center;color: #0080FF;padding: 20px">
                <span>该成本中心未设置审批流,请重新选择成本中心或联系管理员!</span>
            </div>
        </van-action-sheet>

        <van-action-sheet v-model:show="showApproveSelectPopupDept" :actions="approvelList" @select="onSelectDept"
                          style="min-height: 80px">
            <div v-if="isDeptStatus && approvelList.length<1" style="text-align: center;color: #0080FF;padding: 20px">
                <span>该部门未设置审批流,请重新选择成本中心或联系管理员!</span>
            </div>
        </van-action-sheet>
    </div>

    <van-popup v-model:show="bankSelectStatus" style="width:100vw;height:100vh" position="bottom">
        <BankCardSelector :isComponentsStatus="true" :selectBankMesssage="selectBankMesssage"
                          @selectBankCardInformation="selectBankCardMth"/>
    </van-popup>
    <!-- 审批显示 -->
    <van-popup v-model:show="seeApproveInfoPicker" style="width:100vw;height:100vh" position="bottom">
        <approval-info @select-over="seeApproveInfoPicker=false" :approve-info-list="approveInfoList"
                       :code="approveNum"/>
    </van-popup>
</template>

<script setup lang="ts">

    import router from "../../router";
    import {ref, inject, onMounted} from 'vue'
    import moment from 'moment'
    import SelectCostCenterPicker from '../../components/selectComponent/selectCostCenterPicker/index.vue'; // 引入选择成本中心的组件
    import BankCardSelector from '../manageAccounts/index.vue'
    import {
        addLoanApply, checkMaxAmount

    } from '../../services/reimburse'; // 导入需要用到的接口
    import {getDefaultPay} from '../../services/manageAccounts'
    import {getApprovalFlowHistory, getApprovalFlowList, selectPerIDCard, getApprovalFlow} from "../../services/current"
    import {showFailToast, showSuccessToast} from "vant";
    import ApprovalInfo from "../application/form/components/approveInfo/index.vue";

    let userInfoData: any = inject("userInfo");
    let realHeight = (document.documentElement.offsetHeight - 220) + 'px';
    const submitBtnDisable = ref(false)

    //借款事由
    const reasonForBorrowing = ref("")
    const loanAmount = ref()
    const repaymentDate = ref("")
    const borrowingRemarks = ref("")
    const costCenter = ref("")
    const paymentInformation = ref("")
    const approvelFlowObj = ref({})
    const costApprovel = ref("")
    const deptApprovel = ref("")

    const showCalendar = ref(false)
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
    const defaultBankMessage = ref({})
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

    onMounted(() => {
        let selectAccountInformation = JSON.parse(router.currentRoute.value.params.selectAccountInformation as never)
        if (selectAccountInformation != "") {
            defaultBankMessage.value = selectAccountInformation
            paymentInformation.value = selectAccountInformation.payeeName +'  '+ selectAccountInformation.bankcardnum
        }
        getConfiguration()
        getDefaultPayMth()
    })

    function back() {
        router.go(-1)
    }

    const onConfirm = (date: any) => {
        repaymentDate.value = moment(date).format("YYYY-MM-DD")
        showCalendar.value = false;
    };

    // 打开成本中心组件
    const selectCostCenterOpen = () => {
        // costCenterData.value = costCenterData.value;
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

    function selectBankAccount() {
        bankSelectStatus.value = true
    }

    function selectBankCardMth(e: any) {
        if (e) {
            selectBankMesssage.value = e
            paymentInformation.value = e.payeeName +'  '+ e.bankcardnum
        }
        bankSelectStatus.value = false
    }

    function getConfiguration() {
        let functionConfig = userInfoData.functionConfig;
        //判断是否成本中心审批
        if (functionConfig['fk006'].configValue == '1') {
            isDeptStatus.value = false
            isCost.value = true
            approvalFlowType.value = 'costCenterFlow'
        } else {
            isDeptStatus.value = true
            approvalFlowType.value = 'deptFlow'
            //部门审批情况下判断是否必填成本中心
            if (functionConfig['fk010'].configValue == '0') {
                isCost.value = true
            } else {
                isCost.value = false
            }
            let params = {
                applyTypeId: 'FK_LOAN_SQ',
                xmbh: '',
                approvalFlowType: approvalFlowType.value,
            }
            GetCostCenterFlow(params);
        }
    }

    function onSelect(e: any) {
        costApprovel.value = e.name
        flowId.value = e.value
        showApproveSelectPopup.value = false
    }

    function onSelectDept(e: any) {
        deptApprovel.value = e.name
        flowId.value = e.value
        deptSelector.value = e
        showApproveSelectPopupDept.value = false
    }

    async function onSubmit() {
        let params = {
            flowId: flowId.value,
            approvalFlowType: approvalFlowType.value,
            applyno: 'JK' + moment(new Date()).format('YYYYMMDDHHmmss') + Math.round(Math.random() * 10),
            applyreason: reasonForBorrowing.value,
            empcode: userInfoData.userInfo.empCode,
            empname: userInfoData.userInfo.userName,
            loanamount: loanAmount.value,
            deptid: userInfoData.userInfo.deptId,
            deptname: userInfoData.userInfo.deptName,
            bz: borrowingRemarks.value,
            corpcode: userInfoData.userInfo.corpCode,
            xmbh: costCenterData.value.xmbh,
            xmmc: costCenterData.value.xmmc,
            banknum: selectBankMesssage.value.bankcardnum as never,
            payeeName: selectBankMesssage.value.payeeName,
            repaymenttime: repaymentDate.value,
            costType: costCenterData.value.costType,
            budGetingId: costCenterData.value.budGetingId,
            corpCode: userInfoData.userInfo.corpCode
        }
        submitBtnDisable.value = true
        let res: any = await addLoanApply(params)
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
        let flowHistory: any = JSON.parse(result.data);
        approveInfoList.value = flowHistory.data.toUserList;
        approveCCInfo.value = flowHistory.data.ccUserList;
        approveNum.value = flowHistory.data.isVote;
        const object: any = {};
        for (let i = 0; i <= approveInfoList.value.length - 1; i++) {
            let no = approveInfoList.value[i]['no'];
            let isAllPass = approveInfoList.value[i]['isAllPass'];
            if (!object[no]) {
                object[no] = {
                    no,
                    isAllPass,
                    children: []
                }
            }
            if (no === object[no].no) {
                object[no].children.push(approveInfoList.value[i])
            }
        }
        approveInfoList.value = [];
        Object.values(object).forEach((item) => {
            return approveInfoList.value.push(item as never);
        })
        let userName: any = [];
        flowHistory.data.ccUserList.forEach((item: any) => userName.push(`${item.ccUserName}(${item.toRole})`));
    }

    async function checkMaxAmountMth() {
        let params = {
            bxType: "loan",
            corpCode:userInfoData.userInfo.corpCode,
            deptId: userInfoData.userInfo.deptId,
            empCode: userInfoData.userInfo.empCode,
            money: loanAmount.value
        }
        let res: any = await checkMaxAmount(params)
        if (res.code == 200) {
            if (res.data.isControl == '1') {
                switch (res.data.tips) {
                    case"0": //提示信息
                        showFailToast(res.message)
                        setTimeout(()=>{
                            onSubmit()
                        },200)
                        break;
                    case"1": //禁止提交
                        break;

                    case"2": //不控制
                        break;
                }

            }else {
                onSubmit()
            }
        }
    }

    async function getDefaultPayMth() {
        let params = {corpCode: userInfoData.userInfo.corpCode, empCode: userInfoData.userInfo.empCode}
        let res: any = await getDefaultPay(params)
        if (res.code == 200 && res.data != null) {
            selectBankMesssage.value = res.data
            paymentInformation.value = res.data.payeeName +'  '+ res.data.bankcardnum
        }
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

</style>
