<template>
    <van-nav-bar
            title="会议列表"
            left-text="返回"
            left-arrow
            @click-left="onClickLeft"
    />
    <!-- 搜索框 -->
    <van-search v-model="value" placeholder="会议名称、会议号、发起人" />
    <div @click="toDetail">
        <div class="data">
            <div class="year">3月29日 周五</div>
            <div class="year">2024年</div>
        </div>
        <div class="list">
            <div class="list_left">
                <div class="div-bottom">{{number}}</div>
                <div class="f-w-14 div-bottom">特邀会议</div>
                <div><span>时间：16：55</span><span style="margin-left: 20px;">发起人：L。</span></div>
            </div>
            <div class="list_right">
                <van-image :src="arrow" width="20px" height="20px"></van-image>
            </div>
        </div>
    </div>
    <div @click="toDetail">
        <div class="data">
            <div class="year">3月30日 周六</div>
            <div class="year">2024年</div>
        </div>
        <div class="list">
            <div class="list_left">
                <div class="div-bottom">{{number}}</div>
                <div class="f-w-14 div-bottom">网络研讨会</div>
                <div><span>时间：16：55</span><span style="margin-left: 20px;">发起人：L。</span></div>
            </div>
            <div class="list_right">
                <van-image :src="arrow" width="20px" height="20px"></van-image>
            </div>
        </div>
    </div>
    <div @click="toDetail">
        <div class="data">
            <div class="year">3月31日 周日</div>
            <div class="year">2024年</div>
        </div>
        <div class="list">
            <div class="list_left">
                <div class="div-bottom">{{number}}</div>
                <div class="f-w-14 div-bottom">L.的快速会议</div>
                <div><span>时间：16：55</span><span style="margin-left: 20px;">发起人：L。</span></div>
            </div>
            <div class="list_right">
                <van-image :src="arrow" width="20px" height="20px"></van-image>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
    import arrow from '../../assets/img/arrow.png'
    import { ref } from 'vue';
    import { useRouter } from 'vue-router';
    const router = useRouter()
    const value = ref('')
    const number = ref<number>()
    // const list = ref([])
    const onClickLeft = () => history.back();
    const toDetail = ()=>{
        router.push({
            name:'mettingDetail'
        })
    }
    const date = new Date();
    const year = date.getFullYear(); // 年
    const month = date.getMonth() + 1; // 月份从0开始，所以要加1
    const day = date.getDate(); // 日
    const hours = date.getHours(); // 小时
    const minutes = date.getMinutes(); // 分钟
    const seconds = date.getSeconds(); // 秒

    // 将年月日时分秒组合成一个数字
    const combinedNumber = parseInt(`${year}${month}${day}${hours}${minutes}${seconds}`);

    console.log(combinedNumber); // 输出合并后的数字时间]
    number.value = combinedNumber
</script>
<style lang="less">
    .data{
        display: flex;
        justify-content: space-between;
        padding: 10px;
        .year{
            color: #999a9e;
        }
    }
    .list{
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding:10px 20px;
        background-color: #fff;
        .list_left{
            font-size: 14px;
            color: #999a9e;
            .div-bottom{
                margin-bottom: 8px;
            }
            .f-w-14{
                font-size: 16px;
                font-weight: 550;
                color: #000;
            }
        }
    }
</style>
