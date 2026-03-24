<script setup>
import { ref, computed, inject } from "vue";

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

const cart = inject("cart");
const isHovered = ref(false);

const formattedPrice = computed(() => {
    return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "USD",
    }).format(props.product.price);
});

const truncatedDescription = computed(() => {
    const maxLength = 100;
    const description = props.product.description || "";

    if (description.length <= maxLength) {
        return description;
    }
    return description.slice(0, maxLength).trim() + "...";
});

const handleAddToCart = () => {
    cart.addItem(props.product);
};

const getCategoryColor = (category) => {
    const colors = {
        Electronics: 'bg-blue-100 text-blue-700',
        Furniture: 'bg-amber-100 text-amber-700',
        Lighting: 'bg-purple-100 text-purple-700',
    };
    return colors[category] || 'bg-gray-100 text-gray-700';
};
</script>

<template>
    <div
        class="bg-white rounded-xl border border-gray-200 shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all duration-300 overflow-hidden group"
        :class="{ 'shadow-xl': isHovered }"
        @mouseenter="isHovered = true"
        @mouseleave="isHovered = false"
    >
        <div class="h-1.5 bg-gradient-to-r from-blue-500 via-blue-600 to-blue-700"></div>
        
        <div class="p-5">
            <div class="flex justify-between items-start mb-3">
                <span 
                    class="text-xs font-semibold uppercase tracking-wider px-2.5 py-1 rounded-full"
                    :class="getCategoryColor(product.category)"
                >
                    {{ product.category }}
                </span>
            </div>

            <div class="mb-3">
                <h3 class="text-lg font-bold text-gray-900 group-hover:text-blue-700 transition-colors duration-200">
                    {{ product.name }}
                </h3>
            </div>

            <p class="text-sm text-gray-500 leading-relaxed mb-4 line-clamp-2">
                {{ truncatedDescription }}
            </p>

            <div class="flex items-center justify-between pt-3 border-t border-gray-100">
                <p class="text-2xl font-extrabold text-gray-900">
                    {{ formattedPrice }}
                </p>

                <button
                    class="inline-flex items-center gap-2 px-4 py-2.5 bg-gray-900 hover:bg-blue-600 text-white font-medium rounded-lg transition-all duration-200 hover:shadow-lg active:scale-95"
                    @click="handleAddToCart"
                >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
                    </svg>
                    <span class="hidden sm:inline">Add</span>
                </button>
            </div>
        </div>
    </div>
</template>
