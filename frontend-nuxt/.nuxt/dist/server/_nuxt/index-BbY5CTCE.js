import { uuid } from "@primeuix/utils";
import { s as script$2 } from "./index-DrzwqwtL.js";
import { style } from "@primeuix/styles/checkboxgroup";
import { B as BaseStyle } from "../server.mjs";
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
  root: "p-checkbox-group p-component"
};
var CheckboxGroupStyle = BaseStyle.extend({
  name: "checkboxgroup",
  style,
  classes
});
var script$1 = {
  name: "BaseCheckboxGroup",
  "extends": script$2,
  style: CheckboxGroupStyle,
  provide: function provide() {
    return {
      $pcCheckboxGroup: this,
      $parentInstance: this
    };
  }
};
var script = {
  name: "CheckboxGroup",
  "extends": script$1,
  inheritAttrs: false,
  data: function data() {
    return {
      groupName: this.name
    };
  },
  watch: {
    name: function name(newValue) {
      this.groupName = newValue || uuid("checkbox-group-");
    }
  },
  mounted: function mounted() {
    this.groupName = this.groupName || uuid("checkbox-group-");
  }
};
function render(_ctx, _cache, $props, $setup, $data, $options) {
  return openBlock(), createElementBlock("div", mergeProps({
    "class": _ctx.cx("root")
  }, _ctx.ptmi("root")), [renderSlot(_ctx.$slots, "default")], 16);
}
script.render = render;
export {
  script as default
};
//# sourceMappingURL=index-BbY5CTCE.js.map
