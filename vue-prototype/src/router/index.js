import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "@/stores/auth";

/**
 * @typedef {Object} RouteMeta
 * @property {boolean} [requiresAuth] - Route requires authentication
 * @property {boolean} [requiresAdmin] - Route requires admin role
 * @property {boolean} [guest] - Route only accessible to guests
 * @property {string} [title] - Page title for document head
 **/
const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: () => import("@/layouts/DefaultLayout.vue"),
      children: [
        {
          path: "",
          name: "Home",
          component: () => import("@/views/HomeView.vue"),
          meta: { title: "Home" },
        },
        {
          path: "products",
          name: "Products",
          component: () => import("@/views/shop/ProductListView.vue"),
          meta: { title: "Products" },
        },
        {
          path: "product/:id",
          name: "ProductDetail",
          component: () => import("@/views/shop/ProductDetailView.vue"),
          meta: { title: "Product Details" },
        },
        {
          path: "cart",
          name: "Cart",
          component: () => import("@/views/shop/CartView.vue"),
          meta: { title: "Shopping Cart" },
        },
        {
          path: "checkout",
          name: "Checkout",
          component: () => import("@/views/shop/CheckoutView.vue"),
          meta: { requiresAuth: true, title: "Checkout" },
        },
      ],
    },
    {
      path: "/auth",
      component: () => import("@/layouts/AuthLayout.vue"),
      children: [
        {
          path: "login",
          name: "Login",
          component: () => import("@/views/auth/LoginView.vue"),
          meta: { guest: true, title: "Login" },
        },
        {
          path: "register",
          name: "Register",
          component: () => import("@/views/auth/RegisterView.vue"),
          meta: { guest: true, title: "Register" },
        },
      ],
    },
    {
      path: "/user",
      component: () => import("@/layouts/DefaultLayout.vue"),
      meta: { requiresAuth: true },
      children: [
        {
          path: "dashboard",
          name: "UserDashboard",
          component: () => import("@/views/user/DashboardView.vue"),
          meta: { title: "Dashboard" },
        },
        {
          path: "orders",
          name: "UserOrders",
          component: () => import("@/views/user/OrderHistoryView.vue"),
          meta: { title: "Order History" },
        },
        {
          path: "orders/:id",
          name: "UserOrderDetail",
          component: () => import("@/views/user/OrderDetailView.vue"),
          meta: { title: "Order Details" },
        },
        {
          path: "profile",
          name: "UserProfile",
          component: () => import("@/views/user/ProfileView.vue"),
          meta: { title: "Profile" },
        },
      ],
    },
    {
      path: "/admin",
      component: () => import("@/layouts/AdminLayout.vue"),
      meta: { requiresAuth: true, requiresAdmin: true },
      children: [
        {
          path: "",
          name: "AdminDashboard",
          component: () => import("@/views/admin/AdminDashboardView.vue"),
          meta: { title: "Admin Dashboard" },
        },
        {
          path: "products",
          name: "AdminProducts",
          component: () => import("@/views/admin/ProductManageView.vue"),
          meta: { title: "Manage Products" },
        },
        {
          path: "inventory",
          name: "AdminInventory",
          component: () => import("@/views/admin/InventoryView.vue"),
          meta: { title: "Inventory Management" },
        },
        {
          path: "categories",
          name: "AdminCategories",
          component: () => import("@/views/admin/CategoryManageView.vue"),
          meta: { title: "Category Management" },
        },
      ],
    },
    {
      path: "/404",
      name: "NotFound",
      component: () => import("@/views/error/NotFoundView.vue"),
      meta: { title: "404 Not Found" },
    },
    {
      path: "/unauthorized",
      name: "Unauthorized",
      component: () => import("@/views/error/UnauthorizedView.vue"),
      meta: { title: "Unauthorized" },
    },
    {
      path: "/:pathMatch(.*)*",
      redirect: "/404",
    },
  ],
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) return savedPosition;
    if (to.hash) return { el: to.hash, behavior: "smooth", top: 80 };
    return { top: 0 };
  },
});

router.beforeEach(async (to, from, next) => {
  try {
    const authStore = useAuthStore();
    const requiresAuth = to.meta.requiresAuth;
    const requiresAdmin = to.meta.requiresAdmin;
    const isGuestRoute = to.meta.guest;

    if (isGuestRoute && authStore.isAuthenticated) {
      next({ name: "Home" });
    } else if (requiresAuth && !authStore.isAuthenticated) {
      next({ name: "Login", query: { redirect: to.fullPath } });
    } else if (requiresAdmin && !authStore.isAdmin) {
      next({ name: "Unauthorized" });
    } else {
      document.title = to.meta.title
        ? `${to.meta.title} | E-Commerce`
        : "E-Commerce";
      next();
    }
  } catch (error) {
    console.error("Navigation guard error:", error);
    next(false);
  }
});

export default router;
