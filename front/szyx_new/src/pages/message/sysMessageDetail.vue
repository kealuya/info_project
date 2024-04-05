<template>
    <div>
        <van-row>
            <NavBar :title="'消息通知'" :leftText="'返回'"  :leftArrow="true" @leftEvent="back()"/>
        </van-row>
        <van-row>
        <div class="pt10 pl10 pr5" :style="{height:realHeight+'px',overflow:'scroll',display: 'flex',flexDirection: 'column'}">
            <van-row class="jcc">
            <div >
                <van-row class="f20 fw800">
                  {{newTitle}}
                </van-row>
                <van-row class="pt10 jcc">
                    <div>发布时间: {{createTime}}</div>
                </van-row>
                <van-row class="pt10 jcc">
                    <div>来源: {{dept}}</div>
                </van-row>
            </div>
            </van-row>
            <van-row class="pt10" >
                <div v-html="htmlMessage"></div>
            </van-row>

        </div>
        </van-row>
    </div>
</template>

<script setup lang="ts">
    import {ref} from 'vue'
    import route from "../../router";
    import {useRoute} from 'vue-router'
    const router = useRoute();
    import {onMounted} from 'vue'
    import {userInfoData} from '../../store'

    const store = userInfoData();
    const htmlMessage = ref('')
    const newTitle = ref('')
    const newType = ref('')
    const createTime = ref('')
    const dept = ref('')
    let  realwidth = window.document.documentElement.offsetWidth
    let  realHeight = window.document.documentElement.offsetHeight-60

    onMounted(() => {
        htmlMessage.value = route.currentRoute.value.params.htmlMessage as never
        newTitle.value = route.currentRoute.value.params.newTitle as never
        newType.value = route.currentRoute.value.params.newType as never
        createTime.value = route.currentRoute.value.params.createTime as never
        dept.value = route.currentRoute.value.params.dept as never
    })
    function back() {
        route.go(-1)
    }

</script>

<style lang="less" scoped>

    .fw800{
        font-weight: 800;
    }
    .jcc{
        justify-content:center;
    }

</style>
