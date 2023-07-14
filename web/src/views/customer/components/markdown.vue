<template>
  <div>
    <el-upload></el-upload>
    <el-button type="text" @click="changeMdMode()">{{ mdModeText }}</el-button>
    <v-md-editor
      :model-value="md.text"
      :style="{ height: computedHeight }"
      :mode="md.mode"
      @update:model-value="md.text = $event"
      @upload-image="handleUploadImage"
      :disabled-menus="[]"
      @save="save"
    ></v-md-editor>
  </div>
</template>

<script lang="ts" setup>
import {reactive, watchEffect, ref, computed} from "vue"
import {uploadFile} from "@/api/system/base";

const props = defineProps({
  row: {
    type: Object,
    default: 0
  },
  text: {
    type: String,
    default: "部署报告为空"
  },
  mode: {
    type: String,
    default: "preview" // 可选值: edit(纯编辑模式) | editable(编辑与预览模式) | preview(纯预览模式)
  }
})

const emits = defineEmits(["saveDoc"])

const md = reactive({
  text: props.text,
  mode: props.mode
})

const changeMdMode = () => {
  md.mode = md.mode === "editable" ? "preview" : "editable"
}

const mdModeText = computed(() => {
  return md.mode === "preview" ? "编辑" : "取消编辑"
})

const save = () => {
  emits("saveDoc", md.text, props.row)
}

const computedHeight = ref("auto")

watchEffect(() => {
  const editor = document.querySelector(".v-md-editor")
  if (editor) {
    computedHeight.value = editor.clientHeight + "px"
  }
})

const handleUploadImage = async (event, insertImage, files) => {
  const formData = new FormData()
  formData.append("file", files[0])

  try {
    const res = await uploadFile(formData)
    const imageURL = res.data.imageUrl
    insertImage({
      url: imageURL,
      desc: "添加图片描述"
    })
  } catch (error) {
    console.log("文件上传失败:", error)
  }
}

</script>

<style lang="scss" scoped></style>
