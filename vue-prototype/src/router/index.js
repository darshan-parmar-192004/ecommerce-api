import { createRouter, createWebHistory } from "vue-router";
import { isAuthenticated, isAdmin } from "@/lib/auth";

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
          component: () => import("@/pages/HomeView.vue"),
          meta: { title: "Home" },
        },
        {
          path: "products",
          name: "Products",
          component: () => import("@/pages/shop/ProductListView.vue"),
          meta: { title: "Products" },
        },
        {
          path: "product/:id",
          name: "ProductDetail",
          component: () => import("@/pages/shop/ProductDetailView.vue"),
          meta: { title: "Product Details" },
        },
        {
          path: "cart",
          name: "Cart",
          component: () => import("@/pages/shop/CartView.vue"),
          meta: { title: "Shopping Cart" },
        },
        {
          path: "checkout",
          name: "Checkout",
          component: () => import("@/pages/shop/CheckoutView.vue"),
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
          component: () => import("@/pages/auth/LoginView.vue"),
          meta: { guest: true, title: "Login" },
        },
        {
          path: "register",
          name: "Register",
          component: () => import("@/pages/auth/RegisterView.vue"),
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
          component: () => import("@/pages/user/DashboardView.vue"),
          meta: { title: "Dashboard" },
        },
        {
          path: "orders",
          name: "UserOrders",
          component: () => import("@/pages/user/OrderHistoryView.vue"),
          meta: { title: "Order History" },
        },
        {
          path: "orders/:id",
          name: "UserOrderDetail",
          component: () => import("@/pages/user/OrderDetailView.vue"),
          meta: { title: "Order Details" },
        },
        {
          path: "profile",
          name: "UserProfile",
          component: () => import("@/pages/user/ProfileView.vue"),
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
          component: () => import("@/pages/admin/AdminDashboardView.vue"),
          meta: { title: "Admin Dashboard" },
        },
        {
          path: "products",
          name: "AdminProducts",
          component: () => import("@/pages/admin/ProductManageView.vue"),
          meta: { title: "Manage Products" },
        },
        {
          path: "inventory",
          name: "AdminInventory",
          component: () => import("@/pages/admin/InventoryView.vue"),
          meta: { title: "Inventory Management" },
        },
        {
          path: "categories",
          name: "AdminCategories",
          component: () => import("@/pages/admin/CategoryManageView.vue"),
          meta: { title: "Category Management" },
        },
      ],
    },
    {
      path: "/404",
      name: "NotFound",
      component: () => import("@/pages/error/NotFoundView.vue"),
      meta: { title: "404 Not Found" },
    },
    {
      path: "/unauthorized",
      name: "Unauthorized",
      component: () => import("@/pages/error/UnauthorizedView.vue"),
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
    const requiresAuth = to.meta.requiresAuth;
    const requiresAdmin = to.meta.requiresAdmin;
    const isGuestRoute = to.meta.guest;

    if (isGuestRoute && isAuthenticated()) {
      next({ name: "Home" });
    } else if (requiresAuth && !isAuthenticated()) {
      next({ name: "Login", query: { redirect: to.fullPath } });
    } else if (requiresAdmin && !isAdmin()) {
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
