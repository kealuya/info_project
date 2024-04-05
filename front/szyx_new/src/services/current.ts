// 此处用于可复用的选择器 （包含：选择成本中心的列表，选择审批流， 选择人员一类）
import request from "./request";
// 成本中心列表
export function getCostAndBudgetDeptList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/costCenter/getCostAndBudgetDeptList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 根据成本中心获取对应审批列表进行选择
export function getApprovalFlow(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v1/approval/getApprovalFlow',
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
// 查看出行人员是否携带身份证
export function selectPerIDCard(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/user/selectPerIDCard',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
