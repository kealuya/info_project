import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite';
import {VantResolver} from 'unplugin-vue-components/resolvers';

export default defineConfig({
    base: '/costH5_XS',
    plugins: [
        vue(),
        Components({
            resolvers: [VantResolver()],
        }),
    ],
    // 配置反向代理
    server: {
        host: '0.0.0.0',
        port: 8088, // 端口号
        cors: true, //为开发服务器配置 CORS , 默认启用并允许任何源
        hmr: true, // 开启热更新
        proxy: {
            // 'http://124.70.111.57:10012' // 服务器
            // 'http://192.168.120.220:10012' //内网本地
            // http://221.239.127.74:12000/lpf_fp' //穿透本地
            '/ApprovalEcl/v1': {
                // 调用审批系统
                // 配置接口调用目标地址
                //  target: 'http://192.168.120.221:7000', //
                target: 'http://122.9.41.215:8089', // 审批
                // 当进行代理时，Host 的源默认会保留（即Host是浏览器发过来的host），如果将changeOrigin设置为true，则host会变成target的值。
                changeOrigin: true,
                // 前缀 /api 是否被替换为特定目标，不过大多数后端给到的接口都是以/api打头，这个没有完全固定的答案，根据自己的业务需求进行调整
                // rewrite: path => path.replace(/^\/api/, ''),
            },
            '/v4': {
                // 调用审批系统
                // 配置接口调用目标地址
                //  target: 'http://192.168.120.221:7000', //
                target: 'http://192.168.120.220:7001', // 审批
                // 当进行代理时，Host 的源默认会保留（即Host是浏览器发过来的host），如果将changeOrigin设置为true，则host会变成target的值。
                changeOrigin: true,
                // 前缀 /api 是否被替换为特定目标，不过大多数后端给到的接口都是以/api打头，这个没有完全固定的答案，根据自己的业务需求进行调整
                // rewrite: path => path.replace(/^\/api/, ''),
            },
            '/v1': {
                // 配置接口调用目标地址
                target: 'http://124.70.111.57:10003', // 服务器
                // target: 'http://192.168.120.221:10003', // 艳茹本地
                // 当进行代理时，Host 的源默认会保留（即Host是浏览器发过来的host），如果将changeOrigin设置为true，则host会变成target的值。
                changeOrigin: true,
                // 前缀 /api 是否被替换为特定目标，不过大多数后端给到的接口都是以/api打头，这个没有完全固定的答案，根据自己的业务需求进行调整
                // rewrite: path => path.replace(/^\/api/, ''),
            },
            '/v2': {
                // 配置接口调用目标地址
                target: 'http://124.70.111.57:7000', // 服务器
                //target: 'http://172.16.160.43:6686', // 本地启动power
                // target: 'http://192.168.120.221:7000', // 艳茹本地
                // 当进行代理时，Host 的源默认会保留（即Host是浏览器发过来的host），如果将changeOrigin设置为true，则host会变成target的值。
                changeOrigin: true,
                // 前缀 /api 是否被替换为特定目标，不过大多数后端给到的接口都是以/api打头，这个没有完全固定的答案，根据自己的业务需求进行调整
                // rewrite: path => path.replace(/^\/api/, ''),
            },
        }
    },

})
