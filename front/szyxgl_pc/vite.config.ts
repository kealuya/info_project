import { defineConfig } from "vite";
import { resolve } from "path"; //编辑器提示 path 模块找不到，可以cnpm i @types/node --dev 即可
import vue from "@vitejs/plugin-vue";
import * as path from "path";
import { viteMockServe } from "vite-plugin-mock";
//引入svg插件
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";
// import { viteMockServe } from "vite-plugin-mock";
// 当前执行node命令时文件夹的地址（工作目录）
const pathSrc = path.resolve(__dirname, "src"); // 绝对路径src

//自动导入ui-组件 比如说ant-design-vue  element-plus等
import Components from "unplugin-vue-components/vite"; // 封装组件就不用引入了，自动导入
const root: string = process.cwd(); //获取工作目录
const pathRoot = path.resolve(__dirname, "");
// https://vitejs.dev/config/

export default defineConfig({
  base: process.env.NODE_ENV == "production" ? "./" : "/",
  resolve: {
    alias: {
      "@": resolve(process.cwd(), "./src"),
      "#": resolve(process.cwd(), "./types"),
      "~/": `${pathSrc}/`,
      "vue-i18n": "vue-i18n/dist/vue-i18n.runtime.esm-bundler.js",
    },
  },
  css: {
    //全局引入样式
    preprocessorOptions: {
      scss: {
        javascriptEnabled: true,
        additionalData: '@use "./src/styles/variable.scss" as *;',
        // additionalData: `@use "~/styles/index.scss" as *;`, //引入多个
      },
    },
  },
  plugins: [
    vue({
      include: [/\.vue$/, /\.md$/],
      template: {
        compilerOptions: {
          isCustomElement: (tag) => tag.startsWith("ion-"),
        },

        // types: ["jquery"]
      },
    }),
    //引入svg图片
    createSvgIconsPlugin({
      // Specify the icon folder to be cached
      iconDirs: [path.resolve(process.cwd(), "src/assets/icons")],
      // Specify symbolId format
      symbolId: "icon-[dir]-[name]",
    }),
    viteMockServe({
      supportTs: false,
      logger: false,
      mockPath: "./src/mock/",
    }),

    // viteMockServe({
    //   mockPath: "./src/mock/index", // 解析，路径可根据实际变动
    //   localEnabled: true, // 此处可以手动设置为true，也可以根据官方文档格式
    //   supportTs: true,
    //   logger: true
    // }),

    Components({
      // 引入组件的,包括自定义组件
      // 存放的位置
      dts: "src/components.d.ts",
    }),
  ],
  server: {
    host: '0.0.0.0', //ip地址
    port: 5173, //端口号
    open: true //启动后是否自动打开浏览器
  }

});
