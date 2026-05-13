import { _ as __nuxt_component_0 } from "./nuxt-link-1RlcU5D7.js";
import { ref, mergeProps, unref, withCtx, createVNode, createTextVNode, toDisplayString, useSSRContext } from "vue";
import "hookable";
import { u as useRoute, a as useSeoMeta, s as script$4 } from "../server.mjs";
import { u as useAppToast } from "./useAppToast-CxoPk-aP.js";
import { u as useApi } from "./useApi-CR6WE1xi.js";
import script$3 from "./index-DDXXjm8b.js";
import script$2 from "./index-IyD_XvNl.js";
import script$1 from "./index-83EVY8bO.js";
import script from "./index-BAz9mnEE.js";
import { ssrRenderAttrs, ssrRenderComponent, ssrInterpolate } from "vue/server-renderer";
import { _ as _sfc_main$1 } from "./RedirectScreen-mve57wzr.js";
import { useForm } from "@vuehookform/core";
import * as z from "zod";
import "ufo";
import "ofetch";
import "#internal/nuxt/paths";
import "unctx";
import "h3";
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
import "@primeuix/styles/iconfield";
import "@primeuix/styles/message";
const _sfc_main = {
  __name: "login",
  __ssrInlineRender: true,
  setup(__props) {
    const route = useRoute();
    useAppToast();
    const loginSchema = z.object({
      email: z.string().email("Please enter a valid email address"),
      password: z.string().min(1, "Password is required")
    });
    const form = useForm({
      schema: loginSchema,
      defaultValues: {
        email: "",
        password: ""
      }
    });
    const { handleSubmit, reset, setError, watch, setValue, trigger } = form;
    const loading = ref(false);
    const formError = ref("");
    const success = ref(false);
    const redirectUrl = ref(route.query.redirect || "/");
    useApi();
    useSeoMeta({
      title: "Login - E-Commerce Store"
    });
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      const _component_Message = script;
      const _component_IconField = script$1;
      const _component_InputIcon = script$2;
      const _component_InputText = script$3;
      const _component_Button = script$4;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "w-full max-w-md" }, _attrs))}>`);
      if (!unref(success)) {
        _push(`<div class="text-center mb-10">`);
        _push(ssrRenderComponent(_component_NuxtLink, {
          to: "/",
          class: "inline-block"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(`<span class="text-3xl font-bold font-display tracking-tight text-primary"${_scopeId}> The Curator </span>`);
            } else {
              return [
                createVNode("span", { class: "text-3xl font-bold font-display tracking-tight text-primary" }, " The Curator ")
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`<h2 class="mt-8 text-xl font-semibold text-on_surface font-display"> Welcome back to your collection. </h2></div>`);
      } else {
        _push(`<!---->`);
      }
      if (unref(success)) {
        _push(ssrRenderComponent(_sfc_main$1, {
          message: "Welcome back!",
          subtitle: "You are being redirected...",
          "redirect-to": unref(redirectUrl)
        }, null, _parent));
      } else {
        _push(`<div><div class="rounded-2xl p-8 relative overflow-hidden bg-surface-container-lowest shadow-ambient"><form class="space-y-5">`);
        if (unref(formError)) {
          _push(ssrRenderComponent(_component_Message, {
            severity: "error",
            closable: false,
            class: "!mb-0"
          }, {
            default: withCtx((_, _push2, _parent2, _scopeId) => {
              if (_push2) {
                _push2(`<i class="pi pi-exclamation-triangle mr-2"${_scopeId}></i> ${ssrInterpolate(unref(formError))}`);
              } else {
                return [
                  createVNode("i", { class: "pi pi-exclamation-triangle mr-2" }),
                  createTextVNode(" " + toDisplayString(unref(formError)), 1)
                ];
              }
            }),
            _: 1
          }, _parent));
        } else {
          _push(`<!---->`);
        }
        _push(`<div><label for="email" class="block text-xs font-semibold uppercase tracking-wider mb-3 text-on_surface_variant"> Email Address </label>`);
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
                class: "w-full !pl-10",
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
                  class: "w-full !pl-10",
                  placeholder: "you@example.com"
                }, null, 8, ["model-value", "onUpdate:modelValue", "onBlur"])
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div><div><div class="flex items-center justify-between mb-3"><label for="password" class="block text-xs font-semibold uppercase tracking-wider text-on_surface_variant"> Password </label>`);
        _push(ssrRenderComponent(_component_NuxtLink, {
          to: "#",
          class: "text-xs font-medium transition-colors hover:text-primary text-primary"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(` Forgot Password `);
            } else {
              return [
                createTextVNode(" Forgot Password ")
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div>`);
        _push(ssrRenderComponent(_component_IconField, null, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_InputIcon, { class: "pi pi-lock" }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_InputText, {
                id: "password",
                "model-value": unref(watch)("password"),
                "onUpdate:modelValue": ($event) => unref(setValue)("password", $event),
                onBlur: ($event) => unref(trigger)("password"),
                type: "password",
                required: "",
                class: "w-full !pl-10",
                placeholder: "Enter your password"
              }, null, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_InputIcon, { class: "pi pi-lock" }),
                createVNode(_component_InputText, {
                  id: "password",
                  "model-value": unref(watch)("password"),
                  "onUpdate:modelValue": ($event) => unref(setValue)("password", $event),
                  onBlur: ($event) => unref(trigger)("password"),
                  type: "password",
                  required: "",
                  class: "w-full !pl-10",
                  placeholder: "Enter your password"
                }, null, 8, ["model-value", "onUpdate:modelValue", "onBlur"])
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div>`);
        _push(ssrRenderComponent(_component_Button, {
          type: "submit",
          loading: unref(loading),
          label: unref(loading) ? "Signing in..." : "Sign In",
          class: "w-full !py-3"
        }, null, _parent));
        _push(`<p class="text-center pt-2 text-outline"><span class="text-sm">Don&#39;t have an account?</span>`);
        _push(ssrRenderComponent(_component_NuxtLink, {
          to: "/auth/register",
          onClick: () => {
          },
          class: "text-sm font-semibold ml-1 transition-colors hover:text-primary text-primary"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(` Create Account `);
            } else {
              return [
                createTextVNode(" Create Account ")
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</p></form></div><div class="text-center mt-8"><p class="text-xs uppercase tracking-widest text-outline"> Curated by Editorial Team </p></div></div>`);
      }
      _push(`</div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/auth/login.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
export {
  _sfc_main as default
};
//# sourceMappingURL=login-CBO8sI1E.js.map
