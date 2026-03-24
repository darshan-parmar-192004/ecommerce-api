<script setup>
import { ref } from "vue";
import ProductList from "./components/ProductList.vue";

const cartItems = ref([]);

const products = [
    {
        id: 1,
        name: "Wireless Headphones",
        price: 79.99,
        category: "Electronics",
        description:
            "Premium noise-canceling wireless headphones with 30-hour battery life and crystal clear audio quality.",
    },
    {
        id: 2,
        name: "Ergonomic Office Chair",
        price: 299.99,
        category: "Furniture",
        description:
            "Adjustable lumbar support, breathable mesh back, and premium foam cushioning for all-day comfort.",
    },
    {
        id: 3,
        name: "Smart Watch Pro",
        price: 249.99,
        category: "Electronics",
        description:
            "Advanced health tracking, GPS, water resistance up to 50m, and seamless notifications integration.",
    },
    {
        id: 4,
        name: "Minimalist Desk Lamp",
        price: 45.99,
        category: "Lighting",
        description:
            "Touch-controlled LED desk lamp with adjustable brightness and color temperature for focused work.",
    },
    {
        id: 5,
        name: "Mechanical Keyboard",
        price: 129.99,
        category: "Electronics",
        description:
            "RGB backlit mechanical keyboard with Cherry MX switches, programmable macros, and aluminum frame.",
    },
    {
        id: 6,
        name: "Standing Desk Converter",
        price: 189.99,
        category: "Furniture",
        description:
            "Height-adjustable workstation that transforms any desk into a sit-stand desk with gas spring lift.",
    },
];

const handleAddToCart = (product) => {
    cartItems.value.push(product);
    console.log("Added to cart:", product.name);
};
</script>

<template>
    <div class="app-container">
        <header>
            <h1>Vue Product Catalog</h1>
            <div class="cart-indicator">Cart: {{ cartItems.length }} items</div>
        </header>

        <main>
            <ProductList :products="products" @add-to-cart="handleAddToCart" />
        </main>
    </div>
</template>

<style>
* {
    box-sizing: border-box;
}

body {
    margin: 0;
    font-family:
        -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu,
        sans-serif;
    background: #f5f7fa;
}

.app-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 24px;
}

header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 32px;
    padding-bottom: 16px;
    border-bottom: 2px solid #e0e0e0;
}

header h1 {
    margin: 0;
    color: #2c5282;
}

.cart-indicator {
    padding: 8px 16px;
    background: #2c5282;
    color: white;
    border-radius: 20px;
    font-size: 14px;
    font-weight: 500;
}
</style>
