<template>
  <n-card title="关于" :bordered="false" style="height: 100%;">
    GTools是一个集多种功能于一体的跨平台小工具。只为让你在开发过程中，心情愉悦！<br>
    由于本人不是专业的前端开发人员，所以目前的页面略显粗糙，但是我会通过不断地前端知识学习来改善UI和交互。<br>
    该项目完全开源，欢迎各位大佬star或者提交issue进行问题反馈。<br>
    <br>
    Gitee项目地址：<a href="#" @click="openUrl('https://gitee.com/pixiao/gtools')">https://gitee.com/pixiao/gtools</a><br><br>
    QQ交流群：736733848
  </n-card>

</template>
<script setup>
import { ref, watch } from "vue";
// import { EventsOn } from "../../wailsjs/runtime/runtime"
import { BrowserOpenURL  } from "/wailsjs/runtime/runtime"
const app = ref(window.go.gtools.App)
const rt = ref(runtime.runtime)
const result = ref(0)

function openUrl(url) {
  BrowserOpenURL(url)
}
</script>