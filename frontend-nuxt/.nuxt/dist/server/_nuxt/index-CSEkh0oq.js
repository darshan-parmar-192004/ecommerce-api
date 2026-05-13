import { B as BaseStyle, t as script$2 } from "../server.mjs";
import { renderSlot } from "vue";
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
var AccordionTabStyle = BaseStyle.extend({
  name: "accordiontab"
});
var script$1 = {
  name: "BaseAccordionTab",
  "extends": script$2,
  props: {
    header: null,
    headerStyle: null,
    headerClass: null,
    headerProps: null,
    headerActionProps: null,
    contentStyle: null,
    contentClass: null,
    contentProps: null,
    disabled: Boolean
  },
  style: AccordionTabStyle,
  provide: function provide() {
    return {
      $pcAccordionTab: this,
      $parentInstance: this
    };
  }
};
var script = {
  name: "AccordionTab",
  "extends": script$1,
  inheritAttrs: false,
  mounted: function mounted() {
    console.warn("Deprecated since v4. Use the new structure of Accordion instead.");
  }
};
function render(_ctx, _cache, $props, $setup, $data, $options) {
  return renderSlot(_ctx.$slots, "default");
}
script.render = render;
export {
  script as default
};
//# sourceMappingURL=index-CSEkh0oq.js.map
