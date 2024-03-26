<template>
    <div>
        <div>
            <NavBar :leftText="'账户管理'" :leftArrow="true" @leftEvent="back()"></NavBar>
        </div>
        <div style="display: flex">
            <div>
                <van-row>
                    <van-form @submit="onSubmit" class="mt20" :style="{width:realWidth+'px'}">
                        <van-cell-group inset class="pb10">
                            <div class="mt10 ">
                                <div align="center">
                                    <van-icon class="pr5" :name="bank_icon_01"/>
                                    <span>{{titleName}}</span>
                                </div>
                                <span class="f12 font_gray ">
                                     <div align="center" class="pt10">请绑定持卡人本人相关信息</div>
                                </span>
                            </div>
                            <van-field
                                    v-model="bankType"
                                    name="账户类型"
                                    label="账户类型"
                                    placeholder="请选择账户类型"
                                    :rules="[{ required: true, message: '请选择账户类型' }]"
                                    class="mt10"
                                    :right-icon="titleName=='添加账户' || accountBookDetail.isdefaultpayee=='false'?'arrow':''"
                                    required="true"
                                    readonly
                                    @click="titleName=='添加账户' || accountBookDetail.isdefaultpayee=='false'? show=true:''"
                            />

                            <van-field
                                    v-model="bankNumber"
                                    name="银行卡号"
                                    label="银行卡号"
                                    placeholder="请输入银行卡号"
                                    :rules="[{required: true,validator, message: '请输入正确的银行卡号' }]"
                                    class="mt10"
                                    required="true"
                                    :update:model-value="(bankNumber = bankNumber.replace(/\s+/g, ''))"
                            />
                            <van-field
                                    v-model="bankName"
                                    name="开户行名称"
                                    label="开户行名称"
                                    placeholder="请输入开户行名称"
                                    :rules="[{ required: true, message: '请输入开户行名称' }]"
                                    maxlength="30"
                                    class="mt10"
                                    required="true"
                                    :update:model-value="(bankName = bankName.replace(/\s+/g, ''))"
                            />
                            <van-field
                                    v-model="nameOfPayee"
                                    name="收款方名称"
                                    label="收款方名称"
                                    placeholder="请输入收款方名称"
                                    :rules="[{ required: true, message: '请输入收款方名称' }]"
                                    class="mt10"
                                    maxlength="30"
                                    required="true"
                                    :update:model-value="(nameOfPayee = nameOfPayee.replace(/\s+/g, ''))"
                            />
                            <van-field
                                    v-model="areaName"
                                    name="地区"
                                    label="地区"
                                    placeholder="请选择所在地区"
                                    :rules="[{ required: true, message: '请选择所在地区' }]"
                                    class="mt10"
                                    is-link
                                    readonly
                                    required="true"
                                    @click="showAreaMth(true)"
                            >
                            </van-field>
                            <van-field
                                    v-model="openingBankBranch"
                                    name="开户行网点"
                                    label="开户行网点"
                                    placeholder="请输入开户行网点"
                                    :rules="[{ required: true, message: '请输入开户行网点' }]"
                                    class="mt10"
                                    maxlength="35"
                                    required="true"
                                    :update:model-value="(openingBankBranch = openingBankBranch.replace(/\s+/g, ''))"
                            />
                        </van-cell-group>
                        <div class="bottom-wrap">
                            <van-row justify="center" >
                                <div style="margin: 16px;" v-if="titleName=='添加账户'">
                                    <van-button block type="primary" style="width:200px" native-type="submit" :disabled="submitBtnDisable">
                                        确定绑定
                                    </van-button>
                                </div>
                                <div style="margin: 16px; display: flex" v-if="titleName=='编辑账户'">
                                    <van-row gutter="25">
                                        <van-col>
                                            <van-button block style="width:120px" type="primary" @click="deletePayeeInfoMth" :disabled="unbindingBtnDisable">
                                                解除绑定
                                            </van-button>
                                        </van-col>
                                        <van-col></van-col>
                                        <van-col>
                                            <van-button block style="width:120px" type="primary" native-type="submit" :disabled="submitBtnDisable">
                                                确认修改
                                            </van-button>
                                        </van-col>
                                    </van-row>
                                </div>
                            </van-row>
                        </div>
                    </van-form>
                </van-row>
            </div>
        </div>
        <div>
            <van-popup v-model:show="showArea" position="bottom">
                <van-area
                        v-model="cityCode"
                        :area-list="areaList"
                        columns-num="2"
                        @confirm="onConfirm"
                        @cancel="showAreaMth(false)"
                />
            </van-popup>
        </div>
        <van-action-sheet
                v-model:show="show"
                :actions="actions"
                cancel-text="取消"
                description="请点击选择"
                close-on-click-action @select="onSelect">
        </van-action-sheet>
    </div>
</template>

