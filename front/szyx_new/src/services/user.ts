import request from "./request";
//获取用户信息
export function getUserInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/user/getUserInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//修改用户信息
export function updateUserInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/user/updateUserInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 校验更新手机号验证码
export function phoneVerification(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/user/phoneVerification',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 更新新的手机号
export function modifyPhoneNum(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/user/modifyPhoneNum',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 修改用户的token信息
export function modifyUserInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/user/modifyUserInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
// 获取用户个人图表信息 根据时间范围统计个人出行次数、累计报销金额、累计申请次数
export function countAmountNumber(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/statement/countAmountNumber',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
