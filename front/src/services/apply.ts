import request from "./request";
//提交请假申请单
export function createLeaveApply(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createLeaveApply',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取默认抄送人
export function getDefaultCcByTypeAndDept(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/CcUser/getDefaultCcByTypeAndDept',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取审批数组
export function getApprovalFlowList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/examine/getApprovalFlowList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 根据flowId获取审批详情
export function getApprovalFlowHistory(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/examine/getApprovalFlowHistory',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

// 获取人员列表
export function getEmpList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/emp/getEmpList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 申请单详情
export function getLoanApplyDetail(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/getLoanApplyDetail',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 审批历史
export function approvalHistory(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/examine/approvalHistory',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 删除申请单（仅限于加班申请免审）
export function deleteWorkInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/deleteWorkInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 撤销申请单
export function approvalRevoke(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/examine/approvalRevoke',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//提交加班申请单
export function createWorkApply(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createWorkApply',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 提交用章申请单
export function createSealApply(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createSealApply',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 提交补卡申请单
export function createCardReplacement(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createCardReplacement',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 提交物品领用
export function createArticleRequisition(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createArticleRequisition',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 提交外出办公
export function createGoOutToWork(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createGoOutToWork',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 提交董事签批
export function createShareholderSign(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createShareholderSign',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 客户宴请
export function createCustomercBanquet(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createCustomercBanquet',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 会议记录
export function createConferenceRecords(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createConferenceRecords',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 物品申购
export function createShopApply(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createShopApply',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 用车
export function createCarApply(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createCarApply',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 快递寄送
export function createLogApply(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createLogApply',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 费用申请
export function createDailyReimbursemen(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createDailyReimbursement',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取单据列表
export function getAllApplyList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/getAllApplyList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 提交差旅申请
export function createTravelApply(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createTravelApply',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 查看出行人员票据信息
export function getPeopleOrderList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/getPeopleOrderList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取前置申请单列表
export function getApplyList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/apply/getApplyList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
