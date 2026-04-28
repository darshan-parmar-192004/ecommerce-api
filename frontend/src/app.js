import { checkAuth, api, isAuthenticated, getUserName, getUserEmail, resetAuthState } from './api.js'
import { renderHomePage } from './pages/home.js'
import { renderProductPage } from './pages/product.js'
import { renderCartPage, getCart, saveCart, updateCartCount } from './pages/cart.js'
import { renderLoginPage, renderRegisterPage } from './pages/auth.js'
import { renderHeader, updateCartBadge } from './components/Header.js'
import { renderFooter } from './components/Footer.js'

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

// Central SPA navigation with history.pushState
function navigate(path, { replace = false } = {}) {
  if (path === window.location.pathname) return
  
  if (replace) {
    window.history.replaceState(null, '', path)
  } else {
    window.history.pushState(null, '', path)
  }
  router()
}

function updateAuthUI() {
  updateCartBadge()
}

function showToast(message, type = 'info') {
  const container = document.getElementById('toast-container')
  if (!container) return
  
  const toast = document.createElement('div')
  const bgClass = type === 'success' ? 'bg-green-600' : type === 'error' ? 'bg-red-600' : 'bg-indigo-600'
  toast.className = `${bgClass} text-white px-4 py-3 rounded-xl shadow-lg mb-2 transform transition-all duration-300 translate-x-full flex items-center gap-2 active:scale-95`
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

// Global keyboard shortcuts
function setupKeyboardShortcuts() {
  document.addEventListener('keydown', (e) => {
    // Ctrl+K or Cmd+K to focus search
    if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
      e.preventDefault()
      const searchInput = document.getElementById('header-search-input') || 
                         document.getElementById('search-input')
      if (searchInput) {
        searchInput.focus()
        searchInput.select()
      }
    }
    
    // Escape to close dropdowns
    if (e.key === 'Escape') {
      const dropdown = document.getElementById('user-dropdown-menu')
      const mobileMenu = document.getElementById('mobile-menu')
      const avatarBtn = document.getElementById('user-avatar-btn')
      
      if (dropdown && !dropdown.classList.contains('hidden')) {
        dropdown.classList.add('hidden')
        if (avatarBtn) avatarBtn.setAttribute('aria-expanded', 'false')
      }
      
      if (mobileMenu && !mobileMenu.classList.contains('hidden')) {
        mobileMenu.classList.add('hidden')
        const openIcon = document.getElementById('menu-icon-open')
        const closeIcon = document.getElementById('menu-icon-close')
        const btn = document.getElementById('mobile-menu-btn')
        if (openIcon && closeIcon && btn) {
          openIcon.classList.remove('hidden')
          closeIcon.classList.add('hidden')
          btn.setAttribute('aria-expanded', 'false')
        }
      }
    }
  })
}

// Event delegation for data-route links
function setupRouteDelegation() {
  document.addEventListener('click', (e) => {
    const link = e.target.closest('[data-route]')
    if (link && link.tagName === 'A') {
      const href = link.getAttribute('href')
      if (href && href.startsWith('/')) {
        e.preventDefault()
        navigate(href)
      }
    }
  })
}

// Main router
async function router() {
  const { path } = parsePathname()
  
  // Determine header appearance based on path
  const isAuthPage = path === '/login' || path === '/register'
  const isTransparent = false // Always show opaque header
  
  // Render header and footer
  renderHeader({ transparent: isTransparent, hidden: isAuthPage })
  
  // Only show footer on non-auth pages
  if (!isAuthPage) {
    renderFooter()
  } else {
    // Clear footer on auth pages
    const footerRoot = document.getElementById('footer-root')
    if (footerRoot) footerRoot.innerHTML = ''
  }
  
  // Route handling
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
        <button onclick="window.router.navigate('/products')" class="inline-flex items-center gap-2 bg-indigo-600 text-white px-6 py-3 rounded-2xl hover:bg-indigo-700 active:scale-95 transition-all duration-150 font-semibold focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
          </svg>
          Go Home
        </button>
      </article>
    `
  }
}

// Global logout click handler via event delegation (fallback)
document.addEventListener('click', async (e) => {
  if (e.target.id === 'logout-btn' || e.target.id === 'logout-btn-mobile') {
    e.preventDefault()
    await api.auth.logout()
    resetAuthState()
    updateAuthUI()
    navigate('/')
    showToast('Logged out successfully', 'success')
  }
})

// Event listeners
window.addEventListener('popstate', router)
window.addEventListener('auth:change', () => {
  updateAuthUI()
  router()
})
window.addEventListener('cart:update', updateCartBadge)

// Re-render page on theme change
window.addEventListener('theme:change', () => {
  router()
})

// Global exports
window.router = { navigate, showToast }
window.Auth = { updateAuthUI }

// Initialize theme from localStorage or system preference
function initTheme() {
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme === 'dark' || (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
}
initTheme()

// Initialize
setupKeyboardShortcuts()
setupRouteDelegation()

if (window.location.pathname === '/' && window.location.hash !== '#/') {
  navigate('/products', { replace: true })
} else {
  document.getElementById('loading-overlay').classList.add('hidden')
  checkAuth().then(() => {
    updateAuthUI()
    router()
  })
}
