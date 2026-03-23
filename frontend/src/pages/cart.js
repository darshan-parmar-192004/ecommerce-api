import { api } from '../api.js'

const CART_KEY = 'ecommerce_cart'

export function getCart() {
  try {
    return JSON.parse(localStorage.getItem(CART_KEY)) || []
  } catch {
    return []
  }
}

export function saveCart(cart) {
  localStorage.setItem(CART_KEY, JSON.stringify(cart))
}

export function updateCartCount() {
  const cart = getCart()
  const count = cart.reduce((sum, item) => sum + item.quantity, 0)
  const badge = document.getElementById('cart-count')
  const mobileBadge = document.getElementById('cart-count-mobile')
  
  if (count > 0) {
    badge.textContent = count
    badge.classList.remove('hidden')
    if (mobileBadge) {
      mobileBadge.textContent = count
      mobileBadge.classList.remove('hidden')
    }
  } else {
    badge.classList.add('hidden')
    if (mobileBadge) {
      mobileBadge.classList.add('hidden')
    }
  }
}

export async function renderCartPage(container, { navigate, showToast, getCart, saveCart }) {
  const cart = getCart()
  
  if (cart.length === 0) {
    container.innerHTML = `
      <article class="text-center py-16">
        <div class="w-24 h-24 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-6">
          <svg class="w-12 h-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"></path>
          </svg>
        </div>
        <h1 class="text-2xl font-bold text-gray-900 mb-3">Your cart is empty</h1>
        <p class="text-gray-500 mb-8">Looks like you haven't added anything to your cart yet.</p>
        <a 
          href="#/" 
          class="inline-flex items-center gap-2 bg-indigo-600 text-white px-6 py-3 rounded-lg hover:bg-indigo-700 active:bg-indigo-800 font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"></path>
          </svg>
          Browse Products
        </a>
      </article>
    `
    return
  }

  const total = cart.reduce((sum, item) => sum + (item.price * item.quantity), 0)

  container.innerHTML = `
    <section aria-labelledby="cart-heading">
      <h1 id="cart-heading" class="text-3xl font-bold text-gray-900 mb-8">Shopping Cart</h1>
      
      <div class="grid lg:grid-cols-3 gap-8">
        <div class="lg:col-span-2">
          <div class="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
            <ul id="cart-items" class="divide-y divide-gray-100" role="list">
              ${cart.map((item, index) => `
                <li class="flex items-center justify-between p-5 hover:bg-gray-50 transition-colors" data-index="${index}">
                  <div class="flex items-center gap-4">
                    <div class="bg-gray-200 border-2 border-dashed rounded-lg w-20 h-20 flex items-center justify-center flex-shrink-0" aria-label="Product image placeholder">
                      <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                      </svg>
                    </div>
                    <div>
                      <h3 class="font-semibold text-gray-900">${escapeHtml(item.name)}</h3>
                      <p class="text-gray-500 text-sm mt-1">$${item.price?.toFixed(2) || '0.00'} each</p>
                    </div>
                  </div>
                  <div class="flex items-center gap-4">
                    <div class="flex items-center border border-gray-300 rounded-lg overflow-hidden">
                      <button 
                        class="qty-btn px-3 py-2 hover:bg-gray-100 transition-colors focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500" 
                        data-action="decrease" 
                        data-index="${index}"
                        aria-label="Decrease quantity"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4"></path>
                        </svg>
                      </button>
                      <span class="px-4 py-2 font-medium border-x border-gray-300 min-w-[3rem] text-center" aria-label="Quantity: ${item.quantity}">
                        ${item.quantity}
                      </span>
                      <button 
                        class="qty-btn px-3 py-2 hover:bg-gray-100 transition-colors focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500" 
                        data-action="increase" 
                        data-index="${index}"
                        aria-label="Increase quantity"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
                        </svg>
                      </button>
                    </div>
                    <p class="font-bold text-lg w-24 text-right" aria-label="Item total: $${((item.price || 0) * item.quantity).toFixed(2)}">
                      $${((item.price || 0) * item.quantity).toFixed(2)}
                    </p>
                    <button 
                      class="remove-btn text-gray-400 hover:text-red-500 p-2 transition-colors focus:outline-none focus:ring-2 focus:ring-red-500 rounded-lg"
                      data-index="${index}"
                      aria-label="Remove ${escapeHtml(item.name)} from cart"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                      </svg>
                    </button>
                  </div>
                </li>
              `).join('')}
            </ul>
            <div class="p-5 border-t border-gray-100">
              <button 
                id="clear-cart-btn" 
                class="text-red-500 hover:text-red-700 font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-red-500 rounded px-2 py-1"
              >
                Clear Cart
              </button>
            </div>
          </div>
        </div>
        
        <div class="lg:col-span-1">
          <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-6 sticky top-24">
            <h2 class="text-lg font-bold text-gray-900 mb-4">Order Summary</h2>
            <dl class="space-y-3 text-sm">
              <div class="flex justify-between">
                <dt class="text-gray-600">Subtotal</dt>
                <dd class="font-medium">$${total.toFixed(2)}</dd>
              </div>
              <div class="flex justify-between">
                <dt class="text-gray-600">Shipping</dt>
                <dd class="font-medium text-green-600">Free</dd>
              </div>
              <div class="border-t border-gray-200 pt-3 flex justify-between">
                <dt class="font-bold text-gray-900">Total</dt>
                <dd class="font-bold text-xl text-indigo-600">$${total.toFixed(2)}</dd>
              </div>
            </dl>
            <button
              id="checkout-btn"
              class="w-full mt-6 bg-indigo-600 text-white py-3.5 rounded-lg hover:bg-indigo-700 active:bg-indigo-800 font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 flex items-center justify-center gap-2"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
              Proceed to Checkout
            </button>
          </div>
        </div>
      </div>
    </section>
  `

  document.querySelectorAll('.qty-btn').forEach(btn => {
    btn.addEventListener('click', (e) => {
      const index = parseInt(e.currentTarget.dataset.index)
      const action = e.currentTarget.dataset.action
      const cart = getCart()
      
      if (action === 'increase') {
        cart[index].quantity++
      } else if (action === 'decrease') {
        if (cart[index].quantity > 1) {
          cart[index].quantity--
        } else {
          cart.splice(index, 1)
        }
      }
      
      saveCart(cart)
      window.dispatchEvent(new CustomEvent('cart:update'))
      renderCartPage(container, { navigate, showToast, getCart, saveCart })
    })
  })

  document.querySelectorAll('.remove-btn').forEach(btn => {
    btn.addEventListener('click', (e) => {
      const index = parseInt(e.currentTarget.dataset.index)
      const cart = getCart()
      const itemName = cart[index].name
      cart.splice(index, 1)
      saveCart(cart)
      window.dispatchEvent(new CustomEvent('cart:update'))
      renderCartPage(container, { navigate, showToast, getCart, saveCart })
      showToast(`${itemName} removed from cart`, 'info')
    })
  })

  document.getElementById('clear-cart-btn').addEventListener('click', () => {
    saveCart([])
    window.dispatchEvent(new CustomEvent('cart:update'))
    renderCartPage(container, { navigate, showToast, getCart, saveCart })
    showToast('Cart cleared', 'info')
  })

  document.getElementById('checkout-btn').addEventListener('click', async () => {
    showToast('Checkout functionality coming soon!', 'info')
  })
}

function escapeHtml(text) {
  const div = document.createElement('div')
  div.textContent = text
  return div.innerHTML
}
