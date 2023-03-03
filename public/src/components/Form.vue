<template>
  <el-tabs type="border-card">
    <el-tab-pane label="Connect Server Form">
      <ConnectServerFrom @connectServer="connectServer"></ConnectServerFrom>
      <div class="top_place"></div>

      <el-card class="box-card" v-if="basedata.connected">
        <template #header>
          <div class="card-header">
            <span>Select Device</span>
          </div>
        </template>
        <el-form :model="formRequest" ref="formRef">
          <el-row>
            <el-alert title="only activated devices will show in the list" type="success"
              style="margin-bottom: 10px;width:660px;" />
            <el-col :xs="24" :sm="24" :md="16" :lg="16" :xl="16">
              <el-form-item prop="data" :rules="{
                required: true,
                message: 'device can not be null',
                trigger: 'blur',
              }">
                <el-cascader :props="props" v-model="formRequest.data" style="width:550px;" />
              </el-form-item>
            </el-col>
            <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">

            </el-col>
          </el-row>
        </el-form>
      </el-card>

      <el-card class="box-card" v-if="basedata.connected">
        <template #header>
          <div class="card-header">
            <span>Downlink message form</span>
          </div>
        </template>
        <el-tabs type="border-card">
          <el-tab-pane :label="rtype" v-for="rtype in ['settings', 'commands']">
            <div>
              <div>
                <h2>{{ rtype }}</h2>
              </div>
              <el-row v-for="item, key in basedata.deviceTemplate[rtype]">
                <el-form :inline="true" :ref="'formRef_' + key" v-if="key != 'type' && key != 'port'">

                  <el-form-item>
                    <el-input :value="key" disabled></el-input>
                  </el-form-item>
                  <el-form-item label="" prop="content">
                    <el-input-number v-if="isNum(item.conversion) && item.length > 0"
                      v-model="scforms['content' + key + basedata.deviceTemplateVersion]" :max="item.max"
                      :min="item.min"></el-input-number>
                    <el-input v-if="(item.conversion == 'string' || item.conversion == 'byte_array') && item.length > 0"
                      v-model="scforms['content' + key + basedata.deviceTemplateVersion]" prop="content"
                      placeholder="input string" />
                    <el-radio-group v-if="item.conversion == 'bool' && item.length > 0"
                      v-model="scforms['content' + key + basedata.deviceTemplateVersion]">
                      <el-radio-button label="true" />
                      <el-radio-button label="false" />
                    </el-radio-group>
                  </el-form-item>
                  <el-form-item label="confirmed">
                    <el-switch v-model="scforms['confirmed_' + key + basedata.deviceTemplateVersion]"
                      active-color="#13ce66" inactive-color="#ff4949">
                    </el-switch>
                  </el-form-item>
                  <el-form-item>
                    <div class="button_right">
                      <el-button type="primary" @click="doReq(formRef, key, item, rtype)">Send Request</el-button>
                    </div>
                  </el-form-item>
                </el-form>
              </el-row>
            </div>

          </el-tab-pane>
        </el-tabs>
      </el-card>

      <div class="top_place"></div>
    </el-tab-pane>
    <el-tab-pane label="Request Records">
      <el-card class="box-card" v-if="basedata.connected">
        <template #header>
          <div class="card-header">
            <span>Request Records</span>
          </div>
        </template>

        <el-table show-header=true :data="basedata.requestRecords" stripe=true>
          <el-table-column prop="dateTime" label="Date + Time" />
          <el-table-column prop="name" label="Name" />
          <el-table-column prop="devEUI" label="EUI" />
          <el-table-column prop="type" label="Type" />
          <el-table-column prop="port" label="Port" />
          <el-table-column prop="payload" label="Payload (hex)" />
          <el-table-column prop="base64" label="Payload (base64)" />
          <el-table-column prop="fCnt" label="FCnt" />
          <el-table-column prop="confirmed" label="Confirmed" />
        </el-table>


      </el-card>
    </el-tab-pane>
  </el-tabs>
</template>
  
<script lang="ts" setup>
import { reactive, VueElement, ref, watch, nextTick } from 'vue'
import { request } from '../js/request'
import type { CascaderProps } from 'element-plus'
import ConnectServerFrom from './ConnectServerForm.vue'
import type { FormInstance, FormRules } from 'element-plus'
import { lsave, lget, getSet } from '../js/localstore'
const formRef = ref<FormInstance>()


let basedata = reactive({
  connected: false,
  deviceTemplate: {},
  orgList: [],
  serverUrl: "",
  apiKey: "",
  requestRecords: [],
  deviceTemplateVersion: "",
})
let scforms = reactive({})
let formRequest = reactive({
  data: '',
})
const initModel = () => {
  getSet('basedata', basedata)
  getSet('scforms', scforms)
  getSet('formRequest', formRequest)
  basedata.connected = false
}
initModel()

