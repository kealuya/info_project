<template>
  <div class="container">
    <van-nav-bar
        left-arrow
        left-text="返回"
        title="价值申请"
        @click-left="onClickLeft"
    />
    <van-tabs v-model:active="active" @change="tabChange">
      <van-tab title="待申请列表">

      </van-tab>
      <van-tab title="申请记录">

      </van-tab>


    </van-tabs>
    <div :class="{ 'list1': tabIndex }">
      <div>
        <van-pull-refresh v-model="isLoading" success-text="刷新成功" @refresh="onRefresh">
          <div v-if="searchLoading && list.length==0" class="h-67">
            <van-loading :vertical="true" color="#1989fa" type="spinner">加载中...</van-loading>
          </div>
          <van-radio-group v-if="!isShowImg" v-model="selectedItem" @change="handleRadioChange">
            <van-list
                    class="list"
                v-model:loading="loading"
                :finished="finished"
                :immediate-check="false"
                :offset="30"
                finished-text="没有更多了"
                loading-text="加载中..."
                @load="getListLoad"

            >
              <div v-for="item in list" :key="item.worthId" class="list_card">
                <div class="flex-space m-b-10">
                  <div style="display:flex">
                    <van-image :src="Money" height="20" width="20"></van-image>
                    <div class="list_title">{{ item.worthTitle }}</div>
                  </div>
                  <van-radio :name="item.worthId"/>
                  <!--                  <van-checkbox v-model="item.isCheck" @change="changeCheck"></van-checkbox>-->
                </div>
                <van-divider/>
                  <div class="f-z-12 m-b-10">价值ID:{{ item.worthId }}</div>
                <div class="f-z-12 m-b-10">任务评分：{{ item.worthScore }}分</div>
                <div class="flex-space f-z-12 m-b-10">
                  <div>预计可申请金额：{{ item.money }}</div>
                </div>
                <div class="flex-space f-z-12 m-b-10">
                  <!--              <div></div>-->
                  <div>创建时间：{{ item.createTime }}</div>
                </div>
              </div>
            </van-list>
          </van-radio-group>
          <van-row v-else-if="isShowImg" align="center" class="tc" justify="center">
            <div class="noData">
              <img src="../../assets/img/zanwupiaoju.png" style="height:50vw;width:80vw;object-fit: contain"/>
              <div class="banner">暂无数据</div>
            </div>
          </van-row>
        </van-pull-refresh>
      </div>
      <div class="footer_btn" v-if="list.length>0">
        <van-button block type="primary" @click="handleSubmit" :loading="isShow" loading-text="提交中">
          提交申请
        </van-button>
      </div>
    </div>
    <div :class="{ 'list1': !tabIndex}" >
      <van-pull-refresh v-model="isLoading" success-text="刷新成功" @refresh="onRefresh">
        <div v-if="searchLoading && list.length==0" class="h-67">
          <van-loading :vertical="true" color="#1989fa" type="spinner">加载中...</van-loading>
        </div>
        <van-list
                class="listReacord"
            v-if="!isShowImg"
            v-model:loading="loading"
            :finished="finished"
            :immediate-check="false"
            :offset="80"
            finished-text="没有更多了"
            loading-text="加载中..."
            @load="getListLoad"

        >
          <div v-for="item in list" class="list_card" @click="handleJump(item.worthId)">
            <div style="display:flex">
              <van-image :src="Money" height="20" width="20"></van-image>
              <div class="list_title">{{ item.worthTitle }}</div>
            </div>
            <van-divider/>
              <div class="f-z-12 m-b-10">价值ID:{{ item.worthId }}</div>
            <div class="f-z-12 m-b-10">任务评分:{{ item.worthScore }}</div>
            <div class="f-z-12 m-b-10">预计可申请金额:{{ item.money }}</div>
            <div class="flex-space f-z-12">
              <!--            <div>已关联任务100002</div>-->
              <div>已完成时间：2024-04-12 15:20:15</div>
            </div>
            <div v-if="item.status=='0'" class="list_icon">
              <van-image :src="waitshenpi" height="60" width="60"></van-image>
            </div>
            <div v-else-if="item.status =='1'" class="list_icon">
              <van-image :src="yitongyi" height="60" width="60"></van-image>
            </div>
          </div>
        </van-list>
        <van-row v-else-if="isShowImg" align="center" class="tc" justify="center">
          <div>
            <img src="../../assets/img/zanwupiaoju.png" style="height:50vw;width:80vw;object-fit: contain"/>
            <div class="banner">暂无数据</div>
          </div>
        </van-row>
      </van-pull-refresh>
    </div>

  </div>
  <!--  <div class="mb10"></div>-->
