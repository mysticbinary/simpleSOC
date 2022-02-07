
<template>
  <div class="tab2-container">
    <el-table
      :data="tab2"
      stripe
      style="width: 100%"
    >
      <el-table-column
        prop="id"
        label="ID"
        width="180"
      />
      <el-table-column
        prop="name"
        label="姓名"
        width="180"
      />
      <el-table-column
        prop="role"
        label="角色"
        width="180"
      />
      <el-table-column
        prop="section"
        label="部门"
      />
    </el-table>
  </div>
</template>

<script>
import { findAll } from '@/api/table2'

export default {
  data() {
    return {
      tab2: null
    }
  },
  created() {
    this.findAllfunc()
  },
  methods: {
    findAllfunc() {
      findAll().then((response) => {
        console.log(response)
        this.tab2 = response.data
      })
    }
  }
}
</script>

<style scoped>
.tab2-container {
  margin: 20px;
  border: 1px solid #c8c8c8;
}
</style>
