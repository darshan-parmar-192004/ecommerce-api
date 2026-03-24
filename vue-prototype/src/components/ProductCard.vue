<script setup>
import { ref, reactive, computed } from "vue";

const props = defineProps({
    product: {
        type: Object,
        required: true,
        validator: (value) => {
            return ["id", "name", "price", "category"].every(
                (key) => key in value,
            );
        },
    },
});

const emit = defineEmits(["add-to-cart"]);

const isHovered = ref(false);

const cart = reactive({
    items: [],
    count: 0,
    total: 0,
});

const formattedPrice = computed(() => {
    return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "INR",
    }).format(props.product.price);
});

const truncatedDescription = computed(() => {
    const maxLength = 80;
    const description = props.product.description || "";

    if (description.length <= maxLength) {
        return description;
    }
    return description.slice(0, maxLength).trim() + "...";
});

const handleAddToCart = () => {
    cart.items.push(props.product);
    cart.count++;
    cart.total += props.product.price;
    emit("add-to-cart", props.product);
};
</script>

<template>
    <div
        class="product-card"
        :class="{ hovered: isHovered }"
        @mouseenter="isHovered = true"
        @mouseleave="isHovered = false"
    >
        <div class="card-header">
            <span class="category">{{ product.category }}</span>
        </div>

        <div class="card-body">
            <h3 class="product-name">{{ product.name }}</h3>
            <p class="description">{{ truncatedDescription }}</p>
            <p class="price">{{ formattedPrice }}</p>
        </div>

        <div class="card-footer">
            <button class="add-to-cart-btn" @click="handleAddToCart">
                Add to Cart
            </button>
        </div>
    </div>
</template>

<style scoped>
.product-card {
    border: 1px solid #e0e0e0;
    border-radius: 8px;
    padding: 16px;
    background: white;
    transition: all 0.2s ease;
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.product-card.hovered {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    transform: translateY(-2px);
}

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.category {
    font-size: 12px;
    text-transform: uppercase;
    color: #666;
    background: #f5f5f5;
    padding: 4px 8px;
    border-radius: 4px;
}

.card-body {
    flex: 1;
}

.product-name {
    margin: 0 0 8px;
    font-size: 18px;
    color: #333;
}

.description {
    margin: 0 0 12px;
    font-size: 14px;
    color: #666;
    line-height: 1.4;
}

.price {
    margin: 0;
    font-size: 20px;
    font-weight: bold;
    color: #2c5282;
}

.card-footer {
    margin-top: auto;
}

.add-to-cart-btn {
    width: 100%;
    padding: 10px;
    background: #2c5282;
    color: white;
    border: none;
    border-radius: 6px;
    cursor: pointer;
    font-size: 14px;
    font-weight: 500;
    transition: background 0.2s;
}

.add-to-cart-btn:hover {
    background: #1a365d;
}
</style>
