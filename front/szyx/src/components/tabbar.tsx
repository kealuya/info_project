import { defineComponent } from "vue";
import "./tabbar.scss";

export default defineComponent({
  name: "Tabbar",
  setup() {
    return () => (
      <>
        <router-view />
        <van-tabbar active-color="#009688" inactive-color="#595959" route>
          <van-tabbar-item replace to="/home" icon="home-o">
            首页
          </van-tabbar-item>
          <van-tabbar-item replace to="/label" icon="label-o">
            广场
          </van-tabbar-item>
          <van-tabbar-item replace to="/message" icon="notes-o">
            社区
          </van-tabbar-item>
          <van-tabbar-item replace to="/myself" icon="user-o">
            我的
          </van-tabbar-item>
        </van-tabbar>
      </>
    );
  },
});
