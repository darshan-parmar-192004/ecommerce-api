import { isAuthenticated, clearAuth } from './api.js'
import { renderHomePage } from './pages/home.js'
import { renderProductPage } from './pages/product.js'
import { renderCartPage, getCart, saveCart, updateCartCount } from './pages/cart.js'
import { renderLoginPage, renderRegisterPage } from './pages/auth.js'

const routes = {
  '/': renderHomePage,
  '/login': renderLoginPage,
  '/register': renderRegisterPage,
}

const app = document.getElementById('app')

function parseHash() {
  const hash = window.location.hash.slice(1) || '/'
  const [path, queryString] = hash.split('?')
  const params = new URLSearchParams(queryString || '')
  return { path, params }
}

function navigate(path) {
  window.location.hash = path
}

function updateAuthUI() {
  const loggedIn = isAuthenticated()
  
  document.getElementById('login-link').classList.toggle('hidden', loggedIn)
  document.getElementById('register-link').classList.toggle('hidden', !loggedIn)
  document.getElementById('logout-btn').classList.toggle('hidden', !loggedIn)
  
  document.getElementById('login-link-mobile').classList.toggle('hidden', loggedIn)
  document.getElementById('register-link-mobile').classList.toggle('hidden', !loggedIn)
  document.getElementById('logout-btn-mobile').classList.toggle('hidden', !loggedIn)
  
  updateCartCount()
}

function showToast(message, type = 'info') {
  const container = document.getElementById('toast-container')
  const toast = document.createElement('div')
  const colors = {
    success: 'bg-green-500',
    error: 'bg-red-500',
    info: 'bg-indigo-600',
    warning: 'bg-yellow-500',
  }
  toast.className = `${colors[type] || colors.info} text-white px-4 py-3 rounded-lg shadow-lg mb-2 transform transition-all duration-300 translate-x-full flex items-center gap-2`
  toast.innerHTML = `
    <svg class="w-5 h-5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
      ${type === 'success' ? '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>' : 
        type === 'error' ? '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>' :
        '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>'}
    </svg>
    <span>${escapeHtml(message)}</span>
  `
  toast.setAttribute('role', 'alert')
  container.appendChild(toast)
  
  requestAnimationFrame(() => {
    toast.classList.remove('translate-x-full')
  })
  
  setTimeout(() => {
    toast.classList.add('translate-x-full')
    setTimeout(() => toast.remove(), 300)
  }, 3000)
}

function escapeHtml(text) {
  const div = document.createElement('div')
  div.textContent = text
  return div.innerHTML
}

async function router() {
  const { path } = parseHash()
  
  if (path.startsWith('/products/')) {
    const id = path.split('/')[2]
    await renderProductPage(app, id, { navigate, showToast, getCart, saveCart })
    return
  }
  
  if (path === '/cart') {
    await renderCartPage(app, { navigate, showToast, getCart, saveCart })
    return
  }
  
  const handler = routes[path]
  if (handler) {
    await handler(app, { navigate, showToast, isAuthenticated })
  } else {
    app.innerHTML = `
      <article class="text-center py-16">
        <h1 class="text-6xl font-bold text-gray-900 mb-4">404</h1>
        <p class="text-gray-600 mb-8 text-lg">Page not found</p>
        <a href="#/" class="inline-flex items-center gap-2 bg-indigo-600 text-white px-6 py-3 rounded-lg hover:bg-indigo-700 transition-colors focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
          </svg>
          Go Home
        </a>
      </article>
    `
  }
}

function toggleMobileMenu() {
  const menu = document.getElementById('mobile-menu')
  const openIcon = document.getElementById('menu-icon-open')
  const closeIcon = document.getElementById('menu-icon-close')
  const btn = document.getElementById('mobile-menu-btn')
  
  const isOpen = !menu.classList.toggle('hidden')
  
  openIcon.classList.toggle('hidden', isOpen)
  closeIcon.classList.toggle('hidden', !isOpen)
  btn.setAttribute('aria-expanded', isOpen)
}

function closeMobileMenu() {
  const menu = document.getElementById('mobile-menu')
  const openIcon = document.getElementById('menu-icon-open')
  const closeIcon = document.getElementById('menu-icon-close')
  const btn = document.getElementById('mobile-menu-btn')
  
  menu.classList.add('hidden')
  openIcon.classList.remove('hidden')
  closeIcon.classList.add('hidden')
  btn.setAttribute('aria-expanded', 'false')
}

document.getElementById('mobile-menu-btn').addEventListener('click', toggleMobileMenu)

document.getElementById('logout-btn').addEventListener('click', async () => {
  const { api } = await import('./api.js')
  await api.auth.logout()
  updateAuthUI()
  navigate('/')
  showToast('Logged out successfully', 'success')
})

document.getElementById('logout-btn-mobile').addEventListener('click', async () => {
  const { api } = await import('./api.js')
  await api.auth.logout()
  updateAuthUI()
  closeMobileMenu()
  navigate('/')
  showToast('Logged out successfully', 'success')
})

document.querySelectorAll('#mobile-menu a').forEach(link => {
  link.addEventListener('click', closeMobileMenu)
})

window.addEventListener('hashchange', router)
window.addEventListener('auth:change', updateAuthUI)
window.addEventListener('cart:update', updateCartCount)

window.router = { navigate, showToast }

updateAuthUI()
router()
