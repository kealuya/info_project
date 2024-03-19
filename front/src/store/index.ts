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
                storage: localStorage,//sessionStorage / localStorage 存储方式
            }
        ]
    },
    state: () => ({
        userInfo: {}
    }),
    getters: {

    },
    actions: {
        setData(data: any) {
            this.userInfo = data;
        },
    },
})
