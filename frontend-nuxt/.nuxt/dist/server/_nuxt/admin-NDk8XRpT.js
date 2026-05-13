import { l as defineNuxtRouteMiddleware, d as useAuthStore, n as navigateTo, j as createError } from "../server.mjs";
import "vue";
import "ofetch";
import "#internal/nuxt/paths";
import "hookable";
import "unctx";
import "h3";
import "ufo";
import "unhead";
import "@unhead/shared";
import "vue-router";
import "radix3";
import "defu";
import "klona";
import "@primeuix/utils/eventbus";
import "@primeuix/styled";
import "@primeuix/utils";
import "@primeuix/utils/object";
import "@primeuix/styles/base";
import "@primeuix/utils/dom";
import "cookie-es";
import "destr";
import "ohash";
import "vue/server-renderer";
import "@primeuix/utils/zindex";
import "@primeuix/styles/toast";
import "@primeuix/utils/uuid";
import "@primeuix/styles/ripple";
import "@primeuix/styles/badge";
import "@primeuix/styles/button";
const admin = defineNuxtRouteMiddleware((to) => {
  const authStore = useAuthStore();
  {
    authStore.loadAuthFromCookie();
  }
  if (!authStore.isAuthenticated || !authStore.token) {
    return navigateTo(`/auth/login?redirect=${encodeURIComponent(to.fullPath)}`);
  }
  if (!authStore.isAdmin) {
    throw createError({
      statusCode: 403,
      message: "Access denied. Admin privileges required."
    });
  }
});
export {
  admin as default
};
//# sourceMappingURL=admin-NDk8XRpT.js.map
