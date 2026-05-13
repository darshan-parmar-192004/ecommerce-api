import { uuid } from "@primeuix/utils";
import script$2 from "./index-D7DvHjf4.js";
import { B as BaseStyle, t as script$3, I as DynamicDialogEventBus } from "../server.mjs";
import { resolveComponent, openBlock, createElementBlock, Fragment, renderList, createBlock, mergeProps, createSlots, withCtx, resolveDynamicComponent } from "vue";
import "@primeuix/utils/dom";
import "@primeuix/utils/zindex";
import "./index-DzyF1dFx.js";
import "@primeuix/utils/object";
import "./index-Din928lO.js";
import "@primeuix/styled";
import "@primeuix/styles/dialog";
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
import "@primeuix/styles/base";
import "cookie-es";
import "destr";
import "ohash";
import "vue/server-renderer";
import "@primeuix/styles/toast";
import "@primeuix/utils/uuid";
import "@primeuix/styles/ripple";
import "@primeuix/styles/badge";
import "@primeuix/styles/button";
var DynamicDialogStyle = BaseStyle.extend({
  name: "dynamicdialog"
});
var script$1 = {
  name: "BaseDynamicDialog",
  "extends": script$3,
  props: {},
  style: DynamicDialogStyle,
  provide: function provide() {
    return {
      $pcDynamicDialog: this,
      $parentInstance: this
    };
  }
};
var script = {
  name: "DynamicDialog",
  "extends": script$1,
  inheritAttrs: false,
  data: function data() {
    return {
      instanceMap: {}
    };
  },
  openListener: null,
  closeListener: null,
  currentInstance: null,
  mounted: function mounted() {
    var _this = this;
    this.openListener = function(_ref) {
      var instance = _ref.instance;
      var key = uuid() + "_dynamic_dialog";
      instance.visible = true;
      instance.key = key;
      _this.instanceMap[key] = instance;
    };
    this.closeListener = function(_ref2) {
      var instance = _ref2.instance, params = _ref2.params;
      var key = instance.key;
      var currentInstance = _this.instanceMap[key];
      if (currentInstance) {
        currentInstance.visible = false;
        currentInstance.options.onClose && currentInstance.options.onClose({
          data: params,
          type: "config-close"
        });
        _this.currentInstance = currentInstance;
      }
    };
    DynamicDialogEventBus.on("open", this.openListener);
    DynamicDialogEventBus.on("close", this.closeListener);
  },
  beforeUnmount: function beforeUnmount() {
    DynamicDialogEventBus.off("open", this.openListener);
    DynamicDialogEventBus.off("close", this.closeListener);
  },
  methods: {
    onDialogHide: function onDialogHide(instance) {
      !this.currentInstance && instance.options.onClose && instance.options.onClose({
        type: "dialog-close"
      });
    },
    onDialogAfterHide: function onDialogAfterHide(instance) {
      this.currentInstance && delete this.currentInstance;
      this.currentInstance = null;
      delete this.instanceMap[instance.key];
    },
    getTemplateItems: function getTemplateItems(template) {
      return Array.isArray(template) ? template : [template];
    }
  },
  components: {
    DDialog: script$2
  }
};
function render(_ctx, _cache, $props, $setup, $data, $options) {
  var _component_DDialog = resolveComponent("DDialog");
  return openBlock(true), createElementBlock(Fragment, null, renderList($data.instanceMap, function(instance, key) {
    return openBlock(), createBlock(_component_DDialog, mergeProps({
      key,
      visible: instance.visible,
      "onUpdate:visible": function onUpdateVisible($event) {
        return instance.visible = $event;
      },
      _instance: instance
    }, {
      ref_for: true
    }, instance.options.props, {
      onHide: function onHide($event) {
        return $options.onDialogHide(instance);
      },
      onAfterHide: function onAfterHide($event) {
        return $options.onDialogAfterHide(instance);
      }
    }), createSlots({
      "default": withCtx(function() {
        return [(openBlock(), createBlock(resolveDynamicComponent(instance.content), mergeProps({
          ref_for: true
        }, instance.options.emits), null, 16))];
      }),
      _: 2
    }, [instance.options.templates && instance.options.templates.header ? {
      name: "header",
      fn: withCtx(function() {
        return [(openBlock(true), createElementBlock(Fragment, null, renderList($options.getTemplateItems(instance.options.templates.header), function(header, index) {
          return openBlock(), createBlock(resolveDynamicComponent(header), mergeProps({
            key: index + "_header"
          }, {
            ref_for: true
          }, instance.options.emits), null, 16);
        }), 128))];
      }),
      key: "0"
    } : void 0, instance.options.templates && instance.options.templates.footer ? {
      name: "footer",
      fn: withCtx(function() {
        return [(openBlock(true), createElementBlock(Fragment, null, renderList($options.getTemplateItems(instance.options.templates.footer), function(footer, index) {
          return openBlock(), createBlock(resolveDynamicComponent(footer), mergeProps({
            key: index + "_footer"
          }, {
            ref_for: true
          }, instance.options.emits), null, 16);
        }), 128))];
      }),
      key: "1"
    } : void 0]), 1040, ["visible", "onUpdate:visible", "_instance", "onHide", "onAfterHide"]);
  }), 128);
}
script.render = render;
export {
  script as default
};
//# sourceMappingURL=index-DkQNumh6.js.map
