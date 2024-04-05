<template>
    <div>
        <NavBar :leftText="'返回'" :title="tripTitle" :leftArrow="true" @leftEvent="back()"></NavBar>
    </div>
    <div>
        <iframe :src="businessIframeUrl" style="width: 100vw;height: 93.8vh" :height="realHeight"
                class="apprIframe"></iframe>
    </div>
</template>

<script setup lang="ts">
    import {ref, onMounted, inject} from 'vue'
    import {showToast, showFailToast} from "vant";
    import route from "../../router/index";
    import {selectPerIDCard} from "../../services/current"
    import {loginToAuvgo, getToken, getLoginHtShopUrl} from "../../services/trirdPartyBusiness"
    import {userInfoData} from '../../store'
    import moment from "moment";

    let userInfo: any = inject("userInfo");
    const store = userInfoData();
    let userTravelConfig: any = store.userTravelConfig
    const travelType = ref('')
    const token = ref("")
    const businessIframeUrl = ref("")
    const tripTitle = ref("商务出行")
    let realHeight = window.document.documentElement.offsetHeight - 53

    onMounted(() => {
        travelType.value = route.currentRoute.value.params.applicationType as never
        console.log("userTravelConfig", userTravelConfig)
        let type = ""
        tripTitle.value = "商务出行"
        switch (travelType.value) {
            case "plane":
                type = 'air'
                jumpToAuvgo(type)
                break;
            case "hotel":
                type = 'hotel'
                jumpToAuvgo(type)
                break;
            case "train":
                type = 'train'
                jumpToAuvgo(type)
                break;
            case "shop":
                tripTitle.value = "心意选"
                jumpToShop()
        }

    })

    async function jumpToAuvgo(e: any) {
        //判断userTravelConfig 是否拥有直接购票无需申请单的权限
        for (let i = 0; i < userTravelConfig.length; i++) {
            if (userTravelConfig[i].travelPower != 'travel') {
                e = 'tol'
            }
        }
        //查询个人身份证信息接口
        let IDCardRes: any = await selectPerIDCard({
            corpCode: userInfo.userInfo.corpCode,
            empCode: userInfo.userInfo.empCode,
        })
        //判断有无维护身份信息，没有维护不可单点购票
        if (IDCardRes.data) {
            getTokenMth(e).then(() => {
                loginToAuvgoMth(e)
            })
        } else {
            showFailToast('请前往个人中心，完善身份证号信息')
        }
    }

    async function getTokenMth(e: any) {
        let params = {
            target: e,
            empCode: userInfo.userInfo.empCode,
            corpCode: userInfo.userInfo.corpCode,
            appKey: userInfo.userInfo.appKey,
            appSceret: userInfo.userInfo.appSceret,
            customerNo: userInfo.userInfo.customerNo
        }
        let res: any = await getToken(params)
        if (res.success) {
            token.value = res.data.token
        } else {
            console.log("访问第三方业务失败，获取token失败：", res.msg)
            showFailToast("访问第三方业务失败，请联系管理员！")
        }
    }

    //行旅单点登录
    async function loginToAuvgoMth(e: any) {
        let params = {
            jwt: token.value,
            empCode: userInfo.userInfo.empCode,
            corpCode: userInfo.userInfo.corpCode,
            isMobile: true,
            companyCode: userInfo.userInfo.corpCode,
            appKey: userInfo.userInfo.appKey,
            appSceret: userInfo.userInfo.appSceret,
            customerNo: userInfo.userInfo.customerNo,
            ticketType: e
        }
        let res: any = await loginToAuvgo(params)
        if (res.success) {
            let url = res.data.data.loginUrl
            businessIframeUrl.value = url
            console.log("url", url)
        } else {
            console.log("访问第三方业务失败，获取token失败：", res.msg)
            showFailToast("访问第三方业务失败，请联系管理员！")
        }
    }

    async function jumpToShop() {
        let params = {
            empCode: userInfo.userInfo.empCode,
        }
        let res: any = await getLoginHtShopUrl(params)
        if (res.success) {
            let url = res.data.url
            businessIframeUrl.value = url
        }else{
            console.log("访问第三方业务失败，获取token失败：", res.msg)
            showFailToast("访问第三方业务失败，请联系管理员！")
        }
    }

    function back() {
        route.go(-1)
    }

</script>

<style lang="less" scoped>

    .apprIframe {
        border: none
    }
</style>
