import { l as defineNuxtRouteMiddleware, d as useAuthStore, m as useCookie, o as executeAsync, n as navigateTo } from "../server.mjs";
import { nextTick } from "vue";
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
const auth = defineNuxtRouteMiddleware(async (to) => {
  let __temp, __restore;
  const authStore = useAuthStore();
  const token = useCookie("auth_token");
  if (!token.value && !authStore.token) {
    [__temp, __restore] = executeAsync(() => nextTick()), await __temp, __restore();
    if (token.value || authStore.token) return;
    return navigateTo(`/auth/login?redirect=${encodeURIComponent(to.fullPath)}`);
  }
});
export {
  auth as default
};
//# sourceMappingURL=auth-BlGr1aCR.js.map
