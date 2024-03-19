<template>
    <van-nav-bar
            title="发票详情"
            left-text="返回"
            left-arrow
            @click-left="onClickLeft"
    />
    <van-row justify="center" align="center">
      <div class="wrap" style="width: 80vh">
        <p style="color: #A86A64;font-size: 18px">{{getInvoiceTypeList(detail.type)}}</p>
        <!-- 发票展示-->
        <div class="pb10" >
          <quotaTicket v-if="detail.type =='1'" :get-data="detail" />
          <Etron v-else-if="detail.type =='0'"  :get-data="detail" />
          <airTicket v-else-if="detail.type =='5'"  :get-data="detail" />
          
          <ReilwayTickey v-else-if="detail.type =='8'"  :get-data="detail" />
          
          <passableTicket v-else :get-data="detail" />
        </div>
        <div class="mt5"></div>
      </div>
      <div class="wrap_btn" style="width: 80vh">
        <van-row @click="imgDetail" justify="space-between" align="center">
          <van-col span="10"  class="graytxt" :offset="2">查看图片</van-col>
          <van-col span="2" class="detail">
            <van-icon name="arrow" class="graytxt" />
          </van-col>
        </van-row>
<!--        <van-overlay :show="show" @click="show = false">-->
<!--&lt;!&ndash;          <template #default>&ndash;&gt;-->
<!--&lt;!&ndash;            <van-image&ndash;&gt;-->
<!--&lt;!&ndash;                width="100%"&ndash;&gt;-->
<!--&lt;!&ndash;                height="100%"&ndash;&gt;-->
<!--&lt;!&ndash;                fit="scale-down"&ndash;&gt;-->
<!--&lt;!&ndash;                :src="url"&ndash;&gt;-->
<!--&lt;!&ndash;            >&ndash;&gt;-->
<!--&lt;!&ndash;              <template v-slot:error>&ndash;&gt;-->
<!--&lt;!&ndash;                <img src="../../assets/img/jiazaishibai.png" style="object-fit: contain;width:100vw;height:80vh;position:relative"/>&ndash;&gt;-->
<!--&lt;!&ndash;                <div class="f16" style="position: absolute;bottom:190px">获取图片信息超时，请稍后再试</div></template>&ndash;&gt;-->
<!--&lt;!&ndash;              <template v-slot:loading>&ndash;&gt;-->
<!--&lt;!&ndash;                <van-loading type="circular" size="50" />&ndash;&gt;-->
<!--&lt;!&ndash;              </template>&ndash;&gt;-->
<!--&lt;!&ndash;            </van-image>&ndash;&gt;-->
<!--&lt;!&ndash;          </template>&ndash;&gt;-->
<!--        </van-overlay>-->
      </div>
    </van-row>
</template>
<script setup lang="ts">
    import PassableTicket from "./component/passableTicket.vue";
    import 'vant/es/image-preview/style';
    import { ImagePreview } from 'vant';
    const show1=ref(true)
    import { useRouter} from "vue-router";
    import {reactive, onBeforeMount, ref, inject} from 'vue'
    import {getImgDetail} from '../../services/invoice';
    import {invoiceTypeList}from '../invoice/component/ts/invoiceTypeList'
    import moment from "moment";
    import Etron from './component/electronicTicket.vue'
    import ReilwayTickey from './component/reilwayTickey.vue'
    import AirTicket from "./component/airTicket.vue";
    import QuotaTicket from "./component/quotaTicket.vue";
const router=useRouter();
//遍历根据type值展示相应的发票名称
var invoiceTitle='';
    let userInfoData: any = inject("userInfo");
function getInvoiceTypeList (data: any) {
    let invoiceType=invoiceTypeList.find((e)=>data==e.value);
    invoiceTitle=invoiceType==undefined?'':invoiceType.name;
    return invoiceTitle;
    }
let show=ref(false);
//存储图片地址
    let url=ref('')
//存储params数据
    let detail=reactive<any>({});
    let postInvoiceList = reactive({
        'corpCode': userInfoData.userInfo.corpCode,
        'empCode': userInfoData.userInfo.empCode,
        'uuid': '',
        'dzfph':'',
    })
    let imgUrl=ref([])
let imgDetail=async()=>{
    show.value=true
    let res:any=await getImgDetail(postInvoiceList);
        url.value=res;
        imgUrl.value=[];
        let addUrl=url.value as never;
        imgUrl.value.push(addUrl);
        ImagePreview( {
          images:imgUrl.value,
          showIndex:false,
        });
    // router.push({name:'imgDetail'});
};
    let date=ref();
onBeforeMount(async ()=>{
  let invoiceInfo:any=sessionStorage.getItem('invoiceDetail');
    detail =JSON.parse(invoiceInfo);
    postInvoiceList["uuid"]=detail["fjid"]
    postInvoiceList["dzfph"]=detail["dzfph"]
    detail.kprq=moment(detail.kprq).format("yyyy-MM-DD");
    // detail.kprq=formatDateTime(date.value)
})
    //返回
    const onClickLeft = () => history.back();
    //科学计数法
    function saveMoneyTwoNum(num:any){
        if((num+'').indexOf('.') != -1){
            if((num+'').split(".")[1].length == 1){
                num = num+'0'
            }else{
                num = num;
            }
        }else{
            num = Number(num).toFixed(2)
        }
        return num;
    }
</script>

<style scoped>
  .wrap{
      padding-top:5px;
    margin-top: 10px;
    margin-left: 10px;
    margin-right: 10px;
    padding-right: 15px;
    background-color: #fff;
    /*margin:25px 20px 0px 20px;*/
    border-radius: 10px;
    box-sizing: border-box;
      position: relative;
  }
  .wrap_btn{
    margin-top: 8px;
    margin-left: 10px;
    margin-right: 10px;
    padding-top: 10px;
    /*padding-right: 15px;*/
    background-color: #fff;
    /*margin:25px 20px 0px 20px;*/
    border-radius: 10px;
  }
  p{
      background: url("../.././assets/img/fapiao_bg.png") no-repeat center;
      background-size:90px 40px;
      text-align: center;
      padding: 20px;
      position: relative;
  }
  .detail{
    text-align: left;
    font-weight: normal;
    font-size:14px
  }
  /*  .van-row  .van-col:nth-child(2){*/
  /*      color:#ccc;*/
  /*      font-size:15px;*/
  /*      padding:5px 0*/
  /*  }*/
  /*.van-row  .van-col:nth-child(3){*/
  /*    text-align: left;*/
  /*    margin-left:25px;*/
  /*    height: auto;*/
  /*    width:auto;*/
  /*    margin-top: 2.5px;*/
  /*}*/
    .van-row{
        margin-bottom: 10px;
    }
    .yzImg{
        width:130px;
        height:100px;
        position: absolute;
        bottom: 80px;
        right:25px;
    }
 /deep/  .van-image__loading{
     background:rgba(0,0,0,0) ;
 }

</style>
