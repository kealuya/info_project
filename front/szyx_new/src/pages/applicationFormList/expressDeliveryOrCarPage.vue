<template>
    <div>
        <div>
            <NavBar :leftText="'返回'" :leftArrow="true" :title="bussinessTypeName" @leftEvent="back()">
            </NavBar>
            <van-notice-bar color="#1989fa" background="#ecf9ff" left-icon="info-o">
                {{noticeBarMsg}}
            </van-notice-bar>
        </div>
        <div>

            <van-pull-refresh v-model="refreshing" @refresh="onRefresh"
                              :style="{'height':realHeight+'px','overflow':'scroll'}">
                <van-list
                        v-model:loading="loading"
                        :finished="finished"
                        finished-text="没有更多了"
                        offset="100"
                        @load="onLoad"
                >
                    <van-cell v-for="item in dataList" :key="item.approveNo" class="mt10"
                              style="width: 100%; display: flex"  >
                        <div class="pr5 fc" @click="queryApprovalDetail(item)">
                            <van-row justify="space-between">
                                <van-col span="8" align="left">
                                    <van-row>
                                        <van-image :src="iconNameUrl" width="25" height="25"
                                                   style="padding-right: 10px;padding-bottom: 5px"></van-image>
                                        <div>{{approveTitle}}</div>
                                    </van-row>
                                </van-col>
                                <van-col span="10"></van-col>
                                <van-col span="6" class="vanTagCss">
                                    <van-tag type="warning" color="rgb(255 151 106 / 19%)" size="large" text-color="#ff976a"
                                             v-if="item.approvalState == '已撤销'">审批撤销
                                    </van-tag>
                                    <van-tag type="primary" color="#1989fa47" text-color="#1989fa" size="large"
                                             v-if="item.approvalState == '待审批'">等待审批
                                    </van-tag>
                                    <van-tag type="danger" color="#ee0a2424" text-color="rgb(238 10 36 / 86%)" size="large"
                                             v-if="item.approvalState == '已拒绝'">申请拒绝
                                    </van-tag>
                                    <van-tag type="success" color="rgb(7 193 96 / 20%)" text-color="rgb(7 193 96 / 83%)" size="large"
                                             v-if="item.approvalState == '已通过'">审批通过
                                    </van-tag>
                                </van-col>
                            </van-row>
                            <van-row justify="space-between">
                                <div class="ellipsis">申请事由: {{item.approveReason}}</div>
                            </van-row>
                            <van-row justify="space-between">
                                <span>申请人:{{item.creater}}</span>
                            </van-row>
                            <van-row justify="space-between">
                                <span>申请单号:{{item.approveNo}}</span>
                            </van-row>
                            <van-row justify="space-between"
                                     v-if="bussinessTypeName == '用车' && item.approvalState == '已通过' && item.deadline">
                                <span>用车截止时间: 于{{item.deadline}}之前</span>
                            </van-row>
                            <van-row justify="space-between" class="graytxt">
                                <span>{{item.createTime}}</span>
                            </van-row>

                        </div>
                        <van-row justify="space-between" v-if="item.approvalState == '已通过'">
                            <van-button type="primary" size="small"
                                        @click="applicationType==='car'? goTakeTaxi(item):expressDelivery(item)"
                                        :icon="btnIcon">
                                {{btnName}}
                            </van-button>
                        </van-row>
                    </van-cell>
                </van-list>
            </van-pull-refresh>
        </div>
    </div>
</template>

