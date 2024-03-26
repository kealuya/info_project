<template>
    <div>
        <div>
            <NavBar :leftText="'返回'" :title="'订单详情'" :leftArrow="true" @leftEvent="back()" :rightText="'确定'"
                    @rightEvent="okBtn()"></NavBar>
        </div>
        <div>
            <van-grid :column-num="3" class="f14">
                <van-grid-item>
                    <van-dropdown-menu class="dropdownCss">
                        <van-dropdown-item :title="queryLableTitle" v-model="typeName" :options="typeList"
                                           @change="changeTypeMth()"/>
                    </van-dropdown-menu>
                </van-grid-item>
                <van-grid-item @click="showCalenderVisiable">
                    <van-icon name="clock-o">{{queryTimeLable}}</van-icon>
                </van-grid-item>
                <van-grid-item @click="showRight = true">
                    <van-icon name="wap-nav">筛选</van-icon>
                </van-grid-item>
            </van-grid>
            <div class="pl10" style="padding-top: 10px">
                <van-pull-refresh v-model="refreshing" @refresh="onRefresh"
                                  :style="{'height':realHeight+'px','overflow':'scroll'}">
                    <van-list  v-model:loading="loading" :finished="finished" finished-text="没有更多了" offset="100" >
                        <van-checkbox-group v-model="checkedList" shape="square" class="pl20" v-for="(item) in dataList">
                            <van-checkbox :name="item.id" class="mt10">
                                <div :style="{width:realWidth-80+'px', border:'#4BA3FB solid 1px',borderRadius:'5px',minHeight:'30px'}">
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
                            </van-checkbox>
                        </van-checkbox-group>
                    </van-list>
                </van-pull-refresh>
            </div>
        </div>

        <van-popup v-model:show="showCalendar" type="range" closeable>
            <van-calendar
                    title="时间"
                    :poppable="false"
                    type="range"
                    :min-Date="minDateTime"
                    :max-date="maxDateTime"
                    :show-confirm="false"
                    :style="{ height: '500px' }"
                    :default-date="defauleDate"
                    @confirm="onConfirm"
            />
            <van-row style="text-align: center" class="pb10">
                <van-col span="12">
                    <van-button type="primary" size="small" style="width: 75px" @click="submitDateFilterMth">确定</van-button>
                </van-col>
                <van-col span="12">
                    <van-button type="default" size="small" style="width: 75px" @click="restDateFilterMth">重置</van-button>
                </van-col>
            </van-row>
        </van-popup>

        <van-popup v-model:show="showRight" position="right" :style="{ width: '60%', height: '100%' }">
            <div class="pt10 pl10">
                <van-row class="f16">
                    <span>单据搜索</span>
                </van-row>
                <van-row>
                    <van-field
                            v-model="keyWord"
                            name="keyWord"
                            label="出行地点: "
                            placeholder="请输入出行地点"
                            class="f14 mt10 mb10 filterInput"
                            label-width="7.5em"
                            label-align="top"
                            maxlength="30"
                    />
                </van-row>
                <van-row justify="center">
                    <van-col span="12" class="pl10">
                        <van-button type="primary" style="width: 70px" size="small" @click="filterMth">确定</van-button>
                    </van-col>
                    <van-col span="12">
                        <van-button type="primary" style="width: 70px" size="small" @click="reset">重置</van-button>
                    </van-col>
                </van-row>
            </div>

        </van-popup>
    </div>
</template>

