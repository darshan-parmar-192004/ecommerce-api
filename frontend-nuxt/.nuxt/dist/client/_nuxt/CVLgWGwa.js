import{_ as r,aF as n,ah as o,d as s,c as l,aj as d,g as i,ak as a}from"./-e0yDQ9r.js";var p=`
    .p-overlaybadge {
        position: relative;
    }

    .p-overlaybadge .p-badge {
        position: absolute;
        inset-block-start: 0;
        inset-inline-end: 0;
        transform: translate(50%, -50%);
        transform-origin: 100% 0;
        margin: 0;
        outline-width: dt('overlaybadge.outline.width');
        outline-style: solid;
        outline-color: dt('overlaybadge.outline.color');
    }

    .p-overlaybadge .p-badge:dir(rtl) {
        transform: translate(-50%, -50%);
    }
`,v={root:"p-overlaybadge"},g=r.extend({name:"overlaybadge",style:p,classes:v}),c={name:"OverlayBadge",extends:n,style:g,provide:function(){return{$pcOverlayBadge:this,$parentInstance:this}}},y={name:"OverlayBadge",extends:c,inheritAttrs:!1,components:{Badge:n}};function m(e,u,b,B,$,f){var t=o("Badge");return s(),l("div",a({class:e.cx("root")},e.ptmi("root")),[d(e.$slots,"default"),i(t,a(e.$props,{pt:e.ptm("pcBadge")}),null,16,["pt"])],16)}y.render=m;export{y as default};
