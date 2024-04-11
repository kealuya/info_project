import {defineStore} from 'pinia';

//这里官网是单独导出  是可以写成默认导出的  官方的解释为大家一起约定仓库用use打头的单词 固定统一小仓库的名字不易混乱
export const useInvoiceData = defineStore("invoiceData", {
    persist: {
        enabled: true,//开启数据持久化
        strategies: [
            {
                key: 'invoice',//给一个要保存的名称
                storage: sessionStorage,//sessionStorage / localStorage 存储方式
            }
        ]
    },
    state: () => ({
        invoiceData: {}
    }),
    getters: {

    },
    actions: {
        setData(data: any) {
            this.invoiceData = data;
        },
    },
})
// 用户信息
export const userInfoData = defineStore("userInfo", {
    persist: {
        enabled: true,//开启数据持久化
        strategies: [
            {
                key: 'userInfo',//给一个要保存的名称
                storage: sessionStorage,//sessionStorage / localStorage 存储方式
            }
        ]
    },
    state: () => ({
        userInfo: {},
        functionConfig: [],
        userTravelConfig: [],
        unReadMessageNum:0,//消息未读数量
        pullMessage:{},
    }),
    getters: {

    },
    actions: {
        setData(data: any) {
            this.userInfo = data;
        },
        setFunctionConfig(data: any) {
            this.functionConfig = data;
        },
        setUserTravelConfig(data: any) {
            this.userTravelConfig = data;
        },
        setUnReadMessageNum(data: any){
            this.unReadMessageNum = data
        },
        setPullMessage(data: any){
            this.pullMessage = data

        }

    },
})
export const heightData = defineStore("heightData", {
    persist: {
        enabled: true,//开启数据持久化
        strategies: [
            {
                key: 'heightData',//给一个要保存的名称
                storage: sessionStorage,//sessionStorage / localStorage 存储方式
            }
        ]
    },
    state: () => ({
        height:665

    }),
    getters: {

    },
    actions: {
        setFunctionConfig(data: any) {
            this.height = data;
        }
    },
})