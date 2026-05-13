import { reactive, ref, mergeProps, unref, withCtx, createTextVNode, useSSRContext } from "vue";
import { k as defineStore, d as useAuthStore, s as script$2 } from "../server.mjs";
import script$1 from "./index-DDXXjm8b.js";
import script from "./index-BAz9mnEE.js";
import { ssrRenderAttrs, ssrRenderComponent } from "vue/server-renderer";
import { u as useApi } from "./useApi-CR6WE1xi.js";
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
import "@primeuix/utils/zindex";
import "@primeuix/styles/toast";
import "@primeuix/utils/uuid";
import "@primeuix/styles/ripple";
import "@primeuix/styles/badge";
import "@primeuix/styles/button";
import "./index-Dedjwyps.js";
import "./index-DrzwqwtL.js";
import "@primeuix/styles/inputtext";
import "@primeuix/styles/message";
const useUserStore = defineStore("user", {
  state: () => ({
    profile: null,
    orders: [],
    loading: false,
    ordersLoading: false
  }),
  getters: {
    hasProfile: (state) => !!state.profile,
    orderCount: (state) => state.orders.length
  },
  actions: {
    async fetchProfile() {
      this.loading = true;
      try {
        const { auth } = useApi();
        const data = await auth.me();
        this.profile = data.data || data.customer || data;
        return this.profile;
      } catch (error) {
        console.error("Failed to fetch profile:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },
    async fetchOrders() {
      this.ordersLoading = true;
      try {
        const { orders } = useApi();
        const data = await orders.list();
        this.orders = data.data || data || [];
        return this.orders;
      } catch (error) {
        console.error("Failed to fetch orders:", error);
        throw error;
      } finally {
        this.ordersLoading = false;
      }
    },
    async fetchOrder(id) {
      try {
        const { orders } = useApi();
        const data = await orders.get(id);
        return data.data || data;
      } catch (error) {
        console.error("Failed to fetch order:", error);
        throw error;
      }
    },
    async updateProfile(profileData) {
      this.loading = true;
      try {
        const { fetchJson } = useApi();
        const data = await fetchJson("/customers/me", {
          method: "PUT",
          body: JSON.stringify(profileData)
        });
        this.profile = { ...this.profile, ...data };
        return this.profile;
      } catch (error) {
        console.error("Failed to update profile:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },
    clearUserData() {
      this.profile = null;
      this.orders = [];
    },
    async cancelOrder(id) {
      this.loading = true;
      try {
        const { fetchJson } = useApi();
        await fetchJson(`/orders/${id}/cancel`, {
          method: "PUT"
        });
        const orderIndex = this.orders.findIndex((o) => o.order_id === id);
        if (orderIndex !== -1) {
          this.orders[orderIndex].status = "cancelled";
        }
        return true;
      } catch (error) {
        console.error("Failed to cancel order:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    }
  }
});
const _sfc_main = {
  __name: "profile",
  __ssrInlineRender: true,
  setup(__props) {
    useUserStore();
    useAuthStore();
    const form = reactive({ name: "", email: "", phone: "" });
    const loading = ref(false);
    const success = ref(false);
    return (_ctx, _push, _parent, _attrs) => {
      const _component_Message = script;
      const _component_InputText = script$1;
      const _component_Button = script$2;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "max-w-2xl mx-auto px-4 sm:px-6 lg:px-8 py-8" }, _attrs))}><h1 class="text-3xl font-bold text-on_surface mb-8">Profile Settings</h1><div class="bg-surface-container-lowest rounded-lg p-6 shadow-ambient">`);
      if (unref(success)) {
        _push(ssrRenderComponent(_component_Message, {
          severity: "success",
          closable: false,
          class: "mb-4"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(` Profile updated successfully! `);
            } else {
              return [
                createTextVNode(" Profile updated successfully! ")
              ];
            }
          }),
          _: 1
        }, _parent));
      } else {
        _push(`<!---->`);
      }
      _push(`<form class="space-y-6"><div><label class="block text-sm font-medium text-on_surface mb-2">Name</label>`);
      _push(ssrRenderComponent(_component_InputText, {
        modelValue: unref(form).name,
        "onUpdate:modelValue": ($event) => unref(form).name = $event,
        type: "text",
        class: "w-full"
      }, null, _parent));
      _push(`</div><div><label class="block text-sm font-medium text-on_surface mb-2">Email</label>`);
      _push(ssrRenderComponent(_component_InputText, {
        modelValue: unref(form).email,
        "onUpdate:modelValue": ($event) => unref(form).email = $event,
        type: "email",
        disabled: "",
        class: "w-full !bg-surface-container !cursor-not-allowed"
      }, null, _parent));
      _push(`</div><div><label class="block text-sm font-medium text-on_surface mb-2">Phone</label>`);
      _push(ssrRenderComponent(_component_InputText, {
        modelValue: unref(form).phone,
        "onUpdate:modelValue": ($event) => unref(form).phone = $event,
        type: "tel",
        class: "w-full"
      }, null, _parent));
      _push(`</div>`);
      _push(ssrRenderComponent(_component_Button, {
        type: "submit",
        loading: unref(loading),
        label: unref(loading) ? "Saving..." : "Save Changes"
      }, null, _parent));
      _push(`</form></div></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/user/profile.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
export {
  _sfc_main as default
};
//# sourceMappingURL=profile-BwggFFZn.js.map
