import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite';
import {VantResolver} from 'unplugin-vue-components/resolvers';

export default defineConfig({
    base: '/front',
    plugins: [
        vue(),
        Components({
            resolvers: [VantResolver()],
        }),
    ],
    // 配置反向代理
    server: {
        host: '0.0.0.0',
        port: 8080, // 端口号
        cors: true, //为开发服务器配置 CORS , 默认启用并允许任何源
        hmr: true, // 开启热更新
        proxy: {
            '/v4': {
                // 配置接口调用目标地址
                //FIXME 接口地址
                target: 'http://localhost:10012',
                // 当进行代理时，Host 的源默认会保留（即Host是浏览器发过来的host），如果将changeOrigin设置为true，则host会变成target的值。
                changeOrigin: true,
                // 前缀 /api 是否被替换为特定目标，不过大多数后端给到的接口都是以/api打头，这个没有完全固定的答案，根据自己的业务需求进行调整
                // rewrite: path => path.replace(/^\/api/, ''),
            }
        }
    },

})
