<template>
    <div>
        <div>
            <NavBar :leftText="'返回'" :title="'我的委托'" :leftArrow="true" slortName="right" iconName="plus" iconText="添加" @iconEvent="addMyDelegation" @leftEvent="back()"></NavBar>
        </div>
        <div style="margin-top: -1px">
            <van-tabs v-model:active="activeName" background="#0080FF" color="#ffff" title-inactive-color="#ffff" title-active-color="#ffff">
                <van-tab title="委托设定" name="0">
                    <div style="background: #FFFFFF;min-height: 40px">
                        <van-notice-bar color="#1989fa" :scrollable="false" wrapable background="#ecf9ff"
                                        left-icon="info-o">
                            将审批权限委托给另一个用户，另一个用户可通过该模块对委托者业务进行审批
                        </van-notice-bar>
                    </div>
                    <div :style="{height:clientHeight+'px',overflow:'scroll'}">
                        <van-pull-refresh v-model="refreshing" @refresh="onDelegateListRefresh"
                                          :style="{minHeight:clientHeight+'px'}">
                            <van-list
                                    v-model:loading="delegateLoading"
                                    :finished="delegateFinished"
                                    finished-text="没有更多了"
                                    offset="100"
                                    @load="onLoad"
                            >
                                <van-cell v-if="delegateList.length>0" v-for="item in delegateList" :key="item"
                                          @click.stop="queryDelegateDetail(item)"
                                          class="mt10">
                                    <div class="cardCss">
                                        <van-row class="f16 ">
                                            <van-col class="fc fw ta-l" span="5">
                                                <span>审批权限</span>
                                            </van-col>
                                            <van-col span="4">
                                                <van-tag type="success" plain size="medium"
                                                         v-if="cruentDate<=item.agencyTime">未过期
                                                </van-tag>
                                                <van-tag type="danger" plain size="medium"
                                                         v-if="cruentDate>item.agencyTime">已过期
                                                </van-tag>
                                            </van-col>
                                            <van-col span="2"></van-col>
                                            <van-col span="12" class="f16">
                                                <van-switch v-model="item.status"
                                                            @click.stop="updateAgentStateMth(item)" size="16px"/>
                                            </van-col>
                                        </van-row>
                                        <van-row class="f14 fc mt20">
                                            <van-col span="8">
                                                <van-row>{{item.empName}}</van-row>
                                                <van-row class="fc_gary f12">委托人</van-row>
                                            </van-col>
                                            <van-col span="8">
                                                <van-row>{{item.agencyTime}}</van-row>
                                                <van-row class="fc_gary f12">授权截止时间</van-row>
                                            </van-col>
                                            <van-col span="8">
                                                <van-row>{{item.status==true?'启用中':'已停用'}}</van-row>
                                                <van-row class="fc_gary f12">启用状态</van-row>
                                            </van-col>
                                        </van-row>
                                        <van-row class="f14 fc fc_gary mt5">
                                            <div style="float:left;margin-bottom: 5px">{{item.createTime}}</div>
                                        </van-row>
                                    </div>
                                </van-cell>
                            </van-list>
                        </van-pull-refresh>
                    </div>
                </van-tab>

                <van-tab title="委托列表" name="1">
                    <van-notice-bar color="#1989fa" background="#ecf9ff" wrapable :scrollable="false"
                                    left-icon="info-o">
                        用户可以通过本列表来查看其他用户给自己的委托信息
                    </van-notice-bar>
                    <div :style="{height:clientHeight+'px',overflow:'scroll'}">
