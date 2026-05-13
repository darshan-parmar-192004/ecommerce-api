import { ref, mergeProps, unref, withCtx, createVNode, useSSRContext } from "vue";
import "hookable";
import script$2 from "./index-BGLbqIM8.js";
import script$1 from "./index-Q5CgsuHz.js";
import script from "./index-D2BzYgnQ.js";
import { ssrRenderAttrs, ssrRenderList, ssrRenderComponent } from "vue/server-renderer";
import "klona";
import "../server.mjs";
import "@primeuix/utils";
import "@primeuix/styles/tag";
import "@primeuix/utils/dom";
import "@primeuix/utils/object";
import "./index-Bg3yUxEP.js";
import "./index-BGjNF9Dd.js";
import "@primeuix/styles/paginator";
import "./index-mL8pBhlS.js";
import "./index-BdDxcS0p.js";
import "@primeuix/utils/zindex";
import "./index-D2H2HXHf.js";
import "./index-D4owPo0H.js";
import "./index-D255zs2l.js";
import "./index-83EVY8bO.js";
import "@primeuix/styles/iconfield";
import "./index-IyD_XvNl.js";
import "./index-DDXXjm8b.js";
import "./index-Dedjwyps.js";
import "./index-DrzwqwtL.js";
import "@primeuix/styles/inputtext";
import "./index-rAVNvoJo.js";
import "@primeuix/utils/eventbus";
import "./index-C3Ej0YBE.js";
import "@primeuix/styles/virtualscroller";
import "@primeuix/styles/select";
import "./index-CMnz8U93.js";
import "./index-B6Gs2AMF.js";
import "./index-l8XG4jwb.js";
import "@primeuix/styles/inputnumber";
import "./index-BsfLo0vt.js";
import "@primeuix/styles/datatable";
import "./index-DrH3FigG.js";
import "./index-Dik5Ht0d.js";
import "./index-Ccf11gW7.js";
import "./index-BkCm0XJl.js";
import "@primeuix/styles/checkbox";
import "./index-bmcwiX6i.js";
import "@primeuix/styles/radiobutton";
import "./index-CvVdVfGD.js";
import "./index-DzyF1dFx.js";
import "./index-0q6EZZjm.js";
import "ofetch";
import "#internal/nuxt/paths";
import "unctx";
import "h3";
import "ufo";
import "unhead";
import "@unhead/shared";
import "vue-router";
import "radix3";
import "defu";
import "@primeuix/styled";
import "@primeuix/styles/base";
import "cookie-es";
import "destr";
import "ohash";
import "@primeuix/styles/toast";
import "@primeuix/utils/uuid";
import "@primeuix/styles/ripple";
import "@primeuix/styles/badge";
import "@primeuix/styles/button";
const _sfc_main = {
  __name: "inventory",
  __ssrInlineRender: true,
  setup(__props) {
    const inventory = ref([]);
    const loading = ref(true);
    return (_ctx, _push, _parent, _attrs) => {
      const _component_DataTable = script;
      const _component_Column = script$1;
      const _component_Tag = script$2;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "p-6" }, _attrs))}><h1 class="text-3xl font-bold text-on_surface mb-8">Inventory Management</h1>`);
      if (unref(loading)) {
        _push(`<div class="space-y-3"><!--[-->`);
        ssrRenderList(5, (n) => {
          _push(`<div class="h-16 bg-surface-container rounded-xl animate-pulse"></div>`);
        });
        _push(`<!--]--></div>`);
      } else {
        _push(`<div class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 overflow-hidden">`);
        _push(ssrRenderComponent(_component_DataTable, {
          value: unref(inventory),
          stripedRows: "",
          class: "w-full"
        }, {
          default: withCtx((_, _push2, _parent2, _scopeId) => {
            if (_push2) {
              _push2(ssrRenderComponent(_component_Column, {
                field: "product_name",
                header: "Product Name"
              }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, {
                field: "product_id",
                header: "Product ID"
              }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, {
                field: "warehouse_id",
                header: "Warehouse"
              }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, {
                field: "quantity",
                header: "Current Stock"
              }, null, _parent2, _scopeId));
              _push2(ssrRenderComponent(_component_Column, { header: "Status" }, {
                body: withCtx(({ data }, _push3, _parent3, _scopeId2) => {
                  if (_push3) {
                    _push3(ssrRenderComponent(_component_Tag, {
                      value: data.quantity > 10 ? "In Stock" : data.quantity > 0 ? "Low Stock" : "Out of Stock",
                      severity: data.quantity > 10 ? "success" : data.quantity > 0 ? "warn" : "danger"
                    }, null, _parent3, _scopeId2));
                  } else {
                    return [
                      createVNode(_component_Tag, {
                        value: data.quantity > 10 ? "In Stock" : data.quantity > 0 ? "Low Stock" : "Out of Stock",
                        severity: data.quantity > 10 ? "success" : data.quantity > 0 ? "warn" : "danger"
                      }, null, 8, ["value", "severity"])
                    ];
                  }
                }),
                _: 1
              }, _parent2, _scopeId));
            } else {
              return [
                createVNode(_component_Column, {
                  field: "product_name",
                  header: "Product Name"
                }),
                createVNode(_component_Column, {
                  field: "product_id",
                  header: "Product ID"
                }),
                createVNode(_component_Column, {
                  field: "warehouse_id",
                  header: "Warehouse"
                }),
                createVNode(_component_Column, {
                  field: "quantity",
                  header: "Current Stock"
                }),
                createVNode(_component_Column, { header: "Status" }, {
                  body: withCtx(({ data }) => [
                    createVNode(_component_Tag, {
                      value: data.quantity > 10 ? "In Stock" : data.quantity > 0 ? "Low Stock" : "Out of Stock",
                      severity: data.quantity > 10 ? "success" : data.quantity > 0 ? "warn" : "danger"
                    }, null, 8, ["value", "severity"])
                  ]),
                  _: 1
                })
              ];
            }
          }),
          _: 1
        }, _parent));
        _push(`</div>`);
      }
      _push(`</div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/admin/inventory.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
export {
  _sfc_main as default
};
//# sourceMappingURL=inventory-uyq2qMQL.js.map