<script setup lang="ts">

    import router from "../../../router";
    import {ref, inject, onMounted, defineProps, defineEmits} from 'vue'
    import moment from 'moment'
    import {getLeaderBxInfo} from '../../../services/reimburse'; // 导入需要用到的接口
    import {getDefaultPay} from '../../../services/manageAccounts'
    import {showFailToast, showSuccessToast, showToast} from "vant";
    import jiudian from "../../../assets/icon/jiudian.png"
    import feiji from "../../../assets/icon/feiji.png"
    import huoche from "../../../assets/icon/huoche.png"

    let userInfoData: any = inject("userInfo");
    let realWidth = window.document.documentElement.offsetWidth
    const realHeight = window.document.documentElement.offsetHeight - 120
    //类型list
    const typeList = [
        {text: '全部', value: ''},
        {text: '火车', value: 'T'},
        {text: '机票', value: 'F'},
        {text: '酒店', value: 'H'},
    ]
    const checkedList = ref([])
    let startDate = moment().subtract(3, "year").format("YYYY-MM-DD")
    let endDate = moment().add(3, "year").format("YYYY-MM-DD")
    const minDateTime = ref(new Date(startDate))
    const maxDateTime = ref(new Date(endDate))
    const dataList = ref([])
    const total = ref(0)
    //查询条件参数------start
    const page = ref(0)
    const pageSize = ref(8)
    const keyWord = ref("")
    const startTime = ref("")
    const endTime = ref("")
    const queryLableTitle = ref("类型")
    const queryTimeLable = ref("订单时间")
    const queryStartDate = ref("")
    const queryEndDate = ref("")
    const typeName = ref("")
    //查询条件参数-------end
    const defauleDate = ref([])
    const showCalendar = ref(false)
    const showRight = ref(false)
    const loading = ref(false);
    const finished = ref(false);
    const refreshing = ref(false);

    //选中的数据
    const props = defineProps({
        checked: {
            type: Array
        },
    })
    const emit = defineEmits(["close", "selectOver"])
    //初始化数据方法
    defineExpose({
        getDataList
    });

    onMounted(() => {
        getDataList()
    })

    //初始化数据，由于checked状态需要回显，故加载全部数据
    function getDataList() {
        page.value = 1
        refreshing.value = true
        getLeaderBxInfoMth().then(() => {
            checkedMth()
        })
    }

    //获取订单信息
    async function getLeaderBxInfoMth() {
        let list = dataList.value
        setTimeout(async () => {
            if (refreshing.value) {
                list = []
                refreshing.value = false
            }
            let params = {
                empCode: userInfoData.userInfo.empCode,
                corpCode: userInfoData.userInfo.corpCode,
                keyWord: keyWord.value ? keyWord.value : '',
                startDate: startTime.value ? startTime.value : "",
                endDate: endTime.value ? endTime.value : "",
                type: typeName.value,
                page: page.value,
                pageSize: pageSize.value,
            }
            let res: any = await getLeaderBxInfo(params)
            if (res.code == 200) {
                res.data.records.map((item: any) => {
                    list.push(item as never)
                })
                loading.value = false;
                total.value = res.data.total
            }
            dataList.value = list
            if (dataList.value.length >= total.value) {
                finished.value = true;
            } else {
                page.value = page.value + 1
                getLeaderBxInfoMth()
            }
        }, 1000)

    }
    //回显选中的数据
    function checkedMth() {
        let checkList = props.checked
        checkedList.value = checkList as never
    }

    const onRefresh = () => {
        // 清空列表数据
        finished.value = false;
        loading.value = true
        refreshing.value = true
        getDataList();
    };

    //显示时间选择器
    function showCalenderVisiable() {
        let defaultDateArr = []
        if (queryStartDate.value != '' && queryEndDate.value != '') {
            let startDt = new Date(queryStartDate.value + ' 00:00:00')
            let endDt = new Date(queryEndDate.value + ' 23:59:59')
            defaultDateArr[0] = startDt
            defaultDateArr[1] = endDt
        }
        defauleDate.value = defaultDateArr as never
        showCalendar.value = true
    }

    //选择时间
    function onConfirm(e: any) {
        if (e.length > 0) {
            startTime.value = moment(e[0]).format("YYYY-MM-DD")
            endTime.value = moment(e[1]).format("YYYY-MM-DD")
        } else {
            startTime.value = ''
            endTime.value = ''
        }
    }

    //时间筛选方法
    function submitDateFilterMth() {
        queryStartDate.value = startTime.value
        queryEndDate.value = endTime.value
        queryTimeLable.value = startTime.value + " 至 " + endTime.value
        page.value = 1
        refreshing.value = true
        getLeaderBxInfoMth()
        showCalendar.value = false
    }

    //重置时间筛选方法
    function restDateFilterMth() {
        startTime.value = ''
        endTime.value = ''
        queryStartDate.value = ''
        queryEndDate.value = ''
        queryTimeLable.value = "订单时间"
        page.value = 1
        refreshing.value = true
        getLeaderBxInfoMth()
        showCalendar.value = false
    }

    // 出行地点筛选方法
    function filterMth() {
        page.value = 1
        refreshing.value = true
        getLeaderBxInfoMth()
        showRight.value = false
    }

    // 重置出行地点筛选方法
    function reset() {
        keyWord.value = ""
        page.value = 1
        refreshing.value = true
        getLeaderBxInfoMth()
        showRight.value = false
    }

    //类型筛选
    function changeTypeMth() {
        typeList.map((item: any) => {
            item.value == typeName.value ? queryLableTitle.value = item.text : '类型'
        })
        page.value = 1
        refreshing.value = true
        getLeaderBxInfoMth()
    }

    //订单选择确定方法
    function okBtn() {
        let set = new Set(checkedList.value)
        checkedList.value = Array.from(set)
        let list: any = []
        checkedList.value.map((item: any) => {
            dataList.value.map((ite: any) => {
                if (item == ite.id) {
                    list.push(ite)
                }
            })
        })
        emit("selectOver", list)
        queryLableTitle.value = "类型"
        queryTimeLable.value = "订单时间"
        keyWord.value = ""
        startTime.value = ""
        endTime.value = ""
        page.value = 0
    }

    //返回
    function back() {
        queryLableTitle.value = "类型"
        queryTimeLable.value = "订单时间"
        keyWord.value = ""
        page.value = 0
        pageSize.value = 8
        startTime.value = ""
        endTime.value = ""
        checkedList.value = []
        emit("close")
    }

</script>
<style lang="less" scoped>
    .dropdownCss {
        --van-dropdown-menu-shadow: 0;
        --van-dropdown-menu-height: 18px; /*height: 20px;*/
        font-size: 14px;
    }

    .fw600 {
        font-weight: 600;
    }

    .fc {
        color: #1a1a1a;
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

    /deep/ .stepsCss .van-step--vertical:not(:last-child):after {
        border-bottom-width: 0px;
    }

    /deep/ .filterInput .van-cell__value {
        border: #96979978 solid 1px;
        padding: 5px;
    }

</style>
