<script setup lang="ts">
import { ref } from "vue";
const fileName = ref("");
const batchFile = ref("");
const MAX_FILE_SIZE = 10 * 1000 * 1000;
const chooseUploadFile = (e) => {
  const file = e.target.files.item(0);
  if (!file) return;
  if (file.size > MAX_FILE_SIZE) {
    return alert("文件大小不能超过10M");
  }
  batchFile.value = file;
  fileName.value = file.name;
  // 清空，防止上传后再上传没有反应
  e.target.value = "";
};
// 拖拽上传
function fileDragover(e) {
  e.preventDefault();
}
function fileDrop(e) {
  e.preventDefault();
  const file = e.dataTransfer.files[0]; // 获取到第一个上传的文件对象
  console.log(file);
  console.log("拖拽释放鼠标时");

  if (!file) return;
  if (file.size > MAX_FILE_SIZE) {
    return alert("文件大小不能超过10M");
  }
  batchFile.value = file;
  fileName.value = file.name;
}
// 提交
function uploadOk() {
  if (batchFile.value === "") {
    return alert("请选择要上传的文件");
  }
  let data = new FormData();
  data.append("upfile", batchFile.value);
  // ajax
}
</script>

<template>
  <div id="app">
    <div class="content">
      <div class="drag-area" @dragover="fileDragover" @drop="fileDrop">
        <div v-if="fileName" class="file-name">{{ fileName }}</div>
        <div v-else class="uploader-tips">
          <span>将文件拖拽至此，或</span>
          <label for="fileInput" style="color: #11a8ff; cursor: pointer"
            >点此上传</label
          >
        </div>
      </div>
    </div>

    <div class="footer">
      <input
        type="file"
        id="fileInput"
        @change="chooseUploadFile"
        style="display: none"
      />
      <label
        for="fileInput"
        v-if="fileName"
        style="color: #11a8ff; cursor: pointer"
        >选择文件</label
      >

      <button style="background-color: antiquewhite" @click="uploadOk">
        提交
      </button>
    </div>
  </div>
</template>

<style scoped>
* {
  font-size: 14px;
}
.drag-area {
  height: 200px;
  width: 300px;
  border: dashed 1px gray;
  margin-bottom: 10px;
  color: #777;
}
.uploader-tips {
  text-align: center;
  height: 200px;
  line-height: 200px;
}
.file-name {
  text-align: center;
  height: 200px;
  line-height: 200px;
}
</style>
