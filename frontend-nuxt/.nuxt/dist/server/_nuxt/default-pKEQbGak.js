import { _ as __nuxt_component_0 } from "./nuxt-link-1RlcU5D7.js";
import { defineAsyncComponent, ref, mergeProps, withCtx, createTextVNode, unref, createVNode, toDisplayString, openBlock, createBlock, createCommentVNode, useSSRContext } from "vue";
import { k as defineStore, b as useCartStore, d as useAuthStore, u as useRoute, M as storeToRefs, n as navigateTo } from "../server.mjs";
import "hookable";
import "klona";
import "defu";
import "#internal/nuxt/paths";
import { u as useApi } from "./useApi-CR6WE1xi.js";
import script$2 from "./index-ByMj2CEz.js";
import script$1 from "./index-zRb3CJ4f.js";
import { ssrRenderAttrs, ssrRenderComponent, ssrRenderStyle, ssrInterpolate, ssrRenderClass, ssrRenderSlot } from "vue/server-renderer";
import script from "./index-amHC6YS-.js";
import { _ as _export_sfc } from "./_plugin-vue_export-helper-1tPrXgE0.js";
import "ufo";
import "ofetch";
import "unctx";
import "h3";
import "unhead";
import "@unhead/shared";
import "vue-router";
import "radix3";
import "@primeuix/utils/eventbus";
import "@primeuix/styled";
import "@primeuix/utils";
import "@primeuix/utils/object";
import "@primeuix/styles/base";
import "@primeuix/utils/dom";
import "cookie-es";
import "destr";
import "ohash";
import "@primeuix/utils/zindex";
import "@primeuix/styles/toast";
import "@primeuix/utils/uuid";
import "@primeuix/styles/ripple";
import "@primeuix/styles/badge";
import "@primeuix/styles/button";
import "./index-Bz-g1Z_Q.js";
import "./index-Bg3yUxEP.js";
import "./index-DzyF1dFx.js";
import "./index-rAVNvoJo.js";
import "@primeuix/styles/popover";
import "@primeuix/styles/avatar";
import "./index-DrzwqwtL.js";
import "@primeuix/styles/toggleswitch";
const useThemeStore = defineStore("theme", {
  state: () => ({
    themeMode: "light",
    // 'light' or 'dark'
    initialized: false
  }),
  getters: {
    isDark: (state) => state.themeMode === "dark",
    currentMode: (state) => state.themeMode
  },
  actions: {
    initializeTheme() {
    },
    toggleTheme() {
      this.themeMode = this.themeMode === "dark" ? "light" : "dark";
      this.applyTheme();
    },
    setThemeMode(mode) {
      this.themeMode = mode;
      this.applyTheme();
    },
    applyTheme() {
    }
  }
});
const _sfc_main = {
  __name: "default",
  __ssrInlineRender: true,
  setup(__props) {
    const CartDrawer = defineAsyncComponent(() => import("./CartDrawer-cIxyoP0w.js"));
    const cartStore = useCartStore();
    const authStore = useAuthStore();
    const themeStore = useThemeStore();
    useRoute();
    const isMenuOpen = ref(false);
    const userMenu = ref(null);
    const { isAuthenticated, user } = storeToRefs(authStore);
    const closeMenu = () => {
      isMenuOpen.value = false;
    };
    const getUserInitial = (name) => {
      if (!name) return "U";
      return name.charAt(0).toUpperCase();
    };
    const { auth } = useApi();
    const handleLogout = async () => {
      try {
        await auth.logout();
      } catch (e) {
        console.error("Logout error:", e);
      } finally {
        authStore.clearAuth();
        cartStore.clearCart();
        closeMenu();
        navigateTo("/");
      }
    };
    const { isDark } = storeToRefs(themeStore);
    return (_ctx, _push, _parent, _attrs) => {
      var _a;
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Avatar = script$1;
      const _component_OverlayPanel = script$2;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "min-h-screen flex flex-col relative overflow-hidden bg-surface font-body text-on_surface" }, _attrs))} data-v-14edafce><header class="sticky top-0 z-40 backdrop-blur-xl bg-surface/70 border-b border-white/10" data-v-14edafce><nav class="max-w-7xl mx-auto px-6 lg:px-8" data-v-14edafce><div class="flex justify-between items-center h-20" data-v-14edafce><div class="flex items-center gap-12" data-v-14edafce>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/",
        class: "text-2xl font-bold tracking-tight font-display text-primary",
        onClick: closeMenu
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(` The Curator `);
          } else {
            return [
              createTextVNode(" The Curator ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`<div class="hidden md:flex items-center gap-8" data-v-14edafce>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/products",
        class: "text-xs font-semibold uppercase tracking-widest transition-colors duration-300 text-on_surface_variant",
        "active-class": "text-primary"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(` Products `);
          } else {
            return [
              createTextVNode(" Products ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/cart",
        class: "text-xs font-semibold uppercase tracking-widest transition-colors duration-300 text-on_surface_variant",
        "active-class": "text-primary"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(` Cart `);
          } else {
            return [
              createTextVNode(" Cart ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</div></div><div style="${ssrRenderStyle({ "display": "flex", "align-items": "center", "gap": "1.5rem" })}" data-v-14edafce><button class="relative p-2.5 rounded-full transition-colors duration-300 hover:bg-surface-container text-on_surface_variant" aria-label="Shopping cart" data-v-14edafce><i class="pi pi-shopping-cart text-xl" data-v-14edafce></i>`);
      if (unref(cartStore).totalItems > 0) {
        _push(`<span class="absolute -top-0.5 -right-0.5 text-[10px] w-5 h-5 rounded-full flex items-center justify-center font-bold animate-bounce-subtle bg-gradient-to-r from-primary to-primary-container text-white" data-v-14edafce>${ssrInterpolate(unref(cartStore).totalItems)}</span>`);
      } else {
        _push(`<!---->`);
      }
      _push(`</button>`);
      _push(ssrRenderComponent(unref(script), {
        modelValue: unref(isDark),
        "onUpdate:modelValue": ($event) => unref(themeStore).toggleTheme()
      }, {
        handle: withCtx(({ checked }, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`<i class="${ssrRenderClass(["!text-xs pi", checked ? "pi-moon" : "pi-sun"])}" data-v-14edafce${_scopeId}></i>`);
          } else {
            return [
              createVNode("i", {
                class: ["!text-xs pi", checked ? "pi-moon" : "pi-sun"]
              }, null, 2)
            ];
          }
        }),
        _: 1
      }, _parent));
      if (unref(isAuthenticated)) {
        _push(`<div class="relative" data-v-14edafce><button class="flex items-center gap-2 p-1 rounded-full transition-colors duration-300 hover:bg-surface-container" data-v-14edafce>`);
        _push(ssrRenderComponent(_component_Avatar, {
          label: getUserInitial((_a = unref(user)) == null ? void 0 : _a.name),
          shape: "circle",
          class: "bg-gradient-to-r from-primary to-primary-container text-white font-semibold"
        }, null, _parent));
        _push(`</button>`);
        _push(ssrRenderComponent(_component_OverlayPanel, {
          ref_key: "userMenu",
          ref: userMenu,
          style: { width: "16rem" },
          class: "!bg-surface-container-lowest"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            var _a2, _b, _c, _d;
            if (_push2) {
              _push2(`<div class="p-4 border-b border-outline-variant/20" data-v-14edafce${_scopeId}><p class="text-sm font-semibold text-on_surface" data-v-14edafce${_scopeId}>${ssrInterpolate(((_a2 = unref(user)) == null ? void 0 : _a2.name) || "User")}</p><p class="text-xs text-outline" data-v-14edafce${_scopeId}>${ssrInterpolate(((_b = unref(user)) == null ? void 0 : _b.email) || "")}</p></div><div class="py-2" data-v-14edafce${_scopeId}>`);
              _push2(ssrRenderComponent(_component_NuxtLink, {
                to: "/orders",
                class: "flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors"
              }, {
                default: withCtx((_2, _push3, _parent3, _scopeId2) => {
                  if (_push3) {
                    _push3(`<i class="pi pi-shopping-cart text-on_surface_variant" data-v-14edafce${_scopeId2}></i><span class="text-sm font-medium text-on_surface_variant" data-v-14edafce${_scopeId2}>My Orders</span>`);
                  } else {
                    return [
                      createVNode("i", { class: "pi pi-shopping-cart text-on_surface_variant" }),
                      createVNode("span", { class: "text-sm font-medium text-on_surface_variant" }, "My Orders")
                    ];
                  }
                }),
                _: 1
              }, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_NuxtLink, {
                to: "/user/profile",
                class: "flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors"
              }, {
                default: withCtx((_2, _push3, _parent3, _scopeId2) => {
                  if (_push3) {
                    _push3(`<i class="pi pi-user text-on_surface_variant" data-v-14edafce${_scopeId2}></i><span class="text-sm font-medium text-on_surface_variant" data-v-14edafce${_scopeId2}>Profile</span>`);
                  } else {
                    return [
                      createVNode("i", { class: "pi pi-user text-on_surface_variant" }),
                      createVNode("span", { class: "text-sm font-medium text-on_surface_variant" }, "Profile")
                    ];
                  }
                }),
                _: 1
              }, _parent2, _scopeId));
              if (unref(authStore).isAdmin) {
                _push2(ssrRenderComponent(_component_NuxtLink, {
                  to: "/admin",
                  class: "flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors text-primary"
                }, {
                  default: withCtx((_2, _push3, _parent3, _scopeId2) => {
                    if (_push3) {
                      _push3(`<i class="pi pi-shield" data-v-14edafce${_scopeId2}></i><span class="text-sm font-medium" data-v-14edafce${_scopeId2}>Admin Panel</span>`);
                    } else {
                      return [
                        createVNode("i", { class: "pi pi-shield" }),
                        createVNode("span", { class: "text-sm font-medium" }, "Admin Panel")
                      ];
                    }
                  }),
                  _: 1
                }, _parent2, _scopeId));
              } else {
                _push2(`<!---->`);
              }
              _push2(`<hr class="mx-3 my-1 border-outline-variant/20" data-v-14edafce${_scopeId}>`);
              _push2(ssrRenderComponent(_component_NuxtLink, {
                to: "/",
                class: "flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors"
              }, {
                default: withCtx((_2, _push3, _parent3, _scopeId2) => {
                  if (_push3) {
                    _push3(`<i class="pi pi-home text-on_surface_variant" data-v-14edafce${_scopeId2}></i><span class="text-sm font-medium text-on_surface_variant" data-v-14edafce${_scopeId2}>Back to Website</span>`);
                  } else {
                    return [
                      createVNode("i", { class: "pi pi-home text-on_surface_variant" }),
                      createVNode("span", { class: "text-sm font-medium text-on_surface_variant" }, "Back to Website")
                    ];
                  }
                }),
                _: 1
              }, _parent2, _scopeId));
              _push2(`<button class="w-full flex items-center gap-3 px-4 py-3 hover:bg-error-container/30 transition-colors text-error" data-v-14edafce${_scopeId}><i class="pi pi-sign-out" data-v-14edafce${_scopeId}></i><span class="text-sm font-medium" data-v-14edafce${_scopeId}>Logout</span></button></div>`);
            } else {
              return [
                createVNode("div", { class: "p-4 border-b border-outline-variant/20" }, [
                  createVNode("p", { class: "text-sm font-semibold text-on_surface" }, toDisplayString(((_c = unref(user)) == null ? void 0 : _c.name) || "User"), 1),
                  createVNode("p", { class: "text-xs text-outline" }, toDisplayString(((_d = unref(user)) == null ? void 0 : _d.email) || ""), 1)
                ]),
                createVNode("div", { class: "py-2" }, [
                  createVNode(_component_NuxtLink, {
                    to: "/orders",
                    class: "flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors"
                  }, {
                    default: withCtx(() => [
                      createVNode("i", { class: "pi pi-shopping-cart text-on_surface_variant" }),
                      createVNode("span", { class: "text-sm font-medium text-on_surface_variant" }, "My Orders")
                    ]),
                    _: 1
                  }),
                  createVNode(_component_NuxtLink, {
                    to: "/user/profile",
                    class: "flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors"
                  }, {
                    default: withCtx(() => [
                      createVNode("i", { class: "pi pi-user text-on_surface_variant" }),
                      createVNode("span", { class: "text-sm font-medium text-on_surface_variant" }, "Profile")
                    ]),
                    _: 1
                  }),
                  unref(authStore).isAdmin ? (openBlock(), createBlock(_component_NuxtLink, {
                    key: 0,
                    to: "/admin",
                    class: "flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors text-primary"
                  }, {
                    default: withCtx(() => [
                      createVNode("i", { class: "pi pi-shield" }),
                      createVNode("span", { class: "text-sm font-medium" }, "Admin Panel")
                    ]),
                    _: 1
                  })) : createCommentVNode("", true),
                  createVNode("hr", { class: "mx-3 my-1 border-outline-variant/20" }),
                  createVNode(_component_NuxtLink, {
                    to: "/",
                    class: "flex items-center gap-3 px-4 py-3 hover:bg-surface-container-low transition-colors"
                  }, {
                    default: withCtx(() => [
                      createVNode("i", { class: "pi pi-home text-on_surface_variant" }),
                      createVNode("span", { class: "text-sm font-medium text-on_surface_variant" }, "Back to Website")
                    ]),
                    _: 1
                  }),
                  createVNode("button", {
                    onClick: handleLogout,
                    class: "w-full flex items-center gap-3 px-4 py-3 hover:bg-error-container/30 transition-colors text-error"
                  }, [
                    createVNode("i", { class: "pi pi-sign-out" }),
                    createVNode("span", { class: "text-sm font-medium" }, "Logout")
                  ])
                ])
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div>`);
      } else {
        _push(`<!--[-->`);
        _push(ssrRenderComponent(_component_NuxtLink, {
          to: "/auth/login",
          class: "px-5 py-2.5 text-white rounded-md font-medium text-sm transition-all duration-300 bg-gradient-to-r from-primary to-primary-container shadow-glow hover:shadow-xl hover:scale-[1.02] active:scale-[0.98]"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(` Login `);
            } else {
              return [
                createTextVNode(" Login ")
              ];
            }
          }),
          _: 1
        }, _parent));
        if (unref(authStore).isAdmin) {
          _push(ssrRenderComponent(_component_NuxtLink, {
            to: "/admin",
            class: "text-sm font-medium text-primary"
          }, {
            default: withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                _push2(` Admin Panel `);
              } else {
                return [
                  createTextVNode(" Admin Panel ")
                ];
              }
            }),
            _: 1
          }, _parent));
        } else {
          _push(`<!---->`);
        }
        _push(`<!--]-->`);
      }
      _push(`</div></div></nav></header><main class="flex-1" data-v-14edafce>`);
      ssrRenderSlot(_ctx.$slots, "default", {}, null, _push, _parent);
      _push(`</main><footer class="mt-auto bg-surface-container-low" data-v-14edafce><div class="max-w-7xl mx-auto px-6 lg:px-8 py-16" data-v-14edafce><div class="text-center" data-v-14edafce><p class="text-xs uppercase tracking-widest text-outline font-body" data-v-14edafce> © 2024 The Curator. All rights reserved. </p></div></div></footer>`);
      _push(ssrRenderComponent(unref(CartDrawer), null, null, _parent));
      _push(`</div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("layouts/default.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const _default = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-14edafce"]]);
export {
  _default as default
};
//# sourceMappingURL=default-pKEQbGak.js.map
