import { defineStore } from "pinia";
import { useWindowSize } from '@vueuse/core';
const windowSize = useWindowSize();

let useSettingStore = defineStore("Setting", {
  state: () => {
    return {
      fold: false,
      clientWidth: windowSize.width,
      clientHeight: windowSize.height,
    };
  },
  //异步,逻辑区
  actions: {
    // changeIcon(){
    //   this.fold = !this.fold
    // }
  },
});
export default useSettingStore;
