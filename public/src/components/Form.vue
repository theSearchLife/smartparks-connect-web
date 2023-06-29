<template>
  <el-tabs type="border-card">
    <el-tab-pane label="Connect Server Form">
      <ConnectServerFrom @connectServer="connectServer"></ConnectServerFrom>
      <div class="top_place"></div>

      <el-card class="box-card" v-if="basedata.connected">
        <template #header>
          <template v-if="basedata.connectionType === 'chirpstack'">
            <div class="card-header">
              <span>Select Device</span>
            </div>
          </template>
          <template v-else>
            <div class="card-header">
              <span>Device IMEI</span>
            </div>
          </template>

        </template>
        <el-form :model="formRequest" ref="formRef">
          <template v-if="basedata.connectionType === 'chirpstack'">
            <el-row>
              <el-alert title="Only activated devices will show in the list." type="success" />
              <el-col :xs="24" :sm="24" :md="16" :lg="16" :xl="16">
                <el-form-item prop="data" :rules="{
                  required: true,
                  message: 'device can not be null',
                  trigger: 'blur',
                }">
                  <el-cascader :props="available_devices" v-model="formRequest.data" style="width:550px;" />
                </el-form-item>
              </el-col>
            </el-row>
          </template>
          <template v-else>
            <el-row>
              <el-alert title="Please enter device IMEI." type="success" />
              <el-col :xs="24" :sm="24" :md="16" :lg="16" :xl="16">
                <el-form-item label="IMEI" prop="imei" :rules="{
                  validator: validateIMEI,
                  required: true,
                  trigger: 'blur',
                }">
                  <el-input v-model="formRequest.imei" id="imei" name="imei" autocomplete="on" style="width:350px;" />
                </el-form-item>
              </el-col>
            </el-row>
          </template>
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
                  <el-form-item :label="basedata.connectionType === 'chirpstack' ? 'Confirmed' : 'Flush'">
                    <el-switch name="confirmed"
                      v-model="scforms['confirmed_' + key.toString() + basedata.deviceTemplateVersion]"
                      active-color="#13ce66" inactive-color="#ff4949">
                    </el-switch>
                  </el-form-item>
                  <el-form-item>
                    <el-input :value="key" disabled style="min-width: 200px;"></el-input>
                  </el-form-item>

                  <el-form-item label="" prop="content" style="width: 100%; margin-right: 0;">
                    <el-input-number v-if="isNum(item.conversion) && item.length > 0"
                      v-model="scforms['content' + key.toString() + basedata.deviceTemplateVersion]" :max="item.max"
                      :min="item.min" style="width:170px"></el-input-number>
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
      <template v-if="basedata.connected">
        <RockBLOCKRequests :requestRecords="basedata.rockBLOCKRequests" v-if="basedata.connectionType === 'rockblock'">
        </RockBLOCKRequests>
        <ChirpstackRequests :requestRecords="basedata.chipstackRequests" v-if="basedata.connectionType === 'chirpstack'">
        </ChirpstackRequests>
      </template>
    </el-tab-pane>
  </el-tabs>
</template>

<script lang="ts" setup>
import type { CascaderProps, FormInstance } from 'element-plus'
import { ElMessageBox } from 'element-plus'
import { nextTick, reactive, ref, watch } from 'vue'
import { getSet, lsave } from '../js/localstore'
import { request } from "../js/request"

import ConnectServerFrom from './ConnectServerForm.vue'
import ChirpstackRequests from './RequestHistory/ChirpstackRequests.vue'
import RockBLOCKRequests from './RequestHistory/RockBLOCKRequests.vue'
const formRef = ref<FormInstance>()

