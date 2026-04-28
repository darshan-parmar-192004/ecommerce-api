let headerEventsInitialized = false;
let searchTimeout = null;

export function renderHeader({ transparent = false, hidden = false } = {}) {
  const headerRoot = document.getElementById('header-root');
  if (!headerRoot) return;

  const isLoggedIn = window._isAuthenticated || false;
  const userName = window._userName || '';
  const userEmail = window._userEmail || '';

  const isDark = document.documentElement.classList.contains('dark');

  headerRoot.innerHTML = `
    <header class="fixed top-0 left-0 right-0 z-50 transition-all duration-300 ${transparent ? 'bg-transparent border-transparent' : 'bg-white/70 backdrop-blur-md border-gray-200/50 shadow-sm border-b dark:bg-gray-900/70 dark:border-gray-700/50'}" id="main-header">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16 md:h-20">
          <!-- Logo -->
          <a href="/products" class="flex items-center gap-2.5 group active:scale-95 transition-all duration-200" data-route>
            <div class="w-9 h-9 rounded-xl bg-gradient-to-br from-indigo-500 to-indigo-700 flex items-center justify-center shadow-lg shadow-indigo-500/20 group-hover:shadow-lg group-hover:shadow-indigo-500/30 transition-all duration-300 active:scale-95">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"/></svg>
              </div>
              <span class="text-xl font-bold tracking-tight text-indigo-600 hover:text-indigo-700 dark:text-white dark:hover:text-indigo-400 transition-colors duration-300">
                LuxeCart
              </span>
            </a>
           
          <!-- Desktop Navigation -->
          <nav class="hidden md:flex items-center gap-1" role="navigation" aria-label="Main navigation">
            <a href="/products" class="nav-link px-4 py-2 rounded-xl text-sm font-medium transition-all duration-200 active:scale-95 flex items-center gap-2 text-gray-600 hover:text-indigo-600 hover:bg-indigo-50 dark:text-gray-300 dark:hover:text-indigo-400 dark:hover:bg-gray-800" data-route>
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"/></svg>
              Products
            </a>
           
            <!-- Search Bar -->
            <div class="relative group ml-2">
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-gray-400 group-focus-within:text-indigo-500 dark:text-gray-500 dark:group-focus-within:text-indigo-400 transition-colors duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
                </svg>
              </div>
              <input
                type="text"
                id="header-search-input"
                placeholder="Search products..."
                class="w-56 lg:w-72 pl-11 pr-12 py-2.5 bg-white border-gray-200 text-gray-900 placeholder-gray-400 focus:bg-white rounded-xl text-sm border transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent shadow-sm hover:shadow-md active:scale-95 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-100 dark:placeholder-gray-500 dark:focus:bg-gray-800"
                aria-label="Search products"
              />
              <kbd class="absolute inset-y-0 right-2 top-1/2 -translate-y-1/2 hidden sm:flex items-center gap-1 bg-gray-100 text-gray-400 border border-gray-200 rounded px-2 text-[10px] font-semibold uppercase tracking-wider dark:bg-gray-700 dark:text-gray-400 dark:border-gray-600" aria-hidden="true">
                <span class="text-[9px]">⌘</span>K
              </kbd>
              <div id="search-results-dropdown" class="hidden absolute top-full left-0 right-0 mt-2 bg-white rounded-2xl shadow-2xl border border-gray-100 overflow-hidden z-50 max-h-96 overflow-y-auto dark:bg-gray-800 dark:border-gray-700"></div>
            </div>
           
            <!-- Shopping Bag Icon with Badge -->
            <a href="/cart" class="cart-link relative px-4 py-2 rounded-xl text-sm font-medium transition-all duration-200 active:scale-95 flex items-center gap-2 text-gray-600 hover:text-indigo-600 hover:bg-indigo-50 dark:text-gray-300 dark:hover:text-indigo-400 dark:hover:bg-gray-800" data-route aria-label="Shopping cart">
              <div class="relative">
                <svg class="w-5 h-5 transition-transform duration-200 group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"/>
                </svg>
                <span id="header-cart-badge" class="cart-badge absolute -top-1 -right-1 flex h-4 w-4 items-center justify-center rounded-full bg-red-500 text-[10px] text-white font-bold hidden scale-0" aria-label="Cart items count">0</span>
              </div>
              <span class="hidden sm:inline">Bag</span>
            </a>
          </nav>
           
          <!-- Theme Toggle & Auth Section -->
          <div class="flex items-center gap-2 ml-4" id="header-auth-section">
            <!-- Theme Toggle Button -->
            <button 
              id="theme-toggle-btn"
              class="p-2 rounded-xl text-gray-600 hover:text-indigo-600 hover:bg-indigo-50 dark:text-gray-300 dark:hover:text-indigo-400 dark:hover:bg-gray-800 transition-all duration-200 active:scale-95"
              aria-label="Toggle theme"
            >
              <!-- Sun icon (shown in dark mode) -->
              <svg id="theme-icon-sun" class="w-5 h-5 ${isDark ? '' : 'hidden'}" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"/>
              </svg>
              <!-- Moon icon (shown in light mode) -->
              <svg id="theme-icon-moon" class="w-5 h-5 ${isDark ? 'hidden' : ''}" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9 9 0 008.354-5.646z"/>
              </svg>
            </button>

            ${isLoggedIn ? `
              <!-- User Avatar -->
              <div class="relative" id="user-dropdown-container">
                <button 
                  class="flex items-center gap-2 hover:opacity-80 transition-opacity focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 rounded-full p-1 active:scale-95 transition-all duration-200" 
                  aria-label="User menu" 
                  aria-expanded="false" 
                  aria-haspopup="true"
                  id="user-avatar-btn"
                >
                  <div class="w-9 h-9 rounded-full flex items-center justify-center text-white text-sm font-bold shadow-lg transition-transform hover:scale-105 duration-200 bg-gradient-to-br from-indigo-500 to-indigo-700">
                    ${getAvatarInitials(userName)}
                  </div>
                </button>
                
                <!-- Dropdown Menu -->
                <div 
                  class="hidden absolute right-0 top-full mt-2 w-48 bg-white rounded-md shadow-lg py-1 z-50 dark:bg-gray-800 dark:border dark:border-gray-700"
                  id="user-dropdown-menu"
                  role="menu"
                  aria-label="User menu"
                >
                  <div class="px-4 py-3 mb-2 border-b border-gray-100 dark:border-gray-700">
                    <p class="font-semibold text-gray-900 text-sm truncate dark:text-white" id="dropdown-user-name">${escapeHtml(userName)}</p>
                    <p class="text-xs text-gray-500 truncate dark:text-gray-400" id="dropdown-user-email">${escapeHtml(userEmail)}</p>
                  </div>
                  
                  <button 
                    class="dropdown-item w-full text-left px-4 py-2.5 text-sm text-gray-700 hover:bg-indigo-50 hover:text-indigo-700 transition-all duration-200 active:scale-95 flex items-center gap-2 rounded-md focus:outline-none focus:bg-indigo-50 group dark:text-gray-300 dark:hover:bg-gray-700 dark:hover:text-indigo-400"
                    role="menuitem"
                    data-route="/profile"
                  >
                    <svg class="w-4 h-4 text-gray-400 group-hover:text-indigo-600 transition-colors dark:text-gray-500 dark:group-hover:text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                    My Profile
                  </button>
                  
                  <button 
                    class="dropdown-item w-full text-left px-4 py-2.5 text-sm text-gray-700 hover:bg-indigo-50 hover:text-indigo-700 transition-all duration-200 active:scale-95 flex items-center gap-2 rounded-md focus:outline-none focus:bg-indigo-50 group dark:text-gray-300 dark:hover:bg-gray-700 dark:hover:text-indigo-400"
                    role="menuitem"
                    data-route="/orders"
                  >
                    <svg class="w-4 h-4 text-gray-400 group-hover:text-indigo-600 transition-colors dark:text-gray-500 dark:group-hover:text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"/></svg>
                    Order History
                  </button>
                  
                  <div class="my-2 border-t border-gray-100 dark:border-gray-700"></div>
                  
                  <button 
                    class="dropdown-item w-full text-left px-4 py-2.5 text-sm text-red-600 hover:bg-red-50 transition-all duration-200 active:scale-95 flex items-center gap-2 rounded-md focus:outline-none focus:bg-red-50 font-medium group dark:hover:bg-red-900/20"
                    role="menuitem"
                    data-action="logout"
                  >
                    <svg class="w-4 h-4 group-hover:text-red-700 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/></svg>
                    Sign Out
                  </button>
                </div>
              </div>
            ` : `
              <!-- Sign In Pill Button -->
              <a href="/login" 
                 class="auth-signin-btn inline-flex items-center gap-2 bg-indigo-600 hover:bg-indigo-700 text-white px-5 py-2 rounded-full text-sm font-semibold transition-all duration-200 active:scale-95 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 shadow-lg shadow-indigo-500/20" 
                 data-route>
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                </svg>
                Sign In
              </a>
            `}
          </div>
           
          <!-- Mobile Menu Button -->
          <button 
            class="md:hidden p-2 rounded-xl transition-all duration-200 active:scale-95 text-gray-600 hover:text-indigo-600 hover:bg-indigo-50 dark:text-gray-300 dark:hover:text-indigo-400 dark:hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-indigo-500" 
            id="mobile-menu-btn"
            aria-label="Toggle menu"
            aria-expanded="false"
            aria-controls="mobile-menu"
          >
            <svg id="menu-icon-open" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/></svg>
            <svg id="menu-icon-close" class="w-6 h-6 hidden" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/></svg>
          </button>
        </div>
        
        <!-- Mobile Menu -->
        <div 
          class="hidden md:hidden absolute top-full left-0 right-0 bg-white/95 backdrop-blur-md border-b border-gray-100 shadow-2xl animate-slide-down dark:bg-gray-900/95 dark:border-gray-700"
          id="mobile-menu"
          role="menu"
        >
          <div class="p-4 space-y-3">
            <div class="relative mb-3">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-gray-400 dark:text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
                </svg>
              </div>
              <input
                type="text"
                class="w-full pl-10 pr-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all duration-200 active:scale-95 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-100"
                placeholder="Search products..."
                aria-label="Search products"
              />
            </div>
            
            <a href="/products" class="mobile-nav-link flex items-center gap-3 px-4 py-3 text-gray-700 hover:bg-indigo-50 hover:text-indigo-600 rounded-xl transition-all duration-200 active:scale-95 font-medium dark:text-gray-300 dark:hover:bg-gray-800 dark:hover:text-indigo-400" role="menuitem" data-route>
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"/></svg>
              Products
            </a>
            
            <a href="/cart" class="mobile-nav-link flex items-center gap-3 px-4 py-3 text-gray-700 hover:bg-indigo-50 hover:text-indigo-600 rounded-xl transition-all duration-200 active:scale-95 font-medium dark:text-gray-300 dark:hover:bg-gray-800 dark:hover:text-indigo-400" role="menuitem" data-route>
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"/>
              </svg>
              Cart
              <span id="mobile-cart-badge" class="ml-auto bg-red-500 text-white text-[10px] font-bold px-2 py-0.5 rounded-full hidden">0</span>
            </a>

            <!-- Mobile Theme Toggle -->
            <button 
              id="mobile-theme-toggle-btn"
              class="w-full flex items-center gap-3 px-4 py-3 text-gray-700 hover:bg-indigo-50 hover:text-indigo-600 rounded-xl transition-all duration-200 active:scale-95 font-medium dark:text-gray-300 dark:hover:bg-gray-800 dark:hover:text-indigo-400"
              role="menuitem"
            >
              <svg id="mobile-theme-icon-sun" class="w-5 h-5 ${isDark ? '' : 'hidden'}" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"/>
              </svg>
              <svg id="mobile-theme-icon-moon" class="w-5 h-5 ${isDark ? 'hidden' : ''}" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9 9 0 008.354-5.646z"/>
              </svg>
              ${isDark ? 'Light Mode' : 'Dark Mode'}
            </button>
            
            <div id="mobile-auth-section" class="pt-3 mt-3 border-t border-gray-100 dark:border-gray-700">
              ${isLoggedIn ? `
                <div class="flex items-center gap-3 px-4 py-3">
                  <div class="w-10 h-10 rounded-full flex items-center justify-center text-white text-sm font-bold bg-gradient-to-br from-indigo-500 to-indigo-700 shadow-lg">
                    ${getAvatarInitials(userName)}
                  </div>
                  <div>
                    <p class="font-medium text-gray-900 text-sm truncate dark:text-white">${escapeHtml(userName)}</p>
                    <p class="text-xs text-gray-500 dark:text-gray-400">Logged in</p>
                  </div>
                </div>
                <button data-action="logout" class="w-full text-left px-4 py-3 text-red-600 hover:bg-red-50 transition-all duration-200 active:scale-95 font-medium rounded-xl flex items-center gap-2 dark:hover:bg-red-900/20">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/></svg>
                  Sign Out
                </button>
              ` : `
                <a href="/login" class="flex items-center justify-center gap-2 w-full px-4 py-3 text-white bg-indigo-600 hover:bg-indigo-700 rounded-xl font-semibold transition-all duration-200 active:scale-95" role="menuitem" data-route>
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
                  Sign In
                </a>
              `}
            </div>
          </div>
        </div>
      </div>
    </header>
  `;

  setupHeaderEvents();
  updateCartBadge();
}

