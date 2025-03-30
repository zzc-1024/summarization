<script setup lang="ts">
import { onMounted, ref } from "vue";
import { test, uploadFile, getTask, abstractTask } from "../api/task";
import { useRoute } from "vue-router";

const resource = ref([]);

const fileName = ref("");
const batchFile = ref("");
const answer = ref("")
const MAX_FILE_SIZE = 10 * 1000 * 1000;

const route = useRoute();

const init = () => {
  getTask(route.params.id as string).then((res) => {
    if (res.filenames === null) {
      return;
    }
    resource.value = res.filenames.map((item) => {
      return {
        name: item,
      };
    });
    console.log(resource.value);
  });
};

onMounted(() => {
  init();
});

function handleDelete() {
  console.log("删除");
}

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
  const file = e.dataTransfer.files[0];
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
async function uploadOk() {
  const res = await test();
  console.log(res);
  if (batchFile.value === "") {
    return alert("请选择要上传的文件");
  }
  let data = new FormData();
  data.append("file", batchFile.value);
  data.append("taskId", route.params.id as string);
  const fileRes = await uploadFile(data);
  console.log(fileRes);
  // 清空待上传文件
  batchFile.value = "";
  fileName.value = "";
  init();
}
async function onAbstractTask() {
  const res = await abstractTask(route.params.id as string);
  console.log(res);
  answer.value = res.answer.message.content;
}
</script>

<template>
  <div id="app">
    <el-table :data="resource" style="width: 100%">
      <el-table-column prop="name" label="文件名" />
      <el-table-column prop="operator" label="操作">
        <template #default>
          <el-button link type="primary" size="small" @click="handleDelete">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="content">
      <div class="drag-area" @dragover="fileDragover" @drop="fileDrop">
        <div v-if="fileName" class="file-name">{{ fileName }}</div>
        <div v-else class="uploader-tips">
          <span>将文件拖拽至此，或</span>
          <label for="fileInput" style="color: #11a8ff; cursor: pointer">
            点此上传
          </label>
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
      >
        选择文件
      </label>

      <button style="background-color: antiquewhite" @click="uploadOk">
        提交
      </button>
      <button
        style="background-color: antiquewhite; margin: 10px"
        @click="onAbstractTask"
      >
        生成摘要
      </button>
      <button
        style="background-color: firebrick; color: aliceblue"
        @click="$router.go(-1)"
      >
        返回上一页
      </button>
    </div>
    {{ answer }}
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