</template>
<script lang="ts" setup>
import waitshenpi from '../../assets/img/waitshenpi_icon.png'
import yitongyi from '../../assets/img/yitongyi_icon.png'
import Money from '../../assets/img/money.png'
import businessIcon from '../../assets/img/business_icon.png';
import {ref, computed, onMounted, inject} from 'vue';
// import userRouter from 'user'
import {useRouter} from "vue-router";
import {getWorthList, applyWorth} from '../../services/task-processing/index'
import {showSuccessToast} from "vant";
import {showFailToast} from "vant/es";

const selectedItem = ref<any>()
const loading = ref(false);
const finished = ref(false);
const totalCount = ref<number>(0)
const pageCount = ref<number>(0)
const isLoading = ref<boolean>(false)  //控制 下拉刷新
const tabIndex = ref(false)
const searchLoading = ref<Boolean>(false)  //控制 数据没请求回来的 loading
const finishText = ref<string>('没有更多了')
const isShowImg = ref<boolean>(false)  //数据为空的时候 控制图片的显示隐藏
let lastRefreshTime = 0;
const isShow = ref<boolean>(false)
const refreshInterval = 15000; // 15秒
interface listValueType {
  currentPage: number,//	当前页
  pageSize: number,//	每页显示条数
  corpCode: string | undefined | null	//	企业code
  status: string,//	是	状态，0未申请，1已申请，当前接口传0
  userId: string//	用户ID
}

interface paramsValueType {
  worthId: string,//	价值ID，对应任务的ID
  corpCode: string | null,//	MX	String	是	企业code
  userId: string,//	是	用户ID
  userName: string,//	否	用户姓名
  userMobile: string	//否	用户手机号
}
interface applicationType{
    worthId: string, //价值ID，对应任务的ID
    worthTitle: string, //价值title，对应任务的title
    worthScore: string, //价值评分
    status: string, //状态，0未申请，1已申请
    money: string, //价值金额
    userId: string, //用户ID
    userName: string, //用户姓名
    userMobile: string, //用户手机号
    corpName: string, //企业名称
    corpCode: string, //企业code
    createTime: string, //创建时间
    creater: string, //创建人，默认登录用户
    bz1:string, //备注
    bz2:string,
    bz3:string
}
let params = ref<listValueType>({
  currentPage: 1,//	当前页
  pageSize: 10,//	每页显示条数
  corpCode: '',	//	企业code
  status: '0',//	是	状态，0未申请，1已申请，当前接口传0
  userId: ''//	用户ID
})
let paramsValue = ref<paramsValueType>({
  worthId: '',//	价值ID，对应任务的ID
  corpCode: '',
  userId: '',
  userName: '',
  userMobile: ''
})
const checked = ref(false);
const active = ref(0);
const router = useRouter()
const list = ref<applicationType[]>([])
const lists = ref<any>([])
const handleRadioChange = (value: any) => {
  console.log('value', value)
  paramsValue.value.worthId = value
}
//点击了tab切换
const tabChange = async (name: number) => {
  // console.log('检测到了切换')
  if (name == 1) {
    params.value.currentPage = 1
    params.value.status = '1'
    list.value = []
    //这个必须加
    loading.value = false
    finished.value = false
    await getValueList()
    tabIndex.value = true
  } else {
    params.value.status = '0'
    params.value.currentPage = 1
    list.value = []
    //这个必须加
    loading.value = false
    finished.value = false
    await getValueList()
    tabIndex.value = false
  }
}
// 下拉刷新
// const onRefresh = async () => {
//   loading.value = true
//   isLoading.value = true
//   // finished.value = false
//   params.value.currentPage = 1
//   list.value.length = 0
//   await getValueList()
// }
const onRefresh = async () => {
  const currentTime = Date.now();
  if (currentTime - lastRefreshTime < refreshInterval) {
    showFailToast("刷新操作太频繁，请稍后再试！");
    isLoading.value = false; // 关闭下拉刷新的 loading
    return; // 直接返回，不执行加载数据的操作
  }
  lastRefreshTime = currentTime;
  isLoading.value = true; // 打开下拉刷新的 loading
  params.value.currentPage = 1; // 重置页码
  list.value = []; // 清空表单
  await getValueList(); // 重新加载数据
  showSuccessToast("刷新成功！");

  isLoading.value = false; // 关闭下拉刷新的 loading
};







