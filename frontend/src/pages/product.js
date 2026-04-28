import { api } from "../api.js";

function LoadingSpinner() {
  return `
    <div class="flex justify-center items-center py-16">
      <div class="animate-spin rounded-full h-12 w-12 border-4 border-indigo-200 border-t-indigo-600" role="status" aria-label="Loading product">
        <span class="sr-only">Loading...</span>
      </div>
    </div>
  `;
}

export async function renderProductPage(
  container,
  productId,
  { navigate, showToast, getCart, saveCart },
) {
  container.innerHTML = `
    <div>
       <button onclick="router.navigate('/products')" class="inline-flex items-center gap-2 text-indigo-600 hover:text-indigo-800 font-medium transition-all duration-200 active:scale-95 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 rounded-md px-2 py-1 -ml-2">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
        </svg>
        Back to Products
      </button>

      <div id="loading" aria-live="polite">${LoadingSpinner()}</div>

      <div id="error" class="hidden bg-red-50 border border-red-200 text-red-700 px-6 py-5 rounded-2xl text-center" role="alert">
        <svg class="w-12 h-12 mx-auto mb-3 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path>
        </svg>
        <p id="error-message" class="text-lg font-medium mb-2">Failed to load product</p>
        <p id="error-details"></p>
           <button onclick="router.navigate('/products')" class="inline-flex items-center gap-2 mt-4 text-indigo-600 hover:text-indigo-800 font-medium active:scale-95 transition-all duration-200">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
          </svg>
          Go back to products
        </button>
      </div>

      <article id="product-content" class="hidden">
        <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-6 md:p-10">
          <div class="flex flex-col lg:flex-row gap-8">
            <figure class="lg:w-1/2">
              <div class="bg-gradient-to-br from-gray-100 to-gray-200 rounded-2xl w-full aspect-square flex items-center justify-center">
                <svg class="w-32 h-32 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
              </div>
              <figcaption class="sr-only">Product image</figcaption>
            </figure>
            <div class="lg:w-1/2">
              <div id="product-category" class="mb-3"></div>
              <h1 id="product-name" class="text-3xl lg:text-4xl font-bold text-gray-900 mb-2"></h1>
              <p id="product-price" class="text-4xl lg:text-5xl font-bold text-indigo-600 mb-4"></p>
              <p id="product-description" class="text-gray-600 text-lg leading-relaxed mb-8"></p>

              <div class="flex items-center gap-4 mb-6">
                <label for="quantity" class="text-gray-700 font-medium">Quantity:</label>
                <div class="flex items-center border border-gray-300 rounded-xl overflow-hidden">
                   <button
                     type="button"
                     id="qty-decrease"
                     class="px-4 py-2.5 hover:bg-gray-100 transition-all duration-200 active:scale-95 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500"
                     aria-label="Decrease quantity"
                   >
                    <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4"></path>
                    </svg>
                  </button>
                  <input
                    type="number"
                    id="quantity"
                    value="1"
                    min="1"
                    max="99"
                    class="w-16 px-2 py-2.5 text-center border-x border-gray-300 focus:outline-none focus:ring-0 font-medium"
                    aria-label="Product quantity"
                  />
                   <button
                     type="button"
                     id="qty-increase"
                     class="px-4 py-2.5 hover:bg-gray-100 transition-all duration-200 active:scale-95 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500"
                     aria-label="Increase quantity"
                   >
                    <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
                    </svg>
                  </button>
                </div>
              </div>

              <button
                id="add-to-cart-btn"
                 class="w-full sm:w-auto px-10 py-4 bg-indigo-600 text-white font-semibold rounded-2xl hover:bg-indigo-700 active:scale-95 transition-all duration-200 flex items-center justify-center gap-3 shadow-lg shadow-indigo-500/20"
              >
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"></path>
                </svg>
                Add to Cart
              </button>
            </div>
          </div>
        </div>
      </article>
    </div>
  `;

  const loadingEl = document.getElementById("loading");
  const errorEl = document.getElementById("error");
  const errorMessage = document.getElementById("error-message");
  const errorDetails = document.getElementById("error-details");
  const contentEl = document.getElementById("product-content");
  const nameEl = document.getElementById("product-name");
  const priceEl = document.getElementById("product-price");
  const descEl = document.getElementById("product-description");
  const categoryEl = document.getElementById("product-category");
  const quantityEl = document.getElementById("quantity");
  const addBtn = document.getElementById("add-to-cart-btn");
  const qtyDecrease = document.getElementById("qty-decrease");
  const qtyIncrease = document.getElementById("qty-increase");

  qtyDecrease.addEventListener("click", () => {
    const current = parseInt(quantityEl.value) || 1;
    if (current > 1) quantityEl.value = current - 1;
  });

  qtyIncrease.addEventListener("click", () => {
    const current = parseInt(quantityEl.value) || 1;
    if (current < 99) quantityEl.value = current + 1;
  });

  try {
    const data = await api.products.get(productId);
    const product = data.data || data;

    loadingEl.classList.add("hidden");
    contentEl.classList.remove("hidden");

    nameEl.textContent = product.name;
    priceEl.textContent = `$${product.price?.toFixed(2) || "0.00"}`;
    descEl.textContent = product.description || "No description available.";

    if (product.category_name) {
      categoryEl.innerHTML = `<span class="inline-block px-3 py-1.5 text-sm font-semibold text-indigo-700 bg-indigo-50 rounded-full uppercase tracking-wide">${escapeHtml(product.category_name)}</span>`;
    }

    addBtn.addEventListener("click", () => {
      const quantity = parseInt(quantityEl.value) || 1;
      const cart = getCart();

      const existingItem = cart.find(
        (item) => item.product_id === product.product_id,
      );
      if (existingItem) {
        existingItem.quantity += quantity;
      } else {
        cart.push({
          product_id: product.product_id,
          name: product.name,
          price: product.price,
          quantity,
        });
      }

      saveCart(cart);
      window.dispatchEvent(new CustomEvent("cart:update"));
      showToast(`Added ${quantity} ${product.name} to cart`, "success");
    });
  } catch (err) {
    loadingEl.classList.add("hidden");
    errorMessage.textContent = err.message;
    errorDetails.textContent = "Please try again or go back to products.";
    errorEl.classList.remove("hidden");
  }
}

function escapeHtml(text) {
  const div = document.createElement("div");
  div.textContent = text;
  return div.innerHTML;
}