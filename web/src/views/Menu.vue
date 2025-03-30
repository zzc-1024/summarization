<script setup lang="ts">
import { onMounted, ref } from "vue";
import router from "../router";
import { getTasks } from "../api/task";

const tasks = ref([]);

function init() {
  getTasks().then((res) => {
    tasks.value = res.folderNames.map((item) => {
      return {
        id: item,
      };
    });
    console.log(tasks.value);
  });
}

onMounted(() => {
  init();
});

function jumpTo(id: string) {
  // 路由跳转
  router.push({ path: `/task/${id}` });
}
</script>

<template>
  <div>
    <el-table :data="tasks">
      <el-table-column prop="id" label="任务标识" />
      <el-table-column label="操作">
        <template #default="scope">
          <el-button link type="primary" size="small" @click="jumpTo(scope.row.id)">
            查看
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
