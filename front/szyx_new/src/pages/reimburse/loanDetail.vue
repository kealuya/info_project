<template>
    <div>
        <div>
            <NavBar :leftText="'返回'" :title="'借款单详情'" :leftArrow="true" @leftEvent="back()"></NavBar>
        </div>
        <van-loading style="padding-top: 50px" align="center" v-if="loadingStatus"/>
        <div class="mt10" :style="{height:realHeight,overflowY:'scroll'}" v-if="loadingStatus==false">
            <van-row class="ml10" :style="{ background:'#FFFFFF',width:realwidth}">
                <div class="ml10 mt10">
                    <van-row>
                        <van-col span="16">
                            <van-row class="pt10">
                                <span>借款申请单</span>
                            </van-row>
                            <van-row class="pt10 f14 spanCss">
                                <span> 申请单号：</span>
                            </van-row>
                            <van-row class="pt10 f14">
                                <span> {{loanDetail.applyno}}</span>
                            </van-row>
                            <van-row class="pt10 f14 spanCss">
                                <span> 借款事由：</span>
                            </van-row>
                            <van-row class="pt10 f14 pb10">
                                <span> {{loanDetail.applyreason}} </span>
                            </van-row>
                            <van-row class="spanCss f14">
                                <span> 借款金额：</span>
                            </van-row>
                            <van-row class="pt10 pb10 fw600" style="font-size: 18px; color: #FF9200;">
                                <span> ￥{{loanDetail.loanamount}} </span>
                            </van-row>
                        </van-col>
                        <van-col span="8" >
                            <van-icon class="iconCss" a v-if="apprApproval.currentState == 'wait'" size="6rem"
                                      :name="waitshenpi_icon"/>
                            <!--  <van-icon v-if="loanDetail.state == '0'" :name="yichexiao_icon"/>-->
                            <van-icon class="iconCss"
                                      v-if="apprApproval.currentState == 'over' && apprApproval.currentResult == 'ng'"
                                      size="6rem" :name="yijujue_icon"/>
                            <van-icon class="iconCss"
                                      v-if="apprApproval.currentState == 'over' && apprApproval.currentResult == 'ok'"
                                      size="6rem" :name="yitongyi_icon"/>
                        </van-col>
                    </van-row>
                </div>
            </van-row>
            <van-row class="ml10 mt10 f14" :style="{ background:'#FFFFFF',width:realwidth}">
                <div class="ml10 mt10">
                    <van-row class="pt10" :style="{width:realwidth}">
                        <van-col span="8" class="spanCss">申请时间:</van-col>
                        <van-col span="16">{{loanDetail.createtime}}</van-col>
                    </van-row>
                    <van-row class="pt10">
                        <van-col span="8" class="spanCss">还款状态:</van-col>
                        <van-col span="16">{{loanDetail.bxState=='false'?'未还款':'已还款'}}</van-col>
                    </van-row>
                    <van-row class="pt10">
                        <van-col span="8" class="spanCss">还款日期:</van-col>
                        <van-col span="16">{{loanDetail.repaymenttime}}</van-col>
                    </van-row>
                    <van-row class="pt10">
                        <van-col span="8" class="spanCss">银行卡号:</van-col>
                        <van-col span="16">{{loanDetail.banknum}}</van-col>
                    </van-row>
                    <van-row class="pt10">
                        <van-col span="8" class="spanCss">收款名称:</van-col>
                        <van-col span="16">{{loanDetail.payeeName}}</van-col>
                    </van-row>
                    <van-row class="pt10 pb10" v-if="loanDetail.xmmc!=''||loanDetail.xmmc!=null">
                        <van-col span="8" class="spanCss">成本中心:</van-col>
                        <van-col span="16">{{loanDetail.xmmc}}</van-col>
                    </van-row>
                </div>
            </van-row>
            <van-row class="ml10 mt10" :style="{ background:'#FFFFFF',width:realwidth}">
                <div class="ml10 mt10">
                    <van-row class="pt10">
                        <span>冻结记录</span>
                    </van-row>
                    <div style="min-height: 35px">
                        <van-steps direction="vertical" :active="0" v-for="(item,index) in freezingRecords">
                            <van-step>
                                <h>{{item.creationTime}}</h>
                                <div class="cardList" :style="{width:realwidth-40+'px'}">
                                    <div :style="{width:apprDivWidth-20+'px',minHeight:'30px',marginTop:'10px',marginLeft:'10px'}">
                                        <van-row>
                                            <van-col span="15" class="fw800 f14 pt5 fc">
                                                {{item.costName}}
                                            </van-col>
                                            <van-col span="9" align="right" class="pr5">
                                                <van-tag class="tagCss" plain type="danger" text-color="red"
                                                         size="medium">
                                                    冻结经费
                                                </van-tag>
                                            </van-col>
                                        </van-row>
                                        <van-row justify="left" class="pt10 pb10">
                                            <van-col span="2">
                                                <van-icon name="gold-coin" size="16" color="#f4af15"/>
                                            </van-col>
                                            <van-col span="7" class="fc">冻结金额:</van-col>
                                            <van-col span="15" style="color: #f4af15">￥{{item.je}}</van-col>
                                        </van-row>
                                        <van-row justify="left" class=" pb10">
                                            <van-col span="2">
                                                <van-icon :name="Vector"/>
                                            </van-col>
                                            <van-col span="7" class="fc">操作时间:</van-col>
                                            <van-col span="15" class="fc">{{item.creationTime}}</van-col>
                                        </van-row>
                                    </div>
                                </div>
                            </van-step>
                        </van-steps>

                    </div>
                </div>
            </van-row>
            <van-row class="ml10 mt10 mb10" :style="{background:'#FFFFFF',width:realwidth}">
                <div class="ml10 mt10">
                    <van-row class="pt10">
                        <span>审批进度</span>
                    </van-row>
                    <div class="formCard">
                        <div style="margin-top: 10px;margin-bottom: 10px">
                            <van-steps direction="vertical" :active="nextApprIndex" :style="{width: apprDivWidth+'px'}">
                                <van-step>
                                    <van-row>
                                        <van-col :span="4">
                                            <div class="f14 graytxt bold">发起</div>
                                        </van-col>
                                        <van-col :span="20">
                                            <div style="padding-top: 3px;padding-bottom: 3px;background: #f6f6f6">
                                                <div class="approveCard">
                                                    <van-row class="f12 ml5">
                                                        <van-col span="12">
                                                            {{userInfoData.userInfo.userName
                                                            +'('+userInfoData.userInfo.jobName+')'}}
                                                        </van-col>
                                                        <van-col span="12">
                                                            {{loanDetail.createtime}}
                                                        </van-col>
                                                    </van-row>
                                                </div>
                                            </div>
                                        </van-col>
                                    </van-row>
                                </van-step>
                                <van-step v-for="item in approveInfoList">
                                    <van-row>
                                        <van-col :span="4">
                                            <van-row align="center">
                                                <van-col>
                                                    <div class="f14 graytxt bold">
                                                        <div v-if="item.isAllPass==null">单审</div>
                                                        <div v-else>{{ item.isAllPass == '1' ? '或审' : '会审' }}</div>
                                                        <div>
                                                            <span v-if="item.result=='ok'"
                                                                  class="okCss f12 bold">通过</span>
                                                            <span v-else-if="item.result=='ng'"
                                                                  class="ng f12 bold">已拒绝</span>
                                                            <span v-else class="graytxt f12 bold">待审批</span>
                                                        </div>
                                                    </div>
                                                </van-col>
                                                <van-col>
                                                    <van-icon name="info-o"/>
                                                </van-col>
                                            </van-row>
                                        </van-col>
                                        <van-col :span="20">
                                            <div style="padding-top: 3px;padding-bottom: 3px;background: #f6f6f6">
                                                <div v-for="ite in item.list">
                                                    <div class="approveCard">
                                                        <van-row>
                                                            {{ite.toUser+'('+ite.toRole+')'}}
                                                        </van-row>
                                                    </div>
                                                </div>
                                            </div>
                                        </van-col>
                                    </van-row>
                                </van-step>
                            </van-steps>
                        </div>
                    </div>
                </div>
            </van-row>
        </div>
    </div>
