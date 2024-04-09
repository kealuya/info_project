import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite';
import {VantResolver} from 'unplugin-vue-components/resolvers';

export default defineConfig({
    base: '/costH5',
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
            
        }
    },

})