//勾选了价值
const changeCheck = () => {
  let arr = list.value.filter((item: any) => item.isCheck);
  console.log('arr', arr)
  // let worthId = arr.map((item: any) => {
  //   return item.worthId
  // })


}
//触底事件
const getListLoad = async () => {

  if (list.value.length < totalCount.value && params.value.currentPage < pageCount.value) {
    params.value.currentPage = params.value.currentPage + 1
    // setTimeout(() => {
    //   getValueList()
    // },500)
    await getValueList()
  }
}
//业务数据列表
//价值申请数据列表
const handleJump = (id:string | undefined) => {
  router.replace({path:'/valueDetails',query:{id}})
}
const onClickLeft = () => {
  router.replace('/homeNew')
}
//点击了提交申请按钮
const handleSubmit = async () => {
    isShow.value = true
  params.value.status = '1'
  tabIndex.value = true
  params.value.currentPage = 1
  loading.value = true
  finished.value = false
  list.value.length = 0
  applyWorth(paramsValue.value).then(async (res: any) => {
    if (res.success) {
        active.value = 1
        isShow.value = false
      showSuccessToast(res.msg)
    }
  }).finally(async()=>{
      active.value = 1
    await getValueList()
  })
}
//获取价值列表
const getValueList = () => {
    loading.value = true
    // searchLoading.value = true
    // loading.value = true
  getWorthList(params.value).then((res: any) => {
    if (res.success) {
      list.value = list.value.concat(res.data.worthList)
      totalCount.value = res.data.totalCount  //总条数
      //todo 总页
      pageCount.value = res.data.pageCount //总页码
      setInterval(() => {
        isLoading.value = false
      }, 2000)
      if (list.value.length == totalCount.value) {
        loading.value = false   //关闭下拉刷新的加载
        finished.value = true
      }else{
        loading.value = false
        finished.value = false
      }
      if (res.data.totalCount == 0) {
        isShowImg.value = true
        finishText.value = ''
      } else {
        isShowImg.value = false
        finishText.value = '没有更多数据了'
      }
    }
  }).finally(() => {
    searchLoading.value = false
      setTimeout(()=>{
          isLoading.value = false
      },5000)
  })
}
onMounted(async () => {
  let userInfoData: any = inject("userInfo"); // 取出用户信息用于调用接口
  params.value.corpCode = localStorage.getItem('corpCode') //从本地获取corpCode
  //从本地获取corpCode
  params.value.userId = userInfoData.userInfo.userId
  paramsValue.value.userId = userInfoData.userInfo.userId
  paramsValue.value.userName = userInfoData.userInfo.userName
  paramsValue.value.userMobile = userInfoData.userInfo.userMobile
  paramsValue.value.corpCode = localStorage.getItem('corpCode') //从本地获取corpCode
  await getValueList()
})
</script>
<style lang="less" scoped>
.container {
  height: 100vh;

  .header {
    position: fixed;
    z-index: 99;
    width: 100%;
    height: 60px;
    background-color: #fff;
    // padding: 0 10px;
    font-size: 14px;
    display: flex;
    text-align: center;
    height: 50px;
    line-height: 60px;
    border-bottom: 1px solid #f0f0f0;
    //margin-bottom: 50px;
    .menuItems {
      font-size: 14px;
      width: 50%;
      display: flex;
      justify-content: center;
      align-items: center;
      // border: 1px solid red;
    }
  }

  .footer_btn {
    width: 96vw;
    padding: 0 2vw;
    margin-bottom: 2vh;
  }
.noData{
  min-height:calc(100vh - var(--van-nav-bar-height) - var(--van-tabs-line-height) - var(--van-button-default-height) - 2vh)
}
  .box {
    height: 50px;
  }

  .list1 {
    display: none;
  }
  .h-67 {
    height: 20vh;
    align-items: center;
    display: flex;
    justify-content: center;
  }
  .list {
    height: calc(100vh - var(--van-nav-bar-height) - var(--van-tabs-line-height) - var(--van-button-default-height) - 2vh);
    overflow-y: auto;

    .list_card {
      position: relative;
      width: 92vw;
      margin: 2vw;
      padding: 2vw;
      background-color: #FFFFFF;
      border-radius: 10px;

      .flex-space {
        display: flex;
        justify-content: space-between;
      }

      .list_title {
        font-size: 0.875em;
        font-weight: 550;
        margin-left: 2vw;
      }

      .f-z-12 {
        font-size: 0.75em;
        color: #4B5563;
      }

      .m-b-10 {
        margin-bottom: 2vw;
      }

      .list_icon {
        position: absolute;
        right: 30px;
        top: 40px;
      }
    }

  }

  .listReacord {
    height: calc(100vh - var(--van-nav-bar-height) - var(--van-tabs-line-height) - var(--van-tabbar-height));
    overflow-y: auto;

    .list_card {
      position: relative;
      width: 92vw;
      margin: 2vw;
      padding: 2vw;
      background-color: #FFFFFF;
      border-radius: 10px;

      .flex-space {
        display: flex;
        justify-content: space-between;
      }

      .list_title {
        font-size: 0.875em;
        font-weight: 550;
        margin-left: 2vw;
      }

      .f-z-12 {
        font-size: 0.75em;
        color: #4B5563;
      }

      .m-b-10 {
        margin-bottom: 2vw;
      }

      .list_icon {
        position: absolute;
        right: 30px;
        top: 40px;
      }
    }

  }

  .mainButton {
    border-radius: 10px;
    padding: 2vw;
    width: 42vw;
    height: 50px;
    background: url('../../assets/img/radio.png') no-repeat center;
    background-size: 100vw 10vh;
    //background-image: linear-gradient(to right bottom, #0080FF, #6AB5FF);
  }

  .mainButton_sec {
    border-radius: 10px;
    padding: 2vw;
    width: 42vw;
    height: 50px;
    background: url('../../assets/img/radio.png') no-repeat center;
    background-size: 100vw 10vh;
    //background-image: linear-gradient(to right bottom, #33B2AD, #3BD7D3);
  }
}

:deep(.van-dropdown-menu__bar) {
  box-shadow: none;
}

:deep(.van-divider) {
  margin: 2vw;
}

::v-deep(.van-nav-bar__content) {
  background-color: #0088ff;
}

::v-deep(.van-nav-bar__title) {
  color: #ffffff;
}

::v-deep(.van-nav-bar__text) {
  color: #ffffff;
}

::v-deep(.van-icon-arrow-left:before) {
  color: #ffffff;
}
</style>
