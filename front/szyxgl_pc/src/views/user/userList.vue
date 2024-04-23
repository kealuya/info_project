<template>
  <div class="detailbox">
    <el-row :gutter="24">
      <el-col :span="24">
        <el-card class="detailbox-card">
          <div class="title">
            <div class="title-sag">
              <img src="@/assets/layouts/companyInfo/basic.png" alt="">
              <span style="font-size: 16px">角色列表</span>
            </div>
          </div>
          <el-table :data="filterTableData" style="width: 100%">
            <el-table-column label="角色名称" prop="name" />
            <el-table-column label="角色权限描述" prop="ms" />
            <el-table-column label="角色状态" prop="zt" />
            <el-table-column label="创建时间" prop="time" />
            <el-table-column align="right">
              <template #header>
                <el-input v-model="search" size="small" placeholder="输入角色名称" />
              </template>
              <template #default="scope">
                <el-button type="primary" size="small" @click="handleEdit(scope.$index, scope.row)">
                  设置权限
                </el-button>
                <el-button  type="primary" size="small" @click="handleEdit(scope.$index, scope.row)">
                  编辑权限
                </el-button>
                <el-button
                        size="small"
                        type="danger"
                        @click="handleDelete(scope.$index, scope.row)"
                >
                  删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>


        </el-card>
      </el-col>

    </el-row>
  </div>
</template>
<script setup lang="ts">
  import { onMounted, ref,computed } from 'vue';
  interface User {
    ms: string
    name: string
    zt: string
    time: string
  }

  const search = ref('')
  const filterTableData = computed(() =>
          tableData.filter(
                  (data) =>
                          !search.value ||
                          data.name.toLowerCase().includes(search.value.toLowerCase())
          )
  )
  const handleEdit = (index: number, row: User) => {
    console.log(index, row)
  }
  const handleDelete = (index: number, row: User) => {
    console.log(index, row)
  }
  const tableData: User[] = [
    {
      ms: '管理系统人员查看数据',
      name: '系统管理员',
      zt: '开启',
      time: '2024-04-22 09:41:15',
    },{
      ms: '用户默认为普通角色',
      name: '普通用户',
      zt: '开启',
      time: '2024-04-22 09:41:15',
    },{
      ms: '管理维护权限数据',
      name: '分销数据管理员',
      zt: '开启',
      time: '2024-04-22 09:41:15',
    }
    ]
</script>
<style lang="scss" scoped>

</style>
