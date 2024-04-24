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
    getters: {},
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
        unReadMessageNum: 0,//消息未读数量
        pullMessage: {},
    }),
    getters: {},
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
        setUnReadMessageNum(data: any) {
            this.unReadMessageNum = data
        },
        setPullMessage(data: any) {
            this.pullMessage = data

        }

    },
})
//任务详情 数据
export const taskData = defineStore("taskData", {
    persist: {
        enabled: true,//开启数据持久化
        strategies: [
            {
                key: 'taskInformation',//给一个要保存的名称
                storage: sessionStorage,//sessionStorage / localStorage 存储方式
            }
        ]
    },
    state: () => ({
        taskId: {},
        taskTitle: {},
        taskContent:{}

    }),
    getters: {},
    actions: {
        setTaskId(data: any) {
            this.taskId = data
        },
        setTaskTitle(data: string) {
            this.taskTitle = data
        },
        setTaskContent(data: string) {
            this.taskContent = data
        },
        getTaskId() {
            return this.taskId
        },
        getTaskTitle(){
            return this.taskTitle
        },
        getTaskContent(){
            return this.taskContent
        },
    },
})
//会议的数据  勾选的数据（业务内容）
export const meetingData = defineStore("meetingData", {
    persist: {
        enabled: true,//开启数据持久化
        strategies: [
            {
                key: 'meetingCheckList',//给一个要保存的名称
                storage: sessionStorage,//sessionStorage / localStorage 存储方式
            }
        ]
    },
    state: () => ({
      checkList:[]

    }),
    getters: {},
    actions: {
       setMeetingData(data:any){
           console.log('data',data)
           this.checkList = data
       },
        getMeetingData(){
          return this.checkList
        }
    },
})
