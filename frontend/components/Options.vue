<script setup>
    const {options} = defineProps(['options'])
    const focus = ref(false)
    function clickOptions(){
        focus.value = !focus.value
    }
    function blurOptions(e){
        setTimeout(()=>{
            focus.value = false
        }, 200)
    }
</script>

<template>
 <!-- This example requires Tailwind CSS v2.0+ -->
<div class="relative inline-block text-left w-full mb-4">
  <div>
    <button @click="clickOptions" @blur="blurOptions" type="button" class="inline-flex w-full justify-center border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:ring-offset-gray-100" id="menu-button" aria-expanded="true" aria-haspopup="true">
      Opciones
      <!-- Heroicon name: mini/chevron-down -->
      <svg class="-mr-1 ml-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
        <path fill-rule="evenodd" d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z" clip-rule="evenodd" />
      </svg>
    </button>
  </div>

  <!--
    Dropdown menu, show/hide based on menu state.

    Entering: "transition ease-out duration-100"
      From: "transform opacity-0 scale-95"
      To: "transform opacity-100 scale-100"
    Leaving: "transition ease-in duration-75"
      From: "transform opacity-100 scale-100"
      To: "transform opacity-0 scale-95"
  -->
  <div :class="['absolute', 'text-white','z-10','w-full','origin-top-right','divide-y','divide-gray-100',
'bg-gray-800', 'shadow-sm','ring-1','ring-black','ring-opacity-5','focus:outline-none',
!focus ? 'hidden' : 'block']" role="menu" aria-orientation="vertical" aria-labelledby="menu-button" tabindex="-1">
    <div v-for="group in options" class="py-1" role="none">
      <!-- Active: "bg-gray-100 text-gray-900", Not Active: "text-gray-700" -->
      <a v-for="link in group" :href="link.href" class="block px-4 py-2 text-sm hover:bg-gray-500" role="menuitem" tabindex="-1" id="menu-item-0">{{link.text}}</a>
    </div>
  </div>
</div>
</template>