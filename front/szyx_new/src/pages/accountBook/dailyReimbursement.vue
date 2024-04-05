<template>
    <div>
        <van-row>
            <NavBar :title="'日常报销单'" :leftText="'返回'" :leftArrow="true" @leftEvent="backMth()"/>
        </van-row>
            <van-form @submit="onSubmit" style="display: table-cell;vertical-align: middle;background: #F2F2F6">
                <van-cell-group  center :style="{width:realWidth+'px',marginTop:'10px'}">
                    <div class="pl10 pt10 pb10">
                        <van-row class="font_gray f12">
                            <van-col span="5">{{businessId}}</van-col>
                            <van-col span="12"></van-col>
                            <van-col span="7"> {{timeDate}}</van-col>
                        </van-row>
                        <van-row class="pt10 f14">
                            <van-col span="5">{{userInfo.userInfo.userName}}</van-col>
                            <van-col span="12"></van-col>
                            <van-col span="7">{{userInfo.userInfo.deptName}}</van-col>
                        </van-row>
                        <van-row class="pt10 f14">
                            <van-col span="10" class="fw800">报销类型</van-col>
                            <van-col span="14">
                                <van-radio-group v-model="bxType" direction="horizontal">
                                    <van-radio name="1" icon-size="16px">冲借款</van-radio>
                                    <van-radio name="2" icon-size="16px">日常</van-radio>
                                </van-radio-group>
                            </van-col>
                        </van-row>
                    </div>
                </van-cell-group>
                <van-cell-group  center :style="{width:realWidth+'px',marginTop:'10px'}">
                    <div class="pl10 pt10 pb10">
                        <van-row class="pt5 f14 fw800 ">
                            <van-col span="22">
                                费用信息
                                <span style="color: red">*</span>
                            </van-col>
                            <van-col span="2">
                                <van-icon name="add-o" color="#323233" size="18" @click="addCostMessage"/>
                            </van-col>
                        </van-row>
                        <van-row style="margin-top: 8px;margin-bottom: 8px">
                            <div v-for="(item,index) in cardList" >
                                    <div class="cardList" :style="{width:cardWidth+'px'}">
                                        <van-badge  position="top-right" >
                                            <div :style="{width:cardWidth-10+'px',minHeight:'30px',marginTop:'5px',marginLeft:'10px'}">
                                                <van-row>
                                                    <van-col span="17" class="fw800 f14 pt5">
                                                        {{item.expenseTypeName}}
                                                    </van-col>
                                                    <van-col span="7" class="money">
                                                        ￥{{item.money}}
                                                    </van-col>
                                                </van-row>
                                                <van-row justify="left" class="pt10 pb10" >
                                                    <van-icon name="balance-list" color="#1989fa" class="mt3"/>
                                                    <div class="f14 pl5"> {{item.expenseTypeName}}</div>
                                                    <van-icon name="clock" color="#FB5632" class="pl5 mt3" />
                                                    <div class="f14 pl5 mt3" >{{item.createTime}}</div>
                                                </van-row>
                                            </div>
                                            <template  #content >
                                                <van-icon name="cross" class="badge-icon" @click="removeItem(item)"/>
                                            </template>
                                        </van-badge>
                                    </div>
                                    </div>

                        </van-row>
                    </div>
                </van-cell-group>
                <van-cell-group  center :style="{width:realWidth+'px',marginTop:'10px'}">
                    <div class="pl10 pt10 pb10">
                        <van-row class="pt5 f14 fw800 ">
                            <van-col span="22">
                                发票信息
                                <span style="color: red">*</span>
                            </van-col>
                            <van-col span="2">
                                <van-icon name="add-o" color="#323233" size="18" @click="addInvoice()"/>
                            </van-col>
                        </van-row>
                        <van-row style="margin-top: 8px;margin-bottom: 8px">
                            <div v-for="(item,index) in invoiceList" >
                                <div class="cardList" :style="{width:cardWidth+'px'}">
                                    <van-badge  position="top-right" >
                                        <div :style="{width:cardWidth-10+'px',minHeight:'30px',marginTop:'5px',marginLeft:'10px'}">
                                            <van-row>
                                                <van-col span="17" class="fw800 f14 pt5">
                                                    {{item.expenseTypeName}}
                                                </van-col>
                                                <van-col span="7" class="money">
                                                    ￥{{item.money}}
                                                </van-col>
                                            </van-row>
                                            <van-row justify="left" class="pt10 pb10" >
                                                <van-icon name="balance-list" color="#1989fa" class="mt3"/>
                                                <div class="f14 pl5"> {{item.expenseTypeName}}</div>
                                                <van-icon name="clock" color="#FB5632" class="pl5 mt3" />
                                                <div class="f14 pl5 mt3" >{{item.createTime}}</div>
                                            </van-row>
                                        </div>
                                        <template  #content >
                                            <van-icon name="cross" class="badge-icon" @click="removeItem(item)"/>
                                        </template>
                                    </van-badge>
                                </div>
                            </div>

                        </van-row>

                    </div>
                </van-cell-group>
                <van-cell-group  center :style="{width:realWidth+'px',marginTop:'10px'}">
                    <div class="pl10 pt10 pb10">
                        <van-row class="pt5 f14 fw800 ">
                            <van-col span="22">
                                成本中心
                                <span style="color: red">*</span>
                            </van-col>

                        </van-row>

                    </div>
                </van-cell-group>
                <van-cell-group  center :style="{width:realWidth+'px',marginTop:'10px'}">
                    <div class="pl10 pt10 pb10">
                        <van-row class="pt5 f14 fw800 ">
                            <van-col span="22">
                                部门领导审批流程
                                <span style="color: red">*</span>
                            </van-col>

                        </van-row>

                    </div>
                </van-cell-group>
                <van-cell-group  center :style="{width:realWidth+'px',marginTop:'10px'}">
                    <div class="pl10 pt10 pb10">
                        <van-row class="pt5 f14 fw800 ">
                            <van-col span="22">
                                成本中心
                                <span style="color: red">*</span>
                            </van-col>

                        </van-row>

                    </div>
                </van-cell-group>
                <van-cell-group  center :style="{width:realWidth+'px',marginTop:'10px'}">
                    <div class="pl10 pt10 pb10">
                        <van-row class="pt5 f14 fw800 ">
                            <van-col span="22">
                                收款信息(账户信息)
                                <span style="color: red">*</span>
                            </van-col>

                        </van-row>

                    </div>
                </van-cell-group>

            </van-form>
    </div>
    <div>
        <CostInformationSelectionPage></CostInformationSelectionPage>
    </div>
