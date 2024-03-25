<template>
    <div :style="{minHeight:reaHeight+'px', 'background': '#FFFFFF'}">
        <van-row>
            <NavBar :title="'记账本'" :leftArrow="true" @leftEvent="back()"/>
        </van-row>
        <van-row style="display: flex;">
            <div class="iconSizeCss">
                <van-row type="flex" :gutter="8">
                    <van-col span="4">
                        <div :class="expenseType == 'C'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('C','用车')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_01">
                                <div class="f12 fc">用车</div>
                            </van-image>
                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'F'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('F','飞机')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_02">
                                <div class="f12 fc">飞机</div>
                            </van-image>
                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'TR'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('TR','公交')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_03">
                                <div class="f12 fc">公交</div>
                            </van-image>

                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'T'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('T','火车')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_04">
                                <div class="f12 fc">火车</div>
                            </van-image>
                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'RE'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('RE','餐饮')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_05">
                                <div class="f12 fc">餐饮</div>
                            </van-image>

                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'E'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('E','娱乐')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_06">
                                <div class="f12 fc">娱乐</div>
                            </van-image>

                        </div>
                    </van-col>
                </van-row>
                <van-row type="flex" :gutter="8">
                    <van-col span="4">
                        <div :class="expenseType == 'W'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('W','办公')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_07">
                                <div class="f12 fc">办公</div>
                            </van-image>
                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'M'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('M','会议记录')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_08">
                                <div class="f12 fc">会议记录</div>
                            </van-image>
                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'R'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('R','通讯')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_09">
                                <div class="f12 fc">通讯</div>
                            </van-image>

                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'H'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('H','酒店')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_10">
                                <div class="f12 fc">酒店</div>
                            </van-image>

                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'EX'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('EX','快递')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_11">
                                <div class="f12 fc">快递</div>
                            </van-image>
                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'FU'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('FU','燃油')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_12">
                                <div class="f12 fc">燃油</div>
                            </van-image>
                        </div>
                    </van-col>

                </van-row>
                <van-row type="flex" :gutter="8">
                    <van-col span="4">
                        <div :class="expenseType == 'A'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('A','档案')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_13">
                                <div class="f12 fc">档案</div>
                            </van-image>

                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'P'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('P','体检')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_14">
                                <div class="f12 fc">体检</div>
                            </van-image>

                        </div>
                    </van-col>
                    <van-col span="4">
                        <div :class="expenseType == 'V'? 'typeImageChecked':'typeImageUnChecked' "
                             @click="changeExpenseType('V','签证')">
                            <van-image :style="{width:iconWidth+'px',height:iconWidth+'px' }"
                                       :src="cost_icon_15">
                                <div class="f12 fc">签证</div>
                            </van-image>

                        </div>
                    </van-col>
                </van-row>
            </div>
        </van-row>
        <van-row>
            <van-form @submit="onSubmit" class="mt20 f16" :style="{width:realWidth+'px'}"  >
                    <van-field
                            v-model="money"
                            label="金额 (￥)"
                            placeholder="请输入金额"
                            :rules="[{ required: true, message: '请输入金额' }]"
                            class="lableCss"
                            maxlength="10"
                            type="number"
                            @input="formatter"
                            label-width="6em"

                    />
                    <van-field
                            v-model="dateTime"
                            is-link
                            readonly
                            name="calendar"
                            label="发生日期"
                            placeholder="点击选择日期"
                            @click="showCalendar = true"
                            label-width="6em"
                    />
                    <van-calendar v-model:show="showCalendar" :max-date="maxDate" :min-date="minDate" @confirm="onConfirm"/>

                    <van-field
                            v-model="reason"
                            label="费用事由"
                            placeholder="请输入费用事由"
                            :rules="[{ required: true, message: '请输入费用事由' }]"
                            maxlength="20"
                            label-width="6em"
                            :update:model-value="(reason = reason.replace(/\s+/g, ''))"

                    />
                    <van-field
                            v-model="locationName"
                            label="记录地点"
                            placeholder="请输入记录地点"
                            :rules="[{ required: false, message: '请输入记录地点' }]"
                            label-width="6em"

                    >
                    </van-field>
                    <van-field
                            v-model="invoiceMessage"
                            label="发票信息"
                            placeholder="请选择发票信息"
                            readonly="true"
                            input-align="right"
                            label-width="6em"

                    >
                        <template #input>
                            <van-icon name="add-o" color="#323233" size="18" @click="selectInvoiceMth"/>
                        </template>
                    </van-field>
                <van-row class="f14 pl20 pt10">
                    <div v-if="invoiceInformation.length>0">
                        <div v-for="(item,index) in invoiceInformation"
                             :style="{width:realWidth-30+'px',minHeight: '30px',background:'rgb(157 194 232 / 30%)',borderRadius:'3px',marginBottom:'5px'}">
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
                    <van-field name="uploader" label="凭证上传" label-width="6em">

                        <template #input>
                            <van-uploader v-model="fileList" :after-read="getFile"  :before-delete="deleteFile"
                                          max-count="4"/>
                        </template>
                    </van-field>
                <div style="margin: 16px;">
                    <van-button round block type="primary" native-type="submit" :disabled="disableState">
                        提交
                    </van-button>
                </div>
            </van-form>
        </van-row>
    </div>
    <!--  选择发票-->
    <van-popup v-model:show="selectInvoice" style="width:100vw;height:100vh" position="bottom">
        <select-invoice-picker @select-over="selectInvoiceEvent" ref="selectInvoiceOver" :select-list="invoiceInformation"/>
    </van-popup>
