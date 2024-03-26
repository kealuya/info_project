<template>
    <div>
        <van-row>
            <NavBar :title="'选择发票信息'" :leftText="'返回'" :leftArrow="true" :rightText="'确定'" @leftEvent="back()"
                    @rightEvent="submit()"/>
        </van-row>
        <van-row style="margin-top: 8px;margin-bottom: 8px;">
<!--            <van-checkbox-group v-model="checked">-->
<!--                    <van-cell-->
<!--                            v-for="(item, index) in cardList"-->
<!--                            clickable-->
<!--                            :key="item"-->
<!--                            :title="`复选框 ${item}`"-->
<!--                            @click="toggle(index)"-->
<!--                            style="margin-top: 10px;margin-left: 20px"-->
<!--                    >-->
<!--                        <template #right-icon>-->
<!--                            <van-checkbox-->
<!--                                    :name="item"-->
<!--                                    :ref="el => checkboxRefs[index] = el"-->
<!--                                    @click.stop-->
<!--                            />-->
<!--                        </template>-->
<!--                    </van-cell>-->
<!--            </van-checkbox-group>-->
            <!--            <van-checkbox-group  v-model="checked"  style="margin-left: 10px" >-->
            <!--                <div v-for="(item,index) in cardList">-->
            <!--                    <van-checkbox :name="item.fjid"  >-->
            <!--                        <van-col span="21">-->
            <!--                            <div class="cardList" :style="{width:cardWidth-20+'px'}">-->
            <!--                                <div :style="{width:cardWidth-20+'px',minHeight:'30px',marginTop:'10px',marginLeft:'10px'}">-->
            <!--                                    <van-row>-->
            <!--                                        <van-col span="15" class="fw800 f14 pt5">-->
            <!--                                            {{item.dzfph}}-->
            <!--                                        </van-col>-->
            <!--                                        <van-col span="9" class="money fw800">-->
            <!--                                            {{item.invName}}-->
            <!--                                        </van-col>-->
            <!--                                    </van-row>-->
            <!--                                    <van-row justify="left" class="pt10 pb10">-->
            <!--&lt;!&ndash;                                        <van-icon name="balance-list" color="#1989fa" class="mt3"/>&ndash;&gt;-->
            <!--&lt;!&ndash;                                        <div class="f14 pl5">{{item.expenseTypeName}}</div>&ndash;&gt;-->
            <!--&lt;!&ndash;                                        <van-icon name="clock" color="#FB5632" class="pl5 mt3"/>&ndash;&gt;-->
            <!--&lt;!&ndash;                                        <div class="f14 pl5 mt3">{{item.occurrenceTime}}</div>&ndash;&gt;-->
            <!--                                    </van-row>-->
            <!--                                </div>-->
            <!--                            </div>-->
            <!--                        </van-col>-->
            <!--                    </van-checkbox>-->
            <!--                </div>-->
            <!--            </van-checkbox-group >-->
        </van-row>
    </div>
</template>

<script setup lang="ts">
    import {ref, onMounted, onBeforeMount, inject} from 'vue'
    import {getInvoiceWbList} from "../../services/accountBook"
    import route from "../../router";

    let realWidth = window.document.documentElement.offsetWidth
    let cardWidth = window.document.documentElement.offsetWidth - 30
    const page = ref(1)
    const pageSize = ref(8)
    const total = ref(0)
    let userInfoData: any = inject("userInfo");
    const cardList = ref([])
    const checked = ref([])
    // const router = useRouter()
    const bookMessage = ref({})
    const checkboxRefs = ref([]);
    const toggle = (index: any) => {
        // checkboxRefs.value[index].toggle();
    };


    onBeforeMount(() => {
        checkboxRefs.value = [];

        getInvoiceWbListMth()

    })
    onMounted(() => {
        // if(route.currentRoute.value.params.bill){
        //     let billList = JSON.parse(route.currentRoute.value.params.bill as never)
        //
        //     let idList: any[] =[]
        //     billList.map((item: any)=>{
        //         idList.push(item.billId)
        //     })
        //     checked.value = idList as never
        // }
    })


    function back() {

        route.go(-1)
    }

    function submit() {
        let costList: any[] = []
        cardList.value.map((item: any) => {
            if (checked.value.indexOf(item.billId as never) != -1) {
                costList.push(item)
            }
        })
        let bookAccountMessage = {
            identifying: 'costInformation',
            bill: costList
        }

        route.push({
            name: 'dailyReimbursement',
            params: {
                bill: JSON.stringify(bookAccountMessage)
            },
        })
    }

    function clickRadio(item: any) {
        bookMessage.value = item
    }

    async function getInvoiceWbListMth() {

        let param = {
            page: page.value.toString(),
            row: pageSize.value.toString(),
            zt: "0",
            type: "",
            empCode: userInfoData.userInfo.empCode,
            corpCode: userInfoData.userInfo.corpCode,
            equipmentType: "PC"
        }
        let res: any = await getInvoiceWbList(param)
        if (res.success) {
            cardList.value = res.data.data.data
            total.value = res.data.total
        }
    }
</script>

<style lang="less" scoped>

    .mlp4 {
        margin-left: 4%;
    }

    .fw800 {
        font-weight: 800;
    }

    .cardList {
        /*margin-left: 2%;*/
        /*margin-right: 2%;*/
        /*width: 200px;*/
        border: 1px solid #1989fa;
        min-height: 30px;
        border-radius: 3px;
        margin-top: 10px;
    }

    .mt3 {
        margin-top: 3px;

    }

</style>