<!--                        <div class="content_center mt20" v-if="delegatedSettingsList.length<1">-->
<!--                            <van-image width="220" height="220" src="./src/assets/icon/default_img_03.png"/>-->
<!--                        </div>-->
                        <van-pull-refresh v-model="delegatedSettingsRefresh" @refresh="onDelegatedSettingsRefresh"
                                          :style="{minHeight:clientHeight+'px'}"
                        >
                            <van-list
                                    v-model:loading="delegatedSettingsLoading"
                                    :finished="delegatedSettingsFinished"
                                    finished-text="没有更多了"
                                    offset="100"
                                    @load="onLoadMessage">
                                <van-cell v-if="delegatedSettingsList.length>0" v-for="item in delegatedSettingsList"
                                          :key="item"
                                          class="mt10 pl10">
                                    <div class="cardCss">
                                        <van-row class="f16 ">
                                            <van-col class="fc fw ta-l" span="5">
                                                <span>审批权限</span>
                                            </van-col>
                                            <van-col span="4">
                                                <van-tag type="success" plain size="medium"
                                                         v-if="cruentDate <= item.agencyTime">未过期
                                                </van-tag>
                                                <van-tag type="danger" plain size="medium"
                                                         v-if="cruentDate > item.agencyTime">已过期
                                                </van-tag>
                                            </van-col>
                                            <van-col span="2"></van-col>
                                            <van-col span="12" class="f16">
                                                <van-button type="primary" size="small" style="width: 90px"
                                                            v-if="item.state=='true'&& cruentDate<=item.agencyTime"
                                                            @click="jumpApproval(item)">
                                                    委托跳转 >>
                                                </van-button>
                                                <van-button type="danger" size="small" style="width: 90px"
                                                            v-if="item.state=='false'||cruentDate>item.agencyTime">暂无权限
                                                </van-button>
                                            </van-col>
                                        </van-row>
                                        <van-row class="f14 fc mt20">
                                            <van-col span="12">
                                                <van-row class="f14 fw">{{item.principalName}}</van-row>
                                                <van-row class="fc_gary f12">授权人</van-row>
                                            </van-col>
                                            <van-col span="12">
                                                <van-row class="f14">{{item.agencyTime}}</van-row>
                                                <van-row class="fc_gary f12">授权截止时间</van-row>
                                            </van-col>
                                        </van-row>
                                        <van-row class="f14 fc fc_gary pt10">
                                            <div style="float:left;margin-bottom: 5px">{{item.createTime}}</div>
                                        </van-row>
                                    </div>
                                </van-cell>
                            </van-list>
                        </van-pull-refresh>
                    </div>
                </van-tab>
            </van-tabs>
        </div>
    </div>
</template>

