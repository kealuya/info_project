import request from "./request";
// 获取物流地址
export function getWuliuLoginUrl(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/trirdPartyBusiness/getWuliuLoginUrl',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}


// 获取用车地址
export function getCarLoginUrl(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/trirdPartyBusiness/getCarLoginUrl',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

// 登录行旅
export function loginToAuvgo(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/apply/loginToAuvgo',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

// 获取token
export function getToken(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/token/getToken',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}

// 获取心意选路径
export function getLoginHtShopUrl(query :object){
    return new Promise(function (resolve, reject) {
        request({
            url:'/v2/trirdPartyBusiness/getLoginHtShopUrl',
            data: query
        }).then((response: object) => {
            resolve(response)
        }).catch((error :string) => {
            resolve({})
        })
    })
}


