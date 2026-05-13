import { addClass, hasCSSAnimation, createElement } from "@primeuix/utils/dom";
import { ZIndex } from "@primeuix/utils/zindex";
import { u as unblockBodyScroll, b as blockBodyScroll } from "./index-Din928lO.js";
import { B as BaseStyle, t as script$2 } from "../server.mjs";
import { style } from "@primeuix/styles/blockui";
import { openBlock, createElementBlock, mergeProps, renderSlot } from "vue";
import "@primeuix/styled";
import "@primeuix/utils";
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
import "@primeuix/utils/object";
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
var classes = {
  root: "p-blockui"
};
var BlockUIStyle = BaseStyle.extend({
  name: "blockui",
  style,
  classes
});
var script$1 = {
  name: "BaseBlockUI",
  "extends": script$2,
  props: {
    blocked: {
      type: Boolean,
      "default": false
    },
    fullScreen: {
      type: Boolean,
      "default": false
    },
    baseZIndex: {
      type: Number,
      "default": 0
    },
    autoZIndex: {
      type: Boolean,
      "default": true
    }
  },
  style: BlockUIStyle,
  provide: function provide() {
    return {
      $pcBlockUI: this,
      $parentInstance: this
    };
  }
};
var script = {
  name: "BlockUI",
  "extends": script$1,
  inheritAttrs: false,
  emits: ["block", "unblock"],
  mask: null,
  data: function data() {
    return {
      isBlocked: false
    };
  },
  watch: {
    blocked: function blocked(newValue) {
      if (newValue === true) this.block();
      else this.unblock();
    }
  },
  mounted: function mounted() {
    if (this.blocked) {
      this.block();
    }
  },
  methods: {
    block: function block() {
      var styleClass = "p-blockui-mask p-overlay-mask p-overlay-mask-enter-active";
      if (this.fullScreen) {
        styleClass += " p-blockui-mask-document";
        this.mask = createElement("div", {
          style: {
            position: "fixed",
            top: "0",
            left: "0",
            width: "100%",
            height: "100%"
          },
          "class": !this.isUnstyled && styleClass,
          "p-bind": this.ptm("mask")
        });
        (void 0).body.appendChild(this.mask);
        blockBodyScroll();
        (void 0).activeElement.blur();
      } else {
        this.mask = createElement("div", {
          style: {
            position: "absolute",
            top: "0",
            left: "0",
            width: "100%",
            height: "100%"
          },
          "class": !this.isUnstyled && styleClass,
          "p-bind": this.ptm("mask")
        });
        this.$refs.container.appendChild(this.mask);
      }
      if (this.autoZIndex) {
        ZIndex.set("modal", this.mask, this.baseZIndex || this.$primevue.config.zIndex.modal);
      }
      this.isBlocked = true;
      this.$emit("block");
    },
    unblock: function unblock() {
      var _this = this;
      if (this.mask) {
        !this.isUnstyled && addClass(this.mask, "p-overlay-mask-leave-active");
        var _handleAnimationEnd = function handleAnimationEnd() {
          clearTimeout(fallbackTimer);
          _this.mask.removeEventListener("animationend", _handleAnimationEnd);
          _this.mask.removeEventListener("webkitAnimationEnd", _handleAnimationEnd);
        };
        var fallbackTimer = setTimeout(function() {
          _this.removeMask();
        }, 300);
        if (hasCSSAnimation(this.mask) > 0) {
          this.mask.addEventListener("animationend", _handleAnimationEnd);
          this.mask.addEventListener("webkitAnimationEnd", _handleAnimationEnd);
        }
      } else {
        this.removeMask();
      }
    },
    removeMask: function removeMask() {
      ZIndex.clear(this.mask);
      if (this.fullScreen) {
        (void 0).body.removeChild(this.mask);
        unblockBodyScroll();
      } else {
        var _this$$refs$container;
        (_this$$refs$container = this.$refs.container) === null || _this$$refs$container === void 0 || _this$$refs$container.removeChild(this.mask);
      }
      this.isBlocked = false;
      this.$emit("unblock");
    }
  }
};
var _hoisted_1 = ["aria-busy"];
function render(_ctx, _cache, $props, $setup, $data, $options) {
  return openBlock(), createElementBlock("div", mergeProps({
    ref: "container",
    "class": _ctx.cx("root"),
    "aria-busy": $data.isBlocked
  }, _ctx.ptmi("root")), [renderSlot(_ctx.$slots, "default")], 16, _hoisted_1);
}
script.render = render;
export {
  script as default
};
//# sourceMappingURL=index-SLtcaRRk.js.map
