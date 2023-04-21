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
            <el-alert title="Only activated devices will show in the list." type="success" />
            <el-col :xs="24" :sm="24" :md="16" :lg="16" :xl="16">
              <el-form-item prop="data" :rules="{
                  required: true,
                  message: 'device can not be null',
                  trigger: 'blur',
                }">
                <el-cascader :props="props" v-model="formRequest.data" style="width:550px;" />
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </el-card>

      <div class="top_place"></div>
      <el-card class="box-card requests-form" v-if="basedata.connected">
        <template #header>
          <div class="card-header">
            <span>Downlink Message Form</span>
          </div>
        </template>
        <el-tabs type="border-card">
          <el-tab-pane :label="rtype" v-for="rtype in ['settings', 'commands']">
            <div>
              <div>
                <h2>{{ rtype }}</h2>
              </div>
              <el-row v-for="item, key in basedata.deviceTemplate[rtype.toString()]">
                <el-form inline :ref="'formRef_' + key.toString()"
                  v-if="key.toString() != 'type' && key.toString() != 'port'" style="width: 100%;display: flex;">
                  <el-form-item>
                    <div class="button_right">
                      <el-button type="primary" @click="doReq(formRef, key, item, rtype)">Send Request</el-button>
                    </div>
                  </el-form-item>
                  <el-form-item label="confirmed">
                    <el-switch v-model="scforms['confirmed_' + key.toString() + basedata.deviceTemplateVersion]"
                      active-color="#13ce66" inactive-color="#ff4949">
                    </el-switch>
                  </el-form-item>
                  <el-form-item>
                    <el-input :value="key" disabled style="min-width: 200px;"></el-input>
                  </el-form-item>

                  <el-form-item label="" prop="content" style="width: 100%; margin-right: 0;">
                    <el-input-number v-if="isNum(item.conversion) && item.length > 0"
                      v-model="scforms['content' + key.toString() + basedata.deviceTemplateVersion]" :max="item.max"
                      :min="item.min"></el-input-number>
                    <el-input v-if="(item.conversion == 'string' || item.conversion == 'byte_array') && item.length > 0"
                      v-model="scforms['content' + key.toString() + basedata.deviceTemplateVersion]" prop="content"
                      placeholder="input string" style="width: 100%;" />
                    <el-radio-group v-if="item.conversion == 'bool' && item.length > 0"
                      v-model="scforms['content' + key.toString() + basedata.deviceTemplateVersion]">
                      <el-radio-button label="true" />
                      <el-radio-button label="false" />
                    </el-radio-group>
                  </el-form-item>
                </el-form>
              </el-row>
            </div>

          </el-tab-pane>
        </el-tabs>
      </el-card>
    </el-tab-pane>
    <el-tab-pane label="Request Records">
      <el-card class="box-card" v-if="basedata.connected">
        <template #header>
          <div class="card-header">
            <span>Request Records</span>
          </div>
        </template>

        <el-table show-header :data="basedata.requestRecords" stripe>
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
import type { CascaderProps, FormInstance } from 'element-plus'
import { ElNotification } from 'element-plus'
import { nextTick, reactive, ref, watch } from 'vue'
import { getSet, lsave } from '../js/localstore'
import { request } from "../js/request"

import ConnectServerFrom from './ConnectServerForm.vue'
const formRef = ref<FormInstance>()

let basedata = reactive({
  connected: false,
  deviceTemplate: {
    ports: "",
  },
  orgList: [],
  serverUrl: "",
  apiKey: "",
  requestRecords: [],
  deviceTemplateVersion: "",
})
let scforms = reactive({})
let formRequest = reactive({
  data: [],
})
const initModel = () => {
  getSet('basedata', basedata)
  getSet('scforms', scforms)
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
})

const doReq = (formRef, idKey, idValue, rtype) => {
  formRef.validate((valid) => {
    if (valid) {
      lsave('scforms', scforms)
      if (idValue.length > 0 && scforms['content' + idKey + basedata.deviceTemplateVersion] == undefined) {
        alert('content can not be null')
        return
      }

      const promises = []
      for (var i = 0; i < formRequest.data.length; i++) {
        let formData = formRequest.data[i]
        let tt = formData[2].split(":")
        let dev_eui = tt[0]
        let name = tt[1]

        let data = {
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

        const promise = request('v1/device/queue', 'POST', data).then((resp) => {
          addRecord(data.dev_eui, name, data.request_type, data.port, resp.Payload, resp.Base64, resp.FCnt, data.confirmed)
          return {
            dev_eui: data.dev_eui,
            name: name,
            fcnt: resp.FCnt,
            bytes: resp.Paylod,
            base64: resp.Base64,
            confirmed: data.confirmed
          }
        }, (err) => {
          return {
            dev_eui: data.dev_eui,
            name: name,
            error: err
          }
        })
        promises.push(promise)
      }

      Promise.all(promises).then((results) => {
        let errors = []
        let messages = []
        for (var i = 0; i < results.length; i++) {
          let result = results[i]
          if ("error" in results[i]) {
            console.log("error", result)
            errors.push(result.name + `: Error (${result.error})`)
          } else {
            console.log("success", result)
            messages.push(result.name + ": Success")
          }
        }
        if (errors.length == results.length) {
          ElNotification({
            title: 'All requests failed!',
            dangerouslyUseHTMLString: true,
            message: errors.join("<br />"),
            type: 'error',
            duration: 0,
          })
        } else if (errors.length > 0) {
          ElNotification({
            title: 'Some requests failed!',
            dangerouslyUseHTMLString: true,
            message: errors.join("<br />") + "<br />" + messages.join("<br />"),
            type: 'warning',
            duration: 0,
          })
        } else {
          ElNotification({
            title: 'All requests succeded!',
            dangerouslyUseHTMLString: true,
            message: messages.join("<br />"),
            type: 'success',
            duration: 0,
          })
        }
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
const props: CascaderProps = {
  lazy: true,
  lazyLoad(node, resolve) {
    const { level } = node
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
  },
  multiple: true,
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
}
</script>