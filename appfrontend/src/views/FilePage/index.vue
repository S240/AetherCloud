<template>
  <div class="navbar bg-base-100 shadow-sm">
    <div class="navbar-start">
      <div class="dropdown">
        <div tabindex="0" role="button" class="btn btn-ghost btn-circle">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-7 w-7" fill="none" viewBox="0 0 24 24" stroke="currentColor"> <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7" /> </svg>
        </div>
        <ul
          tabindex="0"
          class="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow">
          <li><a class="text-xl">Homepage</a></li>
          <li><a class="text-xl">Portfolio</a></li>
          <li><a class="text-xl">About</a></li>
        </ul>
      </div>
    </div>
    <div class="navbar-center">
      <a class="btn btn-ghost text-xl">CloudStore</a>
    </div>
    <div class="navbar-end">
      <button class="btn btn-ghost btn-circle">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"> <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /> </svg>
      </button>
      <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
          <div class="w-10 rounded-full">
            <img
              alt="Tailwind CSS Navbar component"
              src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
          </div>
      </div>
    </div>
  
  </div>

  <div class="divider"></div>

  <div class="flex">
    <div class="w-16  ..."></div>
    <div class="flex-auto  ...">

      <h1 class="text-2xl font-bold">全部文件</h1>

    <ul class="list bg-base-100 rounded-box shadow-md">

    </ul>


    </div>


    <div class="w-96  ..." >      
      <button class="btn  btn-info"  @click="uploadFile()">
          <PlusOutlined />
      </button>
      <div v-if="showProgressModal" tabindex="0" class="collapse collapse-arrow bg-base-100 border-base-300 border">
        <div class="collapse-title semibold">
          <div class="flex">
            <div class="w-16  ..." >
              <button class="btn" @click="cancelupload"> <PlusOutlined /></button>
            </div>
            <p>正在上传  -  {{ TaskName }}</p>
          </div>
        </div>
        <div class="collapse-content">
          <Flex gap="small" wrap>
            <a-progress type="circle" trailColor="white" :percent="md5Progress.toFixed(0)" />
          </Flex>
        </div>
      </div>
    </div>
  </div>



</template>


<script setup lang="ts">
import { Button } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import { ref } from 'vue';
import SparkMD5 from 'spark-md5';
import { request } from "/@/utils/service";

const uploadstate = ref(false);
const cancelupload = () =>{
  if (!uploadstate.value){
    return
  }
  uploadstate.value = false;
  showProgressModal.value = false;
}

const uploadFile = async() => {
  if (window.showOpenFilePicker) {
    try {
      const [fileHandle] = await window.showOpenFilePicker();
      const file = await fileHandle.getFile();
      TaskName.value = file.name
      uploadstate.value = true;
      showProgressModal.value = true
      const md5 = await calculateFileMD5(file, (progress: number) => {
        md5Progress.value = progress;
      });
      await MinIOFileUpload(file.name,md5,file.size)

      md5Progress.value = 0 
      uploadstate.value = false;
      showProgressModal.value = false;
    } catch (error) {
      console.error('选择文件过程出错或被取消', error);
    }
  } else {
    console.warn('当前浏览器不支持 File System Access API');
  }
}
const md5Progress = ref(0);
const TaskName = ref('')
const showProgressModal = ref(false);
async function calculateFileMD5(file: File, onProgress?: (progress: number) => void): Promise<string> {
  return new Promise((resolve, reject) => {
    const chunkSize = 2 * 1024 * 1024; 
    const chunks = Math.ceil(file.size / chunkSize);
    let currentChunk = 0;
    const spark = new SparkMD5.ArrayBuffer();
    const fileReader = new FileReader();

    fileReader.onload = (e: ProgressEvent<FileReader>) => {
      if (!uploadstate.value) {
        reject(new Error('上传已取消'));
        return;
      }
      if (e.target?.result) {
        spark.append(e.target.result as ArrayBuffer);
        currentChunk++;
        if (onProgress) {
          onProgress((currentChunk / chunks) * 100);
        }
        if (currentChunk < chunks) {
          loadNext();
        } else {
          resolve(spark.end());
        }
      } else {
        reject(new Error('读取文件出错'));
      }
    };

    fileReader.onerror = () => {
      reject(new Error('读取文件出错'));
    };

    function loadNext() {
      const start = currentChunk * chunkSize;
      const end = Math.min(file.size, start + chunkSize);
      const slice = file.slice(start, end);
      fileReader.readAsArrayBuffer(slice);
    }

    loadNext();
  });
}

const MinIOFileUpload = async(filename :string,md5: string, fileSize: number) =>{
  try {
        const response=await request({
            url: '/file/getuploadurl',
            method: 'post',
            data:{
              file_md5:md5,
              file_name: filename,        
              file_size: fileSize,
              is_chunked: true,   
              chunk_size: 10*1024*1024
            }
        })
        const data = await response.data;
    } catch (error) {
        
    }
}
</script>

<style scoped>

</style>