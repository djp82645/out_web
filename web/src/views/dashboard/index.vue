<template>
  <div class="dashboard-container">
    <el-form :inline="true" size="medium" :model="form" class="demo-form-inline">
      <el-form-item style="width: 100px">
        <el-input v-model="form.sn" clearable placeholder="编号"/>
      </el-form-item>
      <el-form-item style="width: 100px">
        <el-select v-model="form.status" @change="search" clearable placeholder="状态">
          <el-option label="已使用" value="0"/>
          <el-option label="新号" value="1"/>
          <el-option label="老号" value="2"/>
          <el-option label="错误" value="3"/>
          <el-option label="已标记" value="4"/>
          <el-option label="未使用" value="-1"/>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-date-picker
          @change="search"
          v-model="form.createRange"
          type="datetimerange"
          range-separator="至"
          start-placeholder="添加时间"
          end-placeholder="结束时间"
        >
        </el-date-picker>
      </el-form-item>
      <el-form-item>
        <el-date-picker
          @change="search"
          v-model="form.updateRange"
          type="datetimerange"
          range-separator="至"
          start-placeholder="更新时间"
          end-placeholder="结束时间"
        >
        </el-date-picker>
      </el-form-item>
      <br>
      <el-form-item>
        <el-button type="primary" @click="search">条件查询</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handleSearchExport">条件导出</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="danger" @click="handleSearchDel">条件删除</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="danger" @click="handleSelectDel">选中删除</el-button>
      </el-form-item>
    </el-form>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
      size="medium"
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        type="selection"
        align="center"
        width="42"
      />
      <el-table-column align="center" prop="id" label="ID" width="95"/>
      <el-table-column align="center" prop="sn" label="编号" width="100"/>
      <el-table-column label="接口地址" show-overflow-tooltip>
        <template slot-scope="scope">
          {{ scope.row.api }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="状态" width="110" align="center">
        <template slot-scope="scope">
          <el-link :underline="false" :type="scope.row.status | statusType">{{ scope.row.status | statusText }}</el-link>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="createTime" label="添加时间" width="200" >
      <template slot-scope="scope">
          {{ scope.row.created_at  }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="updateTime" label="更新时间" width="200">
      <template slot-scope="scope">
          {{ scope.row.updated_at  }}
      </template>
      </el-table-column>
      <el-table-column
        fixed="right"
        align="center"
        label="操作"
        width="100"
      >
        <template slot-scope="scope">
          <!-- <el-link @click="handleDel(scope.row.id)" type="danger" size="small">更新状态</el-link> -->
          <el-link @click="handleDel(scope.row.id)" type="danger" size="small">删除</el-link>
        </template>
      </el-table-column>
    </el-table>
    <br>
    <el-pagination
      background
      hide-on-single-page
      :current-page="form.page"
      :page-size="form.size"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      layout="total, prev, pager, next"
      :total="total"
      :page-count="totalPage"
    />
  </div>
</template>

<script>
import { del, delIds, getList, searchDel } from '@/api/api'

export default {
  name: 'Dashboard',
  filters: {
    statusType(status) {
      switch (status) {
        case 0:
          return 'primary'
        case 1:
          return 'success'
        case 2:
          return 'warning'
        case 3:
          return 'danger'
      }
      return 'info'
    },
    statusText(status) {
      switch (status) {
        case -1:
          return '未使用'
        case 0:
          return '已使用'
        case 1:
          return '新号'
        case 2:
          return '老号'
        case 3:
          return '错误'
      }
      return '未使用'
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      totalPage: 0,
      listLoading: true,
      form: {
        sn: '',
        status: null,
        createRange: null,
        updateRange: null,
        page: 1,
        size: 10
      },
      selected: []
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    handleDel(id) {
      this.$confirm('此操作将永久删除该数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        del(id).then(_ => {
          this.$message.success('Success')
          if (this.list.length === 1 && this.form.page > 1) {
            this.form.page--
          }
          this.fetchData()
          this.$message({
            type: 'success',
            message: '删除成功!'
          });
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
    },
    handleSelectDel() {
      this.$confirm('此操作将永久删除选中数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        if (this.selected.length > 0) {
          this.listLoading = true
          let idArr = []
          for (const selectedKey of this.selected) {
            idArr.push(selectedKey.id)
          }
          let ids = idArr.join('-')
          delIds(ids).then(_ => {
            if (this.selected.length === this.list.length && this.form.page > 1) {
              this.form.page--
            }
            this.selected = []
            this.fetchData()
            this.$message({
              type: 'success',
              message: '删除成功!'
            });
          })
        } else {
          this.$message.warning('Please select')
        }
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });

    },
    handleSearchDel() {
      this.$confirm('此操作将永久删除条件过滤的数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        let query = {
          sn: this.form.sn,
          status: this.form.status,
          createRange: this.form.createRange &&
            this.form.createRange[0].getTime() + '-' + this.form.createRange[1].getTime(),
          updateRange: this.form.updateRange &&
            this.form.updateRange[0].getTime() + '-' + this.form.updateRange[1].getTime()
        }
        searchDel(query).then(_ => {
          this.$message.success('删除成功')
          this.search()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });

    },
    handleSearchExport() {
      let query = {
        sn: this.form.sn,
        status: this.form.status,
        createRange: this.form.createRange &&
          this.form.createRange[0].getTime() + '-' + this.form.createRange[1].getTime(),
        updateRange: this.form.updateRange &&
          this.form.updateRange[0].getTime() + '-' + this.form.updateRange[1].getTime()
      }
      let url = '/api/search-export.csv?sn=' + query.sn + '&status=' + query.status
        + '&createRange=' + (query.createRange ? query.createRange : '')
        + '&updateRange=' + (query.updateRange ? query.updateRange : '')
      window.open(url)
    },
    handleSelectionChange(arr) {
      this.selected = arr
    },
    search() {
      this.form.page = 1
      this.fetchData()
    },
    handleSizeChange(val) {
      this.form.page = 1
      this.form.size = val
      this.fetchData()
    },
    handleCurrentChange(val) {
      this.form.page = val
      this.fetchData()
    },
    fetchData() {
      this.listLoading = true

      let query = {
        page: this.form.page,
        size: this.form.size,
        sn: this.form.sn,
        status: this.form.status,
        createRange: this.form.createRange &&
          this.form.createRange[0].getTime() + '-' + this.form.createRange[1].getTime(),
        updateRange: this.form.updateRange &&
          this.form.updateRange[0].getTime() + '-' + this.form.updateRange[1].getTime()
      }
      try {
        getList(query).then(ret => {
          this.list = ret.data.data
          this.total = ret.data.totalCount
          this.totalPage = ret.data.totalPage
          this.listLoading = false
        }).catch(error => {
          console.error('Error fetching data:', error)
          this.listLoading = false
          // 这里可以添加更多的错误处理逻辑，例如显示错误提示
        })
      } catch (error) {
        console.error('Error in fetchData:', error)
        this.listLoading = false
        // 这里可以添加更多的错误处理逻辑，例如显示错误提示
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }

  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
