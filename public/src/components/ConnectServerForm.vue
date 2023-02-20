<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>Connect Server Form</span>
        <span class="float_right" v-if="config.connected"><el-tag class="ml-2" type="success">Service
            Connected</el-tag></span>
      </div>
    </template>

    <el-form :model="formServer" :inline="true" ref="formSRef">
      <el-row>
        <el-col span="24">
          <el-form-item label="Device Template:" prop="device_template" :rules="{
            required: true,
            message: 'device template can not be null',
            trigger: 'blur',
          }">
            <el-select v-model="formServer.device_template" class="m-2" placeholder="Select" @change="changeTemplate"
              style="width: 200px;">
              <el-option v-for="item, key in deviceTemplates" :key="key" :label="key" :value="key" />
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="GRPC Host:Port" width="200" prop="server_url" :rules="{
        validator: validateGrpcHostPort,
        trigger: 'blur',
      }">
        <el-input v-model="formServer.server_url" autocomplete="on" />
      </el-form-item>


      <el-form-item label="ApiKey" prop="api_key" :rules="{
        required: true,
        message: 'api_key can not be null',
        trigger: 'blur',
      }">
        <el-input v-model="formServer.api_key"></el-input>
      </el-form-item>

      <el-form-item>
        <el-popover placement="bottom" title="Get User's ApiKey" :width="200" trigger="hover"
          content="Click this button, you can obtain apikey through the user's email/password">
          <template #reference>
            <el-button type="danger" :icon="Key" @click="config.dialogFormVisible = true">GetUserApiKey</el-button>
          </template>
        </el-popover>
      </el-form-item>


      <el-form-item>
        <div class="button_right">
          <el-button type="primary" @click="connectServerSubmit(formSRef)">Connect Server</el-button>
        </div>

      </el-form-item>


    </el-form>

  </el-card>

  <el-dialog v-model="config.dialogFormVisible" title="Get APIKey By User">
    <el-form :model="formLogin">
      <el-form-item label="email">
        <el-input v-model="formLogin.email" autocomplete="off" />
      </el-form-item>
      <el-form-item label="password">
        <el-input type="password" v-model="formLogin.password" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="config.dialogFormVisible = false">Cancel</el-button>
        <el-button type="primary" @click="userLogin">
          GetApiKey
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { reactive, VueElement, ref, watch, onBeforeMount } from 'vue'
import { request } from '../js/request'
import { CascaderProps, popoverEmits } from 'element-plus'
import { Key, User } from '@element-plus/icons-vue'
import type { FormInstance, FormRules } from 'element-plus'
import { lsave, lget, getSet } from '../js/localstore'
const emit = defineEmits(["connectServer"]);
const formSRef = ref<FormInstance>()
const deviceTemplates = reactive({})
var validateGrpcHostPort = (rule, value, callback) => {
  let re = /^[a-zA-Z\.\-0-9]{0,}:\d{1,5}$/
  if (value === '' || value == null) {
    callback(new Error('Please input the grpc host:port'));
  } else if (!re.test(value)) {
    callback(new Error('Format must be:  host:port (without any protocol prefix)'));
  } else {
    callback();
  }
};
request('v1/template/list', "GET").then((resp) => {
  for (let i in resp) {
    deviceTemplates[i] = resp[i]
  }
  console.log(deviceTemplates)
  initModel()
}, (err) => {
  alert("device template load error")
})


const formServer = reactive({
  device_template: "",
  server_url: '',
  api_key: ''
})
const formLogin = reactive({
  email: "",
  password: "",
})

const initModel = () => {
  getSet('formServer', formServer)
  getSet('formLogin', formLogin)
}

watch([formServer, formLogin], ([formServer, formLogin]) => {
  lsave('formServer', formServer)
  lsave('formLogin', formLogin)
})
const config = reactive({
  connected: false,
  dialogFormVisible: false,
})

const userLogin = () => {
  formLogin.server_url = formServer.server_url
  request('v1/login', 'POST', formLogin).then((resp) => {
    formServer.api_key = resp
    config.dialogFormVisible = false
  }, (err) => {
    alert(err)
  })
}
var orgList = new Array;
const changeTemplate = () => {
  emit('connectServer', formServer, config, orgList, deviceTemplates)
}
const connectServerSubmit = (formEl) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      let data = {}
      data['server_url'] = formServer.server_url
      data['api_key'] = formServer.api_key
      data['list_type'] = 'org'
      request('v1/list', 'POST', data).then((resp) => {
        orgList = []
        for (var i = 0; i < resp.length; i++) {
          orgList.push({ "id": resp[i].id, "name": resp[i].name })
        }
        config.connected = true
        emit('connectServer', formServer, config, orgList, deviceTemplates)
      }, (err) => {
        alert(err)
      })
    } else {
      return false
    }
  })
}

</script>
<style>
.float_right {
  display: block;
  float: right;
}
</style>