<template>
  <div class="contener">
    <div class="settingMenu">
      <el-button type="primary" :icon="Plus"> 新建</el-button>
      <el-button type="primary" :icon="Delete">删除 </el-button>
      <el-button :icon="Refresh" v-loading.fullscreen.lock="fullscreenLoading" type="primary"
        @click="openFullScreen1">刷新 </el-button>
      <el-button :icon="FolderOpened" @click="exportToExcel">导出 </el-button>
    </div>
    <el-table ref="multipleTableRef" :data="tableData" style="width: 100%" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" />

      <el-table-column label="日期" width="120">
        <template #default="scope">{{ scope.row.date }}</template>
      </el-table-column>
      <el-table-column property="name" label="姓名" width="120" />
      <el-table-column property="address" label="地址" show-overflow-tooltip />
      <el-table-column align="right">
        <!-- <template #header>
        <el-input v-model="search" size="small" placeholder="Type to search" />
      </template> -->
        <template #default="scope">
          <el-button size="small" type="primary" @click="handleRole">权限分配</el-button>
          <el-button size="small" type="success" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
          <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
        </template> 
      </el-table-column>
    </el-table>
    <div class="pagination">
      <el-pagination v-model:current-page="currentPage2" v-model:page-size="pageSize2" :page-sizes="[10, 20, 30, 40]"
        layout="sizes, prev, pager, next" :total="1000" @size-change="handleSizeChange"
        @current-change="handleCurrentChange" />
    </div>
    <!-- 分配权限 -->
    <div class="dialog">
      <el-dialog v-model="centerDialogVisible" title="分配权限" width="500" align-center v-if="centerDialogVisible" destroy-on-close :close-on-click-modal="false">
        <span>Open the dialog from the center from the screen</span>
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="centerDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="centerDialogVisible = false">
              确定
            </el-button>
          </div>
        </template>
      </el-dialog>
    </div>
  </div>

</template>

<script setup lang="ts">
import { reactive, ref, onMounted, nextTick, watch } from 'vue'
import { Plus, Delete, Refresh, FolderOpened } from '@element-plus/icons-vue'
import { ElLoading, ElTable, ElMessage, ElMessageBox } from 'element-plus'
//引入导出excel
// import XLSX from 'xlsx';
interface User {
  date: string
  name: string
  address: string
}
//分配权限模态
const centerDialogVisible = ref(false)
const handleRole = () => {
  centerDialogVisible.value = true
  console.log(' centerDialogVisible.value ', centerDialogVisible.value)
}
//导出excel
const exportToExcel = () => {
  console.log('aaa')

}

//刷新功能
const fullscreenLoading = ref(false)
const openFullScreen1 = () => {
  const loading = ElLoading.service({
    lock: true,
    text: '加载中....',
    background: 'rgba(0, 0, 0, 0.9)',
    target: document.querySelector('main') as unknown as string
  })
  setTimeout(() => {
    loading.close()
  }, 1000)
}
//分页部分
const currentPage2 = ref(5)
const pageSize2 = ref(10)
const handleSizeChange = (val: number) => {
  console.log(`${val} items per page`)
}
const handleCurrentChange = (val: number) => {
  console.log(`current page: ${val}`)
}
//表格修改
const handleEdit = (index: number, row: object) => {
  console.log(index, row)

}
//表格删除
const handleDelete = (index: number, row: any) => {
  console.log(index, row)
  ElMessageBox.confirm(
    '此操作将永久删除选中的数据, 是否继续？',
    '删除操作',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
      center: true,
    }
  )
    .then(() => {
      tableData.splice(index, 1)
      ElMessage({
        type: 'success',
        message: '删除成功',
      })
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '取消删除',
      })
    })


}
// 表格部分
const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<User[]>([])

const handleSelectionChange = (val: User[]) => {
  multipleSelection.value = val
  console.log('val', multipleSelection.value)
}
const tableData = reactive([
  {
    date: '2016-05-03',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-02',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-04',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-08',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-06',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-07',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
])
</script>
<style scoped lang='scss'>
.contener {
  width: 100%;
  // height: 100vh;
  // background: #fff;
  border-radius: 10px;
  padding: 20px;

  .settingMenu {
    margin-bottom: 20px;
  }

  .pagination {
    display: flex;
    justify-content: center;
    margin-top: 20px;
  }
}
</style>