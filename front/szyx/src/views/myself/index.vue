<script setup lang="ts">

</script>

<template>
  <div class="app-container">
    <div class="my-bg">
      <img class="img-avatar" src="@/assets/yaopintouxiang.png" />
      <img class="img-bg" src="@/assets/yaopin.jpg" />
    </div>
    <div class="my-body">
      <div class="my-name">演示一下</div>
      <div class="my-job">高级会员</div>

      <van-grid direction="horizontal" :column-num="3">
        <van-grid-item icon="user-o" text="开直播" />
        <van-grid-item icon="friends-o" text="开会议" />
        <van-grid-item icon="video-o" text="发视频" />
        <van-grid-item icon="comment-o" text="发文章" />
        <van-grid-item icon="music-o" text="发广播" />
        <van-grid-item icon="photo-o" text="更多.." />
      </van-grid>


      <div class="my-desc">
        1、费⽤规范使⽤的全流程、全场景配套服务管理⽅案。
        2、营销费⽤转学术研发的配套⽅案和执⾏服务。
        3、基于新型学术关系管理，数字化学术沟通推⼲、低成本⾼频触达、达成⼲泛学术共识和观念影响带来的增量业务服务。
        4、创新⽀付解决⽅案。
      </div>
    </div>
    <div class="my-footer">
      1、费⽤规范使⽤的全流程、全场景配套服务管理⽅案。
      2、营销费⽤转学术研发的配套⽅案和执⾏服务。
      3、基于新型学术关系管理，数字化学术沟通推⼲、低成本⾼频触达、达成⼲泛学术共识和观念影响带来的增量业务服务。
      4、创新⽀付解决⽅案。
    </div>
  </div>
</template>





<style lang="scss" scoped>
.app-container {
  padding-bottom: 50px;
  .my-bg {
    height: 150px;
    border-radius: 4px;
    position: relative;
    .img-avatar {
      width: 100px;
      height: 100px;
      position: absolute;
      bottom: -50px;
      left: calc(50vw - 50px);
      border-radius: 50%;
      z-index: 99;
    }
    .img-bg {
      position: absolute;
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
  }
  .my-body {
    margin-top: 74px;
    display: flex;
    flex-direction: column;
    align-items: center;
    .my-name {
      font-size: 16px;
      font-weight: bold;
      margin-bottom: 10px;
    }
    .my-job {
      font-size: 14px;
      color: $mainColor;
    }
    .my-desc {
      font-size: 14px;
      color: $baseFontColor;
      padding: 16px 24px;
      line-height: 24px;
    }
  }
  .my-footer {
    background: $baseBackColor;
    border-radius: 12px;
    margin: 0 16px;
    padding: 12px;
    text-align: left;
    line-height: 24px;
    color: $subFontColor;
    font-size: 14px;
    a {
      color: $baseLinkColor;
      text-decoration: underline;
    }
  }
}
</style>
