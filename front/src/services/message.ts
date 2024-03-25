import request from "./request";

//获取信息列表
export function getMessageList(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/common/getNewList',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

//获取推送信息
export function getMessagePull(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/common/getMsgPull',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

//修改信息状态
export function updateMsgPullState(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/common/updateMsgPullState',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}




