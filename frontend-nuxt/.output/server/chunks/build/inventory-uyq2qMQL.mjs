import { ref, mergeProps, unref, withCtx, createVNode, useSSRContext } from 'vue';
import script$2 from './index-BGLbqIM8.mjs';
import script$1 from './index-Q5CgsuHz.mjs';
import script from './index-D2BzYgnQ.mjs';
import { ssrRenderAttrs, ssrRenderList, ssrRenderComponent } from 'vue/server-renderer';
import '@primeuix/utils';
import './server.mjs';
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
import '@primeuix/styles/base';
import '@primeuix/utils/dom';
import '@primeuix/utils/zindex';
import '@primeuix/utils/uuid';
import './index-Bg3yUxEP.mjs';
import './index-BGjNF9Dd.mjs';
import './index-mL8pBhlS.mjs';
import './index-BdDxcS0p.mjs';
import './index-D2H2HXHf.mjs';
import './index-D4owPo0H.mjs';
import './index-D255zs2l.mjs';
import './index-83EVY8bO.mjs';
import './index-IyD_XvNl.mjs';
import './index-DDXXjm8b.mjs';
import './index-Dedjwyps.mjs';
import './index-DrzwqwtL.mjs';
import './index-rAVNvoJo.mjs';
import './index-C3Ej0YBE.mjs';
import './index-CMnz8U93.mjs';
import './index-B6Gs2AMF.mjs';
import './index-l8XG4jwb.mjs';
import './index-BsfLo0vt.mjs';
import './index-DrH3FigG.mjs';
import './index-Dik5Ht0d.mjs';
import './index-Ccf11gW7.mjs';
import './index-BkCm0XJl.mjs';
import './index-bmcwiX6i.mjs';
import './index-CvVdVfGD.mjs';
import './index-DzyF1dFx.mjs';
import './index-0q6EZZjm.mjs';

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

export { _sfc_main as default };
//# sourceMappingURL=inventory-uyq2qMQL.mjs.map
