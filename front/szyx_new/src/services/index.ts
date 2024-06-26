import request from "./request";
//获取验证码接口
export function getMessageCaptcha(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v4/user/getMessageCaptcha',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//登录接口
export function messageLogin(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v4/user/messageLogin',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//首页获取任务池列表的接口
export function getTaskPoolList(query :object){
    // console.log('query',query)
    return new Promise(function (resolve, reject) {
        request({
            url:'/v4/task/getTaskPoolList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//首页任务详情=》去参与的接口
export function applyJoinTask(query :object){
    // console.log('query',query)
    return new Promise(function (resolve, reject) {
        request({
            url:'/v4/task/applyJoinTask',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}
//getTaskList 我的任务列表
export function getTaskList(query :object){
    // console.log('query',query)
    return new Promise(function (resolve, reject) {
        request({
            url:'/v4/task/getTaskList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}