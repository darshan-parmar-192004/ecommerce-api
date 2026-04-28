export function renderFooter() {
  const footerRoot = document.getElementById('footer-root');
  if (!footerRoot) return;
  
  footerRoot.innerHTML = `
    <footer class="bg-white border-t border-gray-100 mt-auto dark:bg-gray-900 dark:border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-8 mb-8">
          <div class="md:col-span-2">
            <div class="flex items-center gap-2.5 mb-3">
              <div class="w-9 h-9 rounded-xl bg-gradient-to-br from-indigo-500 to-indigo-700 flex items-center justify-center shadow-lg active:scale-95 transition-all duration-200">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"/>
                </svg>
              </div>
              <span class="text-xl font-bold text-gray-900 dark:text-white">LuxeCart</span>
            </div>
            <p class="text-gray-500 text-sm leading-relaxed max-w-md dark:text-gray-400">
              Premium shopping experience with curated products, fast shipping, and exceptional customer service.
            </p>
          </div>
          
          <div>
            <h3 class="text-sm font-semibold text-gray-900 uppercase tracking-wider mb-3 dark:text-white">Shop</h3>
            <ul class="space-y-2">
              <li><a href="/products" class="footer-link text-gray-500 hover:text-indigo-600 text-sm transition-all duration-200 active:scale-95 inline-block dark:text-gray-400 dark:hover:text-indigo-400" data-route>All Products</a></li>
              <li><a href="/products" class="footer-link text-gray-500 hover:text-indigo-600 text-sm transition-all duration-200 active:scale-95 inline-block dark:text-gray-400 dark:hover:text-indigo-400" data-route>New Arrivals</a></li>
              <li><a href="/cart" class="footer-link text-gray-500 hover:text-indigo-600 text-sm transition-all duration-200 active:scale-95 inline-block dark:text-gray-400 dark:hover:text-indigo-400" data-route>Shopping Cart</a></li>
            </ul>
          </div>
          
          <div>
            <h3 class="text-sm font-semibold text-gray-900 uppercase tracking-wider mb-3 dark:text-white">Support</h3>
            <ul class="space-y-2">
              <li><a href="#" class="text-gray-500 hover:text-indigo-600 text-sm transition-all duration-200 active:scale-95 inline-block dark:text-gray-400 dark:hover:text-indigo-400">Contact Us</a></li>
              <li><a href="#" class="text-gray-500 hover:text-indigo-600 text-sm transition-all duration-200 active:scale-95 inline-block dark:text-gray-400 dark:hover:text-indigo-400">FAQs</a></li>
              <li><a href="#" class="text-gray-500 hover:text-indigo-600 text-sm transition-all duration-200 active:scale-95 inline-block dark:text-gray-400 dark:hover:text-indigo-400">Shipping Info</a></li>
            </ul>
          </div>
        </div>
        
        <div class="border-t border-gray-100 pt-8 flex flex-col md:flex-row items-center justify-between gap-4 dark:border-gray-700">
          <p class="text-gray-500 text-sm dark:text-gray-400">
            &copy; ${new Date().getFullYear()} LuxeCart. All rights reserved.
          </p>
          <div class="flex items-center gap-4">
            <a href="#" class="text-gray-400 hover:text-indigo-600 transition-all duration-200 active:scale-95 inline-block dark:text-gray-500 dark:hover:text-indigo-400">
              <span class="sr-only">Twitter</span>
              <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24"><path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/></svg>
            </a>
            <a href="#" class="text-gray-400 hover:text-indigo-600 transition-all duration-200 active:scale-95 inline-block dark:text-gray-500 dark:hover:text-indigo-400">
              <span class="sr-only">Instagram</span>
              <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 1.065.049 1.791.218 2.427.465a4.902 4.902 0 011.772 1.153 4.902 4.902 0 011.153 1.772c.247.636.416 1.362.465 2.427.057 1.266.07 1.646.07 4.85s-.015 3.585-.074 4.85c-.049 1.065-.218 1.791-.465 2.427a4.902 4.902 0 01-1.153 1.772 4.902 4.902 0 01-1.772 1.153c-.636.247-1.362.416-2.427.465-1.266.057-1.646.07-4.85.07s-3.585-.015-4.85-.074c-1.065-.049-1.791-.218-2.427-.465a4.902 4.902 0 01-1.772-1.153 4.902 4.902 0 01-1.153-1.772c-.247-.636-.416-1.362-.465-2.427-.057-1.266-.07-1.646-.07-4.85s.015-3.585.074-4.85c.049-1.065.218-1.791.465-2.427a4.902 4.902 0 011.153-1.772A4.902 4.902 0 017.45.63c.636-.247 1.362-.416 2.427-.465C11.085.015 11.465 0 12 0zm0 2.163c-3.18 0-3.544.012-4.803.07-1.09.047-1.78.22-2.26.37a3.15 3.15 0 00-1.36.864 3.15 3.15 0 00-.864 1.36c-.15.48-.322 1.17-.37 2.26-.058 1.26-.07 1.623-.07 4.804s.012 3.544.07 4.804c.048 1.09.22 1.78.37 2.26.148.44.358.818.64 1.117.28.298.642.508 1.087.638.48.15 1.17.322 2.26.37 1.26.058 1.623.07 4.804.07s3.544-.012 4.803-.07c1.09-.048 1.78-.22 2.26-.37.44-.148.818-.358 1.117-.64.298-.28.508-.642.638-1.087.15-.48.322-1.17.37-2.26.058-1.259.07-1.623.07-4.804s-.012-3.544-.07-4.804c-.048-1.09-.22-1.78-.37-2.26a3.15 3.15 0 00-.864-1.36 3.15 3.15 0 00-1.36-.864c-.48-.15-1.17-.322-2.26-.37-1.26-.057-1.623-.07-4.804-.07zm0 5.838a6.162 6.162 0 100 12.324 6.162 6.162 0 000-12.324zM12 16a4 4 0 110-8 4 4 0 010 8zm6.406-11.845a1.44 1.44 0 100 2.881 1.44 1.44 0 000-2.881z"/></svg>
            </a>
          </div>
        </div>
      </div>
    </footer>
  `;
}
