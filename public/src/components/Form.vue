<template>
  <ConnectServerFrom @connectServer="connectServer"></ConnectServerFrom>
  <div class="top_place"></div>

  <el-card class="box-card" v-if="basedata.connected">
    <template #header>
      <div class="card-header">
        <span>GRPC Request Form</span>
      </div>
    </template>
    <el-form :model="formRequest" ref="formRef">
      <el-row>
        <el-alert title="only activated devices will show in the list" type="success" style="margin-bottom: 10px;"/>
        <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
          <el-form-item label="Select Device:" prop="device_id" :rules="{
            required: true,
            message: 'device can not be null',
            trigger: 'blur',
          }">
            <el-cascader :props="props" v-model="formRequest.device_id"  style="width:550px;"/>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
    <div v-for="rtype in ['settings', 'commands']">
      <div>
        <h2>{{ rtype }}</h2>
      </div>
      <el-row v-for="item, key in settings[rtype]">
        <el-form :inline="true" :ref="'formRef_' + key">
          
          <el-form-item>
            <el-input :value="key" disabled></el-input>
          </el-form-item>
          <el-form-item label=""  prop="content" style="width: 160px;">
            <el-input-number v-if="isNum(item.conversion) && item.length > 0" v-model="scforms['content' + key]" :max="item.max" :min="item.min"></el-input-number>
            <el-input v-if="(item.conversion == 'string' || item.conversion == 'byte_array') && item.length > 0" v-model="scforms['content' + key]" :minlength="item.length"
              :maxlength="item.length * 2" prop="content" placeholder="input string"/>
            <el-radio-group v-if="item.conversion == 'bool' && item.length > 0"  v-model="scforms['content' + key]">
              <el-radio-button label="true" />
              <el-radio-button label="false" />
            </el-radio-group>
          </el-form-item>
          <el-form-item>
            <div class="button_right">
              <el-button type="primary" @click="doReq(formRef, key, item, rtype)">Send Request</el-button>
            </div>
          </el-form-item>
        </el-form>
      </el-row>
    </div>



  </el-card>

  <div class="top_place"></div>

  <el-card class="box-card" v-if="basedata.connected">
    <template #header>
      <div class="card-header">
        <span>Request Records</span>
      </div>
    </template>

    <el-table show-header="true" :data="basedata.reqeustRecords" stripe="true">
      <el-table-column prop="dateTime" label="dateTime" />
      <el-table-column prop="deviceID" label="deviceID" />
      <el-table-column prop="Payload" label="Payload" />
      <el-table-column prop="FCnt" label="FCnt" />
    </el-table>


  </el-card>
</template>
  
<script lang="ts" setup>
import { reactive, VueElement, ref, watch } from 'vue'
import settings from '../js/settings_template'
import { request } from '../js/request'
import type { CascaderProps } from 'element-plus'
import ConnectServerFrom from './ConnectServerForm.vue'
import type { FormInstance, FormRules } from 'element-plus'
import {lsave,lget,getSet} from '../js/localstore'
import { filter } from 'lodash'
const formRef = ref<FormInstance>()

let basedata = reactive({
  connected: false,
  template_id:"default",
  orgList: [],
  serverUrl: "",
  apiKey: "",
  reqeustRecords: [],
})
let scforms = reactive({})
let formRequest = reactive({
  device_id: '',
})
getSet('basedata',basedata)
getSet('scforms',scforms)
getSet('formRequest',formRequest)

watch([basedata, scforms,formRequest], ([basedata, scforms,formRequest]) => {
    lsave('basedata',basedata)
    lsave('scforms',scforms)
    lsave('formRequest',formRequest)
})

const doReq = (formRef, idKey, idValue, rtype) => {
  formRef.validate((valid) => {
    if (valid) {
      lsave('scforms',scforms)
      if (idValue.length > 0 && scforms['content' + idKey] == undefined) {
        alert('content can not be null')
        return
      }
      var data = {
        server_url: basedata.serverUrl,
        api_key: basedata.apiKey,
        id: idValue.id,
        device_id: formRequest.device_id[2],
        request_type: rtype,
        content: scforms['content' + idKey],
        content_type: idValue.conversion,
        content_length: idValue.length,
      }
      request('v1/device/queue', 'POST', data).then((resp) => {
        addRecord(data.device_id, resp.Paylod, resp.FCnt)
        alert('request successful! FCnt: ' + resp.FCnt + '; bytes: ' + resp.Paylod)

      }, (err) => {
        alert('request err :' + err)
      })
    }
  })
}

const connectServer = (formServer, config, orgList) => {
  basedata.apiKey = formServer.api_key
  basedata.orgList = orgList
  basedata.connected = false
  setTimeout(() => {
    basedata.connected = true
  }, 100)

  basedata.serverUrl = formServer.server_url
}




const isNum = (value) => {
  return value == 'uint8' || value == 'uint32' || value == 'int32' || value == 'float' || value == 'uint16' ||value == 'int32'
}
let id = 0
const props: CascaderProps = {
  lazy: true,
  lazyLoad(node, resolve) {
    const { level } = node
    console.log(node)
    if (level == 0) {
      const nodes = Array.from(basedata.orgList).map((item) => ({
        value: item.id,
        label: item.name,
        leaf: false,
      }))
      resolve(nodes)
    } else if (level == 1) {
      listRequest(node, resolve, 'app')
    } else if (level == 2) {
      listRequest(node, resolve, 'device')
    }
  }
}
const listRequest = (node, resolve, type) => {
  var data = {
    "server_url": basedata.serverUrl,
    "api_key": basedata.apiKey,
    "list_type": type,
    "id": node.data.value,
  }
  request('v1/list', 'POST', data).then((resp) => {
    if (resp != null && resp.length > 0) {
      
      const nodes = Array.from(resp.filter((item)=>{
        if (type != 'device') {
          return true
        }
        if (item.last_seen_at) {
          return true
        }
        return false
      })).map((item) => ({
        value: (type == 'device') ? item.dev_eui : item.id,
        label: item.name,
        leaf: node.level >= 2,
      }))
      resolve(nodes)
    } else {
      resolve([])
    }
  }, (err) => {
    alert(err)
  })
}
interface Record {
  dateTime: string
  deviceID: string
  Payload: string
  FCnt: Number
}

const addRecord = (deviceID, Payload, FCnt) => {
  basedata.reqeustRecords.push({
    dateTime: (new Date).toISOString(),
    deviceID: deviceID,
    Payload: Payload,
    FCnt: FCnt,
  })
  console.log(basedata.reqeustRecords)
}
</script>
<style>
.button_right {
  width: 100%;
  display: flex;
  justify-content: end;
  align-items: end;
}

.top_place {
  margin-top: 50px;
}</style>
  