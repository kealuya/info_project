<template>
    <router-view/>


    <van-tabbar v-model="active" route>
<!--        <van-tabbar-item icon="home-o" name="home" :to="{name:'home'}" replace>首页</van-tabbar-item>-->
<!--        <van-tabbar-item icon="chat-o" name="message" :badge="store.unReadMessageNum==0?'':store.unReadMessageNum"-->
<!--                         :to="{name:'message'}"  :replace="false">消息-->
<!--        </van-tabbar-item>-->
        <van-tabbar-item icon="records" name="book" :to="{name:'sparkV35'}" replace>智能问诊</van-tabbar-item>
        <van-tabbar-item icon="apps-o" name="system" :to="{name:'speech'}" replace>会议记录</van-tabbar-item>
        <van-tabbar-item icon="friends-o" name="user" :to="{name:'user'}" replace>我的</van-tabbar-item>
    </van-tabbar>
</template>

<script setup lang="ts">
    import {ref, onBeforeMount, inject} from "vue";
    import {getMessagePull} from "./services/message";
    import {userInfoData} from './store'
    import router from "./router/index";

    const active = ref('sparkV35');
    const unReadMessage = ref(0)
    const store = userInfoData();
    const msgCount = ref(0)
    let userInfo:any = inject("userInfo")
    const pullMessageObj = ref("")

    onBeforeMount(() => {
        getPullMessage().then(() => {
            store.$patch({
                unReadMessageNum: msgCount.value
            });
            store.$patch({
                pullMessage: pullMessageObj.value
            });
        })
    })

    async function getPullMessage() {
        let params = {
            page: 1,
            pageSize: 10,
            corpCode: userInfo.userInfo.corpCode,
            deptId: userInfo.userInfo.deptId,
            empCode: userInfo.userInfo.empCode
        }
        let res: any = await getMessagePull(params);
        if (res.success) {
            msgCount.value = res.data.count
            res.data.list.map((item: any) => {
                item.dotReadState = item.isRead == '0' ? true : false
            })
            pullMessageObj.value = res.data
        } else {
            msgCount.value = 0
        }
    }


</script>

<style lang="less" scoped>
</style>
