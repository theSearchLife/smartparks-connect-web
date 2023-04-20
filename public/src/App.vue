
<template>
  <Menu />
  <div class="common-layout">
    <el-container>
      <el-header>
        <h1><img style="width: 40px;position: relative;top: 10px;left: -10px;" src="./assets/logo.png">Smart Parks Connect
        </h1>
      </el-header>
      <el-main>
        <Form />
      </el-main>
    </el-container>
  </div>
</template>
<style>
h1 {
  text-align: center;
}
</style>
<script lang="ts">
import Form from './components/Form.vue';
import Menu from './components/Menu.vue';

export default {
  components: {
    Form,
    Menu
  }
}
</script>