const setDefault = (deviceTemplate) => {
  for (var idKey in deviceTemplate['settings']) {
    if (scforms['content' + idKey + basedata.deviceTemplateVersion] == undefined && deviceTemplate['settings'][idKey].default != undefined) {
      scforms['content' + idKey + basedata.deviceTemplateVersion] = deviceTemplate['settings'][idKey].default
    }

    if (scforms['confirmed_' + idKey + basedata.deviceTemplateVersion] == undefined) {
      scforms['confirmed_' + idKey + basedata.deviceTemplateVersion] = false
    }
  }
  for (var idKey in deviceTemplate['commands']) {
    if (scforms['content' + idKey + basedata.deviceTemplateVersion] == undefined && deviceTemplate['commands'][idKey].default != undefined) {
      scforms['content' + idKey + basedata.deviceTemplateVersion] = deviceTemplate['commands'][idKey].default
    }
    if (scforms['confirmed_' + idKey + basedata.deviceTemplateVersion] == undefined) {
      scforms['confirmed_' + idKey + basedata.deviceTemplateVersion] = false
    }
  }
}


watch([basedata, scforms, formRequest], ([basedata, scforms, formRequest]) => {
  lsave('basedata', basedata)
  lsave('scforms', scforms)
  lsave('formRequest', formRequest)
})

const doReq = (formRef, idKey, idValue, rtype) => {
  formRef.validate((valid) => {
    if (valid) {
      lsave('scforms', scforms)
      if (idValue.length > 0 && scforms['content' + idKey + basedata.deviceTemplateVersion] == undefined) {
        alert('content can not be null')
        return
      }
      console.log(basedata)

      var tt = formRequest.data[2].split(":")
      var dev_eui = tt[0]
      var name = tt[1]

      var data = {
        server_url: basedata.serverUrl,
        api_key: basedata.apiKey,
        id: idValue.id,
        dev_eui: dev_eui,
        confirmed: scforms['confirmed_' + idKey + basedata.deviceTemplateVersion],
        request_type: rtype,
        port: basedata.deviceTemplate.ports[basedata.deviceTemplate[rtype].port],
        content: replaceBytesToStr(scforms['content' + idKey + basedata.deviceTemplateVersion]),
        content_type: idValue.conversion,
        content_length: idValue.length,
      }
      request('v1/device/queue', 'POST', data).then((resp) => {
        addRecord(data.dev_eui, name, data.request_type, data.port, resp.Paylod, resp.Base64, resp.FCnt, data.confirmed)
        alert('request successful! ' + 'eui:' + data.dev_eui + ' name:' + name + 'FCnt: ' + resp.FCnt + '; bytes: ' + resp.Paylod + ' ; base64: ' + resp.Base64 + ' ; confirmed: ' + data.confirmed)
      }, (err) => {
        alert('request err :' + err)
      })
    }
  })
}
const replaceBytesToStr = (bytes) => {
  if (bytes == undefined || bytes.replaceAll == undefined) {
    return bytes
  }
  return bytes.replaceAll('{', '').replaceAll('}', '').replaceAll('0x', '').replaceAll(',', '').replaceAll(' ', '')
}
const connectServer = (formServer, config, orgList, deviceTemplates) => {
  basedata.apiKey = formServer.api_key
  basedata.deviceTemplate = deviceTemplates[formServer.device_template]
  basedata.deviceTemplateVersion = formServer.device_template
  console.log(basedata.deviceTemplate)
  basedata.orgList = orgList
  basedata.connected = false
  nextTick(() => {
    basedata.connected = config.connected
    setDefault(basedata.deviceTemplate)
  })

  basedata.serverUrl = formServer.server_url
}




const isNum = (value) => {
  return value == 'uint8' || value == 'uint32' || value == 'int32' || value == 'float' || value == 'uint16' || value == 'int32'
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

      const nodes = Array.from(resp.filter((item) => {
        if (type != 'device') {
          return true
        }
        if (item.last_seen_at) {
          return true
        }
        return false
      })).map((item) => ({
        value: (type == 'device') ? item.dev_eui + ':' + item.name : item.id,
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
  Base64: string
  FCnt: Number
}

const addRecord = (devEUI, name, type, port, payload, base64, fCnt, confirmed) => {
  basedata.requestRecords.push({
    name: name,
    dateTime: (new Date).toISOString(),
    devEUI: devEUI,
    type: type,
    port: port,
    payload: payload,
    fCnt: fCnt,
    base64: base64,
    confirmed: confirmed,
  })
  console.log(basedata.requestRecords)
}
</script>