function setupHeaderEvents() {
  if (headerEventsInitialized) return;
  
  const headerRoot = document.getElementById('header-root');
  if (!headerRoot) return;
  
  // Theme toggle
  const themeToggleBtn = document.getElementById('theme-toggle-btn');
  const mobileThemeToggleBtn = document.getElementById('mobile-theme-toggle-btn');
  
  function toggleTheme() {
    const isDark = document.documentElement.classList.contains('dark');
    if (isDark) {
      document.documentElement.classList.remove('dark');
      localStorage.setItem('theme', 'light');
      updateThemeIcons(false);
    } else {
      document.documentElement.classList.add('dark');
      localStorage.setItem('theme', 'dark');
      updateThemeIcons(true);
    }
    // Dispatch event to notify app of theme change
    window.dispatchEvent(new CustomEvent('theme:change'));
  }
  
  themeToggleBtn?.addEventListener('click', toggleTheme);
  mobileThemeToggleBtn?.addEventListener('click', toggleTheme);
  
  // Event delegation for all header interactions
  headerRoot.addEventListener('click', async (e) => {
    const target = e.target;
    
    // Mobile menu toggle
    if (target.closest('#mobile-menu-btn')) {
      e.preventDefault();
      toggleMobileMenu();
    }
    
    // User dropdown toggle
    if (target.closest('#user-avatar-btn')) {
      e.preventDefault();
      e.stopPropagation();
      toggleUserDropdown();
    }
    
    // Logout handler
    if (target.closest('[data-action="logout"]')) {
      e.preventDefault();
      await handleLogout();
    }
    
    // Search result click
    if (target.closest('.search-result-item')) {
      const productId = target.closest('.search-result-item').dataset.productId;
      if (productId && window.router) {
        window.router.navigate(`/products/${productId}`);
        hideSearchResults();
      }
    }
  });
  
  // Search input handler
  const searchInput = document.getElementById('header-search-input');
  const searchDropdown = document.getElementById('search-results-dropdown');
  
  searchInput?.addEventListener('input', (e) => {
    const query = e.target.value.trim();
    
    if (query.length < 2) {
      hideSearchResults();
      return;
    }
    
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(async () => {
      try {
        const { api } = await import('../api.js');
        const data = await api.products.list({ search: query, limit: 5 });
        displaySearchResults(data.data || []);
      } catch (err) {
        console.error('Search failed:', err);
      }
    }, 300);
  });
  
  // Close search results on outside click or Escape
  document.addEventListener('click', (e) => {
    const dropdown = document.getElementById('user-dropdown-menu');
    const avatarBtn = document.getElementById('user-avatar-btn');
    
    // Close user dropdown on outside click
    if (dropdown && avatarBtn && 
        !dropdown.contains(e.target) && 
        !avatarBtn.contains(e.target)) {
      dropdown.classList.add('hidden');
      avatarBtn.setAttribute('aria-expanded', 'false');
    }
    
    // Close search results on outside click
    if (searchInput && searchDropdown &&
        !searchInput.contains(e.target) &&
        !searchDropdown.contains(e.target)) {
      hideSearchResults();
    }
  });
  
  // Close search on Escape
  searchInput?.addEventListener('keydown', (e) => {
    if (e.key === 'Escape') {
      hideSearchResults();
      searchInput.blur();
    }
  });
  
  headerEventsInitialized = true;
}

