import { reactive, ref, mergeProps, unref, withCtx, createTextVNode, useSSRContext } from 'vue';
import { d as useAuthStore, s as script$2, k as defineStore } from './server.mjs';
import script$1 from './index-DDXXjm8b.mjs';
import script from './index-BAz9mnEE.mjs';
import { ssrRenderAttrs, ssrRenderComponent } from 'vue/server-renderer';
import { u as useApi } from './useApi-CR6WE1xi.mjs';
import '../nitro/nitro.mjs';
import 'node:http';
import 'node:https';
import 'node:events';
import 'node:buffer';
import 'node:fs';
import 'node:path';
import 'node:crypto';
import '@primevue/core/base/style';
import '@primevue/core/basecomponent/style';
import '@primeuix/styles/autocomplete';
import '@primeuix/utils/object';
import '@primeuix/styles/cascadeselect';
import '@primeuix/styles/checkbox';
import '@primeuix/styles/checkboxgroup';
import '@primeuix/styles/colorpicker';
import '@primeuix/styles/datepicker';
import '@primeuix/styles/floatlabel';
import '@primeuix/styles/iconfield';
import '@primeuix/styles/iftalabel';
import '@primeuix/styles/inputchips';
import '@primeuix/styles/inputgroup';
import '@primeuix/styles/inputnumber';
import '@primeuix/styles/inputotp';
import '@primeuix/styles/inputtext';
import '@primeuix/styles/knob';
import '@primeuix/styles/listbox';
import '@primeuix/styles/multiselect';
import '@primeuix/styles/password';
import '@primeuix/styles/radiobutton';
import '@primeuix/styles/radiobuttongroup';
import '@primeuix/styles/rating';
import '@primeuix/styles/select';
import '@primeuix/styles/selectbutton';
import '@primeuix/styles/slider';
import '@primeuix/styles/textarea';
import '@primeuix/styles/togglebutton';
import '@primeuix/styles/toggleswitch';
import '@primeuix/styles/treeselect';
import '@primeuix/styles/button';
import '@primeuix/styles/buttongroup';
import '@primeuix/styles/speeddial';
import '@primeuix/styles/splitbutton';
import '@primeuix/styles/datatable';
import '@primeuix/styles/dataview';
import '@primeuix/styles/orderlist';
import '@primeuix/styles/organizationchart';
import '@primeuix/styles/paginator';
import '@primeuix/styles/picklist';
import '@primeuix/styles/tree';
import '@primeuix/styles/treetable';
import '@primeuix/styles/timeline';
import '@primeuix/styles/virtualscroller';
import '@primeuix/styles/accordion';
import '@primeuix/styles/card';
import '@primeuix/styles/divider';
import '@primeuix/styles/fieldset';
import '@primeuix/styles/panel';
import '@primeuix/styles/scrollpanel';
import '@primeuix/styles/splitter';
import '@primeuix/styles/stepper';
import '@primeuix/styles/tabview';
import '@primeuix/styles/tabs';
import '@primeuix/styles/toolbar';
import '@primeuix/styles/confirmdialog';
import '@primeuix/styles/confirmpopup';
import '@primeuix/styles/dialog';
import '@primeuix/styles/drawer';
import '@primeuix/styles/popover';
import '@primeuix/styles/fileupload';
import '@primeuix/styles/breadcrumb';
import '@primeuix/styles/contextmenu';
import '@primeuix/styles/dock';
import '@primeuix/styles/menu';
import '@primeuix/styles/menubar';
import '@primeuix/styles/megamenu';
import '@primeuix/styles/panelmenu';
import '@primeuix/styles/steps';
import '@primeuix/styles/tabmenu';
import '@primeuix/styles/tieredmenu';
import '@primeuix/styles/message';
import '@primeuix/styles/inlinemessage';
import '@primeuix/styles/toast';
import '@primeuix/styles/carousel';
import '@primeuix/styles/galleria';
import '@primeuix/styles/image';
import '@primeuix/styles/imagecompare';
import '@primeuix/styles/avatar';
import '@primeuix/styles/badge';
import '@primeuix/styles/blockui';
import '@primeuix/styles/chip';
import '@primeuix/styles/inplace';
import '@primeuix/styles/metergroup';
import '@primeuix/styles/overlaybadge';
import '@primeuix/styles/scrolltop';
import '@primeuix/styles/skeleton';
import '@primeuix/styles/progressbar';
import '@primeuix/styles/progressspinner';
import '@primeuix/styles/tag';
import '@primeuix/styles/terminal';
import '@primevue/forms/form/style';
import '@primevue/forms/formfield/style';
import '@primeuix/styles/tooltip';
import '@primeuix/styles/ripple';
import '@primeuix/styled';
import 'unhead';
import '@unhead/shared';
import 'vue-router';
import '@primeuix/utils/eventbus';
import '@primeuix/utils';
import '@primeuix/styles/base';
import '@primeuix/utils/dom';
import '@primeuix/utils/zindex';
import '@primeuix/utils/uuid';
import './index-Dedjwyps.mjs';
import './index-DrzwqwtL.mjs';

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

export { _sfc_main as default };
//# sourceMappingURL=profile-BwggFFZn.mjs.map
