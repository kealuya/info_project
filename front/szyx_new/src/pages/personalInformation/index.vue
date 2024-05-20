<template>
    <van-nav-bar
            left-arrow
            left-text="返回"
            title="个人信息"
            @click-left="onClickLeft"
    />
   <div class="container">
       <van-form>
           <van-cell-group inset>
               <van-field
                   v-model="username"
                   name="姓名"
                   label="姓名"
                   readonly
               />
               <van-field
                   v-model="userMobile"
                   name="手机号"
                   label="手机号"
                   readonly
               />
<!--               <van-field-->
<!--                   v-model="userInfo"-->
<!--                   name="所属单位"-->
<!--                   label="所属单位"-->
<!--                   readonly-->
<!--               />-->
           </van-cell-group>
       </van-form>
   </div>
</template>
<script setup lang="ts">
import {inject, onMounted, ref} from "vue";
import {useRouter}  from "vue-router"
interface userInfoType{
    userName:string,
    userMobile:string
}
import {convertOptionIdName} from "echarts/types/src/util/model";
const router = useRouter()
const username = ref('');
const userMobile = ref('');
const userInfo = ref<userInfoType>()
const onSubmit = (values:any) => {
    console.log('submit', values);
};
const onClickLeft = ()=>{
router.push('/personal')
}
onMounted(()=>{
    let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
    userInfo.value = {...userInfoData.userInfo}
    username.value =  userInfo.value.userName
    userMobile.value = userInfo.value.userMobile

    // console.log(userInfo.value)
})
</script>
<style scoped lang="less">
.container{
    background-color: #fff;
    margin: 2vw;
    border-radius: 8px;
}
::v-deep(.van-nav-bar__content){
    background-color: #0088ff;
}
::v-deep(.van-nav-bar__title){
    color: #FFFFFF;
}
::v-deep(.van-nav-bar__text){
    color: #FFFFFF;
}
::v-deep(.van-icon:before){
    color: #FFFFFF;
}
</style>