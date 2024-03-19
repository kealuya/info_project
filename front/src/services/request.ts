//引入axios
import axios from 'axios'
import {Toast} from 'vant';
const version='/v4';
//使用指定配置创建axios实例
const instance = axios.create({
    // 基础路径
    baseURL: version+'',
    // 请求超时时间
    timeout: 60000,
    method:'post',

})

/**
 * 请求拦截
 */
instance.interceptors.request.use(
    function (config) {
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
instance.interceptors.response.use(
    function (response) {
        // console.log(response);
        let res = typeof (response.data) === 'string' ? JSON.parse(response.data) : response.data
        if (res.success) {
            // 对响应数据做点什么
            return res.data;
        } else {
            Toast.fail('网络错误,请联系管理员');
        }
    },
    function (error) {
        // 关闭 loading
        if (error.config.loading) {
            Toast.clear();
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
        Toast.fail(error.message);

        return Promise.reject(error);
    },
);

export default instance;
export const request = instance.request;

