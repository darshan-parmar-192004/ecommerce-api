import { api, checkAuth } from '../api.js'
import { isAuthenticated, getUserName } from '../api.js'

let _welcomeShown = false

export async function renderLoginPage(container, { navigate, showToast }) {
  if (isAuthenticated()) {
    navigate('/')
    return
  }

  container.innerHTML = `
    <div class="min-h-screen flex items-center justify-center pt-20 md:pt-24 bg-gray-50 dark:bg-gray-900">
      <div class="max-w-md w-full mx-auto">
        <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-8 dark:bg-gray-800 dark:border-gray-700">
          <div class="text-center mb-8">
            <div class="w-16 h-16 bg-indigo-100 rounded-full flex items-center justify-center mx-auto mb-4 dark:bg-indigo-900/30">
              <svg class="w-8 h-8 text-indigo-600 dark:text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
              </svg>
            </div>
            <h1 class="text-2xl font-bold text-gray-900 dark:text-white">Welcome Back</h1>
            <p class="text-gray-500 mt-2 dark:text-gray-400">Sign in to your account</p>
          </div>
      
        <form id="login-form" class="space-y-5" novalidate>
          <div class="space-y-1">
            <label for="email" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Email</label>
            <input
              type="email"
              id="email"
              name="email"
              required
              autocomplete="email"
              class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors dark:bg-gray-700 dark:border-gray-600 dark:text-white"
              placeholder="you@example.com"
            />
          </div>
          <div class="space-y-1">
            <label for="password" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              required
              autocomplete="current-password"
              class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors dark:bg-gray-700 dark:border-gray-600 dark:text-white"
              placeholder="Enter your password"
            />
          </div>
          <div id="error" class="hidden bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-xl text-sm flex items-center gap-2 dark:bg-red-900/20 dark:border-red-800 dark:text-red-400" role="alert">
            <svg class="w-5 h-5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <span id="error-text"></span>
          </div>
            <button
              type="submit"
              id="submit-btn"
              class="w-full bg-indigo-600 text-white py-3 rounded-xl hover:bg-indigo-700 active:scale-95 transition-all duration-200 font-semibold focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 flex items-center justify-center gap-2"
            >
              Sign In
            </button>
        </form>
        
        <p class="mt-6 text-center text-gray-600 dark:text-gray-400">
          Don't have an account? 
           <a href="/register" class="text-indigo-600 hover:text-indigo-800 font-semibold transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-indigo-500 rounded px-1 inline-block active:scale-95 dark:text-indigo-400 dark:hover:text-indigo-300" data-route>
            Register
          </a>
        </p>
        <p class="mt-4 text-center">
          <button onclick="window.router.navigate('/products')" class="text-gray-500 hover:text-indigo-600 text-sm transition-all duration-200 inline-flex items-center gap-1 active:scale-95 dark:text-gray-400 dark:hover:text-indigo-400">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
            </svg>
            Continue Browsing
          </button>
        </p>
      </div>
    </div>
  </div>
  `

  const form = document.getElementById('login-form')
  const errorEl = document.getElementById('error')
  const errorText = document.getElementById('error-text')
  const submitBtn = document.getElementById('submit-btn')

  form.addEventListener('submit', async (e) => {
    e.preventDefault()
    const formData = new FormData(form)
    const email = formData.get('email')
    const password = formData.get('password')
    
    if (!email || !password) {
      errorText.textContent = 'Please fill in all fields'
      errorEl.classList.remove('hidden')
      return
    }
    
    submitBtn.disabled = true
    submitBtn.innerHTML = `
      <svg class="animate-spin h-5 w-5 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      Signing in...
    `
    errorEl.classList.add('hidden')

    try {
      const response = await api.auth.login({ email, password })
      
       // Verify session with API
       await checkAuth()
       _welcomeShown = true
       
       window.dispatchEvent(new CustomEvent('auth:change'))
       showToast(`Welcome back, ${getUserName()}!`, 'success')
       navigate('/')
    } catch (err) {
      errorText.textContent = err.message
      errorEl.classList.remove('hidden')
      submitBtn.innerHTML = 'Sign In'
    } finally {
      submitBtn.disabled = false
    }
  })
}

