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
    <div className="flex justify-center items-center gap-x-4 text-2xl">
      <router-link to="/login">
        <img :src="lock" alt="cart" className="w-[1.9rem] h-[1.9rem]" />
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
};
</script>
