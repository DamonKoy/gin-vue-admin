
<template>
  <div>
    <el-table :data="tableData" border style="width: 100%">
      <!-- <el-table-column prop="id" label="ID" width: 5% />
    <el-table-column prop="brand" label="品牌" />
    <el-table-column prop="model" label="型号" />
    <el-table-column prop="platform" label="平台" />
    <el-table-column prop="system_version" label="系统版本" />
    <el-table-column prop="auditor" label="管理人" width: 10% />
    <el-table-column prop="holder" label="持有人" width: 10% />
    <el-table-column prop="status" label="状态" />
    <el-table-column prop="update_time" label="更新时间" width: 15% />
    <el-table-column prop="remarks" label="备注" width: 10% /> -->

      <el-table-column align="left" prop="id" label="ID" />
      <el-table-column align="left" prop="brand" label="品牌" />
      <el-table-column align="left" prop="model" label="型号" />
      <el-table-column align="left" prop="platform" label="平台" />
      <el-table-column align="left" prop="system_version" label="系统版本" />
      <el-table-column align="left" prop="auditor" label="管理人" />
      <el-table-column align="left" prop="holder" label="持有人" />
      <el-table-column align="left" prop="status" label="状态" />
      <el-table-column align="left" prop="update_time" label="更新时间" />
      <el-table-column align="left" prop="remarks" label="备注" />

      <el-table-column align="right">
        <template #default="scope">
          <el-button
            size="small"
            @click="updateDeviceHolder(scope.$index, scope.row)"
          >借出</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :total="total"
      :current-page="page"
      :page-size="[10, 30, 50, 100]"
      layout="total,sizes, prev, pager, next, jumper"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />

  </div>
</template>

<script setup>
import { getDeviceList, updateDeviceHolder } from '@/api/device'
import { ref } from 'vue'

const page = ref(1)
const total = ref(0)
const pageSize = ref(50)
const tableData = ref([])
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getDeviceList({
    page: page.value,
    pageSize: pageSize.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    console.log(tableData.value)
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

</script>
