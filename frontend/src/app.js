import { checkAuth, api } from './api.js'
import { renderHomePage } from './pages/home.js'
import { renderProductPage } from './pages/product.js'
import { renderCartPage, getCart, saveCart, updateCartCount } from './pages/cart.js'
import { renderLoginPage, renderRegisterPage } from './pages/auth.js'

const routes = {
  '/products': renderHomePage,
  '/login': renderLoginPage,
  '/register': renderRegisterPage,
}

const app = document.getElementById('app')

function parsePathname() {
  const path = window.location.pathname || '/products'
  const queryString = window.location.search || ''
  const params = new URLSearchParams(queryString || '')
  return { path, params }
}

function navigate(path) {
  const overlay = document.getElementById('loading-overlay')
  overlay.classList.remove('hidden')
  setTimeout(() => {
    window.location.href = path
  }, 300)
}

function updateAuthUI() {
  const loggedIn = window._isAuthenticated
  const userName = window._userName || ''
  
  document.getElementById('login-link').classList.toggle('hidden', loggedIn)
  document.getElementById('register-link').classList.toggle('hidden', loggedIn)
  document.getElementById('logout-btn').classList.toggle('hidden', !loggedIn)
  
  document.getElementById('login-link-mobile').classList.toggle('hidden', loggedIn)
  document.getElementById('register-link-mobile').classList.toggle('hidden', loggedIn)
  document.getElementById('logout-btn-mobile').classList.toggle('hidden', !loggedIn)
  
  const profileContainer = document.getElementById('profile-container')
  const profileContainerMobile = document.getElementById('profile-container-mobile')
  
  if (loggedIn && userName) {
    const initials = userName.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
    const profileHTML = `
      <div class="relative group">
        <button class="flex items-center gap-2 text-gray-600 hover:text-indigo-600 font-medium focus:outline-none focus:ring-2 focus:ring-indigo-500 rounded-md px-3 py-2">
          <div class="w-8 h-8 bg-indigo-600 text-white rounded-full flex items-center justify-center text-sm font-semibold">
            ${initials}
          </div>
          <span class="hidden md:inline">${userName.split(' ')[0]}</span>
        </button>
        <div class="absolute right-0 mt-2 w-48 bg-white rounded-lg shadow-lg border border-gray-200 py-2 hidden group-hover:block">
          <button id="logout-btn-dropdown" class="w-full text-left px-4 py-2 text-gray-700 hover:bg-gray-100 hover:text-red-600 transition-colors">
            Logout
          </button>
        </div>
      </div>
    `
    const profileHTMLMobile = `
      <div class="flex items-center gap-3 py-3 px-4">
        <div class="w-10 h-10 bg-indigo-600 text-white rounded-full flex items-center justify-center font-semibold">
          ${initials}
        </div>
        <div>
          <p class="font-medium text-gray-900">${userName}</p>
          <p class="text-sm text-gray-500">Logged in</p>
        </div>
      </div>
      <button id="logout-btn-mobile-profile" class="w-full text-left text-gray-600 hover:text-red-600 hover:bg-red-50 font-medium py-3 px-4 rounded-lg transition-colors">
        Logout
      </button>
    `
    
    if (profileContainer) {
      profileContainer.innerHTML = profileHTML
      document.getElementById('logout-btn-dropdown')?.addEventListener('click', async () => {
        await api.auth.logout()
        window._isAuthenticated = false
        window._userName = ''
        updateAuthUI()
        navigate('/')
        showToast('Logged out successfully', 'success')
      })
    }
    if (profileContainerMobile) {
      profileContainerMobile.innerHTML = profileHTMLMobile
      document.getElementById('logout-btn-mobile-profile')?.addEventListener('click', async () => {
        await api.auth.logout()
        window._isAuthenticated = false
        window._userName = ''
        updateAuthUI()
        closeMobileMenu()
        navigate('/')
        showToast('Logged out successfully', 'success')
      })
    }
  } else {
    if (profileContainer) profileContainer.innerHTML = ''
    if (profileContainerMobile) profileContainerMobile.innerHTML = ''
  }
  
  updateCartCount()
}

function showToast(message, type = 'info') {
  const container = document.getElementById('toast-container')
  const toast = document.createElement('div')
  toast.className = `toast-${type} text-white px-4 py-3 rounded-lg shadow-lg mb-2 transform transition-all duration-300 translate-x-full flex items-center gap-2`
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
  const { path } = parsePathname()
  
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
    await handler(app, { navigate, showToast })
  } else {
    app.innerHTML = `
      <article class="text-center py-16">
        <h1 class="text-6xl font-bold text-gray-900 mb-4">404</h1>
        <p class="text-gray-600 mb-8 text-lg">Page not found</p>
        <a href="/products" class="inline-flex items-center gap-2 bg-indigo-600 text-white px-6 py-3 rounded-lg hover:bg-indigo-700 transition-colors focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
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
  await api.auth.logout()
  updateAuthUI()
  navigate('/')
  showToast('Logged out successfully', 'success')
})

document.getElementById('logout-btn-mobile').addEventListener('click', async () => {
  await api.auth.logout()
  updateAuthUI()
  closeMobileMenu()
  navigate('/')
  showToast('Logged out successfully', 'success')
})

document.querySelectorAll('#mobile-menu a').forEach(link => {
  link.addEventListener('click', closeMobileMenu)
})

window.addEventListener('popstate', router)
let authCheckInProgress = false

window.addEventListener('auth:change', () => {
  updateAuthUI()
})
window.addEventListener('cart:update', updateCartCount)

window.router = { navigate, showToast }

if (window.location.pathname === '/') {
  window.location.href = '/products'
} else {
  document.getElementById('loading-overlay').classList.add('hidden')
  checkAuth().then(() => {
    updateAuthUI()
    router()
  })
}