<script setup lang="ts">
    import {inject, ref, onMounted} from 'vue'
    import {useRouter} from "vue-router";
    const router = useRouter(); // 路由
    import {getWuliuLoginUrl, getCarLoginUrl} from "../../services/trirdPartyBusiness"
    import {getAllApplyList} from "../../services/apply"
    import application_icon_05 from '../../assets/icon/application_icon_05.png'
    import application_icon_15 from '../../assets/icon/application_icon_15.png'
    import icon_20 from '../../assets/icon/icon_20.png'
    import {showFailToast} from "vant/lib/toast/function-call";
    let userInfo: any = inject("userInfo");
    const realHeight = window.document.documentElement.offsetHeight - 90
    const dataList = ref([]);
    const loading = ref(false);
    const finished = ref(false);
    const refreshing = ref(false);
    const total = ref(0)
    const page = ref(0)
    const pageSize = ref(8)
    const iconNameUrl = ref("")
    const applicationType = ref("")
    const bussinessTypeName = ref("")
    const approveTitle = ref("")
    const btnName = ref("")
    const btnIcon = ref()
    const hrefUrl = ref("")
    const noticeBarMsg = ref("")

    onMounted(() => {
        applicationType.value = router.currentRoute.value.params.applicationType as never
        switch (applicationType.value) {
            case "car":
                bussinessTypeName.value = '用车'
                iconNameUrl.value = application_icon_15
                btnName.value = "去打车"
                approveTitle.value = '用车申请'
                noticeBarMsg.value = '审批通过后在设定的用车时间内进行用车，超过用车时间后将无法使用'
                break;
            case "log":
                bussinessTypeName.value = '快递'
                iconNameUrl.value = application_icon_05
                btnName.value = "寄快递"
                approveTitle.value = '快递物流'
                btnIcon.value = icon_20
                noticeBarMsg.value = '通过审批后再规定时间内进行寄送，超过寄送日期将无法使用'
                break;
        }
    })

    function onLoad() {
        page.value = page.value + 1;
        loading.value = true
        getDataMth()
    }

    const onRefresh = () => {
        // 清空列表数据
        finished.value = false;
        loading.value = true
        page.value = 0
        onLoad();
    };

    async function getDataMth() {
        let list = dataList.value
        setTimeout(async () => {
            if (refreshing.value) {
                list = []
                refreshing.value = false
            }
            let param = {
                corpCode: userInfo.userInfo.corpCode,
                page: page.value,
                pageSize: pageSize.value,
                keyWord: "",
                empCode: userInfo.userInfo.empCode,
                approveType: [applicationType.value],
                approveReason: "",
                createTime: "",
                state: "全部",
                approvalState: "全部"
            }
            let res: any = await getAllApplyList(param)
            if (res.success) {
                res.data.list.map((ite: any, index: any) => {
                    ite.locationData = ite.locationData ? JSON.parse(ite.locationData) : {}
                    if (ite.rowShowJson != undefined && ite.rowShowJson != "") {
                        try {
                            let list = JSON.parse(ite.rowShowJson)
                            list.map((i: any) => {
                                if (i.label == '返回时间') {
                                    ite.deadline = i.value
                                }
                            })
                        } catch (e:any) {
                          // showFailToast(e.msg);
                        }
                    }
                    list.push(ite as never)
                })
                loading.value = false;
                total.value = res.data.total

                if (list.length >= total.value) {
                    finished.value = true;
                }
            } else {
                list = []
            }
            dataList.value = list
        }, 1000)
    }


    //TODO 跳转审批详情(跳转不同详情页)
    function queryApprovalDetail(e: any) {
            const data :any ={
                'approveType':e.approveType,
                'approveNo':e.approveNo,
            };
            sessionStorage.setItem('applicationDetail', JSON.stringify(data))
            router.push('/applicationDetail')
    }

    function back() {
        router.back();
    }

    async function goTakeTaxi(e: any) {
        let param = {
            empCode: userInfo.userInfo.empCode,
            corpCode: userInfo.userInfo.corpCode,
            approveNo: e.approveNo,
            approveType: 'car',
        };
        let res: any = await getCarLoginUrl(param)
        if (res.success) {
            router.push({
                name: 'IframePage',
                params: {
                    barTitle: '用车',
                    iframeUrl: res.data.url
                }
            })
        }
    }

    async function expressDelivery(e: any) {
        let param = {
            empCode: userInfo.userInfo.empCode,
            corpCode: userInfo.userInfo.corpCode,
            approveNo: e.approveNo,
            approveType: 'log',
        };
        let res: any = await getWuliuLoginUrl(param)
        if (res.success) {
            router.push({
                name: 'IframePage',
                params: {
                    barTitle: '快递',
                    iframeUrl: res.data.url
                }
            })
        }

    }


    function getYwDetailMth() {


    }

</script>

<style lang="less" scoped>
    .fc {
        color: #1a1a1a;
    }

    .fw {
        font-weight: 800;
    }

    .iconCss {
        width: 45px;
        height: 45px;
        /*float: left;*/
        /*margin-left: 20%;*/
    }

    .ellipsis {
        width: 80%; /*宽度一定要设置*/
        overflow: hidden; /*文本超出隐藏*/
        text-overflow: ellipsis; /*文本超出显示省略号*/
        white-space: nowrap; /*超出的空白区域不换行*/
        text-align: left
    }

    .moneyColor {
        color: #0080FF;
    }

    .delete-button {
        height: 100%;
    }

    .vanTagCss .van-tag {
        --van-tag-radius: 4px;
        --van-tag-padding: 2px 8px 2px 8px;

    }

    .grayTitleCss {
        font-size: 6px;
        color: #898989;
        margin-top: 10px;
        font-weight: 400;
        margin-left: 5px;
    }

    .btnCss .van-button {
        margin-left: 5px;
        margin-top: 10px;
        width: 90px;
    }

    .fc {
        color: #1a1a1a;
    }

</style>