let basedata = reactive({
  connectionType: "",
  connected: false,
  deviceTemplate: {
    ports: "",
  },
  available_devices: [],
  serverUrl: "",
  apiKey: "",
  rockBLOCKRequests: [],
  chipstackRequests: [],
  deviceTemplateVersion: "",
  username: "",
  password: "",
})
let scforms = reactive({})
let formRequest = reactive({
  data: [],
  imei: ""
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

const validateIMEI = (_, value) => {
  if (value === "" || value == null) {
    return Promise.reject('Device IMEI can not be empty!');
  } else {
    const isValid = /^\d{15}$/.test(value);
    if (!isValid) {
      return Promise.reject('IMEI must be a 15-digit number');
    }
  }
  return Promise.resolve();
};

const doReq = (formRef, idKey, idValue, rtype) => {
  formRef.validate((valid) => {
    if (valid) {
      if (idValue.length > 0 && scforms['content' + idKey + basedata.deviceTemplateVersion] == undefined) {
        ElMessageBox.alert('Content can not be null!', 'Request failed!', {
          confirmButtonText: 'OK',
          dangerouslyUseHTMLString: true,
          type: "error",
        })
        return
      }
      lsave('scforms', scforms)
      if (basedata.connectionType === "chirpstack") {
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
            addChirpstackRequest(data.dev_eui, name, data.request_type, data.port, resp.Payload, resp.Base64, resp.FCnt, data.confirmed)
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
              if (result.error.response) {
                errors.push(result.name + `: Error (${result.error.response.data.err_msg})`)
              } else {
                errors.push(result.name + `: Error (${result.error})`)
              }
            } else {
              console.log("success", result)
              messages.push(result.name + ": Success")
            }
          }
          if (errors.length == results.length) {
            ElMessageBox.alert(errors.join("<br /><br />"), 'All requests failed!', {
              confirmButtonText: 'OK',
              dangerouslyUseHTMLString: true,
              type: "error",
            })
          } else if (errors.length > 0) {
            ElMessageBox.alert([errors.join("<br /><br />"), messages.join("<br /><br />")].join("<br /><br />"), 'Some requests failed!', {
              confirmButtonText: 'OK',
              dangerouslyUseHTMLString: true,
              type: "warning",
            })
          } else {
            ElMessageBox.alert(messages.join("<br /><br />"), 'All requests succeded!', {
              confirmButtonText: 'OK',
              dangerouslyUseHTMLString: true,
              type: "success",
            })
          }
        })
      } else {
        let data = {
          username: basedata.username,
          Password: basedata.password,
          id: idValue.id,
          imei: formRequest.imei,
          confirmed: scforms['confirmed_' + idKey + basedata.deviceTemplateVersion],
          request_type: rtype,
          port: basedata.deviceTemplate.ports[basedata.deviceTemplate[rtype].port],
          content: replaceBytesToStr(scforms['content' + idKey + basedata.deviceTemplateVersion]),
          content_type: idValue.conversion,
          content_length: idValue.length,
        }

        request('v1/rockblock/queue', 'POST', data).then((resp) => {
          console.log(resp)
          addRockBLOCKRequest(data.imei, data.request_type, data.port, resp.Payload, resp.Base64, resp.mtId, data.confirmed)
          ElMessageBox.alert(data.imei + ": Success", 'Request succeded!', {
            confirmButtonText: 'OK',
            dangerouslyUseHTMLString: true,
            type: "success",
          })
        }, (err) => {
          let error_message = err
          if (err.response) {
            error_message = data.imei + ": " + err.response.data.err_msg
          }
          ElMessageBox.alert(error_message, 'Request failed!', {
            confirmButtonText: 'OK',
            dangerouslyUseHTMLString: true,
            type: "error",
          })
        })
      }
    }
  })
}

const replaceBytesToStr = (bytes) => {
  if (bytes == undefined || bytes.replaceAll == undefined) {
    return bytes
  }
  return bytes.replaceAll('{', '').replaceAll('}', '').replaceAll('0x', '').replaceAll(',', '').replaceAll(' ', '')
}
const connectServer = (formServer, config, available_devices, deviceTemplates) => {
  console.log("connected ----", formServer)
  basedata.apiKey = formServer.api_key
  basedata.username = formServer.username
  basedata.password = formServer.password
  basedata.deviceTemplate = deviceTemplates[formServer.device_template]
  basedata.deviceTemplateVersion = formServer.device_template
  basedata.available_devices = available_devices
  basedata.connectionType = formServer.connection_type
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
const available_devices: CascaderProps = {
  lazy: true,
  lazyLoad(node, resolve) {
    const { level } = node
    if (level == 0) {
      const nodes = Array.from(basedata.available_devices).map((item) => ({
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
    let error_message = err
    if (err.response) {
      error_message = err.response.data.err_msg
    }
    ElMessageBox.alert(error_message, 'Request failed!', {
      confirmButtonText: 'OK',
      dangerouslyUseHTMLString: true,
      type: "error",
    })
  })
}

const addChirpstackRequest = (devEUI, name, type, port, payload, base64, fCnt, confirmed) => {
  // Check if chipstackRequests exists in basedata, if not create an ampty array
  if (basedata.chipstackRequests == undefined) {
    basedata.chipstackRequests = []
  }
  basedata.chipstackRequests.push({
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
const addRockBLOCKRequest = (IMEI, type, port, payload, base64, messageID, flush) => {
  // Check if rockBLOCKRequests exists in basedata, if not create an ampty array
  if (basedata.rockBLOCKRequests == undefined) {
    basedata.rockBLOCKRequests = []
  }
  basedata.rockBLOCKRequests.push({
    dateTime: (new Date).toISOString(),
    IMEI: IMEI,
    type: type,
    port: port,
    payload: payload,
    messageID: messageID,
    base64: base64,
    flush: flush,
  })
}
</script>