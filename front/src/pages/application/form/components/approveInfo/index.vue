<template>
  <van-sticky>
    <van-nav-bar
        title="审批信息"
        left-arrow
        @click-left="$emit('selectOver',{})"
    />
  </van-sticky>
  <div class="formCard" >
    <div>
      <van-row align="center">
        <span class="ml10">审批进度</span>
      </van-row>
      <div style="margin-top: 10px;margin-bottom: 10px">
        <van-steps direction="vertical" :active="nextApprIndex">
          <van-step>
            <van-row>
              <van-col :span="4">
                <div class="f14 graytxt bold">发起</div>
              </van-col>
              <van-col :span="20">
                <div style="padding-top: 3px;padding-bottom: 3px;background: #f6f6f6">
                  <div class="approveCard">
                    <van-row>
                      <van-col :span="10" class="f12 ml5">{{ userInfoData.userInfo.userName }}({{userInfoData.userInfo.jobName}})</van-col>
                    </van-row>
                  </div>
                </div>
              </van-col>
            </van-row>
          </van-step>
          <van-step v-for="item in approveInfoList">
            <van-row>
              <van-col :span="4">
                <van-row align="center">
                  <van-col>
                    <div class="f14 graytxt bold">
                      <div v-if="item.isAllPass==null">单审</div>
                      <div v-else>{{ item.isAllPass == '1' ? '或审' : '会审' }}</div>
                      <div>
                        <span v-if="item.result=='ok'" class="ok f12 bold">通过</span>
                        <span v-else-if="item.result=='ng'" class="ng f12 bold">已拒绝</span>
                        <span v-else class="graytxt f12 bold">待审批</span>
                      </div>
                    </div>
                  </van-col>
                  <van-col>
                    <van-icon name="info-o" @click="approveToast(item.isAllPass)"/>
                  </van-col>
                </van-row>
              </van-col>
              <van-col :span="20">
                <div style="padding-top: 3px;padding-bottom: 3px;background: #f6f6f6">
                  <div v-for="key in item.children">
                    <div class="approveCard">
                      <van-row>
                        <van-col :span="12" class="f12 ml5">{{ key.toUserName }}({{ key.toRole }})</van-col>
                      </van-row>
                    </div>
                  </div>
                </div>
              </van-col>
            </van-row>
          </van-step>
        </van-steps>
      </div>

    </div>

  </div>
</template>

<script setup lang="ts">
  import NavBar from '../../../../../components/navBar/index.vue';
  import {inject, onMounted, ref} from "vue";
  import {log} from "echarts/types/src/util/log";
  import {defineEmits} from "vue/dist/vue";
  import {showToast} from "vant";
  const nextApprIndex = ref('');
  let userInfoData: any = inject("userInfo"); // 用户信息
  const emit=defineEmits(['selectOver']);
  interface Props{}
  const props=defineProps<{
    code:any,// 使用类型
    approveInfoList:[], //
  }>()
  const showPopover = ref(false);
  const actions = [
    { text: '选项一' },
  ];
  // 根据显示状态
  const approveToast=(pass:any)=>{
    if(pass==null){
      showToast('单审：该审批单为单人审批，同意即直接通往下级，拒绝则返回');
    }else if(pass='1'){
      showToast('或审：该申请单由当前级别所有审批人进行审批，该级有一位同意即通往下级，全部拒绝则返回');
    }else{
      showToast('会审：该申请单由当前级别所有审批人进行审批，所有审批人都同意则通过')
    }
  }
  onMounted(()=> {
  })
</script>

<style scoped>
.approveCard {
  margin: 5px;
  padding-top: 10px;
  padding-bottom: 5px;
  background: white;
  border-radius: 5px;
}
</style>
