import { B as BaseStyle, t as script$2 } from "../server.mjs";
import { openBlock, createElementBlock, mergeProps, renderSlot } from "vue";
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
var classes = {
  root: function root(_ref) {
    var instance = _ref.instance;
    return ["p-stepitem", {
      "p-stepitem-active": instance.isActive
    }];
  }
};
var StepItemStyle = BaseStyle.extend({
  name: "stepitem",
  classes
});
var script$1 = {
  name: "BaseStepItem",
  "extends": script$2,
  props: {
    value: {
      type: [String, Number],
      "default": void 0
    }
  },
  style: StepItemStyle,
  provide: function provide() {
    return {
      $pcStepItem: this,
      $parentInstance: this
    };
  }
};
var script = {
  name: "StepItem",
  "extends": script$1,
  inheritAttrs: false,
  inject: ["$pcStepper"],
  computed: {
    isActive: function isActive() {
      var _this$$pcStepper;
      return ((_this$$pcStepper = this.$pcStepper) === null || _this$$pcStepper === void 0 ? void 0 : _this$$pcStepper.d_value) === this.value;
    }
  }
};
var _hoisted_1 = ["data-p-active"];
function render(_ctx, _cache, $props, $setup, $data, $options) {
  return openBlock(), createElementBlock("div", mergeProps({
    "class": _ctx.cx("root"),
    "data-p-active": $options.isActive
  }, _ctx.ptmi("root")), [renderSlot(_ctx.$slots, "default")], 16, _hoisted_1);
}
script.render = render;
export {
  script as default
};
//# sourceMappingURL=index-Cn6WLkFk.js.map