<script setup lang="ts">
    import {ref, onMounted, inject} from 'vue'
    import {getMyAgentPerson, getMyAgentInfo, updateAgentState} from "../../services/myDelegation";
    import router from "../../router";
    import {showConfirmDialog, showSuccessToast} from 'vant';
    import moment from 'moment'

    const cruentDate = ref()
    let userInfo: any = inject("userInfo");
    const activeName = ref(0);
    //委托设定列表
    const delegatedSettingsList = ref([])
    //委托列表
    const delegateList = ref([]);
    //弹出筛选框状态
    const show = ref(false)
    //委托列表
    const delegateLoading = ref(false)
    const delegateFinished = ref(false);
    const delegateTotle = ref(0)
    const delegatePage = ref(0)
    const delegatePageSize = ref(10)

    //委托设定
    const delegatedSettingsLoading = ref(false)
    const delegatedSettingsFinished = ref(false);
    const delegatedSettingsTotle = ref(0)
    const delegatedSettingsPage = ref(0)
    const delegatedSettingsPageSize = ref(10)

    //获取高度
    const clientHeight = ref(document.documentElement.offsetHeight - 160)
    const refreshing = ref(false);
    const delegatedSettingsRefresh = ref(false)
    onMounted(() => {
        cruentDate.value = moment(new Date()).format('YYYY-MM-DD')

    })

    const onDelegateListRefresh = () => {
        delegateFinished.value = false;
        // 重新加载数据
        // 将 loading 设置为 true，表示处于加载状态
        delegateLoading.value = true;
        delegatePage.value = 0
        onLoad();
    };

    const onDelegatedSettingsRefresh = () => {
        delegatedSettingsFinished.value = false;
        // 重新加载数据
        // 将 loading 设置为 true，表示处于加载状态
        delegatedSettingsLoading.value = true;
        delegatedSettingsPage.value = 0
        onLoadMessage();

    };

    function onLoadMessage() {
        delegatedSettingsLoading.value = true
        delegatedSettingsPage.value = delegatedSettingsPage.value + 1
        getDelegatedSettingsMth()
    }

    function onLoad() {
        delegatePage.value = delegatePage.value + 1
        delegateLoading.value = true
        geDelegateListMth()
    };

    function jumpApproval(item: any) {
        let url = `http://122.9.41.215/fk_move/MapprList?userid=${item.principalId}&username=${item.principalName}&school=${userInfo.userInfo.corpCode}&isDark=false&type=false`;
        router.push({
            name: 'approvalDetailsIframe',
            params: {
                url: url,
                realHeight: document.documentElement.offsetHeight,
                titleName: '审批'
            }
        })
    }

    //委托详情
    function queryDelegateDetail(e: any) {
        router.push({
            name: 'addMyDelegation', params: {
                typeName: '委托详情',
                commissionDeadline: JSON.stringify(e)
            }
        })
    }

    //委托设定列表
    async function getDelegatedSettingsMth() {
        let sysMsgList = delegatedSettingsList.value
        setTimeout(async () => {
            if (delegatedSettingsRefresh.value) {
                sysMsgList = []
                delegatedSettingsRefresh.value = false
            }
            let param = {
                page: delegatedSettingsPage.value,
                pageSize: delegatedSettingsPageSize.value,
                corpCode: userInfo.userInfo.corpCode,
                deptId: userInfo.userInfo.deptId,
                empCode: userInfo.userInfo.empCode,
            }
            let res: any = await getMyAgentInfo(param);
            if (res.success) {
                res.data.list.map((item: any) => {
                    sysMsgList.push(item as never)
                })
                delegatedSettingsLoading.value = false;
                delegatedSettingsTotle.value = res.data.total
                if (sysMsgList.length >= delegatedSettingsTotle.value) {
                    delegatedSettingsFinished.value = true;
                }
            } else {
                sysMsgList = []
            }
            delegatedSettingsList.value = sysMsgList
        }, 1000)

    }

    //获取推送消息
    async function geDelegateListMth() {
        let dataList = delegateList.value
        setTimeout(async () => {
            if (refreshing.value) {
                dataList = []
                refreshing.value = false
            }
            let param = {
                page: delegatePage.value,
                pageSize: delegatePageSize.value,
                corpCode: userInfo.userInfo.corpCode,
                deptId: userInfo.userInfo.deptId,
                empCode: userInfo.userInfo.empCode,

            }
            let res: any = await getMyAgentPerson(param);
            if (res.success) {
                res.data.list.map((item: any) => {
                    item.status = item.state == 'true' ? true : false
                    dataList.push(item as never)
                })
                delegateLoading.value = false;
                delegateTotle.value = res.data.total
                if (dataList.length >= delegateTotle.value) {
                    delegateFinished.value = true;
                }
            } else {
                dataList = []
            }
            delegateList.value = dataList
        }, 1000)

    }

    async function updateAgentStateMth(item: any) {
        let title = "此操作将 停用 该条委托信息，是否继续?"
        if (item.status) {
            title = "此操作将 启用 该条委托信息，是否继续?"
        }
        let params = item
        params.modifier = userInfo.userInfo.empCode
        let status = item.status == false ? 'false' : 'true'
        params.state = status

        showConfirmDialog({
            title: title
        }).then(async () => {
            let res: any = await updateAgentState(params)
            if(res.success){
                if(item.status){
                    showSuccessToast("启用成功")
                }else{
                    showSuccessToast("停用成功")
                }
            }
        }) .catch(() => {
            item.status = !item.status
        });
    }

    function addMyDelegation() {
        router.push({
            name: 'addMyDelegation', params: {
                typeName: '新建委托',
                commissionDeadline: JSON.stringify("")
            }
        })
    }

    function back() {
        router.go(-1)
    }


</script>

<style lang="less" scoped>

    .cardCss {
        margin-top: 15px;
    }

    .fc {
        color: #1a1a1a;
    }

    .fw {
        font-weight: 800;
    }

    .ta-l {
        text-align: left;
    }

    .fc_gary {
        color: #969799
    }


</style>