function displaySearchResults(products) {
  const dropdown = document.getElementById('search-results-dropdown');
  if (!dropdown) return;
  
  if (products.length === 0) {
    dropdown.innerHTML = `
      <div class="p-4 text-center text-gray-500 dark:text-gray-400">
        <p class="text-sm">No products found</p>
      </div>
    `;
  } else {
    dropdown.innerHTML = products.map(product => `
      <button class="search-result-item w-full text-left px-4 py-3 hover:bg-indigo-50 dark:hover:bg-gray-700 transition-colors duration-150 flex items-center gap-3" data-product-id="${product.product_id}">
        <div class="w-10 h-10 bg-gray-100 dark:bg-gray-700 rounded-lg flex items-center justify-center flex-shrink-0">
          <svg class="w-5 h-5 text-gray-400 dark:text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
          </svg>
        </div>
        <div class="flex-1 min-w-0">
          <p class="text-sm font-medium text-gray-900 dark:text-white truncate">${escapeHtml(product.name)}</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">₹${product.price?.toFixed(2) || '0.00'}${product.category_name ? ` · ${escapeHtml(product.category_name)}` : ''}</p>
        </div>
      </button>
    `).join('');
  }
  
  dropdown.classList.remove('hidden');
}

function hideSearchResults() {
  const dropdown = document.getElementById('search-results-dropdown');
  if (dropdown) {
    dropdown.classList.add('hidden');
    dropdown.innerHTML = '';
  }
}

