import SvgIcon from "./SvgIcon/SvgIcon.vue";
import type { App, Component } from "vue";
// 引入icon
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
// console.log('ElementPlusIconsVue',ElementPlusIconsVue)
const components: { [name: string]: Component } = { SvgIcon };
export default {
  install(app: App) {
    Object.keys(components).forEach((key: string) => {
      app.component(key, components[key]);
    });

    for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
      app.component(key, component)
    }

  },
};