</template>

<script setup lang="ts">

    import router from "../../router";
    import {ref, inject, onMounted, onBeforeMount} from 'vue'
    import moment from 'moment'
    import {
        addLoanApply, getFreezingThawingHistory, getLoanDetail
    } from '../../services/reimburse'; // 导入需要用到的接口
    import {
        approvalHistory,
    } from "../../services/current"
    // import ApprovalInfo from "../application/form/components/approveInfo/index.vue";
    import waitshenpi_icon from "../../assets/img/waitshenpi_icon.png"
    import yichexiao_icon from "../../assets/img/yichexiao_icon.png"
    import yijujue_icon from "../../assets/img/yijujue_icon.png"
    import yitongyi_icon from "../../assets/img/yitongyi_icon.png"
    import Vector from "../../assets/icon/Vector.png"


    let userInfoData: any = inject("userInfo");
    let realHeight = (document.documentElement.offsetHeight - 60) + 'px';
    let realwidth = (document.documentElement.offsetWidth - 20) + 'px';
    let apprDivWidth = (document.documentElement.offsetWidth - 65);
    const loadingStatus = ref(false)
    const freezingRecords = ref([])
    const loanDetail = ref({
        applyno: '',
        applyreason: '',
        createtime: '',
        bxState: '',
        repaymenttime: '',
        banknum: '',
        payeeName: '',
        xmmc: ''
    })
    const approveInfoList = ref([]);
    const nextApprIndex = ref('');
    const apprApproval: any = ref({})

    onBeforeMount(() => {
        let loanId = router.currentRoute.value.params.loanId as never
        loadingStatus.value = true

        getLoanDetailMth(loanId)
        getFreezingThawingMth(loanId)
        getApprovalHistory(loanId)
        setTimeout(() => {
            loadingStatus.value = false

        }, 500)

    })

    function back() {
        router.go(-1)
    }

    async function getFreezingThawingMth(e: any) {
        let params = {
            corpCode: userInfoData.userInfo.corpCode, approveNo: e
        }
        let res: any = await getFreezingThawingHistory(params)
        if (res.code == 200) {
            freezingRecords.value = res.data
        }
    }

    async function getLoanDetailMth(e: any) {
        let param = {applyno: e, corpCode: userInfoData.userInfo.corpCode}
        let res: any = await getLoanDetail(param)
        if (res.code == 200) {
            loanDetail.value = res.data["loan"]
        }
    }

    // 获取申请单审批流程历史
    async function getApprovalHistory(data: any) {
        const payload: any = {
            "applyNo": data,
            "corpCode": userInfoData.userInfo.corpCode,
            "empCode": userInfoData.userInfo.empCode,
            "equipmentType": ""
        }
        let result: any = await approvalHistory(payload);
        if (result.success) {
            let dataObj = JSON.parse(result.data).data;
            apprApproval.value = dataObj
            let apprInfo: any = dataObj.stateMap
            for (let key in apprInfo) {
                const apprData = {
                    no: key,
                    list: apprInfo[key],
                    isAllPass: apprInfo[key][0].isAllPass,
                    state: apprInfo[key][0].state,
                    result: apprInfo[key][0].apprResult,
                }
                if (apprInfo[key][0].state == 'wait') {
                    nextApprIndex.value = key;
                }
                approveInfoList.value.push(apprData as never);
            }
        }
    }


</script>
<style lang="less" scoped>

    .spanCss {
        color: rgb(150, 151, 153);
    }

    .cardList {
        border: 1px solid #1989fa;
        min-height: 30px;
        border-radius: 3px;
        margin-top: 10px;
    }

    .tagCss {
        --van-tag-plain-background: #f1041f1a;
    }

    .fc {
        color: #1a1a1a;
    }

    .fw600 {
        font-weight: 600;
    }

    .approveCard {
        margin: 5px;
        padding-top: 10px;
        padding-bottom: 5px;
        background: white;
        border-radius: 5px;
    }

    .okCss {
        color: #0088ff;
    }

    .ng {
        color: #ff6257;
    }

    .iconCss {
        padding-top: 75px;
        padding-right: 10%;
        /*float: right;*/
    }


</style>
