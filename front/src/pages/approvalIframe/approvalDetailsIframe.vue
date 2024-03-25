<template>
    <div>
        <van-row>
            <NavBar :title="titleName" :leftText="'返回'" :leftArrow="true" @leftEvent="back()"/>
        </van-row>
    </div>
    <div :style="{height:realHeight+'px',width:'100%'}">
        <iframe :src="iframeUrl" width="100%" height="100%" class="apprIframe"></iframe>
    </div>
</template>

<script setup lang="ts">
    import {onBeforeMount} from "vue";
    import {useRoute} from 'vue-router'
    import router from "../../router";
    import {userInfoData} from '../../store'

    const store = userInfoData();
    const route = useRoute();
    let iframeUrl = '';
    let realHeight = 0;
    let realwidth = 0;
    let titleName = '';


    onBeforeMount(async () => {
        iframeUrl = route.params.url as never
        realHeight = route.params.realHeight as never - 60
        realwidth = window.document.documentElement.offsetWidth
        titleName = route.params.titleName as never
    })
    function back() {
        router.go(-1)
    }

</script>

<style scoped>
    .apprIframe {
        border: none
    }
    .barCss {
        --van-nav-bar-icon-color: #fafafa;
        --van-nav-bar-text-color: #fafafa;
        --van-nav-bar-title-text-color: #fafafa;
        --van-nav-bar-background:#0080FF;
    }



</style>
