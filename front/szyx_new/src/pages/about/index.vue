<template>
  <div>
    <nav-bar-component title="功能中心" v-if="editIcon==false" :leftText="''" :leftArrow="true"
                       @rightEvent="editIconList" right-text="编辑" @leftEvent="back()"/>
    <navbar title="管理功能" v-else left-text="取消" right-text="保存" @leftEvent="escIconList"
            @rightEvent="saveIconlist"/>
    <div class="card">
      <van-row justify="space-between" align="center" @click="open=!open">
        <van-col>
          <div class="bold">首页功能</div>
        </van-col>
        <van-col>
          <van-row align="center">
            <span class="graytxt f12" v-if="open">已收起{{ myIconList.length }}个首页应用</span>
            <div class="ml10">
              <van-icon v-if="open" name="arrow"/>
              <van-icon v-else name="arrow-down"/>
            </div>
          </van-row>
        </van-col>
      </van-row>
      <div v-if="!open">
        <div>
          <draggable
              v-model="myIconList"
              :list="myIconList"
              :disabled="!editIcon"
              item-key="title"
              class="list-group"
              ghost-class="ghost"
              :move="checkMove"
              @start="dragging = true"
              @end="dragging = false"
          >
            <template #item="{element}">
              <div
                  class="tc pt10 pb10"
                  @click="editIcon==true?minIconHandle(element.code):goPath(element.path,element.code)">
                <van-row justify="center">
                  <van-icon :name="element.icon" size="50" class="ml15"/>
                  <div @click.stop>
                    <van-icon :name="minIcon" size="20" class="minIcon"
                              v-if="editIcon==true"/>
                    <div style="width: 20px; height: 20px"></div>
                  </div>
                </van-row>
                <div class="f12 mt5">{{ element.title }}</div>
              </div>
            </template>
          </draggable>
        </div>
      </div>

    </div>
    <div class="card">
      <van-row justify="space-between" align="center">
        <van-col>
          <div class="bold">全部功能</div>
        </van-col>
      </van-row>
      <div>
        <van-row>
          <van-col span="6" class="tc pt10" v-for="item in staticIconList"
                   @click="editIcon==true?plusIconHandle(item.code):goPath(item.path,item.code)">
            <van-row justify="center">
              <van-icon :name="item.icon" size="50"
                        class="ml15"/>
              <div v-if="editIcon==true" @click.stop>
                <van-icon v-if="myIconCodeList.indexOf(item.code)==-1" :name="plusIcon" size="20" class="minIcon"
                />
                <div v-else style="width: 20px; height: 20px"></div>
              </div>
              <div v-else style="width: 20px; height: 20px"></div>
            </van-row>
            <div class="f12 mt5">{{ item.title }}</div>
          </van-col>
        </van-row>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import draggable from 'vuedraggable'
import NavBarComponent from '../../components/NavBar.vue';
import Navbar from '../../components/NavBar.vue';
import {useRouter} from "vue-router";
import {ref, onMounted} from "vue";
import {showSuccessToast} from "vant";
import plusIcon from '../../assets/icon/icon_model_add.png';
import minIcon from '../../assets/icon/icon_model_min.png';
import {staticCodeIconList, staticIconListA} from "./ts/iconList";
import {showFailToast} from "vant/lib/toast/function-call";

const router = useRouter(); // 路由
const open = ref(false); // 控制开启关闭

const dragging = false;
const checkMove = (e: any) => {
}
// 返回上一页
const back = () => router.back();
// 根据key跳转到对应页
const goPath = (path: any, code: any) => {
  if (editIcon.value == false) {
    // 跳转至飞机购票页
    if (path == '') {
      showSuccessToast('暂未开放');
      return false
    } else if (path == 'expressDeliveryOrCarPage') {
      router.push({name: path, params: {applicationType: code}})
    } else {
      router.push(path)
    }
  } else {

  }

}
const myIconList = ref(); // 用户个人设置的功能数组
const myIconCodeList = ref(); // 用户图表的标识数组
const staticIconList = ref(); // 静态全部功能数组
const oldIconList = ref(); // 未更改的功能数组
const editIcon = ref(false); // 编辑图标
const srcIndex = ref(); // 拖拽位置
const destIndex = ref(); // 拖拽位置
// 编辑图标
const editIconList = () => {
  // 控制编辑功能显隐
  editIcon.value = true;
  oldIconList.value = myIconCodeList.value;
}
// 取消图标
const escIconList = () => {
  // 隐藏操作按钮
  editIcon.value = false;
  // 还原原来的处理数组
  myIconCodeList.value = oldIconList.value;
  // 重新构建图标
  indexIconFormat();
}
// 保存图标
const saveIconlist = () => {
  let codeList: any[] = [];
  myIconList.value.map((e: any) => codeList.push(e.code));
  // 先关闭
  editIcon.value = false;
  // 存储到浏览器存储中去
  sessionStorage.setItem('myIconCodeList', JSON.stringify(codeList));
}
// 减掉我的功能列表的相应方法
const minIconHandle = (item: any) => {
  // 获取到图标code在我的功能数组中的位置并称为key
  let key = myIconCodeList.value.indexOf(item);
  // 进行删除
  myIconCodeList.value.splice(key, 1);
  // 重新构建图标
  indexIconFormat();
}
// 添加到我的功能列表的相应方法
const plusIconHandle = (item: any) => {
  // 添加code至code图标数组
  if (myIconCodeList.value.length < 7) {
    myIconCodeList.value.push(item);
    indexIconFormat();
  } else {
    showFailToast('首页区块最多添加7个应用超出的部分请先移除后添加');
  }
}

// 处理首页功能
const indexIconFormat = () => {
  let codeList = myIconCodeList.value;
  let comeIconList = [];
  for (let i = 0; i < codeList.length; i++) {
    let codeItem = staticIconList.value.find((e: any) => codeList[i] == e.code);
    comeIconList.push(codeItem);
  }
  myIconList.value = comeIconList;
}
onMounted(() => {
  staticIconList.value = staticIconListA;
  myIconCodeList.value = staticCodeIconList;
  indexIconFormat()
})
</script>

<style scoped>
.card {
  margin: 10px 5px;
  padding: 10px;
  background: white;
  border-radius: 8px;
}

.list-group {
  display: flex;
  flex-wrap: wrap;
}

.list-item {
  flex: 1
}

.minIcon {
  float: right;
  position: relative;
  right: 15px;
}
</style>
