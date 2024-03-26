<template>
    <div>
        <div>
            <NavBar :iconName="'wap-home-o'" :slortName="'left'" :iconText="barTitle"  @leftEvent="back()"/>
        </div>
        <div :style="{height:realHeight+'px',width:'100%' }">
            <iframe :src="iframeUrl" width="100%" height="100%" class="apprIframe"></iframe>
        </div>

    </div>
</template>

<script setup lang="ts">
    import {inject, ref, onMounted} from 'vue'
    import router from "../../router";

    const barTitle = ref("")
    const leftText = ref("")
    const iframeUrl = ref("")
    let realHeight = window.document.documentElement.offsetHeight-50

    onMounted(() => {
        barTitle.value = router.currentRoute.value.params.barTitle as never
        iframeUrl.value = router.currentRoute.value.params.iframeUrl as never

    })
    function back() {
        router.go(-1)
    }


</script>

<style lang="less" scoped>


    .apprIframe {
        border: none
    }

</style>
