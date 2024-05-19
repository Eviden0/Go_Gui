<template>
    <n-space vertical>
        <n-card :bordered="false">
            <n-grid x-gap="12" :cols="2">
                <n-gi>
                    <n-input placeholder="比对文本" type="textarea" size="small" v-model:value="oldStr" :autosize="{
                        minRows: 3,
                        maxRows: 20,
                    }" />
                </n-gi>
                <n-gi>
                    <n-input placeholder="比对文本" type="textarea" size="small" v-model:value="newStr" :autosize="{
                        minRows: 3,
                        maxRows: 20,
                    }" />
                </n-gi>
            </n-grid>
        </n-card>
        <n-card :bordered="false">
            <n-form ref="formRef" inline :label-width="90" size="medium" label-placement="left">
                <n-form-item label="差异化范围">
                    <n-input-number v-model:value="context" style="width: 120px;"/>
                </n-form-item>
                <n-form-item label="差异级别">
                    <n-radio-group v-model:value="diffStyle" name="diffStyle" :on-update="updateDiffStyle(diffStyle)">
                        <n-radio-button v-for="dtype in diffStyleTypes" :key="dtype.value" :value="dtype.value"
                            :label="dtype.label" />
                    </n-radio-group>
                </n-form-item>
                <n-form-item label="移除字符串前后空格" :label-width="150">
                    <n-switch v-model:value="trim" />
                </n-form-item>
                <n-form-item  :label-width="150">
                    <n-button size="medium" secondary strong type="warning" round @click="clearDiffStr()">
                        清空比对文本
                    </n-button>
                </n-form-item>
            </n-form>
        </n-card>
        <n-card :bordered="false">
            <code-diff :old-string="oldStr" :new-string="newStr" file-name="" :diffStyle="diffStyle"
                output-format="side-by-side" :renderNothingWhenEmpty="true" :isShowNoChange="true" :context="context"
                :trim="trim" />
        </n-card>
    </n-space>

</template>
<script setup>
import { CodeDiff } from 'v-code-diff'
import { ref, watch } from 'vue'


const context = ref(5)
const oldStr = ref("")
const newStr = ref("")
const trim = ref(false)
const diffStyle = ref("word")
const diffStyleTypes = ref([{
    label: "词",
    value: "word"
},
{
    label: "字",
    value: "char"
}])

watch(context, (newVal, oldVal) => {
    if (newVal < 0) {
        context.value = 0
    }
})

function updateDiffStyle(param) {
    diffStyle.value = param
}

function clearDiffStr() {
    oldStr.value = ""
    newStr.value = ""
}
</script>
<style>

</style>