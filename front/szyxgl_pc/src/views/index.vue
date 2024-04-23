<template>
    <div class="common-layout">
        <el-container>
            <el-aside class="aside" :class="{ asideFold: !!settingStore.fold }">
                <Aside></Aside>
            </el-aside>
            <el-container>
                <el-header class="header" :class="{ headerFold: !!settingStore.fold }">
                    <Header></Header>
                </el-header>
                <el-main :class="{ mainFold: !!settingStore.fold }">
                    <router-view />
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>


<script setup>
import { ref, onMounted } from 'vue'
import { ElNotification } from 'element-plus'
import useSettingStore from '@/store/module/setting.ts'
let settingStore = useSettingStore()
// console.log('settingStore', settingStore)
let message = ref('')

// onMounted(()=>{
//     const now = new Date();
//       const hours = now.getHours();
//       if (hours >= 0 && hours < 6) {
//       message.value = '凌晨好';
//       } else if (hours >= 6 && hours < 12) {
//       message.value = '上午好';
//       } else if (hours >= 12 && hours < 18) {
//       message.value = '下午好';
//       } else {
//       message.value = '晚上好';
//       }
//     ElNotification({
//     title: 'admin',
//     message:message.value,
//     type: 'success',
//   })
// })
</script>
<style  lang='scss'>
// .el-submenu.is-opened>.el-submenu__title .el-submenu__icon-arrow {
//     display: none;

// }

// .el-submenu>.el-submenu__title .el-submenu__icon-arrow {
//     display: none;

// }

.common-layout {
    width: 100%;
    // height: 100vh;
    overflow: hidden;

    // background: red;
    .el-aside {
        width: $base-aside-width;
        // height: 100vh;
        overflow: hidden;
        // background: $base-color-default;
        transition: all .3s;
        border-right: 1px solid $base-border-color;

        .el-menu--collapse .el-submenu__icon-arrow {
            display: none !important;
        }

    }

    .asideFold {
        width: $base-aside-minWidth;
        border-right: 1px solid $base-border-color;
    }

    .header {
        width: 100%;
        height: 84px;
        // padding: 0 10px;
        // margin-left: 10px;
        // & .fold {
        //     width:calc(100vw - 50px)
        // }
    }

    .el-header{
        padding: 0;
    }
    .headerFold,
    .mainFold {
        width: calc(100vw - $base-aside-minWidth);
    }
}
</style>