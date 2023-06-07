<script setup>
import { alpaca_logo, cart, lock } from "../assets";
</script>

<template>
  <nav className="flex justify-between items-center pb-5">
    <router-link to="/">
      <div
        className="text-2xl font-semibold flex justify-center items-center gap-x-2"
      >
        <img
          :src="alpaca_logo"
          alt="alpaca_logo"
          className="w-[2.5rem] h-[2.5rem]"
        />
        <span>Room Alpaca Shop</span>
      </div>
    </router-link>
    <div className="flex justify-center items-center gap-x-6 text-2xl">
      <div class="relative">
        <input
          type="text"
          class="w-[13rem] pl-3 pr-[2.5rem] py-3 rounded focus:outline-none bg-gray-100 placeholder:text-base text-base"
          placeholder="ค้นหา"
        />
        <button
          class="block w-7 h-7 text-center text-lg leading-0 absolute top-3 right-2 text-gray-400 focus:outline-none"
          @click="find"
        >
          <font-awesome-icon icon="fa-solid fa-magnifying-glass" />
        </button>
      </div>
      <router-link to="/login">
        <img :src="lock" alt="cart" className="w-[1.8rem] h-[1.8rem]" />
      </router-link>
      <router-link to="/cart">
        <img :src="cart" alt="cart" className="w-[1.8rem] h-[1.8rem]" />
      </router-link>
    </div>
  </nav>
</template>

<script>
export default {
  name: "Navbar",
  methods: {
    find() {
      alert("dummy find");
    },
  },
};
</script>
