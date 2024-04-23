<template>
    <div>
        <el-scrollbar>
            <!-- <p v-for="item in 80" :key="item" class="scrollbar-demo-item">菜单:{{ item }}</p> -->
            <!-- <div class="logo" v-if="!settingStore.fold">VueAdmin</div> -->
            <div class="logo">图片</div>
            <el-menu :default-active="route.path" class="menuList" :collapse="settingStore.fold">
                <MenuList :menuList="userMenu.menuRouters"></MenuList>
            </el-menu>
        </el-scrollbar>
    </div>
</template>
<script setup lang="ts">
import useUserStore from '~/store/module/user';
import { useRoute } from 'vue-router'
import useSettingStore from '@/store/module/setting.ts'
let settingStore = useSettingStore()
let userMenu = useUserStore()
// console.log('userMenu', userMenu.menuRouters)


//获取当前路由对象
let route = useRoute()
// console.log(route.path);

</script>
<style lang="scss" scoped>
.aside {
    .logo {
        width: 100%;
        height: 84px;
        line-height: 84px;
        text-align: center;

        color: #333;
        font-size: 36px;
        border-bottom: 1px solid $base-border-color;
    }

    .menuList {
        width: 100%;
        height: calc(100vh - 84px);
        overflow: hidden;
        border: none;
    }

}
</style>
