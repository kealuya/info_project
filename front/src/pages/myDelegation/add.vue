<template>
    <div>
        <div>
            <NavBar :leftText="'返回'" :title="titleName" :leftArrow="true" @leftEvent="back()"></NavBar>
        </div>
        <div style="display: flex">
            <div>
                <van-row>
                    <van-notice-bar color="#1989fa" :scrollable="false" wrapable background="#ecf9ff"
                                    left-icon="info-o" class="f10">
                        委托--新建委托--用户可将业务模块权限委托给其他用户进行托管，在设定日期之内，被委托人可以代替委托人进行模块业务处理
                    </van-notice-bar>
                    <van-form @submit="onSubmit" class="mt10" :style="{width:realWidth+'px'}">
                        <van-cell-group inset class="pb10">
                            <i class="spanCss">*</i>
                            <van-field
                                    v-model="modulePermissions"
                                    name="modulePermissions"
                                    label="委托权限模块"
                                    placeholder="请选择委托权限模块"
                                    :rules="[{ required: true, message: '请选择委托权限模块' }]"
                                    class="mt10 mb10 "
                                    :right-icon="titleName=='新建委托'?'arrow':''"
                                    readonly
                                    label-width="7.5em"
                                    @click="titleName=='新建委托'? showModulePermissionsStatus=true:''"

                            />
                            <i class="spanCss">*</i>
                            <van-field
                                    v-model="entrustedPersonnel"
                                    name="entrustedPersonnel"
                                    label="委托人员名称"
                                    placeholder="请选择委托人员名称"
                                    :rules="[{ required: true, message: '请选择委托人员名称' }]"
                                    class="mt10 mb10 "
                                    :right-icon="titleName=='新建委托'?'arrow':''"
                                    readonly
                                    label-width="7.5em"
                                    @click.stop="titleName=='新建委托'? selectPeoplePickerStatus = true:''"

                            />
                            <i class="spanCss">*</i>

                            <van-field
                                    v-model="commissionDeadline"
                                    name="commissionDeadline"
                                    label="委托截止时间"
                                    placeholder="请选择委托截止时间"
                                    :rules="[{ required: true, message: '请选择委托截止时间' }]"
                                    :is-link="titleName=='新建委托'||titleName=='更改委托'"
                                    readonly
                                    label-width="7.5em"
                                    @click.stop="titleName=='新建委托'||titleName=='更改委托'?showPicker = true:''"
                                    class="mt10 mb10"
                            />

                            <van-cell v-if="titleName=='新建委托'" class="submitCss">
                                <van-button type="primary" size="normal" style="width: 180px" native-type="submit">
                                    提交委托申请
                                </van-button>
                            </van-cell>
                            <van-cell v-if="titleName=='更改委托'" class="submitCss">
                                <van-button type="primary" size="normal" style="width: 180px" native-type="submit">
                                    更改委托
                                </van-button>
                            </van-cell>
                            <van-cell class="detailBtnCss" v-if="titleName=='委托详情'">
                                <van-row>
                                    <van-col span="11">
                                        <van-button type="primary" size="normal"
                                                    style="width: 120px" @click="deleteDelegationItem">
                                            删除
                                        </van-button>
                                    </van-col>
                                    <van-col span="2"></van-col>
                                    <van-col span="11">
                                        <van-button type="primary" size="normal"
                                                    style="width: 120px" @click="updateMth">
                                            更改
                                        </van-button>
                                    </van-col>
                                </van-row>
                            </van-cell>
                        </van-cell-group>
                    </van-form>
                </van-row>
            </div>
        </div>

        <van-action-sheet
                v-model:show="showModulePermissionsStatus"
                :actions="actionsModulePermissions"
                cancel-text="取消"
                description="请点击选择"
                close-on-click-action @select="onSelectModulePermissions">
        </van-action-sheet>
        <!--  选择人员-->
        <van-popup v-model:show="selectPeoplePickerStatus" style="width:100vw;height:100vh" position="bottom">
            <select-people-picker @select-over="selectPeople" ref="selectPeopleOver" :list="selectPersonList"/>
        </van-popup>

        <van-popup v-model:show="showPicker" position="bottom">
            <van-date-picker @confirm="onConfirmPicker" :min-date="minDateFmt" @cancel="showPicker = false"/>
        </van-popup>

    </div>
</template>

