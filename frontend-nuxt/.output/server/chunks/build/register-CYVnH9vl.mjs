import { _ as __nuxt_component_0 } from './nuxt-link-1RlcU5D7.mjs';
import { ref, computed, mergeProps, unref, withCtx, createVNode, createTextVNode, toDisplayString, useSSRContext } from 'vue';
import { u as useAppToast } from './useAppToast-CxoPk-aP.mjs';
import { u as useApi } from './useApi-CR6WE1xi.mjs';
import { a as useSeoMeta, s as script$4 } from './server.mjs';
import script$3 from './index-DDXXjm8b.mjs';
import script$2 from './index-IyD_XvNl.mjs';
import script$1 from './index-83EVY8bO.mjs';
import script from './index-BAz9mnEE.mjs';
import { ssrRenderAttrs, ssrRenderComponent, ssrInterpolate, ssrRenderClass } from 'vue/server-renderer';
import { useForm } from '@vuehookform/core';
import * as z from 'zod';
import { _ as _sfc_main$1 } from './RedirectScreen-mve57wzr.mjs';
import { _ as _export_sfc } from './_plugin-vue_export-helper-1tPrXgE0.mjs';
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

const _sfc_main = {
  __name: "register",
  __ssrInlineRender: true,
  setup(__props) {
    useAppToast();
    const registerSchema = z.object({
      name: z.string().min(2, "Full name is required"),
      email: z.string().email("Please enter a valid email address"),
      password: z.string().min(6, "Password must be at least 6 characters")
    });
    const form = useForm({
      schema: registerSchema,
      defaultValues: {
        name: "",
        email: "",
        password: ""
      }
    });
    const { handleSubmit, reset, watch, setValue, trigger } = form;
    const showPassword = ref(false);
    const error = ref("");
    const loading = ref(false);
    const success = ref(false);
    const passwordValue = ref("");
    useApi();
    const passwordStrength = computed(() => {
      const pwd = passwordValue.value;
      if (!pwd) return { level: 0, text: "", color: "" };
      let score = 0;
      if (pwd.length >= 6) score++;
      if (pwd.length >= 8) score++;
      if (/[a-z]/.test(pwd) && /[A-Z]/.test(pwd)) score++;
      if (/\d/.test(pwd)) score++;
      if (/[^a-zA-Z0-9]/.test(pwd)) score++;
      if (score <= 2) return { level: 1, text: "Weak", color: "bg-red-500" };
      if (score <= 3) return { level: 2, text: "Medium", color: "bg-yellow-500" };
      return { level: 3, text: "Strong", color: "bg-green-500" };
    });
    const isValidEmail = computed(() => {
      const email = watch("email");
      if (!email) return null;
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      return emailRegex.test(email);
    });
    useSeoMeta({
      title: "Register - E-Commerce Store"
    });
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Message = script;
      const _component_IconField = script$1;
      const _component_InputIcon = script$2;
      const _component_InputText = script$3;
      const _component_Button = script$4;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "w-full max-w-md" }, _attrs))} data-v-4eb36d45>`);
      if (!unref(success)) {
        _push(`<div class="text-center mb-10" data-v-4eb36d45>`);
        _push(ssrRenderComponent(_component_NuxtLink, {
          to: "/",
          class: "inline-block"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(`<span class="text-3xl font-bold font-display tracking-tight text-primary" data-v-4eb36d45${_scopeId}> The Curator </span>`);
            } else {
              return [
                createVNode("span", { class: "text-3xl font-bold font-display tracking-tight text-primary" }, " The Curator ")
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`<h2 class="mt-6 text-2xl font-bold text-on_surface font-display" data-v-4eb36d45> Create your account </h2><p class="mt-2 text-on_surface_variant" data-v-4eb36d45> Join us and start shopping </p></div>`);
      } else {
        _push(`<!---->`);
      }
      if (unref(success)) {
        _push(ssrRenderComponent(_sfc_main$1, {
          message: "Registration Successful!",
          subtitle: "Your account has been created.",
          "redirect-to": "/"
        }, null, _parent));
      } else {
        _push(`<div class="rounded-2xl p-8 relative overflow-hidden bg-surface-container-lowest shadow-ambient" data-v-4eb36d45><form class="flex flex-col gap-6" data-v-4eb36d45>`);
        if (unref(error)) {
          _push(ssrRenderComponent(_component_Message, {
            severity: "error",
            closable: false
          }, {
            default: withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                _push2(`<i class="pi pi-exclamation-triangle mr-2" data-v-4eb36d45${_scopeId}></i> ${ssrInterpolate(unref(error))}`);
              } else {
                return [
                  createVNode("i", { class: "pi pi-exclamation-triangle mr-2" }),
                  createTextVNode(" " + toDisplayString(unref(error)), 1)
                ];
              }
            }),
            _: 1
          }, _parent));
        } else {
          _push(`<!---->`);
        }
        _push(`<div data-v-4eb36d45><label for="name" class="block text-xs font-semibold uppercase tracking-wider mb-3 text-on_surface_variant" data-v-4eb36d45> Full Name </label>`);
        _push(ssrRenderComponent(_component_IconField, null, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_InputIcon, { class: "pi pi-user" }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_InputText, {
                id: "name",
                "model-value": unref(watch)("name"),
                "onUpdate:modelValue": ($event) => unref(setValue)("name", $event),
                onBlur: ($event) => unref(trigger)("name"),
                type: "text",
                required: "",
                class: "w-full !pl-10",
                placeholder: "John Doe"
              }, null, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_InputIcon, { class: "pi pi-user" }),
                createVNode(_component_InputText, {
                  id: "name",
                  "model-value": unref(watch)("name"),
                  "onUpdate:modelValue": ($event) => unref(setValue)("name", $event),
                  onBlur: ($event) => unref(trigger)("name"),
                  type: "text",
                  required: "",
                  class: "w-full !pl-10",
                  placeholder: "John Doe"
                }, null, 8, ["model-value", "onUpdate:modelValue", "onBlur"])
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div><div data-v-4eb36d45><label for="email" class="block text-xs font-semibold uppercase tracking-wider mb-3 text-on_surface_variant" data-v-4eb36d45> Email Address </label>`);
        _push(ssrRenderComponent(_component_IconField, null, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_InputIcon, { class: "pi pi-envelope" }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_InputText, {
                id: "email",
                "model-value": unref(watch)("email"),
                "onUpdate:modelValue": ($event) => unref(setValue)("email", $event),
                onBlur: ($event) => unref(trigger)("email"),
                type: "email",
                required: "",
                class: unref(isValidEmail) === false ? "!border-error !bg-error-container" : "",
                placeholder: "you@example.com"
              }, null, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_InputIcon, { class: "pi pi-envelope" }),
                createVNode(_component_InputText, {
                  id: "email",
                  "model-value": unref(watch)("email"),
                  "onUpdate:modelValue": ($event) => unref(setValue)("email", $event),
                  onBlur: ($event) => unref(trigger)("email"),
                  type: "email",
                  required: "",
                  class: unref(isValidEmail) === false ? "!border-error !bg-error-container" : "",
                  placeholder: "you@example.com"
                }, null, 8, ["model-value", "onUpdate:modelValue", "onBlur", "class"])
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`<div class="min-h-[1.25rem] mt-1" data-v-4eb36d45>`);
        if (unref(isValidEmail) === false) {
          _push(`<p class="text-xs text-error" data-v-4eb36d45> Please enter a valid email address </p>`);
        } else {
          _push(`<!---->`);
        }
        _push(`</div></div><div data-v-4eb36d45><label for="password" class="block text-xs font-semibold uppercase tracking-wider mb-3 text-on_surface_variant" data-v-4eb36d45> Password </label>`);
        _push(ssrRenderComponent(_component_IconField, null, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_InputIcon, { class: "pi pi-lock" }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_InputText, {
                id: "password",
                "model-value": unref(watch)("password"),
                "onUpdate:modelValue": ($event) => {
                  unref(setValue)("password", $event);
                  unref(passwordValue).value = $event;
                },
                onBlur: ($event) => unref(trigger)("password"),
                type: unref(showPassword) ? "text" : "password",
                required: "",
                minlength: "6",
                class: "w-full !pl-10",
                placeholder: "Min 6 characters"
              }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Button, {
                type: "button",
                onClick: ($event) => showPassword.value = !unref(showPassword),
                icon: unref(showPassword) ? "pi pi-eye-slash" : "pi pi-eye",
                text: "",
                class: "absolute right-2 top-1/2 -translate-y-1/2"
              }, null, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_InputIcon, { class: "pi pi-lock" }),
                createVNode(_component_InputText, {
                  id: "password",
                  "model-value": unref(watch)("password"),
                  "onUpdate:modelValue": ($event) => {
                    unref(setValue)("password", $event);
                    unref(passwordValue).value = $event;
                  },
                  onBlur: ($event) => unref(trigger)("password"),
                  type: unref(showPassword) ? "text" : "password",
                  required: "",
                  minlength: "6",
                  class: "w-full !pl-10",
                  placeholder: "Min 6 characters"
                }, null, 8, ["model-value", "onUpdate:modelValue", "onBlur", "type"]),
                createVNode(_component_Button, {
                  type: "button",
                  onClick: ($event) => showPassword.value = !unref(showPassword),
                  icon: unref(showPassword) ? "pi pi-eye-slash" : "pi pi-eye",
                  text: "",
                  class: "absolute right-2 top-1/2 -translate-y-1/2"
                }, null, 8, ["onClick", "icon"])
              ];
            }
          }),
          _: 1
        }, _parent));
        if (unref(form).watch("password")) {
          _push(`<div class="mt-2" data-v-4eb36d45><div class="flex gap-1 mb-1" data-v-4eb36d45><div class="${ssrRenderClass(["h-1 flex-1 rounded-full transition-all duration-300", unref(passwordStrength).level >= 1 ? unref(passwordStrength).color : "bg-surface-container-high"])}" data-v-4eb36d45></div><div class="${ssrRenderClass(["h-1 flex-1 rounded-full transition-all duration-300", unref(passwordStrength).level >= 2 ? unref(passwordStrength).color : "bg-surface-container-high"])}" data-v-4eb36d45></div><div class="${ssrRenderClass(["h-1 flex-1 rounded-full transition-all duration-300", unref(passwordStrength).level >= 3 ? unref(passwordStrength).color : "bg-surface-container-high"])}" data-v-4eb36d45></div></div><p class="${ssrRenderClass(["text-xs", unref(passwordStrength).level === 1 ? "text-error" : unref(passwordStrength).level === 2 ? "text-yellow-600" : "text-green-600"])}" data-v-4eb36d45>${ssrInterpolate(unref(passwordStrength).text)}</p></div>`);
        } else {
          _push(`<!---->`);
        }
        _push(`</div>`);
        _push(ssrRenderComponent(_component_Button, {
          type: "submit",
          loading: unref(loading),
          label: unref(loading) ? "Creating account..." : "Create Account",
          class: "w-full !py-3"
        }, null, _parent));
        _push(`<p class="text-center text-on_surface_variant pt-2" data-v-4eb36d45> Already have an account? `);
        _push(ssrRenderComponent(_component_NuxtLink, {
          to: "/auth/login",
          class: "text-primary hover:text-primary/80 font-semibold"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(` Sign in `);
            } else {
              return [
                createTextVNode(" Sign in ")
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</p></form></div>`);
      }
      _push(`</div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/auth/register.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const register = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-4eb36d45"]]);

export { register as default };
//# sourceMappingURL=register-CYVnH9vl.mjs.map