<script setup lang="ts">
    import {inject, onMounted, onBeforeMount, ref, defineEmits, defineProps} from 'vue'
    import {areaList} from '@vant/area-data';
    import router from "../../router";
    import {addPayeeInfo, deletePayeeInfo, updatePayeeInfo} from '../../services/manageAccounts'
    import {showSuccessToast, showFailToast, showConfirmDialog} from "vant";
    import bank_icon_01 from './../../assets/icon/bank_icon_01.png'
    import moment from 'moment';

    //表单参数
    const bankType = ref('')
    const bankNumber = ref('')
    const bankName = ref('')
    const nameOfPayee = ref('')
    const areaName = ref('')
    const openingBankBranch = ref('')
    //自适应宽度
    let realWidth = document.documentElement.offsetWidth
    //弹窗状态
    const show = ref(false)
    let userInfoData: any = inject("userInfo");
    const actions = [
        {name: '个人账户'},
        {name: '对公账户'},
    ];
    const cityCode = ref()
    const cityName = ref("")
    const regionName = ref("")
    //银行卡正则
    const validator = (val: string) => /^([1-9]{1})(\d{15}|\d{16}|\d{18})$/.test(val);
    const result = ref('');
    const showArea = ref(false);
    const submitBtnDisable = ref(false)//添加账户按钮状态
    const unbindingBtnDisable = ref(false)
    let titleName = '添加账户'
    const accountObj = ref()
    const accountBookDetail = ref()
    const isComponents = ref("")

    const emit = defineEmits(["close"])
    const props = defineProps({
        bankCardParams: {
            type: Object,
        }
    })

    onBeforeMount(() => {
        accountObj.value = props.bankCardParams
        titleName = props.bankCardParams.title
        if (titleName == '编辑账户') {
            accountBookDetail.value = accountObj.value.accountBookDetail
            bankType.value = accountBookDetail.value.type == 'personalAccount' ? '个人账户' : '对公账户'
            bankNumber.value = accountBookDetail.value.bankcardnum
            bankName.value = accountBookDetail.value.bankaccountname
            nameOfPayee.value = accountBookDetail.value.payeeName
            areaName.value = accountBookDetail.value.city + "-" + accountBookDetail.value.region
            regionName.value = accountBookDetail.value.region
            cityName.value = accountBookDetail.value.city,
                cityCode.value = accountBookDetail.value.cityCode
            openingBankBranch.value = accountBookDetail.value.accountnetwork
        }
    })

    const onConfirm = ({selectedOptions}: any) => {
        showArea.value = false;
        result.value = selectedOptions.map((item: { text: any; }) => item.text).join('-');
        cityCode.value = selectedOptions[1].value
        cityName.value = selectedOptions[0].text
        regionName.value = selectedOptions[1].text
        areaName.value = result.value

    };

    const onSelect = (item: any) => {
        bankType.value = item.name
        show.value = false;
    };

    function showAreaMth(e: any) {
        showArea.value = e
    }

    const onSubmit = (values: any) => {
        submitBtnDisable.value = true
        if (titleName == '添加账户') {
            addPayeeInfoMth()
        } else {
            updatePayeeInfoMth()
        }
        setTimeout(() => {
            submitBtnDisable.value = false
        }, 3000)
        emit("close", bankType.value)
    };

    async function addPayeeInfoMth() {
        let dateTime = new Date()
        let params = {
            empcode: userInfoData.userInfo.empCode,
            type: bankType.value == '个人账户' ? 'personalAccount' : 'publicAccount',
            bankcardnum: bankNumber.value,
            bankaccountname: bankName.value,
            city: cityName.value,
            region: regionName.value,
            cityCode: cityCode.value,
            accountnetwork: openingBankBranch.value,
            isdefaultpayee: false,
            payeeName: nameOfPayee.value,
            createtime: moment(dateTime).format("YYYY-MM-DD"),
            corpCode: userInfoData.userInfo.corpCode
        }
        let res: any = await addPayeeInfo(params)
        if (res.code == 200) {
            showSuccessToast(res.data)
        }
        if (res.code == 500) {
            showFailToast(res.message)
        }

    }

    function back() {
        let type = bankType.value == '' ? '个人账户' : bankType.value
        emit("close", type)
    }

    async function deletePayeeInfoMth() {
        showConfirmDialog({
            title: '您确认要解绑该银行卡吗?'
        }).then(async () => {
            let param = {
                id: accountBookDetail.value.id,
                corpCode: userInfoData.userInfo.corpCode
            }
            unbindingBtnDisable.value = true
            let res: any = await deletePayeeInfo(param)
            if (res.code == 200) {
                emit("close", bankType.value)
                showSuccessToast("解绑成功")
            } else {
                showFailToast(res.message)
            }
            setTimeout(() => {
                unbindingBtnDisable.value = false
            }, 3000)
        })
    }

    async function updatePayeeInfoMth() {
        let dateTime = new Date()
        let params = {
            empcode: userInfoData.userInfo.empCode,
            type: bankType.value == '个人账户' ? 'personalAccount' : 'publicAccount',
            id: accountBookDetail.value.id,
            bankcardnum: bankNumber.value,
            bankaccountname: bankName.value,
            city: cityName.value,
            region: regionName.value,
            cityCode: cityCode.value,
            accountnetwork: openingBankBranch.value,
            isdefaultpayee: accountBookDetail.value.isdefaultpayee,
            payeeName: nameOfPayee.value,
            createtime: moment(dateTime).format("YYYY-MM-DD"),
            corpCode: userInfoData.userInfo.corpCode
        }
        let res: any = await updatePayeeInfo(params)
        if (res.code == 200) {
            showSuccessToast("修改成功")
            emit("close", bankType.value)
        } else {
            showFailToast(res.message)
        }
    }


</script>

<style lang="less" scoped>

    .bottom-wrap {
        height: 70px;
        width: 100%;
        /*position: fixed;*/
        bottom: 0;
        background: #FFFFFF;
        margin-top: 20%;
    }
</style>