function toggleMobileMenu() {
  const menu = document.getElementById('mobile-menu');
  const openIcon = document.getElementById('menu-icon-open');
  const closeIcon = document.getElementById('menu-icon-close');
  const btn = document.getElementById('mobile-menu-btn');
   
  if (!menu || !openIcon || !closeIcon || !btn) return;
   
  const isOpen = !menu.classList.toggle('hidden');
   
  openIcon.classList.toggle('hidden', isOpen);
  closeIcon.classList.toggle('hidden', !isOpen);
  btn.setAttribute('aria-expanded', isOpen);
}

function toggleUserDropdown() {
  const dropdown = document.getElementById('user-dropdown-menu');
  const avatarBtn = document.getElementById('user-avatar-btn');
   
  if (!dropdown || !avatarBtn) return;
   
  const isOpen = !dropdown.classList.contains('hidden');
   
  if (isOpen) {
    dropdown.classList.add('hidden');
    avatarBtn.setAttribute('aria-expanded', 'false');
  } else {
    dropdown.classList.remove('hidden');
    avatarBtn.setAttribute('aria-expanded', 'true');
  }
}

async function handleLogout() {
  const { api } = await import('../api.js');
  await api.auth.logout();
  window._isAuthenticated = false;
  window._userName = '';
  window._userEmail = '';
   
  // Dispatch auth:change event
  window.dispatchEvent(new CustomEvent('auth:change'));
   
  // Show toast if available
  if (window.router?.showToast) {
    window.router.showToast('Logged out successfully', 'success');
  }
}

