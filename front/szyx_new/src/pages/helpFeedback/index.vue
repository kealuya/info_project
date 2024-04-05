<template>
    <div>
        <div>
            <NavBar :leftText="'返回'" :title="'帮助反馈'" :leftArrow="true" @leftEvent="back()"></NavBar>
        </div>
        <div>
            <van-form @submit="onSubmit">
                <i class="spanCss">*</i>
                <van-field
                        v-model="describeTheProblem"
                        rows="5"
                        autosize
                        label="描述问题"
                        type="textarea"
                        maxlength="100"
                        placeholder="请输入描述问题"
                        show-word-limit
                        :rules="[{ required: true, message: '请输入描述问题' }]"
                        class="mt5"
                        :update:model-value="(describeTheProblem = describeTheProblem.replace(/\s+/g, ''))"
                />
                <van-field class="mt5" name="uploader" label="上传图片" label-width="6em">
                    <template #input>
                        <van-uploader v-model="fileList" :after-read="getFile" :before-delete="deleteFile"
                                      max-count="4"/>
                    </template>
                </van-field>
                <div style="margin: 16px;">
                    <van-button round block type="primary" native-type="submit" :disabled="disableState">
                        提交
                    </van-button>
                </div>
            </van-form>
        </div>
    </div>
</template>

<script setup lang="ts">
    import {ref, inject} from 'vue'
    import {addFeedBack} from "../../services/helpFeedback";
    import router from "../../router";
    import {showFailToast, showSuccessToast } from 'vant';
    import moment from 'moment'
    import {UpLoadImg} from "../../components/hw_obs";

    let userInfo: any = inject("userInfo");
    const fileList = ref([]);
    const fileListUploud = ref([]);
    const disableState = ref(false)
    const describeTheProblem = ref("")

    const getFile = async (file: any) => {
        let obsResult = await UpLoadImg([file]);
        if (obsResult) {
            let index = obsResult[0].lastIndexOf('/')
            let name = ""
            if (index != -1) {
                name = obsResult[0].substring(index + 1, obsResult[0].length)
            }
            fileListUploud.value.push({
                name: file.file.name,
                url: obsResult[0],
                fileName: name
            } as never)
        }
    }

    const deleteFile = async (file: any, index: any) => {
        fileList.value.splice(index.index, 1)
        fileListUploud.value.splice(index.index, 1)
    }

    const onSubmit = (values: any) => {
        disableState.value = true
        submitHelpFeedback()
        setTimeout(() => {
            disableState.value = false
        }, 2000)
    };

    async function submitHelpFeedback() {
        let files = fileListUploud.value.join(",")
        let params = {
            content: describeTheProblem.value,
            corpCode: userInfo.userInfo.corpCode,
            fileUrl: files,
            creater: userInfo.userInfo.empCode,
        }
        let res: any = await addFeedBack(params)
        if (res.success) {
            showSuccessToast(res.data)
            router.go(-1)
        } else {
            showFailToast(res.data)
        }
    }

    function back() {
        router.go(-1)
    }

</script>

<style lang="less" scoped>

    .spanCss {
        color: red;
        position: absolute;
        left: 80px;
        z-index: 999;
        padding-top: 13px;
    }

</style>
