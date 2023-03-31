(() => { var n = { 351: () => { layui.define(["form", "layer", "laydate", "upload", "element", "base"], (function(n) { "use strict"; var e = layui.form,
                        i = void 0 === parent.layer ? layui.layer : top.layer,
                        t = layui.laydate,
                        a = layui.upload,
                        o = (layui.element, layui.base),
                        r = layui.$,
                        l = { edit: function(n, e = 0, i = 0, t = 0, a = [], r = null, s = !1) {
                               var c = e > 0 ? "修改" : "新增";
                                o.isEmpty(n) ? c += "内容" : c += n;
                                var u = e > 0 ? cUrl + "/edit?id=" + e : cUrl + "/add";
                                if (Array.isArray(a))
                                    for (var f in a) u += ( u.indexOf("?") != -1 ? "&" : "?" ) + a[f];

                                l.showWin(c, u, i, t, a, 2, [], (function(n, e) { r && r(n, e) }), s) }, detail: function(n, e, i = 0, t = 0, a = !1) { var o = cUrl + "/detail?id=" + e;
                                l.showWin(n + "详情", o, i, t, [], 2, [], null, a) }, cache: function(n) { var e = cUrl + "/cache";
                                l.ajaxPost(e, { id: n }, (function(n, e) {})) }, copy: function(n, e, i = 0, t = 0) { var a = cUrl + "/copy?id=" + e;
                                l.showWin(n + "复制", a, i, t) }, delete: function(n, e = null) { i.confirm("您确定要删除吗？删除后将无法恢复！", { icon: 3, skin: "layer-ext-moon", btn: ["确认", "取消"] }, (function(t) { var a = cUrl + "/delete/" + n;
                                    console.log(a), l.ajaxPost(a, {}, (function(n, a) { e && (i.close(t), e(n, a)) }), "正在删除。。。") })) }, batchFunc: function(n, e = null) { var t = n.url,
                                    a = n.title,
                                    o = (n.form, n.confirm || !1),
                                    r = n.show_tips || "处理中...",
                                    s = n.data || [],
                                    c = n.param || [],
                                    u = n.type || "POST"; if ("导出数据" != a && 0 == s.length) return i.msg("请选择数据", { icon: 5 }), !1; var f = []; for (var d in s) f.push(s[d].id); var m = f.join(","),
                                    y = {}; if (y.ids = m, Array.isArray(c))
                                    for (var d in c) { var p = c[d].split("=");
                                        y[p[0]] = p[1] } o ? i.confirm("您确定要【" + a + "】选中的数据吗？", { icon: 3, title: "提示信息" }, (function(n) { "POST" == u ? t.indexOf("/delete") >= 0 ? l.ajaxPost(t + "/" + m, {}, e, r) : l.ajaxPost(t, y, e, r) : l.ajaxGet(t + "/" + m, {}, e, r) })) : "POST" == u ? l.ajaxPost(t, y, e, r) : l.ajaxGet(t + "/" + m, {}, e, r) }, verify: function() { e.verify({ number: [/^[0-9]*$/, "请输入数字"], username: function(n, e) { return new RegExp("^[a-zA-Z0-9_一-龥\\s·]+$").test(n) ? /(^\_)|(\__)|(\_+$)/.test(n) ? title + "首尾不能出现下划线'_'" : /^\d+\d+\d$/.test(n) ? title + "不能全为数字" : void 0 : title + "不能含有特殊字符" }, pass: [/^[\S]{6,12}$/, "密码必须6到12位，且不能出现空格"] }) }, submitForm: function(n, e = null, i = null, t = !0) {
                                var a = [],
                                    s = [],
                                    c = n; if (r.each(c, (function(n, e) {
                                        if (console.log(n + ":" + e), /\[|\]|【|】/g.test(n)) {
                                            var i = n.match(/\[(.+?)\]/g);
                                            e = n.match("\\[(.+?)\\]")[1]; var t = n.replace(i, "");
                                            r.inArray(t, a) < 0 && a.push(t), s[t] || (s[t] = []), s[t].push(e) } })), console.log(c), console.log(a), console.log(s), r.each(a, (function(n, e) { var i = [];
                                        r.each(s[e], (function(n, t) { i.push(t), delete c[e + "[" + t + "]"] })), c[e] = i.join(",") })), null == e) { e = cUrl; var u = r("form").attr("action");console.log(n)
                                    o.isEmpty(u) ? ( null == n.id || 0 == n.id ? e += "/add" : n.id > 0 && (e += "/edit")) : e = u } console.log(c), l.ajaxPost(e, c, (function(n, e) { if (e) return t && setTimeout((function() { var n = parent.layer.getFrameIndex(window.name);
                                        parent.layer.close(n) }), 500), i && i(n, e), !1 })) }, searchForm: function(n, e, i = "tableList") { n.reload(i, { page: { curr: 1 }, where: e.field }) }, initDate: function(n, e = null) { if (Array.isArray(n))
                                    for (var i in n) { var a = n[i].split("|"); if (a[2]) var o = a[2].split(","); var r = {}; if (r.elem = "#" + a[0], r.type = a[1], r.theme = "molv", r.range = "true" === a[3] || a[3], r.calendar = !0, r.show = !1, r.position = "absolute", r.trigger = "click", r.btns = ["clear", "now", "confirm"], r.mark = { "0-06-25": "生日", "0-12-31": "跨年" }, r.ready = function(n) {}, r.change = function(n, e, i) {}, r.done = function(n, i, t) { e && e(n, i) }, o) { var l = o[0]; if (l) { var s = !isNaN(l);
                                                r.min = s ? parseInt(l) : l } var c = o[1]; if (c) { var u = !isNaN(c);
                                                r.max = u ? parseInt(c) : c } } t.render(r) } }, showWin: function(n, e, i = 0, t = 0, a = [], o = 2, l = [], s = null, c = !1) { var u = layui.layer.open({ title: n, type: o, area: [i + "px", t + "px"], content: e, shadeClose: c, shade: .4, skin: "layui-layer-admin", success: function(n, e) { if (Array.isArray(a))
                                            for (var i in a) { var t = a[i].split("=");
                                                layui.layer.getChildFrame("body", e).find("#" + t[0]).val(t[1]) } s && s(e, 1) }, end: function() { s(u, 2) } });
                                0 == i && (layui.layer.full(u), r(window).on("resize", (function() { layui.layer.full(u) }))) }, ajaxPost: function(n, e, t = null, a = "处理中,请稍后...") { var o = null;
                                r.ajax({ type: "POST", url: n, data: JSON.stringify(e), contentType: "application/json", dataType: "json", beforeSend: function() { o = i.msg(a, { icon: 16, shade: .01, time: 0 }) }, success: function(n) { if (0 != n.code) return i.close(o), i.msg(n.msg, { icon: 5 }), !1;
                                        i.msg(n.msg, { icon: 1, time: 500 }, (function() { i.close(o), t && t(n, !0) })) }, error: function() { i.close(o), i.msg("AJAX请求异常"), t && t(null, !1) } }) }, ajaxGet: function(n, e, t = null, a = "处理中,请稍后...") { var o = null;
                                r.ajax({ type: "GET", url: n, data: e, contentType: "application/json", dataType: "json", beforeSend: function() { o = i.msg(a, { icon: 16, shade: .01, time: 0 }) }, success: function(n) { if (0 != n.code) return i.msg(n.msg, { icon: 5 }), !1;
                                        i.msg(n.msg, { icon: 1, time: 500 }, (function() { i.close(o), t && t(n, !0) })) }, error: function() { i.msg("AJAX请求异常"), t && t(null, !1) } }) }, formSwitch: function(n, i = "", t = null) { e.on("switch(" + n + ")", (function(e) { var a = this.checked ? "1" : "2";
                                    o.isEmpty(i) && (i = cUrl + "/set" + n.substring(0, 1).toUpperCase() + n.substring(1)); var r = {};
                                    r.id = this.value, r[n] = a; var s = JSON.stringify(r);
                                    JSON.parse(s), l.ajaxPost(i, r, (function(n, e) { t && t(n, e) })) })) }, uploadFile: function(n, e = null, t = "", r = "xls|xlsx", l = 10240, s = {}) { o.isEmpty(t) && (t = cUrl + "/uploadFile"), a.render({ elem: "#" + n, url: t, auto: !1, exts: r, accept: "file", size: l, method: "post", data: s, before: function(n) { i.msg("上传并处理中。。。", { icon: 16, shade: .01, time: 0 }) }, done: function(n) { return i.closeAll(), 0 == n.code ? i.alert(n.msg, { title: "上传反馈", skin: "layui-layer-molv", closeBtn: 1, anim: 0, btn: ["确定", "取消"], icon: 6, yes: function() { e && e(n, !0) }, btn2: function() {} }) : i.msg(n.msg, { icon: 5 }), !1 }, error: function() { return i.msg("数据请求异常") } }) } };
                    n("common", l) })) } },
        e = {};! function i(t) { var a = e[t]; if (void 0 !== a) return a.exports; var o = e[t] = { exports: {} }; return n[t](o, o.exports, i), o.exports }(351) })();