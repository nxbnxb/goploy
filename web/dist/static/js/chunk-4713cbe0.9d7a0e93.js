(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-4713cbe0"],{"02f4":function(e,t,r){var a=r("4588"),o=r("be13");e.exports=function(e){return function(t,r){var n,i,s=String(o(t)),l=a(r),c=s.length;return l<0||l>=c?e?"":void 0:(n=s.charCodeAt(l),n<55296||n>56319||l+1===c||(i=s.charCodeAt(l+1))<56320||i>57343?e?s.charAt(l):n:e?s.slice(l,l+2):i-56320+(n-55296<<10)+65536)}}},"0390":function(e,t,r){"use strict";var a=r("02f4")(!0);e.exports=function(e,t,r){return t+(r?a(e,t).length:1)}},"0a49":function(e,t,r){var a=r("9b43"),o=r("626a"),n=r("4bf8"),i=r("9def"),s=r("cd1c");e.exports=function(e,t){var r=1==e,l=2==e,c=3==e,u=4==e,d=6==e,f=5==e||d,p=t||s;return function(t,s,v){for(var m,b,h=n(t),g=o(h),y=a(s,v,3),S=i(g.length),_=0,j=r?p(t,S):l?p(t,0):void 0;S>_;_++)if((f||_ in g)&&(m=g[_],b=y(m,_,h),e))if(r)j[_]=b;else if(b)switch(e){case 3:return!0;case 5:return m;case 6:return _;case 2:j.push(m)}else if(u)return!1;return d?-1:c||u?u:j}}},"0bfb":function(e,t,r){"use strict";var a=r("cb7c");e.exports=function(){var e=a(this),t="";return e.global&&(t+="g"),e.ignoreCase&&(t+="i"),e.multiline&&(t+="m"),e.unicode&&(t+="u"),e.sticky&&(t+="y"),t}},1169:function(e,t,r){var a=r("2d95");e.exports=Array.isArray||function(e){return"Array"==a(e)}},"11e9":function(e,t,r){var a=r("52a7"),o=r("4630"),n=r("6821"),i=r("6a99"),s=r("69a8"),l=r("c69a"),c=Object.getOwnPropertyDescriptor;t.f=r("9e1e")?c:function(e,t){if(e=n(e),t=i(t,!0),l)try{return c(e,t)}catch(r){}if(s(e,t))return o(!a.f.call(e,t),e[t])}},"214f":function(e,t,r){"use strict";r("b0c5");var a=r("2aba"),o=r("32e9"),n=r("79e5"),i=r("be13"),s=r("2b4c"),l=r("520a"),c=s("species"),u=!n(function(){var e=/./;return e.exec=function(){var e=[];return e.groups={a:"7"},e},"7"!=="".replace(e,"$<a>")}),d=function(){var e=/(?:)/,t=e.exec;e.exec=function(){return t.apply(this,arguments)};var r="ab".split(e);return 2===r.length&&"a"===r[0]&&"b"===r[1]}();e.exports=function(e,t,r){var f=s(e),p=!n(function(){var t={};return t[f]=function(){return 7},7!=""[e](t)}),v=p?!n(function(){var t=!1,r=/a/;return r.exec=function(){return t=!0,null},"split"===e&&(r.constructor={},r.constructor[c]=function(){return r}),r[f](""),!t}):void 0;if(!p||!v||"replace"===e&&!u||"split"===e&&!d){var m=/./[f],b=r(i,f,""[e],function(e,t,r,a,o){return t.exec===l?p&&!o?{done:!0,value:m.call(t,r,a)}:{done:!0,value:e.call(r,t,a)}:{done:!1}}),h=b[0],g=b[1];a(String.prototype,e,h),o(RegExp.prototype,f,2==t?function(e,t){return g.call(e,this,t)}:function(e){return g.call(e,this)})}}},"28a5":function(e,t,r){"use strict";var a=r("aae3"),o=r("cb7c"),n=r("ebd6"),i=r("0390"),s=r("9def"),l=r("5f1b"),c=r("520a"),u=r("79e5"),d=Math.min,f=[].push,p="split",v="length",m="lastIndex",b=4294967295,h=!u(function(){RegExp(b,"y")});r("214f")("split",2,function(e,t,r,u){var g;return g="c"=="abbc"[p](/(b)*/)[1]||4!="test"[p](/(?:)/,-1)[v]||2!="ab"[p](/(?:ab)*/)[v]||4!="."[p](/(.?)(.?)/)[v]||"."[p](/()()/)[v]>1||""[p](/.?/)[v]?function(e,t){var o=String(this);if(void 0===e&&0===t)return[];if(!a(e))return r.call(o,e,t);var n,i,s,l=[],u=(e.ignoreCase?"i":"")+(e.multiline?"m":"")+(e.unicode?"u":"")+(e.sticky?"y":""),d=0,p=void 0===t?b:t>>>0,h=new RegExp(e.source,u+"g");while(n=c.call(h,o)){if(i=h[m],i>d&&(l.push(o.slice(d,n.index)),n[v]>1&&n.index<o[v]&&f.apply(l,n.slice(1)),s=n[0][v],d=i,l[v]>=p))break;h[m]===n.index&&h[m]++}return d===o[v]?!s&&h.test("")||l.push(""):l.push(o.slice(d)),l[v]>p?l.slice(0,p):l}:"0"[p](void 0,0)[v]?function(e,t){return void 0===e&&0===t?[]:r.call(this,e,t)}:r,[function(r,a){var o=e(this),n=void 0==r?void 0:r[t];return void 0!==n?n.call(r,o,a):g.call(String(o),r,a)},function(e,t){var a=u(g,e,this,t,g!==r);if(a.done)return a.value;var c=o(e),f=String(this),p=n(c,RegExp),v=c.unicode,m=(c.ignoreCase?"i":"")+(c.multiline?"m":"")+(c.unicode?"u":"")+(h?"y":"g"),y=new p(h?c:"^(?:"+c.source+")",m),S=void 0===t?b:t>>>0;if(0===S)return[];if(0===f.length)return null===l(y,f)?[f]:[];var _=0,j=0,x=[];while(j<f.length){y.lastIndex=h?j:0;var D,w=l(y,h?f:f.slice(j));if(null===w||(D=d(s(y.lastIndex+(h?0:j)),f.length))===_)j=i(f,j,v);else{if(x.push(f.slice(_,j)),x.length===S)return x;for(var I=1;I<=w.length-1;I++)if(x.push(w[I]),x.length===S)return x;j=_=D}}return x.push(f.slice(_)),x}]})},3846:function(e,t,r){r("9e1e")&&"g"!=/./g.flags&&r("86cc").f(RegExp.prototype,"flags",{configurable:!0,get:r("0bfb")})},"4c08":function(e,t,r){"use strict";r.d(t,"c",function(){return o}),r.d(t,"d",function(){return n}),r.d(t,"a",function(){return i}),r.d(t,"b",function(){return s}),r.d(t,"e",function(){return l});var a=r("b775");function o(){return Object(a["a"])({url:"/projectGroup/getList",method:"get",params:{}})}function n(){return Object(a["a"])({url:"/projectGroup/getOption",method:"get"})}function i(e){return Object(a["a"])({url:"/projectGroup/add",method:"post",data:e})}function s(e){return Object(a["a"])({url:"/projectGroup/edit",method:"post",data:e})}function l(e){return Object(a["a"])({url:"/projectGroup/remove",method:"post",data:{id:e}})}},"520a":function(e,t,r){"use strict";var a=r("0bfb"),o=RegExp.prototype.exec,n=String.prototype.replace,i=o,s="lastIndex",l=function(){var e=/a/,t=/b*/g;return o.call(e,"a"),o.call(t,"a"),0!==e[s]||0!==t[s]}(),c=void 0!==/()??/.exec("")[1],u=l||c;u&&(i=function(e){var t,r,i,u,d=this;return c&&(r=new RegExp("^"+d.source+"$(?!\\s)",a.call(d))),l&&(t=d[s]),i=o.call(d,e),l&&i&&(d[s]=d.global?i.index+i[0].length:t),c&&i&&i.length>1&&n.call(i[0],r,function(){for(u=1;u<arguments.length-2;u++)void 0===arguments[u]&&(i[u]=void 0)}),i}),e.exports=i},"5d58":function(e,t,r){e.exports=r("d8d6")},"5dbc":function(e,t,r){var a=r("d3f4"),o=r("8b97").set;e.exports=function(e,t,r){var n,i=t.constructor;return i!==r&&"function"==typeof i&&(n=i.prototype)!==r.prototype&&a(n)&&o&&o(e,n),e}},"5f1b":function(e,t,r){"use strict";var a=r("23c6"),o=RegExp.prototype.exec;e.exports=function(e,t){var r=e.exec;if("function"===typeof r){var n=r.call(e,t);if("object"!==typeof n)throw new TypeError("RegExp exec method returned something other than an Object or null");return n}if("RegExp"!==a(e))throw new TypeError("RegExp#exec called on incompatible receiver");return o.call(e,t)}},"67bb":function(e,t,r){e.exports=r("f921")},"6b54":function(e,t,r){"use strict";r("3846");var a=r("cb7c"),o=r("0bfb"),n=r("9e1e"),i="toString",s=/./[i],l=function(e){r("2aba")(RegExp.prototype,i,e,!0)};r("79e5")(function(){return"/a/b"!=s.call({source:"a",flags:"b"})})?l(function(){var e=a(this);return"/".concat(e.source,"/","flags"in e?e.flags:!n&&e instanceof RegExp?o.call(e):void 0)}):s.name!=i&&l(function(){return s.call(this)})},7514:function(e,t,r){"use strict";var a=r("5ca1"),o=r("0a49")(5),n="find",i=!0;n in[]&&Array(1)[n](function(){i=!1}),a(a.P+a.F*i,"Array",{find:function(e){return o(this,e,arguments.length>1?arguments[1]:void 0)}}),r("9c6c")(n)},"8b97":function(e,t,r){var a=r("d3f4"),o=r("cb7c"),n=function(e,t){if(o(e),!a(t)&&null!==t)throw TypeError(t+": can't set as prototype!")};e.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(e,t,a){try{a=r("9b43")(Function.call,r("11e9").f(Object.prototype,"__proto__").set,2),a(e,[]),t=!(e instanceof Array)}catch(o){t=!0}return function(e,r){return n(e,r),t?e.__proto__=r:a(e,r),e}}({},!1):void 0),check:n}},9093:function(e,t,r){var a=r("ce10"),o=r("e11e").concat("length","prototype");t.f=Object.getOwnPropertyNames||function(e){return a(e,o)}},a481:function(e,t,r){"use strict";var a=r("cb7c"),o=r("4bf8"),n=r("9def"),i=r("4588"),s=r("0390"),l=r("5f1b"),c=Math.max,u=Math.min,d=Math.floor,f=/\$([$&`']|\d\d?|<[^>]*>)/g,p=/\$([$&`']|\d\d?)/g,v=function(e){return void 0===e?e:String(e)};r("214f")("replace",2,function(e,t,r,m){return[function(a,o){var n=e(this),i=void 0==a?void 0:a[t];return void 0!==i?i.call(a,n,o):r.call(String(n),a,o)},function(e,t){var o=m(r,e,this,t);if(o.done)return o.value;var d=a(e),f=String(this),p="function"===typeof t;p||(t=String(t));var h=d.global;if(h){var g=d.unicode;d.lastIndex=0}var y=[];while(1){var S=l(d,f);if(null===S)break;if(y.push(S),!h)break;var _=String(S[0]);""===_&&(d.lastIndex=s(f,n(d.lastIndex),g))}for(var j="",x=0,D=0;D<y.length;D++){S=y[D];for(var w=String(S[0]),I=c(u(i(S.index),f.length),0),O=[],U=1;U<S.length;U++)O.push(v(S[U]));var k=S.groups;if(p){var F=[w].concat(O,I,f);void 0!==k&&F.push(k);var P=String(t.apply(void 0,F))}else P=b(w,f,I,O,k,t);I>=x&&(j+=f.slice(x,I)+P,x=I+w.length)}return j+f.slice(x)}];function b(e,t,a,n,i,s){var l=a+e.length,c=n.length,u=p;return void 0!==i&&(i=o(i),u=f),r.call(s,u,function(r,o){var s;switch(o.charAt(0)){case"$":return"$";case"&":return e;case"`":return t.slice(0,a);case"'":return t.slice(l);case"<":s=i[o.slice(1,-1)];break;default:var u=+o;if(0===u)return r;if(u>c){var f=d(u/10);return 0===f?r:f<=c?void 0===n[f-1]?o.charAt(1):n[f-1]+o.charAt(1):r}s=n[u-1]}return void 0===s?"":s})}})},aa22:function(e,t,r){"use strict";r.d(t,"c",function(){return o}),r.d(t,"d",function(){return n}),r.d(t,"a",function(){return i}),r.d(t,"b",function(){return s}),r.d(t,"e",function(){return l});var a=r("b775");function o(){return Object(a["a"])({url:"/server/getList",method:"get",params:{}})}function n(){return Object(a["a"])({url:"/server/getOption",method:"get"})}function i(e){return Object(a["a"])({url:"/server/add",method:"post",data:e})}function s(e){return Object(a["a"])({url:"/server/edit",method:"post",data:e})}function l(e){return Object(a["a"])({url:"/server/remove",method:"post",data:{id:e}})}},aa77:function(e,t,r){var a=r("5ca1"),o=r("be13"),n=r("79e5"),i=r("fdef"),s="["+i+"]",l="​",c=RegExp("^"+s+s+"*"),u=RegExp(s+s+"*$"),d=function(e,t,r){var o={},s=n(function(){return!!i[e]()||l[e]()!=l}),c=o[e]=s?t(f):i[e];r&&(o[r]=c),a(a.P+a.F*s,"String",o)},f=d.trim=function(e,t){return e=String(o(e)),1&t&&(e=e.replace(c,"")),2&t&&(e=e.replace(u,"")),e};e.exports=d},b0c5:function(e,t,r){"use strict";var a=r("520a");r("5ca1")({target:"RegExp",proto:!0,forced:a!==/./.exec},{exec:a})},c0f9:function(e,t,r){"use strict";r.r(t);var a=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("el-row",{staticClass:"app-container"},[r("el-row",{staticClass:"app-bar",attrs:{type:"flex",justify:"end"}},[r("el-button",{attrs:{type:"primary",icon:"el-icon-plus"},on:{click:e.handleAdd}},[e._v("添加")])],1),e._v(" "),r("el-table",{staticStyle:{width:"100%"},attrs:{border:"",stripe:"","highlight-current-row":"",data:e.tableData}},[r("el-table-column",{attrs:{prop:"name",label:"项目名称"}}),e._v(" "),r("el-table-column",{attrs:{prop:"url",label:"项目地址"}}),e._v(" "),r("el-table-column",{attrs:{prop:"path",label:"部署路径"}}),e._v(" "),r("el-table-column",{attrs:{prop:"group",label:"分组",width:"100"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v("\n        "+e._s(e.findGroupName(t.row.projectGroupId))+"\n      ")]}}])}),e._v(" "),r("el-table-column",{attrs:{prop:"environment",width:"160",label:"环境"}}),e._v(" "),r("el-table-column",{attrs:{prop:"branch",width:"160",label:"分支"}}),e._v(" "),r("el-table-column",{attrs:{prop:"operation",label:"操作",width:"350"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("el-button",{attrs:{type:"primary"},on:{click:function(r){return e.handleEdit(t.row)}}},[e._v("编辑")]),e._v(" "),r("el-button",{attrs:{type:"success"},on:{click:function(r){return e.handleServer(t.row)}}},[e._v("服务器管理")]),e._v(" "),r("el-button",{attrs:{type:"warning"},on:{click:function(r){return e.handleUser(t.row)}}},[e._v("成员管理")]),e._v(" "),r("el-button",{attrs:{type:"danger"},on:{click:function(r){return e.handleRemove(t.row)}}},[e._v("删除")])]}}])})],1),e._v(" "),r("el-dialog",{attrs:{title:"项目设置",visible:e.dialogVisible},on:{"update:visible":function(t){e.dialogVisible=t}}},[r("el-form",{ref:"form",attrs:{rules:e.formRules,model:e.formData,"label-width":"120px"}},[r("el-form-item",{attrs:{label:"项目名称",prop:"name"}},[r("el-input",{attrs:{autocomplete:"off"},model:{value:e.formData.name,callback:function(t){e.$set(e.formData,"name",t)},expression:"formData.name"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"项目地址",prop:"url"}},[r("el-input",{attrs:{autocomplete:"off"},model:{value:e.formData.url,callback:function(t){e.$set(e.formData,"url",t)},expression:"formData.url"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"部署路径",prop:"path"}},[r("el-input",{attrs:{autocomplete:"off"},model:{value:e.formData.path,callback:function(t){e.$set(e.formData,"path",t)},expression:"formData.path"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"环境",prop:"environment"}},[r("el-select",{staticStyle:{width:"100%"},attrs:{placeholder:"选择环境"},model:{value:e.formData.environment,callback:function(t){e.$set(e.formData,"environment",t)},expression:"formData.environment"}},[r("el-option",{attrs:{label:"生产环境",value:"生产环境"}}),e._v(" "),r("el-option",{attrs:{label:"测试环境",value:"测试环境"}}),e._v(" "),r("el-option",{attrs:{label:"开发环境",value:"开发环境"}})],1)],1),e._v(" "),r("el-form-item",{attrs:{label:"分支",prop:"branch"}},[r("el-input",{attrs:{autocomplete:"off"},model:{value:e.formData.branch,callback:function(t){e.$set(e.formData,"branch",t)},expression:"formData.branch"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"rsync选项",prop:"rsyncOption"}},[r("el-input",{attrs:{autocomplete:"off",placeholder:"-rtv --exclude .git --delete-after"},model:{value:e.formData.rsyncOption,callback:function(t){e.$set(e.formData,"rsyncOption",t)},expression:"formData.rsyncOption"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"拉取后运行脚本",prop:"afterPullScrpit"}},[r("el-input",{attrs:{rows:3,type:"textarea",autocomplete:"off",placeholder:"多个脚本用回车分割"},model:{value:e.formData.afterPullScript,callback:function(t){e.$set(e.formData,"afterPullScript",t)},expression:"formData.afterPullScript"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"部署后运行脚本",prop:"afterDeployScrpit"}},[r("el-input",{attrs:{rows:3,type:"textarea",autocomplete:"off",placeholder:"多个脚本用回车分割"},model:{value:e.formData.afterDeployScript,callback:function(t){e.$set(e.formData,"afterDeployScript",t)},expression:"formData.afterDeployScript"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"绑定分组",prop:"projectGroupId"}},[r("el-select",{staticStyle:{width:"100%"},attrs:{placeholder:"选择分组"},model:{value:e.formData.projectGroupId,callback:function(t){e.$set(e.formData,"projectGroupId",t)},expression:"formData.projectGroupId"}},[r("el-option",{attrs:{label:"默认",value:0}}),e._v(" "),e._l(e.projectGroupOption,function(e,t){return r("el-option",{key:t,attrs:{label:e.name,value:e.id}})})],2)],1),e._v(" "),r("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.formProps.showServers,expression:"formProps.showServers"}],attrs:{label:"绑定服务器",prop:"serverIds"}},[r("el-select",{staticStyle:{width:"100%"},attrs:{multiple:"",placeholder:"选择服务器，可多选"},model:{value:e.formData.serverIds,callback:function(t){e.$set(e.formData,"serverIds",t)},expression:"formData.serverIds"}},e._l(e.serverOption,function(e,t){return r("el-option",{key:t,attrs:{label:e.name,value:e.id}})}),1)],1),e._v(" "),r("el-form-item",{directives:[{name:"show",rawName:"v-show",value:e.formProps.showUsers,expression:"formProps.showUsers"}],attrs:{label:"绑定组员",prop:"userIds"}},[r("el-select",{staticStyle:{width:"100%"},attrs:{multiple:"",placeholder:"选择组员，可多选"},model:{value:e.formData.userIds,callback:function(t){e.$set(e.formData,"userIds",t)},expression:"formData.userIds"}},e._l(e.userOption,function(e,t){return r("el-option",{key:t,attrs:{label:e.name,value:e.id}})}),1)],1)],1),e._v(" "),r("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[r("el-button",{on:{click:function(t){e.dialogVisible=!1}}},[e._v("取 消")]),e._v(" "),r("el-button",{attrs:{disabled:e.formProps.disabled,type:"primary"},on:{click:e.submit}},[e._v("确 定")])],1)],1),e._v(" "),r("el-dialog",{attrs:{title:"服务器管理",visible:e.dialogServerVisible},on:{"update:visible":function(t){e.dialogServerVisible=t}}},[r("el-row",{staticClass:"app-bar",attrs:{type:"flex",justify:"end"}},[r("el-button",{attrs:{type:"primary",icon:"el-icon-plus"},on:{click:e.handleAddServer}},[e._v("添加")])],1),e._v(" "),r("el-table",{staticStyle:{width:"100%"},attrs:{border:"",stripe:"","highlight-current-row":"",data:e.tableServerData}},[r("el-table-column",{attrs:{prop:"serverId",label:"服务器ID"}}),e._v(" "),r("el-table-column",{attrs:{prop:"serverName",label:"服务器名称"}}),e._v(" "),r("el-table-column",{attrs:{prop:"createTime",width:"160",label:"绑定时间"}}),e._v(" "),r("el-table-column",{attrs:{prop:"updateTime",width:"160",label:"更新时间"}}),e._v(" "),r("el-table-column",{attrs:{prop:"operation",label:"操作",width:"75"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("el-button",{attrs:{type:"danger"},on:{click:function(r){return e.removeProjectServer(t.row)}}},[e._v("删除")])]}}])})],1),e._v(" "),r("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[r("el-button",{on:{click:function(t){e.dialogServerVisible=!1}}},[e._v("取 消")])],1)],1),e._v(" "),r("el-dialog",{attrs:{title:"成员管理",visible:e.dialogUserVisible},on:{"update:visible":function(t){e.dialogUserVisible=t}}},[r("el-row",{staticClass:"app-bar",attrs:{type:"flex",justify:"end"}},[r("el-button",{attrs:{type:"primary",icon:"el-icon-plus"},on:{click:e.handleAddUser}},[e._v("添加")])],1),e._v(" "),r("el-table",{staticStyle:{width:"100%"},attrs:{border:"",stripe:"","highlight-current-row":"",data:e.tableUserData}},[r("el-table-column",{attrs:{prop:"userId",label:"用户ID"}}),e._v(" "),r("el-table-column",{attrs:{prop:"userName",label:"用户名称"}}),e._v(" "),r("el-table-column",{attrs:{prop:"createTime",width:"160",label:"绑定时间"}}),e._v(" "),r("el-table-column",{attrs:{prop:"updateTime",width:"160",label:"更新时间"}}),e._v(" "),r("el-table-column",{attrs:{prop:"operation",label:"操作",width:"75"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("el-button",{attrs:{type:"danger"},on:{click:function(r){return e.removeProjectUser(t.row)}}},[e._v("删除")])]}}])})],1),e._v(" "),r("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[r("el-button",{on:{click:function(t){e.dialogUserVisible=!1}}},[e._v("取 消")])],1)],1),e._v(" "),r("el-dialog",{attrs:{title:"添加服务器",visible:e.dialogAddServerVisible},on:{"update:visible":function(t){e.dialogAddServerVisible=t}}},[r("el-form",{ref:"addServerForm",attrs:{rules:e.addServerFormRules,model:e.addServerFormData}},[r("el-form-item",{attrs:{label:"绑定服务器","label-width":"120px",prop:"serverIds"}},[r("el-select",{attrs:{multiple:"",placeholder:"选择服务器，可多选"},model:{value:e.addServerFormData.serverIds,callback:function(t){e.$set(e.addServerFormData,"serverIds",t)},expression:"addServerFormData.serverIds"}},e._l(e.serverOption,function(e,t){return r("el-option",{key:t,attrs:{label:e.name,value:e.id}})}),1)],1)],1),e._v(" "),r("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[r("el-button",{on:{click:function(t){e.dialogAddServerVisible=!1}}},[e._v("取 消")]),e._v(" "),r("el-button",{attrs:{disabled:e.addServerFormProps.disabled,type:"primary"},on:{click:e.addServer}},[e._v("确 定")])],1)],1),e._v(" "),r("el-dialog",{attrs:{title:"添加用户",visible:e.dialogAddUserVisible},on:{"update:visible":function(t){e.dialogAddUserVisible=t}}},[r("el-form",{ref:"addUserForm",attrs:{rules:e.addUserFormRules,model:e.addUserFormData}},[r("el-form-item",{attrs:{label:"绑定组员","label-width":"120px",prop:"userIds"}},[r("el-select",{attrs:{multiple:"",placeholder:"选择组员，可多选"},model:{value:e.addUserFormData.userIds,callback:function(t){e.$set(e.addUserFormData,"userIds",t)},expression:"addUserFormData.userIds"}},e._l(e.userOption,function(e,t){return r("el-option",{key:t,attrs:{label:e.name,value:e.id}})}),1)],1)],1),e._v(" "),r("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[r("el-button",{on:{click:function(t){e.dialogAddUserVisible=!1}}},[e._v("取 消")]),e._v(" "),r("el-button",{attrs:{disabled:e.addUserFormProps.disabled,type:"primary"},on:{click:e.addUser}},[e._v("确 定")])],1)],1)],1)},o=[],n=(r("7f7f"),r("7514"),r("ac6a"),r("c24f")),i=r("aa22"),s=r("4c08"),l=r("b775");function c(){return Object(l["a"])({url:"/project/getList",method:"get",params:{}})}function u(e){return Object(l["a"])({url:"/project/getBindServerList",method:"get",params:{id:e}})}function d(e){return Object(l["a"])({url:"/project/getBindUserList",method:"get",params:{id:e}})}function f(e){return Object(l["a"])({url:"/project/create",method:"post",data:{id:e}})}function p(e){return Object(l["a"])({url:"/project/add",method:"post",data:e})}function v(e){return Object(l["a"])({url:"/project/edit",method:"post",data:e})}function m(e){return Object(l["a"])({url:"/project/remove",method:"post",data:{id:e}})}function b(e){return Object(l["a"])({url:"/project/addServer",method:"post",data:e})}function h(e){return Object(l["a"])({url:"/project/addUser",method:"post",data:e})}function g(e){return Object(l["a"])({url:"/project/removeProjectUser",method:"post",data:{projectUserId:e}})}function y(e){return Object(l["a"])({url:"/project/removeProjectServer",method:"post",data:{projectServerId:e}})}var S=r("ed08"),_={data:function(){return{dialogVisible:!1,dialogServerVisible:!1,dialogUserVisible:!1,dialogAddServerVisible:!1,dialogAddUserVisible:!1,serverOption:[],userOption:[],projectGroupOption:[],tableData:[],tableServerData:[],tableUserData:[],formProps:{disabled:!1,showServers:!0,showUsers:!0},tempFormData:{},formData:{id:0,projectGroupId:0,name:"",url:"",path:"",afterPullScript:"",afterDeployScript:"",environment:"生产环境",branch:"master",rsyncOption:"-rtv --exclude .git --delete-after",serverIds:[],userIds:[]},formRules:{name:[{required:!0,message:"请输入项目名称",trigger:["blur"]}],url:[{required:!0,message:"请输入项目地址",trigger:["blur"]}],path:[{required:!0,message:"请输入部署路径",trigger:["blur"]}],environment:[{required:!0,message:"请选择环境",trigger:["blur"]}],branch:[{required:!0,message:"请输入分支名称",trigger:["blur"]}],serverIds:[{type:"array",message:"请选择服务器",trigger:"change"}],userIds:[{type:"array",message:"请选择组员",trigger:"change"}]},addServerFormProps:{disabled:!1},addServerFormData:{projectId:0,serverIds:[]},addServerFormRules:{serverIds:[{type:"array",required:!0,message:"请选择服务器",trigger:"change"}]},addUserFormProps:{disabled:!1},addUserFormData:{projectId:0,userIds:[]},addUserFormRules:{userIds:[{type:"array",required:!0,message:"请选择用户",trigger:"change"}]}}},created:function(){this.storeFormData(),this.get()},methods:{handleAdd:function(){this.restoreFormData(),this.formProps.showServers=this.formProps.showUsers=!0,this.dialogVisible=!0},handleEdit:function(e){this.formData=Object.assign({},e),this.formData.serverIds=[],this.formData.userIds=[],this.formProps.showServers=this.formProps.showUsers=!1,this.dialogVisible=!0},handleRemove:function(e){var t=this;this.$confirm("此操作将删除该项目, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(function(){m(e.id).then(function(e){t.$message({message:e.message,type:"success",duration:5e3}),t.getProjectList()})}).catch(function(){t.$message({type:"info",message:"已取消删除"})})},handleServer:function(e){this.getBindServerList(e.id),this.addServerFormData.projectId=e.id,this.dialogServerVisible=!0},handleGroup:function(e){this.getBindServerList(e.id),this.addServerFormData.projectId=e.id,this.dialogServerVisible=!0},handleUser:function(e){this.getBindUserList(e.id),this.addUserFormData.projectId=e.id,this.dialogUserVisible=!0},handleAddServer:function(){this.dialogAddServerVisible=!0},handleAddUser:function(){this.dialogAddUserVisible=!0},submit:function(){var e=this;this.$refs.form.validate(function(t){if(!t)return!1;0===e.formData.id?e.add():e.edit()})},add:function(){var e=this;this.formProps.disabled=!0,p(this.formData).then(function(t){e.dialogVisible=!1,e.$message({message:t.message,type:"success",duration:5e3}),e.getProjectList()}).finally(function(){e.formProps.disabled=!1})},edit:function(){var e=this;this.formProps.disabled=!0,v(this.formData).then(function(t){e.dialogVisible=!1,e.$message({message:t.message,type:"success",duration:5e3}),e.getProjectList()}).finally(function(){e.formProps.disabled=!1})},create:function(e){var t=this;f(e).then(function(e){t.$message({message:e.message,type:"success",duration:5e3})})},addServer:function(){var e=this;this.$refs.addServerForm.validate(function(t){if(!t)return!1;e.addServerFormProps.disabled=!0,b(e.addServerFormData).then(function(t){e.dialogAddServerVisible=!1,e.$message({message:t.message,type:"success",duration:5e3}),e.getBindServerList(e.addServerFormData.projectId)}).finally(function(){e.addServerFormProps.disabled=!1})})},addUser:function(){var e=this;this.$refs.addUserForm.validate(function(t){if(!t)return!1;e.addUserFormProps.disabled=!0,h(e.addUserFormData).then(function(t){e.dialogAddUserVisible=!1,e.$message({message:t.message,type:"success",duration:5e3}),e.getBindUserList(e.addUserFormData.projectId)}).finally(function(){e.addUserFormProps.disabled=!1})})},removeProjectServer:function(e){var t=this;this.$confirm("此操作将永久删除该服务器的绑定关系, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(function(){y(e.id).then(function(r){t.$message({message:r.message,type:"success",duration:5e3}),t.getBindServerList(e.projectId)})}).catch(function(){t.$message({type:"info",message:"已取消删除"})})},removeProjectUser:function(e){var t=this;this.$confirm("此操作将永久删除该用户的绑定关系, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(function(){g(e.id).then(function(r){t.$message({message:r.message,type:"success",duration:5e3}),t.getBindUserList(e.projectId)})}).catch(function(){t.$message({type:"info",message:"已取消删除"})})},get:function(){var e=this;this.getProjectList(),Object(i["d"])().then(function(t){e.serverOption=t.data.serverList||[]}),Object(n["f"])().then(function(t){e.userOption=t.data.userList||[]}),Object(s["d"])().then(function(t){e.projectGroupOption=t.data.projectGroupList||[]})},getProjectList:function(){var e=this;c().then(function(t){var r=t.data.projectList;r.forEach(function(e){e.createTime=Object(S["a"])(e.createTime),e.updateTime=Object(S["a"])(e.updateTime)}),e.tableData=r}).catch(function(){})},getBindServerList:function(e){var t=this;u(e).then(function(e){t.tableServerData=e.data.projectServerMap||[],t.tableServerData.forEach(function(e){e.createTime=Object(S["a"])(e.createTime),e.updateTime=Object(S["a"])(e.updateTime)})})},getBindUserList:function(e){var t=this;d(e).then(function(e){t.tableUserData=e.data.projectUserMap||[],t.tableUserData.forEach(function(e){e.createTime=Object(S["a"])(e.createTime),e.updateTime=Object(S["a"])(e.updateTime)})})},findGroupName:function(e){var t=this.projectGroupOption.find(function(t){return t.id===e});return t?t["name"]:"默认"},storeFormData:function(){this.tempFormData=JSON.parse(JSON.stringify(this.formData))},restoreFormData:function(){this.formData=JSON.parse(JSON.stringify(this.tempFormData))}}},j=_,x=r("2877"),D=Object(x["a"])(j,a,o,!1,null,null,null);t["default"]=D.exports},c5f6:function(e,t,r){"use strict";var a=r("7726"),o=r("69a8"),n=r("2d95"),i=r("5dbc"),s=r("6a99"),l=r("79e5"),c=r("9093").f,u=r("11e9").f,d=r("86cc").f,f=r("aa77").trim,p="Number",v=a[p],m=v,b=v.prototype,h=n(r("2aeb")(b))==p,g="trim"in String.prototype,y=function(e){var t=s(e,!1);if("string"==typeof t&&t.length>2){t=g?t.trim():f(t,3);var r,a,o,n=t.charCodeAt(0);if(43===n||45===n){if(r=t.charCodeAt(2),88===r||120===r)return NaN}else if(48===n){switch(t.charCodeAt(1)){case 66:case 98:a=2,o=49;break;case 79:case 111:a=8,o=55;break;default:return+t}for(var i,l=t.slice(2),c=0,u=l.length;c<u;c++)if(i=l.charCodeAt(c),i<48||i>o)return NaN;return parseInt(l,a)}}return+t};if(!v(" 0o1")||!v("0b1")||v("+0x1")){v=function(e){var t=arguments.length<1?0:e,r=this;return r instanceof v&&(h?l(function(){b.valueOf.call(r)}):n(r)!=p)?i(new m(y(t)),r,v):y(t)};for(var S,_=r("9e1e")?c(m):"MAX_VALUE,MIN_VALUE,NaN,NEGATIVE_INFINITY,POSITIVE_INFINITY,EPSILON,isFinite,isInteger,isNaN,isSafeInteger,MAX_SAFE_INTEGER,MIN_SAFE_INTEGER,parseFloat,parseInt,isInteger".split(","),j=0;_.length>j;j++)o(m,S=_[j])&&!o(v,S)&&d(v,S,u(m,S));v.prototype=b,b.constructor=v,r("2aba")(a,p,v)}},cd1c:function(e,t,r){var a=r("e853");e.exports=function(e,t){return new(a(e))(t)}},e853:function(e,t,r){var a=r("d3f4"),o=r("1169"),n=r("2b4c")("species");e.exports=function(e){var t;return o(e)&&(t=e.constructor,"function"!=typeof t||t!==Array&&!o(t.prototype)||(t=void 0),a(t)&&(t=t[n],null===t&&(t=void 0))),void 0===t?Array:t}},ed08:function(e,t,r){"use strict";r("ac6a"),r("c5f6"),r("28a5"),r("a481"),r("6b54");var a=r("5d58"),o=r.n(a),n=r("67bb"),i=r.n(n);function s(e){return s="function"===typeof i.a&&"symbol"===typeof o.a?function(e){return typeof e}:function(e){return e&&"function"===typeof i.a&&e.constructor===i.a&&e!==i.a.prototype?"symbol":typeof e},s(e)}function l(e){return l="function"===typeof i.a&&"symbol"===s(o.a)?function(e){return s(e)}:function(e){return e&&"function"===typeof i.a&&e.constructor===i.a&&e!==i.a.prototype?"symbol":s(e)},l(e)}function c(e,t){if(0===arguments.length)return null;var r,a=t||"{y}-{m}-{d} {h}:{i}:{s}";"object"===l(e)?r=e:("string"===typeof e&&/^[0-9]+$/.test(e)&&(e=parseInt(e)),"number"===typeof e&&10===e.toString().length&&(e*=1e3),r=new Date(e));var o={y:r.getFullYear(),m:r.getMonth()+1,d:r.getDate(),h:r.getHours(),i:r.getMinutes(),s:r.getSeconds(),a:r.getDay()},n=a.replace(/{(y|m|d|h|i|s|a)+}/g,function(e,t){var r=o[t];return"a"===t?["日","一","二","三","四","五","六"][r]:(e.length>0&&r<10&&(r="0"+r),r||0)});return n}r.d(t,"a",function(){return c})},fdef:function(e,t){e.exports="\t\n\v\f\r   ᠎             　\u2028\u2029\ufeff"}}]);