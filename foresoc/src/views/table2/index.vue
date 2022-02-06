<template>
<div class="tab2-container">
  <el-table
    :data="tableData"
    stripe
    style="width: 100%">
    <el-table-column
      prop="id"
      label="编号"
      width="120">
    </el-table-column>
    <el-table-column
      prop="name"
      label="姓名"
      width="180">
    </el-table-column>
    <el-table-column
      prop="role"
      label="角色"
      width="180">
    </el-table-column>
    <el-table-column
      prop="section"
      label="部门">
    </el-table-column>
  </el-table>
  </div>
</template>

<script>
import { findAll } from '@/api/table2'

export default {
  data() {
    name: name
    id:id
  },
  methods: {
    fetchData() {
      this.findAll = true
      findAll().then(response => {
        this.name = response.data.name
        this.id = response.data.id
      })
    }
  }
}
</script>

<style scoped>
  .tab2-container {
    margin: 20px;
    border:1px solid #c8c8c8
  }
</style>