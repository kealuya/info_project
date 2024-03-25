import request from "./request";

//获取委托设定列表
export function getMyAgentPerson(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/getMyAgentPerson',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

//获取委托设定列表详情
export function getMyAgentInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/getMyAgentInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

//委托启用停用
export function updateAgentState(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/updateAgentState',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}


//新建委托
export function createAgentInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/createAgentInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

//删除委托
export function deleteAgentInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/deleteAgentInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

//更改委托
export function updateAgentInfo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/updateAgentInfo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}



