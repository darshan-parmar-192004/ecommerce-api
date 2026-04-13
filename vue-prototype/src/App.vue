<script setup>
import { reactive, provide, computed } from "vue";
import ProductList from "./components/ProductList.vue";
import CartSummary from "./components/CartSummary.vue";

const cart = reactive({
    items: [],
    total: 0,
    addItem(product) {
        this.items.push({ ...product, cartId: Date.now() + Math.random() });
        this.total += product.price;
    },
    removeItem(productId) {
        const index = this.items.findIndex(
            (item) => item.id === productId
        );
        if (index !== -1) {
            this.total -= this.items[index].price;
            this.items.splice(index, 1);
        }
    },
    clearCart() {
        this.items = [];
        this.total = 0;
    }
});

provide("cart", cart);
</script>

<template>
    <div class="min-h-screen bg-gray-50">
        <header class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-50">
            <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div class="flex items-center justify-between h-16">
                    <div class="flex items-center gap-3">
                        <div class="w-9 h-9 bg-gradient-to-br from-blue-500 to-blue-700 rounded-lg flex items-center justify-center shadow-md">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                            </svg>
                        </div>
                        <h1 class="text-xl font-bold text-gray-900 hidden sm:block">ShopVue</h1>
                    </div>
                    
                    <nav class="hidden md:flex items-center gap-6">
                        <a href="#" class="text-sm font-medium text-gray-900 hover:text-blue-600 transition-colors">Home</a>
                        <a href="#" class="text-sm font-medium text-gray-500 hover:text-blue-600 transition-colors">Products</a>
                        <a href="#" class="text-sm font-medium text-gray-500 hover:text-blue-600 transition-colors">About</a>
                        <a href="#" class="text-sm font-medium text-gray-500 hover:text-blue-600 transition-colors">Contact</a>
                    </nav>

                    <div class="flex items-center gap-3">
                        <button class="p-2 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition-colors">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                            </svg>
                        </button>
                        <button class="p-2 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition-colors relative">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                            </svg>
                            <span v-if="cart.items.length > 0" class="absolute -top-1 -right-1 w-5 h-5 bg-blue-600 text-white text-xs font-bold rounded-full flex items-center justify-center">
                                {{ cart.items.length }}
                            </span>
                        </button>
                    </div>
                </div>
            </div>
        </header>

        <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
            <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
                <div class="lg:col-span-3">
                    <ProductList />
                </div>
                <aside class="lg:col-span-1">
                    <div class="sticky top-24">
                        <CartSummary />
                    </div>
                </aside>
            </div>
        </main>

        <footer class="bg-white border-t border-gray-200 mt-auto">
            <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
                <div class="flex flex-col md:flex-row items-center justify-between gap-4">
                    <div class="flex items-center gap-2">
                        <div class="w-6 h-6 bg-gradient-to-br from-blue-500 to-blue-700 rounded flex items-center justify-center">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                            </svg>
                        </div>
                        <span class="text-sm font-medium text-gray-700">ShopVue</span>
                    </div>
                    <p class="text-sm text-gray-500">© 2024 ShopVue. Built with Vue 3 + Tailwind CSS</p>
                </div>
            </div>
        </footer>
    </div>
</template>
