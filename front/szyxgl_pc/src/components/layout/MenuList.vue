<template>
  <div>
    <!-- {{ menuList }} -->
    <template v-for="(item, index ) in menuList" :key="item.path">
      <!-- 没有子路由 -->
      <template v-if="!item.children" :index="item.path">

        <el-menu-item v-if="!item.meta.hidden" @click="goRoute" :index="item.path">
          <el-icon>
            <component :is="item.meta.icon"></component>
          </el-icon>
          <template #title>
            <span>{{ item.meta.title }}</span>
          </template>
        </el-menu-item>
      </template>
      <!-- 有一个子路由 -->
      <template v-if="item.children && item.children.length == 1" :index="item.path">
        <el-menu-item v-if="!item.children[0].meta.hidden" :index="item.children[0].path" @click="goRoute">
          <el-icon>
            <component :is="item.meta.icon"></component>
          </el-icon>
          <template #title>
            <!-- <el-icon>
              <component :is="item.children[0].meta.icon"></component>
            </el-icon> -->
            <span>{{ item.children[0].meta.title }}</span>
          </template>
        </el-menu-item>
      </template>
      <!-- 有多个子路由 -->

      <el-sub-menu :index="item.path" v-if="item.children && item.children.length > 1"
        :class="{ mainFold: settingStore.fold ? true : false }">
        <template #title>
          <el-icon :size="20">
            <Edit />
          </el-icon>
          <!-- <el-icon>
            <component :is="item.meta.icon"></component>
          </el-icon> -->
          <span>{{ item.meta.title }}</span>

        </template>
        <Menu :menuList="item.children"></Menu>
      </el-sub-menu>
    </template>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import useSettingStore from '@/store/module/setting.ts'
let settingStore = useSettingStore()
let router = useRouter()
defineProps(['menuList'])
const goRoute = (path: any) => {
  // console.log(path.index)
  router.push(path.index)
}
// console.log('menuListaaa', menuList)
</script>

<!-- 使用vue2语法，为递归组件做基础 -->
<script lang="ts">
export default {
  name: 'Menu'
}
</script>
<style lang="scss">
.mainFold {

  .el-sub-menu__title{
    width: 50px
  }
  .el-sub-menu__title i:nth-of-type(2) {
  /* 添加你需要的样式 */
  display: none
}
}
</style>