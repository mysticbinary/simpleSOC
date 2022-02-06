<template>
  <div class="tab2-container">
    <el-table :data="tableData" v-loading="loading" stripe style="width: 100%">
      <el-table-column prop="id" label="编号" width="120"
        ><template slot-scope="{ row }">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="name" label="姓名" width="180">
        <template slot-scope="{ row }">
          <span>{{ row.name }}</span>
        </template>
      </el-table-column>
      <!-- <el-table-column prop="role" label="角色" width="180"> </el-table-column>
      <el-table-column prop="section" label="部门"> </el-table-column> -->
    </el-table>
  </div>
</template>

<script>
import { findAll } from "@/api/table2";

export default {
  data() {
    return {
      id,
      name,
    };
  },
  created() {
    this.findAllfunc();
  },
  methods: {
    findAllfunc() {
      console.log("table created 3");
      findAll().then((response) => {
        console.log("table created 4");
        this.id = response.id;
        // Just to simulate the time of the request
        // setTimeout(() => {
        //   this.listLoading = false;
        // }, 1.5 * 1000);
      });
    },
  },
};
</script>

<style scoped>
.tab2-container {
  margin: 20px;
  border: 1px solid #c8c8c8;
}
</style>