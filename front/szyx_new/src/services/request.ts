//引入axios
import axios from 'axios'
import {showToast, closeToast} from 'vant';
import {userInfoData} from "../store";
import router from "../router";
const version='/v2';
//使用指定配置创建axios实例
const instance = axios.create({
    // 基础路径
    // baseURL: version+'',
    // 请求超时时间
    timeout: 10000,
    method:'post',

})

/**
 * 请求拦截
 */
instance.interceptors.request.use(
    function (config) {
        if (localStorage.getItem("token")) {
            // @ts-ignore
            config.headers = {
                'token': localStorage.getItem("token"),
                'refreshToken': localStorage.getItem("refreshToken")
            }
            const userInfoState: any = userInfoData();
            if (Object.keys(userInfoState.userInfo).length !== 0) {
                // config.data['corpCode'] = 'MX';

                // config.data['corpCode'] = userInfoState.userInfo.corpCode || '';
                // config.data['empCode'] = userInfoState.userInfo.empCode || '';
            }
        }
        return config;
    },
    function (error) {
        // 对请求错误做些什么
        return Promise.reject(error);
    },
);

/**
 * 响应拦截
 */
let refresh = false;
instance.interceptors.response.use(
    function (response) {
        let res = typeof (response.data) === 'string' ? JSON.parse(response.data) : response.data
        if (res.code && res.code === 400) {
            // 重新获取token
            if (!refresh) {
                refresh = true;
                return refreshToken(response.config);
            } else {
                // 正在刷新token，返回一个未执行resolve的promise
                return new Promise((resolve) => {
                    // 将resolve放进队列，用一个函数形式来保存，等token刷新后直接执行
                    requests.push(() => {
                        resolve(
                            instance({
                                url: response.config.url,
                                method: 'post',
                                data: JSON.parse(response.config.data),
                            })
                        )
                    })
                });
            }
        } else if (res.code && res.code === 509) {
            closeToast();
            router.push("/login");
        }
        return res;
    },
    function (error) {
        // 关闭 loading
        if (error.config.loading) {
            closeToast();
        }
        // 对响应错误做点什么
        // switch (error.response?.status) {
        //     case 400:
        //         error.message = '请求错误(400)';
        //         break;
        //     case 401:
        //         error.message = '未授权(401)';
        //         break;
        //     case 403:
        //         error.message = '拒绝访问(403)';
        //         break;
        //     case 404:
        //         error.message = '请求出错(404)';
        //         break;
        //     case 408:
        //         error.message = '请求超时(408)';
        //         break;
        //     case 500:
        //         error.message = '服务器错误(500)';
        //         break;
        //     case 501:
        //         error.message = '服务未实现(501)';
        //         break;
        //     case 502:
        //         error.message = '网络错误(502)';
        //         break;
        //     case 503:
        //         error.message = '服务不可用(503)';
        //         break;
        //     case 504:
        //         error.message = '网络超时(504)';
        //         break;
        //     case 505:
        //         error.message = 'HTTP版本不受支持(505)';
        //         break;
        //     default:
        //         error.message = `连接出错(${error.response?.status})!`;
        // }
        showToast(error.message);

        return Promise.reject(error);
    },
);
let requests = <any>[];
const refreshToken = (config: any) => {
    const service1 = axios.create({
        timeout: 5000
    });
    service1.interceptors.request.use(
        config => {
            const userInfoState: any = userInfoData();
            if (Object.keys(userInfoState.userInfo).length !== 0) {
                config.data['corpCode'] = userInfoState.userInfo.corpCode || '';
            }
            config.data['token'] = localStorage.getItem("refreshToken") || '';
            config.data['refreshToken'] = localStorage.getItem("refreshToken") || '';
            config.data['equipmentType'] = 'PC';
            return config;
        },
        error => {
            return Promise.reject();
        }
    );
    service1.interceptors.response.use(
        // @ts-ignore
        function (response) {
            refresh = false;
            let res = typeof (response.data) === 'string' ? JSON.parse(response.data) : response.data;
            if (res.code === 500) {
                window.localStorage.clear();
                router.push("/login");
                return showToast(res.msg.toString());
            } else {
                let result: any= setToken(res);
                if (result) {
                    requests.forEach((cb: any) => cb())
                    requests = [];
                    return instance({
                        url: config.url,
                        method: 'post',
                        data: JSON.parse(config.data),
                    });
                }
            }
        },
        function (error) {
            // 关闭 loading
            if (error.config.loading) {
                closeToast();
            }
            // 对响应错误做点什么
            // switch (error.response?.status) {
            //     case 400:
            //         error.message = '请求错误(400)';
            //         break;
            //     case 401:
            //         error.message = '未授权(401)';
            //         break;
            //     case 403:
            //         error.message = '拒绝访问(403)';
            //         break;
            //     case 404:
            //         error.message = '请求出错(404)';
            //         break;
            //     case 408:
            //         error.message = '请求超时(408)';
            //         break;
            //     case 500:
            //         error.message = '服务器错误(500)';
            //         break;
            //     case 501:
            //         error.message = '服务未实现(501)';
            //         break;
            //     case 502:
            //         error.message = '网络错误(502)';
            //         break;
            //     case 503:
            //         error.message = '服务不可用(503)';
            //         break;
            //     case 504:
            //         error.message = '网络超时(504)';
            //         break;
            //     case 505:
            //         error.message = 'HTTP版本不受支持(505)';
            //         break;
            //     default:
            //         error.message = `连接出错(${error.response?.status})!`;
            // }
            showToast(error.message);

            return Promise.reject(error);
        }
        );
    const userInfoState: any = userInfoData();
    return service1({
        url: '/v2/token/refreshToken',
        method: 'post',
        data: {
            'empCode': userInfoState.userInfo.empCode || ''
        },
    });
}
const setToken = async (res: any) => {
    await localStorage.setItem("token", res.jwt);
    await localStorage.setItem("refreshToken", res.refreshToken);
    return true;
}

export default instance;
export const request = instance.request;

