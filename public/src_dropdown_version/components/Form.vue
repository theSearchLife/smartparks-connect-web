<template>
 <ConnectServerFrom @connectServer="connectServer"></ConnectServerFrom>
 <div class="top_place"></div>

  <el-card class="box-card" v-if="basedata.connected">
    <template #header>
      <div class="card-header">
        <span>GRPC Request Form</span>
      </div>
    </template>
    <el-form :model="formRequest"  ref="formRef" >
        <el-row>
            <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
        <el-form-item label="Select Device:" prop="device_id" :rules="{
        required: true,
        message: 'device can not be null',
        trigger: 'blur',
      }">
            <el-cascader :props="props" v-model="formRequest.device_id"/>
        </el-form-item>
        </el-col>
        <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
        <el-form-item label="RequestType:" prop="request_type" :rules="{
        required: true,
        message: 'request_type must bo choice',
        trigger: 'blur',
      }">
            <el-radio-group v-model="formRequest.request_type" @change="requestTypeChange">
                <el-radio-button label="settings" />
                <el-radio-button label="commands" />
            </el-radio-group>
        </el-form-item>
    </el-col>
    <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
            <el-form-item label="ID:" prop="id" :rules="{
        required: true,
        message: 'ID can not be null',
        trigger: 'blur',
      }">
                <el-select v-model="formRequest.id" class="m-2" placeholder="Select" @change="idChange">
                    <el-option
                    v-for="item,key in IDOptions"
                    :key="key"
                    :label="key"
                    :value="key"
                    />
                </el-select>
            </el-form-item>
            </el-col>
            </el-row>
        <el-row>
            <el-col :span="24">
        <el-form-item label="value" v-if="isNum(basedata.contentType) && basedata.contentLength>0 " prop="content" :rules="{
        required: true,
        message: 'content can not be null',
        trigger: 'blur',
      }">
    <el-input-number v-model="formRequest.content" :max="basedata.contentNumMax" :min="basedata.contentNumMin"></el-input-number>
        </el-form-item>
        <el-form-item label="value" v-if="(basedata.contentType == 'string' || basedata.contentType == 'byte_array')  && basedata.contentLength>0 " prop="content" :rules="{
        required: true,
        message: 'content can not be null',
        trigger: 'blur',
      }">
            <el-input v-model="formRequest.content" type="textarea" :rows="2" :minlength="basedata.contentMinLength" :maxlength="basedata.contentMaxLength" prop="content" :rules="{
        required: true,
        message: 'content can not be null',
        trigger: 'blur',
      }"/>
        </el-form-item>
        <el-form-item label="value" v-if="basedata.contentType == 'bool'  && basedata.contentLength>0 " prop="content" :rules="{
        required: true,
        message: 'content can not be null',
        trigger: 'blur',
      }">
        <el-radio-group v-model="formRequest.content" >
            <el-radio-button label="true" />
            <el-radio-button label="false" />
        </el-radio-group>
    </el-form-item>
    </el-col>
        </el-row>
        <el-form-item >
            <div class="button_right">
                <el-button type="primary" @click="submitForm(formRef)">Send Request</el-button>
            </div>
        </el-form-item>
    
       
    </el-form>
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
    <el-table-column prop="deviceID" label="deviceID"  />
    <el-table-column prop="Payload" label="Payload"  />
    <el-table-column prop="FCnt" label="FCnt" />
  </el-table>

    
</el-card>
    
</template>
  
<script lang="ts" setup>
  import { reactive, VueElement,ref } from 'vue'
  import settings from '../js/settings_template'
  import {request} from '../js/request'
  import type { CascaderProps } from 'element-plus'
  import ConnectServerFrom from './ConnectServerForm.vue'
  import type { FormInstance ,FormRules} from 'element-plus'
import { timePanelSharedProps } from 'element-plus/es/components/time-picker/src/props/shared'
  const formRef = ref<FormInstance>()

  const basedata = reactive({
    connected:false,
    orgList:[],
    serverUrl:"",
    apiKey:"",
    contentType:"",
    contentLength:0,
    contentNumMax:0,
    contentNumMin:0,
    contentMaxLength:0,
    contentMinLength:0,
    reqeustRecords: [],
  })
  

  const connectServer = (formServer,config,orgList)=>{
    basedata.apiKey = formServer.api_key
    basedata.orgList = orgList
    basedata.connected = false
    setTimeout(()=>{
        basedata.connected = true
    },100)
    
    basedata.serverUrl = formServer.server_url
  }
  const formRequest = reactive({
    id:'',
    device_id:'',
    request_type:'',
    content:'',
    content_type:'',
    content_length:'',
  })
  var IDOptions = [{}];
  const requestTypeChange = (value) =>{
    console.log(value,settings)
    IDOptions = settings[formRequest.request_type]
  }
  

const idChange = (value) =>{
    let setting = IDOptions[value]
    if (setting != undefined) {
        console.log(setting)
        basedata.contentType = setting['conversion']
        basedata.contentLength = setting['length']
        if (isNum(basedata.contentType) && setting['length'] > 0) {
            basedata.contentNumMax = setting['max']
            basedata.contentNumMin = setting['min']
        } else if (basedata.contentType == "string") {
            basedata.contentMinLength = 1
            basedata.contentMaxLength = setting['length']
        } else if (basedata.contentType == "byte_array") {
            basedata.contentMinLength = setting['length']*2
            basedata.contentMaxLength = setting['length']*2

        }
        formRequest.content_type = setting['conversion']
        formRequest.content_length = setting['length']
    }
    formRequest.content = ''
}
const isNum = (value)=>{
    return basedata.contentType == 'uint8' || basedata.contentType == 'uint32' || basedata.contentType == 'int32' || basedata.contentType == 'float'
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
    }else if (level == 1) {
        listRequest(node,resolve,'app')
    }else if (level == 2){
        listRequest(node,resolve,'device')
    }
  }
}
const listRequest = (node,resolve,type)=>{
        var data = {
            "server_url":basedata.serverUrl,
            "api_key":basedata.apiKey,
            "list_type":type,
            "id":node.data.value,
        }
        request('v1/list','POST',data).then((resp)=>{
            if (resp != null && resp.length > 0){
                const nodes = Array.from(resp).map((item) => ({
                value: (type =='device') ? item.dev_eui :item.id,
                label: item.name,
                leaf: node.level >= 2,
                }))
                resolve(nodes)
            }else{
                resolve([])
            }
        },(err)=>{
            alert(err)
        })
}
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      var data = {
        server_url:basedata.serverUrl,
        api_key:basedata.apiKey,
        id: IDOptions[formRequest.id].id,
        device_id:formRequest.device_id[2],
        request_type:formRequest.request_type,
        content:formRequest.content,
        content_type:formRequest.content_type,
        content_length:formRequest.content_length,
        }
      request('v1/device/queue','POST',data).then((resp)=>{
        addRecord(data.device_id,resp.Paylod,resp.FCnt)
        alert('request successful! FCnt: '+ resp.FCnt + '; bytes: '+resp.Paylod)
        
      },(err)=>{
        alert('request err :'+err)
      })
    } else {
      console.log('error submit!')
      return false
    }
  })
}
interface Record {
  dateTime: string
  deviceID: string
  Payload: string
  FCnt: Number
}


const addRecord = (deviceID,Payload,FCnt)=>{
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
.button_right{
    width: 100%;
    display: flex;
    justify-content: end;
    align-items: end;
}
.top_place{
    margin-top: 50px;
}
</style>
  