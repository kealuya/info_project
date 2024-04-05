import request from "./request";
import requestApproval from "./request";
// 获取手机号验证码
export function getMessageCaptcha(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/user/getMessageCaptcha',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 登录接口
export function messageLogin(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/user/messageLogin',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

// //获取待审批数量
// 首页单据列表
export function getDocStateList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/accountBook/getApplicationFormList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
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
// 获取账本个数
export function getAccountBookList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/accountBook/getAccountBookList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