</template>

<script setup lang="ts">
    import {ref, onMounted, onBeforeMount, watch, inject} from 'vue'
    import route from "../../router";
    import {getInvoiceBillList,getApprovalFlow,getInvoiceWbList} from "../../services/accountBook"
    import moment from "moment"


    const formDataObj = ref({})
    let realWidth = window.document.documentElement.offsetWidth
    let cardWidth = window.document.documentElement.offsetWidth-30

    const bxType = ref("2")
    const username = ref('')
    const cardList = ref([])
    let userInfo: any = inject("userInfo");
    let businessId = 'RL'+ moment(new Date()).format('YYYYMMDDHHmmss')+ Math.round(Math.random()*10)
    let timeDate = moment(new Date()).format("YYYY-MM-DD")
    const invoiceList = ref([])

    watch(
        () => route.currentRoute.value,
        (newValue: any) => {
            if(newValue.name == 'dailyReimbursement'){
                let bookAccountDetail = JSON.parse(route.currentRoute.value.params.bill as never)
                if(bookAccountDetail.identifying=='costInformation'){
                    cardList.value = bookAccountDetail.bill
                }
                if(bookAccountDetail.identifying=='accountBook'){
                    formDataObj.value = bookAccountDetail.bill
                }

            }

            // if(newValue.name == 'dailyReimbursement' && newValue.params.bill && newValue.params.bill != '' ){
            //    let messageDetail = JSON.parse(newValue.params.bill)
            //     // if(messageDetail.routeName == 'cosInformationSelectionPage'){
            //
            //         cardList.value = messageDetail
            //         // cardList.value.push(messageDetail as never)
            //     // }
            // }

        },
        { immediate: true }
    )

    // onBeforeMount(() => {
    //     formDataObj.value = route.currentRoute.value.params.documentInformation
    //
    // })
    onMounted(()=>{
        //
        // formDataObj.value = userInfo

    })

    function addInvoice(){
        let invoice = {
            invoiceList:invoiceList
        }
        route.push({name:'invoiceInformationList',params:{
                invoice : JSON.stringify(invoice)
            }})
    }





    async function getInvoiceBillListMth() {
        let params = {
            corpCode:userInfo.corpCode,
            billId:[route.currentRoute.value.params.billId],
            empCode:userInfo.empCode
           }
        let res:any = await getInvoiceBillList(params)
        if(res.success){
            // formDataObj.value = res.data
        }
    }

    function removeItem(item: any) {
       let index = cardList.value.indexOf(item as never)
        cardList.value.splice(index,1)
    }

    function backMth() {
        route.push('/accountBook')
    }

    function onSubmit() {
    }
    function addCostMessage(){
        route.push({name:'costInformationSelectionPage',
        params:{
            bill:JSON.stringify(cardList.value)
        }})
    }
</script>

<style lang="less" scoped>

    .mlp4 {
        margin-left: 4%;
    }

    /*.mrp5{*/
    /*    margin-left: 5%;*/
    /*}*/
    /*.divBg{*/
    /*    background: #FFFFFF;*/
    /*}*/
    /*.widthP{*/
    /*    width: 90%;*/
    /*}*/
    .grayCss {
        font-size: 8px;
        color: #898989;
        /*margin-top: 10px;*/
        /*font-weight: 400;*/
        margin-left: 5px;
    }

    .fw800 {
        font-weight: 800;
    }

    .cardList{
        /*margin-left: 2%;*/
        /*margin-right: 2%;*/
        /*width: 200px;*/
        border: 1px solid #1989fa;
        min-height: 30px;
        border-radius:3px;
        margin-top: 10px;
    }
    /*.badge-icon {*/
    /*    display: block;*/
    /*    font-size: 10px;*/
    /*    line-height: 16px;*/
    /*}*/
    .mt3{
        margin-top: 3px;

    }

</style>