<script setup lang="ts">
    import {inject, onBeforeMount, ref} from 'vue'
    import router from "../../router";
    import {createAgentInfo, deleteAgentInfo, updateAgentInfo} from '../../services/myDelegation'
    import SelectPeoplePicker from '../../components/selectComponent/selectPeoplePicker/index.vue' // 引入选择出差人员的组件
    import {showSuccessToast, showFailToast, showConfirmDialog} from "vant";
    import {showToast} from 'vant'
    import moment from 'moment';

    //自适应宽度
    let realWidth = document.documentElement.offsetWidth
    //弹窗状态
    const show = ref(false)
    let userInfoData: any = inject("userInfo");
    const submitBtnDisable = ref(false)//添加账户按钮状态
    const titleName = ref('新建委托')
    let minDateFmt = new Date(moment(new Date()).format("YYYY-MM-DD"))

    //------------------------------------------------------------------
    const showModulePermissionsStatus = ref(false)
    const actionsModulePermissions = ref([{name: '审批权限', value: 'approval'}])
    const modulePermissions = ref("")//委托权限模块
    const modulePermissionsValue = ref("")//委托权限模块

    const entrustedPersonnel = ref("")//委托人员名称
    const entrustedPersonnelEmpCode = ref("")//委托人员code
    const commissionDeadline = ref("")//委托截止时间
    const selectPeoplePickerStatus = ref(false) //选择人员
    let selectPersonList = []; // 用于选择自定义审批流中的选择审批人
    const showPicker = ref(false) //选择时间

    onBeforeMount(() => {
        titleName.value = router.currentRoute.value.params.typeName as never
        if (titleName.value == '委托详情') {
            let obj = JSON.parse(router.currentRoute.value.params.commissionDeadline as never)
            let agentType = actionsModulePermissions.value.map((item) => {
                if (item.value == obj.agentType) {
                    modulePermissions.value = item.name
                }
            })
            entrustedPersonnel.value = obj.empName
            entrustedPersonnelEmpCode.value = obj.empCode
            commissionDeadline.value = obj.agencyTime
        }
    })

    const onConfirmPicker = (e: any) => {
        commissionDeadline.value = e.selectedValues[0] + '-' + e.selectedValues[1] + '-' + e.selectedValues[2]
        showPicker.value = false
    };

    const onSubmit = (values: any) => {
        submitBtnDisable.value = true
        if (titleName.value == '新建委托') {
            createDelegationMth()
        } else {
            updateDelegation()
        }
        setTimeout(() => {
            submitBtnDisable.value = false
        }, 3000)
    };

    function onSelectModulePermissions(e: any) {
        modulePermissions.value = e.name
        modulePermissionsValue.value = e.value
    }

    const selectPeople = (e: any) => {
        if (e.length > 1) {
            showToast('请选择一个人员信息');
            return
        } else if (e.length == 1) {
            entrustedPersonnel.value = e[0].empName
            entrustedPersonnelEmpCode.value = e[0].empCode
            selectPeoplePickerStatus.value = false
        } else {
            selectPeoplePickerStatus.value = false
        }
    }

    const selectPerson = (() => {
        selectPeoplePickerStatus.value = true
    })

    async function createDelegationMth() {
        let params = {
            empCode: entrustedPersonnelEmpCode.value,
            empName: entrustedPersonnel.value,
            agentType: modulePermissionsValue.value,
            principalName: userInfoData.userInfo.userName,
            principalId: userInfoData.userInfo.empCode,
            principalDeptId: userInfoData.userInfo.deptId,
            principalDeptName: userInfoData.userInfo.deptName,
            corpCode: userInfoData.userInfo.corpCode,
            agencyTime: commissionDeadline.value,
            creater: userInfoData.userInfo.empCode
        }
        let res: any = await createAgentInfo(params)
        if (res.success) {
            showSuccessToast(res.data)
            router.go(-1)
        } else {
            showFailToast(res.msg)
        }
    }

    function back() {
        router.go(-1)
    }

    async function deleteDelegationItem(item: any) {
        showConfirmDialog({
            title: "您确定删除该数据吗?"
        }).then(async () => {
            let param = {
                empCode: entrustedPersonnelEmpCode.value,
                principalId: userInfoData.userInfo.empCode,
                corpCode: userInfoData.userInfo.corpCode
            }
            let res: any = await deleteAgentInfo(param)
            if (res.success) {
                showSuccessToast("删除成功")
                router.go(-1)
            } else {
                showFailToast(res.msg)
            }
        })
    }

    function updateMth() {
        titleName.value = '更改委托'
    }

    async function updateDelegation() {
        let params = {
            empCode: entrustedPersonnelEmpCode.value,
            principalId: userInfoData.userInfo.empCode,
            corpCode: userInfoData.userInfo.corpCode,
            agencyTime: commissionDeadline.value,
            modify: userInfoData.userInfo.empCode,
        }
        let res: any = await updateAgentInfo(params)
        if (res.success) {
            showSuccessToast("更改成功")
            router.go(-1)
        } else {
            showFailToast(res.msg)
        }
    }

</script>

<style lang="less" scoped>

    .bottom-wrap {
        height: 70px;
        width: 100%;
        position: fixed;
        bottom: 0;
        background: #FFFFFF;
    }

    /deep/ .submitCss .van-cell__value {
        text-align: center !important;
        margin-top: 100px;

    }

    /deep/ .detailBtnCss .van-cell__value {
        text-align: center;
        margin-top: 100px;

    }

    .spanCss {
        color: red;
        position: absolute;
        left: 120px;
        z-index: 999;
        padding-top: 13px;
    }

</style>