function getAvatarInitials(name) {
  if (!name) return '??';
  const parts = name.trim().split(/\s+/);
  if (parts.length === 1) return parts[0].charAt(0).toUpperCase();
  return (parts[0].charAt(0) + parts[parts.length - 1].charAt(0)).toUpperCase();
}

function escapeHtml(text) {
  const div = document.createElement('div');
  div.textContent = text;
  return div.innerHTML;
}

function updateCartBadge() {
  try {
    const cart = JSON.parse(localStorage.getItem('ecommerce_cart') || '[]');
    const count = cart.reduce((sum, item) => sum + (item.quantity || 0), 0);
     
    // Desktop badge
    const desktopBadge = document.getElementById('header-cart-badge');
    if (desktopBadge) {
      if (count > 0) {
        desktopBadge.textContent = count > 99 ? '99+' : count;
        desktopBadge.classList.remove('hidden', 'scale-0');
        desktopBadge.classList.add('animate-scale-badge');
        setTimeout(() => desktopBadge.classList.remove('animate-scale-badge'), 300);
      } else {
        desktopBadge.classList.add('hidden');
      }
    }
     
    // Mobile badge  
    const mobileBadge = document.getElementById('mobile-cart-badge');
    if (mobileBadge) {
      if (count > 0) {
        mobileBadge.textContent = count > 99 ? '99+' : count;
        mobileBadge.classList.remove('hidden');
      } else {
        mobileBadge.classList.add('hidden');
      }
    }
  } catch (e) {
    // Ignore errors
  }
}

export { updateCartBadge };
