<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="0">
      <el-form-item label="">
        <el-input :autofocus="true" v-model="form.text" :rows="10" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">批量添加</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { add } from '@/api/api'

export default {
  data() {
    return {
      form: {
        text: ''
      }
    }
  },
  methods: {
    onSubmit() {
      if (this.form.text === '') {
        this.$message.error('text cannot be empty')
        return
      }
      let arr = this.form.text.replaceAll('\r\n', '\n').split('\n');
      const arr1 = arr
      add({data: arr1}).then(ret => {
        if (ret && ret.code && ret.code === 20000) {
          this.form.text = ''
          this.$message.success(ret.msg)
        } else {
          this.$message.error('Error occurred while submitting the form')
        }
      })

    }
  }
}
</script>

<style scoped>
.line {
  text-align: center;
}
</style>

