
<template>
  <div>
    <el-table :data="tableData" border style="width: 100%" @row-click="handleUpdate">
      <el-table-column align="left" prop="id" label="ID" />
      <el-table-column align="left" prop="brand" label="品牌" />
      <el-table-column align="left" prop="model" label="型号" />
      <el-table-column align="left" prop="platform" label="平台" />
      <el-table-column align="left" prop="system_version" label="系统版本" />
      <el-table-column align="left" prop="auditor" label="管理人" />
      <el-table-column align="left" prop="holder" label="持有人" />
      <el-table-column align="left" prop="status" label="状态" />
      <el-table-column align="left" prop="updated_at" label="更新时间" />
      <el-table-column align="left" prop="remarks" label="备注" />

      <!-- 更新按钮 -->
      <el-table-column fixed="right" label="操作" width="120">
        <template #default="scope">
          <el-button icon="edit" size="small" @click="openEdit(scope.row)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="gva-pagination">
      <el-pagination
        :total="total"
        :current-page="page"
        :page-size="[10, 30, 50, 100]"
        layout="total,sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>

    <!-- 更新操作弹窗 -->
    <el-dialog
      v-model="editDeviceDialog"
      title="更新设备持有人"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form ref="deviceForm" :model="deviceInfo" label-width="80px">
          <el-form-item label="借出人" prop="holder">
            <el-input v-model="deviceInfo.holder" autocomplete="off" readonly="true" />
          </el-form-item>
          <el-form-item label="借入人" prop="currentHolder">
            <el-input v-model="deviceInfo.currentHolder" placeholder="必填,请输入借入人名字" />
          </el-form-item>
          <el-form-item label="内容记录" prop="remarks">
            <el-input v-model="deviceInfo.describe" placeholder="非必填,可按需输入信息进行记录" />
          </el-form-item>
        </el-form>
      </div>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="closeEditDeviceDialog">取消</el-button>
          <el-button type="primary" @click="updateDeviceHolderDialog">确认</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { getDeviceList, updateDeviceHolder } from '@/api/device'
import { ref } from 'vue'
import { ElMessage } from '_element-plus@2.0.1@element-plus'

const page = ref(1)
const total = ref(0)
const pageSize = ref(50)
const tableData = ref([])

// 查询
const getTableData = async() => {
  const table = await getDeviceList({
    page: page.value,
    pageSize: pageSize.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 弹窗相关

const deviceInfo = ref({
  id: '',
  holder: '',
  currentHolder: '',
  describe: '',
})

// 判断弹窗类型
const deviceForm = ref(null)
const updateDeviceHolderDialog = async() => {
  deviceForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...deviceInfo.value
      }
      if (dialogFlag.value === 'edit') {
        const res = await updateDeviceHolder(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeEditDeviceDialog()
        }
      }
    }
  })
}

// 关闭弹窗
const closeEditDeviceDialog = () => {
  deviceForm.value.resetFields()
  deviceInfo.value.holder = ''
  deviceInfo.value.currentHolder = ''
  deviceInfo.value.describe = ''
  editDeviceDialog.value = false
}

// 打开弹窗
const dialogFlag = ref('edit')
const editDeviceDialog = ref(false)
const openEdit = (row) => {
  dialogFlag.value = 'edit'
  deviceInfo.value = JSON.parse(JSON.stringify(row))
  editDeviceDialog.value = true
}

</script>
