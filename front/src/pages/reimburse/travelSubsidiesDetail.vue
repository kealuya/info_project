<template>
    <div>
        <div>
            <NavBar :leftText="'返回'" :title="'差旅补助报销单'" :leftArrow="true" @leftEvent="back()"></NavBar>
        </div>
        <van-loading style="padding-top: 50px" align="center" v-if="loadingStatus"/>
        <div class="mt10" :style="{height:realHeight,overflowY:'scroll'}" v-if="loadingStatus==false">
            <van-row class="ml10" :style="{ background:'#FFFFFF',width:realWidth+'px'}">
                <div class="ml10 mt10">
                    <van-row>
                        <van-col span="16">
                            <van-row class="pt10">
                                <span>差旅补助报销单</span>
                            </van-row>
                            <van-row class="pt10 f14 spanCss">
                                <span> 申请单号：</span>
                            </van-row>
                            <van-row class="pt10 f14">
                                <span> {{travelDetails.cwxtTravelbx.applyNo}}</span>
                            </van-row>
                            <van-row class="pt10 f14 spanCss">
                                <span> 报销事由：</span>
                            </van-row>
                            <van-row class="pt10 f14 pb10">
                                <span> {{travelDetails.cwxtTravelbx.reason}} </span>
                            </van-row>
                            <van-row class="spanCss f14">
                                <span> 报销总额：</span>
                            </van-row>
                            <van-row class="pt10 pb10 fw600" style="font-size: 18px; color: #FF9200;">
                                <span> ￥{{travelDetails.cwxtTravelbx.bxAmount}} </span>
                            </van-row>
                        </van-col>
                        <van-col span="8">
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
            <van-row class="ml10 mt10 f14" :style="{ background:'#FFFFFF',width:realWidth+'px'}">
                <div class="ml10 mt10" :style="{width:realWidth-20+'px'}">
                    <van-row class="pt10">
                        <van-col span="6" class="spanCss">申请人员:</van-col>
                        <van-col span="18">{{travelDetails.cwxtTravelbx.empName}}</van-col>
                    </van-row>
                    <van-row class="pt10">
                        <van-col span="6" class="spanCss">申请时间:</van-col>
                        <van-col span="18">{{travelDetails.cwxtTravelbx.createTime}}</van-col>
                    </van-row>
                    <van-row class="pt10">
                        <van-col span="6" class="spanCss">申请部门:</van-col>
                        <van-col span="18">{{userInfoData.userInfo.deptName}}</van-col>
                    </van-row>
                    <van-row class="pt10">
                        <van-col span="6" class="spanCss">报销状态:</van-col>
                        <van-col span="18">{{travelDetails.cwxtTravelbx.bx==true?'已报销':'未报销'}}</van-col>
                    </van-row>
                    <van-row class="pt10">
                        <van-col span="6" class="spanCss">出发时间:</van-col>
                        <van-col span="18">{{travelDetails.cwxtTravelbx.startTime}}</van-col>
                    </van-row>
                    <van-row class="pt10">
                        <van-col span="6" class="spanCss">返回时间:</van-col>
                        <van-col span="18">{{travelDetails.cwxtTravelbx.returnTime}}</van-col>
                    </van-row>
                    <van-row class="pt10 pb10"
                             v-if="travelDetails.cwxtTravelbx.xmmc!=''||travelDetails.cwxtTravelbx.xmmc!=null">
                        <van-col span="6" class="spanCss">成本中心:</van-col>
                        <van-col span="18">{{travelDetails.cwxtTravelbx.xmmc+'('+travelDetails.cwxtTravelbx.xmbh+')'}}
                        </van-col>
                    </van-row>
                </div>
            </van-row>

            <van-row>
                <div :style="{width:realWidth+'px',background: '#FFFFFF',marginLeft:'10px',marginTop: '10px',minHeight:'30px'}">
                    <van-row class="f14 pl20 pt10 ">
                        补助信息
                    </van-row>
                    <van-row class="f14 pl20 pt10 pb10" v-for="i in needTypeList">
                        <div :style="{width:realWidth-30+'px',border:' #4BA3FB solid 1px',borderRadius:'5px',paddingBottom:'20px'}">
                            <van-row class="pl20 pt10">
                                <van-col span="8" style="font-weight: 600">{{i.typeName}}</van-col>
                                <van-col span="10"></van-col>
                                <van-col span="6" class="money">￥{{i.needsmoney}}</van-col>
                            </van-row>
                            <van-row class="pl20 pt10 f12">
                                <van-col span="2">
                                    <van-icon name="friends" size="18" color="#1989fa"/>
                                </van-col>
                                <van-col span="6">
                                    人数: {{i.peoplenum}}人
                                </van-col>
                                <van-col span="2">
                                    <!--                                    <van-icon name="notes" />-->
                                    <van-icon name="records" size="18" color="rgb(242 162 37)"/>
                                </van-col>
                                <van-col span="6">
                                    天数: {{i.days}}天
                                </van-col>
                            </van-row>
                            <van-row class="pl20 pt10 f12" style="color: #969799" justify="left">
                                标准补贴: ￥{{i.needsstandard}}/天
                            </van-row>
                        </div>
                    </van-row>
                </div>
            </van-row>

            <van-row v-if="travelDetails.hotelorderList.length>0">
                <div :style="{width:realWidth+'px',background: '#FFFFFF',marginBottom: '10px',marginLeft:'10px',marginTop: '10px',minHeight:'30px'}">
                    <van-row class="f14 pl20 pt10 ">
                        住宿费用信息
                    </van-row>
                    <van-row class="f14 pl20 pt10 pb10" v-for="i in travelDetails.hotelorderList">
                        <div :style="{width:realWidth-30+'px',border:' #4BA3FB solid 1px',borderRadius:'5px',paddingBottom:'20px'}">
                            <van-row class="pt10 pl10 fw600">
                                <van-col span="15">
                                    {{i.hotelname}}
                                </van-col>
                                <van-col span="3"></van-col>
                                <van-col span="6" class="money f16">
                                    {{i.amountpayable}}
                                </van-col>
                            </van-row>
                            <van-row class="pt10 pl10 f12">
                                <van-col span="5">订单编号 :</van-col>
                                <van-col span="19" class="fc_gary">{{i.orderno}}</van-col>
                            </van-row>
                            <van-row class="pt10 pl10 f12">
                                <van-col span="5">入住人员 :</van-col>
                                <van-col span="19" class="fc_gary">{{i.passengernames}}</van-col>
                            </van-row>
                            <van-row class="pt10 pl10 f12">
                                <van-col span="5">入住城市 :</van-col>
                                <van-col span="19" class="fc_gary">{{i.cityname}}</van-col>
                            </van-row>
                            <van-row class="pt10 pl10 f12">
                                <van-col span="5">离住时间 :</van-col>
                                <van-col span="19" class="fc_gary">{{i.gmtoccupancytime+' 至 '+ i.gmtleavetime}}
                                </van-col>
                            </van-row>
                            <van-row class="pt10 pl10 f12">
                                <van-col span="5">入住状态 :</van-col>
                                <van-col span="19" class="fc_gary">{{i.orderstatus=='CONFIRM'?'已确认':
                                    i.orderstatus=='CHECK_OUT'?'已离店':i.orderstatus=='CANCELED'?'已取消':i.orderstatus=='CHECK_IN'?'已入住':''}}
                                </van-col>
                            </van-row>

                        </div>
                    </van-row>
                </div>
            </van-row>
            <van-row v-if="travelList.length>0">
                <div :style="{width:realWidth+'px',background: '#FFFFFF',marginBottom: '10px',marginLeft:'10px',marginTop: '10px',minHeight:'30px'}">
                    <van-row class="f14 pl20 pt10 ">
                        出行费用信息
                    </van-row>
                    <van-row class="f14 pl20 pt10 pb10" v-for="i in travelList">
                        <div :style="{width:realWidth-30+'px',border:' #4BA3FB solid 1px',borderRadius:'5px',paddingBottom:'20px'}">
                            <van-row class="f16 pt10 pl10 fw600" justify="space-between">
                                <van-col span="9">
                                    {{i.startArea}}
                                </van-col>
                                <van-col span="4">
                                    <van-icon size="25" name="exchange"/>
                                </van-col>
                                <van-col span="9">
                                    {{i.endArea}}
                                </van-col>
                            </van-row>
                            <van-row class="pt10 pl10 f12">
                                <van-col span="5">出差人员 :</van-col>
                                <van-col span="19" class="fc_gary">{{i.traveler}}</van-col>
                            </van-row>
                            <van-row class="pt10 pl10 f12">
                                <van-col span="5">交通方式 :</van-col>
                                <van-col span="19" class="fc_gary">{{i.travelType}}</van-col>
                            </van-row>
                            <van-row class="pt10 pl10 f12">
                                <van-col span="5">出发时间 :</van-col>
                                <van-col span="19" class="fc_gary">{{i.startTime}}</van-col>
                            </van-row>
                            <van-row class="pt10 pl10 f12">
                                <van-col span="5">到达时间 :</van-col>
                                <van-col span="19" class="fc_gary">{{i.endTime}}
                                </van-col>
                            </van-row>

                        </div>
                    </van-row>
                </div>
            </van-row>
            <van-row v-if="invoiceInformation.length>0">
                <div :style="{width:realWidth+'px',background: '#FFFFFF',marginBottom: '5px',marginLeft:'10px',minHeight:'30px'}">
                    <van-row class="pt10 pl20 f14">
                        发票明细
                    </van-row>
                    <van-row class="f14 pl20 pt10">
                        <div v-if="invoiceInformation.length>0">
                            <div v-for="(item,index) in invoiceInformation"
                                 :style="{width:realWidth-30+'px',minHeight: '30px',background:'rgb(157 194 232 / 30%)',borderRadius:'3px',marginBottom:'5px'}">
                                <van-row justify="space-between" class="pl10 pt10 pb10">
                                    <van-col span="9" class="fw600 fc">{{ item.spendType }}</van-col>
                                    <van-col span="1">
                                        <van-icon v-if="item.btnStatus" name="arrow-down"
                                                  @click="showInvoice(item,index)"/>
                                        <van-icon v-if="!item.btnStatus" name="arrow"
                                                  @click="showInvoice(item,index)"/>
                                    </van-col>
                                    <van-col span="3"></van-col>
                                    <van-col span="9" align="right" class="fw600 f14" style="color: #0080FF">
                                        ￥{{ item.spendAmount }}
                                    </van-col>
                                    <van-col span="1">
                                    </van-col>
                                </van-row>
                                <div class="pl10 pt10 pb10" v-if="item.btnStatus">
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
                                        <van-col>{{ item.spendType }}</van-col>
                                    </van-row>
                                    <van-row justify="space-between" class="pr5 pt10">
                                        <van-col>开票时间</van-col>
                                        <van-col>{{ item.createtime }}</van-col>
                                    </van-row>
                                </div>
                            </div>
                        </div>
                    </van-row>
                </div>
            </van-row>

            <van-row class="ml10 mt10" :style="{ background:'#FFFFFF',width:realWidth+'px'}">
                <div class="ml10 mt10">
                    <van-row class="pt10">
                        <span>冻结记录</span>
                    </van-row>
                    <div style="min-height: 35px">
                        <van-steps direction="vertical" :active="0" v-for="(item,index) in freezingRecords">
                            <van-step>
                                <h>{{item.creationTime}}</h>
                                <div class="cardList" :style="{width:realWidth-60+'px'}">
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
            <van-row class="ml10 mt10 mb10" :style="{background:'#FFFFFF',width:realWidth+'px'}">
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
                                                            {{travelDetails.cwxtTravelbx.createTime}}
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
        getTravelBxDetail, getFreezingThawingHistory
    } from '../../services/reimburse'; // 导入需要用到的接口
    import {
        approvalHistory,
    } from "../../services/current"
    import waitshenpi_icon from "../../assets/img/waitshenpi_icon.png"
    import yichexiao_icon from "../../assets/img/yichexiao_icon.png"
    import yijujue_icon from "../../assets/img/yijujue_icon.png"
    import yitongyi_icon from "../../assets/img/yitongyi_icon.png"
    import Vector from "../../assets/icon/Vector.png"

    const invoiceInformation = ref([])
    let travelDetails = ref({})
    let userInfoData: any = inject("userInfo");
    let realHeight = (document.documentElement.offsetHeight - 60) + 'px';
    let realWidth = document.documentElement.offsetWidth - 20;
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
    const needTypeList = ref([])
    const travelList = ref([])

    onBeforeMount(() => {
        let applyNo = router.currentRoute.value.params.applyNo as never
        loadingStatus.value = true
        getTravelBxDetailMth(applyNo)
        getFreezingThawingMth(applyNo)
        getApprovalHistory(applyNo)
        setTimeout(() => {
            loadingStatus.value = false
        }, 500)
    })

    function back() {
        router.go(-1)
    }

    function showInvoice(item: any, index: any) {
        invoiceInformation.value[index].btnStatus = !item.btnStatus
    }

    //获取冻结记录
    async function getFreezingThawingMth(e: any) {
        let params = {
            corpCode: userInfoData.userInfo.corpCode, approveNo: e
        }
        let res: any = await getFreezingThawingHistory(params)
        if (res.code == 200) {
            freezingRecords.value = res.data
        }
    }

    //获取差旅补助详情
    async function getTravelBxDetailMth(e: any) {
        let param = {cwxtNo: e, corpCode: userInfoData.userInfo.corpCode, empCode: userInfoData.userInfo.empCode}
        let res: any = await getTravelBxDetail(param)
        if (res.code == 200) {
            travelDetails.value = res.data
            if (res.data.bxneedsinfoList.length > 0) {
                res.data.bxneedsinfoList.map((item: any) => {
                    item.typeName = ""
                    switch (item.type) {
                        case "food":
                            item.typeName = "伙食补助"
                        case "traffic":
                            item.typeName = "交通补助"
                    }
                })
            }
            //补助信息
            needTypeList.value = res.data.bxneedsinfoList
            if (res.data.flightorderList.length > 0) {
                res.data.flightorderList.map((item: any) => {
                    let obj = {
                        startArea: item.endcityname,
                        endArea: item.startcityname,
                        travelType: '飞机',
                        startTime: item.planbegindate,
                        endTime: item.planenddate,
                        traveler: item.passengernames
                    }
                    //出行费用信息
                    travelList.value.push(obj as never)
                })
            }
            if (res.data.trainorderList.length > 0) {
                res.data.trainorderList.map((item: any) => {
                    let obj = {
                        startArea: item.planbegincity,
                        endArea: item.planendcity,
                        travelType: '火车',
                        startTime: item.planbegindate,
                        endTime: item.planenddate,
                        traveler: item.pemployeename
                    }
                    //出行费用信息
                    travelList.value.push(obj as never)
                })
            }
            //发票信息
            res.data.invoiceDetail.map((item: any) => {
                item.btnStatus = false

            })
            invoiceInformation.value =  res.data.invoiceDetail
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
        margin-right: 10px;
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
        /*z-index: 9999;*/
        /*float: right;*/
    }

    .fc_gary {
        color: #969799
    }

</style>
