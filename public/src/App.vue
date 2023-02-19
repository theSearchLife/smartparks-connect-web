
<template>
  <div class="common-layout">
    <el-container>
      <el-header><h1><el-icon><Operation /></el-icon>Chirpstack Grpc Client</h1></el-header>
      <el-main>
        <Form/>
      </el-main>
    </el-container>
  </div>
</template>
<style>
h1{
  text-align: center;
}
</style>
<script>
import Form from './components/Form.vue'

export default {
  components: {
    Form
  }
}
</script>