<template>
    <div class="header">
        <div class="headerLeft">
            <el-icon class="expandIcon" @click="changeIcon">
                <Expand v-if="settingStore.fold" />
                <Fold v-else />
            </el-icon>
            <el-breadcrumb>
                <el-breadcrumb-item v-for="(item, index) in route.matched" :key="index" v-show="item.meta.title">
                    <span>{{ item.meta.title }}</span>
                </el-breadcrumb-item>
            </el-breadcrumb>
        </div>
<!--        <div class="headerRight">-->
<!--            &lt;!&ndash; <el-button :icon="FullScreen" size="small" circle @click="toggleFullscreen()"></el-button>-->
<!--            <el-switch class="switchModel" style="&#45;&#45;el-switch-on-color: #2C2C2C;" v-model="dark" :active-action-icon=" Sunny" :inactive-action-icon="Moon"  @change="switchDark" /> &ndash;&gt;-->
<!--            &lt;!&ndash; <el-button :icon="Setting" size="small" circle></el-button> &ndash;&gt;-->
<!--           <div  >-->
<!--            <img class="pic" src="../../assets/images/login/userPic.gif" alt="">-->
<!--            <el-dropdown>-->
<!--                <span class="el-dropdown-link">-->
<!--                    Admin-->
<!--                    <el-icon class="el-icon&#45;&#45;right">-->
<!--                        <arrow-down />-->
<!--                    </el-icon>-->
<!--                </span>-->
<!--                <template #dropdown>-->
<!--                    <el-dropdown-menu>-->
<!--                        <el-dropdown-item>个人中心</el-dropdown-item>-->
<!--                        <el-dropdown-item @click="logOut">退出</el-dropdown-item>-->
<!--                    </el-dropdown-menu>-->
<!--                </template>-->
<!--            </el-dropdown>-->
<!--           </div>-->
<!--        </div>-->
    </div>
</template>

<script setup>
import { reactive, ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { FullScreen, Setting,Sunny , Moon   } from '@element-plus/icons-vue'
import useSettingStore from '@/store/module/setting.ts'
import screenfull from 'screenfull'//全屏组件
import router from "~/router";
//切换暗黑模式
const dark = ref(false)
const switchDark = ()=>{
    //获取根节点
    let html  = document.documentElement
    //判断是否有dark的类名 
    dark.value?html.className='dark':html.className=''
    
}
let route = useRoute()
let settingStore = useSettingStore()
let fold = ref(false)
//左侧菜单栏缩放
const changeIcon = () => {
    settingStore.fold = !settingStore.fold

}
//退出
const logOut = ()=>{
    console.log('aaa')
    router.push("/login");
}
//全屏
// const fullScreen = ref(false)
const toggleFullscreen = () => {
    if (screenfull.isEnabled) {
        screenfull.toggle()
    }

}

</script>
<style scoped lang='scss'>
.header {
    width: 100%;
    height: 84rpx;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #dcdfe6;
    border-left: none;
    padding: 20px;
    .headerLeft {
        display: flex;
        .expandIcon {
            margin-right: 10px;
        }
    }
    .headerRight {
        display: flex;
        // flex-direction: row-reverse;
        align-items: center;
        .el-switch {
            margin: 0 18px;
        }
        .pic {
            width: 24px;
            height: 24px;
            border-radius: 12px;
            // background: #000;
            // margin-left: 12px;
            margin-right: 5px;
        }
    }



}
</style>