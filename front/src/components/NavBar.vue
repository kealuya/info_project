//头部van-nav-bar样式组件
<template>
    <div class="barCss">
        <van-nav-bar
                :title="title"
                :left-text="leftText"
                :left-arrow="leftArrow"
                :right-text="rightText"
                @click-left="leftEvent"
                @click-right="rightEvent"
                :border="false"
                :style="{width:realWidth+'px'}"
        >
            <template v-if="slortName != undefined && slortName != ''" #[slortName]>
                <van-icon v-if="iconName != undefined && iconName != ''" :name="iconName" size="18" @click="iconMth"/>
                <div class="f16 fc ml5"> {{iconText}}</div>
            </template>
        </van-nav-bar>
    </div>
</template>

<script setup lang="ts">
    import {defineProps, defineEmits, ref} from 'vue'

    interface Props {
        title?: string,
        leftText?: string
        rightText?: string
        leftArrow?: boolean
        iconName?: string
        slortName?: string//插槽位置：'right','title','left'
        iconText?: string

    }

    defineProps<Props>()
    // 此处声明方法写入此处
    const emit = defineEmits(["leftEvent","rightEvent",'iconEvent']);

    //获取当前宽度
    const realWidth = window.document.documentElement.offsetWidth

    const leftEvent = (() => {
        emit("leftEvent")
    })
    const rightEvent = (() => {
      emit("rightEvent" )
    })

    function iconMth() {
        emit("iconEvent" as never)
    }
</script>

<style scoped>
    .barCss .van-nav-bar {
        background: #0080FF;
    }

    .barCss {
        --van-nav-bar-icon-color: #fafafa;
        --van-nav-bar-text-color: #fafafa;
        --van-nav-bar-title-text-color: #fafafa;
    }

    .fc {
        color: white;
    }

    .barCss .van-nav-bar::after {
        border: none;
    }


</style>
