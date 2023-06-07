<script setup>
import Navbar from "./components/Navbar.vue";
</script>

<template>
  <div className="md:w-[1080px] mx-auto py-4">
    <Navbar />
    <router-view></router-view>
  </div>
</template>
