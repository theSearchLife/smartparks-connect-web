<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>Connect Server Form</span>
        <span class="float_right" v-if="config.connected"><el-tag class="ml-2" type="success">Service
            Connected</el-tag></span>
      </div>
    </template>

    <el-form :model="formServer" status-icon inline ref="formSRef" label-width="150px">
      <el-row>
        <el-form-item label="Device Template" :width="500" prop="device_template" :rules="{
          required: true,
          message: 'Device Template is required',
          trigger: 'blur',
        }">
          <el-select v-model="formServer.device_template" class="m-2" placeholder="Select" @change="changeTemplate"
            style="width: 350px;">
            <el-option v-for="item, key in deviceTemplates" :key="key" :label="key" :value="key" />
          </el-select>
        </el-form-item>
      </el-row>


      <el-row>
        <el-form-item label="Server Address" prop="server_url" :rules="{
          validator: validateGrpcHostPort,
          required: true,
          trigger: 'blur',
        }">
          <el-autocomplete v-model="formServer.server_url" name="serverURL" placeholder="host:port"
            :fetch-suggestions="querySearch" @select="handleSelect" style="width: 350px;" />
        </el-form-item>
      </el-row>

      <el-row>
        <el-form-item label="API Key" prop="api_key" :rules="{
          required: true,
          message: 'API Key is required',
          trigger: 'blur',
        }">
          <el-input v-model="formServer.api_key" type="password" show-password style="width: 350px;"></el-input>
        </el-form-item>

        <el-form-item>
          <el-popover placement="bottom" title="Get API Key" :width="200" trigger="hover"
            content="Click this button to obtain a valid API key using your server email and password">
            <template #reference>
              <el-button type="danger" :icon="Key" @click="config.dialogFormVisible = true">Get ApiKey</el-button>
            </template>
          </el-popover>
        </el-form-item>
      </el-row>


      <el-row>
        <el-form-item>
          <div class="button_right">
            <el-button type="primary" @click="connectServerSubmit(formSRef)">Connect Server</el-button>
          </div>
        </el-form-item>
        <div class="float_right">
          <ClearCacheButton />
        </div>
      </el-row>

    </el-form>

  </el-card>

  <el-dialog v-model="config.dialogFormVisible" title="Get API Key By email and password">
    <el-form :model="formLogin">
      <el-form-item label="Email" required>
        <el-input v-model="formLogin.email" />
      </el-form-item>
      <el-form-item label="Password" required>
        <el-input type="password" v-model="formLogin.password" autocomplete="off" show-password />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="config.dialogFormVisible = false">Cancel</el-button>
        <el-button type="primary" @click="userLogin">
          Get API Key
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { getSet, ladd, lget, lsave } from '@/js/localstore';
import { request } from "@/js/request";
import { Key } from '@element-plus/icons-vue';
import type { FormInstance } from 'element-plus';
import { filter } from 'lodash';
import { reactive, ref, watch } from 'vue';
import ClearCacheButton from './ClearCacheButton.vue';
const emit = defineEmits(["connectServer"]);
const formSRef = ref<FormInstance>()
const deviceTemplates = reactive({})

const validateGrpcHostPort = (_, value, callback) => {
  let re = /^[a-zA-Z\.\-0-9]{0,}:\d{1,5}$/
  if (value === '' || value == null) {
    callback(new Error('GRPC host:port is required'));
  } else if (!re.test(value)) {
    callback(new Error('Format must be: host:port (without any protocol prefix)'));
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
  alert("device template load error:" + err)
})


const formServer = reactive({
  device_template: '',
  server_url: '',
  api_key: ''
})
const formLogin = reactive({
  server_url: '',
  email: '',
  password: '',
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
        ladd("connectedServers", formServer)
        emit('connectServer', formServer, config, orgList, deviceTemplates)
      }, (err) => {
        alert(err)
      })
    } else {
      return false
    }
  })
}

const querySearch = (queryString: string, cb: any) => {
  console.log(queryString)
  const results = filter(lget("connectedServers"), createFilter(queryString))
  // call callback function to return suggestions
  cb(results)
}
const createFilter = (queryString: string) => {
  return (item) => {
    return (
      item.server_url.toLowerCase().indexOf(queryString.toLowerCase()) === 0
    )
  }
}

const handleSelect = (item) => {
  console.log(item)
  formServer.api_key = item.api_key
  formServer.device_template = item.device_template
  formServer.server_url = item.server_url
}

</script>
<style>
.float_right {
  display: block;
  float: right;
}
</style>