</template>

<script setup lang="ts">
    import {inject, onMounted, ref,nextTick} from 'vue'
    import router from "../../router/index";
    import moment from "moment";
    import {showToast, Loading, Notify} from 'vant'
    import {fetchLocation} from '../../services/utils'
    import cost_icon_01 from '../../assets/icon/cost_icon_01.png'
    import cost_icon_02 from '../../assets/icon/cost_icon_02.png'
    import cost_icon_03 from '../../assets/icon/cost_icon_03.png'
    import cost_icon_04 from '../../assets/icon/cost_icon_04.png'
    import cost_icon_05 from '../../assets/icon/cost_icon_05.png'
    import cost_icon_06 from '../../assets/icon/cost_icon_06.png'
    import cost_icon_07 from '../../assets/icon/cost_icon_07.png'
    import cost_icon_08 from '../../assets/icon/cost_icon_08.png'
    import cost_icon_09 from '../../assets/icon/cost_icon_09.png'
    import cost_icon_10 from '../../assets/icon/cost_icon_10.png'
    import cost_icon_11 from '../../assets/icon/cost_icon_11.png'
    import cost_icon_12 from '../../assets/icon/cost_icon_12.png'
    import cost_icon_13 from '../../assets/icon/cost_icon_13.png'
    import cost_icon_14 from '../../assets/icon/cost_icon_14.png'
    import cost_icon_15 from '../../assets/icon/cost_icon_15.png'
    import invoiceCloseIcon from '../../assets/icon/close_invoice.png'

    let userInfoData: any = inject("userInfo");
    import {addBillInfo} from '../../services/accountBook'
    import {UpLoadImg} from "../../components/hw_obs";
    const realWidth = window.document.documentElement.offsetWidth
    let iconWidth = (realWidth - 85) / 6 //图片宽高
    const reaHeight = window.document.documentElement.offsetHeight
    const fileList = ref([]);
    const fileListUploud = ref([]);
    let currentTime = new Date()
    const dateTime = ref(moment(currentTime).format('YYYY-MM-DD'))
    const showCalendar = ref(false);
    const onConfirm = (date: any) => {
        dateTime.value = moment(date).format("YYYY-MM-DD")
        showCalendar.value = false;
    };
    const locationName = ref("")
    const expenseType = ref("C")
    const expenseTypeName = ref("用车")
    const money:any= ref("")
    const invoiceMessage = ref([])
    const disableState = ref(false)
    let  maxDate = new Date(dateTime.value)
    let  minDate = new Date('1970-01-01')
    const locationData = ref("")
    const reason = ref("")
    const key=ref('');
    const selectInvoice = ref(false)
    const invoiceInformation = ref([])
    const selectInvoiceOver = ref(); // 获取组件实例
    onMounted(() => {
        getLocation()
        key.value = router.currentRoute.value.query.key as never;
    })

    function formatter(e : any) {
        if(e){
            e.target.value = e.target.value.replace(/[^\d.]/g, "");  //清除“数字”和“.”以外的字符
            e.target.value = e.target.value.replace(/\.{2,}/g, "."); //只保留第一个. 清除多余的
            e.target.value = e.target.value.replace(".", "$#$").replace(/\./g, "").replace("$#$", ".");
            e.target.value = e.target.value.replace(/^(\-)*(\d+)\.(\d\d).*$/, '$1$2.$3');//只能输入两个小数
            e.target.value = e.target.value.replace(/^\./g, ''); //首位不能输入“.”
            if (e.target.value.indexOf(".") < 0 && e.target.value != "") {//如果没有小数点，首位不能为0，如01、02...
                e.target.value = parseFloat(e.target.value);
            }
            nextTick(() => {
                money.value = e.target.value
            })
        }
    }

    const onSubmit = (values: any) => {
        if(money.value*1 <= 0){
            showToast('金额不能小于1');
            return
        }
        addBillMth()
    };

    function back() {
      if(key.value=='add'){
        router.go(-1);
        router.go(1);
      }else{
        router.push({name: 'home'});
      }
    }

    async function getLocation() {
        let res: any = await fetchLocation()
        if (res.info == "OK") {
            locationName.value = res.province + "-" + res.city
            if(!res.address){
                res.address = locationName.value
            }
            locationData.value = JSON.stringify(res)
        }
    }

    function addCostMessage() {

    }

    function changeExpenseType(e: any, n: any) {
        expenseType.value = e
        expenseTypeName.value = n

    }

    async function addBillMth() {
        let voucherList: any = []
        fileListUploud.value.map((item:any)=>{
            voucherList.push(item.fileName)
        })
        let invoiceParams:any = []
        invoiceInformation.value.map((item:any)=>{
            let invoiceObj = {
                invoiceType: item.invName,
                dzfph: item.dzfph,
                amount:item.zje,
                fpdm: item.fpdm,
                fphm:  item.fphm,
                fjid:  item.fjid
            }
            invoiceParams.push(invoiceObj)
        })
        let moneyValue=parseFloat(money.value).toFixed(2);
        let params = {
            empCode: userInfoData.userInfo.empCode,
            creater: userInfoData.userInfo.username,
            expenseType: expenseType.value,
            occurrenceTime: dateTime.value,
            money: moneyValue,
            expenseOwner: userInfoData.userInfo.username,
            startCity: "",
            arriveCity: "",
            startTime: "",
            endTime: "",
            reason: reason.value,
            hotelName: "",
            stayDays: "",
            billType: "CNY",
            tellList: [],
            expenseTypeName: expenseTypeName.value,
            unitPrice: "",
            invoice:invoiceParams,
            state: "0",
            corpCode: userInfoData.userInfo.corpCode,
            voucherList:voucherList,
            locationData:locationData.value
        }
        disableState.value = true
        let res: any = await addBillInfo(params)
        if (res.success) {
            setTimeout(() => {
                disableState.value = false
            }, 300)
            showToast('新增成功');
          if(key.value=='add'){
            router.go(-1);
          }else{
            router.push({name: 'home'});
          }
        }

    }

    const getFile = async (file: any) => {
        let obsResult = await UpLoadImg([file]);
        if (obsResult) {
            let index = obsResult[0].lastIndexOf('/')
            let name = ""
            if(index != -1){
                name = obsResult[0].substring(index+1,obsResult[0].length)
            }
            fileListUploud.value.push({
                name: file.file.name,
                url: obsResult[0],
                fileName:name
            } as never)
        }
    }

    const deleteFile = async (file: any,index: any) => {
        fileList.value.splice(index.index,1)
        fileListUploud.value.splice(index.index,1)
    }

    // 点击打开发票数组
    function selectInvoiceMth() {
        selectInvoice.value = true
        setTimeout(()=> selectInvoiceOver.value.onRefresh('1'),500)
    }

    function selectInvoiceEvent(e: any) {
        if(e.length>0){
            let amount = 0
            e.map((item: any) => {
                item.btnStatus = false
                amount = amount + item.zje*1
            })
            invoiceInformation.value = e
            money.value = amount as never
        }
        // getTotalReimbursementAmount()
        selectInvoice.value = false
    }

    function showInvoice(item: any, index: any) {
        invoiceInformation.value[index].btnStatus = !item.btnStatus
    }

    function removeInvoice(item: any, index: any) {
        let amount = 0
        invoiceInformation.value.splice(index, 1)
        invoiceInformation.value.map((item: any) => {
            item.btnStatus = false
            amount = amount + item.zje*1
        })
        money.value = amount as never
        // getTotalReimbursementAmount()
    }



</script>

<style lang="less" scoped>
    .iconSizeCss .van-image {
        font-size: 8px;
        text-align: center;
    }

    .iconSizeCss {
        padding-top: 15px;
        padding-left: 10px;
        text-align: center;
        justify-content: space-between;
    }

    /deep/ .iconSizeCss .van-row {
        margin-top: 5px;
    }

    /deep/ .lableCss .van-field__label {
        /*text-align: right;*/
        font-weight: 800 !important;
        font-size: 14px !important;
    }

    .van-cell {
        position: initial;
    }

    .typeImageChecked {
        background: #0080ff2b;
        border: 1px solid #067df3f0;
        border-radius: 5px;
        margin-bottom: 11px;
        /*color: #fff;*/
    }

    .typeImageUnChecked {
        border: 1px solid transparent;
        border-radius: 5px;
        margin-bottom: 11px;
    }
    /*.van-field .van-field__label{*/
    /*    width: 100px;*/
    /*}*/

     .van-cell:after{
         border:none;
     }

    .fc{
        color: #1a1a1a;
    }

</style>