export async function renderRegisterPage(container, { navigate, showToast }) {
  if (isAuthenticated()) {
    navigate('/')
    return
  }

  container.innerHTML = `
    <div class="min-h-screen flex items-center justify-center pt-20 md:pt-24">
      <div class="max-w-md w-full mx-auto">
        <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-8 dark:bg-gray-800 dark:border-gray-700">
          <div class="text-center mb-8">
            <div class="w-16 h-16 bg-indigo-100 rounded-full flex items-center justify-center mx-auto mb-4 dark:bg-indigo-900/30">
              <svg class="w-8 h-8 text-indigo-600 dark:text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"></path>
              </svg>
            </div>
            <h1 class="text-2xl font-bold text-gray-900 dark:text-white">Create Account</h1>
            <p class="text-gray-500 mt-2 dark:text-gray-400">Join us and start shopping</p>
          </div>
        
        <form id="register-form" class="space-y-5" novalidate>
          <div class="space-y-1">
            <label for="name" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Full Name</label>
            <input
              type="text"
              id="name"
              name="name"
              required
              autocomplete="name"
              class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors dark:bg-gray-700 dark:border-gray-600 dark:text-white"
              placeholder="John Doe"
            />
          </div>
          <div class="space-y-1">
            <label for="email" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Email</label>
            <input
              type="email"
              id="email"
              name="email"
              required
              autocomplete="email"
              class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors dark:bg-gray-700 dark:border-gray-600 dark:text-white"
              placeholder="you@example.com"
            />
          </div>
          <div class="space-y-1">
            <label for="password" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              required
              minlength="8"
              autocomplete="new-password"
              class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors dark:bg-gray-700 dark:border-gray-600 dark:text-white"
              placeholder="At least 8 characters"
            />
            <p class="text-xs text-gray-500 mt-1 dark:text-gray-400">Must be at least 8 characters long</p>
          </div>
          <div class="space-y-1">
            <label for="confirm-password" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Confirm Password</label>
            <input
              type="password"
              id="confirm-password"
              name="confirmPassword"
              required
              autocomplete="new-password"
              class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors dark:bg-gray-700 dark:border-gray-600 dark:text-white"
              placeholder="Re-enter your password"
            />
          </div>
          <div id="error" class="hidden bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-xl text-sm flex items-center gap-2 dark:bg-red-900/20 dark:border-red-800 dark:text-red-400" role="alert">
            <svg class="w-5 h-5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <span id="error-text"></span>
          </div>
            <button
              type="submit"
              id="submit-btn"
              class="w-full bg-indigo-600 text-white py-3 rounded-xl hover:bg-indigo-700 active:scale-95 transition-all duration-200 font-semibold focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 flex items-center justify-center gap-2"
            >
              Create Account
            </button>
        </form>
        
        <p class="mt-6 text-center text-gray-600">
          Already have an account? 
           <a href="/login" class="text-indigo-600 hover:text-indigo-800 font-semibold transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-indigo-500 rounded px-1 inline-block active:scale-95" data-route>
            Sign in
          </a>
        </p>
        <p class="mt-4 text-center">
          <button onclick="window.router.navigate('/products')" class="text-gray-500 hover:text-indigo-600 text-sm transition-all duration-200 inline-flex items-center gap-1 active:scale-95">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
            </svg>
            Continue Browsing
          </button>
        </p>
      </div>
    </div>
  `

  const form = document.getElementById('register-form')
  const errorEl = document.getElementById('error')
  const errorText = document.getElementById('error-text')
  const submitBtn = document.getElementById('submit-btn')

  form.addEventListener('submit', async (e) => {
    e.preventDefault()
    const formData = new FormData(form)
    const name = formData.get('name')
    const email = formData.get('email')
    const password = formData.get('password')
    const confirmPassword = formData.get('confirmPassword')
    
    if (!name || !email || !password || !confirmPassword) {
      errorText.textContent = 'Please fill in all fields'
      errorEl.classList.remove('hidden')
      return
    }
    
    if (password.length < 8) {
      errorText.textContent = 'Password must be at least 8 characters'
      errorEl.classList.remove('hidden')
      return
    }

    if (password !== confirmPassword) {
      errorText.textContent = 'Passwords do not match'
      errorEl.classList.remove('hidden')
      return
    }

    submitBtn.disabled = true
    submitBtn.innerHTML = `
      <svg class="animate-spin h-5 w-5 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      Creating account...
    `
    errorEl.classList.add('hidden')

    try {
      const response = await api.auth.register({ name, email, password })
      
       // Verify session with API
       await checkAuth()
       _welcomeShown = true
       
       window.dispatchEvent(new CustomEvent('auth:change'))
       showToast(`Welcome, ${name}! Your account has been created.`, 'success')
       navigate('/')
    } catch (err) {
      errorText.textContent = err.message
      errorEl.classList.remove('hidden')
      submitBtn.innerHTML = 'Create Account'
    } finally {
      submitBtn.disabled = false
    }
  })
}